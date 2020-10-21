// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/tracing/v2/custom_tag.proto

package envoy_type_tracing_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/type/metadata/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type CustomTag struct {
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*CustomTag_Literal_
	//	*CustomTag_Environment_
	//	*CustomTag_RequestHeader
	//	*CustomTag_Metadata_
	Type                 isCustomTag_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CustomTag) Reset()         { *m = CustomTag{} }
func (m *CustomTag) String() string { return proto.CompactTextString(m) }
func (*CustomTag) ProtoMessage()    {}
func (*CustomTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0}
}

func (m *CustomTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag.Unmarshal(m, b)
}
func (m *CustomTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag.Marshal(b, m, deterministic)
}
func (m *CustomTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag.Merge(m, src)
}
func (m *CustomTag) XXX_Size() int {
	return xxx_messageInfo_CustomTag.Size(m)
}
func (m *CustomTag) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag proto.InternalMessageInfo

func (m *CustomTag) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type isCustomTag_Type interface {
	isCustomTag_Type()
}

type CustomTag_Literal_ struct {
	Literal *CustomTag_Literal `protobuf:"bytes,2,opt,name=literal,proto3,oneof"`
}

type CustomTag_Environment_ struct {
	Environment *CustomTag_Environment `protobuf:"bytes,3,opt,name=environment,proto3,oneof"`
}

type CustomTag_RequestHeader struct {
	RequestHeader *CustomTag_Header `protobuf:"bytes,4,opt,name=request_header,json=requestHeader,proto3,oneof"`
}

type CustomTag_Metadata_ struct {
	Metadata *CustomTag_Metadata `protobuf:"bytes,5,opt,name=metadata,proto3,oneof"`
}

func (*CustomTag_Literal_) isCustomTag_Type() {}

func (*CustomTag_Environment_) isCustomTag_Type() {}

func (*CustomTag_RequestHeader) isCustomTag_Type() {}

func (*CustomTag_Metadata_) isCustomTag_Type() {}

func (m *CustomTag) GetType() isCustomTag_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *CustomTag) GetLiteral() *CustomTag_Literal {
	if x, ok := m.GetType().(*CustomTag_Literal_); ok {
		return x.Literal
	}
	return nil
}

func (m *CustomTag) GetEnvironment() *CustomTag_Environment {
	if x, ok := m.GetType().(*CustomTag_Environment_); ok {
		return x.Environment
	}
	return nil
}

func (m *CustomTag) GetRequestHeader() *CustomTag_Header {
	if x, ok := m.GetType().(*CustomTag_RequestHeader); ok {
		return x.RequestHeader
	}
	return nil
}

func (m *CustomTag) GetMetadata() *CustomTag_Metadata {
	if x, ok := m.GetType().(*CustomTag_Metadata_); ok {
		return x.Metadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomTag) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomTag_Literal_)(nil),
		(*CustomTag_Environment_)(nil),
		(*CustomTag_RequestHeader)(nil),
		(*CustomTag_Metadata_)(nil),
	}
}

type CustomTag_Literal struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Literal) Reset()         { *m = CustomTag_Literal{} }
func (m *CustomTag_Literal) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Literal) ProtoMessage()    {}
func (*CustomTag_Literal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0, 0}
}

func (m *CustomTag_Literal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Literal.Unmarshal(m, b)
}
func (m *CustomTag_Literal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Literal.Marshal(b, m, deterministic)
}
func (m *CustomTag_Literal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Literal.Merge(m, src)
}
func (m *CustomTag_Literal) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Literal.Size(m)
}
func (m *CustomTag_Literal) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Literal.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Literal proto.InternalMessageInfo

func (m *CustomTag_Literal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CustomTag_Environment struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DefaultValue         string   `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Environment) Reset()         { *m = CustomTag_Environment{} }
func (m *CustomTag_Environment) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Environment) ProtoMessage()    {}
func (*CustomTag_Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0, 1}
}

func (m *CustomTag_Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Environment.Unmarshal(m, b)
}
func (m *CustomTag_Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Environment.Marshal(b, m, deterministic)
}
func (m *CustomTag_Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Environment.Merge(m, src)
}
func (m *CustomTag_Environment) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Environment.Size(m)
}
func (m *CustomTag_Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Environment proto.InternalMessageInfo

func (m *CustomTag_Environment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomTag_Environment) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

type CustomTag_Header struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DefaultValue         string   `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Header) Reset()         { *m = CustomTag_Header{} }
func (m *CustomTag_Header) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Header) ProtoMessage()    {}
func (*CustomTag_Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0, 2}
}

