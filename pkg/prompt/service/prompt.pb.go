// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/havoc-io/mutagen/pkg/prompt/service/prompt.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import prompt "github.com/havoc-io/mutagen/pkg/prompt"

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

type PromptRequest struct {
	Prompter             string         `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	Prompt               *prompt.Prompt `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PromptRequest) Reset()         { *m = PromptRequest{} }
func (m *PromptRequest) String() string { return proto.CompactTextString(m) }
func (*PromptRequest) ProtoMessage()    {}
func (*PromptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_prompt_b00f61b9fa8ec282, []int{0}
}
func (m *PromptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromptRequest.Unmarshal(m, b)
}
func (m *PromptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromptRequest.Marshal(b, m, deterministic)
}
func (dst *PromptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromptRequest.Merge(dst, src)
}
func (m *PromptRequest) XXX_Size() int {
	return xxx_messageInfo_PromptRequest.Size(m)
}
func (m *PromptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PromptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PromptRequest proto.InternalMessageInfo

func (m *PromptRequest) GetPrompter() string {
	if m != nil {
		return m.Prompter
	}
	return ""
}

func (m *PromptRequest) GetPrompt() *prompt.Prompt {
	if m != nil {
		return m.Prompt
	}
	return nil
}

type PromptResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromptResponse) Reset()         { *m = PromptResponse{} }
func (m *PromptResponse) String() string { return proto.CompactTextString(m) }
func (*PromptResponse) ProtoMessage()    {}
func (*PromptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_prompt_b00f61b9fa8ec282, []int{1}
}
func (m *PromptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromptResponse.Unmarshal(m, b)
}
func (m *PromptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromptResponse.Marshal(b, m, deterministic)
}
func (dst *PromptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromptResponse.Merge(dst, src)
}
func (m *PromptResponse) XXX_Size() int {
	return xxx_messageInfo_PromptResponse.Size(m)
}
func (m *PromptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PromptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PromptResponse proto.InternalMessageInfo

func (m *PromptResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*PromptRequest)(nil), "service.PromptRequest")
	proto.RegisterType((*PromptResponse)(nil), "service.PromptResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PromptClient is the client API for Prompt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PromptClient interface {
	Prompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptResponse, error)
}

type promptClient struct {
	cc *grpc.ClientConn
}

func NewPromptClient(cc *grpc.ClientConn) PromptClient {
	return &promptClient{cc}
}

func (c *promptClient) Prompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptResponse, error) {
	out := new(PromptResponse)
	err := c.cc.Invoke(ctx, "/service.Prompt/Prompt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromptServer is the server API for Prompt service.
type PromptServer interface {
	Prompt(context.Context, *PromptRequest) (*PromptResponse, error)
}

func RegisterPromptServer(s *grpc.Server, srv PromptServer) {
	s.RegisterService(&_Prompt_serviceDesc, srv)
}

func _Prompt_Prompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptServer).Prompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Prompt/Prompt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptServer).Prompt(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Prompt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Prompt",
	HandlerType: (*PromptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prompt",
			Handler:    _Prompt_Prompt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/havoc-io/mutagen/pkg/prompt/service/prompt.proto",
}

func init() {
	proto.RegisterFile("github.com/havoc-io/mutagen/pkg/prompt/service/prompt.proto", fileDescriptor_prompt_b00f61b9fa8ec282)
}

var fileDescriptor_prompt_b00f61b9fa8ec282 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x48, 0x2c, 0xcb, 0x4f, 0xd6, 0xcd, 0xcc, 0xd7,
	0xcf, 0x2d, 0x2d, 0x49, 0x4c, 0x4f, 0xcd, 0xd3, 0x2f, 0xc8, 0x4e, 0xd7, 0x2f, 0x28, 0xca, 0xcf,
	0x2d, 0x28, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x85, 0x72, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xd8, 0xa1, 0xa2, 0x52, 0xc6, 0x44, 0x9a, 0x82, 0xac, 0x5b, 0x29, 0x98, 0x8b,
	0x37, 0x00, 0xcc, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x80, 0x28,
	0x48, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0xd4, 0xb8, 0xd8, 0x20,
	0x6c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0xa8, 0x59, 0x50, 0x23, 0xa0, 0xb2,
	0x4a, 0x3a, 0x5c, 0x7c, 0x30, 0x43, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x41, 0xa6, 0x16, 0x41,
	0xd9, 0x30, 0x53, 0x61, 0x7c, 0x23, 0x57, 0x2e, 0x36, 0x88, 0x6a, 0x21, 0x6b, 0x38, 0x4b, 0x4c,
	0x0f, 0xea, 0x2b, 0x3d, 0x14, 0xd7, 0x49, 0x89, 0x63, 0x88, 0x43, 0x0c, 0x51, 0x62, 0x48, 0x62,
	0x03, 0x7b, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xb0, 0x3a, 0x40, 0x4d, 0x01, 0x00,
	0x00,
}
