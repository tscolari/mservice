FROM golang:latest as builder
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go install -a -installsuffix cgo \
    ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/server /bin/server
