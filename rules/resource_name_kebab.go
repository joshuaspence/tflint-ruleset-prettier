package rules

import (
	"fmt"
	"regexp"

	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var KebabRegex = regexp.MustCompile(`^[a-z0-9-]+$`)

type ResourceNameKebabRule struct {
	tflint.DefaultRule
}

func NewResourceNameKebabRule() *ResourceNameKebabRule {
	return &ResourceNameKebabRule{}
}

func (r *ResourceNameKebabRule) Name() string {
	return "resource_name_kebab"
}

func (r *ResourceNameKebabRule) Enabled() bool {
	return true
}

func (r *ResourceNameKebabRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *ResourceNameKebabRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *ResourceNameKebabRule) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{Type: "resource", LabelNames: []string{"type", "name"}, Body: &hclext.BodySchema{
				Attributes: []hclext.AttributeSchema{
					// TODO: Identify other common name attributes
					{Name: "name"},
					{Name: "name_prefix"},
				},
			}},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if len(block.Labels) >= 2 {
			// Check name attribute
			if nameAttr, exists := block.Body.Attributes["name"]; exists {
				if err := r.checkNameAttribute(runner, "name", nameAttr); err != nil {
					return err
				}
			}

			// Check name_prefix attribute
			if namePrefixAttr, exists := block.Body.Attributes["name_prefix"]; exists {
				if err := r.checkNameAttribute(runner, "name_prefix", namePrefixAttr); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *ResourceNameKebabRule) checkNameAttribute(runner tflint.Runner, attrName string, attr *hclext.Attribute) error {
	var nameValue string
	err := runner.EvaluateExpr(attr.Expr, &nameValue, nil)
	if err != nil {
		// Skip if we can't evaluate (variables, references, etc.)
		return nil
	}

	if !KebabRegex.MatchString(nameValue) {
		return runner.EmitIssue(r, fmt.Sprintf("Resource %s attribute '%s' must contain only lowercase letters, numbers, and dashes", attrName, nameValue), attr.Range)
	}

	return nil
}
