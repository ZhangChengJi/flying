// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: release_service.proto

package release

import (
	context "context"
	"flying-config/proto/common"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_release_service_proto protoreflect.FileDescriptor

var file_release_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x1a, 0x13, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xa5, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10,
	0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x10, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10,
	0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x2e, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_release_service_proto_goTypes = []interface{}{
	(*Release)(nil),          // 0: release.Release
	(*common.Request)(nil),   // 1: common.Request
	(*common.Response)(nil),  // 2: common.Response
	(*common.Responses)(nil), // 3: common.Responses
}
var file_release_service_proto_depIdxs = []int32{
	0, // 0: release.ReleaseService.Create:input_type -> release.Release
	0, // 1: release.ReleaseService.Delete:input_type -> release.Release
	0, // 2: release.ReleaseService.DeleteById:input_type -> release.Release
	0, // 3: release.ReleaseService.Update:input_type -> release.Release
	0, // 4: release.ReleaseService.Find:input_type -> release.Release
	1, // 5: release.ReleaseService.Lists:input_type -> common.Request
	2, // 6: release.ReleaseService.Create:output_type -> common.Response
	2, // 7: release.ReleaseService.Delete:output_type -> common.Response
	2, // 8: release.ReleaseService.DeleteById:output_type -> common.Response
	2, // 9: release.ReleaseService.Update:output_type -> common.Response
	2, // 10: release.ReleaseService.Find:output_type -> common.Response
	3, // 11: release.ReleaseService.Lists:output_type -> common.Responses
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_release_service_proto_init() }
func file_release_service_proto_init() {
	if File_release_service_proto != nil {
		return
	}
	file_release_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_release_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_release_service_proto_goTypes,
		DependencyIndexes: file_release_service_proto_depIdxs,
	}.Build()
	File_release_service_proto = out.File
	file_release_service_proto_rawDesc = nil
	file_release_service_proto_goTypes = nil
	file_release_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReleaseServiceClient is the client API for ReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReleaseServiceClient interface {
	Create(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error)
	Delete(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error)
	DeleteById(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error)
	Update(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error)
	Find(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error)
	Lists(ctx context.Context, in *common.Request, opts ...grpc.CallOption) (*common.Responses, error)
}

type releaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReleaseServiceClient(cc grpc.ClientConnInterface) ReleaseServiceClient {
	return &releaseServiceClient{cc}
}

func (c *releaseServiceClient) Create(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/release.ReleaseService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Delete(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/release.ReleaseService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) DeleteById(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/release.ReleaseService/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Update(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/release.ReleaseService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Find(ctx context.Context, in *Release, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/release.ReleaseService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Lists(ctx context.Context, in *common.Request, opts ...grpc.CallOption) (*common.Responses, error) {
	out := new(common.Responses)
	err := c.cc.Invoke(ctx, "/release.ReleaseService/Lists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReleaseServiceServer is the server API for ReleaseService service.
type ReleaseServiceServer interface {
	Create(context.Context, *Release) (*common.Response, error)
	Delete(context.Context, *Release) (*common.Response, error)
	DeleteById(context.Context, *Release) (*common.Response, error)
	Update(context.Context, *Release) (*common.Response, error)
	Find(context.Context, *Release) (*common.Response, error)
	Lists(context.Context, *common.Request) (*common.Responses, error)
}

// UnimplementedReleaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReleaseServiceServer struct {
}

func (*UnimplementedReleaseServiceServer) Create(context.Context, *Release) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedReleaseServiceServer) Delete(context.Context, *Release) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedReleaseServiceServer) DeleteById(context.Context, *Release) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (*UnimplementedReleaseServiceServer) Update(context.Context, *Release) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedReleaseServiceServer) Find(context.Context, *Release) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (*UnimplementedReleaseServiceServer) Lists(context.Context, *common.Request) (*common.Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lists not implemented")
}

func RegisterReleaseServiceServer(s *grpc.Server, srv ReleaseServiceServer) {
	s.RegisterService(&_ReleaseService_serviceDesc, srv)
}

func _ReleaseService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Release)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/release.ReleaseService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Create(ctx, req.(*Release))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Release)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/release.ReleaseService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Delete(ctx, req.(*Release))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Release)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/release.ReleaseService/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).DeleteById(ctx, req.(*Release))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Release)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/release.ReleaseService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Update(ctx, req.(*Release))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Release)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/release.ReleaseService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Find(ctx, req.(*Release))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Lists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Lists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/release.ReleaseService/Lists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Lists(ctx, req.(*common.Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReleaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "release.ReleaseService",
	HandlerType: (*ReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReleaseService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReleaseService_Delete_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _ReleaseService_DeleteById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReleaseService_Update_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _ReleaseService_Find_Handler,
		},
		{
			MethodName: "Lists",
			Handler:    _ReleaseService_Lists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "release_service.proto",
}
