// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: option.proto

package option

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CommandOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CommandOptions) Reset() {
	*x = CommandOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_option_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandOptions) ProtoMessage() {}

func (x *CommandOptions) ProtoReflect() protoreflect.Message {
	mi := &file_option_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandOptions.ProtoReflect.Descriptor instead.
func (*CommandOptions) Descriptor() ([]byte, []int) {
	return file_option_proto_rawDescGZIP(), []int{0}
}

func (x *CommandOptions) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ClientOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *ClientOptions) Reset() {
	*x = ClientOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_option_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientOptions) ProtoMessage() {}

func (x *ClientOptions) ProtoReflect() protoreflect.Message {
	mi := &file_option_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientOptions.ProtoReflect.Descriptor instead.
func (*ClientOptions) Descriptor() ([]byte, []int) {
	return file_option_proto_rawDescGZIP(), []int{1}
}

func (x *ClientOptions) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

var file_option_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*CommandOptions)(nil),
		Field:         50001,
		Name:          "command_options",
		Tag:           "bytes,50001,opt,name=command_options",
		Filename:      "option.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*ClientOptions)(nil),
		Field:         50001,
		Name:          "client_options",
		Tag:           "bytes,50001,opt,name=client_options",
		Filename:      "option.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional CommandOptions command_options = 50001;
	E_CommandOptions = &file_option_proto_extTypes[0]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional ClientOptions client_options = 50001;
	E_ClientOptions = &file_option_proto_extTypes[1]
)

var File_option_proto protoreflect.FileDescriptor

var file_option_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x24, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2f, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x5a, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x55, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_option_proto_rawDescOnce sync.Once
	file_option_proto_rawDescData = file_option_proto_rawDesc
)

func file_option_proto_rawDescGZIP() []byte {
	file_option_proto_rawDescOnce.Do(func() {
		file_option_proto_rawDescData = protoimpl.X.CompressGZIP(file_option_proto_rawDescData)
	})
	return file_option_proto_rawDescData
}

var file_option_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_option_proto_goTypes = []interface{}{
	(*CommandOptions)(nil),             // 0: CommandOptions
	(*ClientOptions)(nil),              // 1: ClientOptions
	(*descriptorpb.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
	(*descriptorpb.FileOptions)(nil),   // 3: google.protobuf.FileOptions
}
var file_option_proto_depIdxs = []int32{
	2, // 0: command_options:extendee -> google.protobuf.MethodOptions
	3, // 1: client_options:extendee -> google.protobuf.FileOptions
	0, // 2: command_options:type_name -> CommandOptions
	1, // 3: client_options:type_name -> ClientOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_option_proto_init() }
func file_option_proto_init() {
	if File_option_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_option_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandOptions); i {
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
		file_option_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientOptions); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_option_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_option_proto_goTypes,
		DependencyIndexes: file_option_proto_depIdxs,
		MessageInfos:      file_option_proto_msgTypes,
		ExtensionInfos:    file_option_proto_extTypes,
	}.Build()
	File_option_proto = out.File
	file_option_proto_rawDesc = nil
	file_option_proto_goTypes = nil
	file_option_proto_depIdxs = nil
}
