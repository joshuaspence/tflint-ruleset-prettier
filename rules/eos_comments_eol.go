package rules

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

const eosAvoidEOLCommentsMessage = "Avoid EOL comments."

// eosCheckEOL checks if EOL comments are used.
func eosCheckEOL(r *EosCommentsRule, config *eosCommentsRuleConfig, _ string, runner tflint.Runner, token hclsyntax.Token, prevToken *hclsyntax.Token) {
	if config.EOL == nil || !*config.EOL {
		return
	}

	if prevToken != nil {
		if prevToken.Type == hclsyntax.TokenNewline {
			return
		}

		if prevToken.Range.End.Line == token.Range.Start.Line {
			message := eosAvoidEOLCommentsMessage
			if err := runner.EmitIssue(r, message, token.Range); err != nil {
				logger.Error(err.Error())
			}
			logger.Debug(message)
		}
	}
}
