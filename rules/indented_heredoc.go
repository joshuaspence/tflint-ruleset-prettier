package rules

import (
	"regexp"

	"github.com/joshuaspence/tflint-ruleset-prettier/project"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var heredocPattern = regexp.MustCompile(`<<(-?)([a-zA-Z0-9]+)\s*$`)

type IndentedHeredocRule struct {
	tflint.DefaultRule
}

func NewIndentedHeredocRule() *IndentedHeredocRule {
	return &IndentedHeredocRule{}
}

func (r *IndentedHeredocRule) Name() string {
	return "indented_heredoc"
}

func (r *IndentedHeredocRule) Enabled() bool {
	return true
}

func (r *IndentedHeredocRule) Severity() tflint.Severity {
	return tflint.NOTICE
}

func (r *IndentedHeredocRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IndentedHeredocRule) Check(runner tflint.Runner) error {
	return walkTokens(runner, r, func(runner tflint.Runner, rule *IndentedHeredocRule, token hclsyntax.Token) {
		checkHeredocToken(runner, rule, token)
	})
}

func checkHeredocToken(runner tflint.Runner, r *IndentedHeredocRule, token hclsyntax.Token) {
	if token.Type != hclsyntax.TokenOHeredoc {
		return
	}

	text := string(token.Bytes)
	matches := heredocPattern.FindStringSubmatch(text)
	if matches == nil {
		return
	}

	indentMarker := matches[1]

	if indentMarker == "" {
		if err := runner.EmitIssue(r, "Avoid standard heredoc (<<). Use indented (<<-) instead.", token.Range); err != nil {
			logger.Error(err.Error())
		}
	}
}
