# map_trailing_comma

Validates that the items in a multi-line map/object use trailing commas consistently. Single-line maps are ignored.

## Configuration

| Name  | Default   | Description |
| ---   | ---       | ---         |
| style | `"match"` | `"all"` requires every item to end with a comma; `"none"` forbids trailing commas; `"match"` infers the desired style from the majority of existing items. |

```hcl
rule "map_trailing_comma" {
  enabled = true
  style   = "match"
}
```

## Example

```hcl
locals {
  tags = {
    Name        = "example"
    Environment = "prod",
    Team        = "platform",
  }
}
```

With the default `"match"` style, the majority of items use a trailing comma, so the item without one is flagged.

Result:
```
$ tflint -f compact --recursive
1 issue(s) found:

main.tf:3:19: Warning - match: majority have comma (map_trailing_comma)
```

## Why

Consistent trailing commas keep multi-line map diffs minimal and uniform. This rule supports autofix.
