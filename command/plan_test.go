package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestPlanCommand_implement(t *testing.T) {
	var _ cli.Command = &PlanCommand{}
}
