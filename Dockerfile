# Go開発環境
FROM golang:1.24

ARG TZ
ENV TZ="$TZ"

# 開発用ユーザーを作成
ARG USERNAME=godev
RUN groupadd --gid 1000 $USERNAME && \
    useradd --uid 1000 --gid $USERNAME --shell /bin/bash --create-home $USERNAME

# Node.js関連ディレクトリを作成（Claude Code用）
RUN mkdir -p /usr/local/share/npm-global && \
    chown -R $USERNAME:$USERNAME /usr/local/share

# bashの履歴を永続化
RUN SNIPPET="export PROMPT_COMMAND='history -a' && export HISTFILE=/commandhistory/.bash_history" \
  && mkdir /commandhistory \
  && touch /commandhistory/.bash_history \
  && chown -R $USERNAME:$USERNAME /commandhistory

# 開発コンテナであることを示す環境変数を設定
ENV DEVCONTAINER=true

# ワークスペースと設定ディレクトリを作成し、権限を設定
RUN mkdir -p /app /home/$USERNAME/.claude && \
  chown -R $USERNAME:$USERNAME /app /home/$USERNAME/.claude

WORKDIR /app

# 必要なツールのインストール
RUN apt-get update && apt-get install -y \
    curl \
    git \
    vim \
    sudo \
    zsh \
    && rm -rf /var/lib/apt/lists/*

# Node.js環境とClaude Codeのインストール（rootユーザーで実行）
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.3/install.sh | bash \
    && . "$HOME/.nvm/nvm.sh" \
    && nvm install 22 \
    && corepack enable pnpm \
    && export SHELL=/bin/bash \
    && pnpm setup \
    && export PNPM_HOME="$HOME/.local/share/pnpm" \
    && export PATH="$PNPM_HOME:$PATH" \
    && pnpm install -g @anthropic-ai/claude-code

# godevユーザーに切り替え
USER $USERNAME

# pnpm環境変数の設定（godevユーザー用）
ENV PNPM_HOME="/root/.local/share/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
ENV SHELL="/bin/zsh"

# COPY go.mod go.sum ./
# RUN go mod download
