#DEV
build-proto:
	protoc internal/gRPC/proto/pipeline.proto --go-grpc_out=internal/gRPC/pb --go_out=internal/gRPC/pb

delete-proto:
	rm internal/gRPC/pb/*.go

server:
	go run main.go

swagger-server:
	swagger serve --no-open --port=8080 -F=swagger docs/swagger.yaml

