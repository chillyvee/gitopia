// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/task.proto

package v3

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

type TaskType int32

const (
	TypeForkRepository      TaskType = 0
	TypeSetPullRequestState TaskType = 1
)

var TaskType_name = map[int32]string{
	0: "TASK_TYPE_FORK_REPOSITORY",
	1: "TASK_TYPE_SET_PULL_REQUEST_STATE",
}

var TaskType_value = map[string]int32{
	"TASK_TYPE_FORK_REPOSITORY":        0,
	"TASK_TYPE_SET_PULL_REQUEST_STATE": 1,
}

func (x TaskType) String() string {
	return proto.EnumName(TaskType_name, int32(x))
}

func (TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6920678987ef43f, []int{0}
}

type TaskState int32

const (
	StatePending TaskState = 0
	StateSuccess TaskState = 1
	StateFailure TaskState = 2
)

var TaskState_name = map[int32]string{
	0: "TASK_STATE_PENDING",
	1: "TASK_STATE_SUCCESS",
	2: "TASK_STATE_FAILURE",
}

var TaskState_value = map[string]int32{
	"TASK_STATE_PENDING": 0,
	"TASK_STATE_SUCCESS": 1,
	"TASK_STATE_FAILURE": 2,
}

func (x TaskState) String() string {
	return proto.EnumName(TaskState_name, int32(x))
}

func (TaskState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6920678987ef43f, []int{1}
}

