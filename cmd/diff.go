package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/git-ex/entity"
)

func Diff(c *cli.Context) error {
	var statusFiles []entity.GitStatusFile
	var err error

	isStage := c.Bool("stage")

	if isStage {
		statusFiles, err = gitUseCase.FetchStage()
	} else {
		statusFiles, err = gitUseCase.FetchUnStage()
	}
	if err != nil {
		return err
	}

	defaultChoices := []string{all, cancel}
	messages := append(defaultChoices, getValues(statusFiles)...)

	prompt := promptui.Select{
		Label:     "Select the file you want to show diff",
		Items:     messages,
		Size:      15,
		CursorPos: len(defaultChoices),
	}
	i, r, err := prompt.Run()

	if err != nil {
		return err
	}

	switch r {
	case all:
		return gitUseCase.DiffAll(isStage)
	case cancel:
		return nil
	default:
		return gitUseCase.Diff(isStage, statusFiles[i-len(defaultChoices)].Path)
	}
}
