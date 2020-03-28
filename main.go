package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/git-ex/cmd"
)

func main() {
	app := &cli.App{
		Name:    "git-ex",
		Usage:   "git-ex is a subcommand that extends Git",
		Version: "0.1.0",
		Commands: []*cli.Command{
			&cli.Command{
				Name:   "discard",
				Usage:  "Executes the removing all changes from the HEAD that include untracked files",
				Action: cmd.Discard,
			},
		},
		Action: cli.ShowAppHelp,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
