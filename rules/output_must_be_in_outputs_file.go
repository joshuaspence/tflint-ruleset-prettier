package rules

import (
	"fmt"
	"path/filepath"

	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type OutputMustBeInOutputsFileRule struct {
	tflint.DefaultRule
}

func NewOutputMustBeInOutputsFileRule() *OutputMustBeInOutputsFileRule {
	return &OutputMustBeInOutputsFileRule{}
}

func (r *OutputMustBeInOutputsFileRule) Name() string {
	return "output_must_be_in_outputs_file"
}

func (r *OutputMustBeInOutputsFileRule) Enabled() bool {
	return true
}

func (r *OutputMustBeInOutputsFileRule) Severity() tflint.Severity {
	return tflint.NOTICE
}

func (r *OutputMustBeInOutputsFileRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *OutputMustBeInOutputsFileRule) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "output",
				LabelNames: []string{"name"},
				Body:       &hclext.BodySchema{},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		filename := filepath.Base(block.DefRange.Filename)
		if filename != "outputs.tf" {
			if err := runner.EmitIssue(
				r,
				fmt.Sprintf("Output %q is defined in %s. All outputs should be in outputs.tf.", block.Labels[0], filename),
				block.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
