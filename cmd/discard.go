package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/git-ex/usecase/git"
)

// Discard executes the removing all changes from the HEAD
func Discard(c *cli.Context) error {
	return (&git.GitUseCase{}).Discard()
}
