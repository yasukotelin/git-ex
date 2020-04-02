package git

import (
	"os/exec"
)

type StageRepository struct{}

func (s StageRepository) Stage(paths []string) error {
	opt := append([]string{"add"}, paths...)
	return exec.Command("git", opt...).Run()
}
