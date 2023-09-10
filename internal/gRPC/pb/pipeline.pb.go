// pkg/proto/upload.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: internal/gRPC/proto/pipeline.proto

package pipeline

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Language int32

const (
	Language_HINDI     Language = 0
	Language_MALAYALAM Language = 1
	Language_GUJARATI  Language = 2
	Language_MARATHI   Language = 3
	Language_TELUGU    Language = 4
	Language_BENGALI   Language = 5
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "HINDI",
		1: "MALAYALAM",
		2: "GUJARATI",
		3: "MARATHI",
		4: "TELUGU",
		5: "BENGALI",
	}
	Language_value = map[string]int32{
		"HINDI":     0,
		"MALAYALAM": 1,
		"GUJARATI":  2,
		"MARATHI":   3,
		"TELUGU":    4,
		"BENGALI":   5,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_gRPC_proto_pipeline_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_internal_gRPC_proto_pipeline_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_internal_gRPC_proto_pipeline_proto_rawDescGZIP(), []int{0}
}

type FileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSize   *int64    `protobuf:"varint,1,opt,name=fileSize,proto3,oneof" json:"fileSize,omitempty"`
	FileName   *string   `protobuf:"bytes,2,opt,name=fileName,proto3,oneof" json:"fileName,omitempty"`
	TargetLang *Language `protobuf:"varint,3,opt,name=targetLang,proto3,enum=proto.Language,oneof" json:"targetLang,omitempty"`
	Chunk      []byte    `protobuf:"bytes,16,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *FileData) Reset() {
	*x = FileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gRPC_proto_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileData) ProtoMessage() {}

func (x *FileData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gRPC_proto_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileData.ProtoReflect.Descriptor instead.
func (*FileData) Descriptor() ([]byte, []int) {
	return file_internal_gRPC_proto_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *FileData) GetFileSize() int64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *FileData) GetFileName() string {
	if x != nil && x.FileName != nil {
		return *x.FileName
	}
	return ""
}

func (x *FileData) GetTargetLang() Language {
	if x != nil && x.TargetLang != nil {
		return *x.TargetLang
	}
	return Language_HINDI
}

func (x *FileData) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     *string `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Successful bool    `protobuf:"varint,2,opt,name=successful,proto3" json:"successful,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_gRPC_proto_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_gRPC_proto_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_internal_gRPC_proto_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *UploadResponse) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

var File_internal_gRPC_proto_pipeline_proto protoreflect.FileDescriptor

var file_internal_gRPC_proto_pipeline_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x48,
	0x02, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x88, 0x01, 0x01,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x22,
	0x58, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x58, 0x0a, 0x08, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x49, 0x4e, 0x44, 0x49, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41, 0x4c, 0x41, 0x59, 0x41, 0x4c, 0x41, 0x4d, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x47, 0x55, 0x4a, 0x41, 0x52, 0x41, 0x54, 0x49, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x4d, 0x41, 0x52, 0x41, 0x54, 0x48, 0x49, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x45,
	0x4c, 0x55, 0x47, 0x55, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x45, 0x4e, 0x47, 0x41, 0x4c,
	0x49, 0x10, 0x05, 0x32, 0x82, 0x01, 0x0a, 0x08, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x3e, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x12, 0x36, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_gRPC_proto_pipeline_proto_rawDescOnce sync.Once
	file_internal_gRPC_proto_pipeline_proto_rawDescData = file_internal_gRPC_proto_pipeline_proto_rawDesc
)

func file_internal_gRPC_proto_pipeline_proto_rawDescGZIP() []byte {
	file_internal_gRPC_proto_pipeline_proto_rawDescOnce.Do(func() {
		file_internal_gRPC_proto_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_gRPC_proto_pipeline_proto_rawDescData)
	})
	return file_internal_gRPC_proto_pipeline_proto_rawDescData
}

var file_internal_gRPC_proto_pipeline_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_gRPC_proto_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_gRPC_proto_pipeline_proto_goTypes = []interface{}{
	(Language)(0),          // 0: proto.Language
	(*FileData)(nil),       // 1: proto.FileData
	(*UploadResponse)(nil), // 2: proto.UploadResponse
}
var file_internal_gRPC_proto_pipeline_proto_depIdxs = []int32{
	0, // 0: proto.FileData.targetLang:type_name -> proto.Language
	1, // 1: proto.Pipeline.SubtitleDownload:input_type -> proto.FileData
	1, // 2: proto.Pipeline.AudioChannel:input_type -> proto.FileData
	2, // 3: proto.Pipeline.SubtitleDownload:output_type -> proto.UploadResponse
	1, // 4: proto.Pipeline.AudioChannel:output_type -> proto.FileData
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_gRPC_proto_pipeline_proto_init() }
func file_internal_gRPC_proto_pipeline_proto_init() {
	if File_internal_gRPC_proto_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_gRPC_proto_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_gRPC_proto_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_internal_gRPC_proto_pipeline_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_internal_gRPC_proto_pipeline_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_gRPC_proto_pipeline_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_gRPC_proto_pipeline_proto_goTypes,
		DependencyIndexes: file_internal_gRPC_proto_pipeline_proto_depIdxs,
		EnumInfos:         file_internal_gRPC_proto_pipeline_proto_enumTypes,
		MessageInfos:      file_internal_gRPC_proto_pipeline_proto_msgTypes,
	}.Build()
	File_internal_gRPC_proto_pipeline_proto = out.File
	file_internal_gRPC_proto_pipeline_proto_rawDesc = nil
	file_internal_gRPC_proto_pipeline_proto_goTypes = nil
	file_internal_gRPC_proto_pipeline_proto_depIdxs = nil
}
