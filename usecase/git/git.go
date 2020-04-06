package git

import (
	"github.com/yasukotelin/git-ex/entity"
	"github.com/yasukotelin/git-ex/repository/git"
)

type GitUseCase interface {
	FetchStatus() ([]entity.GitStatusFile, error)
	FetchStage() ([]entity.GitStatusFile, error)
	FetchUnStage() ([]entity.GitStatusFile, error)
	FetchUnStageWithUntracked() ([]entity.GitStatusFile, error)
	FetchBranch() ([]string, error)
	FetchMergedBranch() ([]string, error)
	Checkout(path string) error
	Stage(path string) error
	Stages(paths []string) error
	UnStage(path string) error
	UnStages(paths []string) error
	DiffAll(isStage bool) error
	Diff(isStage bool, path string) error
	DeleteBranch(branch string) error
	Discard() error
}

type GitUseCaseImpl struct {
	statusRepo git.StatusRepository
	stashRepo  git.StashRepository
	stageRepo  git.StageRepository
	diffRepo   git.DiffRepository
	branchRepo git.BranchRepository
}

func NewGitUseCaseImpl(
	statusRepo git.StatusRepository,
	stashRepo git.StashRepository,
	stageRepo git.StageRepository,
	diffRepo git.DiffRepository,
	branchRepo git.BranchRepository,
) GitUseCase {
	return &GitUseCaseImpl{
		statusRepo: statusRepo,
		stashRepo:  stashRepo,
		stageRepo:  stageRepo,
		diffRepo:   diffRepo,
		branchRepo: branchRepo,
	}
}

func (g *GitUseCaseImpl) FetchStatus() ([]entity.GitStatusFile, error) {
	return g.statusRepo.FetchStatus()
}

func (g *GitUseCaseImpl) FetchStage() ([]entity.GitStatusFile, error) {
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

func (g *GitUseCaseImpl) FetchUnStage() ([]entity.GitStatusFile, error) {
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

func (g *GitUseCaseImpl) FetchUnStageWithUntracked() ([]entity.GitStatusFile, error) {
	status, err := g.statusRepo.FetchStatus()
	if err != nil {
		return nil, err
	}
	result := make([]entity.GitStatusFile, 0, len(status))
	for _, s := range status {
		if s.IsUnstaged || s.IsUntracked {
			result = append(result, s)
		}
	}
	return result, nil
}

func (g *GitUseCaseImpl) FetchBranch() ([]string, error) {
	return g.branchRepo.Fetch()
}

func (g *GitUseCaseImpl) FetchMergedBranch() ([]string, error) {
	return g.branchRepo.FetchMerged()
}

func (g *GitUseCaseImpl) Checkout(path string) error {
	return g.branchRepo.Checkout(path)
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

func (g *GitUseCaseImpl) DiffAll(isStage bool) error {
	return g.diffRepo.DiffAll(isStage)
}

func (g *GitUseCaseImpl) Diff(isStage bool, path string) error {
	return g.diffRepo.Diff(isStage, path)
}

func (g *GitUseCaseImpl) DeleteBranch(branch string) error {
	return nil
}

func (g *GitUseCaseImpl) Discard() error {
	err := g.stashRepo.SaveIncludeUntracked()
	if err != nil {
		return err
	}
	return g.stashRepo.Drop()
}
