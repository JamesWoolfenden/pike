# syntax=docker/dockerfile:1.7
#
# Multi-stage build:
#   1. Build pike from source so the image is reproducible from the checked-out SHA
#      (no fetching the "latest" release at build time).
#   2. Ship a minimal Alpine runtime with a non-root user and only the tools
#      entrypoint.sh actually needs.



# ---- builder ----------------------------------------------------------------
FROM golang:1.27.2-alpine3.24@sha256:f92b6ef800e499660581efdabdf25d9d817a9d124eaf900924f0504e7e27e12d AS builder

WORKDIR /src

# Cache modules separately from source so dep-only changes don't bust the layer.
COPY go.mod go.sum ./
RUN go mod download

# Only the Go sources needed by the package build are copied into the image.
# Keeping the copy explicit avoids sending the entire build context into a layer.
COPY *.go ./

# Note: pike.Version is declared as a `const` in src/version.go, so it cannot
# be overridden at link time. Change it to a `var` if you want build-time
# version injection via `-ldflags -X`.
RUN CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags "-s -w" \
    -o /out/pike .

# ---- runtime ----------------------------------------------------------------
FROM alpine:3.24@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a

# entrypoint.sh needs bash; terraform/tofu flows often want git + ca-certificates.
RUN apk add --no-cache bash=5.2.37-r0 ca-certificates=20250619-r0 git=2.49.1-r0 \
    && adduser -D -u 10001 -h /home/pike pike

LABEL org.opencontainers.image.title="pike" \
      org.opencontainers.image.description="Determine the minimum IAM permissions required to run OpenTofu/Terraform infrastructure code." \
      org.opencontainers.image.source="https://github.com/JamesWoolfenden/pike" \
      org.opencontainers.image.licenses="Apache-2.0" \
      org.opencontainers.image.authors="JamesWoolfenden"

COPY --from=builder /out/pike /usr/bin/pike
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

USER pike
WORKDIR /home/pike

ENTRYPOINT ["/entrypoint.sh"]
