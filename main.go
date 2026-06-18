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
				rules.NewTypeRepetitionRule(),
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
				rules.NewListAlphabeticalOrderRule(),
				rules.NewOutputMustBeInOutputsFileRule(),
				rules.NewResourceNameKebabRule(),
				rules.NewResourceNameNoTypeSubstringRule(),
				rules.NewTypedVariablesRule(),
				rules.NewVariableMustBeInVariablesFileRule(),
				rules.NewCommentsRule(),
				rules.NewDeathMaskRule(),
				rules.NewIndentedHeredocRule(),
				rules.NewMetaRule(),
				rules.NewNamingRule(),
			},
		},
	})
}
