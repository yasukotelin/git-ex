package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Discard executes the removing all changes from the HEAD
func Discard(c *cli.Context) error {
	err := gitUseCase.Discard()
	if err != nil {
		fmt.Println("Discard failed.")
		return err
	}
	fmt.Println("Discard successed.")
	return nil
}
