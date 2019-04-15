// Code generated by protoc-gen-go. DO NOT EDIT.
// source: url/url.proto

package url

import (
	fmt "fmt"
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

// Kind indicates the kind of a URL.
type Kind int32

const (
	// Synchronization indicates a synchronization URL.
	Kind_Synchronization Kind = 0
	// Forwarding indicates a forwarding URL.
	Kind_Forwarding Kind = 1
)

var Kind_name = map[int32]string{
	0: "Synchronization",
	1: "Forwarding",
}

var Kind_value = map[string]int32{
	"Synchronization": 0,
	"Forwarding":      1,
}

func (x Kind) String() string {
	return proto.EnumName(Kind_name, int32(x))
}

func (Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce31eacd751d7393, []int{0}
}

// Protocol indicates a location type.
type Protocol int32

const (
	// Local indicates that the resource is on the local system.
	Protocol_Local Protocol = 0
	// SSH indicates that the resource is accessible via SSH.
	Protocol_SSH Protocol = 1
	// Docker indicates that the resource is inside a Docker container.
	Protocol_Docker  Protocol = 11
	Protocol_Kubectl Protocol = 12
)

var Protocol_name = map[int32]string{
	0:  "Local",
	1:  "SSH",
	11: "Docker",
	12: "Kubectl",
}

var Protocol_value = map[string]int32{
	"Local":   0,
	"SSH":     1,
	"Docker":  11,
	"Kubectl": 12,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}

func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce31eacd751d7393, []int{1}
}

