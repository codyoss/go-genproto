// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1/gcs_fileset_spec.proto

package datacatalog

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Describes a Cloud Storage fileset entry.
type GcsFilesetSpec struct {
	// Required. Patterns to identify a set of files in Google Cloud Storage. See [Cloud
	// Storage documentation](/storage/docs/gsutil/addlhelp/WildcardNames) for
	// more information. Note that bucket wildcards are currently not supported.
	//
	// Examples of valid file_patterns:
	//
	//  * `gs://bucket_name/dir/*`: matches all files within `bucket_name/dir`
	//                              directory.
	//  * `gs://bucket_name/dir/**`: matches all files in `bucket_name/dir`
	//                               spanning all subdirectories.
	//  * `gs://bucket_name/file*`: matches files prefixed by `file` in
	//                              `bucket_name`
	//  * `gs://bucket_name/??.txt`: matches files with two characters followed by
	//                               `.txt` in `bucket_name`
	//  * `gs://bucket_name/[aeiou].txt`: matches files that contain a single
	//                                    vowel character followed by `.txt` in
	//                                    `bucket_name`
	//  * `gs://bucket_name/[a-m].txt`: matches files that contain `a`, `b`, ...
	//                                  or `m` followed by `.txt` in `bucket_name`
	//  * `gs://bucket_name/a/*/b`: matches all files in `bucket_name` that match
	//                              `a/*/b` pattern, such as `a/c/b`, `a/d/b`
	//  * `gs://another_bucket/a.txt`: matches `gs://another_bucket/a.txt`
	//
	// You can combine wildcards to provide more powerful matches, for example:
	//
	//  * `gs://bucket_name/[a-m]??.j*g`
	FilePatterns []string `protobuf:"bytes,1,rep,name=file_patterns,json=filePatterns,proto3" json:"file_patterns,omitempty"`
	// Output only. Sample files contained in this fileset, not all files
	// contained in this fileset are represented here.
	SampleGcsFileSpecs   []*GcsFileSpec `protobuf:"bytes,2,rep,name=sample_gcs_file_specs,json=sampleGcsFileSpecs,proto3" json:"sample_gcs_file_specs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GcsFilesetSpec) Reset()         { *m = GcsFilesetSpec{} }
func (m *GcsFilesetSpec) String() string { return proto.CompactTextString(m) }
func (*GcsFilesetSpec) ProtoMessage()    {}
func (*GcsFilesetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef0234cfb0655d3, []int{0}
}

func (m *GcsFilesetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsFilesetSpec.Unmarshal(m, b)
}
func (m *GcsFilesetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsFilesetSpec.Marshal(b, m, deterministic)
}
func (m *GcsFilesetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsFilesetSpec.Merge(m, src)
}
func (m *GcsFilesetSpec) XXX_Size() int {
	return xxx_messageInfo_GcsFilesetSpec.Size(m)
}
func (m *GcsFilesetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsFilesetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GcsFilesetSpec proto.InternalMessageInfo

func (m *GcsFilesetSpec) GetFilePatterns() []string {
	if m != nil {
		return m.FilePatterns
	}
	return nil
}

func (m *GcsFilesetSpec) GetSampleGcsFileSpecs() []*GcsFileSpec {
	if m != nil {
		return m.SampleGcsFileSpecs
	}
	return nil
}

// Specifications of a single file in Cloud Storage.
type GcsFileSpec struct {
	// Required. The full file path. Example: `gs://bucket_name/a/b.txt`.
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// Output only. Timestamps about the Cloud Storage file.
	GcsTimestamps *SystemTimestamps `protobuf:"bytes,2,opt,name=gcs_timestamps,json=gcsTimestamps,proto3" json:"gcs_timestamps,omitempty"`
	// Output only. The size of the file, in bytes.
	SizeBytes            int64    `protobuf:"varint,4,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GcsFileSpec) Reset()         { *m = GcsFileSpec{} }
func (m *GcsFileSpec) String() string { return proto.CompactTextString(m) }
func (*GcsFileSpec) ProtoMessage()    {}
func (*GcsFileSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef0234cfb0655d3, []int{1}
}

func (m *GcsFileSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsFileSpec.Unmarshal(m, b)
}
func (m *GcsFileSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsFileSpec.Marshal(b, m, deterministic)
}
func (m *GcsFileSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsFileSpec.Merge(m, src)
}
func (m *GcsFileSpec) XXX_Size() int {
	return xxx_messageInfo_GcsFileSpec.Size(m)
}
func (m *GcsFileSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsFileSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GcsFileSpec proto.InternalMessageInfo

func (m *GcsFileSpec) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *GcsFileSpec) GetGcsTimestamps() *SystemTimestamps {
	if m != nil {
		return m.GcsTimestamps
	}
	return nil
}

func (m *GcsFileSpec) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*GcsFilesetSpec)(nil), "google.cloud.datacatalog.v1.GcsFilesetSpec")
	proto.RegisterType((*GcsFileSpec)(nil), "google.cloud.datacatalog.v1.GcsFileSpec")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1/gcs_fileset_spec.proto", fileDescriptor_5ef0234cfb0655d3)
}

var fileDescriptor_5ef0234cfb0655d3 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x6a, 0xe2, 0x40,
	0x18, 0xc7, 0x99, 0x64, 0x59, 0xd6, 0x71, 0xf5, 0x10, 0x58, 0x08, 0x0a, 0x6b, 0xf0, 0x94, 0xc3,
	0x6e, 0x42, 0xdc, 0x5b, 0xf6, 0x54, 0x2d, 0x7a, 0x15, 0x2d, 0x42, 0x8b, 0x90, 0x8e, 0xe3, 0x38,
	0x06, 0x12, 0x67, 0xc8, 0x37, 0x15, 0xec, 0xb3, 0xf4, 0x09, 0x7a, 0xe8, 0x83, 0xb4, 0x2f, 0xd1,
	0x73, 0x9f, 0xa0, 0xc7, 0x92, 0x4c, 0x52, 0x73, 0x28, 0x39, 0xce, 0xff, 0x37, 0xff, 0xef, 0xff,
	0xff, 0xf8, 0xf0, 0x88, 0x0b, 0xc1, 0x13, 0xe6, 0xd3, 0x44, 0xdc, 0x6d, 0xfd, 0x2d, 0x51, 0x84,
	0x12, 0x45, 0x12, 0xc1, 0xfd, 0x63, 0xe0, 0x73, 0x0a, 0xd1, 0x2e, 0x4e, 0x18, 0x30, 0x15, 0x81,
	0x64, 0xd4, 0x93, 0x99, 0x50, 0xc2, 0xea, 0x6b, 0x8f, 0x57, 0x78, 0xbc, 0x9a, 0xc7, 0x3b, 0x06,
	0xbd, 0x41, 0x39, 0x90, 0xc8, 0xd8, 0xdf, 0xc5, 0x2c, 0xd9, 0x46, 0x1b, 0xb6, 0x27, 0xc7, 0x58,
	0x64, 0xda, 0xdd, 0xfb, 0xd3, 0x94, 0xa8, 0xe2, 0x94, 0x81, 0x22, 0xa9, 0x04, 0xfd, 0x7b, 0xf8,
	0x80, 0x70, 0x77, 0x46, 0x61, 0xaa, 0x5b, 0x2c, 0x25, 0xa3, 0x96, 0x8b, 0x3b, 0x79, 0xa9, 0x48,
	0x12, 0xa5, 0x58, 0x76, 0x00, 0x1b, 0x39, 0xa6, 0xdb, 0x1a, 0x9b, 0xaf, 0x17, 0xc6, 0xe2, 0x67,
	0x4e, 0xe6, 0x25, 0xb0, 0x6e, 0xf1, 0x2f, 0x20, 0xa9, 0x4c, 0x58, 0x54, 0x6d, 0x52, 0xac, 0x01,
	0xb6, 0xe1, 0x98, 0x6e, 0x7b, 0xe4, 0x7a, 0x0d, 0x8b, 0x78, 0x65, 0x6a, 0x1e, 0x99, 0xcf, 0x36,
	0x17, 0x96, 0x9e, 0x55, 0xd3, 0x61, 0xf8, 0x84, 0x70, 0xbb, 0x26, 0x58, 0x0e, 0x6e, 0x55, 0xdd,
	0xf6, 0x36, 0x72, 0x50, 0xd5, 0xeb, 0x47, 0xd9, 0x6b, 0x6f, 0x5d, 0xe3, 0x6e, 0x5e, 0xe6, 0xbc,
	0xa8, 0x6d, 0x38, 0xc8, 0x6d, 0x8f, 0xfe, 0x36, 0x96, 0x59, 0x9e, 0x40, 0xb1, 0xf4, 0xea, 0xd3,
	0xa4, 0x1b, 0x75, 0x38, 0x85, 0xb3, 0x66, 0x0d, 0x31, 0x86, 0xf8, 0x9e, 0x45, 0x9b, 0x93, 0x62,
	0x60, 0x7f, 0x73, 0x90, 0x6b, 0xea, 0x7f, 0xad, 0x5c, 0x1e, 0xe7, 0xea, 0xf8, 0x05, 0xe1, 0x01,
	0x15, 0x69, 0x53, 0xd8, 0x1c, 0xdd, 0x4c, 0x4b, 0xcc, 0x45, 0x42, 0x0e, 0xdc, 0x13, 0x19, 0xf7,
	0x39, 0x3b, 0x14, 0x17, 0xf1, 0x35, 0x22, 0x32, 0x86, 0x2f, 0x4f, 0xf8, 0xbf, 0xf6, 0x7c, 0x47,
	0xe8, 0xd1, 0xe8, 0xcf, 0xf4, 0xac, 0x49, 0x11, 0x75, 0x49, 0x14, 0x99, 0x94, 0x51, 0xab, 0xe0,
	0xb9, 0xa2, 0xeb, 0x82, 0xae, 0x6b, 0x74, 0xbd, 0x0a, 0xde, 0x8c, 0xdf, 0x9a, 0x86, 0x61, 0x81,
	0xc3, 0xb0, 0xc6, 0xc3, 0x70, 0x15, 0x6c, 0xbe, 0x17, 0x95, 0xfe, 0x7d, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x01, 0x41, 0x92, 0xcd, 0xc6, 0x02, 0x00, 0x00,
}
