# 2208-ace-go-server

## Installation
Dockerをインストールしておく <br>
https://docs.docker.com/get-docker/ <br><br>

.env.exampleをコピーして.envを作成
```shell script
$ cp .env.example .env
```

## Usage

DB起動
```shell script
$ make run-db
```

API Server起動(mysqlの起動に少し時間がかかるため↑実行後に少し待ってから実行)
```shell script
$ make run-go
```

シャットダウン
```shell script
$ make down
```

# ブランチ名規約(適宜追加してください)

| ブランチ名| 説明|
| - | - |
|feat/hoge| 新しい機能|
|fix/hoge| バグの修正|
|refactor/hoge| 仕様に影響がないコード改善(リファクタ)|
|chore/hoge| ビルド、補助ツール、ライブラリ関連|
|docs/hoge| ドキュメントのみの変更|
|style/hoge| 空白、フォーマット、セミコロン追加など|
|perf/hoge| パフォーマンス向上関連|
|test/hoge| テスト関連|

# コミットメッセージ(適宜追加してください)
* 日本語
* プレフィックスつける
    * 例
        | フレフィックス| 説明|
        | - | - |
        |feat:| 新しい機能|
        |fix:| バグの修正|
        |refactor:| 仕様に影響がないコード改善(リファクタ)|
        |chore:| ビルド、補助ツール、ライブラリ関連|
        |docs:| ドキュメントのみの変更|
        |style:| 空白、フォーマット、セミコロン追加など|
        |perf:| パフォーマンス向上関連|
        |test:| テスト関連|

* [Qiita記事](https://qiita.com/konatsu_p/items/dfe199ebe3a7d2010b3e)
* [Conventional Commits](https://www.conventionalcommits.org/ja/v1.0.0/#%e4%bb%95%e6%a7%98)
* [例文集](https://gist.github.com/mono0926/e6ffd032c384ee4c1cef5a2aa4f778d7#%E8%A1%A8%E7%8F%BE%E5%82%BE%E5%90%91%E3%81%A8%E3%81%BE%E3%81%A8%E3%82%81)
