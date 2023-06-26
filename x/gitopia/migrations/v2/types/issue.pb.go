// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/issue.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Issue_State int32

const (
	Issue_OPEN   Issue_State = 0
	Issue_CLOSED Issue_State = 1
)

var Issue_State_name = map[int32]string{
	0: "OPEN",
	1: "CLOSED",
}

var Issue_State_value = map[string]int32{
	"OPEN":   0,
	"CLOSED": 1,
}

func (x Issue_State) String() string {
	return proto.EnumName(Issue_State_name, int32(x))
}

func (Issue_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cf64e56e9098bda, []int{0, 0}
}

type Issue struct {
	Creator       string            `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id            uint64            `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Iid           uint64            `protobuf:"varint,3,opt,name=iid,proto3" json:"iid,omitempty"`
	Title         string            `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	State         Issue_State       `protobuf:"varint,5,opt,name=state,proto3,enum=gitopia.gitopia.gitopia.v2.Issue_State" json:"state,omitempty"`
	Description   string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	CommentsCount uint64            `protobuf:"varint,7,opt,name=commentsCount,proto3" json:"commentsCount,omitempty"`
	PullRequests  []*PullRequestIid `protobuf:"bytes,8,rep,name=pullRequests,proto3" json:"pullRequests,omitempty"`
	RepositoryId  uint64            `protobuf:"varint,9,opt,name=repositoryId,proto3" json:"repositoryId,omitempty"`
	Labels        []uint64          `protobuf:"varint,10,rep,packed,name=labels,proto3" json:"labels,omitempty"`
	Weight        uint64            `protobuf:"varint,11,opt,name=weight,proto3" json:"weight,omitempty"`
	Assignees     []string          `protobuf:"bytes,12,rep,name=assignees,proto3" json:"assignees,omitempty"`
	Bounties      []uint64          `protobuf:"varint,13,rep,packed,name=bounties,proto3" json:"bounties,omitempty"`
	CreatedAt     int64             `protobuf:"varint,14,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     int64             `protobuf:"varint,15,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	ClosedAt      int64             `protobuf:"varint,16,opt,name=closedAt,proto3" json:"closedAt,omitempty"`
	ClosedBy      string            `protobuf:"bytes,17,opt,name=closedBy,proto3" json:"closedBy,omitempty"`
}

func (m *Issue) Reset()         { *m = Issue{} }
func (m *Issue) String() string { return proto.CompactTextString(m) }
func (*Issue) ProtoMessage()    {}
func (*Issue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf64e56e9098bda, []int{0}
}
func (m *Issue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Issue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Issue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Issue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Issue.Merge(m, src)
}
func (m *Issue) XXX_Size() int {
	return m.Size()
}
func (m *Issue) XXX_DiscardUnknown() {
	xxx_messageInfo_Issue.DiscardUnknown(m)
}

var xxx_messageInfo_Issue proto.InternalMessageInfo

func (m *Issue) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Issue) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Issue) GetIid() uint64 {
	if m != nil {
		return m.Iid
	}
	return 0
}

func (m *Issue) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Issue) GetState() Issue_State {
	if m != nil {
		return m.State
	}
	return Issue_OPEN
}

func (m *Issue) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Issue) GetCommentsCount() uint64 {
	if m != nil {
		return m.CommentsCount
	}
	return 0
}

func (m *Issue) GetPullRequests() []*PullRequestIid {
	if m != nil {
		return m.PullRequests
	}
	return nil
}

func (m *Issue) GetRepositoryId() uint64 {
	if m != nil {
		return m.RepositoryId
	}
	return 0
}

func (m *Issue) GetLabels() []uint64 {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Issue) GetWeight() uint64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Issue) GetAssignees() []string {
	if m != nil {
		return m.Assignees
	}
	return nil
}

func (m *Issue) GetBounties() []uint64 {
	if m != nil {
		return m.Bounties
	}
	return nil
}

func (m *Issue) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Issue) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Issue) GetClosedAt() int64 {
	if m != nil {
		return m.ClosedAt
	}
	return 0
}