// URL represents a pointer to a resource.
type URL struct {
	// Kind indicates the URL kind.
	// NOTE: This field number is out of order for historical reasons.
	Kind Kind `protobuf:"varint,7,opt,name=kind,proto3,enum=url.Kind" json:"kind,omitempty"`
	// Protocol indicates a location type.
	Protocol Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=url.Protocol" json:"protocol,omitempty"`
	// User is the user under which a resource should be accessed.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// Host is protocol-specific, but generally indicates the location of the
	// remote.
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// Port indicates a TCP port via which to access the remote location, if
	// applicable.
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// Path indicates the path of a resource.
	Path string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	// Environment is used to capture environment variable information (if
	// necessary) for transports which operate by executing a command.
	Environment          map[string]string `protobuf:"bytes,6,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *URL) Reset()         { *m = URL{} }
func (m *URL) String() string { return proto.CompactTextString(m) }
func (*URL) ProtoMessage()    {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce31eacd751d7393, []int{0}
}

func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (m *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(m, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetKind() Kind {
	if m != nil {
		return m.Kind
	}
	return Kind_Synchronization
}

func (m *URL) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_Local
}

func (m *URL) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *URL) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *URL) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *URL) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *URL) GetEnvironment() map[string]string {
	if m != nil {
		return m.Environment
	}
	return nil
}

func init() {
	proto.RegisterEnum("url.Kind", Kind_name, Kind_value)
	proto.RegisterEnum("url.Protocol", Protocol_name, Protocol_value)
	proto.RegisterType((*URL)(nil), "url.URL")
	proto.RegisterMapType((map[string]string)(nil), "url.URL.EnvironmentEntry")
}

func init() { proto.RegisterFile("url/url.proto", fileDescriptor_ce31eacd751d7393) }

var fileDescriptor_ce31eacd751d7393 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x4b, 0xf3, 0x30,
	0x1c, 0xc7, 0xd7, 0xb5, 0xfb, 0xd3, 0x5f, 0x9f, 0xed, 0x09, 0xd1, 0x43, 0x15, 0x84, 0x22, 0x88,
	0x75, 0x62, 0x07, 0xf3, 0xa0, 0x28, 0x78, 0x10, 0x27, 0xc2, 0x76, 0x90, 0x8c, 0x5d, 0xbc, 0x75,
	0x5d, 0x68, 0x43, 0xbb, 0x64, 0x64, 0xc9, 0x64, 0xbe, 0x16, 0x5f, 0xac, 0x24, 0xdb, 0x54, 0xbc,
	0x7d, 0xf3, 0xf9, 0x7d, 0x42, 0x92, 0x6f, 0xa0, 0xa3, 0x65, 0xd5, 0xd7, 0xb2, 0x4a, 0x96, 0x52,
	0x28, 0x81, 0x5d, 0x2d, 0xab, 0xd3, 0xcf, 0x3a, 0xb8, 0x53, 0x32, 0xc6, 0x27, 0xe0, 0x95, 0x8c,
	0xcf, 0xc3, 0x56, 0xe4, 0xc4, 0xdd, 0x81, 0x9f, 0x18, 0x6d, 0xc4, 0xf8, 0x9c, 0x58, 0x8c, 0x2f,
	0xa0, 0x6d, 0x37, 0x65, 0xa2, 0x0a, 0x1d, 0xab, 0x74, 0xac, 0xf2, 0xba, 0x83, 0xe4, 0x7b, 0x8c,
	0x31, 0x78, 0x7a, 0x45, 0x65, 0x58, 0x8f, 0x9c, 0xd8, 0x27, 0x36, 0x1b, 0x56, 0x88, 0x95, 0x0a,
	0xdd, 0x2d, 0x33, 0xd9, 0xb0, 0xa5, 0x90, 0x2a, 0xf4, 0x22, 0x27, 0xee, 0x10, 0x9b, 0x2d, 0x4b,
	0x55, 0x11, 0x36, 0xb6, 0x9e, 0xc9, 0xf8, 0x1e, 0x02, 0xca, 0xd7, 0x4c, 0x0a, 0xbe, 0xa0, 0x5c,
	0x85, 0xcd, 0xc8, 0x8d, 0x83, 0xc1, 0x91, 0x3d, 0x7d, 0x4a, 0xc6, 0xc9, 0xf0, 0x67, 0x36, 0xe4,
	0x4a, 0x6e, 0xc8, 0x6f, 0xfb, 0xf8, 0x01, 0xd0, 0x5f, 0x01, 0x23, 0x70, 0x4b, 0xba, 0xb1, 0xcf,
	0xf0, 0x89, 0x89, 0xf8, 0x10, 0x1a, 0xeb, 0xb4, 0xd2, 0x74, 0x77, 0xe7, 0xed, 0xe2, 0xae, 0x7e,
	0xeb, 0xf4, 0x2e, 0xc1, 0x33, 0x2d, 0xe0, 0x03, 0xf8, 0x3f, 0xd9, 0xf0, 0xac, 0x90, 0x82, 0xb3,
	0x8f, 0x54, 0x31, 0xc1, 0x51, 0x0d, 0x77, 0x01, 0x9e, 0x85, 0x7c, 0x4f, 0xe5, 0x9c, 0xf1, 0x1c,
	0x39, 0xbd, 0x1b, 0x68, 0xef, 0xfb, 0xc0, 0x3e, 0x34, 0xc6, 0x22, 0x4b, 0x2b, 0x54, 0xc3, 0x2d,
	0x70, 0x27, 0x93, 0x17, 0xe4, 0x60, 0x80, 0xe6, 0x93, 0xc8, 0x4a, 0x2a, 0x51, 0x80, 0x03, 0x68,
	0x8d, 0xf4, 0x8c, 0x66, 0xaa, 0x42, 0xff, 0x1e, 0xcf, 0xdf, 0xce, 0x72, 0xa6, 0x0a, 0x3d, 0x4b,
	0x32, 0xb1, 0xe8, 0x2f, 0xb4, 0x4a, 0x73, 0xca, 0xaf, 0x98, 0xd8, 0xc7, 0xfe, 0xb2, 0xcc, 0xcd,
	0xc7, 0xcd, 0x9a, 0xb6, 0xe5, 0xeb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xeb, 0xdb, 0xd8,
	0xca, 0x01, 0x00, 0x00,
}
