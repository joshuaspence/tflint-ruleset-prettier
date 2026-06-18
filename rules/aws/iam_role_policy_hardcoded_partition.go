package aws

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AwsIamRolePolicyHardcodedPartitionRule struct {
	tflint.DefaultRule
}

func NewAwsIamRolePolicyHardcodedPartitionRule() *AwsIamRolePolicyHardcodedPartitionRule {
	return &AwsIamRolePolicyHardcodedPartitionRule{}
}

func (r *AwsIamRolePolicyHardcodedPartitionRule) Name() string {
	return "aws_iam_role_policy_hardcoded_partition"
}

func (r *AwsIamRolePolicyHardcodedPartitionRule) Enabled() bool {
	return true
}

func (r *AwsIamRolePolicyHardcodedPartitionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *AwsIamRolePolicyHardcodedPartitionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *AwsIamRolePolicyHardcodedPartitionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_role_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedPartitions(runner, r, "IAM role policy", policy, rng)
	})
}
