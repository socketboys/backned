// pkg/proto/upload.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: internal/gRPC/proto/pipeline.proto

package pipeline

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Pipeline_SubtitleDownload_FullMethodName = "/proto.Pipeline/SubtitleDownload"
	Pipeline_AudioChannel_FullMethodName     = "/proto.Pipeline/AudioChannel"
)

// PipelineClient is the client API for Pipeline service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineClient interface {
	SubtitleDownload(ctx context.Context, opts ...grpc.CallOption) (Pipeline_SubtitleDownloadClient, error)
	AudioChannel(ctx context.Context, opts ...grpc.CallOption) (Pipeline_AudioChannelClient, error)
}

type pipelineClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineClient(cc grpc.ClientConnInterface) PipelineClient {
	return &pipelineClient{cc}
}

func (c *pipelineClient) SubtitleDownload(ctx context.Context, opts ...grpc.CallOption) (Pipeline_SubtitleDownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Pipeline_ServiceDesc.Streams[0], Pipeline_SubtitleDownload_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineSubtitleDownloadClient{stream}
	return x, nil
}

type Pipeline_SubtitleDownloadClient interface {
	Send(*FileData) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type pipelineSubtitleDownloadClient struct {
	grpc.ClientStream
}

func (x *pipelineSubtitleDownloadClient) Send(m *FileData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pipelineSubtitleDownloadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pipelineClient) AudioChannel(ctx context.Context, opts ...grpc.CallOption) (Pipeline_AudioChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &Pipeline_ServiceDesc.Streams[1], Pipeline_AudioChannel_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pipelineAudioChannelClient{stream}
	return x, nil
}

type Pipeline_AudioChannelClient interface {
	Send(*FileData) error
	Recv() (*FileData, error)
	grpc.ClientStream
}

type pipelineAudioChannelClient struct {
	grpc.ClientStream
}

func (x *pipelineAudioChannelClient) Send(m *FileData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pipelineAudioChannelClient) Recv() (*FileData, error) {
	m := new(FileData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PipelineServer is the server API for Pipeline service.
// All implementations must embed UnimplementedPipelineServer
// for forward compatibility
type PipelineServer interface {
	SubtitleDownload(Pipeline_SubtitleDownloadServer) error
	AudioChannel(Pipeline_AudioChannelServer) error
	mustEmbedUnimplementedPipelineServer()
}

// UnimplementedPipelineServer must be embedded to have forward compatible implementations.
type UnimplementedPipelineServer struct {
}

func (UnimplementedPipelineServer) SubtitleDownload(Pipeline_SubtitleDownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method SubtitleDownload not implemented")
}
func (UnimplementedPipelineServer) AudioChannel(Pipeline_AudioChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method AudioChannel not implemented")
}
func (UnimplementedPipelineServer) mustEmbedUnimplementedPipelineServer() {}

// UnsafePipelineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineServer will
// result in compilation errors.
type UnsafePipelineServer interface {
	mustEmbedUnimplementedPipelineServer()
}

func RegisterPipelineServer(s grpc.ServiceRegistrar, srv PipelineServer) {
	s.RegisterService(&Pipeline_ServiceDesc, srv)
}

func _Pipeline_SubtitleDownload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PipelineServer).SubtitleDownload(&pipelineSubtitleDownloadServer{stream})
}

type Pipeline_SubtitleDownloadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*FileData, error)
	grpc.ServerStream
}

type pipelineSubtitleDownloadServer struct {
	grpc.ServerStream
}

func (x *pipelineSubtitleDownloadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pipelineSubtitleDownloadServer) Recv() (*FileData, error) {
	m := new(FileData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Pipeline_AudioChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PipelineServer).AudioChannel(&pipelineAudioChannelServer{stream})
}

type Pipeline_AudioChannelServer interface {
	Send(*FileData) error
	Recv() (*FileData, error)
	grpc.ServerStream
}

type pipelineAudioChannelServer struct {
	grpc.ServerStream
}

func (x *pipelineAudioChannelServer) Send(m *FileData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pipelineAudioChannelServer) Recv() (*FileData, error) {
	m := new(FileData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Pipeline_ServiceDesc is the grpc.ServiceDesc for Pipeline service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pipeline_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Pipeline",
	HandlerType: (*PipelineServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubtitleDownload",
			Handler:       _Pipeline_SubtitleDownload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AudioChannel",
			Handler:       _Pipeline_AudioChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/gRPC/proto/pipeline.proto",
}