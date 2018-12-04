package main

import (
	"github.com/mitchellh/cli"
	"github.com/stakagi/kongger/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"plan": func() (cli.Command, error) {
			return &command.PlanCommand{
				Meta: *meta,
			}, nil
		},
		"apply": func() (cli.Command, error) {
			return &command.ApplyCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
