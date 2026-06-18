package aws

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AwsIamPolicyHardcodedPartitionRule struct {
	tflint.DefaultRule
}

func NewAwsIamPolicyHardcodedPartitionRule() *AwsIamPolicyHardcodedPartitionRule {
	return &AwsIamPolicyHardcodedPartitionRule{}
}

func (r *AwsIamPolicyHardcodedPartitionRule) Name() string {
	return "aws_iam_policy_hardcoded_partition"
}

func (r *AwsIamPolicyHardcodedPartitionRule) Enabled() bool {
	return true
}

func (r *AwsIamPolicyHardcodedPartitionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *AwsIamPolicyHardcodedPartitionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *AwsIamPolicyHardcodedPartitionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedPartitions(runner, r, "IAM policy", policy, rng)
	})
}
