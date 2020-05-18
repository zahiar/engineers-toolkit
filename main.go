package main

import (
	"./command/logging"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "Engineer's Toolkit",
		Description: "Engineer's Toolkit - to make some of the more \"opsy\" tasks a lot easier...",
		Version:     "v0.0.1",
		Commands: []*cli.Command{
			logging.GetCommands(),
		},
	}

	_ = app.Run(os.Args)
}
