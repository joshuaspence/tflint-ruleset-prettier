package rules

import (
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestEosMetaRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "count guard",
			Content: heredoc.Doc(`
				locals {
				  prod_bool  = true
				  prod_count = 1
				}

				resource "null_resource" "loop" {
				  count = 2
				}

				resource "terraform_data" "no01guard" {
				  count = local.prod_bool ? 0 : 2
				  input = count.index
				}

				resource "terraform_data" "emit2" {
				  count = local.prod_bool ? 2 : 0
				  input = count.index
				}

				resource "bad" "emit3" {
				  count = local.prod_count
				}

				resource "bad" "emit4" {
				  count = length([1, 2])
				}

				resource "good" "no_emit0" {
				  count = 0
				}

				resource "good" "no_emit1" {
				  count = 1
				}

				resource "good" "no_emit_cond" {
				  count = local.prod_bool ? 0 : 1
				}
			`),
			Expected: helper.Issues{
				{Rule: NewMetaRule(), Message: eosOnlyDynamicGuardMessage},
				{Rule: NewMetaRule(), Message: eosGuardMustReturn10Message},
				{Rule: NewMetaRule(), Message: eosGuardMustReturn10Message},
				{Rule: NewMetaRule(), Message: eosOnlyDynamicGuardMessage},
				{Rule: NewMetaRule(), Message: eosOnlyDynamicGuardMessage},
			},
		},
		{
			Name: "order",
			Content: heredoc.Doc(`
				resource "terraform_data" "first_after_other" {
				  input    = "x"
				  for_each = { za = 1, kp = 2 }
				}

				resource "terraform_data" "last_before_other" {
				  depends_on = [terraform_data.first_after_other]
				  input      = "test"
				}

				resource "terraform_data" "mixed_violation" {
				  for_each   = { za = 1, kp = 2 }
				  depends_on = [terraform_data.first_after_other]
				  input      = "x"
				}

				resource "terraform_data" "pass_first_correct" {
				  for_each = { za = 1, kp = 2 }
				  input    = "x"
				}

				resource "terraform_data" "pass_last_correct" {
				  input      = "test"
				  depends_on = [terraform_data.first_after_other]
				}
			`),
			Expected: helper.Issues{
				{Rule: NewMetaRule(), Message: eosMisOrderedMessage},
				{Rule: NewMetaRule(), Message: eosMisOrderedMessage},
				{Rule: NewMetaRule(), Message: eosMisOrderedMessage},
			},
		},
	}

	rule := NewMetaRule()

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
