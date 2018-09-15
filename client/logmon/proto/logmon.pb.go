// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logmon.proto

package proto

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

type StartRequest struct {
	LogDir               string   `protobuf:"bytes,1,opt,name=log_dir,json=logDir,proto3" json:"log_dir,omitempty"`
	StdoutFileName       string   `protobuf:"bytes,2,opt,name=stdout_file_name,json=stdoutFileName,proto3" json:"stdout_file_name,omitempty"`
	StderrFileName       string   `protobuf:"bytes,3,opt,name=stderr_file_name,json=stderrFileName,proto3" json:"stderr_file_name,omitempty"`
	MaxFiles             uint32   `protobuf:"varint,4,opt,name=max_files,json=maxFiles,proto3" json:"max_files,omitempty"`
	MaxFileSizeMb        uint32   `protobuf:"varint,5,opt,name=max_file_size_mb,json=maxFileSizeMb,proto3" json:"max_file_size_mb,omitempty"`
	Uid                  uint32   `protobuf:"varint,6,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid                  uint32   `protobuf:"varint,7,opt,name=gid,proto3" json:"gid,omitempty"`
	StdoutFifo           string   `protobuf:"bytes,8,opt,name=stdout_fifo,json=stdoutFifo,proto3" json:"stdout_fifo,omitempty"`
	StderrFifo           string   `protobuf:"bytes,9,opt,name=stderr_fifo,json=stderrFifo,proto3" json:"stderr_fifo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_logmon_2ea9756cd5b55bdb, []int{0}
}
func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (dst *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(dst, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetLogDir() string {
	if m != nil {
		return m.LogDir
	}
	return ""
}

func (m *StartRequest) GetStdoutFileName() string {
	if m != nil {
		return m.StdoutFileName
	}
	return ""
}

func (m *StartRequest) GetStderrFileName() string {
	if m != nil {
		return m.StderrFileName
	}
	return ""
}

func (m *StartRequest) GetMaxFiles() uint32 {
	if m != nil {
		return m.MaxFiles
	}
	return 0
}

func (m *StartRequest) GetMaxFileSizeMb() uint32 {
	if m != nil {
		return m.MaxFileSizeMb
	}
	return 0
}

func (m *StartRequest) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *StartRequest) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *StartRequest) GetStdoutFifo() string {
	if m != nil {
		return m.StdoutFifo
	}
	return ""
}

func (m *StartRequest) GetStderrFifo() string {
	if m != nil {
		return m.StderrFifo
	}
	return ""
}

type StartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_logmon_2ea9756cd5b55bdb, []int{1}
}
func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (dst *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(dst, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

type StopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_logmon_2ea9756cd5b55bdb, []int{2}
}
func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (dst *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(dst, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

type StopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_logmon_2ea9756cd5b55bdb, []int{3}
}
func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (dst *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(dst, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StartRequest)(nil), "hashicorp.nomad.client.logmon.proto.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "hashicorp.nomad.client.logmon.proto.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "hashicorp.nomad.client.logmon.proto.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "hashicorp.nomad.client.logmon.proto.StopResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogMonClient is the client API for LogMon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogMonClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type logMonClient struct {
	cc *grpc.ClientConn
}

func NewLogMonClient(cc *grpc.ClientConn) LogMonClient {
	return &logMonClient{cc}
}

func (c *logMonClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.client.logmon.proto.LogMon/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logMonClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.client.logmon.proto.LogMon/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogMonServer is the server API for LogMon service.
type LogMonServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
}

func RegisterLogMonServer(s *grpc.Server, srv LogMonServer) {
	s.RegisterService(&_LogMon_serviceDesc, srv)
}

func _LogMon_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogMonServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.client.logmon.proto.LogMon/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogMonServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogMon_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogMonServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.client.logmon.proto.LogMon/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogMonServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogMon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.client.logmon.proto.LogMon",
	HandlerType: (*LogMonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _LogMon_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _LogMon_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logmon.proto",
}

func init() { proto.RegisterFile("logmon.proto", fileDescriptor_logmon_2ea9756cd5b55bdb) }

var fileDescriptor_logmon_2ea9756cd5b55bdb = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbf, 0x6e, 0xf2, 0x30,
	0x14, 0xc5, 0xbf, 0xf0, 0x27, 0xc0, 0x85, 0xf0, 0x21, 0x2f, 0xb5, 0xe8, 0x50, 0x94, 0x0e, 0x65,
	0x8a, 0x0a, 0x7d, 0x83, 0xaa, 0xea, 0x54, 0x3a, 0xc0, 0xd6, 0x25, 0x32, 0xc4, 0x09, 0x96, 0xe2,
	0xdc, 0xd4, 0x36, 0x12, 0xe2, 0x85, 0xfa, 0x6c, 0x7d, 0x8b, 0x2a, 0x8e, 0x89, 0x18, 0x61, 0x4a,
	0x7c, 0xce, 0xef, 0xc8, 0xc7, 0xf7, 0xc2, 0x28, 0xc7, 0x4c, 0x62, 0x11, 0x95, 0x0a, 0x0d, 0x92,
	0xc7, 0x3d, 0xd3, 0x7b, 0xb1, 0x43, 0x55, 0x46, 0x05, 0x4a, 0x96, 0x44, 0xbb, 0x5c, 0xf0, 0xc2,
	0x44, 0x97, 0x50, 0xf8, 0xd3, 0x82, 0xd1, 0xc6, 0x30, 0x65, 0xd6, 0xfc, 0xfb, 0xc0, 0xb5, 0x21,
	0x77, 0xd0, 0xcb, 0x31, 0x8b, 0x13, 0xa1, 0xa8, 0x37, 0xf3, 0xe6, 0x83, 0xb5, 0x9f, 0x63, 0xf6,
	0x26, 0x14, 0x99, 0xc3, 0x44, 0x9b, 0x04, 0x0f, 0x26, 0x4e, 0x45, 0xce, 0xe3, 0x82, 0x49, 0x4e,
	0x5b, 0x96, 0x18, 0xd7, 0xfa, 0xbb, 0xc8, 0xf9, 0x27, 0x93, 0xdc, 0x91, 0x5c, 0xa9, 0x0b, 0xb2,
	0xdd, 0x90, 0x5c, 0xa9, 0x86, 0xbc, 0x87, 0x81, 0x64, 0x47, 0x8b, 0x69, 0xda, 0x99, 0x79, 0xf3,
	0x60, 0xdd, 0x97, 0xec, 0x58, 0xf9, 0x9a, 0x3c, 0xc1, 0xe4, 0x6c, 0xc6, 0x5a, 0x9c, 0x78, 0x2c,
	0xb7, 0xb4, 0x6b, 0x99, 0xc0, 0x31, 0x1b, 0x71, 0xe2, 0xab, 0x2d, 0x99, 0x40, 0xfb, 0x20, 0x12,
	0xea, 0x5b, 0xaf, 0xfa, 0xad, 0x94, 0x4c, 0x24, 0xb4, 0x57, 0x2b, 0x99, 0x48, 0xc8, 0x03, 0x0c,
	0x9b, 0xf6, 0x29, 0xd2, 0xbe, 0xad, 0x03, 0xe7, 0xe2, 0x29, 0x3a, 0xa0, 0x2e, 0x9d, 0x22, 0x1d,
	0x34, 0x80, 0xed, 0x9b, 0x62, 0xf8, 0x1f, 0x02, 0x37, 0x28, 0x5d, 0x62, 0xa1, 0x79, 0x18, 0xc0,
	0x70, 0x63, 0xb0, 0x74, 0x83, 0x0b, 0xc7, 0xd5, 0x20, 0xab, 0x63, 0x6d, 0x2f, 0x7f, 0x3d, 0xf0,
	0x3f, 0x30, 0x5b, 0x61, 0x41, 0x4a, 0xe8, 0xda, 0x28, 0x59, 0x44, 0x57, 0xec, 0x24, 0xba, 0xdc,
	0xc7, 0x74, 0x79, 0x4b, 0xc4, 0x35, 0xfb, 0x47, 0x24, 0x74, 0xaa, 0x32, 0xe4, 0xf9, 0xca, 0x74,
	0xf3, 0x8c, 0xe9, 0xe2, 0x86, 0xc4, 0xf9, 0xba, 0xd7, 0xde, 0x57, 0xd7, 0xea, 0x5b, 0xdf, 0x7e,
	0x5e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x00, 0x4c, 0x02, 0x82, 0x8a, 0x02, 0x00, 0x00,
}
