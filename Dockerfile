FROM golang:1.17.7-bullseye AS builder
WORKDIR /app/src
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -trimpath -o /app/app

FROM scratch
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
