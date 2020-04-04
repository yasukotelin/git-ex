package git

import (
	"os/exec"
)

type StashRepository interface {
	SaveIncludeUntracked() error
	Drop() error
}

type StashRepositoryImpl struct{}

func NewStashRepositoryImpl() StashRepository {
	return &StashRepositoryImpl{}
}

func (g *StashRepositoryImpl) SaveIncludeUntracked() error {
	return exec.Command("git", "stash", "save", "--include-untracked").Run()
}

// Drop removes a latest stash.
func (g *StashRepositoryImpl) Drop() error {
	return exec.Command("git", "stash", "drop").Run()
}
