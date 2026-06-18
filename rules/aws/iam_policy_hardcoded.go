package aws

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/myerscode/tflint-ruleset-aws-meta/rules/awsmeta"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// checkIamPolicyAttributes evaluates the "policy" attribute of every resource of the given type and invokes check for
// each policy string that can be statically evaluated.
func checkIamPolicyAttributes(runner tflint.Runner, resourceType string, check func(policy string, rng hcl.Range) error) error {
	resources, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "policy"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attr, exists := resource.Body.Attributes["policy"]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attr.Expr, func(policy string) error {
			return check(policy, attr.Expr.Range())
		}, nil)
		if err != nil && !isExpectedEvalError(err) {
			return err
		}
	}

	return nil
}

// iamPolicyScanTarget normalizes a policy string for scanning. If the policy is valid JSON it is re-marshaled
// (canonicalizing whitespace) and the label is suffixed with " document" to match the message used for structured
// policies.
func iamPolicyScanTarget(policy, policyLabel string) (target, label string, ok bool) {
	var policyDoc map[string]interface{}
	if err := json.Unmarshal([]byte(policy), &policyDoc); err != nil {
		return policy, policyLabel, true
	}

	docBytes, err := json.Marshal(policyDoc)
	if err != nil {
		// Skip if we can't marshal back.
		return "", "", false
	}

	return string(docBytes), policyLabel + " document", true
}

// checkIamPolicyForHardcodedRegions emits an issue for each hardcoded region found either as a bare region string or
// within an ARN. policyLabel is the human-readable resource description used in messages (e.g. "IAM policy").
func checkIamPolicyForHardcodedRegions(runner tflint.Runner, rule tflint.Rule, policyLabel, policy string, rng hcl.Range) error {
	regionInStringPattern := awsmeta.GetRegionInStringPattern()
	arnRegionPattern := awsmeta.GetARNRegionPattern()

	target, label, ok := iamPolicyScanTarget(policy, policyLabel)
	if !ok {
		return nil
	}

	for _, match := range regionInStringPattern.FindAllString(target, -1) {
		if err := runner.EmitIssue(
			rule,
			fmt.Sprintf("Hardcoded AWS region '%s' found in %s. Consider using variables or data.aws_region.current.name", match, label),
			rng,
		); err != nil {
			return err
		}
	}

	for _, match := range arnRegionPattern.FindAllStringSubmatch(target, -1) {
		if len(match) > 1 {
			region := match[1]
			if err := runner.EmitIssue(
				rule,
				fmt.Sprintf("Hardcoded AWS region '%s' found in ARN within %s. Consider using variables or data.aws_region.current.name", region, label),
				rng,
			); err != nil {
				return err
			}
		}
	}

	return nil
}

// checkIamPolicyForHardcodedPartitions emits an issue for each hardcoded partition found within an ARN. policyLabel is
// the human-readable resource description used in messages (e.g. "IAM policy").
func checkIamPolicyForHardcodedPartitions(runner tflint.Runner, rule tflint.Rule, policyLabel, policy string, rng hcl.Range) error {
	arnPartitionPattern := awsmeta.GetPartitionPattern()

	target, label, ok := iamPolicyScanTarget(policy, policyLabel)
	if !ok {
		return nil
	}

	for _, match := range arnPartitionPattern.FindAllStringSubmatch(target, -1) {
		if len(match) > 1 {
			partition := match[1]
			if err := runner.EmitIssue(
				rule,
				fmt.Sprintf("Hardcoded AWS partition '%s' found in ARN within %s. Consider using data.aws_partition.current.partition", partition, label),
				rng,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
