// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package register

import (
	context "context"
	common "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/skywalking/v2/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ServiceInstancePingClient is the client API for ServiceInstancePing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceInstancePingClient interface {
	DoPing(ctx context.Context, in *ServiceInstancePingPkg, opts ...grpc.CallOption) (*common.Commands, error)
}

type serviceInstancePingClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceInstancePingClient(cc grpc.ClientConnInterface) ServiceInstancePingClient {
	return &serviceInstancePingClient{cc}
}

func (c *serviceInstancePingClient) DoPing(ctx context.Context, in *ServiceInstancePingPkg, opts ...grpc.CallOption) (*common.Commands, error) {
	out := new(common.Commands)
	err := c.cc.Invoke(ctx, "/skywalking.v2.ServiceInstancePing/doPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceInstancePingServer is the server API for ServiceInstancePing service.
// All implementations must embed UnimplementedServiceInstancePingServer
// for forward compatibility
type ServiceInstancePingServer interface {
	DoPing(context.Context, *ServiceInstancePingPkg) (*common.Commands, error)
	mustEmbedUnimplementedServiceInstancePingServer()
}

// UnimplementedServiceInstancePingServer must be embedded to have forward compatible implementations.
type UnimplementedServiceInstancePingServer struct {
}

func (UnimplementedServiceInstancePingServer) DoPing(context.Context, *ServiceInstancePingPkg) (*common.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoPing not implemented")
}
func (UnimplementedServiceInstancePingServer) mustEmbedUnimplementedServiceInstancePingServer() {}

// UnsafeServiceInstancePingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceInstancePingServer will
// result in compilation errors.
type UnsafeServiceInstancePingServer interface {
	mustEmbedUnimplementedServiceInstancePingServer()
}

func RegisterServiceInstancePingServer(s grpc.ServiceRegistrar, srv ServiceInstancePingServer) {
	s.RegisterService(&_ServiceInstancePing_serviceDesc, srv)
}

func _ServiceInstancePing_DoPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInstancePingPkg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceInstancePingServer).DoPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywalking.v2.ServiceInstancePing/doPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceInstancePingServer).DoPing(ctx, req.(*ServiceInstancePingPkg))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceInstancePing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.v2.ServiceInstancePing",
	HandlerType: (*ServiceInstancePingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "doPing",
			Handler:    _ServiceInstancePing_DoPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins/inputs/skywalking/v2/register/InstancePing.proto",
}
