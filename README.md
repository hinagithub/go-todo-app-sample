CRUD 操作の基本を確認するための TODO アプリケーション

# 　完成画面イメージ
### 一覧ページ

<img width="400" alt="image" src="https://user-images.githubusercontent.com/44778704/216805058-d4ca1cc5-508a-4955-b1f5-78b4dc163941.png">

### 新規作成フォームモーダル

<img width="200" alt="image" src="https://user-images.githubusercontent.com/44778704/216805076-fd0dceee-0c28-4e45-bbee-e5f09896db1d.png">

### 編集フォームモーダル

<img width="200" alt="image" src="https://user-images.githubusercontent.com/44778704/216805107-466b12c1-4659-47e1-9172-61cbaeb3b0f1.png">

### 仕様について
- Addボタンをクリックすると新規作成モーダルが開きます
- TODOリストの行をクリックすると該当アイテムの編集モーダルが開きます
- Completeボタンを押すとその行のcompleteボタンが消えて、文字がグレーアウトされます
- バツボタンをクリックすると、その行のアイテムが表示されなくなります

# 起動手順
### MySQL 起動

db ディレクトリに移動し Docker 起動

```
cd db
docker compose up -d
```

### .env の読み込み

```
source .env
```

### サーバの開始

ルートディレクトリで run する

```
go run main.go
```

### ブラウザで以下にアクセス

```
localhost:3000/todo
```

