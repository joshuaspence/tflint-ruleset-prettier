package aws

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type IamPolicyHardcodedRegionRule struct {
	tflint.DefaultRule
}

func NewIamPolicyHardcodedRegionRule() *IamPolicyHardcodedRegionRule {
	return &IamPolicyHardcodedRegionRule{}
}

func (r *IamPolicyHardcodedRegionRule) Name() string {
	return "aws_iam_policy_hardcoded_region"
}

func (r *IamPolicyHardcodedRegionRule) Enabled() bool {
	return true
}

func (r *IamPolicyHardcodedRegionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *IamPolicyHardcodedRegionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IamPolicyHardcodedRegionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedRegions(runner, r, "IAM policy", policy, rng)
	})
}
