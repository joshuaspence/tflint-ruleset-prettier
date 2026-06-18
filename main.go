package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/joshuaspence/tflint-ruleset-prettier/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "prettier",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewTerraformListsTrailingCommaRule(),
				rules.NewTerraformMapTrailingCommaRule(),
				rules.NewTerraformListOrderRule(),
				rules.NewTerraformVariablesOrderRule(),
				rules.NewStyleGuideTypeVariablesExceptAnyRule(),
				rules.NewAwsMetaHardcodedRule(),
				rules.NewAwsHardcodedIDsRule(),
				rules.NewAwsIamRolePolicyHardcodedRegionRule(),
				rules.NewAwsIamRolePolicyHardcodedPartitionRule(),
				rules.NewAwsIamPolicyHardcodedRegionRule(),
				rules.NewAwsIamPolicyHardcodedPartitionRule(),
				rules.NewAwsProviderHardcodedRegionRule(),
				rules.NewAwsServicePrincipalHardcodedRule(),
				rules.NewAwsServicePrincipalDNSSuffixRule(),
			},
		},
	})
}
