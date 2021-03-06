// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coprocess_object.proto

package coprocess

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Object struct {
	HookType             HookType           `protobuf:"varint,1,opt,name=hook_type,json=hookType,proto3,enum=coprocess.HookType" json:"hook_type,omitempty"`
	HookName             string             `protobuf:"bytes,2,opt,name=hook_name,json=hookName,proto3" json:"hook_name,omitempty"`
	Request              *MiniRequestObject `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	Session              *SessionState      `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	Metadata             map[string]string  `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 map[string]string  `protobuf:"bytes,6,rep,name=spec,proto3" json:"spec,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_72698a2223f86099, []int{0}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetHookType() HookType {
	if m != nil {
		return m.HookType
	}
	return HookType_Unknown
}

func (m *Object) GetHookName() string {
	if m != nil {
		return m.HookName
	}
	return ""
}

func (m *Object) GetRequest() *MiniRequestObject {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Object) GetSession() *SessionState {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Object) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSpec() map[string]string {
	if m != nil {
		return m.Spec
	}
	return nil
}

type Event struct {
	Payload              string   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_72698a2223f86099, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type EventReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventReply) Reset()         { *m = EventReply{} }
func (m *EventReply) String() string { return proto.CompactTextString(m) }
func (*EventReply) ProtoMessage()    {}
func (*EventReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_72698a2223f86099, []int{2}
}

func (m *EventReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventReply.Unmarshal(m, b)
}
func (m *EventReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventReply.Marshal(b, m, deterministic)
}
func (m *EventReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventReply.Merge(m, src)
}
func (m *EventReply) XXX_Size() int {
	return xxx_messageInfo_EventReply.Size(m)
}
func (m *EventReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EventReply.DiscardUnknown(m)
}

var xxx_messageInfo_EventReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Object)(nil), "coprocess.Object")
	proto.RegisterMapType((map[string]string)(nil), "coprocess.Object.MetadataEntry")
	proto.RegisterMapType((map[string]string)(nil), "coprocess.Object.SpecEntry")
	proto.RegisterType((*Event)(nil), "coprocess.Event")
	proto.RegisterType((*EventReply)(nil), "coprocess.EventReply")
}

func init() { proto.RegisterFile("coprocess_object.proto", fileDescriptor_72698a2223f86099) }

