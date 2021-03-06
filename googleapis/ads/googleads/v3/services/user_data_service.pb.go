// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/user_data_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
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

// Request message for [UserDataService.UploadUserData][google.ads.googleads.v3.services.UserDataService.UploadUserData]
type UploadUserDataRequest struct {
	// Required. The ID of the customer for which to update the user data.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to be done.
	Operations []*UserDataOperation `protobuf:"bytes,3,rep,name=operations,proto3" json:"operations,omitempty"`
	// Metadata of the request.
	//
	// Types that are valid to be assigned to Metadata:
	//	*UploadUserDataRequest_CustomerMatchUserListMetadata
	Metadata             isUploadUserDataRequest_Metadata `protobuf_oneof:"metadata"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *UploadUserDataRequest) Reset()         { *m = UploadUserDataRequest{} }
func (m *UploadUserDataRequest) String() string { return proto.CompactTextString(m) }
func (*UploadUserDataRequest) ProtoMessage()    {}
func (*UploadUserDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f394d2fdab45a7d, []int{0}
}

func (m *UploadUserDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadUserDataRequest.Unmarshal(m, b)
}
func (m *UploadUserDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadUserDataRequest.Marshal(b, m, deterministic)
}
func (m *UploadUserDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadUserDataRequest.Merge(m, src)
}
func (m *UploadUserDataRequest) XXX_Size() int {
	return xxx_messageInfo_UploadUserDataRequest.Size(m)
}
func (m *UploadUserDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadUserDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadUserDataRequest proto.InternalMessageInfo

func (m *UploadUserDataRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *UploadUserDataRequest) GetOperations() []*UserDataOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

type isUploadUserDataRequest_Metadata interface {
	isUploadUserDataRequest_Metadata()
}

type UploadUserDataRequest_CustomerMatchUserListMetadata struct {
	CustomerMatchUserListMetadata *common.CustomerMatchUserListMetadata `protobuf:"bytes,2,opt,name=customer_match_user_list_metadata,json=customerMatchUserListMetadata,proto3,oneof"`
}

func (*UploadUserDataRequest_CustomerMatchUserListMetadata) isUploadUserDataRequest_Metadata() {}

func (m *UploadUserDataRequest) GetMetadata() isUploadUserDataRequest_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *UploadUserDataRequest) GetCustomerMatchUserListMetadata() *common.CustomerMatchUserListMetadata {
	if x, ok := m.GetMetadata().(*UploadUserDataRequest_CustomerMatchUserListMetadata); ok {
		return x.CustomerMatchUserListMetadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadUserDataRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadUserDataRequest_CustomerMatchUserListMetadata)(nil),
	}
}

// Operation to be made for the UploadUserDataRequest.
type UserDataOperation struct {
	// Operation to be made for the UploadUserDataRequest.
	//
	// Types that are valid to be assigned to Operation:
	//	*UserDataOperation_Create
	Operation            isUserDataOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *UserDataOperation) Reset()         { *m = UserDataOperation{} }
func (m *UserDataOperation) String() string { return proto.CompactTextString(m) }
func (*UserDataOperation) ProtoMessage()    {}
func (*UserDataOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f394d2fdab45a7d, []int{1}
}

func (m *UserDataOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDataOperation.Unmarshal(m, b)
}
func (m *UserDataOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDataOperation.Marshal(b, m, deterministic)
}
func (m *UserDataOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDataOperation.Merge(m, src)
}
func (m *UserDataOperation) XXX_Size() int {
	return xxx_messageInfo_UserDataOperation.Size(m)
}
func (m *UserDataOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDataOperation.DiscardUnknown(m)
}

var xxx_messageInfo_UserDataOperation proto.InternalMessageInfo

type isUserDataOperation_Operation interface {
	isUserDataOperation_Operation()
}

type UserDataOperation_Create struct {
	Create *common.UserData `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

func (*UserDataOperation_Create) isUserDataOperation_Operation() {}

func (m *UserDataOperation) GetOperation() isUserDataOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *UserDataOperation) GetCreate() *common.UserData {
	if x, ok := m.GetOperation().(*UserDataOperation_Create); ok {
		return x.Create
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UserDataOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UserDataOperation_Create)(nil),
	}
}

