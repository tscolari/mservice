.PHONY: server client clean deploy

all: server client

server:
	go build -o server ./cmd/server/main.go

client:
	go build -o client ./cmd/client/main.go

clean:
	rm -rf server client

deploy:
	kubectl apply -f deploy/deployment.yaml
	kubectl apply -f deploy/service.yaml
	kubectl apply -f deploy/ingress.yaml

proto:
	protoc pkg/pb/taxsvc.proto --go_out=plugins=grpc:.
generate:
	go generate ./...
