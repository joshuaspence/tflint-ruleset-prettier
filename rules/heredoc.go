package rules

import (
	"regexp"

	"github.com/joshuaspence/tflint-ruleset-prettier/project"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

const eosAvoidEOFHeredocMessage = "Avoid using 'EOF' as the heredoc delimiter."
const eosAvoidStandardHeredocMessage = "Avoid standard heredoc (<<). Use indented (<<-) instead."

// eosHeredocPattern is a regex to match heredoc delimiters.
var eosHeredocPattern = regexp.MustCompile(`<<(-?)([a-zA-Z0-9]+)\s*$`)

// eosHeredocRuleConfig represents the configuration for the HeredocRule.
type eosHeredocRuleConfig struct {
	EOF   *bool  `hclext:"EOF,optional"`
	Level string `hclext:"level,optional"`
}

// HeredocRule checks for standard heredoc usage.
type HeredocRule struct {
	tflint.DefaultRule
}

// NewHeredocRule returns a new rule.
func NewHeredocRule() *HeredocRule {
	return &HeredocRule{}
}

// Name returns the rule name.
func (r *HeredocRule) Name() string {
	return "heredoc"
}

// Enabled returns whether the rule is enabled by default.
func (r *HeredocRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *HeredocRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *HeredocRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether the rule conditions are met.
func (r *HeredocRule) Check(runner tflint.Runner) error {
	config := &eosHeredocRuleConfig{
		EOF:   eosBoolPtr(true),
		Level: "warning",
	}
	if err := runner.DecodeRuleConfig(r.Name(), config); err != nil {
		return err
	}

	return eosWalkTokens(runner, r, func(runner tflint.Runner, rule *HeredocRule, token hclsyntax.Token) {
		eosCheckHeredocToken(runner, rule, config, token)
	})
}

// eosCheckHeredocToken checks for heredoc style violations in a token.
func eosCheckHeredocToken(runner tflint.Runner, r *HeredocRule, config *eosHeredocRuleConfig, token hclsyntax.Token) {
	if token.Type != hclsyntax.TokenOHeredoc {
		return
	}

	text := string(token.Bytes)
	matches := eosHeredocPattern.FindStringSubmatch(text)
	if matches == nil {
		return
	}

	indentMarker := matches[1]
	heredocLabel := matches[2]

	if indentMarker == "" {
		if err := runner.EmitIssue(r, eosAvoidStandardHeredocMessage, token.Range); err != nil {
			logger.Error(err.Error())
		}
	}

	if config.EOF != nil && *config.EOF {
		if heredocLabel == "EOF" {
			if err := runner.EmitIssue(r, eosAvoidEOFHeredocMessage, token.Range); err != nil {
				logger.Error(err.Error())
			}
		}
	}
}
