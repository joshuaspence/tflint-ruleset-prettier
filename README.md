# TFLint Ruleset Prettier
[![Build Status](https://github.com/joshuaspence/tflint-ruleset-prettier/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/joshuaspence/tflint-ruleset-prettier/actions)

A TFLint ruleset that enforces formatting and style conventions for Terraform code.

## Requirements

- TFLint v0.46+
- Go v1.26

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "prettier" {
  enabled = true

  version = "0.2.0"
  source  = "github.com/joshuaspence/tflint-ruleset-prettier"
}
```

## Rules

|Name|Description|Severity|Enabled By Default|Link|
| --- | --- | --- | --- | --- |
|aws_hardcoded_ids|Validates that there are no hardcoded AWS account IDs or AMI IDs|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_hardcoded_ids.md)|
|aws_iam_policy_hardcoded_partition|Validates that there are no hardcoded AWS partitions in IAM policy documents|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_iam_policy_hardcoded_partition.md)|
|aws_iam_policy_hardcoded_region|Validates that there are no hardcoded AWS regions in IAM policy documents|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_iam_policy_hardcoded_region.md)|
|aws_iam_role_policy_hardcoded_partition|Validates that there are no hardcoded AWS partitions in IAM role policy documents|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_iam_role_policy_hardcoded_partition.md)|
|aws_iam_role_policy_hardcoded_region|Validates that there are no hardcoded AWS regions in IAM role policy documents|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_iam_role_policy_hardcoded_region.md)|
|aws_meta_hardcoded|Validates that there are no hardcoded AWS regions or partitions in ARN values across all resource types|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_meta_hardcoded.md)|
|aws_provider_hardcoded_region|Validates that there are no hardcoded AWS regions in provider configuration|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_provider_hardcoded_region.md)|
|aws_service_principal_dns_suffix|Validates that service principals don't use dns_suffix interpolation|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_service_principal_dns_suffix.md)|
|aws_service_principal_hardcoded|Validates that service principals don't use hardcoded DNS suffixes (e.g., amazonaws.com)|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_service_principal_hardcoded.md)|
|aws_policy_no_jsonencode|Flags use of `jsonencode()` in policy attributes on AWS resources|NOTICE|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_policy_no_jsonencode.md)|
|output_must_be_in_outputs_file|Ensures all `output` blocks are declared in `outputs.tf`|NOTICE|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/output_must_be_in_outputs_file.md)|
|typed_variables|Ensures all variables have an explicit `type` constraint|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/typed_variables.md)|
|variable_must_be_in_variables_file|Ensures all variable blocks are declared in `variables.tf`|NOTICE|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/variable_must_be_in_variables_file.md)|
|indented_heredoc|Suggests indented heredoc syntax (`<<-`) instead of standard heredoc (`<<`)|NOTICE|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/indented_heredoc.md)|
|list_trailing_comma|Validates that the last item in a multi-line list ends with a trailing comma|NOTICE|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/list_trailing_comma.md)|
|map_trailing_comma|Validates that maps have consistent trailing commas|NOTICE|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/map_trailing_comma.md)|
|variables_order|Validates that variable blocks are sorted alphabetically|NOTICE|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/variables_order.md)|

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
