package aws

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type IamPolicyHardcodedPartitionRule struct {
	tflint.DefaultRule
}

func NewIamPolicyHardcodedPartitionRule() *IamPolicyHardcodedPartitionRule {
	return &IamPolicyHardcodedPartitionRule{}
}

func (r *IamPolicyHardcodedPartitionRule) Name() string {
	return "aws_iam_policy_hardcoded_partition"
}

func (r *IamPolicyHardcodedPartitionRule) Enabled() bool {
	return true
}

func (r *IamPolicyHardcodedPartitionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *IamPolicyHardcodedPartitionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IamPolicyHardcodedPartitionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedPartitions(runner, r, "IAM policy", policy, rng)
	})
}
