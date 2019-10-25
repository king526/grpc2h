// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Item struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetItemReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemReq) Reset()         { *m = GetItemReq{} }
func (m *GetItemReq) String() string { return proto.CompactTextString(m) }
func (*GetItemReq) ProtoMessage()    {}
func (*GetItemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{1}
}

func (m *GetItemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemReq.Unmarshal(m, b)
}
func (m *GetItemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemReq.Marshal(b, m, deterministic)
}
func (m *GetItemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemReq.Merge(m, src)
}
func (m *GetItemReq) XXX_Size() int {
	return xxx_messageInfo_GetItemReq.Size(m)
}
func (m *GetItemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemReq proto.InternalMessageInfo

func (m *GetItemReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{2}
}

func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (m *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(m, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Item)(nil), "pb.Item")
	proto.RegisterType((*GetItemReq)(nil), "pb.GetItemReq")
	proto.RegisterType((*Null)(nil), "pb.Null")
}

func init() { proto.RegisterFile("item.proto", fileDescriptor_6007f868cf6553df) }

var fileDescriptor_6007f868cf6553df = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2c, 0x49, 0xcd,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe2, 0x62, 0xf1, 0x2c,
	0x49, 0xcd, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62,
	0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0x64, 0xb8, 0xb8, 0xdc, 0x53, 0x4b, 0x40, 0xca, 0x83, 0x52, 0x0b, 0xd1,
	0x75, 0x28, 0xb1, 0x71, 0xb1, 0xf8, 0x95, 0xe6, 0xe4, 0x18, 0x05, 0x70, 0x71, 0x83, 0x94, 0x04,
	0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x29, 0x72, 0x31, 0xbb, 0xa7, 0x96, 0x08, 0xf1, 0xe9,
	0x15, 0x24, 0xe9, 0x21, 0x74, 0x4b, 0x71, 0x80, 0xf8, 0x20, 0x8e, 0x12, 0x83, 0x90, 0x1c, 0x17,
	0x8b, 0x4f, 0x66, 0x71, 0x89, 0x10, 0x58, 0x0c, 0x64, 0x06, 0xb2, 0xac, 0x01, 0x63, 0x12, 0x1b,
	0xd8, 0xb9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0xdc, 0x14, 0x8a, 0xbc, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemServiceClient interface {
	Get(ctx context.Context, in *GetItemReq, opts ...grpc.CallOption) (*Item, error)
	List(ctx context.Context, in *Null, opts ...grpc.CallOption) (ItemService_ListClient, error)
}

type itemServiceClient struct {
	cc *grpc.ClientConn
}

func NewItemServiceClient(cc *grpc.ClientConn) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) Get(ctx context.Context, in *GetItemReq, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/pb.ItemService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) List(ctx context.Context, in *Null, opts ...grpc.CallOption) (ItemService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ItemService_serviceDesc.Streams[0], "/pb.ItemService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &itemServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemService_ListClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type itemServiceListClient struct {
	grpc.ClientStream
}

func (x *itemServiceListClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ItemServiceServer is the server API for ItemService service.
type ItemServiceServer interface {
	Get(context.Context, *GetItemReq) (*Item, error)
	List(*Null, ItemService_ListServer) error
}

func RegisterItemServiceServer(s *grpc.Server, srv ItemServiceServer) {
	s.RegisterService(&_ItemService_serviceDesc, srv)
}

func _ItemService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ItemService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).Get(ctx, req.(*GetItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Null)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemServiceServer).List(m, &itemServiceListServer{stream})
}

type ItemService_ListServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type itemServiceListServer struct {
	grpc.ServerStream
}

func (x *itemServiceListServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

var _ItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ItemService_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ItemService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "item.proto",
}
