package git

import (
	"os/exec"
	"strings"
)

type BranchRepository interface {
	Fetch() ([]string, error)
	FetchMerged() ([]string, error)
	Checkout(path string) error
}

type BranchRepositoryImpl struct{}

func NewBranchRepositoryImpl() BranchRepository {
	return &BranchRepositoryImpl{}
}

func (b *BranchRepositoryImpl) Fetch() ([]string, error) {
	return fetch()
}

func (b *BranchRepositoryImpl) FetchMerged() ([]string, error) {
	return fetch("--merged")
}

func fetch(option ...string) ([]string, error) {
	arg := append([]string{"branch"}, option...)
	out, err := exec.Command("git", arg...).Output()
	if err != nil {
		return nil, err
	}
	rows := strings.Split(string(out), "\n")
	// Remove the latest empty row.
	rows = rows[0 : len(rows)-1]

	result := make([]string, 0, len(rows))
	for _, row := range rows {
		if row[0:1] == "*" {
			// 先頭1文字目が*の時はカレントブランチなので対象から外す
			continue
		}
		branch := strings.TrimLeft(row, " ")
		result = append(result, branch)
	}

	return result, nil
}

func (b *BranchRepositoryImpl) Checkout(path string) error {
	return exec.Command("git", "checkout", path).Run()
}
