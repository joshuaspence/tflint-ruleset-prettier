package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-template/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "template",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewAwsInstanceExampleTypeRule(),
				rules.NewAwsS3BucketExampleLifecycleRule(),
				rules.NewGoogleComputeSSLPolicyRule(),
				rules.NewTerraformBackendTypeRule(),
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
