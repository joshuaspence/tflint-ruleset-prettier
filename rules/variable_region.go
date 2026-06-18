package rules

import (
	"fmt"

	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type VariableRegionRule struct {
	BaseRule
}

func NewVariableRegionRule() *VariableRegionRule {
	return &VariableRegionRule{
		BaseRule: BaseRule{ruleName: "variable_region"},
	}
}

func (r *VariableRegionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *VariableRegionRule) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{Type: "variable", LabelNames: []string{"name"}},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		label := block.Labels[0]
		if label == "region" {
			err := EmitIssue(runner, r, fmt.Sprintf("Variable '%s' is not allowed; use provider default region", label), block.DefRange)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
