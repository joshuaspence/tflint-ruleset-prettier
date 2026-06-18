package aws

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type IamRolePolicyHardcodedPartitionRule struct {
	tflint.DefaultRule
}

func NewIamRolePolicyHardcodedPartitionRule() *IamRolePolicyHardcodedPartitionRule {
	return &IamRolePolicyHardcodedPartitionRule{}
}

func (r *IamRolePolicyHardcodedPartitionRule) Name() string {
	return "aws_iam_role_policy_hardcoded_partition"
}

func (r *IamRolePolicyHardcodedPartitionRule) Enabled() bool {
	return true
}

func (r *IamRolePolicyHardcodedPartitionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *IamRolePolicyHardcodedPartitionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IamRolePolicyHardcodedPartitionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_role_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedPartitions(runner, r, "IAM role policy", policy, rng)
	})
}
