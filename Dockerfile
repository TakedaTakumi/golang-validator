# Go開発環境
FROM golang:1.24

ARG TZ
ENV TZ="$TZ"

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# 開発コンテナであることを示す環境変数を設定
ENV DEVCONTAINER=true

WORKDIR /app

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME

# 必要なツールのインストール
RUN apt-get update && apt-get install -y \
    curl \
    nodejs \
    npm \
    && rm -rf /var/lib/apt/lists/*

RUN npm install -g @anthropic-ai/claude-code

USER $USERNAME

RUN mkdir /home/vscode/.claude

# COPY go.mod go.sum ./
# RUN go mod download
