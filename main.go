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
				rules.NewTerraformListsTrailingCommaRule(),
				rules.NewTerraformMapTrailingCommaRule(),
				rules.NewTerraformListOrderRule(),
				rules.NewTerraformVariablesOrderRule(),
				rules.NewStyleGuideTypeVariablesExceptAnyRule(),
				rules.NewStyleGuideTypeRepetitionRule(),
				rules.NewAwsMetaHardcodedRule(),
				rules.NewAwsHardcodedIDsRule(),
				rules.NewAwsIamRolePolicyHardcodedRegionRule(),
				rules.NewAwsIamRolePolicyHardcodedPartitionRule(),
				rules.NewAwsIamPolicyHardcodedRegionRule(),
				rules.NewAwsIamPolicyHardcodedPartitionRule(),
				rules.NewAwsProviderHardcodedRegionRule(),
				rules.NewAwsServicePrincipalHardcodedRule(),
				rules.NewAwsServicePrincipalDNSSuffixRule(),
				rules.NewDaveAwsPolicyNoJsonencodeRule(),
				rules.NewDaveLabelNoTypeSubstringRule(),
				rules.NewDaveLabelSnakeRule(),
				rules.NewDaveListAlphabeticalOrderRule(),
				rules.NewDaveOutputMustBeInOutputsFileRule(),
				rules.NewDaveResourceNameKebabRule(),
				rules.NewDaveResourceNameNoTypeSubstringRule(),
				rules.NewDaveVariableAlphabeticalOrderRule(),
				rules.NewDaveVariableHasTypeRule(),
				rules.NewDaveVariableMustBeInVariablesFileRule(),
				rules.NewDaveVariableRegionRule(),
				rules.NewEosCommentsRule(),
				rules.NewEosDeathMaskRule(),
				rules.NewEosHeredocRule(),
				rules.NewEosMetaRule(),
				rules.NewEosNamingRule(),
			},
		},
	})
}
