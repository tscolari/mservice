[![CircleCI](https://circleci.com/gh/tscolari/mservice/tree/master.svg?style=svg)](https://circleci.com/gh/tscolari/mservice/tree/master)

# Mservice

Example of using go-kit gRPC for a micro-service.

Very simple client-server binaries, the server expose 2 methods for tax calculation:

* add(value) -> Returns the taxes add given value
* sub(value) -> Returns the value after having taxes subtracted from it

The value of the tax % is set on the server initialisation.

## Compilation:

```
  make
```

## Server:


```
  ./server --port 4443 --tax-value 0.2 --tls-cert ./certs/server.crt --tls-key ./certs/server.key --ca-cert ./certs/ca.crt
```

## Client:

```
  ./client --addr 127.0.0.1:4443 -m add --value 100 --tls-cert ./certs/client.crt --tls-key ./certs/client.key --ca-cert ./certs/ca.crt
  > Result: 120


  ./client --addr 127.0.0.1:4443 -m sub --value 120 --tls-cert ./certs/client.crt --tls-key ./certs/client.key --ca-cert ./certs/ca.crt
  > Result: 100
```

## Tests:

```
 go get github.com/onsi/ginkgo/ginkgo
 ginkgo -r

 # or

 go test ./...
```

## Authentication:

The service has a basic TLS server/client authentication:

```
# Requires certstrap (brew install certstrap)
make certs
```

Note that these certificates have hardcoded ips/dnses, so you might need to tune it if you want to run it outside 127.0.0.1. 
There are some example certs committed in `./certs`.

There's also a `--insecure` flag to both server and client to skip all TLS setup.

## Deployment:

To deploy locally to minikube:

```
minikube start
minikube addons enable ingress

# once the cluster is up

make deploy
```

Be aware that this will create a secret using the certificates from `./certs`.
