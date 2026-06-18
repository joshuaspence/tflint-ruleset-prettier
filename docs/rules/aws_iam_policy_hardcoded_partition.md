---
title: IAM Policy Hardcoded Partitions
description: Detects hardcoded AWS partitions in aws_iam_policy resources.
ruleName: aws_iam_policy_hardcoded_partition
---

**Rule:** `aws_iam_policy_hardcoded_partition`

This rule checks `aws_iam_policy` resources for hardcoded AWS partitions in policy documents. It detects:

- Hardcoded partitions in ARNs within policy statements

## Example violations

```hcl
resource "aws_iam_policy" "bad" {
  policy = jsonencode({
    Statement = [{
      Effect   = "Allow"
      Action   = "sqs:*"
      Resource = "arn:aws-us-gov:sqs:*:*:*"  # ❌ Hardcoded partition
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
      Action   = "sqs:*"
      Resource = "arn:${data.aws_partition.current.partition}:sqs:*:*:*"  # ✅ Dynamic partition
    }]
  })
}
```

## Enabling this rule

This rule is **enabled by default** when you install the prettier plugin. No additional configuration is needed.

If you want to disable this rule, add it to your `.tflint.hcl`:

```hcl
rule "aws_iam_policy_hardcoded_partition" {
  enabled = false
}
```
