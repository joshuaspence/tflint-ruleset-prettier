package rules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AwsPolicyNoJsonencodeRule struct {
	BaseRule
}

func NewAwsPolicyNoJsonencodeRule() *AwsPolicyNoJsonencodeRule {
	return &AwsPolicyNoJsonencodeRule{
		BaseRule: BaseRule{ruleName: "aws_policy_no_jsonencode"},
	}
}

func (r *AwsPolicyNoJsonencodeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *AwsPolicyNoJsonencodeRule) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{Type: "resource", LabelNames: []string{"type", "name"}, Body: &hclext.BodySchema{
				Attributes: []hclext.AttributeSchema{{Name: "assume_role_policy"}, {Name: "policy"}, {Name: "bucket_policy"}},
			}},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if len(block.Labels) >= 2 {
			resourceType := block.Labels[0]

			// Only check AWS resources
			if !strings.HasPrefix(resourceType, "aws_") {
				continue
			}

			// Check specific policy attributes
			for name, attr := range block.Body.Attributes {
				if r.isPolicyAttribute(name) {
					if r.containsJsonencode(attr.Expr) {
						err := EmitIssue(runner, r, fmt.Sprintf("AWS resource '%s' attribute '%s' should reference an aws_iam_policy_document data source instead of using jsonencode()", resourceType, name), attr.Range)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func (r *AwsPolicyNoJsonencodeRule) isPolicyAttribute(name string) bool {
	return name == "policy" || name == "assume_role_policy" || name == "bucket_policy" || strings.HasSuffix(name, "_policy")
}

func (r *AwsPolicyNoJsonencodeRule) containsJsonencode(expr hcl.Expression) bool {
	switch e := expr.(type) {
	case *hclsyntax.FunctionCallExpr:
		return e.Name == "jsonencode"
	case *hclsyntax.ParenthesesExpr:
		return r.containsJsonencode(e.Expression)
	}
	return false
}
