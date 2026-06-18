package rules

import (
	"testing"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func TestEosCommentsRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name: "eol",
			Content: heredoc.Doc(`
				locals {
				  eol1 = 1 # This is an eol comment.
				  eol2 = 2 // This is an eol comment.
				  eol3 = {
				    eola = 1 # This is an eol comment.
				  }
				  eol4 = "4" /* This is an eol comment. */
				}

				# This is not an eol comment.
				// This is not an eol comment.

				locals {
				  # This is not an eol comment.
				  eol5 = 1
				}
			`),
			Expected: helper.Issues{
				{Rule: NewEosCommentsRule(), Message: eosAvoidEOLCommentsMessage},
				{Rule: NewEosCommentsRule(), Message: eosAvoidEOLCommentsMessage},
				{Rule: NewEosCommentsRule(), Message: eosAvoidEOLCommentsMessage},
				{Rule: NewEosCommentsRule(), Message: eosAvoidEOLCommentsMessage},
			},
		},
		{
			Name: "jammed",
			Content: heredoc.Doc(`
				#Jammed comment
				##Jammed comment
				//Jammed comment
				///Jammed comment

				# Good comment
				// Good comment
			`),
			Expected: helper.Issues{
				{Rule: NewEosCommentsRule(), Message: "Avoid jammed comment ('#Jamm ...')."},
				{Rule: NewEosCommentsRule(), Message: "Avoid jammed comment ('##Jam ...')."},
				{Rule: NewEosCommentsRule(), Message: "Avoid jammed comment ('//Jam ...')."},
				{Rule: NewEosCommentsRule(), Message: "Avoid jammed comment ('///Ja ...')."},
			},
		},
		{
			Name: "length",
			Content: heredoc.Doc(`
				# This comment is way too long and it will definitely extend beyond the eighty character limit that we have set for this rule.

				resource "foo" "bar" {
				  # Indented comment that is also way too long and should trigger the rule because it goes past column 80.
				}

				# Good comment

				# This comment is very long but it contains a url http://example.com/very/long/url/that/makes/this/line/exceed/the/limit so it should be ignored.
			`),
			Expected: helper.Issues{
				{Rule: NewEosCommentsRule(), Message: "Wrap comment at column 80 (currently 126)."},
				{Rule: NewEosCommentsRule(), Message: "Wrap comment at column 80 (currently 106)."},
			},
		},
		{
			Name: "jammed disabled",
			Content: heredoc.Doc(`
				#Jammed comment
			`),
			Config: heredoc.Doc(`
				rule "eos_comments" {
				  enabled = true
				  jammed  = false
				}
			`),
			Expected: helper.Issues{},
		},
		{
			Name: "eol disabled",
			Content: heredoc.Doc(`
				locals {
				  eol1 = 1 # This is an eol comment.
				}
			`),
			Config: heredoc.Doc(`
				rule "eos_comments" {
				  enabled = true
				  eol     = false
				}
			`),
			Expected: helper.Issues{},
		},
	}

	rule := NewEosCommentsRule()

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
