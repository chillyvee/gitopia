// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/user.proto

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

type User struct {
	Creator        string              `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id             uint64              `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name           string              `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Username       string              `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	UsernameGithub string              `protobuf:"bytes,5,opt,name=usernameGithub,proto3" json:"usernameGithub,omitempty"`
	AvatarUrl      string              `protobuf:"bytes,6,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	Followers      []string            `protobuf:"bytes,7,rep,name=followers,proto3" json:"followers,omitempty"`
	Following      []string            `protobuf:"bytes,8,rep,name=following,proto3" json:"following,omitempty"`
	Repositories   []*UserRepository   `protobuf:"bytes,9,rep,name=repositories,proto3" json:"repositories,omitempty"`
	Organizations  []*UserOrganization `protobuf:"bytes,10,rep,name=organizations,proto3" json:"organizations,omitempty"`
	StarredRepos   []uint64            `protobuf:"varint,11,rep,packed,name=starred_repos,json=starredRepos,proto3" json:"starred_repos,omitempty"`
	Subscriptions  string              `protobuf:"bytes,12,opt,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	Email          string              `protobuf:"bytes,13,opt,name=email,proto3" json:"email,omitempty"`
	Bio            string              `protobuf:"bytes,14,opt,name=bio,proto3" json:"bio,omitempty"`
	CreatedAt      int64               `protobuf:"varint,15,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      int64               `protobuf:"varint,16,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Extensions     string              `protobuf:"bytes,17,opt,name=extensions,proto3" json:"extensions,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7f4b301dc3d162, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetUsernameGithub() string {
	if m != nil {
		return m.UsernameGithub
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *User) GetFollowers() []string {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *User) GetFollowing() []string {
	if m != nil {
		return m.Following
	}
	return nil
}

func (m *User) GetRepositories() []*UserRepository {
	if m != nil {
		return m.Repositories
	}
	return nil
}

func (m *User) GetOrganizations() []*UserOrganization {
	if m != nil {
		return m.Organizations
	}
	return nil
}

func (m *User) GetStarredRepos() []uint64 {
	if m != nil {
		return m.StarredRepos
	}
	return nil
}

func (m *User) GetSubscriptions() string {
	if m != nil {
		return m.Subscriptions
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *User) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *User) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *User) GetExtensions() string {
	if m != nil {
		return m.Extensions
	}
	return ""
}

type UserRepository struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *UserRepository) Reset()         { *m = UserRepository{} }
func (m *UserRepository) String() string { return proto.CompactTextString(m) }
func (*UserRepository) ProtoMessage()    {}
func (*UserRepository) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7f4b301dc3d162, []int{1}
}
func (m *UserRepository) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserRepository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserRepository.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserRepository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRepository.Merge(m, src)
}
func (m *UserRepository) XXX_Size() int {
	return m.Size()
}
func (m *UserRepository) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRepository.DiscardUnknown(m)
}

var xxx_messageInfo_UserRepository proto.InternalMessageInfo

func (m *UserRepository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRepository) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserOrganization struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *UserOrganization) Reset()         { *m = UserOrganization{} }
func (m *UserOrganization) String() string { return proto.CompactTextString(m) }
func (*UserOrganization) ProtoMessage()    {}
func (*UserOrganization) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7f4b301dc3d162, []int{2}
}
func (m *UserOrganization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserOrganization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserOrganization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserOrganization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOrganization.Merge(m, src)
}
func (m *UserOrganization) XXX_Size() int {
	return m.Size()
}
func (m *UserOrganization) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOrganization.DiscardUnknown(m)
}

var xxx_messageInfo_UserOrganization proto.InternalMessageInfo

func (m *UserOrganization) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserOrganization) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "gitopia.gitopia.gitopia.User")
	proto.RegisterType((*UserRepository)(nil), "gitopia.gitopia.gitopia.UserRepository")
	proto.RegisterType((*UserOrganization)(nil), "gitopia.gitopia.gitopia.UserOrganization")
}

func init() { proto.RegisterFile("gitopia/user.proto", fileDescriptor_bf7f4b301dc3d162) }

