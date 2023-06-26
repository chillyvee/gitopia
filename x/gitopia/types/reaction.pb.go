// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/gitopia/reaction.proto

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

type Emoji int32

const (
	EmojiThumbsUp   Emoji = 0
	EmojiThumbsDown Emoji = 1
)

var Emoji_name = map[int32]string{
	0: "EMOJI_THUMBS_UP",
	1: "EMOJI_THUMBS_DOWN",
}

var Emoji_value = map[string]int32{
	"EMOJI_THUMBS_UP":   0,
	"EMOJI_THUMBS_DOWN": 1,
}

func (x Emoji) String() string {
	return proto.EnumName(Emoji_name, int32(x))
}

func (Emoji) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_693bbc73a7a9ebea, []int{0}
}

type Reaction struct {
	Address string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Emojis  []Emoji `protobuf:"varint,2,rep,packed,name=emojis,proto3,enum=gitopia.gitopia.gitopia.Emoji" json:"emojis,omitempty"`
}

func (m *Reaction) Reset()         { *m = Reaction{} }
func (m *Reaction) String() string { return proto.CompactTextString(m) }
func (*Reaction) ProtoMessage()    {}
func (*Reaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_693bbc73a7a9ebea, []int{0}
}
func (m *Reaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Reaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reaction.Merge(m, src)
}
func (m *Reaction) XXX_Size() int {
	return m.Size()
}
func (m *Reaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Reaction.DiscardUnknown(m)
}

var xxx_messageInfo_Reaction proto.InternalMessageInfo

func (m *Reaction) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Reaction) GetEmojis() []Emoji {
	if m != nil {
		return m.Emojis
	}
	return nil
}

func init() {
	proto.RegisterEnum("gitopia.gitopia.gitopia.Emoji", Emoji_name, Emoji_value)
	proto.RegisterType((*Reaction)(nil), "gitopia.gitopia.gitopia.Reaction")
}

func init() { proto.RegisterFile("gitopia/gitopia/reaction.proto", fileDescriptor_693bbc73a7a9ebea) }

var fileDescriptor_693bbc73a7a9ebea = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0x2c, 0xc9,
	0x2f, 0xc8, 0x4c, 0xd4, 0x87, 0xd1, 0x45, 0xa9, 0x89, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0xe2, 0x50, 0x71, 0x3d, 0x34, 0x5a, 0x4a, 0x24, 0x3d, 0x3f, 0x3d,
	0x1f, 0xac, 0x46, 0x1f, 0xc4, 0x82, 0x28, 0x57, 0x8a, 0xe1, 0xe2, 0x08, 0x82, 0x1a, 0x20, 0x24,
	0xc1, 0xc5, 0x9e, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0xe3, 0x0a, 0x99, 0x71, 0xb1, 0xa5, 0xe6, 0xe6, 0x67, 0x65, 0x16, 0x4b, 0x30, 0x29, 0x30,
	0x6b, 0xf0, 0x19, 0xc9, 0xe9, 0xe1, 0xb0, 0x45, 0xcf, 0x15, 0xa4, 0x2c, 0x08, 0xaa, 0x5a, 0x2b,
	0x91, 0x8b, 0x15, 0x2c, 0x20, 0xa4, 0xc6, 0xc5, 0xef, 0xea, 0xeb, 0xef, 0xe5, 0x19, 0x1f, 0xe2,
	0x11, 0xea, 0xeb, 0x14, 0x1c, 0x1f, 0x1a, 0x20, 0xc0, 0x20, 0x25, 0xd8, 0x35, 0x57, 0x81, 0x17,
	0x2c, 0x1f, 0x92, 0x51, 0x9a, 0x9b, 0x54, 0x1c, 0x5a, 0x20, 0xa4, 0xc5, 0x25, 0x88, 0xa2, 0xce,
	0xc5, 0x3f, 0xdc, 0x4f, 0x80, 0x51, 0x4a, 0xb8, 0x6b, 0xae, 0x02, 0x3f, 0x92, 0x4a, 0x97, 0xfc,
	0xf2, 0x3c, 0x29, 0x96, 0x8e, 0xc5, 0x72, 0x0c, 0x4e, 0x2e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc,
	0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab,
	0x8f, 0x1e, 0x68, 0x15, 0x70, 0x56, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x34, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0xd6, 0x6c, 0x53, 0x5e, 0x01, 0x00, 0x00,
}

func (m *Reaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Reaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Emojis) > 0 {
		dAtA2 := make([]byte, len(m.Emojis)*10)
		var j1 int
		for _, num := range m.Emojis {
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
		i = encodeVarintReaction(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintReaction(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReaction(dAtA []byte, offset int, v uint64) int {
	offset -= sovReaction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Reaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovReaction(uint64(l))
	}
	if len(m.Emojis) > 0 {
		l = 0
		for _, e := range m.Emojis {
			l += sovReaction(uint64(e))
		}
		n += 1 + sovReaction(uint64(l)) + l
	}
	return n
}

func sovReaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReaction(x uint64) (n int) {
	return sovReaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Reaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReaction
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
			return fmt.Errorf("proto: Reaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaction
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
				return ErrInvalidLengthReaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v Emoji
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowReaction
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Emoji(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Emojis = append(m.Emojis, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowReaction
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
					return ErrInvalidLengthReaction
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthReaction
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Emojis) == 0 {
					m.Emojis = make([]Emoji, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Emoji
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowReaction
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Emoji(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Emojis = append(m.Emojis, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Emojis", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReaction
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
func skipReaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReaction
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
					return 0, ErrIntOverflowReaction
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
					return 0, ErrIntOverflowReaction
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
				return 0, ErrInvalidLengthReaction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReaction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReaction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReaction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReaction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReaction = fmt.Errorf("proto: unexpected end of group")
)
