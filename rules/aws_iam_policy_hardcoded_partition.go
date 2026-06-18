package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIamPolicyHardcodedPartitionRule checks for hardcoded AWS partitions in IAM policies
type AwsIamPolicyHardcodedPartitionRule struct {
	tflint.DefaultRule
}

// NewAwsIamPolicyHardcodedPartitionRule returns a new rule
func NewAwsIamPolicyHardcodedPartitionRule() *AwsIamPolicyHardcodedPartitionRule {
	return &AwsIamPolicyHardcodedPartitionRule{}
}

// Name returns the rule name
func (r *AwsIamPolicyHardcodedPartitionRule) Name() string {
	return "aws_iam_policy_hardcoded_partition"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIamPolicyHardcodedPartitionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIamPolicyHardcodedPartitionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AwsIamPolicyHardcodedPartitionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks for hardcoded AWS partitions in IAM policies
func (r *AwsIamPolicyHardcodedPartitionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedPartitions(runner, r, "IAM policy", policy, rng)
	})
}