func (m *Issue) GetClosedBy() string {
	if m != nil {
		return m.ClosedBy
	}
	return ""
}

func init() {
	proto.RegisterEnum("gitopia.gitopia.gitopia.v2.Issue_State", Issue_State_name, Issue_State_value)
	proto.RegisterType((*Issue)(nil), "gitopia.gitopia.gitopia.v2.Issue")
}

func init() { proto.RegisterFile("gitopia/issue.proto", fileDescriptor_4cf64e56e9098bda) }

var fileDescriptor_4cf64e56e9098bda = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xcd, 0xc6, 0x76, 0x9a, 0x4c, 0xd2, 0x7c, 0xf9, 0x96, 0x0a, 0xad, 0x22, 0xb0, 0xac, 0x08,
	0x09, 0x8b, 0x83, 0x2d, 0x85, 0x23, 0x02, 0x89, 0x96, 0x1e, 0x22, 0xa1, 0xb6, 0x72, 0x2f, 0x88,
	0x9b, 0x63, 0xaf, 0xdc, 0x91, 0x9c, 0xac, 0xf1, 0xae, 0x0b, 0xf9, 0x17, 0xfc, 0x2c, 0x8e, 0xe5,
	0xc6, 0x11, 0x25, 0x7f, 0x04, 0xed, 0x3a, 0xb1, 0x09, 0x12, 0x5c, 0xb2, 0xf3, 0xde, 0x9b, 0xf7,
	0x32, 0xde, 0x59, 0x78, 0x94, 0xa1, 0x12, 0x05, 0xc6, 0x21, 0x4a, 0x59, 0xf1, 0xa0, 0x28, 0x85,
	0x12, 0x74, 0xba, 0x27, 0x83, 0x3f, 0xcf, 0xfb, 0xf9, 0xf4, 0x2c, 0x13, 0x99, 0x30, 0x6d, 0xa1,
	0xae, 0x6a, 0xc7, 0x94, 0x1d, 0x62, 0x4a, 0x5e, 0x08, 0x89, 0x4a, 0x94, 0x9b, 0x5a, 0x99, 0x7d,
	0xb7, 0xc1, 0x59, 0xe8, 0x6c, 0xca, 0xe0, 0x24, 0x29, 0x79, 0xac, 0x44, 0xc9, 0x88, 0x47, 0xfc,
	0x41, 0x74, 0x80, 0x74, 0x0c, 0x5d, 0x4c, 0x59, 0xd7, 0x23, 0xbe, 0x1d, 0x75, 0x31, 0xa5, 0x13,
	0xb0, 0x10, 0x53, 0x66, 0x19, 0x42, 0x97, 0xf4, 0x0c, 0x1c, 0x85, 0x2a, 0xe7, 0xcc, 0x36, 0xce,
	0x1a, 0xd0, 0xd7, 0xe0, 0x48, 0x15, 0x2b, 0xce, 0x1c, 0x8f, 0xf8, 0xe3, 0xf9, 0xf3, 0xe0, 0xef,
	0x73, 0x07, 0x66, 0x86, 0xe0, 0x56, 0xb7, 0x47, 0xb5, 0x8b, 0x7a, 0x30, 0x4c, 0xb9, 0x4c, 0x4a,
	0x2c, 0x14, 0x8a, 0x35, 0xeb, 0x99, 0xe8, 0xdf, 0x29, 0xfa, 0x0c, 0x4e, 0x13, 0xb1, 0x5a, 0xf1,
	0xb5, 0x92, 0x17, 0xa2, 0x5a, 0x2b, 0x76, 0x62, 0x46, 0x3a, 0x26, 0xe9, 0x15, 0x8c, 0x8a, 0x2a,
	0xcf, 0x23, 0xfe, 0xa9, 0xe2, 0x52, 0x49, 0xd6, 0xf7, 0x2c, 0x7f, 0x38, 0x7f, 0xf1, 0xaf, 0x69,
	0x6e, 0xda, 0xfe, 0x05, 0xa6, 0xd1, 0x91, 0x9f, 0xce, 0x60, 0xd4, 0x5e, 0xe3, 0x22, 0x65, 0x03,
	0xf3, 0xa7, 0x47, 0x1c, 0x7d, 0x0c, 0xbd, 0x3c, 0x5e, 0xf2, 0x5c, 0x32, 0xf0, 0x2c, 0xdf, 0x8e,
	0xf6, 0x48, 0xf3, 0x9f, 0x39, 0x66, 0x77, 0x8a, 0x0d, 0x8d, 0x6b, 0x8f, 0xe8, 0x13, 0x18, 0xc4,
	0x52, 0x62, 0xb6, 0xe6, 0x5c, 0xb2, 0x91, 0x67, 0xf9, 0x83, 0xa8, 0x25, 0xe8, 0x14, 0xfa, 0x4b,
	0xfd, 0x29, 0xc8, 0x25, 0x3b, 0x35, 0x79, 0x0d, 0xd6, 0x4e, 0xb3, 0x27, 0x9e, 0xbe, 0x55, 0x6c,
	0xec, 0x11, 0xdf, 0x8a, 0x5a, 0x42, 0xab, 0x55, 0x91, 0xee, 0xd5, 0xff, 0x6a, 0xb5, 0x21, 0x74,
	0x6e, 0x92, 0x0b, 0x69, 0xc4, 0x89, 0x11, 0x1b, 0xdc, 0x6a, 0xe7, 0x1b, 0xf6, 0xbf, 0xb9, 0xfa,
	0x06, 0xcf, 0x9e, 0x82, 0x63, 0x36, 0x45, 0xfb, 0x60, 0x5f, 0xdf, 0x5c, 0x5e, 0x4d, 0x3a, 0x14,
	0xa0, 0x77, 0xf1, 0xfe, 0xfa, 0xf6, 0xf2, 0xdd, 0x84, 0x9c, 0x7f, 0xf8, 0xb6, 0x75, 0xc9, 0xc3,
	0xd6, 0x25, 0x3f, 0xb7, 0x2e, 0xf9, 0xba, 0x73, 0x3b, 0x0f, 0x3b, 0xb7, 0xf3, 0x63, 0xe7, 0x76,
	0x3e, 0xbe, 0xc9, 0x50, 0xdd, 0x55, 0xcb, 0x20, 0x11, 0xab, 0xf0, 0xf0, 0x24, 0x0f, 0xe7, 0x97,
	0xa6, 0x5a, 0x61, 0x56, 0xc6, 0x7a, 0xc1, 0x32, 0xbc, 0x9f, 0x87, 0x6a, 0x53, 0x70, 0xf9, 0xca,
	0xfc, 0x2e, 0x7b, 0xe6, 0xd1, 0xbe, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xa7, 0x2c, 0x7c,
	0x17, 0x03, 0x00, 0x00,
}

