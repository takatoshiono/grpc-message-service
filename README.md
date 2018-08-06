# grpc-message-service

## server

```
$ make docker-build
$ make docker-run
```

## client

gRPC API

```
$ go run cmd/client/main.go -api-key API_KEY
```

JSON API

```
$ curl -H "x-api-key: API_KEY" -X POST http://HOST:PORT/v1/conversations
```

## deploy

### Prerequisites

- Install [Google Cloud SDK](https://cloud.google.com/sdk/) and setup
- Install [ghq](https://github.com/motemen/ghq)
- Clone [googleapis](https://github.com/googleapis/googleapis) using `ghq get`

### Create cluster

```
$ make cluster-create
```

### Create Endpoint

```
$ make deploy-endpoint
```

### Create Service

```
$ make deploy-with-esp
```
