// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/analytics/a_rpc.proto

package analytics

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	io "stash.passkit.com/io/model/sdk/go/io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("io/analytics/a_rpc.proto", fileDescriptor_76450c92650569b9)
}

var fileDescriptor_76450c92650569b9 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xbd, 0x6e, 0x13, 0x41,
	0x10, 0x96, 0x1d, 0x09, 0x29, 0x07, 0x08, 0x74, 0xc4, 0x04, 0x5f, 0x10, 0x5a, 0xb9, 0x44, 0xf8,
	0x36, 0x38, 0x29, 0x50, 0x1a, 0x30, 0x4d, 0x94, 0xb8, 0xb1, 0x0c, 0x08, 0x89, 0x26, 0x9a, 0xec,
	0x8d, 0xcf, 0x8b, 0xf7, 0x76, 0x96, 0x9d, 0xbd, 0x44, 0x26, 0xa2, 0xe1, 0x11, 0x42, 0x47, 0x49,
	0xc9, 0xeb, 0xd0, 0xf0, 0x00, 0x88, 0x67, 0xa0, 0x03, 0xf9, 0x2c, 0x1f, 0x27, 0x48, 0xb5, 0x9a,
	0x6f, 0xbe, 0xfd, 0xe6, 0xe7, 0x9b, 0xe8, 0x9e, 0x26, 0x09, 0x16, 0xcc, 0x22, 0x68, 0xc5, 0x12,
	0x4e, 0xbc, 0x53, 0xa9, 0xf3, 0x14, 0x28, 0xde, 0xac, 0xe1, 0xe4, 0x7e, 0x4e, 0x94, 0x1b, 0x94,
	0xe0, 0xb4, 0x04, 0x6b, 0x29, 0x40, 0xd0, 0x64, 0x79, 0x45, 0x4c, 0x1e, 0x55, 0x8f, 0xea, 0xe7,
	0x68, 0xfb, 0x7c, 0x0e, 0x79, 0x8e, 0x5e, 0x92, 0xab, 0x18, 0x57, 0xb0, 0xbb, 0x9a, 0xa4, 0xa2,
	0xa2, 0x20, 0x2b, 0x3d, 0x3a, 0xf2, 0x41, 0xdb, 0x7c, 0x95, 0x1a, 0xfc, 0x6c, 0x45, 0x9b, 0xc3,
	0x75, 0xd1, 0xf8, 0x7b, 0x2b, 0xba, 0x91, 0x63, 0xf8, 0x0b, 0x6c, 0xa5, 0x9a, 0xd2, 0x3a, 0x9c,
	0xe0, 0xbb, 0x12, 0x39, 0x24, 0x9d, 0x7f, 0x50, 0x76, 0x64, 0x19, 0x7b, 0x9f, 0x5b, 0x97, 0xc3,
	0x8b, 0xa6, 0xdc, 0xcd, 0x43, 0x0c, 0xa2, 0x0e, 0x93, 0x9d, 0x09, 0x06, 0xaf, 0xf1, 0x0c, 0x59,
	0x80, 0x15, 0xf5, 0xa8, 0x22, 0x83, 0x00, 0xc7, 0xbd, 0x68, 0x63, 0x7f, 0x77, 0x2f, 0xde, 0x89,
	0xba, 0xaf, 0x18, 0xbd, 0x30, 0xa0, 0xe6, 0x2c, 0xa0, 0x0c, 0x33, 0xf2, 0xfa, 0x7d, 0x35, 0x4d,
	0x7a, 0xfc, 0x60, 0xc9, 0xd9, 0x8f, 0xb7, 0xa3, 0xce, 0x04, 0x15, 0xf9, 0x4c, 0x9c, 0x03, 0x0b,
	0x4b, 0x41, 0x4c, 0xa9, 0xb4, 0x59, 0xfa, 0xf1, 0xdb, 0x8f, 0x4f, 0xed, 0xbb, 0xf1, 0x56, 0x63,
	0xb9, 0x17, 0xca, 0x00, 0xf3, 0x51, 0xf6, 0xe1, 0xf9, 0xef, 0xf6, 0xe5, 0xf0, 0x4b, 0x3b, 0xfe,
	0xd5, 0x8a, 0x3a, 0x63, 0x60, 0x1e, 0xe9, 0x46, 0x67, 0x62, 0x38, 0x3e, 0x8a, 0x4f, 0x5e, 0xce,
	0x50, 0x5c, 0x99, 0x12, 0x06, 0x03, 0x8b, 0x05, 0x95, 0x22, 0x78, 0x50, 0x73, 0x11, 0x66, 0x28,
	0x1c, 0xfa, 0x29, 0xf9, 0x02, 0xac, 0x42, 0x41, 0x53, 0x31, 0x74, 0xce, 0xa0, 0x78, 0x0d, 0xc6,
	0x60, 0x10, 0x60, 0x33, 0x71, 0x58, 0xb9, 0x27, 0xc6, 0xb0, 0x10, 0x0e, 0x98, 0x91, 0xd3, 0xe4,
	0xc9, 0x2c, 0x04, 0xc7, 0x07, 0x52, 0x2e, 0x81, 0xb9, 0x0e, 0xa9, 0xa2, 0x42, 0x1a, 0xcc, 0xc1,
	0xc8, 0x80, 0xbe, 0xe0, 0x3e, 0x4d, 0xfb, 0x5c, 0x9e, 0xb2, 0xf2, 0xba, 0x72, 0xb2, 0xcf, 0xe8,
	0xcf, 0xb4, 0x42, 0xd9, 0x7b, 0x1a, 0xdd, 0x5a, 0x37, 0xf6, 0xa2, 0x74, 0x4b, 0x03, 0xe3, 0xed,
	0xb5, 0x58, 0x46, 0x8a, 0xd3, 0xb5, 0xa2, 0xa6, 0xe4, 0x0e, 0xaf, 0x18, 0xcf, 0x1a, 0x55, 0x06,
	0x1b, 0xbb, 0xe9, 0xe3, 0x87, 0xad, 0xf6, 0xe0, 0x36, 0x38, 0x67, 0xb4, 0xaa, 0xd6, 0x29, 0xdf,
	0x32, 0xd9, 0x83, 0xff, 0x90, 0xa8, 0xab, 0xa9, 0xd6, 0x1b, 0x8f, 0x9a, 0x46, 0xbf, 0x49, 0x39,
	0x00, 0xcf, 0xd2, 0x66, 0xff, 0x9a, 0x64, 0x41, 0x19, 0x1a, 0xc9, 0xd9, 0x5c, 0xe6, 0x24, 0x9b,
	0xa7, 0xfd, 0xb5, 0x7d, 0x7d, 0x3c, 0xaa, 0x7f, 0x9f, 0x5e, 0xab, 0x2e, 0x6e, 0xef, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x30, 0xee, 0xb5, 0x7b, 0xff, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnalyticsClient is the client API for Analytics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnalyticsClient interface {
	// Retrieve a daily, monthly or yearly record.
	GetAnalytics(ctx context.Context, in *io.AnalyticsRequest, opts ...grpc.CallOption) (*io.AnalyticsResponse, error)
}

type analyticsClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsClient(cc grpc.ClientConnInterface) AnalyticsClient {
	return &analyticsClient{cc}
}

func (c *analyticsClient) GetAnalytics(ctx context.Context, in *io.AnalyticsRequest, opts ...grpc.CallOption) (*io.AnalyticsResponse, error) {
	out := new(io.AnalyticsResponse)
	err := c.cc.Invoke(ctx, "/analytics.Analytics/getAnalytics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsServer is the server API for Analytics service.
type AnalyticsServer interface {
	// Retrieve a daily, monthly or yearly record.
	GetAnalytics(context.Context, *io.AnalyticsRequest) (*io.AnalyticsResponse, error)
}

// UnimplementedAnalyticsServer can be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServer struct {
}

func (*UnimplementedAnalyticsServer) GetAnalytics(ctx context.Context, req *io.AnalyticsRequest) (*io.AnalyticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnalytics not implemented")
}

func RegisterAnalyticsServer(s *grpc.Server, srv AnalyticsServer) {
	s.RegisterService(&_Analytics_serviceDesc, srv)
}

func _Analytics_GetAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(io.AnalyticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServer).GetAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.Analytics/GetAnalytics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServer).GetAnalytics(ctx, req.(*io.AnalyticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analytics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "analytics.Analytics",
	HandlerType: (*AnalyticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAnalytics",
			Handler:    _Analytics_GetAnalytics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "io/analytics/a_rpc.proto",
}
