# git-ex

git-exはGitコマンドを拡張するサブコマンドです。
サブコマンドなのでインストールすると `git ex` の形式で使えます。

## Installation

```
go get install github.com/yasukotelin/git-ex
```

## Commands

- `git ex discard`

## Features

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