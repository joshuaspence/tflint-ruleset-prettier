---
title: AWS Provider Hardcoded Regions
description: Checks AWS provider configurations for hardcoded regions.
ruleName: aws_provider_hardcoded_region
---

**Rule:** `aws_provider_hardcoded_region`

This rule checks AWS provider configurations for hardcoded regions. It detects:

- Hardcoded regions in provider `region` attribute
- Hardcoded regions in `assume_role` ARNs

## Example violations

```hcl
provider "aws" {
  region = "us-east-1"  # ❌ Hardcoded region
}
```

## Recommended fix

```hcl
provider "aws" {
  region = var.aws_region  # ✅ Use variables
}
```

## Enabling this rule

This rule is **enabled by default** when you install the prettier plugin. No additional configuration is needed.

If you want to disable this rule, add it to your `.tflint.hcl`:

```hcl
rule "aws_provider_hardcoded_region" {
  enabled = false
}
```
