// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/user_interest.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A user interest: a particular interest-based vertical to be targeted.
type UserInterest struct {
	// Output only. The resource name of the user interest.
	// User interest resource names have the form:
	//
	// `customers/{customer_id}/userInterests/{user_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Taxonomy type of the user interest.
	TaxonomyType enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType `protobuf:"varint,2,opt,name=taxonomy_type,json=taxonomyType,proto3,enum=google.ads.googleads.v3.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType" json:"taxonomy_type,omitempty"`
	// Output only. The ID of the user interest.
	UserInterestId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=user_interest_id,json=userInterestId,proto3" json:"user_interest_id,omitempty"`
	// Output only. The name of the user interest.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The parent of the user interest.
	UserInterestParent *wrappers.StringValue `protobuf:"bytes,5,opt,name=user_interest_parent,json=userInterestParent,proto3" json:"user_interest_parent,omitempty"`
	// Output only. True if the user interest is launched to all channels and locales.
	LaunchedToAll *wrappers.BoolValue `protobuf:"bytes,6,opt,name=launched_to_all,json=launchedToAll,proto3" json:"launched_to_all,omitempty"`
	// Output only. Availability information of the user interest.
	Availabilities       []*common.CriterionCategoryAvailability `protobuf:"bytes,7,rep,name=availabilities,proto3" json:"availabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *UserInterest) Reset()         { *m = UserInterest{} }
func (m *UserInterest) String() string { return proto.CompactTextString(m) }
func (*UserInterest) ProtoMessage()    {}
func (*UserInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_853f2530b9dfe755, []int{0}
}

func (m *UserInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInterest.Unmarshal(m, b)
}
func (m *UserInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInterest.Marshal(b, m, deterministic)
}
func (m *UserInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInterest.Merge(m, src)
}
func (m *UserInterest) XXX_Size() int {
	return xxx_messageInfo_UserInterest.Size(m)
}
func (m *UserInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInterest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInterest proto.InternalMessageInfo

func (m *UserInterest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserInterest) GetTaxonomyType() enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType {
	if m != nil {
		return m.TaxonomyType
	}
	return enums.UserInterestTaxonomyTypeEnum_UNSPECIFIED
}

func (m *UserInterest) GetUserInterestId() *wrappers.Int64Value {
	if m != nil {
		return m.UserInterestId
	}
	return nil
}

func (m *UserInterest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserInterest) GetUserInterestParent() *wrappers.StringValue {
	if m != nil {
		return m.UserInterestParent
	}
	return nil
}

func (m *UserInterest) GetLaunchedToAll() *wrappers.BoolValue {
	if m != nil {
		return m.LaunchedToAll
	}
	return nil
}

func (m *UserInterest) GetAvailabilities() []*common.CriterionCategoryAvailability {
	if m != nil {
		return m.Availabilities
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInterest)(nil), "google.ads.googleads.v3.resources.UserInterest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/user_interest.proto", fileDescriptor_853f2530b9dfe755)
}

var fileDescriptor_853f2530b9dfe755 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0xe3, 0xb6, 0x9f, 0x3e, 0xf7, 0x07, 0xb0, 0x58, 0x98, 0x52, 0x41, 0x8a, 0x54, 0x91,
	0x0d, 0x33, 0x52, 0x4c, 0x59, 0x18, 0x21, 0xe4, 0x14, 0x54, 0xb5, 0x0b, 0x54, 0x85, 0x90, 0x05,
	0x44, 0x58, 0x13, 0x7b, 0xea, 0x0e, 0x1a, 0xcf, 0x58, 0x33, 0xe3, 0x40, 0x54, 0x95, 0x87, 0x61,
	0xc9, 0xa3, 0xf0, 0x14, 0x5d, 0xf7, 0x0d, 0x60, 0x81, 0x50, 0xec, 0xb1, 0x3b, 0xa1, 0x0a, 0x94,
	0xdd, 0x4d, 0xee, 0x39, 0xe7, 0x1e, 0x9f, 0xab, 0x3b, 0xce, 0x6e, 0xca, 0x79, 0x4a, 0x31, 0x44,
	0x89, 0x84, 0x55, 0x39, 0xab, 0x26, 0x3e, 0x14, 0x58, 0xf2, 0x42, 0xc4, 0x58, 0xc2, 0x42, 0x62,
	0x11, 0x11, 0xa6, 0xb0, 0xc0, 0x52, 0x81, 0x5c, 0x70, 0xc5, 0xdd, 0xed, 0x0a, 0x0b, 0x50, 0x22,
	0x41, 0x43, 0x03, 0x13, 0x1f, 0x34, 0xb4, 0xcd, 0x17, 0x8b, 0x94, 0x63, 0x9e, 0x65, 0x9c, 0xc1,
	0x58, 0x10, 0x85, 0x05, 0xe1, 0x2c, 0x8a, 0x91, 0xc2, 0x29, 0x17, 0xd3, 0x08, 0x4d, 0x10, 0xa1,
	0x68, 0x4c, 0x28, 0x51, 0xd3, 0x6a, 0xd0, 0xe6, 0xf3, 0x45, 0x2a, 0x98, 0x15, 0xd9, 0x6f, 0xde,
	0x22, 0x85, 0x3e, 0x71, 0xc6, 0xb3, 0x69, 0xa4, 0xa6, 0x39, 0xd6, 0x02, 0xf7, 0x6b, 0x81, 0x9c,
	0xc0, 0x63, 0x82, 0x69, 0x12, 0x8d, 0xf1, 0x09, 0x9a, 0x10, 0x2e, 0x34, 0xe0, 0x8e, 0x01, 0xa8,
	0xdd, 0xeb, 0xd6, 0x3d, 0xdd, 0x2a, 0x7f, 0x8d, 0x8b, 0x63, 0xf8, 0x51, 0xa0, 0x3c, 0xc7, 0x42,
	0xea, 0xfe, 0x96, 0x41, 0x45, 0x8c, 0x71, 0x85, 0x14, 0xe1, 0x4c, 0x77, 0x1f, 0x7c, 0x5f, 0x76,
	0xd6, 0xde, 0x48, 0x2c, 0x0e, 0xb4, 0x3d, 0xb7, 0xef, 0xac, 0xd7, 0x03, 0x22, 0x86, 0x32, 0xec,
	0x59, 0x6d, 0xab, 0xf3, 0x7f, 0xef, 0xd1, 0x79, 0x68, 0xff, 0x08, 0x1f, 0x3a, 0x3b, 0x97, 0x41,
	0xea, 0x2a, 0x27, 0x12, 0xc4, 0x3c, 0x83, 0xa6, 0x4a, 0x7f, 0xad, 0xd6, 0x78, 0x85, 0x32, 0xec,
	0x7e, 0x76, 0xd6, 0xe7, 0xbe, 0xda, 0x6b, 0xb5, 0xad, 0xce, 0x46, 0x77, 0x00, 0x16, 0x2d, 0xa8,
	0xcc, 0x0d, 0x98, 0x8a, 0x03, 0xcd, 0x1f, 0x4c, 0x73, 0xfc, 0x92, 0x15, 0xd9, 0xc2, 0x66, 0xcf,
	0x3e, 0x0f, 0xed, 0xfe, 0x9a, 0x32, 0xfe, 0x72, 0x0f, 0x9d, 0x9b, 0xf3, 0x3b, 0x20, 0x89, 0x67,
	0xb7, 0xad, 0xce, 0x6a, 0xf7, 0x6e, 0x6d, 0xa1, 0x4e, 0x0f, 0x1c, 0x30, 0xf5, 0xe4, 0xf1, 0x10,
	0xd1, 0x42, 0x2b, 0x6d, 0x14, 0xc6, 0xa0, 0x83, 0xc4, 0xdd, 0x75, 0x96, 0xca, 0x58, 0x96, 0x4a,
	0xfe, 0xd6, 0x15, 0xfe, 0x6b, 0x25, 0x08, 0x4b, 0x0d, 0x81, 0x12, 0xee, 0x9e, 0x39, 0xb7, 0xe7,
	0x2d, 0xe4, 0x48, 0x60, 0xa6, 0xbc, 0xe5, 0x6b, 0xc8, 0xfc, 0x63, 0xf6, 0xae, 0xe9, 0xf8, 0xa8,
	0x1c, 0xe3, 0xee, 0x3b, 0x37, 0x28, 0x2a, 0x58, 0x7c, 0x82, 0x93, 0x48, 0xf1, 0x08, 0x51, 0xea,
	0xad, 0x94, 0x93, 0x37, 0xaf, 0x4c, 0xee, 0x71, 0x4e, 0x0d, 0xfb, 0xeb, 0x35, 0x6f, 0xc0, 0x43,
	0x4a, 0xdd, 0x0f, 0xce, 0x86, 0x71, 0x00, 0x04, 0x4b, 0xef, 0xbf, 0xb6, 0xdd, 0x59, 0xed, 0x3e,
	0x5b, 0xb8, 0xcb, 0xea, 0x92, 0xc0, 0x5e, 0x7d, 0x49, 0x7b, 0xfa, 0x90, 0x42, 0xe3, 0x8e, 0x74,
	0xd4, 0xf3, 0xca, 0xc1, 0xfb, 0x8b, 0xf0, 0xdd, 0x35, 0x3f, 0xda, 0xed, 0xc6, 0x85, 0x54, 0x3c,
	0xc3, 0x42, 0xc2, 0xd3, 0xba, 0x3c, 0x83, 0x66, 0x12, 0x12, 0x9e, 0xce, 0x6d, 0xe0, 0xac, 0xf7,
	0xd3, 0x72, 0x76, 0x62, 0x9e, 0x81, 0xbf, 0x3e, 0x13, 0xbd, 0x5b, 0xe6, 0xac, 0xa3, 0x59, 0x52,
	0x47, 0xd6, 0xdb, 0x43, 0xcd, 0x4b, 0x39, 0x45, 0x2c, 0x05, 0x5c, 0xa4, 0x30, 0xc5, 0xac, 0xcc,
	0x11, 0x5e, 0x5a, 0xfd, 0xc3, 0xa3, 0xf5, 0xb4, 0xa9, 0xbe, 0xb4, 0xec, 0xfd, 0x30, 0xfc, 0xda,
	0xda, 0xde, 0xaf, 0x24, 0xc3, 0x44, 0x82, 0xaa, 0x9c, 0x55, 0x43, 0x1f, 0xf4, 0x6b, 0xe4, 0xb7,
	0x1a, 0x33, 0x0a, 0x13, 0x39, 0x6a, 0x30, 0xa3, 0xa1, 0x3f, 0x6a, 0x30, 0x17, 0xad, 0x9d, 0xaa,
	0x11, 0x04, 0x61, 0x22, 0x83, 0xa0, 0x41, 0x05, 0xc1, 0xd0, 0x0f, 0x82, 0x06, 0x37, 0x5e, 0x29,
	0xcd, 0xfa, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x03, 0xb8, 0x55, 0x25, 0x60, 0x05, 0x00, 0x00,
}
