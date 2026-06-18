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
				rules.NewListsTrailingCommaRule(),
				rules.NewMapTrailingCommaRule(),
				rules.NewVariablesOrderRule(),
				rules.NewTypeVariablesExceptAnyRule(),
				rules.NewTypeRepetitionRule(),
				aws.NewAwsMetaHardcodedRule(),
				aws.NewAwsHardcodedIDsRule(),
				aws.NewAwsIamRolePolicyHardcodedRegionRule(),
				aws.NewAwsIamRolePolicyHardcodedPartitionRule(),
				aws.NewAwsIamPolicyHardcodedRegionRule(),
				aws.NewAwsIamPolicyHardcodedPartitionRule(),
				aws.NewAwsProviderHardcodedRegionRule(),
				aws.NewAwsServicePrincipalHardcodedRule(),
				aws.NewAwsServicePrincipalDNSSuffixRule(),
				aws.NewAwsPolicyNoJsonencodeRule(),
				rules.NewListAlphabeticalOrderRule(),
				rules.NewOutputMustBeInOutputsFileRule(),
				rules.NewResourceNameKebabRule(),
				rules.NewResourceNameNoTypeSubstringRule(),
				rules.NewVariableHasTypeRule(),
				rules.NewVariableMustBeInVariablesFileRule(),
				rules.NewVariableRegionRule(),
				rules.NewCommentsRule(),
				rules.NewDeathMaskRule(),
				rules.NewHeredocRule(),
				rules.NewMetaRule(),
				rules.NewNamingRule(),
			},
		},
	})
}
