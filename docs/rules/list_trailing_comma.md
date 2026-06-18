# list_trailing_comma

Recommends that the last item in a multi-line list ends with a trailing comma. Single-line lists are ignored.

## Example

```hcl
locals {
  names = [
    "Alice",
    "Bob",
    "Charlie"
  ]
}
```

Result:
```
$ tflint -f compact --recursive
1 issue(s) found:

main.tf:2:11: Notice - Last item in lists should always end with a trailing comma (list_trailing_comma)
```

## Why

A trailing comma keeps multi-line list diffs minimal: adding or reordering an element does not modify the previously-final line. This rule supports autofix.
