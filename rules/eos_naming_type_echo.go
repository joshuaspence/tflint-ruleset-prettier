package rules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosCheckTypeEcho checks if a word in the type is echoed in the name.
func eosCheckTypeEcho(runner tflint.Runner, ctx eosNamingContext, defRange hcl.Range, typ string, name string, synonym string) {
	// Assume there is no echo.
	echo := false

	lowerTyp := strings.ToLower(typ)   // aws_s3_bucket
	lowerName := strings.ToLower(name) // my_bucket
	synonymText := ""

	// For each word in type, see if it exists in name.
	for part := range strings.SplitSeq(lowerTyp, "_") {
		if strings.Contains(lowerName, part) {
			echo = true
			break
		}

		// Get synonyms for the word.
		var synonyms []string
		te := ctx.config.TypeEcho
		if te != nil && te.Synonyms != nil {
			synonyms = te.Synonyms[part]
			if synonym != "" {
				synonyms = append(synonyms, synonym)
			}
		}

		// Check synonyms on word boundaries.
		nameParts := strings.SplitSeq(lowerName, "_-")
		for _, syn := range synonyms {
			for n := range nameParts {
				if strings.Contains(n, syn) {
					echo = true
					synonymText = fmt.Sprintf(" (via synonym '%s')", syn)
					break
				}
			}

			if echo {
				break
			}
		}
	}

	if echo {
		if err := runner.EmitIssue(
			ctx.rule,
			fmt.Sprintf("Avoid echoing type \"%s\"%s in label \"%s\".", typ, synonymText, name),
			defRange,
		); err != nil {
			logger.Error(err.Error())
		}
	}
}
