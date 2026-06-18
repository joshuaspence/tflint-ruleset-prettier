package rules

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"
)

const eosOnlyDynamicGuardMessage = "Avoid using count for anything other than dynamic guarding (condition ? 1 : 0)."
const eosGuardMustReturn10Message = "Count guard must return 1 or 0."

// eosCheckCountGuard checks for proper count guard usage.
func eosCheckCountGuard(runner tflint.Runner, r *MetaRule, attr *hclsyntax.Attribute) {
	expr := attr.Expr
	condExpr, ok := expr.(*hclsyntax.ConditionalExpr)

	// We want to check if it is a conditional expression:
	// condition ? true_val : false_val
	if !ok {
		// Allow literal 0 or 1.
		if lit, ok := expr.(*hclsyntax.LiteralValueExpr); ok {
			val := eosGetLiteralValue(lit)
			if val == 0 || val == 1 {
				return
			}
		}

		r.emitIssue(runner, eosOnlyDynamicGuardMessage, attr.Range())
		return
	}

	// Check true/false results.
	if !eosIsValidGuardResult(condExpr.TrueResult) || !eosIsValidGuardResult(condExpr.FalseResult) {
		r.emitIssue(runner, eosGuardMustReturn10Message, attr.Range())
	}
}

// eosIsValidGuardResult checks if the expression is a valid guard result (0 or
// 1).
func eosIsValidGuardResult(expr hclsyntax.Expression) bool {
	val := eosGetLiteralValue(expr)
	return val == 0 || val == 1
}

// eosGetLiteralValue extracts the integer value from a literal expression.
func eosGetLiteralValue(expr hclsyntax.Expression) int {
	if lit, ok := expr.(*hclsyntax.LiteralValueExpr); ok {
		if lit.Val.Type() == cty.Number {
			f, _ := lit.Val.AsBigFloat().Float64()
			if f == 0 {
				return 0
			}
			if f == 1 {
				return 1
			}
		}
	}
	return -1
}
