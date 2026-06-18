package project

import "fmt"

var (
  Version = "dev"
)

func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/joshuaspence/tflint-ruleset-prettier/blob/v%s/docs/rules/%s.md", Version, name)
}
