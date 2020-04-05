package git

import (
	"os/exec"
	"strings"

	"github.com/yasukotelin/git-ex/entity"
)

type StatusRepository interface {
	FetchStatus() ([]entity.GitStatusFile, error)
}

type StatusRepositoryImpl struct{}

func NewStatusRepositoryImpl() StatusRepository {
	return &StatusRepositoryImpl{}
}

func (r *StatusRepositoryImpl) FetchStatus() ([]entity.GitStatusFile, error) {
	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return nil, err
	}
	rows := strings.Split(string(out), "\n")
	// Remove the latest empty row.
	rows = rows[0 : len(rows)-1]

	result := make([]entity.GitStatusFile, len(rows))
	for i, row := range rows {
		result[i] = makeGitStatusFile(row)
	}
	return result, nil
}

func makeGitStatusFile(row string) entity.GitStatusFile {
	result := &entity.GitStatusFile{}

	statusMark := row[0:2]
	if statusMark == "??" {
		result.IsUnstaged = false
	} else {
		if statusMark[:1] != " " {
			result.IsStaged = true
		}
		if statusMark[1:2] != " " {
			result.IsUnstaged = true
		}
	}

	result.Value = row
	result.Path = row[3:]

	return *result
}
