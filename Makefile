#DEV
build-proto:
	protoc internal/gRPC/proto/pipeline.proto --go-grpc_out=internal/gRPC

run-app:
	go run cmd/main.go

