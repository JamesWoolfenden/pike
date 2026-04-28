# syntax=docker/dockerfile:1.7
#
# Multi-stage build:
#   1. Build pike from source so the image is reproducible from the checked-out SHA
#      (no fetching the "latest" release at build time).
#   2. Ship a minimal Alpine runtime with a non-root user and only the tools
#      entrypoint.sh actually needs.



# ---- builder ----------------------------------------------------------------
FROM golang:1.26.2-alpine3.22@sha256:7ef941168f213aa115df2e61364d67682129e99dc8188b734139dea862cc7d31 AS builder

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
FROM alpine:3.23@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11

# entrypoint.sh needs bash; terraform/tofu flows often want git + ca-certificates.
RUN apk add --no-cache bash ca-certificates git \
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
