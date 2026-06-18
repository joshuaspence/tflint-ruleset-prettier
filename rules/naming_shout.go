package rules

import (
	"fmt"
	"unicode"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosCheckShout checks if the name is all uppercase.
func eosCheckShout(runner tflint.Runner, ctx eosNamingContext, defRange hcl.Range, _ string, name string, _ string) {
	hasAlpha := false
	allUpper := true

	for _, ch := range name {
		if unicode.IsLetter(ch) {
			hasAlpha = true
			if !unicode.IsUpper(ch) {
				allUpper = false
			}
		}
	}

	if hasAlpha && allUpper {
		message := fmt.Sprintf("Avoid SHOUTED names (%s)", name)
		if err := runner.EmitIssue(ctx.rule, message, defRange); err != nil {
			logger.Error(err.Error())
		}
		logger.Debug(message)
	}
}
