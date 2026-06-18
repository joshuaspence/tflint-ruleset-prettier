package rules

import (
	"github.com/joshuaspence/tflint-ruleset-prettier/project"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosTypeEchoConfig represents the configuration for type echo checks.
type eosTypeEchoConfig struct {
	Enabled  *bool               `hclext:"enabled,optional"`
	Level    string              `hclext:"level,optional"`
	Synonyms map[string][]string `hclext:"synonyms,optional"`
}

// eosNamingRuleConfig represents the configuration for the NamingRule.
type eosNamingRuleConfig struct {
	Level    string             `hclext:"level,optional"`
	Snake    *bool              `hclext:"snake,optional"`
	TypeEcho *eosTypeEchoConfig `hclext:"type_echo,block"`
}

// eosNamingContext carries the rule and its resolved config through the block
// walk.
type eosNamingContext struct {
	rule   *NamingRule
	config *eosNamingRuleConfig
}

// NamingRule checks whether a block's name is excessively long or otherwise
// violates naming conventions.
type NamingRule struct {
	tflint.DefaultRule
}

// NewNamingRule returns a new rule.
func NewNamingRule() *NamingRule {
	return &NamingRule{}
}

// Name returns the rule name.
func (r *NamingRule) Name() string {
	return "naming"
}

// Enabled returns whether the rule is enabled by default.
func (r *NamingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *NamingRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *NamingRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether the rule conditions are met.
func (r *NamingRule) Check(runner tflint.Runner) error {
	config := &eosNamingRuleConfig{
		Level: "warning",
	}
	if err := runner.DecodeRuleConfig(r.Name(), config); err != nil {
		return err
	}

	ctx := eosNamingContext{rule: r, config: config}

	var checks []func(tflint.Runner, eosNamingContext, hcl.Range, string, string, string)

	if config.Snake == nil || *config.Snake {
		checks = append(checks, eosCheckSnake)
	}

	te := config.TypeEcho
	if te == nil || te.Enabled == nil || *te.Enabled {
		checks = append(checks, eosCheckTypeEcho)
	}

	return eosWalkBlocks(runner, eosAllLintableBlocks, ctx, checks...)
}
