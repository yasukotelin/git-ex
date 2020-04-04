package git

import (
	"os/exec"
)

type StageRepository interface {
	Stage(path string) error
	Stages(paths []string) error
	UnStage(path string) error
	UnStages(paths []string) error
}

type StageRepositoryImpl struct{}

func NewStageRepositoryImpl() StageRepository {
	return &StageRepositoryImpl{}
}

func (s StageRepositoryImpl) Stage(path string) error {
	return exec.Command("git", "add", path).Run()
}

func (s StageRepositoryImpl) Stages(paths []string) error {
	opt := append([]string{"add"}, paths...)
	return exec.Command("git", opt...).Run()
}

func (s StageRepositoryImpl) UnStage(path string) error {
	return exec.Command("git", "reset", "HEAD", path).Run()
}

func (s StageRepositoryImpl) UnStages(paths []string) error {
	opt := append([]string{"reset", "HEAD"}, paths...)
	return exec.Command("git", opt...).Run()
}
