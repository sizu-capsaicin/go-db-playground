# go-db-playground

Go 言語を用いて様々なデータベースに操作を加えてみる．

- mongodb
  - MongoDB 向けのソースコードや json ファイル
- mysql
  - MySQL 向けのソースコード

## ディレクトリ構成

- mongo
  - json
    - 各車両形式をモデルにした json 形式のデータを格納
  - queries
    - drop
      - 既存のデータを全て削除する
    - find
      - find-all
        - 全てのデータを取得して表示する
      - find-sample
        - 条件を付けてデータを取得して表示するサンプル
    - insert
      - json 形式のファイルを読み込んで mongoDB へデータを挿入する
  - db.go
    - mongoDB への操作をするための関数をまとめたファイル
  - struct.go
    - 使用する構造体をまとめたファイル
- mysql
  - queries
    - create
      - データベースやテーブルを作成する
    - insert
      - テーブルにデータを挿入する
    - select
      - データを取得して表示するサンプル
  - struct.go
    - 使用する構造体をまとめたファイル
