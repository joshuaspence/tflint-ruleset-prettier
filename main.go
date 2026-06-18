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
				rules.NewListTrailingCommaRule(),
				rules.NewMapTrailingCommaRule(),
				rules.NewVariablesOrderRule(),
				aws.NewMetaHardcodedRule(),
				aws.NewHardcodedIDsRule(),
				aws.NewIamRolePolicyHardcodedRegionRule(),
				aws.NewIamRolePolicyHardcodedPartitionRule(),
				aws.NewIamPolicyHardcodedRegionRule(),
				aws.NewIamPolicyHardcodedPartitionRule(),
				aws.NewProviderHardcodedRegionRule(),
				aws.NewServicePrincipalHardcodedRule(),
				aws.NewServicePrincipalDNSSuffixRule(),
				aws.NewPolicyNoJsonencodeRule(),
				rules.NewOutputMustBeInOutputsFileRule(),
				rules.NewTypedVariablesRule(),
				rules.NewVariableMustBeInVariablesFileRule(),
				rules.NewIndentedHeredocRule(),
			},
		},
	})
}
