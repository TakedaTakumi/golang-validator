# Go開発環境
FROM golang:1.24

ARG TZ
ENV TZ="$TZ"

# 開発コンテナであることを示す環境変数を設定
ENV DEVCONTAINER=true

WORKDIR /app

# 必要なツールのインストール
RUN apt-get update && apt-get install -y \
    curl \
    nodejs \
    npm \
    && rm -rf /var/lib/apt/lists/*

RUN npm install -g @anthropic-ai/claude-code

# COPY go.mod go.sum ./
# RUN go mod download
