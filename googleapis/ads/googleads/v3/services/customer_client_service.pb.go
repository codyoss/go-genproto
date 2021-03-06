// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/customer_client_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [CustomerClientService.GetCustomerClient][google.ads.googleads.v3.services.CustomerClientService.GetCustomerClient].
type GetCustomerClientRequest struct {
	// Required. The resource name of the client to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerClientRequest) Reset()         { *m = GetCustomerClientRequest{} }
func (m *GetCustomerClientRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerClientRequest) ProtoMessage()    {}
func (*GetCustomerClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_275643255b459061, []int{0}
}

func (m *GetCustomerClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerClientRequest.Unmarshal(m, b)
}
func (m *GetCustomerClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerClientRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerClientRequest.Merge(m, src)
}
func (m *GetCustomerClientRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerClientRequest.Size(m)
}
func (m *GetCustomerClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerClientRequest proto.InternalMessageInfo

func (m *GetCustomerClientRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerClientRequest)(nil), "google.ads.googleads.v3.services.GetCustomerClientRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/customer_client_service.proto", fileDescriptor_275643255b459061)
}

var fileDescriptor_275643255b459061 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x4a, 0xe3, 0x40,
	0x18, 0x27, 0x59, 0x58, 0xd8, 0xb0, 0x7b, 0xd8, 0xc0, 0xb2, 0xdd, 0xec, 0xc2, 0x96, 0x52, 0x50,
	0x7b, 0x98, 0xa1, 0xe6, 0x20, 0x8c, 0x28, 0x4c, 0x7b, 0xa8, 0x5e, 0xa4, 0xa8, 0xf4, 0x20, 0x81,
	0x30, 0x4d, 0xc6, 0x18, 0x48, 0x32, 0x31, 0x93, 0xe6, 0x22, 0x1e, 0x14, 0xdf, 0xc0, 0x37, 0xf0,
	0xe8, 0xa3, 0x14, 0x3c, 0x79, 0xf3, 0xe4, 0xc1, 0x93, 0x8f, 0xe0, 0x49, 0xd2, 0xc9, 0xa4, 0x8d,
	0x36, 0xf4, 0xf6, 0x23, 0xbf, 0x3f, 0xf3, 0x7d, 0xdf, 0x2f, 0xda, 0xae, 0xc7, 0x98, 0x17, 0x50,
	0x48, 0x5c, 0x0e, 0x05, 0xcc, 0x51, 0x66, 0x42, 0x4e, 0x93, 0xcc, 0x77, 0x28, 0x87, 0xce, 0x84,
	0xa7, 0x2c, 0xa4, 0x89, 0xed, 0x04, 0x3e, 0x8d, 0x52, 0xbb, 0x20, 0x40, 0x9c, 0xb0, 0x94, 0xe9,
	0x4d, 0x61, 0x02, 0xc4, 0xe5, 0xa0, 0xf4, 0x83, 0xcc, 0x04, 0xd2, 0x6f, 0x6c, 0xd5, 0xbd, 0x90,
	0x50, 0xce, 0x26, 0xc9, 0x92, 0x27, 0x44, 0xb4, 0xf1, 0x4f, 0x1a, 0x63, 0x1f, 0x92, 0x28, 0x62,
	0x29, 0x49, 0x7d, 0x16, 0xf1, 0x82, 0xfd, 0xbd, 0xc0, 0x56, 0x6c, 0xff, 0x17, 0x88, 0x53, 0x9f,
	0x06, 0xae, 0x3d, 0xa6, 0x67, 0x24, 0xf3, 0x59, 0x52, 0x08, 0xfe, 0x2c, 0x08, 0xe4, 0x0c, 0x82,
	0x6a, 0xc5, 0x5a, 0x63, 0x40, 0xd3, 0x7e, 0x31, 0x4e, 0x7f, 0x16, 0x7b, 0x48, 0xcf, 0x27, 0x94,
	0xa7, 0xfa, 0xb1, 0xf6, 0x43, 0xaa, 0xed, 0x88, 0x84, 0xb4, 0xa1, 0x34, 0x95, 0xf5, 0x6f, 0x3d,
	0xf8, 0x8c, 0xd5, 0x37, 0xbc, 0xa1, 0xad, 0xcd, 0xb7, 0x2f, 0x50, 0xec, 0x73, 0xe0, 0xb0, 0x10,
	0x7e, 0x88, 0xfb, 0x2e, 0x53, 0x0e, 0x48, 0x48, 0x37, 0x6f, 0x54, 0xed, 0x57, 0x55, 0x70, 0x24,
	0x0e, 0xa7, 0x3f, 0x28, 0xda, 0xcf, 0x4f, 0xc3, 0xe8, 0x08, 0xac, 0x3a, 0x38, 0xa8, 0xdb, 0xc0,
	0xe8, 0xd6, 0x7a, 0xcb, 0x2a, 0x40, 0xd5, 0xd9, 0xda, 0x7f, 0xc2, 0xd5, 0xad, 0xaf, 0x1f, 0x5f,
	0x6e, 0x55, 0x53, 0xef, 0xe6, 0x05, 0x5e, 0x54, 0x98, 0x1d, 0xd9, 0x22, 0x87, 0x9d, 0xb2, 0x51,
	0x11, 0xc3, 0x61, 0xe7, 0xd2, 0xf8, 0x3b, 0xc5, 0x8d, 0xba, 0x1b, 0xf5, 0xae, 0x54, 0xad, 0xed,
	0xb0, 0x70, 0xe5, 0x72, 0x3d, 0x63, 0xe9, 0xb1, 0x86, 0x79, 0x7b, 0x43, 0xe5, 0x64, 0xaf, 0xf0,
	0x7b, 0x2c, 0x20, 0x91, 0x07, 0x58, 0xe2, 0x41, 0x8f, 0x46, 0xb3, 0x6e, 0xe1, 0xfc, 0xc5, 0xfa,
	0x9f, 0x7d, 0x5b, 0x82, 0x3b, 0xf5, 0xcb, 0x00, 0xe3, 0x7b, 0xb5, 0x39, 0x10, 0x81, 0xd8, 0xe5,
	0x40, 0xc0, 0x1c, 0x8d, 0x4c, 0x50, 0x3c, 0xcc, 0xa7, 0x52, 0x62, 0x61, 0x97, 0x5b, 0xa5, 0xc4,
	0x1a, 0x99, 0x96, 0x94, 0xbc, 0xaa, 0x6d, 0xf1, 0x1d, 0x21, 0xec, 0x72, 0x84, 0x4a, 0x11, 0x42,
	0x23, 0x13, 0x21, 0x29, 0x1b, 0x7f, 0x9d, 0xcd, 0x69, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x5a,
	0xc3, 0x5b, 0x79, 0x93, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerClientServiceClient is the client API for CustomerClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClientServiceClient interface {
	// Returns the requested client in full detail.
	GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error)
}

type customerClientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClientServiceClient(cc grpc.ClientConnInterface) CustomerClientServiceClient {
	return &customerClientServiceClient{cc}
}

func (c *customerClientServiceClient) GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error) {
	out := new(resources.CustomerClient)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerClientService/GetCustomerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientServiceServer is the server API for CustomerClientService service.
type CustomerClientServiceServer interface {
	// Returns the requested client in full detail.
	GetCustomerClient(context.Context, *GetCustomerClientRequest) (*resources.CustomerClient, error)
}

// UnimplementedCustomerClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerClientServiceServer struct {
}

func (*UnimplementedCustomerClientServiceServer) GetCustomerClient(ctx context.Context, req *GetCustomerClientRequest) (*resources.CustomerClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerClient not implemented")
}

func RegisterCustomerClientServiceServer(s *grpc.Server, srv CustomerClientServiceServer) {
	s.RegisterService(&_CustomerClientService_serviceDesc, srv)
}

func _CustomerClientService_GetCustomerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerClientService/GetCustomerClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, req.(*GetCustomerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CustomerClientService",
	HandlerType: (*CustomerClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClient",
			Handler:    _CustomerClientService_GetCustomerClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/customer_client_service.proto",
}
