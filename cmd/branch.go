package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func Branch(c *cli.Context) error {
	branchs, err := gitUseCase.FetchBranch()
	if err != nil {
		return err
	}
	defaultChoices := []string{cancel}
	messages := append(defaultChoices, branchs...)
	selectedIndex := len(defaultChoices)

	prompt := promptui.Select{
		Label:     "Select the file you want to remove",
		Items:     messages,
		Size:      15,
		CursorPos: selectedIndex,
	}
	_, r, err := prompt.Run()

	if err != nil {
		return err
	}

	switch r {
	case cancel:
		return nil
	default:
		return gitUseCase.Checkout(r)
	}
}
