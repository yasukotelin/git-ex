package cmd

import (
	repoGit "github.com/yasukotelin/git-ex/repository/git"
	"github.com/yasukotelin/git-ex/usecase/git"
	"github.com/yasukotelin/git-ex/util"
)

var gitUseCase = git.NewGitUseCaseImpl(
	repoGit.NewStatusRepositoryImpl(),
	repoGit.NewStashRepositoryImpl(),
	repoGit.NewStageRepositoryImpl(),
)
var stringSliceUtil = &util.StringSliceUtil{}
var gitStatusFileUtil = &util.GitStatusFileSliceUtil{}
