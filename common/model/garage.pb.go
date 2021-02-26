// Code generated by protoc-gen-go. DO NOT EDIT.
// source: garage.proto

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

type GarageCoordinate struct {
	Latitude             float32  `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GarageCoordinate) Reset()         { *m = GarageCoordinate{} }
func (m *GarageCoordinate) String() string { return proto.CompactTextString(m) }
func (*GarageCoordinate) ProtoMessage()    {}
func (*GarageCoordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{0}
}

func (m *GarageCoordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageCoordinate.Unmarshal(m, b)
}
func (m *GarageCoordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageCoordinate.Marshal(b, m, deterministic)
}
func (m *GarageCoordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageCoordinate.Merge(m, src)
}
func (m *GarageCoordinate) XXX_Size() int {
	return xxx_messageInfo_GarageCoordinate.Size(m)
}
func (m *GarageCoordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageCoordinate.DiscardUnknown(m)
}

var xxx_messageInfo_GarageCoordinate proto.InternalMessageInfo

func (m *GarageCoordinate) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *GarageCoordinate) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type Garage struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Coordinate           *GarageCoordinate `protobuf:"bytes,3,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Garage) Reset()         { *m = Garage{} }
func (m *Garage) String() string { return proto.CompactTextString(m) }
func (*Garage) ProtoMessage()    {}
func (*Garage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{1}
}

func (m *Garage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Garage.Unmarshal(m, b)
}
func (m *Garage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Garage.Marshal(b, m, deterministic)
}
func (m *Garage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Garage.Merge(m, src)
}
func (m *Garage) XXX_Size() int {
	return xxx_messageInfo_Garage.Size(m)
}
func (m *Garage) XXX_DiscardUnknown() {
	xxx_messageInfo_Garage.DiscardUnknown(m)
}

var xxx_messageInfo_Garage proto.InternalMessageInfo

func (m *Garage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Garage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Garage) GetCoordinate() *GarageCoordinate {
	if m != nil {
		return m.Coordinate
	}
	return nil
}

type GarageList struct {
	List                 []*Garage `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GarageList) Reset()         { *m = GarageList{} }
func (m *GarageList) String() string { return proto.CompactTextString(m) }
func (*GarageList) ProtoMessage()    {}
func (*GarageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{2}
}

func (m *GarageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageList.Unmarshal(m, b)
}
func (m *GarageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageList.Marshal(b, m, deterministic)
}
func (m *GarageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageList.Merge(m, src)
}
func (m *GarageList) XXX_Size() int {
	return xxx_messageInfo_GarageList.Size(m)
}
func (m *GarageList) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageList.DiscardUnknown(m)
}

var xxx_messageInfo_GarageList proto.InternalMessageInfo

func (m *GarageList) GetList() []*Garage {
	if m != nil {
		return m.List
	}
	return nil
}

type GarageListByUser struct {
	List                 map[string]*GarageList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GarageListByUser) Reset()         { *m = GarageListByUser{} }
func (m *GarageListByUser) String() string { return proto.CompactTextString(m) }
func (*GarageListByUser) ProtoMessage()    {}
func (*GarageListByUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{3}
}

func (m *GarageListByUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageListByUser.Unmarshal(m, b)
}
func (m *GarageListByUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageListByUser.Marshal(b, m, deterministic)
}
func (m *GarageListByUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageListByUser.Merge(m, src)
}
func (m *GarageListByUser) XXX_Size() int {
	return xxx_messageInfo_GarageListByUser.Size(m)
}
func (m *GarageListByUser) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageListByUser.DiscardUnknown(m)
}

var xxx_messageInfo_GarageListByUser proto.InternalMessageInfo

func (m *GarageListByUser) GetList() map[string]*GarageList {
	if m != nil {
		return m.List
	}
	return nil
}

type GarageUserId struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GarageUserId) Reset()         { *m = GarageUserId{} }
func (m *GarageUserId) String() string { return proto.CompactTextString(m) }
func (*GarageUserId) ProtoMessage()    {}
func (*GarageUserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{4}
}

func (m *GarageUserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageUserId.Unmarshal(m, b)
}
func (m *GarageUserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageUserId.Marshal(b, m, deterministic)
}
func (m *GarageUserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageUserId.Merge(m, src)
}
func (m *GarageUserId) XXX_Size() int {
	return xxx_messageInfo_GarageUserId.Size(m)
}
func (m *GarageUserId) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageUserId.DiscardUnknown(m)
}

var xxx_messageInfo_GarageUserId proto.InternalMessageInfo

func (m *GarageUserId) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GarageAndUserId struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Garage               *Garage  `protobuf:"bytes,2,opt,name=garage,proto3" json:"garage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GarageAndUserId) Reset()         { *m = GarageAndUserId{} }
func (m *GarageAndUserId) String() string { return proto.CompactTextString(m) }
func (*GarageAndUserId) ProtoMessage()    {}
func (*GarageAndUserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_d35248283e3db3c1, []int{5}
}

func (m *GarageAndUserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageAndUserId.Unmarshal(m, b)
}
func (m *GarageAndUserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageAndUserId.Marshal(b, m, deterministic)
}
func (m *GarageAndUserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageAndUserId.Merge(m, src)
}
func (m *GarageAndUserId) XXX_Size() int {
	return xxx_messageInfo_GarageAndUserId.Size(m)
}
func (m *GarageAndUserId) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageAndUserId.DiscardUnknown(m)
}

var xxx_messageInfo_GarageAndUserId proto.InternalMessageInfo

func (m *GarageAndUserId) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GarageAndUserId) GetGarage() *Garage {
	if m != nil {
		return m.Garage
	}
	return nil
}

