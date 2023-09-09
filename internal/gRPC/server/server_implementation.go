package server

import (
	"bytes"
	"errors"
	"google.golang.org/grpc"
	"io"
	"net"
	pipeline "project-x/internal/gRPC/pb"
	"project-x/internal/http_server/utils"
)

type server struct {
	pipeline.UnimplementedPipelineServer
}

func (s server) GetSubtitles(stream pipeline.Pipeline_SubtitleDownloadServer) error {
	//ctx := stream.Context()
	f := bytes.Buffer{}

	var fileSize float32
	i := 0
	for {
		fileData, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if i == 0 {
			fileSize = fileData.GetFileSize()
			if fileSize == 0 {
				return errors.New("empty file returned from pipeline")
			}
		} else if fileSize != float32(f.Len()) {
			n, err := f.Write(fileData.Chunk)
			if err == io.EOF {
				return nil
			} else if err == bytes.ErrTooLarge {
				return err
			} else if err != nil || n == 0 {
				return err
			}
		}

		i++
	}

	return nil
}

//func (s server) ProcessAudio(stream pipeline.Pipeline_SubtitleDownloadServer) error {
//	ctx := stream.Context()
//	file, err := stream.Recv()
//
//	if file.GetMime()
//
//	return nil
//}

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
