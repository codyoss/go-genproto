// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/genomics/v1/readalignment.proto
// DO NOT EDIT!

package google_genomics_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf3 "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A linear alignment can be represented by one CIGAR string. Describes the
// mapped position and local alignment of the read to the reference.
type LinearAlignment struct {
	// The position of this alignment.
	Position *Position `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	// The mapping quality of this alignment. Represents how likely
	// the read maps to this position as opposed to other locations.
	//
	// Specifically, this is -10 log10 Pr(mapping position is wrong), rounded to
	// the nearest integer.
	MappingQuality int32 `protobuf:"varint,2,opt,name=mapping_quality,json=mappingQuality" json:"mapping_quality,omitempty"`
	// Represents the local alignment of this sequence (alignment matches, indels,
	// etc) against the reference.
	Cigar []*CigarUnit `protobuf:"bytes,3,rep,name=cigar" json:"cigar,omitempty"`
}

func (m *LinearAlignment) Reset()                    { *m = LinearAlignment{} }
func (m *LinearAlignment) String() string            { return proto.CompactTextString(m) }
func (*LinearAlignment) ProtoMessage()               {}
func (*LinearAlignment) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *LinearAlignment) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *LinearAlignment) GetCigar() []*CigarUnit {
	if m != nil {
		return m.Cigar
	}
	return nil
}

// A read alignment describes a linear alignment of a string of DNA to a
// [reference sequence][google.genomics.v1.Reference], in addition to metadata
// about the fragment (the molecule of DNA sequenced) and the read (the bases
// which were read by the sequencer). A read is equivalent to a line in a SAM
// file. A read belongs to exactly one read group and exactly one
// [read group set][google.genomics.v1.ReadGroupSet].
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
//
// ### Reverse-stranded reads
//
// Mapped reads (reads having a non-null `alignment`) can be aligned to either
// the forward or the reverse strand of their associated reference. Strandedness
// of a mapped read is encoded by `alignment.position.reverseStrand`.
//
// If we consider the reference to be a forward-stranded coordinate space of
// `[0, reference.length)` with `0` as the left-most position and
// `reference.length` as the right-most position, reads are always aligned left
// to right. That is, `alignment.position.position` always refers to the
// left-most reference coordinate and `alignment.cigar` describes the alignment
// of this read to the reference from left to right. All per-base fields such as
// `alignedSequence` and `alignedQuality` share this same left-to-right
// orientation; this is true of reads which are aligned to either strand. For
// reverse-stranded reads, this means that `alignedSequence` is the reverse
// complement of the bases that were originally reported by the sequencing
// machine.
//
// ### Generating a reference-aligned sequence string
//
// When interacting with mapped reads, it's often useful to produce a string
// representing the local alignment of the read to reference. The following
// pseudocode demonstrates one way of doing this:
//
//     out = ""
//     offset = 0
//     for c in read.alignment.cigar {
//       switch c.operation {
//       case "ALIGNMENT_MATCH", "SEQUENCE_MATCH", "SEQUENCE_MISMATCH":
//         out += read.alignedSequence[offset:offset+c.operationLength]
//         offset += c.operationLength
//         break
//       case "CLIP_SOFT", "INSERT":
//         offset += c.operationLength
//         break
//       case "PAD":
//         out += repeat("*", c.operationLength)
//         break
//       case "DELETE":
//         out += repeat("-", c.operationLength)
//         break
//       case "SKIP":
//         out += repeat(" ", c.operationLength)
//         break
//       case "CLIP_HARD":
//         break
//       }
//     }
//     return out
//
// ### Converting to SAM's CIGAR string
//
// The following pseudocode generates a SAM CIGAR string from the
// `cigar` field. Note that this is a lossy conversion
// (`cigar.referenceSequence` is lost).
//
//     cigarMap = {
//       "ALIGNMENT_MATCH": "M",
//       "INSERT": "I",
//       "DELETE": "D",
//       "SKIP": "N",
//       "CLIP_SOFT": "S",
//       "CLIP_HARD": "H",
//       "PAD": "P",
//       "SEQUENCE_MATCH": "=",
//       "SEQUENCE_MISMATCH": "X",
//     }
//     cigarStr = ""
//     for c in read.alignment.cigar {
//       cigarStr += c.operationLength + cigarMap[c.operation]
//     }
//     return cigarStr
type Read struct {
	// The server-generated read ID, unique across all reads. This is different
	// from the `fragmentName`.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The ID of the read group this read belongs to. A read belongs to exactly
	// one read group. This is a server-generated ID which is distinct from SAM's
	// RG tag (for that value, see
	// [ReadGroup.name][google.genomics.v1.ReadGroup.name]).
	ReadGroupId string `protobuf:"bytes,2,opt,name=read_group_id,json=readGroupId" json:"read_group_id,omitempty"`
	// The ID of the read group set this read belongs to. A read belongs to
	// exactly one read group set.
	ReadGroupSetId string `protobuf:"bytes,3,opt,name=read_group_set_id,json=readGroupSetId" json:"read_group_set_id,omitempty"`
	// The fragment name. Equivalent to QNAME (query template name) in SAM.
	FragmentName string `protobuf:"bytes,4,opt,name=fragment_name,json=fragmentName" json:"fragment_name,omitempty"`
	// The orientation and the distance between reads from the fragment are
	// consistent with the sequencing protocol (SAM flag 0x2).
	ProperPlacement bool `protobuf:"varint,5,opt,name=proper_placement,json=properPlacement" json:"proper_placement,omitempty"`
	// The fragment is a PCR or optical duplicate (SAM flag 0x400).
	DuplicateFragment bool `protobuf:"varint,6,opt,name=duplicate_fragment,json=duplicateFragment" json:"duplicate_fragment,omitempty"`
	// The observed length of the fragment, equivalent to TLEN in SAM.
	FragmentLength int32 `protobuf:"varint,7,opt,name=fragment_length,json=fragmentLength" json:"fragment_length,omitempty"`
	// The read number in sequencing. 0-based and less than numberReads. This
	// field replaces SAM flag 0x40 and 0x80.
	ReadNumber int32 `protobuf:"varint,8,opt,name=read_number,json=readNumber" json:"read_number,omitempty"`
	// The number of reads in the fragment (extension to SAM flag 0x1).
	NumberReads int32 `protobuf:"varint,9,opt,name=number_reads,json=numberReads" json:"number_reads,omitempty"`
	// Whether this read did not pass filters, such as platform or vendor quality
	// controls (SAM flag 0x200).
	FailedVendorQualityChecks bool `protobuf:"varint,10,opt,name=failed_vendor_quality_checks,json=failedVendorQualityChecks" json:"failed_vendor_quality_checks,omitempty"`
	// The linear alignment for this alignment record. This field is null for
	// unmapped reads.
	Alignment *LinearAlignment `protobuf:"bytes,11,opt,name=alignment" json:"alignment,omitempty"`
	// Whether this alignment is secondary. Equivalent to SAM flag 0x100.
	// A secondary alignment represents an alternative to the primary alignment
	// for this read. Aligners may return secondary alignments if a read can map
	// ambiguously to multiple coordinates in the genome. By convention, each read
	// has one and only one alignment where both `secondaryAlignment`
	// and `supplementaryAlignment` are false.
	SecondaryAlignment bool `protobuf:"varint,12,opt,name=secondary_alignment,json=secondaryAlignment" json:"secondary_alignment,omitempty"`
	// Whether this alignment is supplementary. Equivalent to SAM flag 0x800.
	// Supplementary alignments are used in the representation of a chimeric
	// alignment. In a chimeric alignment, a read is split into multiple
	// linear alignments that map to different reference contigs. The first
	// linear alignment in the read will be designated as the representative
	// alignment; the remaining linear alignments will be designated as
	// supplementary alignments. These alignments may have different mapping
	// quality scores. In each linear alignment in a chimeric alignment, the read
	// will be hard clipped. The `alignedSequence` and
	// `alignedQuality` fields in the alignment record will only
	// represent the bases for its respective linear alignment.
	SupplementaryAlignment bool `protobuf:"varint,13,opt,name=supplementary_alignment,json=supplementaryAlignment" json:"supplementary_alignment,omitempty"`
	// The bases of the read sequence contained in this alignment record,
	// **without CIGAR operations applied** (equivalent to SEQ in SAM).
	// `alignedSequence` and `alignedQuality` may be
	// shorter than the full read sequence and quality. This will occur if the
	// alignment is part of a chimeric alignment, or if the read was trimmed. When
	// this occurs, the CIGAR for this read will begin/end with a hard clip
	// operator that will indicate the length of the excised sequence.
	AlignedSequence string `protobuf:"bytes,14,opt,name=aligned_sequence,json=alignedSequence" json:"aligned_sequence,omitempty"`
	// The quality of the read sequence contained in this alignment record
	// (equivalent to QUAL in SAM).
	// `alignedSequence` and `alignedQuality` may be shorter than the full read
	// sequence and quality. This will occur if the alignment is part of a
	// chimeric alignment, or if the read was trimmed. When this occurs, the CIGAR
	// for this read will begin/end with a hard clip operator that will indicate
	// the length of the excised sequence.
	AlignedQuality []int32 `protobuf:"varint,15,rep,packed,name=aligned_quality,json=alignedQuality" json:"aligned_quality,omitempty"`
	// The mapping of the primary alignment of the
	// `(readNumber+1)%numberReads` read in the fragment. It replaces
	// mate position and mate strand in SAM.
	NextMatePosition *Position `protobuf:"bytes,16,opt,name=next_mate_position,json=nextMatePosition" json:"next_mate_position,omitempty"`
	// A map of additional read alignment information. This must be of the form
	// map<string, string[]> (string key mapping to a list of string values).
	Info map[string]*google_protobuf3.ListValue `protobuf:"bytes,17,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Read) Reset()                    { *m = Read{} }
func (m *Read) String() string            { return proto.CompactTextString(m) }
func (*Read) ProtoMessage()               {}
func (*Read) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *Read) GetAlignment() *LinearAlignment {
	if m != nil {
		return m.Alignment
	}
	return nil
}

