package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIamPolicyHardcodedRegionRule checks for hardcoded AWS regions in IAM policies
type AwsIamPolicyHardcodedRegionRule struct {
	tflint.DefaultRule
}

// NewAwsIamPolicyHardcodedRegionRule returns a new rule
func NewAwsIamPolicyHardcodedRegionRule() *AwsIamPolicyHardcodedRegionRule {
	return &AwsIamPolicyHardcodedRegionRule{}
}

// Name returns the rule name
func (r *AwsIamPolicyHardcodedRegionRule) Name() string {
	return "aws_iam_policy_hardcoded_region"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIamPolicyHardcodedRegionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIamPolicyHardcodedRegionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AwsIamPolicyHardcodedRegionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks for hardcoded AWS regions in IAM policies
func (r *AwsIamPolicyHardcodedRegionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedRegions(runner, r, "IAM policy", policy, rng)
	})
}
