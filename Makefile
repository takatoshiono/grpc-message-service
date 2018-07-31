IMAGE_VERSION := 0.2.0

protoc:
	protoc --descriptor_set_out=proto/api_descriptor.pb --go_out=plugins=grpc:. proto/message.proto

run-server:
	go run cmd/server/main.go cmd/server/server.go

docker-build:
	docker build -t grpc-message-service-server:$(IMAGE_VERSION) .

docker-run:
	docker run --rm -p 50101:50101 grpc-message-service-server:$(IMAGE_VERSION)

gcloud-docker-push:
	docker tag grpc-message-service-server:$(IMAGE_VERSION) asia.gcr.io/grpc-message-service/server:$(IMAGE_VERSION)
	gcloud docker -- push asia.gcr.io/grpc-message-service/server:$(IMAGE_VERSION)

gcloud-cluster-create:
	gcloud container clusters create message-cluster --machine-type=g1-small
	gcloud container clusters get-credentials message-cluster

gcloud-deploy:
	kubectl create -f deployment.yaml
	kubectl create -f service.yaml
	kubectl get service -w message-server

gcloud-deploy-delete:
	kubectl delete service message-server
	kubectl delete deployment message-server
	gcloud container clusters delete message-cluster

gcloud-deploy-endpoint: # only once
	gcloud endpoints services deploy proto/api_descriptor.pb api_config.yaml

gcloud-delete-endpoint: # will delete after 30 days
	gcloud endpoints services delete message-server.endpoints.grpc-message-service.cloud.goog

gcloud-deploy-with-endpoint:
	kubectl create -f deployment-esp.yml
	kubectl create -f service-esp.yml
	kubectl get service -w

gcloud-deploy-with-endpoint-delete:
	kubectl delete service message-server
	kubectl delete deployment message-server
	gcloud container clusters delete message-cluster

deploy-image:
	kubectl set image deployment/message-server message-server=asia.gcr.io/grpc-message-service/server:$(IMAGE_VERSION)
