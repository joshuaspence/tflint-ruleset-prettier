package rules

import (
	"github.com/joshuaspence/tflint-ruleset-prettier/project"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosNamingDefaultLimit is the default maximum length for names.
const eosNamingDefaultLimit = 16

// eosTypeEchoConfig represents the configuration for type echo checks.
type eosTypeEchoConfig struct {
	Enabled  *bool               `hclext:"enabled,optional"`
	Level    string              `hclext:"level,optional"`
	Synonyms map[string][]string `hclext:"synonyms,optional"`
}

// eosNamingRuleConfig represents the configuration for the EosNamingRule.
type eosNamingRuleConfig struct {
	Level    string             `hclext:"level,optional"`
	Length   *int               `hclext:"length,optional"`
	Shout    *bool              `hclext:"shout,optional"`
	Snake    *bool              `hclext:"snake,optional"`
	TypeEcho *eosTypeEchoConfig `hclext:"type_echo,block"`
}

// eosNamingContext carries the rule and its resolved config through the block
// walk.
type eosNamingContext struct {
	rule   *EosNamingRule
	config *eosNamingRuleConfig
}

// EosNamingRule checks whether a block's name is excessively long or otherwise
// violates naming conventions.
type EosNamingRule struct {
	tflint.DefaultRule
}

// NewEosNamingRule returns a new rule.
func NewEosNamingRule() *EosNamingRule {
	return &EosNamingRule{}
}

// Name returns the rule name.
func (r *EosNamingRule) Name() string {
	return "eos_naming"
}

// Enabled returns whether the rule is enabled by default.
func (r *EosNamingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *EosNamingRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *EosNamingRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether the rule conditions are met.
func (r *EosNamingRule) Check(runner tflint.Runner) error {
	defaultLength := eosNamingDefaultLimit
	config := &eosNamingRuleConfig{
		Level:  "warning",
		Length: &defaultLength,
	}
	if err := runner.DecodeRuleConfig(r.Name(), config); err != nil {
		return err
	}

	ctx := eosNamingContext{rule: r, config: config}

	var checks []func(tflint.Runner, eosNamingContext, hcl.Range, string, string, string)
	length := eosNamingDefaultLimit
	if config.Length != nil {
		length = *config.Length
	}
	if length > 0 {
		checks = append(checks, eosCheckNameLength)
	}

	if config.Shout == nil || *config.Shout {
		checks = append(checks, eosCheckShout)
	}

	if config.Snake == nil || *config.Snake {
		checks = append(checks, eosCheckSnake)
	}

	te := config.TypeEcho
	if te == nil || te.Enabled == nil || *te.Enabled {
		checks = append(checks, eosCheckTypeEcho)
	}

	return eosWalkBlocks(runner, eosAllLintableBlocks, ctx, checks...)
}
