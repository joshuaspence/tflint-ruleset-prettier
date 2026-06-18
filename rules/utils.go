package rules

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

// isExpectedEvalError reports whether an error from runner.EvaluateExpr is an
// expected "this expression can't be statically evaluated to a string" case
// that rules should silently skip rather than propagate. This covers the SDK
// sentinel errors (unknown/null/sensitive/ephemeral values) as well as cty
// type-conversion failures, which surface as a "cannot convert" message when
// the expression evaluates to a non-string value.
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

// isFileInCurrentModule checks if the file belongs to the current module context.
func isFileInCurrentModule(filename string) bool {
	dir := filepath.Dir(filename)
	if dir == "." {
		return true
	}

	cwd, err := os.Getwd()
	if err != nil {
		return true
	}

	return strings.HasSuffix(cwd, dir)
}

// checkExpression is a generic helper to validate expressions and files before processing.
func checkExpression(e hcl.Expression, files map[string]*hcl.File, callback func(*hcl.File) hcl.Diagnostics) hcl.Diagnostics {
	filename := e.Range().Filename

	if !isFileInCurrentModule(filename) {
		return nil
	}

	file, ok := files[filename]
	if !ok {
		return nil
	}

	if len(file.Bytes) == 0 {
		return nil
	}

	return callback(file)
}
