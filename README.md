CRUD 操作の基本を確認するための TODO アプリケーション

## MySQL 起動

db ディレクトリに移動し Docker 起動

```
cd db
docker compose up -d
```

## .env の読み込み

```
source .env
```

## GoWeb サーバの開始

ルートディレクトリで run する

```
go run main.go
```

## リクエスト例

```
curl localhost:3000/todo
```

## レスポンス例

```
{
    "result":[
        {
            "ID":1,
            "Title":"wash",
            "Body":"toilet"
        }
    ]
}
```
