# データベース設計

## MySQL

``` mermaid
---
title: ユーザー情報
---
erDiagram
    users ||--|| passwords: ""

    users {
        string id PK "ID"
        string username UK "ユーザー名"
        string email "メールアドレス"
        boolean is_deleted "論理削除のフラグ"
        timestamp created_at "作成日時"
        timestamp updated_at "更新日時"
    }

    passwords {
        string id PK "ID"
        string user_id FK "ユーザーID"
        string hashed_password "パスワード"
        timestamp created_at "作成日時"
        timestamp updated_at "更新日時"
    }
```

## Redis

``` mermaid
---
title: キャッシュ
---
erDiagram
    tokens {
        string user_id
        string token
    }
```