package aws

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/myerscode/tflint-ruleset-aws-meta/rules/awsmeta"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AwsServicePrincipalHardcodedRule struct {
	tflint.DefaultRule
}

func NewAwsServicePrincipalHardcodedRule() *AwsServicePrincipalHardcodedRule {
	return &AwsServicePrincipalHardcodedRule{}
}

func (r *AwsServicePrincipalHardcodedRule) Name() string {
	return "aws_service_principal_hardcoded"
}

func (r *AwsServicePrincipalHardcodedRule) Enabled() bool {
	return true
}

func (r *AwsServicePrincipalHardcodedRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *AwsServicePrincipalHardcodedRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *AwsServicePrincipalHardcodedRule) Check(runner tflint.Runner) error {
	dnsSuffixPattern := awsmeta.GetDNSSuffixPattern()

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	checked := make(map[string]bool)

	diags := runner.WalkExpressions(tflint.ExprWalkFunc(func(expr hcl.Expression) hcl.Diagnostics {
		exprKey := fmt.Sprintf("%s:%d:%d", expr.Range().Filename, expr.Range().Start.Line, expr.Range().Start.Column)
		if checked[exprKey] {
			return nil
		}
		checked[exprKey] = true

		// Pre-filter: check raw source for any known DNS suffix before making gRPC call
		exprRange := expr.Range()
		if file, ok := files[exprRange.Filename]; ok {
			src := file.Bytes
			if exprRange.Start.Byte < len(src) && exprRange.End.Byte <= len(src) {
				sourceText := string(src[exprRange.Start.Byte:exprRange.End.Byte])
				if !dnsSuffixPattern.MatchString(sourceText) {
					return nil
				}
			}
		}

		err := runner.EvaluateExpr(expr, func(value string) error {
			if matches := dnsSuffixPattern.FindStringSubmatch(value); len(matches) > 0 {
				serviceName := matches[1]

				if err := runner.EmitIssue(
					r,
					fmt.Sprintf("Hardcoded service principal '%s' found. Consider using data.aws_service_principal.%s.name for multi-partition compatibility", value, strings.ReplaceAll(serviceName, "-", "_")),
					expr.Range(),
				); err != nil {
					return err
				}
			}

			return nil
		}, nil)

		// This walks every expression in the module, so most will not evaluate to a string (function calls, tuples,
		// unknown references, etc.). Those evaluation errors are expected and ignored.
		_ = err

		return nil
	}))

	if diags.HasErrors() {
		return diags
	}

	return nil
}