func init() {
	proto.RegisterType((*GarageCoordinate)(nil), "model.GarageCoordinate")
	proto.RegisterType((*Garage)(nil), "model.Garage")
	proto.RegisterType((*GarageList)(nil), "model.GarageList")
	proto.RegisterType((*GarageListByUser)(nil), "model.GarageListByUser")
	proto.RegisterMapType((map[string]*GarageList)(nil), "model.GarageListByUser.ListEntry")
	proto.RegisterType((*GarageUserId)(nil), "model.GarageUserId")
	proto.RegisterType((*GarageAndUserId)(nil), "model.GarageAndUserId")
}

func init() { proto.RegisterFile("garage.proto", fileDescriptor_d35248283e3db3c1) }

var fileDescriptor_d35248283e3db3c1 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x69, 0x0b, 0xc5, 0x0e, 0xa8, 0x38, 0x26, 0x40, 0xaa, 0x07, 0xd8, 0xc4, 0xc0, 0xa9,
	0x98, 0x1a, 0x83, 0xf1, 0x86, 0x86, 0x18, 0x0d, 0x17, 0x9b, 0x78, 0x36, 0xc5, 0x5d, 0x9b, 0xc6,
	0xd2, 0x25, 0xed, 0xd6, 0xa4, 0x0f, 0xe2, 0xfb, 0x9a, 0xee, 0x42, 0xb1, 0x70, 0xf0, 0xb6, 0x3b,
	0xbf, 0xef, 0x9b, 0x7f, 0x19, 0x68, 0x07, 0x7e, 0xe2, 0x07, 0xcc, 0x59, 0x27, 0x5c, 0x70, 0x6c,
	0xac, 0x38, 0x65, 0x91, 0x7d, 0x11, 0x70, 0x1e, 0x44, 0x6c, 0x22, 0x83, 0xcb, 0xec, 0x73, 0xc2,
	0x56, 0x6b, 0x91, 0x2b, 0x0d, 0x59, 0x40, 0xe7, 0x49, 0x7a, 0x1e, 0x39, 0x4f, 0x68, 0x18, 0xfb,
	0x82, 0xa1, 0x0d, 0x47, 0x91, 0x2f, 0x42, 0x91, 0x51, 0xd6, 0xd7, 0x06, 0xda, 0x58, 0xf7, 0xca,
	0x3f, 0x5e, 0x82, 0x15, 0xf1, 0x38, 0x50, 0x50, 0x97, 0x70, 0x17, 0x20, 0x0c, 0x4c, 0x95, 0x0d,
	0x4f, 0x40, 0x0f, 0xa9, 0x74, 0x5b, 0x9e, 0x1e, 0x52, 0x44, 0xa8, 0xc7, 0xfe, 0x4a, 0x59, 0x2c,
	0x4f, 0xbe, 0x71, 0x0a, 0xf0, 0x51, 0x56, 0xed, 0x1b, 0x03, 0x6d, 0xdc, 0x72, 0x7b, 0x8e, 0x6c,
	0xda, 0xd9, 0x6f, 0xca, 0xfb, 0x23, 0x25, 0x13, 0x00, 0xc5, 0x17, 0x61, 0x2a, 0x70, 0x08, 0xf5,
	0x28, 0x4c, 0x45, 0x5f, 0x1b, 0x18, 0xe3, 0x96, 0x7b, 0x5c, 0x49, 0xe0, 0x49, 0x44, 0x7e, 0xb4,
	0xed, 0x98, 0x85, 0xe3, 0x21, 0x7f, 0x4b, 0x59, 0x82, 0xb7, 0x15, 0xdf, 0xb0, 0xe2, 0xdb, 0xc9,
	0x9c, 0xe2, 0x39, 0x8f, 0x45, 0x92, 0xab, 0x5c, 0xf6, 0x0b, 0x58, 0x65, 0x08, 0x3b, 0x60, 0x7c,
	0xb1, 0x7c, 0x33, 0x67, 0xf1, 0xc4, 0x11, 0x34, 0xbe, 0xfd, 0x28, 0x53, 0x93, 0xb6, 0xdc, 0xb3,
	0x83, 0xb4, 0x9e, 0xe2, 0xf7, 0xfa, 0x9d, 0x46, 0x46, 0xd0, 0x56, 0xa0, 0xa8, 0xf4, 0x4c, 0xb1,
	0x07, 0xcd, 0x2c, 0x65, 0xc9, 0x7b, 0xb9, 0x3a, 0x33, 0x93, 0x80, 0xbc, 0xc2, 0xa9, 0x12, 0xce,
	0x62, 0xfa, 0x8f, 0x16, 0xaf, 0xc0, 0x54, 0x67, 0xb0, 0x69, 0x61, 0x6f, 0x23, 0x1b, 0xe8, 0x0a,
	0x68, 0xaa, 0x48, 0x8a, 0xd7, 0x50, 0x97, 0x9b, 0x3c, 0xaf, 0x28, 0x55, 0x1d, 0xfb, 0x70, 0x02,
	0x52, 0xc3, 0x29, 0x18, 0x33, 0x4a, 0xb1, 0x5b, 0x61, 0x65, 0x6f, 0x76, 0xd7, 0x51, 0x37, 0xe7,
	0x6c, 0x6f, 0xce, 0x99, 0x17, 0x37, 0x47, 0x6a, 0x4b, 0x53, 0x46, 0x6e, 0x7e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xda, 0x61, 0x00, 0xfe, 0xaa, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GaragesClient is the client API for Garages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GaragesClient interface {
	List(ctx context.Context, in *GarageUserId, opts ...grpc.CallOption) (*GarageList, error)
	Add(ctx context.Context, in *GarageAndUserId, opts ...grpc.CallOption) (*empty.Empty, error)
}

