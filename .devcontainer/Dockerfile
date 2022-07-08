ARG GOVERSION=${GOVERSION:-1.18}
FROM mcr.microsoft.com/vscode/devcontainers/go:${GOVERSION}-bullseye

RUN echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | tee /etc/apt/sources.list.d/goreleaser.list \
    && apt update \
    && apt install  --no-install-recommends -y goreleaser \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*

