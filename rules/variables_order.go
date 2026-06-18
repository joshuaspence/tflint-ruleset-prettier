package rules

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type VariablesOrderRule struct {
	tflint.DefaultRule
}

type VariablesOrderRuleConfig struct {
	GroupRequired bool `hclext:"group_required"`
}

func NewVariablesOrderRule() *VariablesOrderRule {
	return &VariablesOrderRule{}
}

func (r *VariablesOrderRule) Name() string {
	return "variables_order"
}

func (r *VariablesOrderRule) Enabled() bool {
	return true
}

func (r *VariablesOrderRule) Severity() tflint.Severity {
	return tflint.NOTICE
}

func (r *VariablesOrderRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *VariablesOrderRule) Check(runner tflint.Runner) error {
	config := &VariablesOrderRuleConfig{}
	if err := runner.DecodeRuleConfig(r.Name(), config); err != nil {
		return err
	}

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		if subErr := r.checkVariablesOrder(runner, config.GroupRequired, file); subErr != nil {
			err = multierror.Append(err, subErr)
		}
	}
	return err
}

func (r *VariablesOrderRule) checkVariablesOrder(runner tflint.Runner, sortRequired bool, file *hcl.File) error {
	body, ok := file.Body.(*hclsyntax.Body)
	if !ok {
		logger.Debug("skip terraform_variables_order check since it's not a valid hcl file")
		return nil
	}
	blocks := body.Blocks

	var sortedVariableNames []string

	if sortRequired {
		requiredVars := r.getSortedVariableNames(blocks, true)
		optionalVars := r.getSortedVariableNames(blocks, false)
		sortedVariableNames = append(requiredVars, optionalVars...)
	} else {
		sortedVariableNames = r.getVariableNames(blocks)
		sort.Strings(sortedVariableNames)
	}

	variableNames := r.getVariableNames(blocks)
	if reflect.DeepEqual(variableNames, sortedVariableNames) {
		return nil
	}

	firstRange := r.firstVariableRange(blocks)
	sortedVariableHclTxts := r.sortedVariableCodeTxts(blocks, file, sortedVariableNames)
	sortedVariableString := strings.Join(sortedVariableHclTxts, "\n\n")
	sortedVariableHclBytes := hclwrite.Format([]byte(sortedVariableString))

	return runner.EmitIssueWithFix(
		r,
		fmt.Sprintf("Recommended variables order:\n%s", sortedVariableHclBytes),
		*firstRange,
		func(f tflint.Fixer) error {
			// We can't fix the file if it contains a mix of variables and other blocks
			if !r.isAllVariables(blocks) {
				logger.Debug("Fix is not supported for files with non-variable blocks")
				return tflint.ErrFixNotSupported
			}

			fullRange := body.Range()
			err := f.ReplaceText(fullRange, sortedVariableString+"\n")
			if err != nil {
				return err
			}
			return nil
		},
	)
}

func (r *VariablesOrderRule) sortedVariableCodeTxts(blocks hclsyntax.Blocks, file *hcl.File, sortedVariableNames []string) []string {
	variableHclTxts := r.variableCodeTxts(blocks, file)
	var sortedVariableHclTxts []string
	for _, name := range sortedVariableNames {
		sortedVariableHclTxts = append(sortedVariableHclTxts, variableHclTxts[name])
	}
	return sortedVariableHclTxts
}

func (r *VariablesOrderRule) variableCodeTxts(blocks hclsyntax.Blocks, file *hcl.File) map[string]string {
	variableHclTxts := make(map[string]string)
	r.forVariables(blocks, func(v *hclsyntax.Block) {
		name := v.Labels[0]
		variableHclTxts[name] = string(v.Range().SliceBytes(file.Bytes))
	})
	return variableHclTxts
}

func (r *VariablesOrderRule) firstVariableRange(blocks hclsyntax.Blocks) *hcl.Range {
	for _, block := range blocks {
		if block.Type == "variable" {
			rng := block.DefRange()
			return &rng
		}
	}

	return nil

}

func (r *VariablesOrderRule) getVariableNames(blocks hclsyntax.Blocks) []string {
	var variableNames []string
	r.forVariables(blocks, func(v *hclsyntax.Block) {
		variableNames = append(variableNames, v.Labels[0])
	})
	return variableNames
}

func (r *VariablesOrderRule) getSortedVariableNames(blocks hclsyntax.Blocks, required bool) []string {
	var variableNames []string

	r.forVariables(blocks, func(v *hclsyntax.Block) {
		if _, hasDefault := v.Body.Attributes["default"]; hasDefault != required {
			variableNames = append(variableNames, v.Labels[0])
		}
	})

	sort.Strings(variableNames)
	return variableNames
}

func (r *VariablesOrderRule) forVariables(blocks hclsyntax.Blocks, action func(v *hclsyntax.Block)) {
	for _, block := range blocks {
		if block.Type == "variable" {
			action(block)
		}
	}
}

// Checks if all blocks are variable blocks
func (r *VariablesOrderRule) isAllVariables(blocks hclsyntax.Blocks) bool {
	for _, block := range blocks {
		if block.Type != "variable" {
			return false
		}
	}

	return true
}
