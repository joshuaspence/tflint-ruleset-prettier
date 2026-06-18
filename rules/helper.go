package rules

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// These helpers are flattened from the upstream
// tfctl/tflint-ruleset-elements-of-style "internal/rulehelper" package. They
// are prefixed with "eos" to avoid colliding with the other rules bundled in
// this package.

// eosBoolPtr returns a pointer to the given bool value.
func eosBoolPtr(b bool) *bool {
	return &b
}

// eosToSeverity converts a string level to a tflint.Severity.
func eosToSeverity(level string) tflint.Severity {
	switch strings.ToLower(level) {
	case "notice":
		return tflint.NOTICE
	case "warning":
		return tflint.WARNING
	}

	return tflint.ERROR
}

// eosBlockDef represents a block definition for schema generation.
type eosBlockDef struct {
	Typ     string
	Labels  []string
	Synonym string
}

// eosAllLintableBlocks defines all block types and their label structures to
// check.
var eosAllLintableBlocks = []eosBlockDef{
	{Typ: "variable", Labels: []string{"name"}},
	{Typ: "check", Labels: []string{"name"}},
	{Typ: "data", Labels: []string{"type", "name"}},
	{Typ: "ephemeral", Labels: []string{"type", "name"}},
	{Typ: "module", Labels: []string{"name"}},
	{Typ: "output", Labels: []string{"name"}},
	{Typ: "resource", Labels: []string{"type", "name"}},
}

// eosNormalizeBlock extracts the type, name, and synonym from a block.
func eosNormalizeBlock(block *hclsyntax.Block, myBlocks []eosBlockDef) (string, string, string) {
	var name string
	var typ string

	switch len(block.Labels) {
	case 2:
		typ = block.Labels[0]
		name = block.Labels[1]
	case 1:
		typ = block.Type
		name = block.Labels[0]
	default:
		typ = block.Type
		name = ""
	}

	synonym := ""
	for _, def := range myBlocks {
		if def.Typ == typ && def.Synonym != "" {
			synonym = def.Synonym
			break
		}
	}
	return typ, name, synonym
}

// eosWalkBlocks iterates over blocks and locals using the AST and applies the
// check functions.
func eosWalkBlocks[T any](
	runner tflint.Runner,
	myBlocks []eosBlockDef,
	rule T,
	checkFuncs ...func(tflint.Runner, T, hcl.Range, string, string, string),
) error {
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
			// Handle locals specifically.
			if block.Type == "locals" {
				for _, attr := range block.Body.Attributes {
					for _, checkFunc := range checkFuncs {
						checkFunc(runner, rule, attr.Range(), "local", attr.Name, "")
					}
				}
				continue
			}

			typ, name, synonym := eosNormalizeBlock(block, myBlocks)

			// Filter by myBlocks to ensure we only lint what we expect.
			found := false
			for _, def := range myBlocks {
				if def.Typ == block.Type {
					found = true
					break
				}
			}
			if !found {
				continue
			}

			defRange := block.TypeRange
			if len(block.LabelRanges) > 0 {
				lastLabel := block.LabelRanges[len(block.LabelRanges)-1]
				defRange = hcl.Range{
					Filename: defRange.Filename,
					Start:    defRange.Start,
					End:      lastLabel.End,
				}
			}

			for _, checkFunc := range checkFuncs {
				checkFunc(runner, rule, defRange, typ, name, synonym)
			}
		}
	}

	return nil
}

// eosWalkTokens iterates over all files in the root module, lexes them, and
// applies the check function to each token.
func eosWalkTokens[T any](
	runner tflint.Runner,
	rule T,
	checkFunc func(tflint.Runner, T, hclsyntax.Token),
) error {
	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		return nil
	}

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	for filename, file := range files {
		tokens, diags := hclsyntax.LexConfig(file.Bytes, filename, hcl.InitialPos)
		if diags.HasErrors() {
			return diags
		}

		for _, token := range tokens {
			checkFunc(runner, rule, token)
		}
	}

	return nil
}
