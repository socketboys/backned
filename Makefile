#DEV
build-proto:
	protoc internal/gRPC/proto/pipeline.proto --go-grpc_out=internal/gRPC/pb --go_out=internal/gRPC/pb

delete-proto:
	rm internal/gRPC/pb/*.go

run-app:
	go run cmd/main.go