type garagesClient struct {
	cc *grpc.ClientConn
}

func NewGaragesClient(cc *grpc.ClientConn) GaragesClient {
	return &garagesClient{cc}
}

func (c *garagesClient) List(ctx context.Context, in *GarageUserId, opts ...grpc.CallOption) (*GarageList, error) {
	out := new(GarageList)
	err := c.cc.Invoke(ctx, "/model.Garages/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *garagesClient) Add(ctx context.Context, in *GarageAndUserId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.Garages/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GaragesServer is the server API for Garages service.
type GaragesServer interface {
	List(context.Context, *GarageUserId) (*GarageList, error)
	Add(context.Context, *GarageAndUserId) (*empty.Empty, error)
}

// UnimplementedGaragesServer can be embedded to have forward compatible implementations.
type UnimplementedGaragesServer struct {
}

func (*UnimplementedGaragesServer) List(ctx context.Context, req *GarageUserId) (*GarageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedGaragesServer) Add(ctx context.Context, req *GarageAndUserId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterGaragesServer(s *grpc.Server, srv GaragesServer) {
	s.RegisterService(&_Garages_serviceDesc, srv)
}

func _Garages_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GarageUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaragesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Garages/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaragesServer).List(ctx, req.(*GarageUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Garages_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GarageAndUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaragesServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Garages/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaragesServer).Add(ctx, req.(*GarageAndUserId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Garages_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Garages",
	HandlerType: (*GaragesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Garages_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Garages_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "garage.proto",
}
