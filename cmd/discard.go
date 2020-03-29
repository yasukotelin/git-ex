package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/git-ex/usecase/git"
)

// Discard executes the removing all changes from the HEAD
func Discard(c *cli.Context) error {
	err := (&git.GitUseCase{}).Discard()
	if err != nil {
		fmt.Println("Discard failed.")
		return err
	}
	fmt.Println("Discard successed.")
	return nil
}
