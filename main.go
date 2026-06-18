package main

import (
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/joshuaspence/tflint-ruleset-prettier/rules"
	"github.com/joshuaspence/tflint-ruleset-prettier/rules/aws"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "prettier",
			Version: project.Version,
			Rules: []tflint.Rule{
				aws.NewHardcodedIDsRule(),
				aws.NewIamPolicyHardcodedPartitionRule(),
				aws.NewIamPolicyHardcodedRegionRule(),
				aws.NewIamRolePolicyHardcodedPartitionRule(),
				aws.NewIamRolePolicyHardcodedRegionRule(),
				aws.NewMetaHardcodedRule(),
				aws.NewPolicyNoJsonencodeRule(),
				aws.NewProviderHardcodedRegionRule(),
				aws.NewServicePrincipalDNSSuffixRule(),
				aws.NewServicePrincipalHardcodedRule(),
				rules.NewIndentedHeredocRule(),
				rules.NewListTrailingCommaRule(),
				rules.NewMapTrailingCommaRule(),
				rules.NewOutputMustBeInOutputsFileRule(),
				rules.NewTypedVariablesRule(),
				rules.NewVariableMustBeInVariablesFileRule(),
				rules.NewVariablesOrderRule(),
			},
		},
	})
}
