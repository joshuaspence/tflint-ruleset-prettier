package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIamRolePolicyHardcodedPartitionRule checks for hardcoded AWS partitions in IAM role policies
type AwsIamRolePolicyHardcodedPartitionRule struct {
	tflint.DefaultRule
}

// NewAwsIamRolePolicyHardcodedPartitionRule returns a new rule
func NewAwsIamRolePolicyHardcodedPartitionRule() *AwsIamRolePolicyHardcodedPartitionRule {
	return &AwsIamRolePolicyHardcodedPartitionRule{}
}

// Name returns the rule name
func (r *AwsIamRolePolicyHardcodedPartitionRule) Name() string {
	return "aws_iam_role_policy_hardcoded_partition"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIamRolePolicyHardcodedPartitionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIamRolePolicyHardcodedPartitionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AwsIamRolePolicyHardcodedPartitionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks for hardcoded AWS partitions in IAM role policies
func (r *AwsIamRolePolicyHardcodedPartitionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_role_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedPartitions(runner, r, "IAM role policy", policy, rng)
	})
}