type Task struct {
	Id       uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type     TaskType  `protobuf:"varint,2,opt,name=type,proto3,enum=gitopia.gitopia.gitopia.TaskType" json:"type,omitempty"`
	State    TaskState `protobuf:"varint,3,opt,name=state,proto3,enum=gitopia.gitopia.gitopia.TaskState" json:"state,omitempty"`
	Message  string    `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Creator  string    `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Provider string    `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6920678987ef43f, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Task.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return m.Size()
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetType() TaskType {
	if m != nil {
		return m.Type
	}
	return TypeForkRepository
}

func (m *Task) GetState() TaskState {
	if m != nil {
		return m.State
	}
	return StatePending
}

func (m *Task) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Task) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Task) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func init() {
	// proto.RegisterEnum("gitopia.gitopia.gitopia.TaskType", TaskType_name, TaskType_value)
	// proto.RegisterEnum("gitopia.gitopia.gitopia.TaskState", TaskState_name, TaskState_value)
	// proto.RegisterType((*Task)(nil), "gitopia.gitopia.gitopia.Task")
}

func init() { proto.RegisterFile("gitopia/task.proto", fileDescriptor_a6920678987ef43f) }

var fileDescriptor_a6920678987ef43f = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0x92, 0x8d, 0xcd, 0x42, 0x53, 0x64, 0x21, 0x16, 0x82, 0x14, 0x85, 0x9d, 0xaa,
	0x1e, 0x52, 0x09, 0x34, 0x89, 0x6b, 0xd9, 0x5c, 0x54, 0xad, 0x5a, 0x83, 0xed, 0x1e, 0xc6, 0x25,
	0xca, 0x5a, 0x2b, 0x58, 0xed, 0xe6, 0x60, 0x3b, 0x88, 0x7e, 0x03, 0xd4, 0x0b, 0x7c, 0x81, 0x9e,
	0xf8, 0x32, 0x1c, 0x77, 0x83, 0x23, 0x6a, 0xbf, 0x08, 0x8a, 0xdb, 0x94, 0xa9, 0x12, 0x3b, 0xbd,
	0xf7, 0xfc, 0x7e, 0xff, 0xbf, 0xde, 0x93, 0x1f, 0x44, 0xb9, 0x30, 0xb2, 0x10, 0x59, 0xdb, 0x64,
	0x7a, 0x12, 0x17, 0x4a, 0x1a, 0x89, 0x8e, 0x37, 0x6f, 0xf1, 0x4e, 0x0c, 0x9e, 0xe6, 0x32, 0x97,
	0x96, 0x69, 0x57, 0xd9, 0x1a, 0x3f, 0xf9, 0x05, 0xa0, 0xcb, 0x32, 0x3d, 0x41, 0x47, 0xb0, 0x21,
	0xc6, 0x3e, 0x88, 0x40, 0xd3, 0x25, 0x0d, 0x31, 0x46, 0xa7, 0xd0, 0x35, 0xb3, 0x82, 0xfb, 0x8d,
	0x08, 0x34, 0x8f, 0x5e, 0xbd, 0x8c, 0xff, 0x63, 0x1b, 0x57, 0x62, 0x36, 0x2b, 0x38, 0xb1, 0x38,
	0x7a, 0x03, 0xf7, 0xb4, 0xc9, 0x0c, 0xf7, 0x1f, 0x59, 0xdd, 0xc9, 0x83, 0x3a, 0x5a, 0x91, 0x64,
	0x2d, 0x40, 0x3e, 0x7c, 0x7c, 0xc3, 0xb5, 0xce, 0x72, 0xee, 0xbb, 0x11, 0x68, 0x1e, 0x92, 0xba,
	0xac, 0x3a, 0x23, 0xc5, 0x33, 0x23, 0x95, 0xbf, 0xb7, 0xee, 0x6c, 0x4a, 0x14, 0xc0, 0x83, 0x42,
	0xc9, 0xcf, 0x62, 0xcc, 0x95, 0xbf, 0x6f, 0x5b, 0xdb, 0xba, 0x35, 0x07, 0xf0, 0xa0, 0x1e, 0x0e,
	0x9d, 0xc2, 0xe7, 0xac, 0x43, 0x2f, 0x52, 0x76, 0x95, 0xe0, 0xb4, 0x3b, 0x20, 0x17, 0x29, 0xc1,
	0xc9, 0x80, 0xf6, 0xd8, 0x80, 0x5c, 0x79, 0x4e, 0xf0, 0x6c, 0xbe, 0x88, 0x50, 0x05, 0x76, 0xa5,
	0x9a, 0x10, 0x5e, 0x48, 0x2d, 0x8c, 0x54, 0x33, 0xd4, 0x81, 0xd1, 0x3f, 0x19, 0xc5, 0x2c, 0x4d,
	0x86, 0xfd, 0x7e, 0x4a, 0xf0, 0xfb, 0x21, 0xa6, 0x2c, 0xa5, 0xac, 0xc3, 0xb0, 0x07, 0x82, 0x17,
	0xf3, 0x45, 0x74, 0x5c, 0xa9, 0x29, 0x37, 0x49, 0x39, 0x9d, 0x12, 0xfe, 0xa9, 0xe4, 0xda, 0xd8,
	0xed, 0x02, 0xf7, 0xeb, 0x8f, 0xd0, 0x69, 0x7d, 0x03, 0xf0, 0x70, 0xbb, 0x31, 0x6a, 0x42, 0x64,
	0x6d, 0xad, 0x41, 0x9a, 0xe0, 0xcb, 0xf3, 0xde, 0xe5, 0x3b, 0xcf, 0x09, 0xbc, 0xf9, 0x22, 0x7a,
	0x62, 0x91, 0x84, 0xdf, 0x8e, 0xc5, 0x6d, 0xbe, 0x43, 0xd2, 0xe1, 0xd9, 0x19, 0xa6, 0xd4, 0x03,
	0xf7, 0x48, 0x5a, 0x8e, 0x46, 0x5c, 0xeb, 0x1d, 0xb2, 0xdb, 0xe9, 0xf5, 0x87, 0x04, 0x7b, 0x8d,
	0x7b, 0x64, 0x37, 0x13, 0xd3, 0x52, 0x6d, 0x26, 0x7a, 0x7b, 0xfe, 0x73, 0x19, 0x82, 0xbb, 0x65,
	0x08, 0xfe, 0x2c, 0x43, 0xf0, 0x7d, 0x15, 0x3a, 0x77, 0xab, 0xd0, 0xf9, 0xbd, 0x0a, 0x9d, 0x0f,
	0xad, 0x5c, 0x98, 0x8f, 0xe5, 0x75, 0x3c, 0x92, 0x37, 0xed, 0xfa, 0xc0, 0xea, 0xf8, 0x65, 0x9b,
	0x55, 0xbf, 0xad, 0xaf, 0xf7, 0xed, 0x15, 0xbd, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x69, 0xfa,
	0x3f, 0x24, 0x8a, 0x02, 0x00, 0x00,
}

func (m *Task) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Task) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Task) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x22
	}
	if m.State != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if m.Type != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTask(dAtA []byte, offset int, v uint64) int {
	offset -= sovTask(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Task) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTask(uint64(m.Id))
	}
	if m.Type != 0 {
		n += 1 + sovTask(uint64(m.Type))
	}
	if m.State != 0 {
		n += 1 + sovTask(uint64(m.State))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	return n
}

func sovTask(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTask(x uint64) (n int) {
	return sovTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Task) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
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
			return fmt.Errorf("proto: Task: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= TaskType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= TaskState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTask
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
func skipTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTask
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
					return 0, ErrIntOverflowTask
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
					return 0, ErrIntOverflowTask
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
				return 0, ErrInvalidLengthTask
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTask
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTask
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTask        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTask          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTask = fmt.Errorf("proto: unexpected end of group")
)
