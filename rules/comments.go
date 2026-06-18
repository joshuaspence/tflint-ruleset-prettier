package rules

import (
	"github.com/joshuaspence/tflint-ruleset-prettier/project"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosCommentsLengthConfig represents the configuration for comment length
// checks.
type eosCommentsLengthConfig struct {
	// AllowURL allows comments with a URL to bust the column limit.
	AllowURL *bool `hclext:"allow_url,optional"`
	// Column is the maximum allowed column for comments. <=0 effectively
	// disables the check.
	Column int `hclext:"column,optional"`
}

// eosCommentsRuleConfig represents the configuration for the CommentsRule.
type eosCommentsRuleConfig struct {
	// EOL enables the end-of-line comment check.
	EOL *bool `hclext:"eol,optional"`
	// Jammed enables the jammed comment check.
	Jammed *bool `hclext:"jammed,optional"`
	// Length configures the comment length check.
	Length *eosCommentsLengthConfig `hclext:"length,block"`
	// Level is the issue severity level.
	Level string `hclext:"level,optional"`
}

// CommentsRule checks for comment style.
type CommentsRule struct {
	tflint.DefaultRule
}

// NewCommentsRule returns a new rule.
func NewCommentsRule() *CommentsRule {
	return &CommentsRule{}
}

// Name returns the rule name.
func (r *CommentsRule) Name() string {
	return "comments"
}

// Enabled returns whether the rule is enabled by default.
func (r *CommentsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *CommentsRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *CommentsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether the rule conditions are met.
func (r *CommentsRule) Check(runner tflint.Runner) error {
	config := &eosCommentsRuleConfig{
		EOL:    eosBoolPtr(true),
		Jammed: eosBoolPtr(true),
		Length: &eosCommentsLengthConfig{
			AllowURL: eosBoolPtr(true),
			Column:   80,
		},
		Level: "warning",
	}
	if err := runner.DecodeRuleConfig(r.Name(), config); err != nil {
		return err
	}

	return eosCheckCommentsWithContext(runner, r, config,
		eosCheckEOL,
		eosCheckJammed,
		eosCheckLength)
}

// eosCheckCommentsWithContext iterates over all files in the root module,
// parses them, and applies the check functions to each comment token,
// providing the previous token for context.
func eosCheckCommentsWithContext(
	runner tflint.Runner,
	rule *CommentsRule,
	config *eosCommentsRuleConfig,
	checkFuncs ...func(*CommentsRule, *eosCommentsRuleConfig, string, tflint.Runner, hclsyntax.Token, *hclsyntax.Token),
) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	for filename, file := range files {
		tokens, diags := hclsyntax.LexConfig(file.Bytes, filename, hcl.InitialPos)
		if diags.HasErrors() {
			return diags
		}

		for i, token := range tokens {
			if token.Type != hclsyntax.TokenComment {
				continue
			}

			var prevToken *hclsyntax.Token
			for j := i - 1; j >= 0; j-- {
				if tokens[j].Type != hclsyntax.TokenComment {
					prevToken = &tokens[j]
					break
				}
			}

			for _, checkFunc := range checkFuncs {
				checkFunc(rule, config, string(token.Bytes), runner, token, prevToken)
			}
		}
	}

	return nil
}
