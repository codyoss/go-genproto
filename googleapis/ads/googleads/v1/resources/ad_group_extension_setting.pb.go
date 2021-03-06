// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_extension_setting.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// An ad group extension setting.
type AdGroupExtensionSetting struct {
	// Immutable. The resource name of the ad group extension setting.
	// AdGroupExtensionSetting resource names have the form:
	//
	// `customers/{customer_id}/adGroupExtensionSettings/{ad_group_id}~{extension_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The extension type of the ad group extension setting.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v1.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// Immutable. The resource name of the ad group. The linked extension feed items will
	// serve under this ad group.
	// AdGroup resource names have the form:
	//
	// `customers/{customer_id}/adGroups/{ad_group_id}`
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The resource names of the extension feed items to serve under the ad group.
	// ExtensionFeedItem resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ExtensionFeedItems []*wrappers.StringValue `protobuf:"bytes,4,rep,name=extension_feed_items,json=extensionFeedItems,proto3" json:"extension_feed_items,omitempty"`
	// The device for which the extensions will serve. Optional.
	Device               enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice `protobuf:"varint,5,opt,name=device,proto3,enum=google.ads.googleads.v1.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *AdGroupExtensionSetting) Reset()         { *m = AdGroupExtensionSetting{} }
func (m *AdGroupExtensionSetting) String() string { return proto.CompactTextString(m) }
func (*AdGroupExtensionSetting) ProtoMessage()    {}
func (*AdGroupExtensionSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_48cf797d6357afb2, []int{0}
}

func (m *AdGroupExtensionSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupExtensionSetting.Unmarshal(m, b)
}
func (m *AdGroupExtensionSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupExtensionSetting.Marshal(b, m, deterministic)
}
func (m *AdGroupExtensionSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupExtensionSetting.Merge(m, src)
}
func (m *AdGroupExtensionSetting) XXX_Size() int {
	return xxx_messageInfo_AdGroupExtensionSetting.Size(m)
}
func (m *AdGroupExtensionSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupExtensionSetting.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupExtensionSetting proto.InternalMessageInfo

func (m *AdGroupExtensionSetting) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupExtensionSetting) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *AdGroupExtensionSetting) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupExtensionSetting) GetExtensionFeedItems() []*wrappers.StringValue {
	if m != nil {
		return m.ExtensionFeedItems
	}
	return nil
}

func (m *AdGroupExtensionSetting) GetDevice() enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice {
	if m != nil {
		return m.Device
	}
	return enums.ExtensionSettingDeviceEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupExtensionSetting)(nil), "google.ads.googleads.v1.resources.AdGroupExtensionSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_extension_setting.proto", fileDescriptor_48cf797d6357afb2)
}

var fileDescriptor_48cf797d6357afb2 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0x93, 0x3f, 0xfd, 0x61, 0xa1, 0x3d, 0x58, 0x48, 0x84, 0xaa, 0x82, 0xb4, 0x52, 0xa5,
	0x08, 0xa1, 0x5d, 0x12, 0x2e, 0xc8, 0x20, 0xa4, 0xb5, 0x28, 0x11, 0x3d, 0xa0, 0xc8, 0x45, 0x39,
	0xa0, 0x20, 0x6b, 0x93, 0x9d, 0xb8, 0x46, 0xf1, 0xae, 0xe5, 0x5d, 0x07, 0xaa, 0xaa, 0x20, 0x8e,
	0xbc, 0x06, 0x47, 0x1e, 0x80, 0x87, 0xe0, 0x29, 0x7a, 0xee, 0x23, 0xf4, 0x84, 0x62, 0x7b, 0x9d,
	0x04, 0x30, 0x29, 0xb7, 0xcf, 0x3b, 0xdf, 0x7c, 0x33, 0xf3, 0xcd, 0xae, 0x91, 0x1b, 0x48, 0x19,
	0x4c, 0x81, 0x30, 0xae, 0x48, 0x0e, 0xe7, 0x68, 0xd6, 0x21, 0x09, 0x28, 0x99, 0x26, 0x63, 0x50,
	0x84, 0x71, 0x3f, 0x48, 0x64, 0x1a, 0xfb, 0xf0, 0x41, 0x83, 0x50, 0xa1, 0x14, 0xbe, 0x02, 0xad,
	0x43, 0x11, 0xe0, 0x38, 0x91, 0x5a, 0xda, 0xbb, 0x79, 0x22, 0x66, 0x5c, 0xe1, 0x52, 0x03, 0xcf,
	0x3a, 0xb8, 0xd4, 0xd8, 0x7e, 0x5a, 0x55, 0x06, 0x44, 0x1a, 0x29, 0xf2, 0x9b, 0xb2, 0xcf, 0x61,
	0x16, 0x8e, 0x21, 0x2f, 0xb0, 0xdd, 0xbd, 0x6a, 0xb6, 0x3e, 0x89, 0x4d, 0xce, 0x3d, 0x93, 0x13,
	0x87, 0x64, 0x12, 0xc2, 0x94, 0xfb, 0x23, 0x38, 0x66, 0xb3, 0x50, 0x26, 0x05, 0xe1, 0xce, 0x12,
	0xc1, 0x34, 0x5a, 0x84, 0xee, 0x16, 0xa1, 0xec, 0x6b, 0x94, 0x4e, 0xc8, 0xfb, 0x84, 0xc5, 0x31,
	0x24, 0xaa, 0x88, 0xef, 0x2c, 0xa5, 0x32, 0x21, 0xa4, 0x66, 0x3a, 0x94, 0xa2, 0x88, 0xee, 0x7d,
	0x6f, 0xa0, 0xdb, 0x94, 0xf7, 0xe6, 0x96, 0x1d, 0x98, 0xce, 0x8e, 0xf2, 0xb1, 0xec, 0xb7, 0x68,
	0xd3, 0xd4, 0xf2, 0x05, 0x8b, 0xa0, 0x69, 0xb5, 0xac, 0xf6, 0x75, 0xf7, 0xf1, 0x39, 0x6d, 0x5c,
	0xd2, 0x2e, 0x7a, 0xb8, 0xb0, 0xaf, 0x40, 0x71, 0xa8, 0xf0, 0x58, 0x46, 0xa4, 0x42, 0xd0, 0xbb,
	0x69, 0xe4, 0x5e, 0xb1, 0x08, 0xec, 0x77, 0x68, 0x6b, 0xd5, 0x8c, 0x66, 0xad, 0x65, 0xb5, 0xb7,
	0xba, 0xcf, 0x70, 0xd5, 0x8a, 0x32, 0x07, 0x71, 0x29, 0xfb, 0xfa, 0x24, 0x86, 0x03, 0x91, 0x46,
	0xab, 0x27, 0x6e, 0xfd, 0x9c, 0x36, 0xbc, 0x4d, 0x58, 0x3e, 0xb3, 0x19, 0xba, 0x66, 0x6e, 0x46,
	0xb3, 0xde, 0xb2, 0xda, 0x37, 0xba, 0x3b, 0xa6, 0x8a, 0xf1, 0x0d, 0x1f, 0xe9, 0x24, 0x14, 0xc1,
	0x80, 0x4d, 0x53, 0x70, 0xdb, 0xd9, 0x8c, 0x7b, 0xa8, 0xb5, 0x6e, 0x46, 0xef, 0x7f, 0x96, 0x03,
	0xfb, 0x13, 0xba, 0xb5, 0x18, 0x67, 0x02, 0xc0, 0xfd, 0x50, 0x43, 0xa4, 0x9a, 0xff, 0xb5, 0xea,
	0x6b, 0xcb, 0x91, 0x4b, 0xfa, 0x00, 0xdd, 0xaf, 0xac, 0x55, 0xce, 0xf7, 0x02, 0x80, 0xbf, 0xd4,
	0x10, 0x79, 0x36, 0xfc, 0x7a, 0xa4, 0xec, 0x63, 0xb4, 0x91, 0x5f, 0xc4, 0x66, 0x23, 0xf3, 0xb1,
	0x7f, 0x55, 0x1f, 0x8b, 0xf5, 0x3c, 0xcf, 0x92, 0x57, 0x0d, 0x5d, 0x09, 0x79, 0x85, 0xbe, 0xf3,
	0xd9, 0xba, 0xa0, 0x1f, 0xff, 0x7d, 0xff, 0xf6, 0xe1, 0x38, 0x55, 0x5a, 0x46, 0x90, 0x28, 0x72,
	0x6a, 0xe0, 0x19, 0x61, 0x7f, 0x66, 0x2b, 0x72, 0x5a, 0xfd, 0x98, 0xcf, 0xdc, 0x2f, 0x35, 0xb4,
	0x3f, 0x96, 0x11, 0x5e, 0xfb, 0x9c, 0xdd, 0x9d, 0x8a, 0x76, 0xfa, 0xf3, 0x4d, 0xf4, 0xad, 0x37,
	0x87, 0x85, 0x44, 0x20, 0xa7, 0x4c, 0x04, 0x58, 0x26, 0x01, 0x09, 0x40, 0x64, 0x7b, 0x22, 0x8b,
	0xc1, 0xfe, 0xf2, 0xd3, 0x79, 0x52, 0xa2, 0xaf, 0xb5, 0x7a, 0x8f, 0xd2, 0x6f, 0xb5, 0xdd, 0x5e,
	0x2e, 0x49, 0xb9, 0xc2, 0x39, 0x9c, 0xa3, 0x41, 0x07, 0x7b, 0x86, 0xf9, 0xc3, 0x70, 0x86, 0x94,
	0xab, 0x61, 0xc9, 0x19, 0x0e, 0x3a, 0xc3, 0x92, 0x73, 0x51, 0xdb, 0xcf, 0x03, 0x8e, 0x43, 0xb9,
	0x72, 0x9c, 0x92, 0xe5, 0x38, 0x83, 0x8e, 0xe3, 0x94, 0xbc, 0xd1, 0x46, 0xd6, 0xec, 0xa3, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0xdf, 0xc1, 0x70, 0x20, 0x05, 0x00, 0x00,
}
