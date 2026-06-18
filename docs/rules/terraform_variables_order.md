# terraform_variables_order

Recommends alphabetical order for variables.
Autofix requires a file to only have variable blocks. If it contains other types of blocks, the fix will not be applied.

## Configuration

```hcl
rule "terraform_variables_order" {
  enabled = true
  group_required = false # Set to true if you want required variables to be sorted separately from the optional ones
}
```


## Example

```hcl

variable "b" {
  type = string
}

variable "a" {
  type = string
}

```

Result:
```
$ tflint --recursive
1 issue(s) found:

Notice: Recommended variables order:
variable "a" {
  type = string
}

variable "b" {
  type = string
}
 (terraform_variables_order)

  on main.tf line 2:
   2: variable "b" {

```
