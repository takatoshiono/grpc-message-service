# grpc-message-service

## server

```
$ make docker-build
$ make docker-run
```

## client

```
$ go run cmd/client/main.go
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