func (m *CustomTag_Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Header.Unmarshal(m, b)
}
func (m *CustomTag_Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Header.Marshal(b, m, deterministic)
}
func (m *CustomTag_Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Header.Merge(m, src)
}
func (m *CustomTag_Header) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Header.Size(m)
}
func (m *CustomTag_Header) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Header.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Header proto.InternalMessageInfo

func (m *CustomTag_Header) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomTag_Header) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

type CustomTag_Metadata struct {
	Kind                 *v2.MetadataKind `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	MetadataKey          *v2.MetadataKey  `protobuf:"bytes,2,opt,name=metadata_key,json=metadataKey,proto3" json:"metadata_key,omitempty"`
	DefaultValue         string           `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CustomTag_Metadata) Reset()         { *m = CustomTag_Metadata{} }
func (m *CustomTag_Metadata) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Metadata) ProtoMessage()    {}
func (*CustomTag_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0, 3}
}

func (m *CustomTag_Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Metadata.Unmarshal(m, b)
}
func (m *CustomTag_Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Metadata.Marshal(b, m, deterministic)
}
func (m *CustomTag_Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Metadata.Merge(m, src)
}
func (m *CustomTag_Metadata) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Metadata.Size(m)
}
func (m *CustomTag_Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Metadata proto.InternalMessageInfo

func (m *CustomTag_Metadata) GetKind() *v2.MetadataKind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *CustomTag_Metadata) GetMetadataKey() *v2.MetadataKey {
	if m != nil {
		return m.MetadataKey
	}
	return nil
}

func (m *CustomTag_Metadata) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomTag)(nil), "envoy.type.tracing.v2.CustomTag")
	proto.RegisterType((*CustomTag_Literal)(nil), "envoy.type.tracing.v2.CustomTag.Literal")
	proto.RegisterType((*CustomTag_Environment)(nil), "envoy.type.tracing.v2.CustomTag.Environment")
	proto.RegisterType((*CustomTag_Header)(nil), "envoy.type.tracing.v2.CustomTag.Header")
	proto.RegisterType((*CustomTag_Metadata)(nil), "envoy.type.tracing.v2.CustomTag.Metadata")
}

func init() {
	proto.RegisterFile("envoy/type/tracing/v2/custom_tag.proto", fileDescriptor_d6b7ba7857eb6848)
}

