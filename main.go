package main

import (
	"github.com/joshuaspence/tflint-ruleset-prettier/project"
	"github.com/joshuaspence/tflint-ruleset-prettier/rules"
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
				rules.NewListOrderRule(),
				rules.NewVariablesOrderRule(),
				rules.NewTypeVariablesExceptAnyRule(),
				rules.NewTypeRepetitionRule(),
				rules.NewAwsMetaHardcodedRule(),
				rules.NewAwsHardcodedIDsRule(),
				rules.NewAwsIamRolePolicyHardcodedRegionRule(),
				rules.NewAwsIamRolePolicyHardcodedPartitionRule(),
				rules.NewAwsIamPolicyHardcodedRegionRule(),
				rules.NewAwsIamPolicyHardcodedPartitionRule(),
				rules.NewAwsProviderHardcodedRegionRule(),
				rules.NewAwsServicePrincipalHardcodedRule(),
				rules.NewAwsServicePrincipalDNSSuffixRule(),
				rules.NewAwsPolicyNoJsonencodeRule(),
				rules.NewLabelNoTypeSubstringRule(),
				rules.NewLabelSnakeRule(),
				rules.NewListAlphabeticalOrderRule(),
				rules.NewOutputMustBeInOutputsFileRule(),
				rules.NewResourceNameKebabRule(),
				rules.NewResourceNameNoTypeSubstringRule(),
				rules.NewVariableAlphabeticalOrderRule(),
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
