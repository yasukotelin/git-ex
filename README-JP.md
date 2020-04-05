# git-ex

git-exはGitコマンドを拡張するサブコマンドです。
サブコマンドなのでインストールすると `git ex` の形式で使えます。

## Installation

```
go get -u github.com/yasukotelin/git-ex
```

## Commands

- `git ex stage`
- `git ex unstage`
- `git ex diff`
- `git ex discard`

## Features

### Stage

```
git ex stage
```

ステージングしたいファイルを選択してステージングすることができます。

![stage.gif](./images/stage.gif)

### UnStage

```
git ex unstage
```

ワーキングエリアに戻したいファイルを選択してアンステージングすることができます。

![unstage.gif](./images/unstage.gif)

### Diff

```
git ex diff
```

You can see the diff selected file.By default, you can choose from the unstage file.
ファイルを選択してdiffを表示することができます。デフォルトでは、アンステージングファイルの中から選択可能です。

`--satage` オプションを付ければステージングファイルの中から選択できます。

```
git ex diff --stage
```

### Discard

```
git ex discard
```

何かしらの修正を一度破棄してHEADまで戻したい場合、通常 `git checkout .` を使うと思いますが、未追跡なファイル（Untracked files）は破棄されません。

このDiscardを実行すると、未追跡ファイルも含めて綺麗にHEADの状態まで変更を破棄することができます。

## Author

yasukotelin

## LICENCE

MIT LICENCE