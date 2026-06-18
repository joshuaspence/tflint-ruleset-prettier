package rules

import (
	"os"
	"path/filepath"
	"strings"

  "github.com/hashicorp/hcl/v2"
  "github.com/hashicorp/hcl/v2/hclsyntax"
  "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

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

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func walkTokens[T any](
  runner tflint.Runner,
  rule T,
  checkFunc func(tflint.Runner, T, hclsyntax.Token),
) error {
  path, err := runner.GetModulePath()
  if err != nil {
    return err
  }
  if !path.IsRoot() {
    return nil
  }

  files, err := runner.GetFiles()
  if err != nil {
    return err
  }

  for filename, file := range files {
    tokens, diags := hclsyntax.LexConfig(file.Bytes, filename, hcl.InitialPos)
    if diags.HasErrors() {
      return diags
    }

    for _, token := range tokens {
      checkFunc(runner, rule, token)
    }
  }

  return nil
}
