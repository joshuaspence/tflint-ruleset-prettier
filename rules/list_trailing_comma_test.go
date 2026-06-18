package rules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ListTrailingCommaRule(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Files    map[string]string
		Expected helper.Issues
	}{
    {                                                                                                                   
      Name: "trailing comma",                                                                                        
      Content: `                                                                                                        
resource "vault_generic_endpoint" "user" {                                                                              
  depends_on = [                                                                                                        
    random_password.svc_acc_pass,                                                                                       
  ]                                                                                                                     
}`,                                                                                                                     
      Expected: helper.Issues{},                                                                                         
    },
		{
			Name: "no trailing comma",
			Content: `
resource "vault_generic_endpoint" "user" {
  depends_on = [
    random_password.svc_acc_pass
  ]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewListTrailingCommaRule(),
					Message: "Last item in lists should always end with a trailing comma",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 16},
						End:      hcl.Pos{Line: 5, Column: 4},
					},
				},
			},
		},
		{
			Name: "heredoc without trailing comma",
			Content: `
resource "terraform_data" "test" {
  input = [
    "test",
    <<-HERE
      Lorem ipsum
    HERE
  ]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewListTrailingCommaRule(),
					Message: "Last item in lists should always end with a trailing comma",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 11},
						End:      hcl.Pos{Line: 8, Column: 4},
					},
				},
			},
		},
		{
			Name: "heredoc with trailing comma",
			Content: `
resource "terraform_data" "test" {
  input = [
    "test",
    <<-HERE
      Lorem ipsum
    HERE
    ,
  ]
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewListTrailingCommaRule()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			files := test.Files
			if files == nil {
				files = map[string]string{"resource.tf": test.Content}
			} else {
				if test.Content != "" {
					files["resource.tf"] = test.Content
				}
			}

			runner := helper.TestRunner(t, files)

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, test.Expected, runner.Issues)
		})
	}
}
