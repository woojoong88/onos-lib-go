// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/logging/logging.proto

package logging

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ResponseStatus int32

const (
	ResponseStatus_OK                  ResponseStatus = 0
	ResponseStatus_FAILED              ResponseStatus = 1
	ResponseStatus_PRECONDITION_FAILED ResponseStatus = 2
)

var ResponseStatus_name = map[int32]string{
	0: "OK",
	1: "FAILED",
	2: "PRECONDITION_FAILED",
}

var ResponseStatus_value = map[string]int32{
	"OK":                  0,
	"FAILED":              1,
	"PRECONDITION_FAILED": 2,
}

func (x ResponseStatus) String() string {
	return proto.EnumName(ResponseStatus_name, int32(x))
}

func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_282a7fd5e5377550, []int{0}
}

// Logger level
type Level int32

const (
	// Debug log level
	Level_DEBUG Level = 0
	// Info log level
	Level_INFO Level = 1
	// Warn log level
	Level_WARN Level = 2
	// Error log level
	Level_ERROR Level = 3
	// DPanic log level
	Level_DPANIC Level = 4
	// Panic log level
	Level_PANIC Level = 5
	// Fatal log level
	Level_FATAL Level = 6
)

var Level_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARN",
	3: "ERROR",
	4: "DPANIC",
	5: "PANIC",
	6: "FATAL",
}

var Level_value = map[string]int32{
	"DEBUG":  0,
	"INFO":   1,
	"WARN":   2,
	"ERROR":  3,
	"DPANIC": 4,
	"PANIC":  5,
	"FATAL":  6,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}

func (Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_282a7fd5e5377550, []int{1}
}

