package cmd

import (
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func Branch(c *cli.Context) error {
	isAll := c.Bool("all")

	var branchs []string
	var err error
	if isAll {
		branchs, err = gitUseCase.FetchAllBranch()
	} else {
		branchs, err = gitUseCase.FetchBranch()
	}
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
		{
			branch := r
			if strings.HasPrefix(r, "remotes") {
				// remotes/{リモート名}/{ブランチ名}からブランチ名だけを取り出す
				branch = strings.SplitN(r, "/", 3)[2]
			}
			return gitUseCase.Checkout(branch)
		}
	}
}
