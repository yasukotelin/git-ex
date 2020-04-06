# git-ex

[Japanese page](./README-JP.md)

git-ex is a subcommand that extends Git.

## Installation

```
go get -u github.com/yasukotelin/git-ex
```

## Commands

- `git ex branch`
- `git ex stage`
- `git ex unstage`
- `git ex diff`
- `git ex discard`

## Features

### branch

```
git ex branch
```

You can checkout a branch with selecter.

You can use the `--all` or `-a` option to select from local and remote file.

```
git ex branch --all
```

### Stage

```
git ex stage
```

You can stage the files with selecter.

![stage.gif](./images/stage.gif)

### UnStage

```
git ex unstage
```

You can unstage the files with selecter.

![unstage.gif](./images/unstage.gif)

### Diff

```
git ex diff
```

You can see the diff selected file.By default, you can choose from the unstage file.

You can use the `--stage` option to select from a stage file.

```
git ex diff --stage
```

### Discard

```
git ex discard
```

This executes the removing all changes from the HEAD that include untracked files.

`git checkout .` doesn't remove untracked files.

## Author

yasukotelin

## LICENCE

MIT LICENCE