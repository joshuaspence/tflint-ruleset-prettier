package rules

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_TypedVariables_Valid(t *testing.T) {
	rule := NewTypedVariablesRule()

	runner := helper.TestRunner(t, map[string]string{
		"variables.tf": `
variable "name" {
  type = string
}

variable "tags" {
  type = map(string)
}
`,
	})

	if err := rule.Check(runner); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(runner.Issues) != 0 {
		t.Errorf("expected no issues, got %d", len(runner.Issues))
	}
}

func Test_TypedVariables_Missing(t *testing.T) {
	rule := NewTypedVariablesRule()

	runner := helper.TestRunner(t, map[string]string{
		"variables.tf": `
variable "name" {
  description = "The name"
}
`,
	})

	if err := rule.Check(runner); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(runner.Issues) != 1 {
		t.Fatalf("expected 1 issue, got %d", len(runner.Issues))
	}

	expected := `Variable "name" is missing a type constraint.`
	if runner.Issues[0].Message != expected {
		t.Errorf("expected message %q, got %q", expected, runner.Issues[0].Message)
	}
}
