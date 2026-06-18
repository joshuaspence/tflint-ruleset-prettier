package aws

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type IamRolePolicyHardcodedRegionRule struct {
	tflint.DefaultRule
}

func NewIamRolePolicyHardcodedRegionRule() *IamRolePolicyHardcodedRegionRule {
	return &IamRolePolicyHardcodedRegionRule{}
}

func (r *IamRolePolicyHardcodedRegionRule) Name() string {
	return "aws_iam_role_policy_hardcoded_region"
}

func (r *IamRolePolicyHardcodedRegionRule) Enabled() bool {
	return true
}

func (r *IamRolePolicyHardcodedRegionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *IamRolePolicyHardcodedRegionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IamRolePolicyHardcodedRegionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_role_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedRegions(runner, r, "IAM role policy", policy, rng)
	})
}
