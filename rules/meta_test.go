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
		{
			Name: "source version",
			Content: heredoc.Doc(`
				module "fail_gh_no_ref" {
				  source = "github.com/eos/test.git"
				}

				module "fail_git_no_ref" {
				  source = "git::github.com/eos/test.git"
				}

				module "fail_hg_no_ref" {
				  source = "hg::https://example.com/eos/test"
				}

				module "fail_https_no_extension" {
				  source = "https://example.com/eos/test"
				}

				module "fail_https_no_extension2" {
				  source = "https://example.com/eos/test.bad"
				}

				module "fail_no_version" {
				  source = "eos/module/zakpxy"
				}

				module "fail_no_version2" {
				  source = "app.terraform.io/eos/module/zakpxy"
				}

				module "fail_version_gt" {
				  source  = "eos/module/zakpxy"
				  version = "> 1.2.0"
				}

				module "fail_version_gte" {
				  source  = "eos/module/zakpxy"
				  version = ">= 1.2.0"
				}

				module "fail_version_mixed_gte" {
				  source  = "eos/module/zakpxy"
				  version = "~> 1.2.0, >= 1.0"
				}

				module "fail_version_pessimistic_short" {
				  source  = "eos/module/zakpxy"
				  version = "~> 1"
				}

				module "pass_git_ref" {
				  source = "github.com/eos/test.git?ref=v0.0.1"
				}

				module "pass_version_pessimistic" {
				  source  = "eos/module/zakpxy"
				  version = "~> 1.2"
				}

				module "pass_https_zip_extension" {
				  source = "https://example.com/eos/test.zip"
				}

				module "pass_https_tar_gz_extension" {
				  source = "https://example.com/eos/test.tar.gz"
				}

				module "pass_https_archive_query" {
				  source = "https://example.com/eos/test?archive=.zip"
				}
			`),
			Expected: helper.Issues{
				{Rule: NewMetaRule(), Message: "Git module source should specify ref parameter."},
				{Rule: NewMetaRule(), Message: "Git module source should specify ref parameter."},
				{Rule: NewMetaRule(), Message: "Mercurial module source should specify #revision."},
				{Rule: NewMetaRule(), Message: "https module source should specify a valid archive extension."},
				{Rule: NewMetaRule(), Message: "https module source should specify a valid archive extension."},
				{Rule: NewMetaRule(), Message: "Module from registry should specify version."},
				{Rule: NewMetaRule(), Message: "Module from registry should specify version."},
				{Rule: NewMetaRule(), Message: "Version constraint > or >= should not be used. Use ~> or exact version."},
				{Rule: NewMetaRule(), Message: "Version constraint > or >= should not be used. Use ~> or exact version."},
				{Rule: NewMetaRule(), Message: "Version constraint > or >= should not be used. Use ~> or exact version."},
				{Rule: NewMetaRule(), Message: "Pessimistic version constraint should specify at least major and minor version."},
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
