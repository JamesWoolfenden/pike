# syntax=docker/dockerfile:1.7
#
# Multi-stage build:
#   1. Build pike from source so the image is reproducible from the checked-out SHA
#      (no fetching the "latest" release at build time).
#   2. Ship a minimal Alpine runtime with a non-root user and only the tools
#      entrypoint.sh actually needs.



# ---- builder ----------------------------------------------------------------
FROM golang:1.25.9-alpine3.22@sha256:ea77c38bc50df598f22ae02b729b9d37eb0d70ed72d6dd336b8d6c02ae2b8b09 AS builder

WORKDIR /src

# Cache modules separately from source so dep-only changes don't bust the layer.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Note: pike.Version is declared as a `const` in src/version.go, so it cannot
# be overridden at link time. Change it to a `var` if you want build-time
# version injection via `-ldflags -X`.
RUN CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags "-s -w" \
    -o /out/pike .

# ---- runtime ----------------------------------------------------------------
FROM alpine:3.22@sha256:310c62b5e7ca5b08167e4384c68db0fd2905dd9c7493756d356e893909057601

# entrypoint.sh needs bash; terraform/tofu flows often want git + ca-certificates.
RUN apk add --no-cache bash ca-certificates git \
    && adduser -D -u 10001 -h /home/pike pike

COPY --from=builder /out/pike /usr/bin/pike
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

USER pike
WORKDIR /home/pike

ENTRYPOINT ["/entrypoint.sh"]

LABEL org.opencontainers.image.title="pike" \
      org.opencontainers.image.description="Determine the minimum IAM permissions required to run OpenTofu/Terraform infrastructure code." \
      org.opencontainers.image.source="https://github.com/JamesWoolfenden/pike" \
      org.opencontainers.image.licenses="Apache-2.0" \
      org.opencontainers.image.authors="JamesWoolfenden"
