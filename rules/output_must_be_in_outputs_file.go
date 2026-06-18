package rules

import (
	"fmt"
	"path/filepath"

	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// OutputMustBeInOutputsFileRule ensures all output blocks are declared in outputs.tf.
type OutputMustBeInOutputsFileRule struct {
	BaseRule
}

func NewOutputMustBeInOutputsFileRule() *OutputMustBeInOutputsFileRule {
	return &OutputMustBeInOutputsFileRule{
		BaseRule: BaseRule{ruleName: "output_must_be_in_outputs_file"},
	}
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
			if err := EmitIssue(runner, r,
				fmt.Sprintf("Output %q is defined in %s. All outputs should be in outputs.tf for consistent file organization.", block.Labels[0], filename),
				block.DefRange,
			); err != nil {
				return err
			}
		}
	}
	return nil
}
