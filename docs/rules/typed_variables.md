# typed_variables

Ensures all variables have an explicit `type` constraint.

## Why

Without a type constraint, Terraform cannot catch invalid inputs at plan time, pushing errors to apply where they are more expensive and disruptive.

## Example

```hcl
# ❌ Invalid — missing type
variable "environment" {
  description = "The deployment environment."
}

# ✅ Valid
variable "environment" {
  description = "The deployment environment."
  type        = string
}
```

```
$ tflint
1 issue(s) found:

Warning: Variable "environment" is missing a type constraint. (typed_variables)

  on variables.tf line 2:
   2: variable "environment" {

Reference: https://github.com/joshuaspence/tflint-ruleset-prettier/blob/v0.2.0/docs/rules/typed_variables.md
```

## How To Fix

Add an explicit `type` to the variable. See the [Terraform type constraints documentation](https://developer.hashicorp.com/terraform/language/values/variables#type-constraints) for the available types.

## Configuration

This rule is enabled by default and can be disabled with:

```hcl
rule "typed_variables" {
  enabled = false
}
```
