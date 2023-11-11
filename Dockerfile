FROM golang:1.21.4-alpine

WORKDIR /app

RUN --mount=type=cache,mode=0755,target=/go/pkg/mod \
     --mount=type=bind,source=go.sum,target=go.sum \
     --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

CMD ["/app/main"]