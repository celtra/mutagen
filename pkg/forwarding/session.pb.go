// Code generated by protoc-gen-go. DO NOT EDIT.
// source: forwarding/session.proto

package forwarding

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	url "github.com/mutagen-io/mutagen/pkg/url"
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

type Session struct {
	// Identifier is the (unique) session identifier. It is static. It cannot be
	// empty.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Version is the session version. It is static.
	Version Version `protobuf:"varint,2,opt,name=version,proto3,enum=forwarding.Version" json:"version,omitempty"`
	// CreationTime is the creation time of the session. It is static. It cannot
	// be nil.
	CreationTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	// CreatingVersionMajor is the major version component of the version of
	// Mutagen which created the session. It is static.
	CreatingVersionMajor uint32 `protobuf:"varint,4,opt,name=creatingVersionMajor,proto3" json:"creatingVersionMajor,omitempty"`
	// CreatingVersionMinor is the minor version component of the version of
	// Mutagen which created the session. It is static.
	CreatingVersionMinor uint32 `protobuf:"varint,5,opt,name=creatingVersionMinor,proto3" json:"creatingVersionMinor,omitempty"`
	// CreatingVersionPatch is the patch version component of the version of
	// Mutagen which created the session. It is static.
	CreatingVersionPatch uint32 `protobuf:"varint,6,opt,name=creatingVersionPatch,proto3" json:"creatingVersionPatch,omitempty"`
	// Source is the source endpoint URL. It is static. It cannot be nil.
	Source *url.URL `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	// Destination is the destination endpoint URL. It is static. It cannot be
	// nil.
	Destination *url.URL `protobuf:"bytes,8,opt,name=destination,proto3" json:"destination,omitempty"`
	// Configuration is the flattened session configuration. It is static. It
	// cannot be nil.
	Configuration *Configuration `protobuf:"bytes,9,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// ConfigurationSource are the source-specific session configuration
	// overrides. It is static.
	ConfigurationSource *Configuration `protobuf:"bytes,10,opt,name=configurationSource,proto3" json:"configurationSource,omitempty"`
	// ConfigurationDestination are the destination-specific session
	// configuration overrides. It is static.
	ConfigurationDestination *Configuration `protobuf:"bytes,11,opt,name=configurationDestination,proto3" json:"configurationDestination,omitempty"`
	// Name is a user-friendly name for the session. It may be empty and is not
	// guaranteed to be unique across all sessions. It is only used as a simpler
	// handle for specifying sessions. It is static.
	Name string `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	// Labels are the session labels. They are static.
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Paused indicates whether or not the session is marked as paused.
	Paused               bool     `protobuf:"varint,14,opt,name=paused,proto3" json:"paused,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_017cb915c0809617, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Session) GetVersion() Version {
	if m != nil {
		return m.Version
	}
	return Version_Invalid
}

func (m *Session) GetCreationTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Session) GetCreatingVersionMajor() uint32 {
	if m != nil {
		return m.CreatingVersionMajor
	}
	return 0
}

func (m *Session) GetCreatingVersionMinor() uint32 {
	if m != nil {
		return m.CreatingVersionMinor
	}
	return 0
}

func (m *Session) GetCreatingVersionPatch() uint32 {
	if m != nil {
		return m.CreatingVersionPatch
	}
	return 0
}

func (m *Session) GetSource() *url.URL {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Session) GetDestination() *url.URL {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *Session) GetConfiguration() *Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *Session) GetConfigurationSource() *Configuration {
	if m != nil {
		return m.ConfigurationSource
	}
	return nil
}

func (m *Session) GetConfigurationDestination() *Configuration {
	if m != nil {
		return m.ConfigurationDestination
	}
	return nil
}

func (m *Session) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Session) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Session) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func init() {
	proto.RegisterType((*Session)(nil), "forwarding.Session")
	proto.RegisterMapType((map[string]string)(nil), "forwarding.Session.LabelsEntry")
}

func init() { proto.RegisterFile("forwarding/session.proto", fileDescriptor_017cb915c0809617) }

