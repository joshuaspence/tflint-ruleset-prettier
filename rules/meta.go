package rules

import (
	"github.com/joshuaspence/tflint-ruleset-prettier/project"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosMetaOrderConfig defines which arguments must appear first and last in a
// block.
type eosMetaOrderConfig struct {
	First []string `hclext:"first,optional"`
	Last  []string `hclext:"last,optional"`
}

// eosMetaRuleConfig represents the configuration for the MetaRule.
type eosMetaRuleConfig struct {
	Level         string               `hclext:"level,optional"`
	Order         []eosMetaOrderConfig `hclext:"order,block"`
	SourceVersion *bool                `hclext:"source_version,optional"`
}

// MetaRule checks for meta-argument style violations.
type MetaRule struct {
	tflint.DefaultRule
}

// NewMetaRule returns a new rule.
func NewMetaRule() *MetaRule {
	return &MetaRule{}
}

// Name returns the rule name.
func (r *MetaRule) Name() string {
	return "meta"
}

// Enabled returns whether the rule is enabled by default.
func (r *MetaRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *MetaRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *MetaRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether the rule conditions are met.
func (r *MetaRule) Check(runner tflint.Runner) error {
	config := &eosMetaRuleConfig{
		Level: "warning",
		Order: []eosMetaOrderConfig{{
			First: []string{"for_each", "count"},
			Last:  []string{"depends_on", "provider", "lifecycle"},
		}},
		SourceVersion: eosBoolPtr(true),
	}
	if err := runner.DecodeRuleConfig(r.Name(), config); err != nil {
		return err
	}

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		body, ok := file.Body.(*hclsyntax.Body)
		if !ok {
			continue
		}
		for _, block := range body.Blocks {
			eosCheckMetaOrder(runner, r, config, block)
			if attr, exists := block.Body.Attributes["count"]; exists {
				eosCheckCountGuard(runner, r, attr)
			}
			if block.Type == "module" && (config.SourceVersion == nil || *config.SourceVersion) {
				eosCheckModuleSourceVersion(runner, r, block)
			}
		}
	}

	return nil
}

func (r *MetaRule) emitIssue(runner tflint.Runner, message string, rng hcl.Range) {
	if err := runner.EmitIssue(r, message, rng); err != nil {
		logger.Error(err.Error())
	}
}
