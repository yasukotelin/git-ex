package git

import (
	"github.com/yasukotelin/git-ex/entity"
	"github.com/yasukotelin/git-ex/repository/git"
)

type GitUseCase interface {
	FetchStatus() ([]entity.GitStatusFile, error)
	FetchUnStageStatusFiles() ([]entity.GitStatusFile, error)
	FetchStageStatusFiles() ([]entity.GitStatusFile, error)
	Stage(path string) error
	Stages(paths []string) error
	UnStage(path string) error
	UnStages(paths []string) error
	Discard() error
}

type GitUseCaseImpl struct {
	statusRepo git.StatusRepository
	stashRepo  git.StashRepository
	stageRepo  git.StageRepository
}

func NewGitUseCaseImpl(
	statusRepo git.StatusRepository,
	stashRepo git.StashRepository,
	stageRepo git.StageRepository,
) GitUseCase {
	return &GitUseCaseImpl{
		statusRepo: statusRepo,
		stashRepo:  stashRepo,
		stageRepo:  stageRepo,
	}
}

func (g *GitUseCaseImpl) FetchStatus() ([]entity.GitStatusFile, error) {
	return g.statusRepo.FetchStatus()
}

func (g *GitUseCaseImpl) FetchUnStageStatusFiles() ([]entity.GitStatusFile, error) {
	status, err := g.statusRepo.FetchStatus()
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

func (g *GitUseCaseImpl) FetchStageStatusFiles() ([]entity.GitStatusFile, error) {
	status, err := g.statusRepo.FetchStatus()
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

func (g *GitUseCaseImpl) Stage(path string) error {
	return g.stageRepo.Stage(path)
}

func (g *GitUseCaseImpl) Stages(paths []string) error {
	return g.stageRepo.Stages(paths)
}

func (g *GitUseCaseImpl) UnStage(path string) error {
	return g.stageRepo.UnStage(path)
}

func (g *GitUseCaseImpl) UnStages(paths []string) error {
	return g.stageRepo.UnStages(paths)
}

func (g *GitUseCaseImpl) Discard() error {
	err := g.stashRepo.SaveIncludeUntracked()
	if err != nil {
		return err
	}
	return g.stashRepo.Drop()
}
