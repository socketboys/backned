package gRPC

import (
	"google.golang.org/grpc"
	"net"
	"project-x/internal/http_server/utils"
)

func RungRPC() {
	listener, err := net.Listen("tcp", ":36710")
	if err != nil {
		utils.Logger.Error(err.Error() + "Error running gRPC server!")
	}

	grpcServer := grpc.NewServer()
	RegisterPipelineServer(grpcServer, UnimplementedPipelineServer{})
	if err = grpcServer.Serve(listener); err != nil {
		utils.Logger.Error(err.Error() + "Error listening on 36710!")
	}
}
