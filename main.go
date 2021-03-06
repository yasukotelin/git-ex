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
		Version: "0.6.0",
		Commands: []*cli.Command{
			&cli.Command{
				Name:   "discard",
				Usage:  "Executes the removing all changes from the HEAD that include untracked files",
				Action: cmd.Discard,
			},
			&cli.Command{
				Name:   "stage",
				Usage:  "Executes the staging files with selecter.",
				Action: cmd.Stage,
			},
			&cli.Command{
				Name:   "unstage",
				Usage:  "Executes the unstaging files with selecter.",
				Action: cmd.UnStage,
			},
			&cli.Command{
				Name:  "diff",
				Usage: "Executes the diff with selecter.",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "stage",
						Value: false,
						Usage: "true is showing stage files, false is showing unstaging files",
					},
				},
				Action: cmd.Diff,
			},
			&cli.Command{
				Name:  "branch",
				Usage: "Checkout a branch with selecter",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
						Value:   false,
						Usage:   "true is showing local and remote branchs, false is local only",
					},
				},
				Action: cmd.Branch,
			},
		},
		Action: cli.ShowAppHelp,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