var fileDescriptor_017cb915c0809617 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6f, 0xd3, 0x30,
	0x14, 0x57, 0xd6, 0x2d, 0x6d, 0x5f, 0xd6, 0x09, 0x79, 0x13, 0x32, 0x39, 0x6c, 0x11, 0xa7, 0x08,
	0x31, 0x47, 0x2a, 0x07, 0x3e, 0x0e, 0x20, 0xf1, 0x71, 0x62, 0x48, 0xc8, 0xdb, 0x38, 0x70, 0x73,
	0x53, 0x37, 0x33, 0x4b, 0xec, 0xca, 0xb1, 0x87, 0xf6, 0x17, 0xf0, 0x6f, 0xa3, 0xda, 0x29, 0x75,
	0x50, 0xa6, 0xdd, 0xfc, 0xde, 0xef, 0xe3, 0xbd, 0x5f, 0xf2, 0x00, 0xaf, 0x94, 0xfe, 0xcd, 0xf4,
	0x52, 0xc8, 0xaa, 0x68, 0x79, 0xdb, 0x0a, 0x25, 0xc9, 0x5a, 0x2b, 0xa3, 0x10, 0xec, 0x90, 0xf4,
	0xac, 0x52, 0xaa, 0xaa, 0x79, 0xe1, 0x90, 0x85, 0x5d, 0x15, 0x46, 0x34, 0xbc, 0x35, 0xac, 0x59,
	0x7b, 0x72, 0x7a, 0x1a, 0xd8, 0x94, 0x4a, 0xae, 0x44, 0x65, 0x35, 0x33, 0xff, 0xcc, 0xd2, 0x70,
	0xcc, 0x1d, 0xd7, 0xbb, 0x31, 0xe9, 0xcc, 0xea, 0xba, 0xb0, 0xba, 0xf6, 0xe5, 0xf3, 0x3f, 0x31,
	0x8c, 0x2f, 0xfd, 0x1e, 0xe8, 0x14, 0x40, 0x2c, 0xb9, 0x34, 0x62, 0x25, 0xb8, 0xc6, 0x51, 0x16,
	0xe5, 0x53, 0x1a, 0x74, 0xd0, 0x39, 0x8c, 0x3b, 0x2f, 0xbc, 0x97, 0x45, 0xf9, 0xd1, 0xfc, 0x98,
	0xec, 0xc6, 0x90, 0x1f, 0x1e, 0xa2, 0x5b, 0x0e, 0x7a, 0x0f, 0x87, 0xa5, 0xe6, 0x6e, 0xab, 0x2b,
	0xd1, 0x70, 0x3c, 0xca, 0xa2, 0x3c, 0x99, 0xa7, 0xc4, 0x67, 0x23, 0xdb, 0x6c, 0xe4, 0x6a, 0x9b,
	0x8d, 0xf6, 0xf8, 0x68, 0x0e, 0x27, 0xbe, 0x96, 0x55, 0xe7, 0xfd, 0x8d, 0xfd, 0x52, 0x1a, 0xef,
	0x67, 0x51, 0x3e, 0xa3, 0x83, 0xd8, 0x90, 0x46, 0x48, 0xa5, 0xf1, 0xc1, 0xb0, 0x66, 0x83, 0x0d,
	0x68, 0xbe, 0x33, 0x53, 0xde, 0xe0, 0x78, 0x50, 0xe3, 0x30, 0x94, 0x41, 0xdc, 0x2a, 0xab, 0x4b,
	0x8e, 0xc7, 0x2e, 0xd5, 0x84, 0x6c, 0x3e, 0xe9, 0x35, 0xbd, 0xa0, 0x5d, 0x1f, 0xbd, 0x80, 0x64,
	0xc9, 0x5b, 0x23, 0xa4, 0x0b, 0x84, 0x27, 0xff, 0xd1, 0x42, 0x10, 0x7d, 0x80, 0x59, 0xef, 0x27,
	0xe2, 0xa9, 0x63, 0x3f, 0x0b, 0x3f, 0xef, 0xa7, 0x90, 0x40, 0xfb, 0x7c, 0xf4, 0x15, 0x8e, 0x7b,
	0x8d, 0x4b, 0xbf, 0x1b, 0x3c, 0x66, 0x33, 0xa4, 0x42, 0xd7, 0x80, 0x7b, 0xed, 0xcf, 0x41, 0x8c,
	0xe4, 0x31, 0xc7, 0x07, 0xa5, 0x08, 0xc1, 0xbe, 0x64, 0x0d, 0xc7, 0x87, 0xee, 0xae, 0xdc, 0x1b,
	0xbd, 0x86, 0xb8, 0x66, 0x0b, 0x5e, 0xb7, 0x78, 0x96, 0x8d, 0xf2, 0x64, 0x7e, 0x16, 0x1a, 0x77,
	0x67, 0x49, 0x2e, 0x1c, 0xe3, 0x8b, 0x34, 0xfa, 0x9e, 0x76, 0x74, 0xf4, 0x14, 0xe2, 0x35, 0xb3,
	0x2d, 0x5f, 0xe2, 0xa3, 0x2c, 0xca, 0x27, 0xb4, 0xab, 0xd2, 0xb7, 0x90, 0x04, 0x74, 0xf4, 0x04,
	0x46, 0xb7, 0xfc, 0xbe, 0x3b, 0xe5, 0xcd, 0x13, 0x9d, 0xc0, 0xc1, 0x1d, 0xab, 0x2d, 0x77, 0x17,
	0x3c, 0xa5, 0xbe, 0x78, 0xb7, 0xf7, 0x26, 0xfa, 0x48, 0x7e, 0xbe, 0xac, 0x84, 0xb9, 0xb1, 0x0b,
	0x52, 0xaa, 0xa6, 0x68, 0xac, 0x61, 0x15, 0x97, 0xe7, 0x42, 0x6d, 0x9f, 0xc5, 0xfa, 0xb6, 0x2a,
	0x76, 0xeb, 0x2d, 0x62, 0x77, 0xc0, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xbc, 0xdf,
	0x26, 0xd2, 0x03, 0x00, 0x00,
}
