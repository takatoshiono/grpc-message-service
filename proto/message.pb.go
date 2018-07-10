// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/message.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateConversationRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConversationRequest) Reset()         { *m = CreateConversationRequest{} }
func (m *CreateConversationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConversationRequest) ProtoMessage()    {}
func (*CreateConversationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_521c3b15e9248762, []int{0}
}
func (m *CreateConversationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConversationRequest.Unmarshal(m, b)
}
func (m *CreateConversationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConversationRequest.Marshal(b, m, deterministic)
}
func (dst *CreateConversationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConversationRequest.Merge(dst, src)
}
func (m *CreateConversationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConversationRequest.Size(m)
}
func (m *CreateConversationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConversationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConversationRequest proto.InternalMessageInfo

type Conversation struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Conversation) Reset()         { *m = Conversation{} }
func (m *Conversation) String() string { return proto.CompactTextString(m) }
func (*Conversation) ProtoMessage()    {}
func (*Conversation) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_521c3b15e9248762, []int{1}
}
func (m *Conversation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conversation.Unmarshal(m, b)
}
func (m *Conversation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conversation.Marshal(b, m, deterministic)
}
func (dst *Conversation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conversation.Merge(dst, src)
}
func (m *Conversation) XXX_Size() int {
	return xxx_messageInfo_Conversation.Size(m)
}
func (m *Conversation) XXX_DiscardUnknown() {
	xxx_messageInfo_Conversation.DiscardUnknown(m)
}

var xxx_messageInfo_Conversation proto.InternalMessageInfo

