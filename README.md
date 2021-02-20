# CloudWatch Synthetics 定時起動用Lambda

## 背景

- 2021年2月20日現在、CloudWatch Syntheticsは時刻指定での実行ができない
- Syntheticsを実行するだけのLambdaを作成、それを時刻指定で実行するもの

## 開発環境構築

- Macによる開発を前提とする。
- Docker for Macをインストールする。
- リポジトリをクローンする

```bash
git clone git@github.com:mmmsasaki/start_synthetics_lambda.git
cd start_synthetics_lambda
```

- `env.sample` ファイルをコピーして、 `.env` を作成し、そこにAWSアクセスキー、Canary名など入力する

```bash
cp env.sample .env
```

- デプロイ方法

```
make deploy
```
