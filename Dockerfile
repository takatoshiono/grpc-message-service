FROM golang:1.10 AS build

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/takatoshiono/grpc-message-service
COPY . .

RUN dep ensure

WORKDIR /go/src/github.com/takatoshiono/grpc-message-service/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go install -v

FROM alpine:latest

WORKDIR /app
COPY --from=build /go/bin/server .

EXPOSE 50101

CMD ["./server"]
