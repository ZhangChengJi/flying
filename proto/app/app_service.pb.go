// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: app_service.proto

package proto

import (
	context "context"
	"flying/proto/common"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

var File_app_service_proto protoreflect.FileDescriptor

var file_app_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x70, 0x1a, 0x0f, 0x61, 0x70, 0x70, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x87, 0x02, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x1a, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x08, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x1a, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x42, 0x0d,
	0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_app_service_proto_goTypes = []interface{}{
	(*App)(nil),              // 0: app.App
	(*emptypb.Empty)(nil),    // 1: google.protobuf.Empty
	(*common.Request)(nil),   // 2: common.Request
	(*common.Response)(nil),  // 3: common.Response
	(*common.Responses)(nil), // 4: common.Responses
}
var file_app_service_proto_depIdxs = []int32{
	0, // 0: app.AppService.Create:input_type -> app.App
	0, // 1: app.AppService.Delete:input_type -> app.App
	0, // 2: app.AppService.DeleteById:input_type -> app.App
	0, // 3: app.AppService.Update:input_type -> app.App
	1, // 4: app.AppService.Find:input_type -> google.protobuf.Empty
	2, // 5: app.AppService.Lists:input_type -> common.Request
	3, // 6: app.AppService.Create:output_type -> common.Response
	3, // 7: app.AppService.Delete:output_type -> common.Response
	3, // 8: app.AppService.DeleteById:output_type -> common.Response
	3, // 9: app.AppService.Update:output_type -> common.Response
	3, // 10: app.AppService.Find:output_type -> common.Response
	4, // 11: app.AppService.Lists:output_type -> common.Responses
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_service_proto_init() }
func file_app_service_proto_init() {
	if File_app_service_proto != nil {
		return
	}
	file_app_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_service_proto_goTypes,
		DependencyIndexes: file_app_service_proto_depIdxs,
	}.Build()
	File_app_service_proto = out.File
	file_app_service_proto_rawDesc = nil
	file_app_service_proto_goTypes = nil
	file_app_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppServiceClient interface {
	Create(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Response, error)
	Delete(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Response, error)
	DeleteById(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Response, error)
	Update(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Response, error)
	Find(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.Response, error)
	Lists(ctx context.Context, in *common.Request, opts ...grpc.CallOption) (*common.Responses, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) Create(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/app.AppService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Delete(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/app.AppService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteById(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/app.AppService/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Update(ctx context.Context, in *App, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/app.AppService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Find(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/app.AppService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) Lists(ctx context.Context, in *common.Request, opts ...grpc.CallOption) (*common.Responses, error) {
	out := new(common.Responses)
	err := c.cc.Invoke(ctx, "/app.AppService/Lists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
type AppServiceServer interface {
	Create(context.Context, *App) (*common.Response, error)
	Delete(context.Context, *App) (*common.Response, error)
	DeleteById(context.Context, *App) (*common.Response, error)
	Update(context.Context, *App) (*common.Response, error)
	Find(context.Context, *emptypb.Empty) (*common.Response, error)
	Lists(context.Context, *common.Request) (*common.Responses, error)
}

// UnimplementedAppServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (*UnimplementedAppServiceServer) Create(context.Context, *App) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAppServiceServer) Delete(context.Context, *App) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedAppServiceServer) DeleteById(context.Context, *App) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (*UnimplementedAppServiceServer) Update(context.Context, *App) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAppServiceServer) Find(context.Context, *emptypb.Empty) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (*UnimplementedAppServiceServer) Lists(context.Context, *common.Request) (*common.Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lists not implemented")
}

func RegisterAppServiceServer(s *grpc.Server, srv AppServiceServer) {
	s.RegisterService(&_AppService_serviceDesc, srv)
}

func _AppService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Create(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Delete(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteById(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Update(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Find(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_Lists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Lists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/Lists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Lists(ctx, req.(*common.Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AppService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AppService_Delete_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _AppService_DeleteById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AppService_Update_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _AppService_Find_Handler,
		},
		{
			MethodName: "Lists",
			Handler:    _AppService_Lists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app_service.proto",
}
