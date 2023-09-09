# GoでGraphQLを導入する

- GoのGraphQLライブラリとして`gqlgen`があります
- 今回は`gqlgen`を使ってGraphQLのサーバーを立ち上げるまでのチュートリアルを作成していきます

### gqlgenのインストール
```
    go get github.com/99designs/gqlgen@latest
    go install github.com/99designs/gqlgen@latest
```

### GraphQLプロジェクトの作成

```
    gqlgen init
```

デフォルトではこのような構成でファイルが生成されます
```
    .
    ├── Makefile
    ├── docs
    │   └── graph.md
    ├── go.mod
    ├── go.sum
    ├── gqlgen.yml
    ├── graph
    │   ├── generated.go
    │   ├── model
    │   │   └── models_gen.go
    │   ├── resolver.go
    │   ├── schema.graphqls
    │   └── schema.resolvers.go
    └── server.go
```