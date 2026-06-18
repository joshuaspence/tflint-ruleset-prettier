package rules

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// eosMisOrderedMessage is the message emitted when meta arguments are not in
// the configured order.
const eosMisOrderedMessage = "Meta arguments should be ordered consistently"

// eosCheckMetaOrder verifies that arguments in a block respect the Order
// config.
func eosCheckMetaOrder(runner tflint.Runner, r *EosMetaRule, config *eosMetaRuleConfig, block *hclsyntax.Block) {
	if len(config.Order) == 0 {
		return
	}
	order := config.Order[0]
	if len(order.First) == 0 && len(order.Last) == 0 {
		return
	}

	// Build sets for First and Last arguments.
	firstSet := make(map[string]bool, len(order.First))
	for _, name := range order.First {
		firstSet[name] = true
	}
	lastSet := make(map[string]bool, len(order.Last))
	for _, name := range order.Last {
		lastSet[name] = true
	}

	// Collect all attributes with their line numbers and classification.
	type attrInfo struct {
		name       string
		lineNumber int
		isFirst    bool
		isLast     bool
	}

	var attrs []attrInfo
	for name, attr := range block.Body.Attributes {
		info := attrInfo{
			name:       name,
			lineNumber: attr.SrcRange.Start.Line,
			isFirst:    firstSet[name],
			isLast:     lastSet[name],
		}
		attrs = append(attrs, info)
	}

	if len(attrs) == 0 {
		return
	}

	// Sort by line number to get the actual order in the source.
	for i := 0; i < len(attrs)-1; i++ {
		for j := i + 1; j < len(attrs); j++ {
			if attrs[i].lineNumber > attrs[j].lineNumber {
				attrs[i], attrs[j] = attrs[j], attrs[i]
			}
		}
	}

	// Check that First arguments appear before all non-First arguments.
	seenNonFirst := false
	for _, attr := range attrs {
		if !attr.isFirst {
			seenNonFirst = true
		} else if seenNonFirst {
			a := block.Body.Attributes[attr.name]
			r.emitIssue(runner, eosMisOrderedMessage, a.SrcRange)
			return
		}
	}

	// Check that Last arguments appear after all non-Last arguments.
	seenLast := false
	for _, attr := range attrs {
		if attr.isLast {
			seenLast = true
		} else if seenLast {
			a := block.Body.Attributes[attr.name]
			r.emitIssue(runner, eosMisOrderedMessage, a.SrcRange)
			return
		}
	}
}
