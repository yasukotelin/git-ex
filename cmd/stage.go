package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"github.com/yasukotelin/git-ex/entity"
	"github.com/yasukotelin/git-ex/usecase/git"
	"github.com/yasukotelin/git-ex/util"
)

const all = "...(All)"
const finish = "...(Finish)"
const defaultChoices = 2

var gitUseCase = &git.GitUseCase{}
var stringSliceUtil = &util.StringSliceUtil{}
var gitStatusFileUtil = &util.GitStatusFileSliceUtil{}

// Stage stages the files with selecter
func Stage(c *cli.Context) error {
	unStages, err := gitUseCase.FetchUnStageStatusFiles()
	if err != nil {
		return err
	}

	messages := append([]string{all, finish}, getValues(unStages)...)
	selectedIndex := defaultChoices

	for {
		if len(messages) == defaultChoices {
			// 選択肢が無くなったら終了
			return nil
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
		case all:
			gitUseCase.Stages(getPaths(unStages))
			return nil
		case finish:
			return nil
		default:
			{
				gitUseCase.Stage(unStages[i-defaultChoices].Path)
				messages = stringSliceUtil.Remove(messages, i)
				unStages = gitStatusFileUtil.Remove(unStages, i-defaultChoices)
				selectedIndex = i
			}
		}
	}
}

// UnStage unstages the files with selecter
func UnStage(c *cli.Context) error {
	stages, err := gitUseCase.FetchStageStatusFiles()
	if err != nil {
		return err
	}

	messages := append([]string{all, finish}, getValues(stages)...)
	selectedIndex := defaultChoices

	for {
		if len(messages) == defaultChoices {
			// 選択肢が無くなったら終了
			return nil
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
		case all:
			gitUseCase.UnStages(getPaths(stages))
			return nil
		case finish:
			return nil
		default:
			{
				gitUseCase.UnStage(stages[i-defaultChoices].Path)
				messages = stringSliceUtil.Remove(messages, i)
				stages = gitStatusFileUtil.Remove(stages, i-defaultChoices)
				selectedIndex = i
			}
		}
	}
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