var fileDescriptor_bf7f4b301dc3d162 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xeb, 0x26, 0xfb, 0xa7, 0xb3, 0x6d, 0x29, 0xa3, 0x95, 0xb0, 0x56, 0x28, 0x8a, 0x0a,
	0x82, 0xc0, 0x21, 0x95, 0x00, 0x71, 0x07, 0x21, 0x71, 0xe0, 0xb0, 0x52, 0xa4, 0xbd, 0x70, 0x41,
	0x6e, 0x63, 0x82, 0xa5, 0x36, 0x8e, 0x6c, 0x17, 0x76, 0x79, 0x0a, 0x1e, 0x8b, 0x1b, 0x7b, 0xe4,
	0x88, 0xda, 0x17, 0x41, 0x9e, 0x36, 0x69, 0x1b, 0x09, 0xf6, 0x94, 0xf1, 0xef, 0x9b, 0xcf, 0xb6,
	0xbe, 0x8c, 0x01, 0x0b, 0xe5, 0x74, 0xa5, 0xc4, 0x64, 0x69, 0xa5, 0x49, 0x2b, 0xa3, 0x9d, 0xc6,
	0x07, 0x5b, 0x96, 0xb6, 0xbe, 0x17, 0xe7, 0x85, 0x2e, 0x34, 0xf5, 0x4c, 0x7c, 0xb5, 0x69, 0x1f,
	0xff, 0x0a, 0x21, 0xbc, 0xb2, 0xd2, 0x20, 0x87, 0x93, 0x99, 0x91, 0xc2, 0x69, 0xc3, 0x59, 0xcc,
	0x92, 0x5e, 0x56, 0x2f, 0x71, 0x08, 0x5d, 0x95, 0xf3, 0x6e, 0xcc, 0x92, 0x30, 0xeb, 0xaa, 0x1c,
	0x11, 0xc2, 0x52, 0x2c, 0x24, 0x0f, 0xa8, 0x8d, 0x6a, 0xbc, 0x80, 0x53, 0x7f, 0x07, 0xe2, 0x21,
	0xf1, 0x66, 0x8d, 0x4f, 0x60, 0x58, 0xd7, 0xef, 0x95, 0xfb, 0xb2, 0x9c, 0xf2, 0x23, 0xea, 0x68,
	0x51, 0x7c, 0x08, 0x3d, 0xf1, 0x55, 0x38, 0x61, 0xae, 0xcc, 0x9c, 0x1f, 0x53, 0xcb, 0x0e, 0x78,
	0xf5, 0xb3, 0x9e, 0xcf, 0xf5, 0x37, 0x69, 0x2c, 0x3f, 0x89, 0x03, 0xaf, 0x36, 0x60, 0xa7, 0xaa,
	0xb2, 0xe0, 0xa7, 0xfb, 0xaa, 0x2a, 0x0b, 0xfc, 0x00, 0x7d, 0x23, 0x2b, 0x6d, 0x95, 0xd3, 0x46,
	0x49, 0xcb, 0x7b, 0x71, 0x90, 0x9c, 0xbd, 0x78, 0x9a, 0xfe, 0x23, 0xaa, 0xd4, 0x07, 0x92, 0xd5,
	0x86, 0x9b, 0xec, 0xc0, 0x8c, 0x97, 0x30, 0xd0, 0xa6, 0x10, 0xa5, 0xfa, 0x2e, 0x9c, 0xd2, 0xa5,
	0xe5, 0x40, 0xbb, 0x3d, 0xfb, 0xef, 0x6e, 0x97, 0x7b, 0x8e, 0xec, 0xd0, 0x8f, 0x8f, 0x60, 0x60,
	0x9d, 0x30, 0x46, 0xe6, 0x9f, 0xe8, 0x20, 0x7e, 0x16, 0x07, 0x49, 0x98, 0xf5, 0xb7, 0x90, 0x2e,
	0x82, 0x8f, 0x61, 0x60, 0x97, 0x53, 0x3b, 0x33, 0xaa, 0xda, 0x9c, 0xda, 0xa7, 0x80, 0x0e, 0x21,
	0x9e, 0xc3, 0x91, 0x5c, 0x08, 0x35, 0xe7, 0x03, 0x52, 0x37, 0x0b, 0x1c, 0x41, 0x30, 0x55, 0x9a,
	0x0f, 0x89, 0xf9, 0xd2, 0xc7, 0x45, 0x7f, 0x57, 0xe6, 0x6f, 0x1c, 0xbf, 0x17, 0xb3, 0x24, 0xc8,
	0x76, 0xc0, 0xab, 0xcb, 0x2a, 0xdf, 0xaa, 0xa3, 0x8d, 0xda, 0x00, 0x8c, 0x00, 0xe4, 0xb5, 0x93,
	0xa5, 0xa5, 0x6b, 0xdc, 0xa7, 0x4d, 0xf7, 0xc8, 0xf8, 0x15, 0x0c, 0x0f, 0xf3, 0x6b, 0x06, 0x86,
	0xed, 0x0d, 0x4c, 0x6b, 0xa8, 0xc6, 0xaf, 0x61, 0xd4, 0xce, 0xe9, 0x0e, 0x5f, 0xcf, 0xfb, 0xde,
	0xbe, 0xfb, 0xb9, 0x8a, 0xd8, 0xed, 0x2a, 0x62, 0x7f, 0x56, 0x11, 0xfb, 0xb1, 0x8e, 0x3a, 0xb7,
	0xeb, 0xa8, 0xf3, 0x7b, 0x1d, 0x75, 0x3e, 0x3e, 0x2f, 0x68, 0xac, 0xd2, 0x99, 0x5e, 0x4c, 0xea,
	0x77, 0x52, 0x7f, 0xaf, 0x9b, 0xca, 0xdd, 0x54, 0xd2, 0x4e, 0x8f, 0xe9, 0x31, 0xbc, 0xfc, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x39, 0x26, 0x41, 0xf4, 0x51, 0x03, 0x00, 0x00,
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Extensions) > 0 {
		i -= len(m.Extensions)
		copy(dAtA[i:], m.Extensions)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Extensions)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.CreatedAt != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x78
	}
	if len(m.Bio) > 0 {
		i -= len(m.Bio)
		copy(dAtA[i:], m.Bio)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Bio)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Subscriptions) > 0 {
		i -= len(m.Subscriptions)
		copy(dAtA[i:], m.Subscriptions)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Subscriptions)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.StarredRepos) > 0 {
		dAtA2 := make([]byte, len(m.StarredRepos)*10)
		var j1 int
		for _, num := range m.StarredRepos {
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
		i = encodeVarintUser(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Organizations) > 0 {
		for iNdEx := len(m.Organizations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Organizations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUser(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Repositories) > 0 {
		for iNdEx := len(m.Repositories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Repositories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUser(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Following) > 0 {
		for iNdEx := len(m.Following) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Following[iNdEx])
			copy(dAtA[i:], m.Following[iNdEx])
			i = encodeVarintUser(dAtA, i, uint64(len(m.Following[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Followers) > 0 {
		for iNdEx := len(m.Followers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Followers[iNdEx])
			copy(dAtA[i:], m.Followers[iNdEx])
			i = encodeVarintUser(dAtA, i, uint64(len(m.Followers[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AvatarUrl) > 0 {
		i -= len(m.AvatarUrl)
		copy(dAtA[i:], m.AvatarUrl)
		i = encodeVarintUser(dAtA, i, uint64(len(m.AvatarUrl)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.UsernameGithub) > 0 {
		i -= len(m.UsernameGithub)
		copy(dAtA[i:], m.UsernameGithub)
		i = encodeVarintUser(dAtA, i, uint64(len(m.UsernameGithub)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserRepository) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserRepository) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserRepository) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserOrganization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserOrganization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserOrganization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovUser(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.UsernameGithub)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.AvatarUrl)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if len(m.Followers) > 0 {
		for _, s := range m.Followers {
			l = len(s)
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if len(m.Following) > 0 {
		for _, s := range m.Following {
			l = len(s)
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if len(m.Repositories) > 0 {
		for _, e := range m.Repositories {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if len(m.Organizations) > 0 {
		for _, e := range m.Organizations {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	if len(m.StarredRepos) > 0 {
		l = 0
		for _, e := range m.StarredRepos {
			l += sovUser(uint64(e))
		}
		n += 1 + sovUser(uint64(l)) + l
	}
	l = len(m.Subscriptions)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Bio)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovUser(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 2 + sovUser(uint64(m.UpdatedAt))
	}
	l = len(m.Extensions)
	if l > 0 {
		n += 2 + l + sovUser(uint64(l))
	}
	return n
}

func (m *UserRepository) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovUser(uint64(m.Id))
	}
	return n
}

func (m *UserOrganization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
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
					return ErrIntOverflowUser
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsernameGithub", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UsernameGithub = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvatarUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AvatarUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Followers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Followers = append(m.Followers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Following", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Following = append(m.Following, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repositories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Repositories = append(m.Repositories, &UserRepository{})
			if err := m.Repositories[len(m.Repositories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organizations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Organizations = append(m.Organizations, &UserOrganization{})
			if err := m.Organizations[len(m.Organizations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUser
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
				m.StarredRepos = append(m.StarredRepos, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowUser
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
					return ErrInvalidLengthUser
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthUser
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
				if elementCount != 0 && len(m.StarredRepos) == 0 {
					m.StarredRepos = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowUser
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
					m.StarredRepos = append(m.StarredRepos, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field StarredRepos", wireType)
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriptions = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extensions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extensions = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func (m *UserRepository) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserRepository: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserRepository: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func (m *UserOrganization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserOrganization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserOrganization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