var fileDescriptor_72698a2223f86099 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0xcf, 0x93, 0x40,
	0x10, 0xc6, 0x4b, 0xe9, 0x3f, 0xa6, 0xd6, 0xd4, 0xf5, 0x1f, 0xa1, 0x1a, 0x11, 0x2f, 0x9c, 0x50,
	0x31, 0x51, 0xd3, 0x5e, 0x6d, 0xe2, 0xa5, 0x9a, 0x50, 0xef, 0x64, 0x4b, 0x27, 0x29, 0xb6, 0xb0,
	0x2b, 0xbb, 0x6d, 0x42, 0xe2, 0xe7, 0x79, 0x3f, 0xe7, 0x9b, 0xee, 0x02, 0xa5, 0x6f, 0x4f, 0xef,
	0x6d, 0xe7, 0x99, 0xe7, 0x37, 0x0f, 0x33, 0x01, 0x5e, 0x25, 0x8c, 0x17, 0x2c, 0x41, 0x21, 0x62,
	0xb6, 0xf9, 0x8b, 0x89, 0x0c, 0x78, 0xc1, 0x24, 0x23, 0x56, 0xa3, 0x3b, 0x1f, 0x2e, 0x96, 0x2c,
	0xcd, 0xd3, 0xb8, 0xc0, 0x7f, 0x47, 0x14, 0xf2, 0xca, 0xef, 0xbc, 0xbd, 0x98, 0x04, 0x0a, 0x91,
	0xb2, 0x3c, 0x16, 0x92, 0x4a, 0xac, 0xda, 0xad, 0x98, 0x84, 0x65, 0x19, 0xcb, 0xb5, 0xee, 0xdd,
	0x99, 0x30, 0xf8, 0xad, 0xe6, 0x90, 0x4f, 0x60, 0xed, 0x18, 0xdb, 0xc7, 0xb2, 0xe4, 0x68, 0x1b,
	0xae, 0xe1, 0x3f, 0x0d, 0x9f, 0x07, 0x0d, 0x16, 0xfc, 0x64, 0x6c, 0xff, 0xa7, 0xe4, 0x18, 0x8d,
	0x76, 0xd5, 0x8b, 0xcc, 0x2a, 0x22, 0xa7, 0x19, 0xda, 0x5d, 0xd7, 0xf0, 0x2d, 0xdd, 0xfc, 0x45,
	0x33, 0x24, 0x5f, 0x61, 0x58, 0x7d, 0xa8, 0x6d, 0xba, 0x86, 0x3f, 0x0e, 0xdf, 0xb4, 0x86, 0xad,
	0xd2, 0x3c, 0x8d, 0x74, 0x57, 0xa7, 0x47, 0xb5, 0x99, 0x7c, 0x86, 0x61, 0xb5, 0x80, 0xdd, 0x53,
	0xdc, 0xeb, 0x16, 0xb7, 0xd6, 0x9d, 0xf5, 0x79, 0xb3, 0xa8, 0xf6, 0x91, 0x05, 0x8c, 0x32, 0x94,
	0x74, 0x4b, 0x25, 0xb5, 0xfb, 0xae, 0xe9, 0x8f, 0xc3, 0x77, 0x2d, 0x46, 0x07, 0x04, 0xab, 0xca,
	0xb1, 0xcc, 0x65, 0x51, 0x46, 0x0d, 0x40, 0x3e, 0x42, 0x4f, 0x70, 0x4c, 0xec, 0x81, 0x02, 0x67,
	0xb7, 0xe0, 0x9a, 0x63, 0xa2, 0x21, 0x65, 0x74, 0x16, 0x30, 0xb9, 0x9a, 0x45, 0xa6, 0x60, 0xee,
	0xb1, 0x54, 0x27, 0xb3, 0xa2, 0xf3, 0x93, 0xbc, 0x80, 0xfe, 0x89, 0x1e, 0x8e, 0xf5, 0x51, 0x74,
	0x31, 0xef, 0x7e, 0x37, 0x9c, 0x6f, 0x60, 0x35, 0xf3, 0x1e, 0x03, 0x7a, 0xef, 0xa1, 0xbf, 0x3c,
	0x61, 0x2e, 0x89, 0x0d, 0x43, 0x4e, 0xcb, 0x03, 0xa3, 0xdb, 0x0a, 0xac, 0x4b, 0xef, 0x09, 0x80,
	0xb2, 0x44, 0xc8, 0x0f, 0x65, 0xf8, 0x1f, 0xe0, 0x47, 0x2a, 0x38, 0x95, 0xc9, 0x0e, 0x0b, 0x12,
	0xc2, 0xa8, 0xae, 0xc8, 0xb3, 0x9b, 0x1d, 0x9d, 0x5b, 0xc9, 0xeb, 0x90, 0x39, 0x4c, 0x6a, 0x46,
	0x47, 0x4f, 0x5b, 0x2e, 0xa5, 0x38, 0x2f, 0x1f, 0x2a, 0x2a, 0xdb, 0xeb, 0x6c, 0x06, 0xea, 0xf7,
	0xfa, 0x72, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xaf, 0x58, 0x65, 0xdf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DispatcherClient is the client API for Dispatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DispatcherClient interface {
	Dispatch(ctx context.Context, in *Object, opts ...grpc.CallOption) (*Object, error)
	DispatchEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventReply, error)
}

type dispatcherClient struct {
	cc *grpc.ClientConn
}

func NewDispatcherClient(cc *grpc.ClientConn) DispatcherClient {
	return &dispatcherClient{cc}
}

func (c *dispatcherClient) Dispatch(ctx context.Context, in *Object, opts ...grpc.CallOption) (*Object, error) {
	out := new(Object)
	err := c.cc.Invoke(ctx, "/coprocess.Dispatcher/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) DispatchEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventReply, error) {
	out := new(EventReply)
	err := c.cc.Invoke(ctx, "/coprocess.Dispatcher/DispatchEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServer is the server API for Dispatcher service.
type DispatcherServer interface {
	Dispatch(context.Context, *Object) (*Object, error)
	DispatchEvent(context.Context, *Event) (*EventReply, error)
}

// UnimplementedDispatcherServer can be embedded to have forward compatible implementations.
type UnimplementedDispatcherServer struct {
}

func (*UnimplementedDispatcherServer) Dispatch(ctx context.Context, req *Object) (*Object, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}
func (*UnimplementedDispatcherServer) DispatchEvent(ctx context.Context, req *Event) (*EventReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DispatchEvent not implemented")
}

func RegisterDispatcherServer(s *grpc.Server, srv DispatcherServer) {
	s.RegisterService(&_Dispatcher_serviceDesc, srv)
}

func _Dispatcher_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coprocess.Dispatcher/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).Dispatch(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_DispatchEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).DispatchEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coprocess.Dispatcher/DispatchEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).DispatchEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dispatcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coprocess.Dispatcher",
	HandlerType: (*DispatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dispatch",
			Handler:    _Dispatcher_Dispatch_Handler,
		},
		{
			MethodName: "DispatchEvent",
			Handler:    _Dispatcher_DispatchEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coprocess_object.proto",
}
