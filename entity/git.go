package entity

type GitStatusFile struct {
	Value       string
	Path        string
	IsStaged    bool
	IsUnstaged  bool
	IsUntracked bool
}
