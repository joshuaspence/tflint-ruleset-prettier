package aws

import (
	"errors"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// isExpectedEvalError reports whether an error from runner.EvaluateExpr is an expected "this expression can't be
// statically evaluated to a string" case that rules should silently skip rather than propagate. This covers the SDK
// sentinel errors (unknown/null/sensitive/ephemeral values) as well as cty type-conversion failures, which surface as
// a "cannot convert" message when the expression evaluates to a non-string value.
func isExpectedEvalError(err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, tflint.ErrUnknownValue) ||
		errors.Is(err, tflint.ErrNullValue) ||
		errors.Is(err, tflint.ErrSensitive) ||
		errors.Is(err, tflint.ErrEphemeral) {
		return true
	}

	return strings.Contains(err.Error(), "cannot convert")
}