func (m *Issue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Issue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Issue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClosedBy) > 0 {
		i -= len(m.ClosedBy)
		copy(dAtA[i:], m.ClosedBy)
		i = encodeVarintIssue(dAtA, i, uint64(len(m.ClosedBy)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if m.ClosedAt != 0 {
		i = encodeVarintIssue(dAtA, i, uint64(m.ClosedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintIssue(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x78
	}
	if m.CreatedAt != 0 {
		i = encodeVarintIssue(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x70
	}
	if len(m.Bounties) > 0 {
		dAtA2 := make([]byte, len(m.Bounties)*10)
		var j1 int
		for _, num := range m.Bounties {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintIssue(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Assignees) > 0 {
		for iNdEx := len(m.Assignees) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Assignees[iNdEx])
			copy(dAtA[i:], m.Assignees[iNdEx])
			i = encodeVarintIssue(dAtA, i, uint64(len(m.Assignees[iNdEx])))
			i--
			dAtA[i] = 0x62
		}
	}
	if m.Weight != 0 {
		i = encodeVarintIssue(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x58
	}
	if len(m.Labels) > 0 {
		dAtA4 := make([]byte, len(m.Labels)*10)
		var j3 int
		for _, num := range m.Labels {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintIssue(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x52
	}
	if m.RepositoryId != 0 {
		i = encodeVarintIssue(dAtA, i, uint64(m.RepositoryId))
		i--
		dAtA[i] = 0x48
	}
	if len(m.PullRequests) > 0 {
		for iNdEx := len(m.PullRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PullRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIssue(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.CommentsCount != 0 {
		i = encodeVarintIssue(dAtA, i, uint64(m.CommentsCount))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintIssue(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if m.State != 0 {
		i = encodeVarintIssue(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintIssue(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x22
	}
	if m.Iid != 0 {
		i = encodeVarintIssue(dAtA, i, uint64(m.Iid))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintIssue(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintIssue(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIssue(dAtA []byte, offset int, v uint64) int {
	offset -= sovIssue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Issue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovIssue(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovIssue(uint64(m.Id))
	}
	if m.Iid != 0 {
		n += 1 + sovIssue(uint64(m.Iid))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovIssue(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovIssue(uint64(m.State))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovIssue(uint64(l))
	}
	if m.CommentsCount != 0 {
		n += 1 + sovIssue(uint64(m.CommentsCount))
	}
	if len(m.PullRequests) > 0 {
		for _, e := range m.PullRequests {
			l = e.Size()
			n += 1 + l + sovIssue(uint64(l))
		}
	}
	if m.RepositoryId != 0 {
		n += 1 + sovIssue(uint64(m.RepositoryId))
	}
	if len(m.Labels) > 0 {
		l = 0
		for _, e := range m.Labels {
			l += sovIssue(uint64(e))
		}
		n += 1 + sovIssue(uint64(l)) + l
	}
	if m.Weight != 0 {
		n += 1 + sovIssue(uint64(m.Weight))
	}
	if len(m.Assignees) > 0 {
		for _, s := range m.Assignees {
			l = len(s)
			n += 1 + l + sovIssue(uint64(l))
		}
	}
	if len(m.Bounties) > 0 {
		l = 0
		for _, e := range m.Bounties {
			l += sovIssue(uint64(e))
		}
		n += 1 + sovIssue(uint64(l)) + l
	}
	if m.CreatedAt != 0 {
		n += 1 + sovIssue(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovIssue(uint64(m.UpdatedAt))
	}
	if m.ClosedAt != 0 {
		n += 2 + sovIssue(uint64(m.ClosedAt))
	}
	l = len(m.ClosedBy)
	if l > 0 {
		n += 2 + l + sovIssue(uint64(l))
	}
	return n
}

func sovIssue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssue(x uint64) (n int) {
	return sovIssue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Issue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue
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
			return fmt.Errorf("proto: Issue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Issue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
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
				return ErrInvalidLengthIssue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssue
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
					return ErrIntOverflowIssue
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
				return fmt.Errorf("proto: wrong wireType = %d for field Iid", wireType)
			}
			m.Iid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Iid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
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
				return ErrInvalidLengthIssue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Issue_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
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
				return ErrInvalidLengthIssue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommentsCount", wireType)
			}
			m.CommentsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommentsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PullRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
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
				return ErrInvalidLengthIssue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIssue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PullRequests = append(m.PullRequests, &PullRequestIid{})
			if err := m.PullRequests[len(m.PullRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepositoryId", wireType)
			}
			m.RepositoryId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
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
		case 10:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Labels = append(m.Labels, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthIssue
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthIssue
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Labels) == 0 {
					m.Labels = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowIssue
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Labels = append(m.Labels, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assignees", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
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
				return ErrInvalidLengthIssue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assignees = append(m.Assignees, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Bounties = append(m.Bounties, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowIssue
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthIssue
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthIssue
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Bounties) == 0 {
					m.Bounties = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowIssue
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Bounties = append(m.Bounties, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Bounties", wireType)
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
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
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
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
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosedAt", wireType)
			}
			m.ClosedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClosedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue
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
				return ErrInvalidLengthIssue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClosedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIssue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIssue
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
func skipIssue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue
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
					return 0, ErrIntOverflowIssue
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
					return 0, ErrIntOverflowIssue
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
				return 0, ErrInvalidLengthIssue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIssue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIssue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIssue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIssue = fmt.Errorf("proto: unexpected end of group")
)
