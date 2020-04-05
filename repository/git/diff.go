package git

import (
	"os"
	"os/exec"
)

type DiffRepository interface {
	DiffAll(isStage bool) error
	Diff(isStage bool, path string) error
}

type DiffRepositoryImpl struct{}

func NewDiffRepositoryImpl() DiffRepository {
	return &DiffRepositoryImpl{}
}

func (d *DiffRepositoryImpl) DiffAll(isStage bool) error {
	var cmd *exec.Cmd
	if isStage {
		cmd = exec.Command("git", "diff", "--cached")
	} else {
		cmd = exec.Command("git", "diff")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (d *DiffRepositoryImpl) Diff(isStage bool, path string) error {
	var cmd *exec.Cmd
	if isStage {
		cmd = exec.Command("git", "diff", "--cached", path)
	} else {
		cmd = exec.Command("git", "diff", path)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
