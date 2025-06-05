# 自動テストのワークフロー設計

# CI に自動テストを組み込む手順（Go + GitHub Actions）

このドキュメントでは、Goプロジェクトにおいて GitHub Actions を用いた自動テストの組み込み方法を記載します。初期セットアップから、トリガーイベント別の運用目的、Goのパッケージに関する補足までをカバーします。

---

## 1. 初期セットアップ手順

### Step 1: `.github/workflows` ディレクトリを作成

```bash
mkdir -p .github/workflows
```

### Step 2: 自動テストワークフロー（例: `test.yml`）を作成

```yaml
# .github/workflows/test.yml
name: Go Test

on:
  pull_request:
    paths:
      - 'src/**/*.go'
  push:
    branches:
      - main
      - 'feature/**'

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.20'
      - run: go test ./src
```

---

## 2. 自動テストを走らせるイベントの分類と目的（具体的な記法付き）

| イベント | 使用例 | YAML記法 | 採択目的 |
|----------|--------|----------|----------|
| Pull Request 作成・更新時 | 他ブランチから main に PR | `pull_request:` | コードレビュー前にテストを通して品質担保 |
| ブランチへ push 時（main） | mainブランチへの直pushやマージ | `push:\n  branches:\n    - main` | CIの強制実行（本番ブランチでの安全確認） |
| featureブランチ push時 | `feature/xxx` のブランチでの開発 | `push:\n  branches:\n    - 'feature/**'` | 機能開発中でも早期にテスト失敗を検出 |
| 手動実行 | 任意のタイミングで確認 | `workflow_dispatch:` | 手動でワークフロー確認・検証したいとき |
---

## 3. Go パッケージと `go test` の注意点

- `go test ./src` のように**ディレクトリ単位**で実行することで、同じパッケージ内の `.go` / `_test.go` ファイル間の依存関係が正しく解決される
- `go test src/*.go` のように個別ファイル指定をすると、パッケージ単位でのビルドがされず、**関数未定義エラー**になる恐れがある

---

## 4. 開発フローの推奨順序

1. `.github/workflows` を `main` に追加（CIの土台を整備）
2. `main.go` ＋ `main_test.go` を含むブランチを作成してPR
3. PR時にCIが実行され、テスト結果でレビューの判断が可能
4. 今後はどのPRでも自動的にCIが走る体制が整う

---

## 5. 補足

- `go mod init <module名>` でGo Modulesを使い始めると依存管理がしやすい
- GitHub Actions の `workflow_dispatch:` を使うと、手動でもワークフローを実行できる（開発・検証に便利）

---