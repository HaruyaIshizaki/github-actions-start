# GitHub Actions の基本

## ワークフローファイルの構造

以下は基本的な GitHub Actions のワークフローファイル（`.github/workflows/hello.yml`）の例です：

```yaml
name: learn-github-actions # ワークフローの名前
run-name: ${{ github.actor }} is learning GitHub Actions # 実行時の名前
on: [push] # トリガー
jobs:
  hello-world: # ジョブの名前
    runs-on: ubuntu-latest # ジョブの実行環境（ランナー）
    steps: # ジョブのステップ
      - name: Checkout # ステップの名前
        uses: actions/checkout@v4 # ステップの実行（アクション）
      - name: Hello world # ステップの名前
        run: echo "Hello, world!" # ステップの実行（シェルコマンド）
```

## 主要な要素の説明

### 1. ワークフローの基本設定

- `name`: ワークフローの名前を指定します
- `run-name`: 実行時の表示名を指定します（`${{ github.actor }}`は実行者のユーザー名を表示）
- `on`: ワークフローを実行するトリガーを指定します（この例では`push`イベントで実行）

### 2. ジョブの設定

- `jobs`: ワークフロー内の実行単位を定義します
- `hello-world`: ジョブの名前
- `runs-on`: ジョブを実行する環境を指定します（この例では`ubuntu-latest`）

### 3. ステップの実行

- `steps`: ジョブ内で実行する個々のタスクを定義します
- `name`: ステップの名前
- `uses`: 再利用可能なアクションを指定します
- `run`: シェルコマンドを実行します

## 重要なアクション: actions/checkout@v4

`actions/checkout@v4`は、GitHub Actions の最も基本的で重要なアクションの 1 つです：

- リポジトリのコードをワークフローの実行環境（ランナー）にチェックアウト（ダウンロード）します
- これにより、ワークフロー内でリポジトリのコードにアクセスできるようになります
- `@v4`は、このアクションのバージョンを指定しています（現在の最新の安定バージョン）
- リポジトリのコードを操作する必要がある場合は、ほぼ必ず使用することになります

このアクションがなければ、ワークフローは空の環境で実行されることになり、リポジトリのコードにアクセスすることができません。

## ランナー

ランナーは、GitHub Actions の実行環境です。
Hosted Runner は GitHub が提供するランナーで、GitHub Actions の実行環境を提供します。

GitHub Actions では、以下のような Hosted Runner を使用できます：

- `ubuntu-latest`: Ubuntu の最新バージョン
- `ubuntu-20.04`: Ubuntu の 20.04 バージョン
- `ubuntu-18.04`: Ubuntu の 18.04 バージョン
- `windows-latest`: Windows の最新バージョン
- `macos-latest`: macOS の最新バージョン

また、エフェメラルという特性があり、ランナーはジョブ開始時に起動し、終了時に破棄されます。
これにより、毎回同じ環境でワークフローを実行できます。

## トリガーの種類

GitHub Actions では、以下のようなトリガーを使用できます：

- `push`: リポジトリにコードがプッシュされたとき
- `pull_request`: リポジトリにプルリクエストが作成されたとき
- `schedule`: 定期的に実行されるように設定されたスケジュール
- `workflow_dispatch`: 手動でワークフローを実行するとき
- `workflow_run`: 他のワークフローが実行されたとき
