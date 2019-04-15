// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync/ignore.proto

package sync // import "github.com/mutagen-io/mutagen/pkg/sync"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// IgnoreVCSMode specifies the mode for ignoring VCS directories.
type IgnoreVCSMode int32

const (
	// IgnoreVCSMode_IgnoreVCSDefault represents an unspecified VCS ignore
	// mode. It is not valid for use with Scan. It should be converted to one of
	// the following values based on the desired default behavior.
	IgnoreVCSMode_IgnoreVCSDefault IgnoreVCSMode = 0
	// IgnoreVCSMode_IgnoreVCS indicates that VCS directories should be ignored.
	IgnoreVCSMode_IgnoreVCS IgnoreVCSMode = 1
	// IgnoreVCSMode_PropagateVCS indicates that VCS directories should be
	// propagated.
	IgnoreVCSMode_PropagateVCS IgnoreVCSMode = 2
)

var IgnoreVCSMode_name = map[int32]string{
	0: "IgnoreVCSDefault",
	1: "IgnoreVCS",
	2: "PropagateVCS",
}
var IgnoreVCSMode_value = map[string]int32{
	"IgnoreVCSDefault": 0,
	"IgnoreVCS":        1,
	"PropagateVCS":     2,
}

func (x IgnoreVCSMode) String() string {
	return proto.EnumName(IgnoreVCSMode_name, int32(x))
}
func (IgnoreVCSMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ignore_69b171cb9d287eab, []int{0}
}

func init() {
	proto.RegisterEnum("sync.IgnoreVCSMode", IgnoreVCSMode_name, IgnoreVCSMode_value)
}

func init() { proto.RegisterFile("sync/ignore.proto", fileDescriptor_ignore_69b171cb9d287eab) }

var fileDescriptor_ignore_69b171cb9d287eab = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xae, 0xcc, 0x4b,
	0xd6, 0xcf, 0x4c, 0xcf, 0xcb, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01,
	0x09, 0x69, 0xb9, 0x71, 0xf1, 0x7a, 0x82, 0x45, 0xc3, 0x9c, 0x83, 0x7d, 0xf3, 0x53, 0x52, 0x85,
	0x44, 0xb8, 0x04, 0xe0, 0x02, 0x2e, 0xa9, 0x69, 0x89, 0xa5, 0x39, 0x25, 0x02, 0x0c, 0x42, 0xbc,
	0x5c, 0x9c, 0x70, 0x51, 0x01, 0x46, 0x21, 0x01, 0x2e, 0x9e, 0x80, 0xa2, 0xfc, 0x82, 0xc4, 0xf4,
	0xc4, 0x12, 0xb0, 0x08, 0x93, 0x93, 0x5a, 0x94, 0x4a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x7e, 0x46, 0x62, 0x59, 0x7e, 0xb2, 0x6e, 0x66, 0xbe, 0x7e, 0x6e, 0x69, 0x49,
	0x62, 0x7a, 0x6a, 0x9e, 0x7e, 0x41, 0x76, 0xba, 0x3e, 0xc8, 0xbe, 0x24, 0x36, 0xb0, 0xe5, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0xb9, 0x44, 0x50, 0x91, 0x00, 0x00, 0x00,
}