func (m *Conversation) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Conversation) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type CreateMessageRequest struct {
	ConversationId       uint32   `protobuf:"varint,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageRequest) Reset()         { *m = CreateMessageRequest{} }
func (m *CreateMessageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMessageRequest) ProtoMessage()    {}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_521c3b15e9248762, []int{2}
}
func (m *CreateMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageRequest.Unmarshal(m, b)
}
func (m *CreateMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageRequest.Marshal(b, m, deterministic)
}
func (dst *CreateMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageRequest.Merge(dst, src)
}
func (m *CreateMessageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMessageRequest.Size(m)
}
func (m *CreateMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageRequest proto.InternalMessageInfo

func (m *CreateMessageRequest) GetConversationId() uint32 {
	if m != nil {
		return m.ConversationId
	}
	return 0
}

func (m *CreateMessageRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Message struct {
	ConversationId       uint32   `protobuf:"varint,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt            int64    `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_521c3b15e9248762, []int{3}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetConversationId() uint32 {
	if m != nil {
		return m.ConversationId
	}
	return 0
}

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Message) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type GetMessagesRequest struct {
	ConversationId       uint32   `protobuf:"varint,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessagesRequest) Reset()         { *m = GetMessagesRequest{} }
func (m *GetMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*GetMessagesRequest) ProtoMessage()    {}
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_521c3b15e9248762, []int{4}
}
func (m *GetMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessagesRequest.Unmarshal(m, b)
}
func (m *GetMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessagesRequest.Marshal(b, m, deterministic)
}
func (dst *GetMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessagesRequest.Merge(dst, src)
}
func (m *GetMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_GetMessagesRequest.Size(m)
}
func (m *GetMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessagesRequest proto.InternalMessageInfo

func (m *GetMessagesRequest) GetConversationId() uint32 {
	if m != nil {
		return m.ConversationId
	}
	return 0
}

type Messages struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Messages) Reset()         { *m = Messages{} }
func (m *Messages) String() string { return proto.CompactTextString(m) }
func (*Messages) ProtoMessage()    {}
func (*Messages) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_521c3b15e9248762, []int{5}
}
func (m *Messages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Messages.Unmarshal(m, b)
}
func (m *Messages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Messages.Marshal(b, m, deterministic)
}
func (dst *Messages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Messages.Merge(dst, src)
}
func (m *Messages) XXX_Size() int {
	return xxx_messageInfo_Messages.Size(m)
}
func (m *Messages) XXX_DiscardUnknown() {
	xxx_messageInfo_Messages.DiscardUnknown(m)
}

var xxx_messageInfo_Messages proto.InternalMessageInfo

func (m *Messages) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateConversationRequest)(nil), "message.CreateConversationRequest")
	proto.RegisterType((*Conversation)(nil), "message.Conversation")
	proto.RegisterType((*CreateMessageRequest)(nil), "message.CreateMessageRequest")
	proto.RegisterType((*Message)(nil), "message.Message")
	proto.RegisterType((*GetMessagesRequest)(nil), "message.GetMessagesRequest")
	proto.RegisterType((*Messages)(nil), "message.Messages")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	CreateConversation(ctx context.Context, in *CreateConversationRequest, opts ...grpc.CallOption) (*Conversation, error)
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Message, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*Messages, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) CreateConversation(ctx context.Context, in *CreateConversationRequest, opts ...grpc.CallOption) (*Conversation, error) {
	out := new(Conversation)
	err := c.cc.Invoke(ctx, "/message.MessageService/CreateConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/message.MessageService/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*Messages, error) {
	out := new(Messages)
	err := c.cc.Invoke(ctx, "/message.MessageService/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	CreateConversation(context.Context, *CreateConversationRequest) (*Conversation, error)
	CreateMessage(context.Context, *CreateMessageRequest) (*Message, error)
	GetMessages(context.Context, *GetMessagesRequest) (*Messages, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_CreateConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CreateConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/CreateConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CreateConversation(ctx, req.(*CreateConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConversation",
			Handler:    _MessageService_CreateConversation_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _MessageService_CreateMessage_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _MessageService_GetMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/message.proto",
}

func init() { proto.RegisterFile("proto/message.proto", fileDescriptor_message_521c3b15e9248762) }

var fileDescriptor_message_521c3b15e9248762 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xbb, 0x89, 0xd8, 0x76, 0x6a, 0xa3, 0x8e, 0x0a, 0x31, 0xa5, 0x10, 0xf6, 0x62, 0x0e,
	0x52, 0xa1, 0x5e, 0xbc, 0x14, 0xd1, 0x1e, 0xc4, 0x83, 0x07, 0xd3, 0x1f, 0x50, 0xd2, 0xec, 0x20,
	0x39, 0xd8, 0xd5, 0xec, 0x5a, 0xf0, 0x1f, 0xfb, 0x33, 0x84, 0xed, 0x26, 0x6e, 0x13, 0x3c, 0x88,
	0xb7, 0xec, 0xcb, 0xcc, 0xb7, 0x8f, 0xf7, 0x16, 0x4e, 0xde, 0x4a, 0xa9, 0xe5, 0xd5, 0x2b, 0x29,
	0x95, 0xbd, 0xd0, 0xc4, 0x9c, 0xb0, 0x6b, 0x8f, 0x7c, 0x04, 0xe7, 0xf3, 0x92, 0x32, 0x4d, 0x73,
	0xb9, 0xde, 0x50, 0xa9, 0x32, 0x5d, 0xc8, 0x75, 0x4a, 0xef, 0x1f, 0xa4, 0x34, 0x9f, 0xc1, 0x81,
	0x2b, 0x63, 0x00, 0x5e, 0x21, 0x42, 0x16, 0xb3, 0x64, 0x98, 0x7a, 0x85, 0xc0, 0x31, 0x40, 0x6e,
	0x96, 0xc5, 0x32, 0xd3, 0xa1, 0x17, 0xb3, 0xc4, 0x4f, 0xfb, 0x56, 0xb9, 0xd3, 0x7c, 0x01, 0xa7,
	0x5b, 0xf6, 0xd3, 0xf6, 0x32, 0x8b, 0xc5, 0x0b, 0x38, 0xcc, 0x1d, 0xec, 0xb2, 0x66, 0x06, 0xae,
	0xfc, 0x28, 0x10, 0x61, 0x6f, 0x25, 0xc5, 0xa7, 0x21, 0xf7, 0x53, 0xf3, 0xcd, 0x09, 0xba, 0x16,
	0xf7, 0x2f, 0x4e, 0xc3, 0xbb, 0xdf, 0xf4, 0x3e, 0x03, 0x7c, 0x20, 0x6d, 0x6f, 0x52, 0x7f, 0x75,
	0xce, 0x6f, 0xa0, 0x57, 0xed, 0xe2, 0x25, 0xf4, 0x6c, 0xda, 0x2a, 0x64, 0xb1, 0x9f, 0x0c, 0xa6,
	0x47, 0x93, 0xaa, 0x8d, 0x2a, 0x99, 0x7a, 0x62, 0xfa, 0xc5, 0x20, 0xb0, 0xea, 0x82, 0xca, 0x4d,
	0x91, 0x13, 0x3e, 0x03, 0xb6, 0x3b, 0x42, 0x5e, 0x43, 0x7e, 0x2d, 0x30, 0x3a, 0xfb, 0x99, 0x71,
	0xfe, 0xf2, 0x0e, 0xde, 0xc3, 0x70, 0xa7, 0x1a, 0x1c, 0x37, 0x68, 0xbb, 0x95, 0x45, 0x2d, 0xc7,
	0xbc, 0x83, 0xb7, 0x30, 0x70, 0x22, 0xc2, 0x51, 0x3d, 0xd2, 0x0e, 0x2e, 0x3a, 0x6e, 0xee, 0x2b,
	0xde, 0x59, 0xed, 0x9b, 0xb7, 0x78, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x6a, 0x1b, 0xe0,
	0xa2, 0x02, 0x00, 0x00,
}
