# syntax=docker/dockerfile:1.7
#
# Multi-stage build:
#   1. Build pike from source so the image is reproducible from the checked-out SHA
#      (no fetching the "latest" release at build time).
#   2. Ship a minimal Alpine runtime with a non-root user and only the tools
#      entrypoint.sh actually needs.

ARG GO_VERSION=1.25.9
ARG ALPINE_VERSION=3.20

# ---- builder ----------------------------------------------------------------
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

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
FROM alpine:${ALPINE_VERSION}

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
