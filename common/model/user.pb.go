// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type UserGender int32

const (
	UserGender_UNDEFINED UserGender = 0
	UserGender_MALE      UserGender = 1
	UserGender_FEMALE    UserGender = 2
)

var UserGender_name = map[int32]string{
	0: "UNDEFINED",
	1: "MALE",
	2: "FEMALE",
}

var UserGender_value = map[string]int32{
	"UNDEFINED": 0,
	"MALE":      1,
	"FEMALE":    2,
}

func (x UserGender) String() string {
	return proto.EnumName(UserGender_name, int32(x))
}

func (UserGender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type User struct {
	ID                   string     `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Gender               UserGender `protobuf:"varint,4,opt,name=gender,proto3,enum=model.UserGender" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetGender() UserGender {
	if m != nil {
		return m.Gender
	}
	return UserGender_UNDEFINED
}

type UserList struct {
	List                 []*User  `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("model.UserGender", UserGender_name, UserGender_value)
	proto.RegisterType((*User)(nil), "model.User")
	proto.RegisterType((*UserList)(nil), "model.UserList")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x6b, 0x83, 0x40,
	0x10, 0xc5, 0xd5, 0x6c, 0xc4, 0x4c, 0x68, 0x9a, 0xce, 0xa1, 0x88, 0x3d, 0x54, 0x3c, 0xd9, 0x16,
	0x36, 0x68, 0x3f, 0x41, 0x41, 0x53, 0x84, 0x34, 0x07, 0x21, 0x1f, 0x20, 0xc1, 0xad, 0x08, 0x9a,
	0xb5, 0xbb, 0x86, 0xd2, 0x6f, 0x5f, 0x76, 0xd2, 0x3f, 0x5e, 0x72, 0x9b, 0x7d, 0xf3, 0x7b, 0x6f,
	0x87, 0x07, 0x70, 0xd2, 0x42, 0xf1, 0x5e, 0xc9, 0x41, 0xe2, 0xb4, 0x93, 0x95, 0x68, 0x83, 0xbb,
	0x5a, 0xca, 0xba, 0x15, 0x2b, 0x12, 0x0f, 0xa7, 0xf7, 0x95, 0xe8, 0xfa, 0xe1, 0xeb, 0xcc, 0x44,
	0x1f, 0xc0, 0x76, 0x5a, 0x28, 0x5c, 0x80, 0x53, 0x64, 0xbe, 0x1d, 0xda, 0xf1, 0xac, 0x74, 0x8a,
	0x0c, 0x11, 0xd8, 0x71, 0xdf, 0x09, 0xdf, 0x21, 0x85, 0x66, 0x0c, 0xc0, 0xeb, 0xf7, 0x5a, 0x7f,
	0x4a, 0x55, 0xf9, 0x13, 0xd2, 0xff, 0xde, 0xf8, 0x00, 0x6e, 0x2d, 0x8e, 0x95, 0x50, 0x3e, 0x0b,
	0xed, 0x78, 0x91, 0xde, 0x70, 0xfa, 0x9c, 0x9b, 0xf0, 0x57, 0x5a, 0x94, 0x3f, 0x40, 0xf4, 0x04,
	0x9e, 0x51, 0x37, 0x8d, 0x1e, 0xf0, 0x1e, 0x58, 0xdb, 0xe8, 0xc1, 0xb7, 0xc3, 0x49, 0x3c, 0x4f,
	0xe7, 0x23, 0x53, 0x49, 0x8b, 0xc7, 0x04, 0xe0, 0x3f, 0x02, 0xaf, 0x60, 0xb6, 0xdb, 0x66, 0xf9,
	0xba, 0xd8, 0xe6, 0xd9, 0xd2, 0x42, 0x0f, 0xd8, 0xdb, 0xcb, 0x26, 0x5f, 0xda, 0x08, 0xe0, 0xae,
	0x73, 0x9a, 0x9d, 0xb4, 0x83, 0xa9, 0xb1, 0x68, 0x4c, 0xc0, 0x2b, 0x45, 0xdd, 0xe8, 0x41, 0x28,
	0x1c, 0x47, 0x07, 0xb7, 0xfc, 0x5c, 0x09, 0xff, 0xad, 0x84, 0xe7, 0xa6, 0x92, 0xc8, 0xc2, 0x04,
	0x18, 0xdd, 0x75, 0x81, 0x08, 0xae, 0x47, 0x31, 0x06, 0x8c, 0xac, 0x83, 0x4b, 0xc8, 0xf3, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x2b, 0xff, 0xf0, 0x7a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserList, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.Users/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/model.Users/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Register(context.Context, *User) (*empty.Empty, error)
	List(context.Context, *empty.Empty) (*UserList, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) Register(ctx context.Context, req *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUsersServer) List(ctx context.Context, req *empty.Empty) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Users_Register_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Users_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
