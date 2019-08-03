all: server client

server:
	go build -o server ./cmd/server/main.go

client:
	go build -o client ./cmd/client/main.go

clean:
	rm -rf server client

proto:
	protoc pkg/pb/taxsvc.proto --go_out=plugins=grpc:.
generate:
	go generate ./...
