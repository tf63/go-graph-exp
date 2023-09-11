# syntax=docker/dockerfile:1
FROM golang:1.19-alpine
# FROM golang:1.19-alpine as builder

# ワークディレクトリの指定
WORKDIR /app

# パッケージのインストール
RUN apk update --no-cache \
    && apk add --no-cache \
    git \
    gcc \
    musl-dev \
    make \
    wget

# モジュールのダウンロード
COPY backend/go.mod backend/go.sum backend/Makefile ./
RUN go mod download
RUN make install-tools

# ポートの開放
EXPOSE 8080


# ARG CGO_ENABLED=0
# ARG GOOS=linux
# ARG GOARCH=amd64
# RUN go build -ldflags '-s -w' ./server.go

# FROM alpine
# COPY --from=builder /build/server.go /app
# ENTRYPOINT ["/app/server"]