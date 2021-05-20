// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: client_model.proto

package client

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     string `protobuf:"bytes,100,opt,name=AppId,proto3" json:"AppId,omitempty"`
	Namespace string `protobuf:"bytes,101,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_client_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_client_model_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Client) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type FlyingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId         string `protobuf:"bytes,100,opt,name=AppId,proto3" json:"AppId,omitempty"`
	NamespaceName string `protobuf:"bytes,101,opt,name=NamespaceName,proto3" json:"NamespaceName,omitempty"`
	ReleaseKey    string `protobuf:"bytes,102,opt,name=ReleaseKey,proto3" json:"ReleaseKey,omitempty"`
	Format        string `protobuf:"bytes,103,opt,name=Format,proto3" json:"Format,omitempty"`
	Value         string `protobuf:"bytes,104,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *FlyingConfig) Reset() {
	*x = FlyingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlyingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlyingConfig) ProtoMessage() {}

func (x *FlyingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_client_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlyingConfig.ProtoReflect.Descriptor instead.
func (*FlyingConfig) Descriptor() ([]byte, []int) {
	return file_client_model_proto_rawDescGZIP(), []int{1}
}

func (x *FlyingConfig) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *FlyingConfig) GetNamespaceName() string {
	if x != nil {
		return x.NamespaceName
	}
	return ""
}

func (x *FlyingConfig) GetReleaseKey() string {
	if x != nil {
		return x.ReleaseKey
	}
	return ""
}

func (x *FlyingConfig) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *FlyingConfig) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_client_model_proto protoreflect.FileDescriptor

var file_client_model_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x06,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x46,
	0x6c, 0x79, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2f, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79,
	0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_model_proto_rawDescOnce sync.Once
	file_client_model_proto_rawDescData = file_client_model_proto_rawDesc
)

func file_client_model_proto_rawDescGZIP() []byte {
	file_client_model_proto_rawDescOnce.Do(func() {
		file_client_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_model_proto_rawDescData)
	})
	return file_client_model_proto_rawDescData
}

var file_client_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_model_proto_goTypes = []interface{}{
	(*Client)(nil),       // 0: client.Client
	(*FlyingConfig)(nil), // 1: client.FlyingConfig
}
var file_client_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_model_proto_init() }
func file_client_model_proto_init() {
	if File_client_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_client_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlyingConfig); i {
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
			RawDescriptor: file_client_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_model_proto_goTypes,
		DependencyIndexes: file_client_model_proto_depIdxs,
		MessageInfos:      file_client_model_proto_msgTypes,
	}.Build()
	File_client_model_proto = out.File
	file_client_model_proto_rawDesc = nil
	file_client_model_proto_goTypes = nil
	file_client_model_proto_depIdxs = nil
}
