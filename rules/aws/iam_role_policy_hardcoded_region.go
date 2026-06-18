package aws

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AwsIamRolePolicyHardcodedRegionRule struct {
	tflint.DefaultRule
}

func NewAwsIamRolePolicyHardcodedRegionRule() *AwsIamRolePolicyHardcodedRegionRule {
	return &AwsIamRolePolicyHardcodedRegionRule{}
}

func (r *AwsIamRolePolicyHardcodedRegionRule) Name() string {
	return "aws_iam_role_policy_hardcoded_region"
}

func (r *AwsIamRolePolicyHardcodedRegionRule) Enabled() bool {
	return true
}

func (r *AwsIamRolePolicyHardcodedRegionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *AwsIamRolePolicyHardcodedRegionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *AwsIamRolePolicyHardcodedRegionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_role_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedRegions(runner, r, "IAM role policy", policy, rng)
	})
}
