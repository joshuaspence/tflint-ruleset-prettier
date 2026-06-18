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
|aws_hardcoded_ids|Validates that there are no hardcoded AWS account IDs or AMI IDs|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_hardcoded_ids.md)|
|aws_iam_policy_hardcoded_partition|Validates that there are no hardcoded AWS partitions in IAM policy documents|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_iam_policy_hardcoded_partition.md)|
|aws_iam_policy_hardcoded_region|Validates that there are no hardcoded AWS regions in IAM policy documents|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_iam_policy_hardcoded_region.md)|
|aws_iam_role_policy_hardcoded_partition|Validates that there are no hardcoded AWS partitions in IAM role policy documents|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_iam_role_policy_hardcoded_partition.md)|
|aws_iam_role_policy_hardcoded_region|Validates that there are no hardcoded AWS regions in IAM role policy documents|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_iam_role_policy_hardcoded_region.md)|
|aws_meta_hardcoded|Validates that there are no hardcoded AWS regions or partitions in ARN values across all resource types|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_meta_hardcoded.md)|
|aws_provider_hardcoded_region|Validates that there are no hardcoded AWS regions in provider configuration|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_provider_hardcoded_region.md)|
|aws_service_principal_dns_suffix|Validates that service principals don't use dns_suffix interpolation|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_service_principal_dns_suffix.md)|
|aws_service_principal_hardcoded|Validates that service principals don't use hardcoded DNS suffixes (e.g., amazonaws.com)|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/aws_service_principal_hardcoded.md)|
|dave_aws_policy_no_jsonencode|Flags use of `jsonencode()` in policy attributes on AWS resources|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_aws_policy_no_jsonencode.md)|
|dave_label_no_type_substring|Prevents labels from containing words that already appear in the resource type|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_label_no_type_substring.md)|
|dave_label_snake|Ensures labels on resource, data, ephemeral, module, output, and variable blocks use snake_case|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_label_snake.md)|
|dave_list_alphabetical_order|Ensures the string elements of a list literal are sorted alphabetically, for a configurable set of attribute names|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_list_alphabetical_order.md)|
|dave_output_must_be_in_outputs_file|Ensures all `output` blocks are declared in `outputs.tf`|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_output_must_be_in_outputs_file.md)|
|dave_resource_name_kebab|Ensures `name` and `name_prefix` attributes use kebab-case|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_resource_name_kebab.md)|
|dave_resource_name_no_type_substring|Prevents `name` and `name_prefix` attributes from containing words that appear in the resource type|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_resource_name_no_type_substring.md)|
|dave_variable_alphabetical_order|Ensures variables within each file are sorted alphabetically by name|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_variable_alphabetical_order.md)|
|dave_variable_has_type|Ensures all variables have an explicit `type` constraint|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_variable_has_type.md)|
|dave_variable_must_be_in_variables_file|Ensures all variable blocks are declared in `variables.tf`|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_variable_must_be_in_variables_file.md)|
|dave_variable_region|Flags any variable named `region`|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/dave_variable_region.md)|
|eos_comments|Enforces comment style: no end-of-line comments, no jammed comments, and a maximum line length|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/eos_comments.md)|
|eos_death_mask|Identifies commented-out blocks of code left behind ("death masks")|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/eos_death_mask.md)|
|eos_heredoc|Suggests indented heredoc syntax (`<<-`) and optionally flags `EOF` delimiters|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/eos_heredoc.md)|
|eos_meta|Enforces Terraform meta-argument syntax conventions|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/eos_meta.md)|
|eos_naming|Enforces naming conventions on Terraform blocks and locals|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/eos_naming.md)|
|style_guide_type_repetition|Validates that resource and data source names do not repeat a word from their type|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/style_guide_type_repetition.md)|
|style_guide_typed_variables_except_any|Validates that variables do not use `any` as their type, including inside composite types|WARNING|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/style_guide_typed_variables_except_any.md)|
|terraform_list_order|Validates that list items are sorted alphabetically|NOTICE|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/terraform_list_order.md)|
|terraform_lists_trailing_comma|Validates that the last item in a multi-line list ends with a trailing comma|WARNING|✅||
|terraform_map_trailing_comma|Validates that maps have consistent trailing commas|WARNING|✅||
|terraform_variables_order|Validates that variable blocks are sorted alphabetically|NOTICE|✅|[docs](https://github.com/joshuaspence/tflint-ruleset-prettier/blob/main/docs/rules/terraform_variables_order.md)|

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
