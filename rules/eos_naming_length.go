package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosCheckNameLength checks if the name is too long.
func eosCheckNameLength(runner tflint.Runner, ctx eosNamingContext, defRange hcl.Range, _ string, name string, _ string) {
	limit := eosNamingDefaultLimit
	if ctx.config.Length != nil {
		limit = *ctx.config.Length
	}

	if len(name) > limit {
		message := fmt.Sprintf("Avoid names longer than %d ('%s' is %d).", limit, name, len(name))
		if err := runner.EmitIssue(ctx.rule, message, defRange); err != nil {
			logger.Error(err.Error())
		}
		logger.Debug(message)
	}
}
