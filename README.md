# grpc-message-service

## Development

### server

```
$ docker-compose up --build
```

### client

gRPC API

```
$ go run cmd/client/main.go -api-key API_KEY
```

JSON API

```
$ curl -H "x-api-key: API_KEY" -X POST http://HOST:PORT/v1/conversations
```

## Deploy

### Prerequisites

- Install [Google Cloud SDK](https://cloud.google.com/sdk/) and setup
- Install [ghq](https://github.com/motemen/ghq)
- Clone [googleapis/googleapis](https://github.com/googleapis/googleapis) using `ghq get`
- Clone [google/protobuf](https://github.com/google/protobuf) using `ghq get`

### Docker Image

```
$ make docker-build
$ make deploy-image
```

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
