package git

import (
	"github.com/yasukotelin/git-ex/entity"
	"github.com/yasukotelin/git-ex/repository/git"
)

type GitUseCase struct {
}

func (g *GitUseCase) FetchStatus() ([]entity.GitStatusFile, error) {
	return (&git.StatusRepository{}).FetchStatus()
}

func (g *GitUseCase) FetchUnStageStatusFiles() ([]entity.GitStatusFile, error) {
	status, err := (&git.StatusRepository{}).FetchStatus()
	if err != nil {
		return nil, err
	}
	result := make([]entity.GitStatusFile, 0, len(status))
	for _, s := range status {
		if s.IsUnstaged {
			result = append(result, s)
		}
	}
	return result, nil
}

func (g *GitUseCase) FetchStageStatusFiles() ([]entity.GitStatusFile, error) {
	status, err := (&git.StatusRepository{}).FetchStatus()
	if err != nil {
		return nil, err
	}
	result := make([]entity.GitStatusFile, 0, len(status))
	for _, s := range status {
		if s.IsStaged {
			result = append(result, s)
		}
	}
	return result, nil
}

func (g *GitUseCase) Stage(path string) error {
	return (&git.StageRepository{}).Stage(path)
}

func (g *GitUseCase) Stages(paths []string) error {
	return (&git.StageRepository{}).Stages(paths)
}

func (g *GitUseCase) UnStage(path string) error {
	return (&git.StageRepository{}).UnStage(path)
}

func (g *GitUseCase) UnStages(paths []string) error {
	return (&git.StageRepository{}).UnStages(paths)
}

func (g *GitUseCase) Discard() error {
	stashRepo := &git.StashRepository{}

	err := stashRepo.SaveIncludeUntracked()
	if err != nil {
		return err
	}
	return stashRepo.Drop()
}
