FROM --platform=$BUILDPLATFORM golang:1.23.4 AS builder

WORKDIR /app

RUN --mount=target=. \
        --mount=type=cache,target=/root/.cache/go-build \
        --mount=type=cache,target=/go/pkg \
        GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 \
    go build -o /deploy/server/service ./cmd/service/main.go && \
    cp -r config /deploy/server/config && \
    cp -r migrations /deploy/server/migrations
FROM alpine

WORKDIR /app
LABEL MATRESHKA_CONFIG_ENABLED=true

COPY --from=builder /deploy/server/ .

EXPOSE 80

ENTRYPOINT ["./service"]