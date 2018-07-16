protoc:
	protoc --go_out=plugins=grpc:. proto/message.proto

run-server:
	go run cmd/server/main.go cmd/server/server.go

docker-build:
	docker build -t grpc-message-service-server .

docker-run:
	docker run --rm -p 50101:50101 grpc-message-service-server:latest

docker-tag:
	docker tag grpc-message-service-server:latest asia.gcr.io/grpc-message-service/server:latest

gcloud-docker-push:
	gcloud docker -- push asia.gcr.io/grpc-message-service/server:latest
