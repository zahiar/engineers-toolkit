package pipeline

import (
	"./create"
	"github.com/urfave/cli"
)

func GetCommands() *cli.Command {
	return &cli.Command{
		Name:        "pipeline",
		Description: "All things to do with Jenkins Pipelines!",
		Subcommands: []*cli.Command{
			create.GetCommand(),
		},
	}
}
