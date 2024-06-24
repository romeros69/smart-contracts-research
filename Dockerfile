# Step 1: Modules caching
FROM golang:1.18-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.18-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY smart-contracts-research /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/main

# Step 3: Final
FROM scratch
COPY --from=builder /app/configs /configs
COPY --from=builder /bin/app /app
CMD ["/app"]