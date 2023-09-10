package server

import (
	"bytes"
	"errors"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
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
			fileSize = float32(fileData.GetFileSize())
			if fileSize == 0 {
				return errors.New("empty file returned from pipeline")
			}
		} else if fileSize != float32(f.Len()) {
			n, err := f.Write(fileData.Chunk)
			if err == io.EOF {
				break
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

func (s server) ProcessAudioClient(client pipeline.Pipeline_AudioChannelClient, filePath string, targetLang *pipeline.Language) error {
	//ctx := stream.Context()
	var f *os.File
	var err error
	fileBytes := make([]byte, 4*1024)

	f, err = os.Open(filePath)
	if err != nil {
		return err
	}

	// send file to server
	i := 0
	for {
		// first pass
		if i == 0 {
			s, err := f.Stat()
			if s.Size() == 0 || err != nil {
				return errors.New("empty file found from user")
			} else {
				_, err := f.Read(fileBytes)
				if err == io.EOF {
					size := s.Size()
					name := s.Name()
					client.Send(&pipeline.FileData{
						FileSize:   &size,
						FileName:   &name,
						TargetLang: targetLang,
						Chunk:      fileBytes,
					})
					break
				} else if err != nil {
					return err
				} else {
					size := s.Size()
					name := s.Name()
					client.Send(&pipeline.FileData{
						FileSize:   &size,
						FileName:   &name,
						TargetLang: targetLang,
						Chunk:      fileBytes,
					})
				}
			}
		} else {
			_, err := f.Read(fileBytes)
			if err == io.EOF {
				client.Send(&pipeline.FileData{
					Chunk: fileBytes,
				})
				break
			} else if err != nil {
				return err
			} else {
				client.Send(&pipeline.FileData{
					Chunk: fileBytes,
				})
			}
		}

		i++
	}
	f.Close()
	os.Remove(filePath)

	fileBytes = nil
	var fileSize int64

	// ! Recieve from server now

	i = 0
	for {
		fileData, err := client.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if i == 0 {
			fileSize = fileData.GetFileSize()
			if fileSize <= 0 {
				return errors.New("empty file returned from pipeline")
			} else {
				f, err = os.CreateTemp("", "*.srt")
				n, err := f.Write(fileData.Chunk)
				if err == io.EOF {
					break
				} else if err == bytes.ErrTooLarge {
					return err
				} else if err != nil || n == 0 {
					return err
				}
			}
		} else if fileSize > 0 {
			n, err := f.Write(fileData.Chunk)
			if err == io.EOF {
				break
			} else if err == bytes.ErrTooLarge {
				return err
			} else if err != nil || n == 0 {
				return err
			}
			fileSize -= int64(len(fileData.Chunk))
		}

		i++
	}

	return nil
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
