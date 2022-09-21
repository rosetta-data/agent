// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ingest_svc.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	events "github.com/nginx/agent/sdk/v2/proto/events"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("ingest_svc.proto", fileDescriptor_e87a1d7991134362) }

var fileDescriptor_e87a1d7991134362 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcc, 0x4b, 0x4f,
	0x2d, 0x2e, 0x89, 0x2f, 0x2e, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0x33,
	0xd5, 0xcb, 0x4b, 0xcf, 0xcc, 0xab, 0xd0, 0x4b, 0x4c, 0x4f, 0xcd, 0x2b, 0xd1, 0x2b, 0x4e, 0xc9,
	0x96, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xab, 0x48, 0x2a, 0x4d, 0xd3, 0x4f,
	0xcd, 0x2d, 0x28, 0xa9, 0x84, 0x68, 0x90, 0xe2, 0xcd, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e, 0x86,
	0x72, 0x85, 0x52, 0xcb, 0x52, 0xf3, 0x4a, 0x8a, 0xf5, 0xc1, 0x14, 0x44, 0xcc, 0x68, 0x27, 0x23,
	0x17, 0x87, 0x27, 0xd8, 0xa2, 0xd4, 0x22, 0xa1, 0x10, 0x2e, 0xe1, 0xe0, 0x92, 0xa2, 0xd4, 0xc4,
	0x5c, 0x5f, 0x88, 0xbe, 0xa0, 0xd4, 0x82, 0xfc, 0xa2, 0x12, 0x21, 0x45, 0x3d, 0x4c, 0x8b, 0xf5,
	0x50, 0x94, 0x48, 0x89, 0xe9, 0x41, 0xdc, 0xa1, 0x07, 0x73, 0x87, 0x9e, 0x2b, 0xc8, 0x1d, 0x4a,
	0x0c, 0x1a, 0x8c, 0x42, 0xe1, 0x5c, 0x82, 0x10, 0x53, 0x5d, 0x41, 0xf6, 0x42, 0xcd, 0x54, 0xc3,
	0x66, 0x26, 0xc4, 0x7d, 0x7a, 0x48, 0xea, 0xf0, 0x19, 0xec, 0x64, 0x7e, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46, 0x69, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0x0d, 0xd6, 0x07, 0x1b, 0xac, 0x5f, 0x9c, 0x92, 0xad, 0x5f,
	0x66, 0x04, 0x09, 0x20, 0x6b, 0x88, 0x29, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x56, 0x21, 0x22, 0x28, 0x63, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IngesterClient is the client API for Ingester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IngesterClient interface {
	// A client-to-server streaming RPC to deliver high volume metrics reports.
	StreamMetricsReport(ctx context.Context, opts ...grpc.CallOption) (Ingester_StreamMetricsReportClient, error)
	// A client-to-server streaming RPC to deliver high volume event reports.
	StreamEventReport(ctx context.Context, opts ...grpc.CallOption) (Ingester_StreamEventReportClient, error)
}

type ingesterClient struct {
	cc *grpc.ClientConn
}

func NewIngesterClient(cc *grpc.ClientConn) IngesterClient {
	return &ingesterClient{cc}
}

func (c *ingesterClient) StreamMetricsReport(ctx context.Context, opts ...grpc.CallOption) (Ingester_StreamMetricsReportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ingester_serviceDesc.Streams[0], "/f5.nginx.agent.sdk.Ingester/StreamMetricsReport", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingesterStreamMetricsReportClient{stream}
	return x, nil
}

type Ingester_StreamMetricsReportClient interface {
	Send(*MetricsReport) error
	CloseAndRecv() (*types.Empty, error)
	grpc.ClientStream
}

type ingesterStreamMetricsReportClient struct {
	grpc.ClientStream
}

func (x *ingesterStreamMetricsReportClient) Send(m *MetricsReport) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingesterStreamMetricsReportClient) CloseAndRecv() (*types.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(types.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ingesterClient) StreamEventReport(ctx context.Context, opts ...grpc.CallOption) (Ingester_StreamEventReportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ingester_serviceDesc.Streams[1], "/f5.nginx.agent.sdk.Ingester/StreamEventReport", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingesterStreamEventReportClient{stream}
	return x, nil
}

type Ingester_StreamEventReportClient interface {
	Send(*events.EventReport) error
	CloseAndRecv() (*types.Empty, error)
	grpc.ClientStream
}

type ingesterStreamEventReportClient struct {
	grpc.ClientStream
}

func (x *ingesterStreamEventReportClient) Send(m *events.EventReport) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingesterStreamEventReportClient) CloseAndRecv() (*types.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(types.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IngesterServer is the server API for Ingester service.
type IngesterServer interface {
	// A client-to-server streaming RPC to deliver high volume metrics reports.
	StreamMetricsReport(Ingester_StreamMetricsReportServer) error
	// A client-to-server streaming RPC to deliver high volume event reports.
	StreamEventReport(Ingester_StreamEventReportServer) error
}

// UnimplementedIngesterServer can be embedded to have forward compatible implementations.
type UnimplementedIngesterServer struct {
}

func (*UnimplementedIngesterServer) StreamMetricsReport(srv Ingester_StreamMetricsReportServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMetricsReport not implemented")
}
func (*UnimplementedIngesterServer) StreamEventReport(srv Ingester_StreamEventReportServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEventReport not implemented")
}

func RegisterIngesterServer(s *grpc.Server, srv IngesterServer) {
	s.RegisterService(&_Ingester_serviceDesc, srv)
}

func _Ingester_StreamMetricsReport_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngesterServer).StreamMetricsReport(&ingesterStreamMetricsReportServer{stream})
}

type Ingester_StreamMetricsReportServer interface {
	SendAndClose(*types.Empty) error
	Recv() (*MetricsReport, error)
	grpc.ServerStream
}

type ingesterStreamMetricsReportServer struct {
	grpc.ServerStream
}

func (x *ingesterStreamMetricsReportServer) SendAndClose(m *types.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingesterStreamMetricsReportServer) Recv() (*MetricsReport, error) {
	m := new(MetricsReport)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ingester_StreamEventReport_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngesterServer).StreamEventReport(&ingesterStreamEventReportServer{stream})
}

type Ingester_StreamEventReportServer interface {
	SendAndClose(*types.Empty) error
	Recv() (*events.EventReport, error)
	grpc.ServerStream
}

type ingesterStreamEventReportServer struct {
	grpc.ServerStream
}

func (x *ingesterStreamEventReportServer) SendAndClose(m *types.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingesterStreamEventReportServer) Recv() (*events.EventReport, error) {
	m := new(events.EventReport)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Ingester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "f5.nginx.agent.sdk.Ingester",
	HandlerType: (*IngesterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMetricsReport",
			Handler:       _Ingester_StreamMetricsReport_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamEventReport",
			Handler:       _Ingester_StreamEventReport_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "ingest_svc.proto",
}
