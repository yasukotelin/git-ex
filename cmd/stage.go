package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/git-ex/entity"
	"github.com/yasukotelin/git-ex/usecase/git"
	"github.com/yasukotelin/git-ex/util"
)

const all = "...(All)"
const cancel = "...(Cancel)"
const apply = "...(Apply)"
const defaultChoices = 3

var gitUseCase = &git.GitUseCase{}
var stringSliceUtil = &util.StringSliceUtil{}
var gitStatusFileUtil = &util.GitStatusFileSliceUtil{}

// Stage stages the files with selecter
func Stage(c *cli.Context) error {
	unStages, err := gitUseCase.FetchUnStagePath()
	if err != nil {
		return err
	}

	messages := append([]string{cancel, all, apply}, getValues(unStages)...)
	selected := make([]entity.GitStatusFile, 0, len(unStages))
	selectedIndex := defaultChoices

ASK:
	for {
		if len(messages) == defaultChoices {
			// 選択肢が無くなったら確定
			break
		}

		prompt := promptui.Select{
			Label:     "Select the file you want to stage",
			Items:     messages,
			Size:      15,
			CursorPos: selectedIndex,
		}
		i, r, err := prompt.Run()

		if err != nil {
			return err
		}

		switch r {
		case apply:
			break ASK
		case cancel:
			return nil
		case all:
			selected = append(selected, unStages...)
			break ASK
		default:
			{
				selected = append(selected, unStages[i-defaultChoices])
				messages = stringSliceUtil.Remove(messages, i)
				unStages = gitStatusFileUtil.Remove(unStages, i-defaultChoices)
				selectedIndex = i
			}
		}
	}

	return (&git.GitUseCase{}).Stage(getPaths(selected))
}

func getValues(s []entity.GitStatusFile) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = v.Value
	}
	return r
}

func getPaths(s []entity.GitStatusFile) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = v.Path
	}
	return r
}
