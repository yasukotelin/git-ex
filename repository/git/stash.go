package git

import (
	"os/exec"
)

type StashRepository struct{}

func (g *StashRepository) SaveIncludeUntracked() error {
	return exec.Command("git", "stash", "save", "--include-untracked").Run()
}

// Drop removes a latest stash.
func (g *StashRepository) Drop() error {
	return exec.Command("git", "stash", "drop").Run()
}
