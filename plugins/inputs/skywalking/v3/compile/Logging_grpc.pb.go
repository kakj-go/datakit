// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package compile

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LogReportServiceClient is the client API for LogReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogReportServiceClient interface {
	// Recommend to report log data in a stream mode.
	// The service/instance/endpoint of the log could share the previous value if
	// they are not set. Reporting the logs of same service in the batch mode
	// could reduce the network cost.
	Collect(ctx context.Context, opts ...grpc.CallOption) (LogReportService_CollectClient, error)
}

type logReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogReportServiceClient(cc grpc.ClientConnInterface) LogReportServiceClient {
	return &logReportServiceClient{cc}
}

func (c *logReportServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (LogReportService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogReportService_serviceDesc.Streams[0], "/skywalking.v3.LogReportService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &logReportServiceCollectClient{stream}
	return x, nil
}

type LogReportService_CollectClient interface {
	Send(*LogData) error
	CloseAndRecv() (*Commands, error)
	grpc.ClientStream
}

type logReportServiceCollectClient struct {
	grpc.ClientStream
}

func (x *logReportServiceCollectClient) Send(m *LogData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logReportServiceCollectClient) CloseAndRecv() (*Commands, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Commands)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogReportServiceServer is the server API for LogReportService service.
// All implementations must embed UnimplementedLogReportServiceServer
// for forward compatibility.
type LogReportServiceServer interface {
	// Recommend to report log data in a stream mode.
	// The service/instance/endpoint of the log could share the previous value if
	// they are not set. Reporting the logs of same service in the batch mode
	// could reduce the network cost.
	Collect(LogReportService_CollectServer) error
	mustEmbedUnimplementedLogReportServiceServer()
}

// UnimplementedLogReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogReportServiceServer struct{}

func (UnimplementedLogReportServiceServer) Collect(LogReportService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedLogReportServiceServer) mustEmbedUnimplementedLogReportServiceServer() {}

// UnsafeLogReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogReportServiceServer will
// result in compilation errors.
type UnsafeLogReportServiceServer interface {
	mustEmbedUnimplementedLogReportServiceServer()
}

func RegisterLogReportServiceServer(s grpc.ServiceRegistrar, srv LogReportServiceServer) {
	s.RegisterService(&_LogReportService_serviceDesc, srv)
}

func _LogReportService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogReportServiceServer).Collect(&logReportServiceCollectServer{stream})
}

type LogReportService_CollectServer interface {
	SendAndClose(*Commands) error
	Recv() (*LogData, error)
	grpc.ServerStream
}

type logReportServiceCollectServer struct {
	grpc.ServerStream
}

func (x *logReportServiceCollectServer) SendAndClose(m *Commands) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logReportServiceCollectServer) Recv() (*LogData, error) {
	m := new(LogData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LogReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.v3.LogReportService",
	HandlerType: (*LogReportServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _LogReportService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "plugins/inputs/skywalking/v3/proto/logging/Logging.proto",
}
