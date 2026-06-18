---
title: IAM Policy Hardcoded Regions
description: Detects hardcoded AWS regions in aws_iam_policy resources.
ruleName: aws_iam_policy_hardcoded_region
---

**Rule:** `aws_iam_policy_hardcoded_region`

This rule checks `aws_iam_policy` resources for hardcoded AWS regions in policy documents. Similar to the role policy rule, it examines:

- Hardcoded regions in ARNs within policy statements
- Direct region references in policy JSON

## Example violations

```hcl
resource "aws_iam_policy" "bad" {
  policy = jsonencode({
    Statement = [{
      Effect   = "Allow"
      Action   = "lambda:InvokeFunction"
      Resource = "arn:aws:lambda:eu-west-1:123456789012:function:*"  # ❌ Hardcoded region
    }]
  })
}
```

## Recommended fix

```hcl
resource "aws_iam_policy" "good" {
  policy = jsonencode({
    Statement = [{
      Effect   = "Allow"
      Action   = "lambda:InvokeFunction"
      Resource = "arn:aws:lambda:${data.aws_region.current.name}:123456789012:function:*"  # ✅ Dynamic region
    }]
  })
}
```

## Enabling this rule

This rule is **enabled by default** when you install the prettier plugin. No additional configuration is needed.

If you want to disable this rule, add it to your `.tflint.hcl`:

```hcl
rule "aws_iam_policy_hardcoded_region" {
  enabled = false
}
```
