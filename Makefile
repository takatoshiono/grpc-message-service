IMAGE_VERSION := 0.2.0
GOOGLEAPIS_DIR := $(shell ghq list -p googleapis)

protoc:
	protoc \
		--include_imports \
		--include_source_info \
		--proto_path=$(GOOGLEAPIS_DIR) \
		--proto_path=. \
		--descriptor_set_out=proto/api_descriptor.pb \
		--go_out=plugins=grpc:. \
		proto/message.proto

run-server:
	go run cmd/server/main.go cmd/server/server.go

docker-build:
	docker build -t grpc-message-service-server:$(IMAGE_VERSION) .

docker-run:
	docker run --rm -p 50101:50101 grpc-message-service-server:$(IMAGE_VERSION)

docker-push:
	docker tag grpc-message-service-server:$(IMAGE_VERSION) asia.gcr.io/grpc-message-service/server:$(IMAGE_VERSION)
	gcloud docker -- push asia.gcr.io/grpc-message-service/server:$(IMAGE_VERSION)

cluster-create:
	gcloud container clusters create message-cluster --num-nodes=1 --machine-type=g1-small
	gcloud container clusters get-credentials message-cluster

cluster-delete:
	gcloud container clusters delete message-cluster

deploy:
	echo "DEPRECATED: use deploy-with-endpoint"
	#kubectl create -f deployment.yaml
	#kubectl create -f service.yaml
	#kubectl get service -w message-server

delete-deploy:
	echo "DEPRECATED: use delete-deploy-with-endpoint"
	#kubectl delete service message-server
	#kubectl delete deployment message-server

deploy-endpoint:
	gcloud endpoints services deploy proto/api_descriptor.pb api_config.yaml

delete-endpoint: # will delete after 30 days
	gcloud endpoints services delete message-server.endpoints.grpc-message-service.cloud.goog

deploy-with-esp:
	kubectl create -f deployment-esp.yml
	kubectl create -f service-esp.yml
	kubectl get service -w

delete-deploy-with-esp:
	kubectl delete service message-server
	kubectl delete deployment message-server

deploy-image:
	kubectl set image deployment/message-server message-server=asia.gcr.io/grpc-message-service/server:$(IMAGE_VERSION)
