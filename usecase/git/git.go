package git

import (
	"github.com/yasukotelin/git-ex/repository/git"
)

type GitUseCase struct {
}

func (g *GitUseCase) Discard() error {
	stashRepo := &git.StashRepository{}

	err := stashRepo.SaveIncludeUntracked()
	if err != nil {
		return err
	}
	return stashRepo.Drop()
}
