package command

import (
	"strings"
)

type PlanCommand struct {
	Meta
}

func (c *PlanCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *PlanCommand) Synopsis() string {
	return ""
}

func (c *PlanCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
