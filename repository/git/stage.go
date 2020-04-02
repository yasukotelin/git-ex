package git

import (
	"os/exec"
)

type StageRepository struct{}

func (s StageRepository) Stage(path string) error {
	return exec.Command("git", "add", path).Run()
}

func (s StageRepository) Stages(paths []string) error {
	opt := append([]string{"add"}, paths...)
	return exec.Command("git", opt...).Run()
}
