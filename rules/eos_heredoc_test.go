package rules

import (
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestEosHeredocRule(t *testing.T) {
	content := heredoc.Doc(`
		variable "heredoc" {
		  value = <<EOF
		This is the heredoc.
		EOF
		}

		locals {
		  heredoc = <<EOF
		This is the heredoc.
		EOF
		}

		check "heredoc" {
		  assert {
		    condition     = local.heredoc == <<-EOF
		This is the heredoc.
		EOF
		    error_message = "Wrong heredoc syntax."
		  }
		}

		ephemeral "random_password" "heredoc" {
		  lower = <<-EOF
		This is the heredoc.
		EOF
		}

		module "heredoc" {
		  source = <<EOF
		This is the heredoc.
		EOF
		}

		output "heredoc" {
		  value = <<-EOF
		This is the heredoc.
		EOF
		}

		variable "no_heredoc" {
		  value = <<-VAR
		    This is the heredoc.
		    VAR
		}
	`)

	// Issues are emitted in file order. For non-indented (<<EOF) delimiters,
	// the standard-heredoc issue is emitted before the EOF issue.
	expected := helper.Issues{
		// variable: <<EOF
		{Rule: NewEosHeredocRule(), Message: eosAvoidStandardHeredocMessage},
		{Rule: NewEosHeredocRule(), Message: eosAvoidEOFHeredocMessage},
		// locals: <<EOF
		{Rule: NewEosHeredocRule(), Message: eosAvoidStandardHeredocMessage},
		{Rule: NewEosHeredocRule(), Message: eosAvoidEOFHeredocMessage},
		// check: <<-EOF
		{Rule: NewEosHeredocRule(), Message: eosAvoidEOFHeredocMessage},
		// ephemeral: <<-EOF
		{Rule: NewEosHeredocRule(), Message: eosAvoidEOFHeredocMessage},
		// module: <<EOF
		{Rule: NewEosHeredocRule(), Message: eosAvoidStandardHeredocMessage},
		{Rule: NewEosHeredocRule(), Message: eosAvoidEOFHeredocMessage},
		// output: <<-EOF
		{Rule: NewEosHeredocRule(), Message: eosAvoidEOFHeredocMessage},
	}

	rule := NewEosHeredocRule()
	runner := helper.TestRunner(t, map[string]string{"main.tf": content})

	if err := rule.Check(runner); err != nil {
		t.Fatalf("Unexpected error occurred: %s", err)
	}

	AssertIssues(t, expected, runner.Issues)
}
