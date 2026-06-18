# comments

Enforces comment style guidelines: no end-of-line comments, no jammed comments, and a maximum line length.

## Sub-rules

| Sub-rule | Identifies | Default |
|----------|------------|---------|
| `eol` | End-of-line comments. | `true` |
| `jammed` | Comments without space after marker. | `true` |
| `length` | Comments exceeding line length. | `true` (column 80) |

## Example

```hcl
#This is a jammed comment
//This is also jammed

# This comment is way too long and exceeds the configured column limit which defaults to 80 characters so it will trigger a warning.

resource "aws_instance" "example" {
  ami = var.ami # EOL comment
}
```

```
$ tflint
4 issue(s) found:

Warning: Avoid jammed comment ('#This ...'). (comments)

  on main.tf line 1:
   1: #This is a jammed comment

Warning: Avoid jammed comment ('//Thi ...'). (comments)

  on main.tf line 2:
   2: //This is also jammed

Warning: Wrap comment at column 80 (currently 126). (comments)

  on main.tf line 4:
   4: # This comment is way too long and exceeds the configured column limit which defaults to 80 characters so it will trigger a warning.

Warning: Avoid EOL comments. (comments)

  on main.tf line 7:
   7:   ami = var.ami # EOL comment

Reference: https://github.com/joshuaspence/tflint-ruleset-prettier/blob/v0.2.0/docs/rules/comments.md
```

## Why

Readable comments improve code maintainability. "Jammed" comments (without a space after the `#` or `//` marker) are harder to read. Nothing disrupts readability more than a comment that disappears off the right side of the editor pane or wraps unnaturally. End-of-line comments can clutter code.

## How To Fix

Add a space after the comment marker, break long comments across multiple lines, or move end-of-line comments to their own line.

```hcl
# This is a properly spaced comment.

# This comment is broken across multiple lines so it does not
# exceed the configured column limit.

resource "aws_instance" "example" {
  # This comment is on its own line.
  ami = var.ami
}
```

The rule can be ignored with:

```hcl
# tflint-ignore: comments
#This jammed comment is intentional
resource "aws_instance" "example" {
  ami = var.ami
}
```

## Configuration

This rule is enabled by default and can be disabled with:

```hcl
rule "comments" {
  enabled = false
}
```

Configure sub-rules individually:

```hcl
rule "comments" {
  eol    = false  # Allow EOL comments
  jammed = false  # Allow jammed comments
  length {
    column    = 100    # Set max column to 100
    allow_url = false  # Don't allow URLs to exceed limit
  }
  level = "error"  # Change severity to error
}
```
