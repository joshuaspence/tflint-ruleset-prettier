package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosJammedCommentParser is a regex to detect jammed comments.
// https://regex101.com/r/5HRrLc/1
var eosJammedCommentParser = regexp.MustCompile(`^\s*(///*|##*|/\*\**)([^\s/#])`)

// eosCheckJammed checks if comments are jammed (no space after delimiter).
func eosCheckJammed(r *EosCommentsRule, config *eosCommentsRuleConfig, text string, runner tflint.Runner, token hclsyntax.Token, _ *hclsyntax.Token) {
	if config.Jammed == nil || !*config.Jammed {
		return
	}

	if eosJammedCommentParser.MatchString(text) {
		trimmed := strings.TrimSpace(text)
		rns := []rune(trimmed)
		snippet := trimmed
		if len(rns) > 5 {
			snippet = string(rns[:5])
		}
		message := fmt.Sprintf("Avoid jammed comment ('%s ...').", snippet)
		if err := runner.EmitIssue(r, message, token.Range); err != nil {
			logger.Error(err.Error())
		}
		logger.Debug(message)
	}
}
