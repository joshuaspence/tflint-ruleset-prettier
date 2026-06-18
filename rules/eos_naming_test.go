package rules

import (
	"fmt"
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestEosNamingRule(t *testing.T) {
	longNameMsg := fmt.Sprintf("Avoid names longer than %d ('really_a_very_long_name' is 23).", eosNamingDefaultLimit)

	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name: "length",
			Config: heredoc.Doc(`
				rule "eos_naming" {
				  enabled = true
				  shout   = false
				  snake   = false
				  type_echo {
				    enabled = false
				  }
				}
			`),
			Content: heredoc.Doc(`
				variable "really_a_very_long_name" {}

				locals {
				  really_a_very_long_name = 1
				}

				output "really_a_very_long_name" {
				  value = "test"
				}

				resource "aws_instance" "really_a_very_long_name" {
				  ami = "ami-12345678"
				}

				variable "short" {}
			`),
			Expected: helper.Issues{
				{Rule: NewEosNamingRule(), Message: longNameMsg},
				{Rule: NewEosNamingRule(), Message: longNameMsg},
				{Rule: NewEosNamingRule(), Message: longNameMsg},
				{Rule: NewEosNamingRule(), Message: longNameMsg},
			},
		},
		{
			Name: "shout",
			Config: heredoc.Doc(`
				rule "eos_naming" {
				  enabled = true
				  snake   = false
				  type_echo {
				    enabled = false
				  }
				}
			`),
			Content: heredoc.Doc(`
				variable "SHOUT" {}

				locals {
				  SHOUT = 1
				}

				output "SHOUT" {
				  value = 1
				}

				variable "no_shout" {}
			`),
			Expected: helper.Issues{
				{Rule: NewEosNamingRule(), Message: "Avoid SHOUTED names (SHOUT)"},
				{Rule: NewEosNamingRule(), Message: "Avoid SHOUTED names (SHOUT)"},
				{Rule: NewEosNamingRule(), Message: "Avoid SHOUTED names (SHOUT)"},
			},
		},
		{
			Name: "snake",
			Config: heredoc.Doc(`
				rule "eos_naming" {
				  enabled = true
				  shout   = false
				  type_echo {
				    enabled = false
				  }
				}
			`),
			Content: heredoc.Doc(`
				variable "CamelCase" {}
				variable "kebab-case" {}
				variable "with.dots" {}

				variable "snake_case" {}
				variable "lower123" {}
			`),
			Expected: helper.Issues{
				{Rule: NewEosNamingRule(), Message: "Names should be snake_case (CamelCase)."},
				{Rule: NewEosNamingRule(), Message: "Names should be snake_case (kebab-case)."},
				{Rule: NewEosNamingRule(), Message: "Names should be snake_case (with.dots)."},
			},
		},
		{
			Name: "type echo",
			Config: heredoc.Doc(`
				rule "eos_naming" {
				  enabled = true
				  length  = -1
				  shout   = false
				  snake   = false
				}
			`),
			Content: heredoc.Doc(`
				variable "variable_echo" {}

				locals {
				  local_echo = 1
				}

				data "aws_caller_identity" "caller_echo" {}

				output "output_echo" {
				  value = 1
				}

				variable "clean_var" {}
			`),
			Expected: helper.Issues{
				{Rule: NewEosNamingRule(), Message: eosMakeTypeEchoMessage("variable", "variable_echo")},
				{Rule: NewEosNamingRule(), Message: eosMakeTypeEchoMessage("local", "local_echo")},
				{Rule: NewEosNamingRule(), Message: eosMakeTypeEchoMessage("aws_caller_identity", "caller_echo")},
				{Rule: NewEosNamingRule(), Message: eosMakeTypeEchoMessage("output", "output_echo")},
			},
		},
		{
			Name: "type echo via synonym",
			Config: heredoc.Doc(`
				rule "eos_naming" {
				  enabled = true
				  length  = -1
				  shout   = false
				  snake   = false
				  type_echo {
				    synonyms = {
				      bucket = ["pail"]
				    }
				  }
				}
			`),
			Content: heredoc.Doc(`
				resource "aws_s3_bucket" "my_pail" {}

				resource "aws_s3_bucket" "clean_name" {}
			`),
			Expected: helper.Issues{
				{Rule: NewEosNamingRule(), Message: eosMakeTypeEchoSynonymMessage("aws_s3_bucket", "pail", "my_pail")},
			},
		},
	}

	rule := NewEosNamingRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			files := map[string]string{"main.tf": tc.Content}
			if tc.Config != "" {
				files[".tflint.hcl"] = tc.Config
			}
			runner := helper.TestRunner(t, files)

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}

func eosMakeTypeEchoMessage(typ string, name string) string {
	return fmt.Sprintf("Avoid echoing type \"%s\" in label \"%s\".", typ, name)
}

func eosMakeTypeEchoSynonymMessage(typ string, synonym string, name string) string {
	return fmt.Sprintf("Avoid echoing type \"%s\" (via synonym '%s') in label \"%s\".", typ, synonym, name)
}
