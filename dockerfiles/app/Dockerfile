FROM unit:1.34.2-go1.24 AS base

FROM base AS devserver

# volume
RUN chmod 777 /mnt

# add normal user
ARG USER_UID=1000
ARG USER_GID=1000
RUN groupadd --gid $USER_GID unit2 \
  && useradd --uid $USER_UID --gid $USER_GID -M -s /usr/sbin/nologin unit2

# 開発コンテナから設定変更できるようにする
CMD ["unitd","--no-daemon","--control","0.0.0.0:7999","--user","unit2","--group","unit2"]

FROM base AS workspace

# 定義済みARG変数を使用する
# https://docs.docker.com/engine/reference/builder/#automatic-platform-args-in-the-global-scope
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

CMD ["bash"]

# prerequisite/utility packages
RUN apt-get update \
  && DEBIAN_FRONTEND=noninteractive apt-get install -y \
  acl \
  make \
  curl \
  git \
  less \
  sudo \
  jq \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

# adjust GOPATH permission
RUN chgrp -R users /go && chmod -R g+ws /go && setfacl -R -d -m g::rwx /go

# postgresql-client
RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(cat /etc/os-release | grep VERSION_CODENAME | head -n 1 | cut -d = -f 2)-pgdg main" > /etc/apt/sources.list.d/pgdg.list' \
  && curl -sSL https://www.postgresql.org/media/keys/ACCC4CF8.asc | tee /etc/apt/trusted.gpg.d/ACCC4CF8.asc > /dev/null \
  && apt-get update \
  && apt-get install -y locales postgresql-client-16 \
  && rm -rf /var/lib/apt/lists/*
RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8

# go tools for vscode
#   https://pkg.go.dev/github.com/go-delve/delve
ARG GOTOOLS_DLV_VERSION=1.24
RUN --mount=type=cache,target=/go/pkg/mod \
  go install github.com/go-delve/delve/cmd/dlv@v${GOTOOLS_DLV_VERSION} \
  && go clean -cache

# add normal user
ARG USER_NAME=dev
ARG USER_UID=1000
ARG USER_GID=1000
RUN groupadd --gid $USER_GID $USER_NAME \
  && useradd --uid $USER_UID --gid $USER_GID -m $USER_NAME \
  && usermod -aG users $USER_NAME

# passwordless sudo
RUN echo $USER_NAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USER_NAME \
  && chmod 0440 /etc/sudoers.d/$USER_NAME

# history
# ref: https://code.visualstudio.com/remote/advancedcontainers/persist-bash-history
RUN SNIPPET="export PROMPT_COMMAND='history -a' && export HISTFILE=/usr/local/history/.bash_history" \
  && mkdir /usr/local/history \
  && chown -R $USER_NAME:$USER_NAME /usr/local/history \
  && echo "$SNIPPET" >> "/home/$USER_NAME/.bashrc"

# locales
ENV LANG=C.UTF-8

# timezone
ENV TZ Japan

# to get right permission
RUN mkdir -p /go/pkg/mod

# switch to normal user
USER $USER_NAME

# to get right permission
RUN mkdir -p /home/$USER_NAME/.vscode-server/extensions
RUN mkdir -p /home/$USER_NAME/.cache/go-build

# .env に書いた環境変数をシェルに取り込む処理を追加
# 使用する前に開発環境で `go install ./cmd/exportenv` としてコマンドをインストールする必要がある
RUN printf 'function expenv() {\n    if [ ! -f /go/bin/exportenv ]; then echo run go install ./cmd/exportenv; fi\n    eval $(exportenv -e)\n}' >> "/home/$USER_NAME/.bashrc"