var fileDescriptor_d6b7ba7857eb6848 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcb, 0x8a, 0xd4, 0x40,
	0x14, 0x86, 0x53, 0xdd, 0xe9, 0xdb, 0xc9, 0xf4, 0x20, 0x05, 0x62, 0x6c, 0x19, 0x68, 0x6d, 0x2f,
	0x2d, 0x48, 0x02, 0x71, 0xe3, 0x4e, 0x88, 0xb7, 0x80, 0x8a, 0x21, 0x88, 0xdb, 0xe6, 0x38, 0x29,
	0xdb, 0x30, 0x49, 0xa5, 0xad, 0x54, 0x82, 0xd9, 0xf9, 0x4a, 0x6e, 0x5d, 0xb9, 0x9c, 0xad, 0x4f,
	0xe1, 0x2b, 0xc8, 0xac, 0x24, 0x95, 0xcb, 0xb4, 0x76, 0x43, 0x33, 0xbb, 0xca, 0xa9, 0xff, 0xff,
	0x72, 0xea, 0x3f, 0x55, 0x70, 0x9f, 0xf1, 0x22, 0x2d, 0x6d, 0x59, 0x6e, 0x98, 0x2d, 0x05, 0x9e,
	0x46, 0x7c, 0x6d, 0x17, 0x8e, 0x7d, 0x9a, 0x67, 0x32, 0x4d, 0x56, 0x12, 0xd7, 0xd6, 0x46, 0xa4,
	0x32, 0xa5, 0xd7, 0x95, 0xce, 0xaa, 0x74, 0x56, 0xa3, 0xb3, 0x0a, 0x67, 0x76, 0x6f, 0xcb, 0x9e,
	0x30, 0x89, 0x21, 0x4a, 0xac, 0xfc, 0xed, 0xba, 0x76, 0xcf, 0x4e, 0xf2, 0x70, 0x83, 0x36, 0x72,
	0x9e, 0x4a, 0x94, 0x51, 0xca, 0x33, 0x3b, 0x93, 0x28, 0xf3, 0xac, 0xd9, 0xbe, 0x51, 0x60, 0x1c,
	0x85, 0x28, 0x99, 0xdd, 0x2e, 0xea, 0x8d, 0x3b, 0xbf, 0x07, 0x30, 0x79, 0xa6, 0x5a, 0x79, 0x8f,
	0x6b, 0x7a, 0x13, 0xfa, 0x12, 0xd7, 0x26, 0x99, 0x93, 0xe5, 0xc4, 0x1d, 0x5d, 0xb8, 0xba, 0xe8,
	0xcd, 0x49, 0x50, 0xd5, 0xe8, 0x73, 0x18, 0xc5, 0x91, 0x64, 0x02, 0x63, 0xb3, 0x37, 0x27, 0x4b,
	0xc3, 0x59, 0x5a, 0x7b, 0x1b, 0xb6, 0x3a, 0x9a, 0xf5, 0xa6, 0xd6, 0x7b, 0x5a, 0xd0, 0x5a, 0xa9,
	0x0f, 0x06, 0xe3, 0x45, 0x24, 0x52, 0x9e, 0x30, 0x2e, 0xcd, 0xbe, 0x22, 0x3d, 0x3a, 0x48, 0x7a,
	0x71, 0xe9, 0xf1, 0xb4, 0x60, 0x1b, 0x41, 0x7d, 0x38, 0x16, 0xec, 0x4b, 0xce, 0x32, 0xb9, 0xfa,
	0xcc, 0x30, 0x64, 0xc2, 0xd4, 0x15, 0xf4, 0xc1, 0x41, 0xa8, 0xa7, 0xe4, 0x9e, 0x16, 0x4c, 0x1b,
	0x40, 0x5d, 0xa0, 0xaf, 0x60, 0xdc, 0x86, 0x6b, 0x0e, 0x14, 0xeb, 0xe1, 0x41, 0xd6, 0xdb, 0xc6,
	0xe0, 0x69, 0x41, 0x67, 0x9e, 0x2d, 0x61, 0xd4, 0x44, 0x40, 0x4f, 0x60, 0x50, 0x60, 0x9c, 0xb3,
	0xff, 0xa3, 0xad, 0xab, 0xb3, 0x77, 0x60, 0x6c, 0x1d, 0x91, 0xde, 0x02, 0x9d, 0x63, 0xb2, 0x23,
	0x56, 0x45, 0xba, 0x80, 0x69, 0xc8, 0x3e, 0x61, 0x1e, 0xcb, 0x55, 0x8d, 0xac, 0xc6, 0x31, 0x09,
	0x8e, 0x9a, 0xe2, 0x07, 0x05, 0xf4, 0x61, 0xd8, 0x9c, 0xe6, 0xf6, 0x3f, 0xac, 0xe9, 0x85, 0x0b,
	0x62, 0x3c, 0x27, 0x3f, 0x09, 0x39, 0x27, 0xda, 0x55, 0x88, 0xdf, 0x09, 0x8c, 0xdb, 0x53, 0xd2,
	0x27, 0xa0, 0x9f, 0x45, 0x3c, 0x54, 0x50, 0xc3, 0xb9, 0xbb, 0x1d, 0x4f, 0x77, 0x2f, 0x0b, 0xa7,
	0x4b, 0xe5, 0x75, 0xc4, 0xc3, 0x40, 0x39, 0xe8, 0x4b, 0x38, 0x6a, 0x15, 0xab, 0x33, 0x56, 0x36,
	0x77, 0x69, 0x71, 0x90, 0xc0, 0xca, 0xc0, 0x48, 0x2e, 0x3f, 0x76, 0x7b, 0xee, 0xef, 0xf6, 0xec,
	0x1a, 0xa0, 0x57, 0x44, 0xda, 0xff, 0xe3, 0x12, 0xf7, 0xe9, 0x8f, 0x6f, 0xe7, 0xbf, 0x86, 0xbd,
	0x6b, 0x04, 0x16, 0x51, 0x5a, 0xff, 0x6f, 0x23, 0xd2, 0xaf, 0xe5, 0xfe, 0xd9, 0xba, 0xc7, 0xdd,
	0x70, 0xfd, 0xea, 0xa1, 0xf8, 0xe4, 0xe3, 0x50, 0xbd, 0x98, 0xc7, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xea, 0x2c, 0x64, 0xf5, 0xd1, 0x03, 0x00, 0x00,
}
