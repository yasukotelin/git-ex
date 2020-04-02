package util

import "github.com/yasukotelin/git-ex/entity"

type StringSliceUtil struct{}

func (u *StringSliceUtil) Remove(s []string, i int) []string {
	s = append(s[:i], s[i+1:]...)
	n := make([]string, len(s))
	copy(n, s)
	return n
}

type GitStatusFileSliceUtil struct{}

func (u *GitStatusFileSliceUtil) Remove(g []entity.GitStatusFile, i int) []entity.GitStatusFile {
	g = append(g[:i], g[i+1:]...)
	n := make([]entity.GitStatusFile, len(g))
	copy(n, g)
	return n
}
