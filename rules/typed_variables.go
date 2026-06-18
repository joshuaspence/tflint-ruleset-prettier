package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type TypedVariablesRule struct {
	tflint.DefaultRule
}

func NewTypedVariablesRule() *TypedVariablesRule {
	return &TypedVariablesRule{}
}

func (r *TypedVariablesRule) Name() string {
	return "typed_variables"
}

func (r *TypedVariablesRule) Enabled() bool {
	return true
}

func (r *TypedVariablesRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *TypedVariablesRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *TypedVariablesRule) Check(runner tflint.Runner) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		body, ok := file.Body.(*hclsyntax.Body)
		if !ok {
			continue
		}

		for _, block := range body.Blocks {
			if block.Type != "variable" || len(block.Labels) == 0 {
				continue
			}

			hasType := false
			for _, attr := range block.Body.Attributes {
				if attr.Name == "type" {
					hasType = true
					break
				}
			}

			if !hasType {
				if err := EmitIssue(runner, r,
					fmt.Sprintf("Variable %q is missing a type constraint.", block.Labels[0]),
					block.DefRange(),
				); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
