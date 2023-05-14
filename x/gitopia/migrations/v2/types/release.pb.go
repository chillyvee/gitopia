// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/release.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Release struct {
	Creator      string        `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id           uint64        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	RepositoryId uint64        `protobuf:"varint,3,opt,name=repositoryId,proto3" json:"repositoryId,omitempty"`
	TagName      string        `protobuf:"bytes,4,opt,name=tagName,proto3" json:"tagName,omitempty"`
	Target       string        `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	Name         string        `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Description  string        `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Attachments  []*Attachment `protobuf:"bytes,8,rep,name=attachments,proto3" json:"attachments,omitempty"`
	Draft        bool          `protobuf:"varint,9,opt,name=draft,proto3" json:"draft,omitempty"`
	PreRelease   bool          `protobuf:"varint,10,opt,name=preRelease,proto3" json:"preRelease,omitempty"`
	IsTag        bool          `protobuf:"varint,11,opt,name=isTag,proto3" json:"isTag,omitempty"`
	CreatedAt    int64         `protobuf:"varint,12,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    int64         `protobuf:"varint,13,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	PublishedAt  int64         `protobuf:"varint,14,opt,name=publishedAt,proto3" json:"publishedAt,omitempty"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fd2e3bb1bcbe1fb, []int{0}
}
func (m *Release) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Release) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Release.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Release) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Release.Merge(m, src)
}
func (m *Release) XXX_Size() int {
	return m.Size()
}
func (m *Release) XXX_DiscardUnknown() {
	xxx_messageInfo_Release.DiscardUnknown(m)
}

var xxx_messageInfo_Release proto.InternalMessageInfo

func (m *Release) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Release) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Release) GetRepositoryId() uint64 {
	if m != nil {
		return m.RepositoryId
	}
	return 0
}

func (m *Release) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

func (m *Release) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Release) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Release) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Release) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Release) GetDraft() bool {
	if m != nil {
		return m.Draft
	}
	return false
}

func (m *Release) GetPreRelease() bool {
	if m != nil {
		return m.PreRelease
	}
	return false
}

func (m *Release) GetIsTag() bool {
	if m != nil {
		return m.IsTag
	}
	return false
}

func (m *Release) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Release) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Release) GetPublishedAt() int64 {
	if m != nil {
		return m.PublishedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Release)(nil), "gitopia.gitopia.gitopia.v2.Release")
}

func init() { proto.RegisterFile("gitopia/release.proto", fileDescriptor_5fd2e3bb1bcbe1fb) }

var fileDescriptor_5fd2e3bb1bcbe1fb = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x75, 0x7f, 0xd3, 0xb9, 0x43, 0x98, 0x12, 0x86, 0x94, 0xb2, 0x83, 0xf4, 0xd4,
	0xc2, 0x3c, 0x0a, 0xc2, 0x3c, 0xe9, 0xc5, 0x43, 0xf1, 0x20, 0xde, 0xb2, 0x36, 0x66, 0x81, 0x75,
	0x09, 0x49, 0x36, 0xdc, 0xb7, 0xf0, 0xbb, 0xf8, 0x25, 0x3c, 0xee, 0xe8, 0x51, 0xb6, 0x2f, 0x22,
	0x4d, 0xdb, 0xad, 0x0a, 0x5e, 0x9a, 0xf7, 0x79, 0x7e, 0xef, 0xfb, 0x42, 0x9e, 0x06, 0x9e, 0x33,
	0x6e, 0x84, 0xe4, 0x24, 0x52, 0x74, 0x49, 0x89, 0xa6, 0xa1, 0x54, 0xc2, 0x08, 0x34, 0x2e, 0xed,
	0xf0, 0xef, 0xb9, 0x99, 0x8e, 0x47, 0x4c, 0x30, 0x61, 0xdb, 0xa2, 0xbc, 0x2a, 0x26, 0xc6, 0xb8,
	0x5a, 0x44, 0x8c, 0x21, 0xc9, 0x22, 0xa3, 0x2b, 0x53, 0x90, 0xc9, 0x87, 0x03, 0xbb, 0x71, 0xb1,
	0x1d, 0x61, 0xd8, 0x4d, 0x14, 0x25, 0x46, 0x28, 0x0c, 0x7c, 0x10, 0xf4, 0xe3, 0x4a, 0xa2, 0x21,
	0x6c, 0xf2, 0x14, 0x37, 0x7d, 0x10, 0xb4, 0xe2, 0x26, 0x4f, 0xd1, 0x04, 0x0e, 0x14, 0x95, 0x42,
	0x73, 0x23, 0xd4, 0xf6, 0x21, 0xc5, 0x8e, 0x25, 0xbf, 0xbc, 0x7c, 0x9b, 0x21, 0xec, 0x91, 0x64,
	0x14, 0xb7, 0x8a, 0x6d, 0xa5, 0x44, 0x17, 0xb0, 0x63, 0x88, 0x62, 0xd4, 0xe0, 0xb6, 0x05, 0xa5,
	0x42, 0x08, 0xb6, 0x56, 0x79, 0x7b, 0xc7, 0xba, 0xb6, 0x46, 0x3e, 0x74, 0x53, 0xaa, 0x13, 0xc5,
	0xa5, 0xe1, 0x62, 0x85, 0xbb, 0x16, 0xd5, 0x2d, 0x74, 0x0f, 0xdd, 0xd3, 0xad, 0x34, 0xee, 0xf9,
	0x4e, 0xe0, 0x4e, 0xaf, 0xc2, 0xff, 0x33, 0x0a, 0x67, 0xc7, 0xf6, 0xb8, 0x3e, 0x8a, 0x46, 0xb0,
	0x9d, 0x2a, 0xf2, 0x6a, 0x70, 0xdf, 0x07, 0x41, 0x2f, 0x2e, 0x04, 0xf2, 0x20, 0x94, 0x8a, 0x96,
	0x19, 0x61, 0x68, 0x51, 0xcd, 0xc9, 0xa7, 0xb8, 0x7e, 0x22, 0x0c, 0xbb, 0xc5, 0x94, 0x15, 0xe8,
	0x12, 0xf6, 0x6d, 0x78, 0x34, 0x9d, 0x19, 0x3c, 0xf0, 0x41, 0xe0, 0xc4, 0x27, 0x23, 0xa7, 0x6b,
	0x99, 0x96, 0xf4, 0xac, 0xa0, 0x47, 0x23, 0xbf, 0xb3, 0x5c, 0xcf, 0x97, 0x5c, 0x2f, 0x2c, 0x1f,
	0x5a, 0x5e, 0xb7, 0xee, 0x9e, 0x3f, 0xf7, 0x1e, 0xd8, 0xed, 0x3d, 0xf0, 0xbd, 0xf7, 0xc0, 0xfb,
	0xc1, 0x6b, 0xec, 0x0e, 0x5e, 0xe3, 0xeb, 0xe0, 0x35, 0x5e, 0x6e, 0x19, 0x37, 0x8b, 0xf5, 0x3c,
	0x4c, 0x44, 0x16, 0x55, 0x3f, 0xbd, 0x3a, 0xdf, 0x8e, 0x55, 0xc6, 0x99, 0x22, 0x79, 0x7a, 0x3a,
	0xda, 0x4c, 0x23, 0xb3, 0x95, 0x54, 0xdf, 0xd8, 0xef, 0xbc, 0x63, 0x9f, 0xc5, 0xf5, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x17, 0xfa, 0x91, 0xb8, 0x7b, 0x02, 0x00, 0x00,
}

