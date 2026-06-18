package rules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosCheckLength checks if comments exceed the column limit.
func eosCheckLength(r *EosCommentsRule, config *eosCommentsRuleConfig, text string, runner tflint.Runner, token hclsyntax.Token, _ *hclsyntax.Token) {
	if config.Length == nil || config.Length.Column <= 0 {
		return
	}

	trimmedText := strings.TrimRight(text, "\r\n")
	end := token.Range.Start.Column + len(trimmedText) - 1

	if config.Length.AllowURL != nil && *config.Length.AllowURL {
		// Simple URL detection.
		if strings.Contains(trimmedText, "http://") || strings.Contains(trimmedText, "https://") {
			return
		}
	}

	if end > config.Length.Column {
		message := fmt.Sprintf("Wrap comment at column %d (currently %d).", config.Length.Column, end)
		if err := runner.EmitIssue(r, message, token.Range); err != nil {
			logger.Error(err.Error())
		}
		logger.Debug(message)
	}
}
