package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type MapTrailingCommaRule struct {
	tflint.DefaultRule
}

func NewMapTrailingCommaRule() *MapTrailingCommaRule {
	return &MapTrailingCommaRule{}
}

func (r *MapTrailingCommaRule) Name() string {
	return "map_trailing_comma"
}

func (r *MapTrailingCommaRule) Enabled() bool {
	return true
}

func (r *MapTrailingCommaRule) Severity() tflint.Severity {
	return tflint.NOTICE
}

func (r *MapTrailingCommaRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *MapTrailingCommaRule) Check(runner tflint.Runner) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	diags := runner.WalkExpressions(tflint.ExprWalkFunc(func(e hcl.Expression) hcl.Diagnostics {
		return checkExpression(e, files, func(file *hcl.File) hcl.Diagnostics {
			fileLength := len(file.Bytes)
			filename := e.Range().Filename

			expr, ok := e.(*hclsyntax.ObjectConsExpr)
			if !ok || len(expr.Items) == 0 {
				return nil
			}

			listRange := expr.Range()
			if listRange.Start.Line == listRange.End.Line {
				return nil
			}

			var itemsWithComma []int
			var itemsWithoutComma []int

			for i, item := range expr.Items {
				valRange := item.ValueExpr.Range()
				commaPos := valRange.End.Byte

				for commaPos < fileLength && isWhitespace(file.Bytes[commaPos]) {
					commaPos++
				}

				if commaPos < fileLength && file.Bytes[commaPos] == ',' {
					itemsWithComma = append(itemsWithComma, i)
				} else {
					itemsWithoutComma = append(itemsWithoutComma, i)
				}
			}

			for _, i := range itemsWithComma {
				// If the next item is on the same line, the comma is a separator and cannot be removed
				if i+1 < len(expr.Items) {
					currentEndLine := expr.Items[i].ValueExpr.Range().End.Line
					nextStartLine := expr.Items[i+1].KeyExpr.Range().Start.Line
					if currentEndLine == nextStartLine {
						continue
					}
				}

				item := expr.Items[i]
				startPos := item.ValueExpr.Range().End
				curr := startPos.Byte

				for curr < fileLength && isWhitespace(file.Bytes[curr]) {
					if file.Bytes[curr] == '\n' {
						startPos.Line++
						startPos.Column = 1
					} else {
						startPos.Column++
					}
					startPos.Byte++
					curr++
				}

				if curr < fileLength && file.Bytes[curr] == ',' {
					endPos := startPos
					endPos.Column++
					endPos.Byte++

					if err := runner.EmitIssueWithFix(
						r,
						"Map values should not have trailing commas",
						item.ValueExpr.Range(),
						func(f tflint.Fixer) error {
							return f.Remove(hcl.Range{
								Filename: filename,
								Start:    startPos,
								End:      endPos,
							})
						},
					); err != nil {
						return hcl.Diagnostics{
							{
								Severity: hcl.DiagError,
								Summary:  "failed to call EmitIssueWithFix()",
								Detail:   err.Error(),
							},
						}
					}
				}
			}

			return nil
		})
	}))

	if diags.HasErrors() {
		return diags
	}

	return nil
}