func (m *Release) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Release) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Release) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PublishedAt != 0 {
		i = encodeVarintRelease(dAtA, i, uint64(m.PublishedAt))
		i--
		dAtA[i] = 0x70
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintRelease(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x68
	}
	if m.CreatedAt != 0 {
		i = encodeVarintRelease(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x60
	}
	if m.IsTag {
		i--
		if m.IsTag {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.PreRelease {
		i--
		if m.PreRelease {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.Draft {
		i--
		if m.Draft {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.Attachments) > 0 {
		for iNdEx := len(m.Attachments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attachments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRelease(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintRelease(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRelease(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintRelease(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TagName) > 0 {
		i -= len(m.TagName)
		copy(dAtA[i:], m.TagName)
		i = encodeVarintRelease(dAtA, i, uint64(len(m.TagName)))
		i--
		dAtA[i] = 0x22
	}
	if m.RepositoryId != 0 {
		i = encodeVarintRelease(dAtA, i, uint64(m.RepositoryId))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintRelease(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRelease(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRelease(dAtA []byte, offset int, v uint64) int {
	offset -= sovRelease(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Release) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRelease(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovRelease(uint64(m.Id))
	}
	if m.RepositoryId != 0 {
		n += 1 + sovRelease(uint64(m.RepositoryId))
	}
	l = len(m.TagName)
	if l > 0 {
		n += 1 + l + sovRelease(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovRelease(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRelease(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovRelease(uint64(l))
	}
	if len(m.Attachments) > 0 {
		for _, e := range m.Attachments {
			l = e.Size()
			n += 1 + l + sovRelease(uint64(l))
		}
	}
	if m.Draft {
		n += 2
	}
	if m.PreRelease {
		n += 2
	}
	if m.IsTag {
		n += 2
	}
	if m.CreatedAt != 0 {
		n += 1 + sovRelease(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovRelease(uint64(m.UpdatedAt))
	}
	if m.PublishedAt != 0 {
		n += 1 + sovRelease(uint64(m.PublishedAt))
	}
	return n
}

func sovRelease(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRelease(x uint64) (n int) {
	return sovRelease(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Release) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelease
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Release: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Release: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepositoryId", wireType)
			}
			m.RepositoryId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RepositoryId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attachments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRelease
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRelease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attachments = append(m.Attachments, &Attachment{})
			if err := m.Attachments[len(m.Attachments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Draft", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Draft = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreRelease", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PreRelease = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsTag", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsTag = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublishedAt", wireType)
			}
			m.PublishedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PublishedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRelease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelease
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRelease(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRelease
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRelease
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRelease
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRelease
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRelease
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRelease        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRelease          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRelease = fmt.Errorf("proto: unexpected end of group")
)