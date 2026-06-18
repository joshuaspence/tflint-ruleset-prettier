package aws

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AwsIamPolicyHardcodedRegionRule struct {
	tflint.DefaultRule
}

func NewAwsIamPolicyHardcodedRegionRule() *AwsIamPolicyHardcodedRegionRule {
	return &AwsIamPolicyHardcodedRegionRule{}
}

func (r *AwsIamPolicyHardcodedRegionRule) Name() string {
	return "aws_iam_policy_hardcoded_region"
}

func (r *AwsIamPolicyHardcodedRegionRule) Enabled() bool {
	return true
}

func (r *AwsIamPolicyHardcodedRegionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *AwsIamPolicyHardcodedRegionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *AwsIamPolicyHardcodedRegionRule) Check(runner tflint.Runner) error {
	return checkIamPolicyAttributes(runner, "aws_iam_policy", func(policy string, rng hcl.Range) error {
		return checkIamPolicyForHardcodedRegions(runner, r, "IAM policy", policy, rng)
	})
}
