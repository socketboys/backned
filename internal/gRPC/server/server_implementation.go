package server

import (
	"google.golang.org/grpc"
	"net"
	pipeline "project-x/internal/gRPC/pb"
	"project-x/internal/http_server/utils"
)

type server struct {
	pipeline.UnimplementedPipelineServer
}

func RungRPC() {
	listener, err := net.Listen("tcp", ":36710")
	if err != nil {
		utils.Logger.Error(err.Error() + "Error running gRPC server!")
	}

	grpcServer := grpc.NewServer()
	pipeline.RegisterPipelineServer(grpcServer, &pipeline.UnimplementedPipelineServer{})
	if err = grpcServer.Serve(listener); err != nil {
		utils.Logger.Error(err.Error() + "Error listening on 36710!")
	}
}
