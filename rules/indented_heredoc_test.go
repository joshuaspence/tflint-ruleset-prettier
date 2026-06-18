package rules

import (
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestIndentedHeredocRule(t *testing.T) {
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
	expectedCount := 3

	rule := NewIndentedHeredocRule()
	runner := helper.TestRunner(t, map[string]string{"main.tf": content})

	if err := rule.Check(runner); err != nil {
		t.Fatalf("Unexpected error occurred: %s", err)
	}

  if len(runner.Issues) != expectedCount {
    t.Errorf("Expected %d issues, got %d", expectedCount, len(runner.Issues))
    for i, issue := range runner.Issues {
      t.Logf("Issue %d: %s", i+1, issue.Message)
    }
  }
}
