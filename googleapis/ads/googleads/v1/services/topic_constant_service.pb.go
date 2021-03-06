// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/topic_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for [TopicConstantService.GetTopicConstant][google.ads.googleads.v1.services.TopicConstantService.GetTopicConstant].
type GetTopicConstantRequest struct {
	// Required. Resource name of the Topic to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicConstantRequest) Reset()         { *m = GetTopicConstantRequest{} }
func (m *GetTopicConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicConstantRequest) ProtoMessage()    {}
func (*GetTopicConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49736dbc2500820f, []int{0}
}

func (m *GetTopicConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicConstantRequest.Unmarshal(m, b)
}
func (m *GetTopicConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetTopicConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicConstantRequest.Merge(m, src)
}
func (m *GetTopicConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicConstantRequest.Size(m)
}
func (m *GetTopicConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicConstantRequest proto.InternalMessageInfo

func (m *GetTopicConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTopicConstantRequest)(nil), "google.ads.googleads.v1.services.GetTopicConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/topic_constant_service.proto", fileDescriptor_49736dbc2500820f)
}

var fileDescriptor_49736dbc2500820f = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0xcb, 0xd3, 0x40,
	0x18, 0x26, 0x11, 0x04, 0x83, 0x82, 0x04, 0xa1, 0x6d, 0x14, 0x2c, 0xa5, 0x94, 0xe2, 0x70, 0x67,
	0x14, 0x04, 0x4f, 0x3b, 0x5c, 0x1d, 0xea, 0x24, 0xc5, 0x4a, 0x07, 0x09, 0x84, 0x6b, 0x72, 0xc6,
	0x83, 0xe4, 0x2e, 0xe6, 0xae, 0x59, 0x44, 0x10, 0xff, 0x82, 0xff, 0xc0, 0xd1, 0xff, 0xe0, 0x1f,
	0xe8, 0xea, 0xe6, 0xe4, 0xe0, 0xe4, 0xea, 0xe6, 0x24, 0xe9, 0xe5, 0xd2, 0xc4, 0xef, 0x0b, 0xdd,
	0x1e, 0xee, 0x79, 0x9e, 0xf7, 0x7d, 0xde, 0xf7, 0x3d, 0x67, 0x91, 0x08, 0x91, 0xa4, 0x14, 0x92,
	0x58, 0x42, 0x0d, 0x2b, 0x54, 0xfa, 0x50, 0xd2, 0xa2, 0x64, 0x11, 0x95, 0x50, 0x89, 0x9c, 0x45,
	0x61, 0x24, 0xb8, 0x54, 0x84, 0xab, 0xb0, 0x7e, 0x07, 0x79, 0x21, 0x94, 0x70, 0xc7, 0xda, 0x03,
	0x48, 0x2c, 0x41, 0x63, 0x07, 0xa5, 0x0f, 0x8c, 0xdd, 0x7b, 0xd4, 0xd7, 0xa0, 0xa0, 0x52, 0xec,
	0x8b, 0x8b, 0x1d, 0x74, 0x65, 0xef, 0x8e, 0xf1, 0xe5, 0x0c, 0x12, 0xce, 0x85, 0x22, 0x8a, 0x09,
	0x2e, 0x6b, 0x76, 0xd0, 0x62, 0xa3, 0x94, 0xd1, 0xc6, 0x76, 0xb7, 0x45, 0xbc, 0x61, 0x34, 0x8d,
	0xc3, 0x1d, 0x7d, 0x4b, 0x4a, 0x26, 0x8a, 0x5a, 0x30, 0x6a, 0x09, 0x4c, 0x04, 0x4d, 0x4d, 0xb8,
	0x33, 0x58, 0x51, 0xf5, 0xaa, 0x4a, 0xf3, 0xac, 0x0e, 0xf3, 0x92, 0xbe, 0xdb, 0x53, 0xa9, 0xdc,
	0x8d, 0x73, 0xc3, 0x88, 0x43, 0x4e, 0x32, 0x3a, 0xb4, 0xc6, 0xd6, 0xfc, 0xda, 0x12, 0xfc, 0xc4,
	0xf6, 0x5f, 0x3c, 0x77, 0x66, 0xa7, 0xd9, 0x6b, 0x94, 0x33, 0x09, 0x22, 0x91, 0xc1, 0x6e, 0xb5,
	0xeb, 0xa6, 0xc8, 0x0b, 0x92, 0xd1, 0x07, 0x7f, 0x2c, 0xe7, 0x56, 0x87, 0xdf, 0xe8, 0xa5, 0xb9,
	0xdf, 0x2c, 0xe7, 0xe6, 0xff, 0x49, 0xdc, 0xc7, 0xe0, 0xdc, 0xae, 0x41, 0x4f, 0x7a, 0xef, 0x7e,
	0xaf, 0xb5, 0x39, 0x02, 0xe8, 0x18, 0x27, 0x4f, 0x7f, 0xe0, 0xee, 0xc0, 0x9f, 0xbe, 0xff, 0xfa,
	0x6c, 0xcf, 0xdc, 0x69, 0x75, 0xb9, 0xf7, 0x1d, 0x66, 0xa1, 0xda, 0x4e, 0x09, 0xef, 0x7d, 0xf0,
	0x6e, 0x1f, 0xf0, 0xb0, 0x6f, 0x23, 0xcb, 0x8f, 0xb6, 0x33, 0x8d, 0x44, 0x76, 0x76, 0x9a, 0xe5,
	0xe8, 0xb2, 0xdd, 0xac, 0xab, 0x4b, 0xad, 0xad, 0xd7, 0xcf, 0x6b, 0x7b, 0x22, 0x52, 0xc2, 0x13,
	0x20, 0x8a, 0x04, 0x26, 0x94, 0x1f, 0xef, 0x08, 0x4f, 0x0d, 0xfb, 0xbf, 0xf5, 0x13, 0x03, 0xbe,
	0xd8, 0x57, 0x56, 0x18, 0x7f, 0xb5, 0xc7, 0x2b, 0x5d, 0x10, 0xc7, 0x12, 0x68, 0x58, 0xa1, 0xad,
	0x0f, 0xea, 0xc6, 0xf2, 0x60, 0x24, 0x01, 0x8e, 0x65, 0xd0, 0x48, 0x82, 0xad, 0x1f, 0x18, 0xc9,
	0x6f, 0x7b, 0xaa, 0xdf, 0x11, 0xc2, 0xb1, 0x44, 0xa8, 0x11, 0x21, 0xb4, 0xf5, 0x11, 0x32, 0xb2,
	0xdd, 0xd5, 0x63, 0xce, 0x87, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x8a, 0xdd, 0xdc, 0x7d,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TopicConstantServiceClient is the client API for TopicConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicConstantServiceClient interface {
	// Returns the requested topic constant in full detail.
	GetTopicConstant(ctx context.Context, in *GetTopicConstantRequest, opts ...grpc.CallOption) (*resources.TopicConstant, error)
}

type topicConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicConstantServiceClient(cc grpc.ClientConnInterface) TopicConstantServiceClient {
	return &topicConstantServiceClient{cc}
}

func (c *topicConstantServiceClient) GetTopicConstant(ctx context.Context, in *GetTopicConstantRequest, opts ...grpc.CallOption) (*resources.TopicConstant, error) {
	out := new(resources.TopicConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.TopicConstantService/GetTopicConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicConstantServiceServer is the server API for TopicConstantService service.
type TopicConstantServiceServer interface {
	// Returns the requested topic constant in full detail.
	GetTopicConstant(context.Context, *GetTopicConstantRequest) (*resources.TopicConstant, error)
}

// UnimplementedTopicConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopicConstantServiceServer struct {
}

func (*UnimplementedTopicConstantServiceServer) GetTopicConstant(ctx context.Context, req *GetTopicConstantRequest) (*resources.TopicConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicConstant not implemented")
}

func RegisterTopicConstantServiceServer(s *grpc.Server, srv TopicConstantServiceServer) {
	s.RegisterService(&_TopicConstantService_serviceDesc, srv)
}

func _TopicConstantService_GetTopicConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicConstantServiceServer).GetTopicConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.TopicConstantService/GetTopicConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicConstantServiceServer).GetTopicConstant(ctx, req.(*GetTopicConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.TopicConstantService",
	HandlerType: (*TopicConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopicConstant",
			Handler:    _TopicConstantService_GetTopicConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/topic_constant_service.proto",
}
