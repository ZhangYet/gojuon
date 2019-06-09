// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gojuon.proto

package gojuon_dict

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SearchRequest struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f19b88db5b32fd64, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type WordRecord struct {
	Japanese             string               `protobuf:"bytes,1,opt,name=japanese,proto3" json:"japanese,omitempty"`
	Furigana             string               `protobuf:"bytes,2,opt,name=furigana,proto3" json:"furigana,omitempty"`
	English              string               `protobuf:"bytes,3,opt,name=english,proto3" json:"english,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WordRecord) Reset()         { *m = WordRecord{} }
func (m *WordRecord) String() string { return proto.CompactTextString(m) }
func (*WordRecord) ProtoMessage()    {}
func (*WordRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f19b88db5b32fd64, []int{1}
}

func (m *WordRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordRecord.Unmarshal(m, b)
}
func (m *WordRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordRecord.Marshal(b, m, deterministic)
}
func (m *WordRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordRecord.Merge(m, src)
}
func (m *WordRecord) XXX_Size() int {
	return xxx_messageInfo_WordRecord.Size(m)
}
func (m *WordRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_WordRecord.DiscardUnknown(m)
}

var xxx_messageInfo_WordRecord proto.InternalMessageInfo

func (m *WordRecord) GetJapanese() string {
	if m != nil {
		return m.Japanese
	}
	return ""
}

func (m *WordRecord) GetFurigana() string {
	if m != nil {
		return m.Furigana
	}
	return ""
}

func (m *WordRecord) GetEnglish() string {
	if m != nil {
		return m.English
	}
	return ""
}

func (m *WordRecord) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type SearchResponse struct {
	Record               *WordRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f19b88db5b32fd64, []int{2}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetRecord() *WordRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type RecordRequest struct {
	Record               *WordRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RecordRequest) Reset()         { *m = RecordRequest{} }
func (m *RecordRequest) String() string { return proto.CompactTextString(m) }
func (*RecordRequest) ProtoMessage()    {}
func (*RecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f19b88db5b32fd64, []int{3}
}

func (m *RecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordRequest.Unmarshal(m, b)
}
func (m *RecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordRequest.Marshal(b, m, deterministic)
}
func (m *RecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordRequest.Merge(m, src)
}
func (m *RecordRequest) XXX_Size() int {
	return xxx_messageInfo_RecordRequest.Size(m)
}
func (m *RecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecordRequest proto.InternalMessageInfo

func (m *RecordRequest) GetRecord() *WordRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type RecordResponse struct {
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,1,opt,name=createTime,proto3" json:"createTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RecordResponse) Reset()         { *m = RecordResponse{} }
func (m *RecordResponse) String() string { return proto.CompactTextString(m) }
func (*RecordResponse) ProtoMessage()    {}
func (*RecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f19b88db5b32fd64, []int{4}
}

func (m *RecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordResponse.Unmarshal(m, b)
}
func (m *RecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordResponse.Marshal(b, m, deterministic)
}
func (m *RecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordResponse.Merge(m, src)
}
func (m *RecordResponse) XXX_Size() int {
	return xxx_messageInfo_RecordResponse.Size(m)
}
func (m *RecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecordResponse proto.InternalMessageInfo

func (m *RecordResponse) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "gojuon.dict.SearchRequest")
	proto.RegisterType((*WordRecord)(nil), "gojuon.dict.WordRecord")
	proto.RegisterType((*SearchResponse)(nil), "gojuon.dict.SearchResponse")
	proto.RegisterType((*RecordRequest)(nil), "gojuon.dict.RecordRequest")
	proto.RegisterType((*RecordResponse)(nil), "gojuon.dict.RecordResponse")
}

func init() { proto.RegisterFile("gojuon.proto", fileDescriptor_f19b88db5b32fd64) }

var fileDescriptor_f19b88db5b32fd64 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4f, 0xf3, 0x30,
	0x18, 0x84, 0x9b, 0xef, 0x43, 0x05, 0xde, 0xd0, 0x0e, 0x5e, 0x88, 0xc2, 0x40, 0xe5, 0xa9, 0x2c,
	0xae, 0x54, 0x36, 0x26, 0x10, 0x8c, 0x4c, 0x69, 0x25, 0x66, 0xd7, 0x79, 0x9b, 0xba, 0xb4, 0x71,
	0xb0, 0x1d, 0x10, 0xbf, 0x04, 0x89, 0x5f, 0x8b, 0x1a, 0xdb, 0x50, 0x23, 0x96, 0x8e, 0xe7, 0x3b,
	0x9d, 0xef, 0xb1, 0xe1, 0xac, 0x52, 0xeb, 0x56, 0xd5, 0xac, 0xd1, 0xca, 0x2a, 0x92, 0x7a, 0x55,
	0x4a, 0x61, 0xf3, 0xcb, 0x4a, 0xa9, 0x6a, 0x83, 0x93, 0xce, 0x5a, 0xb4, 0xcb, 0x89, 0x95, 0x5b,
	0x34, 0x96, 0x6f, 0x1b, 0x97, 0xa6, 0x57, 0x30, 0x98, 0x21, 0xd7, 0x62, 0x55, 0xe0, 0x4b, 0x8b,
	0xc6, 0x92, 0x0c, 0x8e, 0x9f, 0xf1, 0xfd, 0x4d, 0xe9, 0x32, 0x4b, 0x46, 0xc9, 0xf8, 0xb4, 0x08,
	0x92, 0x7e, 0x26, 0x00, 0x4f, 0x4a, 0x97, 0x05, 0x0a, 0xa5, 0x4b, 0x92, 0xc3, 0xc9, 0x9a, 0x37,
	0xbc, 0x46, 0x83, 0x3e, 0xf9, 0xad, 0x77, 0xde, 0xb2, 0xd5, 0xb2, 0xe2, 0x35, 0xcf, 0xfe, 0x39,
	0x2f, 0xe8, 0xdd, 0x05, 0x58, 0x57, 0x1b, 0x69, 0x56, 0xd9, 0x7f, 0x77, 0x81, 0x97, 0xe4, 0x06,
	0x40, 0x68, 0xe4, 0x16, 0xe7, 0x72, 0x8b, 0xd9, 0xd1, 0x28, 0x19, 0xa7, 0xd3, 0x9c, 0x39, 0x02,
	0x16, 0x08, 0xd8, 0x3c, 0x10, 0x14, 0x7b, 0x69, 0x7a, 0x07, 0xc3, 0xc0, 0x61, 0x1a, 0x55, 0x1b,
	0x24, 0x13, 0xe8, 0xeb, 0x6e, 0x69, 0xb7, 0x2e, 0x9d, 0x9e, 0xb3, 0xbd, 0x87, 0x61, 0x3f, 0x20,
	0x85, 0x8f, 0xd1, 0x5b, 0x18, 0xf8, 0x13, 0xff, 0x14, 0x07, 0x37, 0x3c, 0xc2, 0x30, 0x34, 0xf8,
	0x11, 0x31, 0x52, 0x72, 0x08, 0xd2, 0xf4, 0x23, 0x81, 0xf4, 0x41, 0x0a, 0x3b, 0x43, 0xfd, 0x2a,
	0x05, 0x92, 0x7b, 0xe8, 0x3b, 0x44, 0x92, 0x47, 0x43, 0xa2, 0xff, 0xcb, 0x2f, 0xfe, 0xf4, 0xdc,
	0x1c, 0xda, 0xdb, 0x95, 0x84, 0xff, 0x8b, 0x82, 0x11, 0xf9, 0xaf, 0x92, 0x98, 0x89, 0xf6, 0x16,
	0xfd, 0x6e, 0xf9, 0xf5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x48, 0x11, 0xd7, 0x79, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DictServiceClient is the client API for DictService server.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DictServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Record(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*RecordResponse, error)
}

type dictServiceClient struct {
	cc *grpc.ClientConn
}

func NewDictServiceClient(cc *grpc.ClientConn) DictServiceClient {
	return &dictServiceClient{cc}
}

func (c *dictServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/gojuon.dict.DictService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) Record(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := c.cc.Invoke(ctx, "/gojuon.dict.DictService/Record", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictServiceServer is the server API for DictService server.
type DictServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Record(context.Context, *RecordRequest) (*RecordResponse, error)
}

func RegisterDictServiceServer(s *grpc.Server, srv DictServiceServer) {
	s.RegisterService(&_DictService_serviceDesc, srv)
}

func _DictService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gojuon.dict.DictService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gojuon.dict.DictService/Record",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).Record(ctx, req.(*RecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DictService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gojuon.dict.DictService",
	HandlerType: (*DictServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _DictService_Search_Handler,
		},
		{
			MethodName: "Record",
			Handler:    _DictService_Record_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gojuon.proto",
}
