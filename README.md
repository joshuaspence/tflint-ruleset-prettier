# TFLint Ruleset Prettier
[![Build Status](https://github.com/joshuaspence/tflint-ruleset-prettier/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/joshuaspence/tflint-ruleset-prettier/actions)

A TFLint ruleset that enforces formatting and style conventions for Terraform code.

## Requirements

- TFLint v0.46+
- Go v1.26

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "prettier" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/joshuaspence/tflint-ruleset-prettier"
}
```

## Rules

|Name|Description|Severity|Enabled By Default|Link|
| --- | --- | --- | --- | --- |
|terraform_lists_trailing_comma|Validates that the last item in a multi-line list ends with a trailing comma|WARNING|✅||
|terraform_map_trailing_comma|Validates that maps have consistent trailing commas|WARNING|✅||
|terraform_list_order|Validates that list items are sorted alphabetically|NOTICE|✅||
|terraform_variables_order|Validates that variable blocks are sorted alphabetically|NOTICE|✅||
|style_guide_typed_variables_except_any|Validates that variables do not use `any` as their type, including inside composite types|WARNING|✅||
|aws_meta_hardcoded|Validates that there are no hardcoded AWS regions or partitions in ARN values across all resource types|WARNING|✅|[docs](https://myerscode.github.io/tflint-ruleset-aws-meta/rules/aws_meta_hardcoded)|
|aws_hardcoded_ids|Validates that there are no hardcoded AWS account IDs or AMI IDs|WARNING|✅|[docs](https://myerscode.github.io/tflint-ruleset-aws-meta/rules/aws_hardcoded_ids)|
|aws_iam_role_policy_hardcoded_region|Validates that there are no hardcoded AWS regions in IAM role policy documents|WARNING|✅|[docs](https://myerscode.github.io/tflint-ruleset-aws-meta/rules/aws_iam_role_policy_hardcoded_region)|
|aws_iam_role_policy_hardcoded_partition|Validates that there are no hardcoded AWS partitions in IAM role policy documents|WARNING|✅|[docs](https://myerscode.github.io/tflint-ruleset-aws-meta/rules/aws_iam_role_policy_hardcoded_partition)|
|aws_iam_policy_hardcoded_region|Validates that there are no hardcoded AWS regions in IAM policy documents|WARNING|✅|[docs](https://myerscode.github.io/tflint-ruleset-aws-meta/rules/aws_iam_policy_hardcoded_region)|
|aws_iam_policy_hardcoded_partition|Validates that there are no hardcoded AWS partitions in IAM policy documents|WARNING|✅|[docs](https://myerscode.github.io/tflint-ruleset-aws-meta/rules/aws_iam_policy_hardcoded_partition)|
|aws_provider_hardcoded_region|Validates that there are no hardcoded AWS regions in provider configuration|WARNING|✅|[docs](https://myerscode.github.io/tflint-ruleset-aws-meta/rules/aws_provider_hardcoded_region)|
|aws_service_principal_hardcoded|Validates that service principals don't use hardcoded DNS suffixes (e.g., amazonaws.com)|WARNING|✅|[docs](https://myerscode.github.io/tflint-ruleset-aws-meta/rules/aws_service_principal_hardcoded)|
|aws_service_principal_dns_suffix|Validates that service principals don't use dns_suffix interpolation|WARNING|✅|[docs](https://myerscode.github.io/tflint-ruleset-aws-meta/rules/aws_service_principal_dns_suffix)|

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "prettier" {
  enabled = true
}
EOS
$ tflint
```