func (m *Read) GetNextMatePosition() *Position {
	if m != nil {
		return m.NextMatePosition
	}
	return nil
}

func (m *Read) GetInfo() map[string]*google_protobuf3.ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*LinearAlignment)(nil), "google.genomics.v1.LinearAlignment")
	proto.RegisterType((*Read)(nil), "google.genomics.v1.Read")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/genomics/v1/readalignment.proto", fileDescriptor6)
}

var fileDescriptor6 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0x51, 0x4f, 0xdb, 0x3a,
	0x14, 0xc7, 0xd5, 0x96, 0x72, 0xa9, 0x0b, 0x6d, 0xf1, 0x95, 0xb8, 0xb9, 0x15, 0x57, 0x97, 0x95,
	0x87, 0x6d, 0x0f, 0x4b, 0x06, 0x68, 0x1b, 0x9a, 0x34, 0x6d, 0x80, 0xd8, 0xc4, 0xc4, 0x50, 0x17,
	0x34, 0x5e, 0x23, 0x37, 0x71, 0x8d, 0x45, 0x62, 0x07, 0xdb, 0xa9, 0xd6, 0x8f, 0xb4, 0xef, 0xb6,
	0x0f, 0xb0, 0xc7, 0xd9, 0x4e, 0x9c, 0x8e, 0xad, 0x0f, 0xf0, 0x02, 0xd5, 0xff, 0xfc, 0xce, 0x3f,
	0xce, 0xf9, 0xfb, 0x04, 0x9c, 0x12, 0xce, 0x49, 0x8a, 0x7d, 0xc2, 0x53, 0xc4, 0x88, 0xcf, 0x05,
	0x09, 0x08, 0x66, 0xb9, 0xe0, 0x8a, 0x07, 0x65, 0x09, 0xe5, 0x54, 0x1a, 0x8d, 0x67, 0x34, 0x96,
	0xc1, 0x6c, 0x2f, 0x10, 0x18, 0x25, 0x28, 0xa5, 0x84, 0x65, 0x98, 0x29, 0xdf, 0xa2, 0x10, 0x3a,
	0x9b, 0x8a, 0xf3, 0x67, 0x7b, 0xc3, 0xb3, 0xfb, 0x59, 0xeb, 0x3f, 0x81, 0xc4, 0x62, 0x46, 0x63,
	0x1c, 0x73, 0x36, 0xa5, 0x24, 0x40, 0x8c, 0x71, 0x85, 0x14, 0xe5, 0x4c, 0x96, 0xf6, 0xc3, 0x37,
	0x0f, 0x3f, 0x65, 0x4c, 0x09, 0x12, 0x55, 0xfb, 0xbb, 0x87, 0xb7, 0xe7, 0x5c, 0x52, 0x73, 0x82,
	0xca, 0xe1, 0x05, 0xa1, 0xea, 0xba, 0x98, 0xf8, 0x31, 0xcf, 0x82, 0xd2, 0x25, 0xb0, 0x85, 0x49,
	0x31, 0x0d, 0x72, 0x35, 0xcf, 0xb1, 0x0c, 0xa4, 0x12, 0x45, 0xac, 0xaa, 0x7f, 0x65, 0xdb, 0xe8,
	0x5b, 0x03, 0xf4, 0xcf, 0x29, 0xc3, 0x48, 0x1c, 0xb9, 0x81, 0xc1, 0x43, 0xb0, 0xe6, 0xcc, 0xbd,
	0xc6, 0x4e, 0xe3, 0x49, 0x77, 0x7f, 0xdb, 0xff, 0x73, 0x7a, 0xfe, 0xb8, 0x62, 0xc2, 0x9a, 0x86,
	0x8f, 0x41, 0x3f, 0x43, 0x79, 0x4e, 0x19, 0x89, 0x6e, 0x0b, 0x9d, 0x80, 0x9a, 0x7b, 0x4d, 0x6d,
	0xd0, 0x0e, 0x7b, 0x95, 0xfc, 0xb9, 0x54, 0xe1, 0x01, 0x68, 0xdb, 0xd7, 0xf7, 0x5a, 0x3b, 0x2d,
	0xed, 0xff, 0xdf, 0x32, 0xff, 0x13, 0x03, 0x7c, 0x61, 0x54, 0x85, 0x25, 0x3b, 0xfa, 0xbe, 0x0a,
	0x56, 0x42, 0x1d, 0x2d, 0xec, 0x81, 0x26, 0x4d, 0xec, 0xd1, 0x3a, 0xa1, 0xfe, 0x05, 0x47, 0x60,
	0xc3, 0x44, 0x1e, 0x11, 0xc1, 0x8b, 0x3c, 0xd2, 0xa5, 0xa6, 0x2d, 0x75, 0x8d, 0xf8, 0xc1, 0x68,
	0x67, 0x09, 0x7c, 0x0a, 0x36, 0x7f, 0x61, 0x24, 0x56, 0x86, 0x6b, 0x59, 0xae, 0x57, 0x73, 0x97,
	0x58, 0x69, 0x74, 0x17, 0x6c, 0x4c, 0x05, 0x22, 0x66, 0x16, 0x11, 0x43, 0x19, 0xf6, 0x56, 0x2c,
	0xb6, 0xee, 0xc4, 0x0b, 0xad, 0x69, 0xbf, 0x81, 0x9e, 0x60, 0x8e, 0x45, 0x94, 0xa7, 0x28, 0xc6,
	0x46, 0xf7, 0xda, 0x9a, 0x5b, 0x0b, 0xfb, 0xa5, 0x3e, 0x76, 0x32, 0x7c, 0x06, 0x60, 0x52, 0xe4,
	0x29, 0x8d, 0x91, 0xc2, 0x91, 0x33, 0xf1, 0x56, 0x2d, 0xbc, 0x59, 0x57, 0xde, 0x57, 0x05, 0x33,
	0xc4, 0xfa, 0xf1, 0x29, 0x66, 0x44, 0x5d, 0x7b, 0x7f, 0x95, 0x43, 0x74, 0xf2, 0xb9, 0x55, 0xe1,
	0xff, 0xc0, 0xbe, 0x61, 0xc4, 0x8a, 0x6c, 0x82, 0x85, 0xb7, 0x66, 0x21, 0x60, 0xa4, 0x0b, 0xab,
	0xc0, 0x47, 0x60, 0xbd, 0xac, 0x45, 0x46, 0x94, 0x5e, 0xc7, 0x12, 0xdd, 0x52, 0x33, 0x93, 0x94,
	0xf0, 0x2d, 0xd8, 0x9e, 0x22, 0x9a, 0xe2, 0x24, 0x9a, 0x61, 0x96, 0x70, 0xe1, 0x72, 0x8b, 0xe2,
	0x6b, 0x1c, 0xdf, 0x48, 0x0f, 0xd8, 0x53, 0xfe, 0x5b, 0x32, 0x57, 0x16, 0xa9, 0x32, 0x3c, 0xb1,
	0x00, 0x3c, 0x02, 0x9d, 0x7a, 0xd5, 0xbc, 0xae, 0xbd, 0x2d, 0xbb, 0xcb, 0xd2, 0xfc, 0xed, 0x92,
	0x85, 0x8b, 0x2e, 0x18, 0x80, 0xbf, 0xa5, 0xd9, 0xac, 0x04, 0x89, 0x79, 0xb4, 0x30, 0x5b, 0xb7,
	0x8f, 0x86, 0x75, 0x69, 0x71, 0x41, 0x5f, 0x81, 0x7f, 0x64, 0x91, 0xe7, 0xa9, 0x1d, 0xef, 0xdd,
	0xa6, 0x0d, 0xdb, 0xb4, 0x75, 0xa7, 0xbc, 0x68, 0xd4, 0xa1, 0x59, 0x54, 0xbf, 0xae, 0xc4, 0xb7,
	0x05, 0x66, 0x31, 0xf6, 0x7a, 0x36, 0xdc, 0x7e, 0xa5, 0x5f, 0x56, 0xb2, 0x49, 0xc1, 0xa1, 0xee,
	0x2a, 0xf7, 0xf5, 0x5d, 0xd5, 0x29, 0x54, 0xb2, 0xbb, 0xca, 0x1f, 0x01, 0x64, 0xf8, 0xab, 0x8a,
	0x32, 0x93, 0x6e, 0xbd, 0x37, 0x83, 0x7b, 0xec, 0xcd, 0xc0, 0xf4, 0x7d, 0xd2, 0x6d, 0x4e, 0x81,
	0x2f, 0xc1, 0x0a, 0x65, 0x53, 0xee, 0x6d, 0xda, 0xad, 0x18, 0x2d, 0xeb, 0x36, 0xb1, 0xf9, 0x67,
	0x1a, 0x3a, 0x65, 0x4a, 0xcc, 0x43, 0xcb, 0x0f, 0x2f, 0x41, 0xa7, 0x96, 0xe0, 0x00, 0xb4, 0x6e,
	0xf0, 0xbc, 0x5a, 0x0f, 0xf3, 0x13, 0x3e, 0x07, 0xed, 0x19, 0x4a, 0x0b, 0x6c, 0xf7, 0xa2, 0xbb,
	0x3f, 0x74, 0xbe, 0xee, 0x03, 0xa1, 0xc3, 0x91, 0xea, 0xca, 0x10, 0x61, 0x09, 0xbe, 0x6e, 0x1e,
	0x36, 0x8e, 0xf7, 0xc0, 0x96, 0xfe, 0x98, 0x2c, 0x39, 0xc3, 0x31, 0x34, 0x87, 0xa8, 0xa7, 0x3a,
	0x36, 0x2e, 0xe3, 0xc6, 0x8f, 0x46, 0x63, 0xb2, 0x6a, 0x1d, 0x0f, 0x7e, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xb5, 0xa0, 0xb1, 0xbc, 0xb4, 0x05, 0x00, 0x00,
}