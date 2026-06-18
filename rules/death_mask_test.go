package rules

import (
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestEosDeathMaskRule(t *testing.T) {
	content := heredoc.Doc(`
		# TEST
		# Although invalid syntax (no parent block), this should still emit an issue.
		# x = 1

		# TEST
		# Successive commented lines that collectively represent a valid expression.
		# resource "resource" "dead" {
		#   x = 1
		# }

		# TEST
		# A single dead line embedded in a larger, "live" block.
		resource "resource" "dead" {
		  # x = 1
		}

		# TEST
		# Mixed content, all of which is dead.
		# resource "foo" "bar" {
		#   # This is a nested comment
		#   name = "baz"
		# }

		resource "resource" "live" {
		  # A comment x = 1
		}
	`)

	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name:    "commented-out code",
			Content: content,
			Expected: helper.Issues{
				{Rule: NewDeathMaskRule(), Message: eosAvoidDeathMaskMessage},
				{Rule: NewDeathMaskRule(), Message: eosAvoidDeathMaskMessage},
				{Rule: NewDeathMaskRule(), Message: eosAvoidDeathMaskMessage},
				{Rule: NewDeathMaskRule(), Message: eosAvoidDeathMaskMessage},
			},
		},
	}

	rule := NewDeathMaskRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tc.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
