package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_MapTrailingCommaRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "without trailing commas",
			Content: `
locals {
  a_dictionary = {
    one  = "fish"
    two  = "fish"
    red  = "fish"
    blue = "fish"
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "with trailing commas",
			Content: `
locals {
  a_dictionary = {
    one  = "fish",
    two  = "fish",
    red  = "fish",
    blue = "fish",
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewMapTrailingCommaRule(),
					Message: "Map values should not have trailing commas",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 12},
						End:      hcl.Pos{Line: 4, Column: 18},
					},
				},
				{
					Rule:    NewMapTrailingCommaRule(),
					Message: "Map values should not have trailing commas",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 12},
						End:      hcl.Pos{Line: 5, Column: 18},
					},
				},
				{
					Rule:    NewMapTrailingCommaRule(),
					Message: "Map values should not have trailing commas",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 12},
						End:      hcl.Pos{Line: 6, Column: 18},
					},
				},
				{
					Rule:    NewMapTrailingCommaRule(),
					Message: "Map values should not have trailing commas",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 7, Column: 12},
						End:      hcl.Pos{Line: 7, Column: 18},
					},
				},
			},
		},
		{
			Name: "do not remove separator comma on same line",
			Content: `
locals {
  a_dictionary = {
    "one" = "fish", "two" = "fish"
    "red" = "fish"
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name:     "single line map",
			Content:  `b_dictionary = { "one" = "fish", "two" = "fish", "red" = "fish", "blue" = "fish" }`,
			Expected: helper.Issues{},
		},
	}

	rule := NewMapTrailingCommaRule()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, test.Expected, runner.Issues)
		})
	}
}
