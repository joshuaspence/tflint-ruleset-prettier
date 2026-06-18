package rules

import (
	"fmt"
	"path/filepath"

	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type VariableMustBeInVariablesFileRule struct {
	tflint.DefaultRule
}

func NewVariableMustBeInVariablesFileRule() *VariableMustBeInVariablesFileRule {
	return &VariableMustBeInVariablesFileRule{}
}

func (r *VariableMustBeInVariablesFileRule) Name() string {
  return "variables_must_be_in_variables_file"
}

func (r *HardcodedIDsRule) Enabled() bool {
  return true
}

func (r *HardcodedIDsRule) Severity() tflint.Severity {
  return tflint.NOTICE
}

func (r *VariableMustBeInVariablesFileRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *VariableMustBeInVariablesFileRule) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{Type: "variable", LabelNames: []string{"name"}},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if len(block.Labels) >= 1 {
			varName := block.Labels[0]
			filename := filepath.Base(block.DefRange.Filename)

			if filename != "variables.tf" {
				err := EmitIssue(runner, r, fmt.Sprintf("Variable '%s' must be declared in variables.tf, not in %s", varName, filename), block.DefRange)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
