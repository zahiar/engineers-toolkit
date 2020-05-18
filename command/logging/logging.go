package logging

import (
	"./create"
	"./update"
	"github.com/urfave/cli"
)

func GetCommands() *cli.Command {
	return &cli.Command{
		Name:        "logging",
		Description: "All things to do with Graylog logging!",
		Subcommands: []*cli.Command{
			create.GetCommand(),
			update.GetCommand(),
		},
	}
}
