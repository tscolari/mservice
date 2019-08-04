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
  ./server --port 8080 --tax-value 0.2
```

## Client:

```
  ./server --addr :8080 --method add --value 100
  > Result: 120


  ./server --addr :8080 --method sub --value 120
  > Result: 100
```

## Tests:

```
 go get github.com/onsi/ginkgo/ginkgo
 ginkgo -r

 # or

 go test ./...
```

## Deployment:

To deploy locally to minikube:

```
minikube start
minikube addons enable ingress

# once the cluster is up

make deploy
```
