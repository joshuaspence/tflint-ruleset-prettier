package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type ListTrailingCommaRule struct {
	tflint.DefaultRule
}

func NewListTrailingCommaRule() *ListTrailingCommaRule {
	return &ListTrailingCommaRule{}
}

func (r *ListTrailingCommaRule) Name() string {
	return "list_trailing_comma"
}

func (r *ListTrailingCommaRule) Enabled() bool {
	return true
}

func (r *ListTrailingCommaRule) Severity() tflint.Severity {
	return tflint.NOTICE
}

func (r *ListTrailingCommaRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *ListTrailingCommaRule) Check(runner tflint.Runner) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	diags := runner.WalkExpressions(tflint.ExprWalkFunc(func(e hcl.Expression) hcl.Diagnostics {
		return checkExpression(e, files, func(file *hcl.File) hcl.Diagnostics {
			fileLength := len(file.Bytes)

			list, ok := e.(*hclsyntax.TupleConsExpr)
			if !ok || len(list.Exprs) == 0 {
				return nil
			}

			listRange := list.Range()
			lastItem := list.Exprs[len(list.Exprs)-1]
			lastItemRange := lastItem.Range()

			if listRange.Start.Line == lastItemRange.Start.Line {
				return nil
			}

			// Check if there's already a trailing comma after the last item. We need to skip whitespace and newlines to 
      // handle heredoc cases.
			commaPos := lastItemRange.End.Byte

			// Skip whitespace and newlines after the last item to look for a comma
			for commaPos < fileLength && isWhitespace(file.Bytes[commaPos]) {
				commaPos++
			}

			if commaPos < fileLength && file.Bytes[commaPos] == ',' {
				// It already has a trailling comma
				return nil
			}

			insertText := ","
			// Check if the last item is a heredoc.
			// A heredoc is a TemplateExpr with a single LiteralValueExpr part.
			if template, ok := lastItem.(*hclsyntax.TemplateExpr); ok {
				if len(template.Parts) == 1 {
					if _, isLiteral := template.Parts[0].(*hclsyntax.LiteralValueExpr); isLiteral {
						// This is a strong indicator of a heredoc, especially if it spans multiple lines.
						if template.Range().Start.Line != template.Range().End.Line {
							insertText = "\n,"
						}
					}
				}
			}

			if err := runner.EmitIssueWithFix(
				r,
				"Last item in lists should always end with a trailing comma",
				listRange,
				func(f tflint.Fixer) error {
					return f.InsertTextAfter(lastItemRange, insertText)
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

			return nil
		})
	}))

	if diags.HasErrors() {
		return diags
	}

	return nil
}
