package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIamRolePolicyHardcodedRegionRule checks for hardcoded AWS regions in IAM role policies
type AwsIamRolePolicyHardcodedRegionRule struct {
	tflint.DefaultRule
}

// NewAwsIamRolePolicyHardcodedRegionRule returns a new rule
func NewAwsIamRolePolicyHardcodedRegionRule() *AwsIamRolePolicyHardcodedRegionRule {
	return &AwsIamRolePolicyHardcodedRegionRule{}
}

// Name returns the rule name
func (r *AwsIamRolePolicyHardcodedRegionRule) Name() string {
	return "aws_iam_role_policy_hardcoded_region"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIamRolePolicyHardcodedRegionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIamRolePolicyHardcodedRegionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AwsIamRolePolicyHardcodedRegionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks for hardcoded AWS regions in IAM role policies
func (r *AwsIamRolePolicyHardcodedRegionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_role_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedRegions(runner, r, "IAM role policy", policy, rng)
	})
}
