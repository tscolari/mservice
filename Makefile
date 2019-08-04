.PHONY: server client clean deploy certs proto generate

all: server client

server:
	go build -o server ./cmd/server/*

client:
	go build -o client ./cmd/client/main.go

clean:
	rm -rf server client
	rm -rf certs/*

deploy:
	kubectl delete secret server-tls
	kubectl create secret generic server-tls --from-file=tls-cert=./certs/server.crt --from-file=tls-key=./certs/server.key --from-file=ca-cert=./certs/ca.crt
	kubectl apply -f deploy/deployment.yaml
	kubectl apply -f deploy/service.yaml
	kubectl apply -f deploy/ingress.yaml

proto:
	protoc pkg/pb/taxsvc.proto --go_out=plugins=grpc:.

generate:
	go generate ./...

image:
	docker build --rm -t tscolari/mservice:latest .
	docker push tscolari/mservice:latest

certs:
	rm -rf out/*
	certstrap init --common-name "ca" --passphrase ""
	certstrap request-cert --common-name server --passphrase "" -ip "127.0.0.1,10.102.74.251" -domain "localhost,mservice.default.svc.cluster.local"
	certstrap sign server --CA ca
	certstrap request-cert --common-name client --passphrase ""
	certstrap sign client --CA ca
	mv -f out/* certs
