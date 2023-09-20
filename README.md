# GoでGraphQLのサーバーを立ち上げるまで

### 環境設定
backend
- Goのコンテナ
- devcontainerでVSCodeにアタッチするように設定
- 8080:8080でフォワーディング
- `docker/backend/Dockerfile`

postgres
- DBコンテナ
- 5432:5432

pgadmin4
- pgadminのコンテナ
- 8888:80

環境変数は`.env.dev`参照

### コンテナの立ち上げ
```bash
    docker compose up -d
```