// Response message for [UserDataService.UploadUserData][google.ads.googleads.v3.services.UserDataService.UploadUserData]
type UploadUserDataResponse struct {
	// The date time at which the request was received by API, formatted as
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", e.g. "2019-01-01 12:32:45-08:00".
	UploadDateTime *wrappers.StringValue `protobuf:"bytes,1,opt,name=upload_date_time,json=uploadDateTime,proto3" json:"upload_date_time,omitempty"`
	// Number of upload data operations received by API.
	ReceivedOperationsCount *wrappers.Int32Value `protobuf:"bytes,2,opt,name=received_operations_count,json=receivedOperationsCount,proto3" json:"received_operations_count,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *UploadUserDataResponse) Reset()         { *m = UploadUserDataResponse{} }
func (m *UploadUserDataResponse) String() string { return proto.CompactTextString(m) }
func (*UploadUserDataResponse) ProtoMessage()    {}
func (*UploadUserDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f394d2fdab45a7d, []int{2}
}

func (m *UploadUserDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadUserDataResponse.Unmarshal(m, b)
}
func (m *UploadUserDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadUserDataResponse.Marshal(b, m, deterministic)
}
func (m *UploadUserDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadUserDataResponse.Merge(m, src)
}
func (m *UploadUserDataResponse) XXX_Size() int {
	return xxx_messageInfo_UploadUserDataResponse.Size(m)
}
func (m *UploadUserDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadUserDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadUserDataResponse proto.InternalMessageInfo

func (m *UploadUserDataResponse) GetUploadDateTime() *wrappers.StringValue {
	if m != nil {
		return m.UploadDateTime
	}
	return nil
}

func (m *UploadUserDataResponse) GetReceivedOperationsCount() *wrappers.Int32Value {
	if m != nil {
		return m.ReceivedOperationsCount
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadUserDataRequest)(nil), "google.ads.googleads.v3.services.UploadUserDataRequest")
	proto.RegisterType((*UserDataOperation)(nil), "google.ads.googleads.v3.services.UserDataOperation")
	proto.RegisterType((*UploadUserDataResponse)(nil), "google.ads.googleads.v3.services.UploadUserDataResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/user_data_service.proto", fileDescriptor_3f394d2fdab45a7d)
}

var fileDescriptor_3f394d2fdab45a7d = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xae, 0x1d, 0xa9, 0xba, 0x9d, 0x48, 0xbd, 0xf7, 0x5a, 0x40, 0x43, 0x5a, 0x20, 0x44, 0x5d,
	0x44, 0x15, 0xb2, 0x45, 0x2c, 0xd1, 0xca, 0xa8, 0x0b, 0xa7, 0x15, 0x6d, 0x25, 0x2a, 0xaa, 0x94,
	0x06, 0x84, 0x22, 0x59, 0x53, 0xfb, 0xc4, 0x1d, 0xc9, 0xf6, 0x98, 0x99, 0x71, 0x58, 0x20, 0x36,
	0xec, 0x58, 0xf3, 0x06, 0x2c, 0xd9, 0xb0, 0x46, 0xe2, 0x09, 0xba, 0xe5, 0x05, 0x58, 0xb0, 0x42,
	0xe2, 0x09, 0xd8, 0x20, 0x67, 0x66, 0x9c, 0xfe, 0x50, 0x0a, 0xbb, 0x93, 0x99, 0xef, 0xe7, 0xcc,
	0x77, 0x4e, 0x8c, 0xd6, 0x62, 0x4a, 0xe3, 0x04, 0x1c, 0x1c, 0x71, 0x47, 0x96, 0x65, 0x35, 0x76,
	0x1d, 0x0e, 0x6c, 0x4c, 0x42, 0xe0, 0x4e, 0xc1, 0x81, 0x05, 0x11, 0x16, 0x38, 0x50, 0x47, 0x76,
	0xce, 0xa8, 0xa0, 0x56, 0x4b, 0xc2, 0x6d, 0x1c, 0x71, 0xbb, 0x62, 0xda, 0x63, 0xd7, 0xd6, 0xcc,
	0xe6, 0xbd, 0x8b, 0xb4, 0x43, 0x9a, 0xa6, 0x34, 0x73, 0xe8, 0x68, 0x94, 0x90, 0x0c, 0x82, 0xca,
	0x41, 0x2a, 0x37, 0x97, 0x34, 0x2f, 0x27, 0x0e, 0xce, 0x32, 0x2a, 0xb0, 0x20, 0x34, 0xe3, 0xea,
	0xf6, 0xd6, 0x89, 0xdb, 0x11, 0x81, 0x24, 0x0a, 0x0e, 0xe1, 0x08, 0x8f, 0x09, 0x65, 0x0a, 0x70,
	0x53, 0x01, 0x26, 0xbf, 0x0e, 0x8b, 0x91, 0xf3, 0x82, 0xe1, 0x3c, 0x07, 0xa6, 0x05, 0x16, 0x4e,
	0x08, 0x84, 0x09, 0x81, 0x4c, 0xc8, 0x8b, 0xf6, 0x07, 0x13, 0x5d, 0x3d, 0xc8, 0x13, 0x8a, 0xa3,
	0x03, 0x0e, 0x6c, 0x13, 0x0b, 0xdc, 0x87, 0xe7, 0x05, 0x70, 0x61, 0x2d, 0xa3, 0x7a, 0x58, 0x70,
	0x41, 0x53, 0x60, 0x01, 0x89, 0x1a, 0x46, 0xcb, 0xe8, 0xcc, 0xf5, 0x6a, 0x5f, 0x7c, 0xb3, 0x8f,
	0xf4, 0xf9, 0x4e, 0x64, 0x3d, 0x45, 0x88, 0xe6, 0xc0, 0x64, 0xb7, 0x8d, 0x5a, 0xab, 0xd6, 0xa9,
	0x77, 0x5d, 0xfb, 0xb2, 0x98, 0x6c, 0x6d, 0xf6, 0x48, 0x73, 0x95, 0xf2, 0x54, 0xcb, 0x7a, 0x63,
	0xa0, 0xdb, 0x55, 0x03, 0x29, 0x16, 0xe1, 0x91, 0x0c, 0x2d, 0x21, 0x5c, 0x04, 0x29, 0x08, 0x5c,
	0xa6, 0xd7, 0x30, 0x5b, 0x46, 0xa7, 0xde, 0x5d, 0xbf, 0xd0, 0x51, 0xc6, 0x6e, 0x6f, 0x28, 0xa1,
	0xdd, 0x52, 0xa7, 0x34, 0x7f, 0x48, 0xb8, 0xd8, 0x55, 0x22, 0xdb, 0x33, 0xfd, 0x1b, 0xe1, 0xef,
	0x00, 0x3d, 0x84, 0xfe, 0xd1, 0x8e, 0xed, 0x08, 0xfd, 0x7f, 0xae, 0x7b, 0xab, 0x87, 0x66, 0x43,
	0x06, 0x58, 0xc0, 0x24, 0xa7, 0x7a, 0xb7, 0x73, 0x59, 0x43, 0x5a, 0x62, 0x7b, 0xa6, 0xaf, 0x98,
	0xbd, 0x3a, 0x9a, 0xab, 0x9e, 0xdf, 0xfe, 0x68, 0xa0, 0x6b, 0x67, 0xe7, 0xc2, 0x73, 0x9a, 0x71,
	0xb0, 0x1e, 0xa0, 0xff, 0x8a, 0xc9, 0x4d, 0xb9, 0x3f, 0x10, 0x08, 0x92, 0x6a, 0xd7, 0x25, 0xed,
	0xaa, 0xd7, 0xc0, 0xde, 0x17, 0x8c, 0x64, 0xf1, 0x00, 0x27, 0x05, 0xf4, 0xe7, 0x25, 0x6b, 0x13,
	0x0b, 0x78, 0x4c, 0x52, 0xb0, 0x9e, 0xa0, 0xeb, 0x0c, 0x42, 0x20, 0x63, 0x88, 0x82, 0x69, 0xee,
	0x41, 0x48, 0x8b, 0x4c, 0xa8, 0x5c, 0x17, 0xcf, 0x09, 0xee, 0x64, 0xc2, 0xed, 0x4a, 0xbd, 0x05,
	0xcd, 0xae, 0x62, 0xe0, 0x1b, 0x25, 0xb7, 0xfb, 0xdd, 0x40, 0xff, 0xea, 0xae, 0xf7, 0xe5, 0xc4,
	0xad, 0x4f, 0x06, 0x9a, 0x3f, 0xfd, 0x1e, 0x6b, 0xf5, 0x0f, 0xd6, 0xe4, 0x57, 0x9b, 0xd9, 0x5c,
	0xfb, 0x7b, 0xa2, 0x8c, 0xae, 0xbd, 0xfa, 0xfa, 0xf3, 0xd7, 0xb7, 0xe6, 0xdd, 0xf6, 0x9d, 0xc9,
	0x3f, 0x52, 0x8d, 0x9c, 0x3b, 0x2f, 0x4f, 0x2c, 0xfa, 0xfa, 0xca, 0x2b, 0xaf, 0x38, 0xc5, 0xf6,
	0x8c, 0x95, 0xe6, 0xe2, 0xb1, 0xdf, 0x98, 0x3a, 0xa9, 0x2a, 0x27, 0xbc, 0x1c, 0x67, 0xef, 0x87,
	0x81, 0x96, 0x43, 0x9a, 0x5e, 0xda, 0x55, 0xef, 0xca, 0x99, 0x54, 0xf6, 0xca, 0x54, 0xf7, 0x8c,
	0x67, 0xdb, 0x8a, 0x19, 0xd3, 0x04, 0x67, 0xb1, 0x4d, 0x59, 0xec, 0xc4, 0x90, 0x4d, 0x32, 0x77,
	0xa6, 0x5e, 0x17, 0x7f, 0xaf, 0xee, 0xeb, 0xe2, 0x9d, 0x59, 0xdb, 0xf2, 0xfd, 0xf7, 0x66, 0x6b,
	0x4b, 0x0a, 0xfa, 0x11, 0xb7, 0x65, 0x59, 0x56, 0x03, 0xd7, 0x56, 0xc6, 0xfc, 0x58, 0x43, 0x86,
	0x7e, 0xc4, 0x87, 0x15, 0x64, 0x38, 0x70, 0x87, 0x1a, 0xf2, 0xcd, 0x5c, 0x96, 0xe7, 0x9e, 0xe7,
	0x47, 0xdc, 0xf3, 0x2a, 0x90, 0xe7, 0x0d, 0x5c, 0xcf, 0xd3, 0xb0, 0xc3, 0xd9, 0x49, 0x9f, 0xee,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xe1, 0x1f, 0x70, 0x56, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserDataServiceClient is the client API for UserDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserDataServiceClient interface {
	// Uploads the given user data.
	UploadUserData(ctx context.Context, in *UploadUserDataRequest, opts ...grpc.CallOption) (*UploadUserDataResponse, error)
}

type userDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataServiceClient(cc grpc.ClientConnInterface) UserDataServiceClient {
	return &userDataServiceClient{cc}
}

func (c *userDataServiceClient) UploadUserData(ctx context.Context, in *UploadUserDataRequest, opts ...grpc.CallOption) (*UploadUserDataResponse, error) {
	out := new(UploadUserDataResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.UserDataService/UploadUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataServiceServer is the server API for UserDataService service.
type UserDataServiceServer interface {
	// Uploads the given user data.
	UploadUserData(context.Context, *UploadUserDataRequest) (*UploadUserDataResponse, error)
}

// UnimplementedUserDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserDataServiceServer struct {
}

func (*UnimplementedUserDataServiceServer) UploadUserData(ctx context.Context, req *UploadUserDataRequest) (*UploadUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadUserData not implemented")
}

func RegisterUserDataServiceServer(s *grpc.Server, srv UserDataServiceServer) {
	s.RegisterService(&_UserDataService_serviceDesc, srv)
}

func _UserDataService_UploadUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).UploadUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.UserDataService/UploadUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).UploadUserData(ctx, req.(*UploadUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.UserDataService",
	HandlerType: (*UserDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadUserData",
			Handler:    _UserDataService_UploadUserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/user_data_service.proto",
}
