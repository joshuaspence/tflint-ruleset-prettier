# indented_heredoc

Identifies usage of standard heredoc syntax (`<<`) and suggests indented heredoc syntax (`<<-`).

## Example

```hcl
# Bad.
resource "terraform_data" "example" {
  input = <<EOF
#!/bin/bash
echo "hello"
EOF
}

# Good.
resource "terraform_data" "example" {
  input = <<-SHELL
    #!/bin/bash
    echo "hello"
  SHELL
}
```

```
$ tflint
1 issue(s) found:

Notice: Avoid standard heredoc (<<). Use indented (<<-) instead. (indented_heredoc)

  on main.tf line 3:
   3:   input = <<EOF

Reference: https://github.com/joshuaspence/tflint-ruleset-prettier/blob/v0.2.0/docs/rules/indented_heredoc.md
```

## Why

Standard heredocs (`<<`) require the content to be left-aligned, which breaks the visual indentation hierarchy of the Terraform code. Indented heredocs (`<<-`) allow the content to be indented relative to the surrounding code, improving readability and maintainability.

## How To Fix

Change `<<` to `<<-` and indent the content to match the surrounding code.

```hcl
resource "terraform_data" "example" {
  input = <<-SHELL
    #!/bin/bash
    echo "hello"
  SHELL
}
```

The rule can be ignored with:

```hcl
# tflint-ignore: indented_heredoc
resource "terraform_data" "example" {
  input = <<EOF
#!/bin/bash
echo "hello"
EOF
}
```

## Configuration

This rule is enabled by default and can be disabled with:

```hcl
rule "indented_heredoc" {
  enabled = false
}
```
