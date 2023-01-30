CRUD操作の基本を確認するためのTODOアプリケーション

## MySQL起動
```
docker compose up -d
```

## GoWebサーバの開始
```
go run main.go
```

## リクエスト
```
curl localhost:3000/todo
```

##　レスポンス例
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