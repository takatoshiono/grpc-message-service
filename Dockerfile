FROM golang:1.10

RUN go get -u -v google.golang.org/grpc

WORKDIR /go/src/github.com/takatoshiono/grpc-message-service
COPY . .

WORKDIR /go/src/github.com/takatoshiono/grpc-message-service/cmd/server
RUN go install -v

EXPOSE 50101

CMD ["server"]
