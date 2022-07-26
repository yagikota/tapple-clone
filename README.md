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