// SetLevelRequest request for setting a logger level
type SetLevelRequest struct {
	// logger name
	LoggerName string `protobuf:"bytes,1,opt,name=logger_name,json=loggerName,proto3" json:"logger_name,omitempty"`
	// logger level
	Level                Level    `protobuf:"varint,2,opt,name=level,proto3,enum=onos.lib.go.logging.Level" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetLevelRequest) Reset()         { *m = SetLevelRequest{} }
func (m *SetLevelRequest) String() string { return proto.CompactTextString(m) }
func (*SetLevelRequest) ProtoMessage()    {}
func (*SetLevelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_282a7fd5e5377550, []int{0}
}
func (m *SetLevelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetLevelRequest.Unmarshal(m, b)
}
func (m *SetLevelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetLevelRequest.Marshal(b, m, deterministic)
}
func (m *SetLevelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetLevelRequest.Merge(m, src)
}
func (m *SetLevelRequest) XXX_Size() int {
	return xxx_messageInfo_SetLevelRequest.Size(m)
}
func (m *SetLevelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetLevelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetLevelRequest proto.InternalMessageInfo

func (m *SetLevelRequest) GetLoggerName() string {
	if m != nil {
		return m.LoggerName
	}
	return ""
}

func (m *SetLevelRequest) GetLevel() Level {
	if m != nil {
		return m.Level
	}
	return Level_DEBUG
}

// SetLevelResponse response for setting a logger level
type SetLevelResponse struct {
	ResponseStatus       ResponseStatus `protobuf:"varint,1,opt,name=response_status,json=responseStatus,proto3,enum=onos.lib.go.logging.ResponseStatus" json:"response_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SetLevelResponse) Reset()         { *m = SetLevelResponse{} }
func (m *SetLevelResponse) String() string { return proto.CompactTextString(m) }
func (*SetLevelResponse) ProtoMessage()    {}
func (*SetLevelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_282a7fd5e5377550, []int{1}
}
func (m *SetLevelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetLevelResponse.Unmarshal(m, b)
}
func (m *SetLevelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetLevelResponse.Marshal(b, m, deterministic)
}
func (m *SetLevelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetLevelResponse.Merge(m, src)
}
func (m *SetLevelResponse) XXX_Size() int {
	return xxx_messageInfo_SetLevelResponse.Size(m)
}
func (m *SetLevelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetLevelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetLevelResponse proto.InternalMessageInfo

func (m *SetLevelResponse) GetResponseStatus() ResponseStatus {
	if m != nil {
		return m.ResponseStatus
	}
	return ResponseStatus_OK
}

// SetDebugModeRequest enable/disable debug mode of the logger
type SetDebugModeRequest struct {
	Debug                bool     `protobuf:"varint,1,opt,name=debug,proto3" json:"debug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDebugModeRequest) Reset()         { *m = SetDebugModeRequest{} }
func (m *SetDebugModeRequest) String() string { return proto.CompactTextString(m) }
func (*SetDebugModeRequest) ProtoMessage()    {}
func (*SetDebugModeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_282a7fd5e5377550, []int{2}
}
func (m *SetDebugModeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDebugModeRequest.Unmarshal(m, b)
}
func (m *SetDebugModeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDebugModeRequest.Marshal(b, m, deterministic)
}
func (m *SetDebugModeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDebugModeRequest.Merge(m, src)
}
func (m *SetDebugModeRequest) XXX_Size() int {
	return xxx_messageInfo_SetDebugModeRequest.Size(m)
}
func (m *SetDebugModeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDebugModeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDebugModeRequest proto.InternalMessageInfo

func (m *SetDebugModeRequest) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

// SetDebugModeResponse response for setting debug mode of a logger
type SetDebugModeResponse struct {
	ResponseStatus       ResponseStatus `protobuf:"varint,1,opt,name=response_status,json=responseStatus,proto3,enum=onos.lib.go.logging.ResponseStatus" json:"response_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SetDebugModeResponse) Reset()         { *m = SetDebugModeResponse{} }
func (m *SetDebugModeResponse) String() string { return proto.CompactTextString(m) }
func (*SetDebugModeResponse) ProtoMessage()    {}
func (*SetDebugModeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_282a7fd5e5377550, []int{3}
}
func (m *SetDebugModeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDebugModeResponse.Unmarshal(m, b)
}
func (m *SetDebugModeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDebugModeResponse.Marshal(b, m, deterministic)
}
func (m *SetDebugModeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDebugModeResponse.Merge(m, src)
}
func (m *SetDebugModeResponse) XXX_Size() int {
	return xxx_messageInfo_SetDebugModeResponse.Size(m)
}
func (m *SetDebugModeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDebugModeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetDebugModeResponse proto.InternalMessageInfo

func (m *SetDebugModeResponse) GetResponseStatus() ResponseStatus {
	if m != nil {
		return m.ResponseStatus
	}
	return ResponseStatus_OK
}

func init() {
	proto.RegisterEnum("onos.lib.go.logging.ResponseStatus", ResponseStatus_name, ResponseStatus_value)
	proto.RegisterEnum("onos.lib.go.logging.Level", Level_name, Level_value)
	proto.RegisterType((*SetLevelRequest)(nil), "onos.lib.go.logging.SetLevelRequest")
	proto.RegisterType((*SetLevelResponse)(nil), "onos.lib.go.logging.SetLevelResponse")
	proto.RegisterType((*SetDebugModeRequest)(nil), "onos.lib.go.logging.SetDebugModeRequest")
	proto.RegisterType((*SetDebugModeResponse)(nil), "onos.lib.go.logging.SetDebugModeResponse")
}

func init() { proto.RegisterFile("api/logging/logging.proto", fileDescriptor_282a7fd5e5377550) }

var fileDescriptor_282a7fd5e5377550 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0x4e, 0xb3, 0x40,
	0x14, 0x2d, 0x7c, 0x85, 0xb4, 0xf7, 0x4b, 0xe8, 0x64, 0xda, 0xc4, 0xda, 0x8d, 0x0d, 0x6a, 0x52,
	0x6b, 0x82, 0xa6, 0xae, 0x5d, 0x60, 0xa1, 0x86, 0x88, 0xd0, 0x0c, 0x35, 0x5d, 0x22, 0x0d, 0x13,
	0xd2, 0x84, 0x32, 0x15, 0xa8, 0x0f, 0xe8, 0x93, 0x99, 0xe1, 0x27, 0xa6, 0x4d, 0x8d, 0x6e, 0x5c,
	0x71, 0xe6, 0x9e, 0x73, 0xcf, 0xe5, 0xfe, 0xc0, 0x69, 0xb0, 0x5d, 0xdf, 0xc4, 0x2c, 0x8a, 0xd6,
	0x49, 0x54, 0x7f, 0xb5, 0x6d, 0xca, 0x72, 0x86, 0xbb, 0x2c, 0x61, 0x99, 0x16, 0xaf, 0x57, 0x5a,
	0xc4, 0xb4, 0x8a, 0x52, 0x43, 0xe8, 0x78, 0x34, 0xb7, 0xe9, 0x3b, 0x8d, 0x09, 0x7d, 0xdb, 0xd1,
	0x2c, 0xc7, 0x67, 0xf0, 0x9f, 0xb3, 0x34, 0xf5, 0x93, 0x60, 0x43, 0xfb, 0xc2, 0x50, 0x18, 0xb5,
	0x09, 0x94, 0x21, 0x27, 0xd8, 0x50, 0x7c, 0x0b, 0x52, 0xcc, 0x13, 0xfa, 0xe2, 0x50, 0x18, 0x29,
	0x93, 0x81, 0x76, 0xc4, 0x58, 0x2b, 0x2d, 0x4b, 0xa1, 0xfa, 0x0a, 0xe8, 0xab, 0x4a, 0xb6, 0x65,
	0x49, 0x46, 0xb1, 0x0d, 0x9d, 0xb4, 0xc2, 0x7e, 0x96, 0x07, 0xf9, 0x2e, 0x2b, 0x4a, 0x29, 0x93,
	0xf3, 0xa3, 0x7e, 0x75, 0x9e, 0x57, 0x48, 0x89, 0x92, 0xee, 0xbd, 0xd5, 0x6b, 0xe8, 0x7a, 0x34,
	0x37, 0xe8, 0x6a, 0x17, 0x3d, 0xb3, 0x90, 0xd6, 0xbd, 0xf4, 0x40, 0x0a, 0x79, 0xac, 0xb0, 0x6e,
	0x91, 0xf2, 0xa1, 0x86, 0xd0, 0xdb, 0x17, 0xff, 0xc5, 0x2f, 0x8d, 0xef, 0x41, 0xd9, 0x57, 0x60,
	0x19, 0x44, 0xf7, 0x09, 0x35, 0x30, 0x80, 0x3c, 0xd3, 0x2d, 0xdb, 0x34, 0x90, 0x80, 0x4f, 0xa0,
	0x3b, 0x27, 0xe6, 0xd4, 0x75, 0x0c, 0x6b, 0x61, 0xb9, 0x8e, 0x5f, 0x11, 0xe2, 0xd8, 0x03, 0xa9,
	0x18, 0x18, 0x6e, 0x83, 0x64, 0x98, 0x0f, 0x2f, 0x8f, 0xa8, 0x81, 0x5b, 0xd0, 0xb4, 0x9c, 0x99,
	0x8b, 0x04, 0x8e, 0x96, 0x3a, 0x71, 0x90, 0xc8, 0x69, 0x93, 0x10, 0x97, 0xa0, 0x7f, 0xdc, 0xd7,
	0x98, 0xeb, 0x8e, 0x35, 0x45, 0x4d, 0x1e, 0x2e, 0xa1, 0xc4, 0xe1, 0x4c, 0x5f, 0xe8, 0x36, 0x92,
	0x27, 0x1f, 0x02, 0xc8, 0xe5, 0x26, 0xf1, 0x12, 0x5a, 0xf5, 0x4e, 0xf0, 0xc5, 0xd1, 0xfe, 0x0e,
	0x0e, 0x63, 0x70, 0xf9, 0x83, 0xaa, 0x9a, 0xa2, 0x5f, 0x18, 0x17, 0xd3, 0xc5, 0xa3, 0xef, 0x52,
	0x0e, 0x37, 0x35, 0xb8, 0xfa, 0x85, 0xb2, 0x2c, 0xb0, 0x92, 0x8b, 0x7b, 0xbe, 0xfb, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0xfa, 0x64, 0x3d, 0xe0, 0xec, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoggerClient is the client API for Logger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoggerClient interface {
	// Sets a logger level
	SetLevel(ctx context.Context, in *SetLevelRequest, opts ...grpc.CallOption) (*SetLevelResponse, error)
	// Sets debug mode to debug logging package
	SetDebug(ctx context.Context, in *SetDebugModeRequest, opts ...grpc.CallOption) (*SetDebugModeResponse, error)
}

type loggerClient struct {
	cc *grpc.ClientConn
}

func NewLoggerClient(cc *grpc.ClientConn) LoggerClient {
	return &loggerClient{cc}
}

func (c *loggerClient) SetLevel(ctx context.Context, in *SetLevelRequest, opts ...grpc.CallOption) (*SetLevelResponse, error) {
	out := new(SetLevelResponse)
	err := c.cc.Invoke(ctx, "/onos.lib.go.logging.logger/SetLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerClient) SetDebug(ctx context.Context, in *SetDebugModeRequest, opts ...grpc.CallOption) (*SetDebugModeResponse, error) {
	out := new(SetDebugModeResponse)
	err := c.cc.Invoke(ctx, "/onos.lib.go.logging.logger/SetDebug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServer is the server API for Logger service.
type LoggerServer interface {
	// Sets a logger level
	SetLevel(context.Context, *SetLevelRequest) (*SetLevelResponse, error)
	// Sets debug mode to debug logging package
	SetDebug(context.Context, *SetDebugModeRequest) (*SetDebugModeResponse, error)
}

// UnimplementedLoggerServer can be embedded to have forward compatible implementations.
type UnimplementedLoggerServer struct {
}

func (*UnimplementedLoggerServer) SetLevel(ctx context.Context, req *SetLevelRequest) (*SetLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLevel not implemented")
}
func (*UnimplementedLoggerServer) SetDebug(ctx context.Context, req *SetDebugModeRequest) (*SetDebugModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDebug not implemented")
}

func RegisterLoggerServer(s *grpc.Server, srv LoggerServer) {
	s.RegisterService(&_Logger_serviceDesc, srv)
}

func _Logger_SetLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).SetLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onos.lib.go.logging.logger/SetLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).SetLevel(ctx, req.(*SetLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logger_SetDebug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDebugModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).SetDebug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onos.lib.go.logging.logger/SetDebug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).SetDebug(ctx, req.(*SetDebugModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onos.lib.go.logging.logger",
	HandlerType: (*LoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLevel",
			Handler:    _Logger_SetLevel_Handler,
		},
		{
			MethodName: "SetDebug",
			Handler:    _Logger_SetDebug_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/logging/logging.proto",
}
