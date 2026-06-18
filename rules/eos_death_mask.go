package rules

import (
	"strings"

	"github.com/joshuaspence/tflint-ruleset-prettier/project"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosAvoidDeathMaskMessage is the message emitted when commented-out code is
// detected.
const eosAvoidDeathMaskMessage = "Avoid commented-out code."

// eosDeathMaskRuleConfig represents the configuration for the EosDeathMaskRule.
type eosDeathMaskRuleConfig struct {
	Level string `hclext:"level,optional"`
}

// EosDeathMaskRule checks for commented-out code.
type EosDeathMaskRule struct {
	tflint.DefaultRule
}

// NewEosDeathMaskRule returns a new rule.
func NewEosDeathMaskRule() *EosDeathMaskRule {
	return &EosDeathMaskRule{}
}

// Name returns the rule name.
func (r *EosDeathMaskRule) Name() string {
	return "eos_death_mask"
}

// Enabled returns whether the rule is enabled by default.
func (r *EosDeathMaskRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *EosDeathMaskRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *EosDeathMaskRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether the rule conditions are met.
func (r *EosDeathMaskRule) Check(runner tflint.Runner) error {
	config := &eosDeathMaskRuleConfig{Level: "warning"}
	if err := runner.DecodeRuleConfig(r.Name(), config); err != nil {
		return err
	}

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	for name, file := range files {
		if err := r.checkDeathMask(runner, name, file); err != nil {
			return err
		}
	}

	return nil
}

// checkDeathMask checks for commented-out code in a file.
func (r *EosDeathMaskRule) checkDeathMask(runner tflint.Runner, filename string, file *hcl.File) error {
	tokens, diags := hclsyntax.LexConfig(file.Bytes, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return diags
	}

	var commentBlock []hclsyntax.Token

	for _, token := range tokens {
		switch token.Type {
		case hclsyntax.TokenComment:
			if len(commentBlock) > 0 {
				last := commentBlock[len(commentBlock)-1]
				if token.Range.Start.Line > last.Range.End.Line {
					// Detected a gap, so flush the previous block.
					r.processCommentBlock(runner, commentBlock)
					commentBlock = nil
				}
			}
			commentBlock = append(commentBlock, token)
		case hclsyntax.TokenNewline, hclsyntax.TokenEOF:
			// Continue and let newlines pass.
		default:
			// A non-comment, non-newline token breaks the block.
			if len(commentBlock) > 0 {
				r.processCommentBlock(runner, commentBlock)
				commentBlock = nil
			}
		}
	}

	if len(commentBlock) > 0 {
		r.processCommentBlock(runner, commentBlock)
	}

	return nil
}

// processCommentBlock unwraps and validates a block of comments.
func (r *EosDeathMaskRule) processCommentBlock(runner tflint.Runner, tokens []hclsyntax.Token) {
	var lines []string
	for _, token := range tokens {
		text := string(token.Bytes)

		if s, cut := strings.CutPrefix(text, "//"); cut {
			s = strings.TrimPrefix(s, " ")
			lines = append(lines, s)
			continue
		}

		if s, cut := strings.CutPrefix(text, "#"); cut {
			s = strings.TrimPrefix(s, " ")
			lines = append(lines, s)
			continue
		}

		if s, cut := strings.CutPrefix(text, "/*"); cut {
			s = strings.TrimPrefix(s, "/*")
			s = strings.TrimSuffix(s, "*/")
			blockLines := strings.Split(s, "\n")
			lines = append(lines, blockLines...)
		}
	}

	// Try to parse subsets of lines to handle header text.
	for i := 0; i < len(lines); i++ {
		candidate := strings.Join(lines[i:], "\n")
		file, diags := hclsyntax.ParseConfig([]byte(candidate), "candidate.tf", hcl.InitialPos)
		if diags.HasErrors() {
			continue
		}

		if body, ok := file.Body.(*hclsyntax.Body); ok {
			if len(body.Attributes) > 0 || len(body.Blocks) > 0 {
				// It is valid code. Flag the whole block.
				start := tokens[0].Range.Start
				end := tokens[len(tokens)-1].Range.End
				issueRange := hcl.Range{
					Filename: tokens[0].Range.Filename,
					Start:    start,
					End:      end,
				}

				if err := runner.EmitIssue(r, eosAvoidDeathMaskMessage, issueRange); err != nil {
					logger.Error(err.Error())
				}
				return // We found a match, so we stop checking this block.
			}
		}
	}
}
