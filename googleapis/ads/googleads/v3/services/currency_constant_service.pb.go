// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/currency_constant_service.proto

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

// Request message for [CurrencyConstantService.GetCurrencyConstant][google.ads.googleads.v3.services.CurrencyConstantService.GetCurrencyConstant].
type GetCurrencyConstantRequest struct {
	// Required. Resource name of the currency constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrencyConstantRequest) Reset()         { *m = GetCurrencyConstantRequest{} }
func (m *GetCurrencyConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrencyConstantRequest) ProtoMessage()    {}
func (*GetCurrencyConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_823f63dbbd8d219a, []int{0}
}

func (m *GetCurrencyConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrencyConstantRequest.Unmarshal(m, b)
}
func (m *GetCurrencyConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrencyConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrencyConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrencyConstantRequest.Merge(m, src)
}
func (m *GetCurrencyConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrencyConstantRequest.Size(m)
}
func (m *GetCurrencyConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrencyConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrencyConstantRequest proto.InternalMessageInfo

func (m *GetCurrencyConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCurrencyConstantRequest)(nil), "google.ads.googleads.v3.services.GetCurrencyConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/currency_constant_service.proto", fileDescriptor_823f63dbbd8d219a)
}

var fileDescriptor_823f63dbbd8d219a = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0xaa, 0xd4, 0x30,
	0x14, 0xa6, 0x15, 0x04, 0x8b, 0x6e, 0xea, 0xe2, 0x5e, 0xeb, 0x05, 0x87, 0xcb, 0x05, 0x1d, 0x85,
	0x04, 0xed, 0xca, 0xa8, 0x68, 0x66, 0x16, 0xe3, 0x4a, 0x06, 0x85, 0x2e, 0xa4, 0x50, 0x32, 0x69,
	0xac, 0x85, 0x36, 0x19, 0x93, 0x4c, 0x41, 0xc4, 0xcd, 0xf8, 0x08, 0xbe, 0x81, 0x4b, 0xdf, 0xc4,
	0xd9, 0xba, 0x73, 0xe5, 0xc2, 0x95, 0x8f, 0xe0, 0x4a, 0x3a, 0x49, 0x3a, 0x3f, 0x77, 0xca, 0xec,
	0x3e, 0xfa, 0xfd, 0x9c, 0x93, 0xef, 0x34, 0x78, 0x51, 0x08, 0x51, 0x54, 0x0c, 0x92, 0x5c, 0x41,
	0x03, 0x5b, 0xd4, 0xc4, 0x50, 0x31, 0xd9, 0x94, 0x94, 0x29, 0x48, 0x17, 0x52, 0x32, 0x4e, 0x3f,
	0x66, 0x54, 0x70, 0xa5, 0x09, 0xd7, 0x99, 0xa5, 0xc0, 0x5c, 0x0a, 0x2d, 0xc2, 0x81, 0xb1, 0x01,
	0x92, 0x2b, 0xd0, 0x25, 0x80, 0x26, 0x06, 0x2e, 0x21, 0x7a, 0xdc, 0x37, 0x43, 0x32, 0x25, 0x16,
	0xf2, 0xe0, 0x10, 0x13, 0x1e, 0x9d, 0x39, 0xeb, 0xbc, 0x84, 0x84, 0x73, 0xa1, 0x89, 0x2e, 0x05,
	0x57, 0x96, 0x3d, 0xd9, 0x62, 0x69, 0x55, 0xb2, 0xce, 0x76, 0x67, 0x8b, 0x78, 0x57, 0xb2, 0x2a,
	0xcf, 0x66, 0xec, 0x3d, 0x69, 0x4a, 0x21, 0xad, 0xe0, 0xd6, 0x96, 0xc0, 0x6d, 0x61, 0xa8, 0x73,
	0x1d, 0x44, 0x13, 0xa6, 0xc7, 0x76, 0xa1, 0xb1, 0xdd, 0xe7, 0x35, 0xfb, 0xb0, 0x60, 0x4a, 0x87,
	0x49, 0x70, 0xc3, 0xe9, 0x33, 0x4e, 0x6a, 0x76, 0xea, 0x0d, 0xbc, 0x7b, 0xd7, 0x46, 0x0f, 0x7f,
	0x63, 0xff, 0x1f, 0x7e, 0x10, 0x0c, 0x37, 0x0d, 0x58, 0x34, 0x2f, 0x15, 0xa0, 0xa2, 0x86, 0x97,
	0x02, 0xaf, 0xbb, 0x9c, 0x57, 0xa4, 0x66, 0x8f, 0x96, 0x7e, 0x70, 0xb2, 0x2f, 0x79, 0x63, 0x0a,
	0x0c, 0x7f, 0x78, 0xc1, 0xcd, 0x03, 0x2b, 0x85, 0x4f, 0xc1, 0xb1, 0xea, 0x41, 0xff, 0x4b, 0xa2,
	0xb8, 0xd7, 0xdd, 0x9d, 0x05, 0xec, 0x7b, 0xcf, 0x9f, 0xff, 0xc2, 0xbb, 0xef, 0x5f, 0xfe, 0xfc,
	0xf3, 0xd5, 0x1f, 0x86, 0x77, 0xdb, 0x73, 0x7e, 0xda, 0x61, 0x9e, 0xd1, 0x3d, 0xb3, 0x82, 0xf7,
	0x3f, 0x47, 0xb7, 0x57, 0xf8, 0xb4, 0xaf, 0xa3, 0xd1, 0x17, 0x3f, 0xb8, 0xa0, 0xa2, 0x3e, 0xfa,
	0xac, 0xd1, 0x59, 0x4f, 0x55, 0xd3, 0xf6, 0x82, 0x53, 0xef, 0xed, 0x4b, 0x9b, 0x50, 0x88, 0x8a,
	0xf0, 0x02, 0x08, 0x59, 0xc0, 0x82, 0xf1, 0xf5, 0x7d, 0xe1, 0x66, 0x66, 0xff, 0x4f, 0xff, 0xc4,
	0x81, 0x6f, 0xfe, 0x95, 0x09, 0xc6, 0xdf, 0xfd, 0xc1, 0xc4, 0x04, 0xe2, 0x5c, 0x01, 0x03, 0x5b,
	0x94, 0xc4, 0xc0, 0x0e, 0x56, 0x2b, 0x27, 0x49, 0x71, 0xae, 0xd2, 0x4e, 0x92, 0x26, 0x71, 0xea,
	0x24, 0x7f, 0xfd, 0x0b, 0xf3, 0x1d, 0x21, 0x9c, 0x2b, 0x84, 0x3a, 0x11, 0x42, 0x49, 0x8c, 0x90,
	0x93, 0xcd, 0xae, 0xae, 0xf7, 0x8c, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xc6, 0xca, 0xc8,
	0x9b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyConstantServiceClient is the client API for CurrencyConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyConstantServiceClient interface {
	// Returns the requested currency constant.
	GetCurrencyConstant(ctx context.Context, in *GetCurrencyConstantRequest, opts ...grpc.CallOption) (*resources.CurrencyConstant, error)
}

type currencyConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyConstantServiceClient(cc grpc.ClientConnInterface) CurrencyConstantServiceClient {
	return &currencyConstantServiceClient{cc}
}

func (c *currencyConstantServiceClient) GetCurrencyConstant(ctx context.Context, in *GetCurrencyConstantRequest, opts ...grpc.CallOption) (*resources.CurrencyConstant, error) {
	out := new(resources.CurrencyConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CurrencyConstantService/GetCurrencyConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyConstantServiceServer is the server API for CurrencyConstantService service.
type CurrencyConstantServiceServer interface {
	// Returns the requested currency constant.
	GetCurrencyConstant(context.Context, *GetCurrencyConstantRequest) (*resources.CurrencyConstant, error)
}

// UnimplementedCurrencyConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyConstantServiceServer struct {
}

func (*UnimplementedCurrencyConstantServiceServer) GetCurrencyConstant(ctx context.Context, req *GetCurrencyConstantRequest) (*resources.CurrencyConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencyConstant not implemented")
}

func RegisterCurrencyConstantServiceServer(s *grpc.Server, srv CurrencyConstantServiceServer) {
	s.RegisterService(&_CurrencyConstantService_serviceDesc, srv)
}

func _CurrencyConstantService_GetCurrencyConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrencyConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyConstantServiceServer).GetCurrencyConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CurrencyConstantService/GetCurrencyConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyConstantServiceServer).GetCurrencyConstant(ctx, req.(*GetCurrencyConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CurrencyConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CurrencyConstantService",
	HandlerType: (*CurrencyConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrencyConstant",
			Handler:    _CurrencyConstantService_GetCurrencyConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/currency_constant_service.proto",
}
