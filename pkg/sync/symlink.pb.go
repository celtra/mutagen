// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync/symlink.proto

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

// SymlinkMode specifies the mode for handling the propagation of symlinks.
type SymlinkMode int32

const (
	// SymlinkMode_SymlinkDefault represents an unspecified symlink mode. It is
	// not valid for use with Scan or Transition. It should be converted to one
	// of the following values based on the desired default behavior.
	SymlinkMode_SymlinkDefault SymlinkMode = 0
	// SymlinkMode_SymlinkIgnore specifies that all symlinks should be ignored.
	SymlinkMode_SymlinkIgnore SymlinkMode = 1
	// SymlinkMode_SymlinkPortable specifies that only portable symlinks should
	// be synchronized. If a symlink is found during a scan operation that it is
	// not portable, it halts the scan and synchronization. The reason for this
	// is that it can't simply be ignored/unignored as desired without breaking
	// the three-way merge.
	SymlinkMode_SymlinkPortable SymlinkMode = 2
	// SymlinkMode_SymlinkPOSIXRaw specifies that symlinks should be propagated
	// in their raw form. It is only valid on POSIX systems and only makes sense
	// in the context of POSIX-to-POSIX synchronization.
	SymlinkMode_SymlinkPOSIXRaw SymlinkMode = 3
)

var SymlinkMode_name = map[int32]string{
	0: "SymlinkDefault",
	1: "SymlinkIgnore",
	2: "SymlinkPortable",
	3: "SymlinkPOSIXRaw",
}
var SymlinkMode_value = map[string]int32{
	"SymlinkDefault":  0,
	"SymlinkIgnore":   1,
	"SymlinkPortable": 2,
	"SymlinkPOSIXRaw": 3,
}

func (x SymlinkMode) String() string {
	return proto.EnumName(SymlinkMode_name, int32(x))
}
func (SymlinkMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_symlink_d2d4a30957604e58, []int{0}
}

func init() {
	proto.RegisterEnum("sync.SymlinkMode", SymlinkMode_name, SymlinkMode_value)
}

func init() { proto.RegisterFile("sync/symlink.proto", fileDescriptor_symlink_d2d4a30957604e58) }

var fileDescriptor_symlink_d2d4a30957604e58 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xae, 0xcc, 0x4b,
	0xd6, 0x2f, 0xae, 0xcc, 0xcd, 0xc9, 0xcc, 0xcb, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x89, 0x69, 0xc5, 0x71, 0x71, 0x07, 0x43, 0x84, 0x7d, 0xf3, 0x53, 0x52, 0x85, 0x84, 0xb8,
	0xf8, 0xa0, 0x5c, 0x97, 0xd4, 0xb4, 0xc4, 0xd2, 0x9c, 0x12, 0x01, 0x06, 0x21, 0x41, 0x2e, 0x5e,
	0xa8, 0x98, 0x67, 0x7a, 0x5e, 0x7e, 0x51, 0xaa, 0x00, 0xa3, 0x90, 0x30, 0x17, 0x3f, 0x54, 0x28,
	0x20, 0xbf, 0xa8, 0x24, 0x31, 0x29, 0x27, 0x55, 0x80, 0x09, 0x59, 0xd0, 0x3f, 0xd8, 0x33, 0x22,
	0x28, 0xb1, 0x5c, 0x80, 0xd9, 0x49, 0x2d, 0x4a, 0x25, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0x3f, 0x23, 0xb1, 0x2c, 0x3f, 0x59, 0x37, 0x33, 0x5f, 0x3f, 0xb7, 0xb4, 0x24,
	0x31, 0x3d, 0x35, 0x4f, 0xbf, 0x20, 0x3b, 0x5d, 0x1f, 0xe4, 0x8e, 0x24, 0x36, 0xb0, 0xa3, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x07, 0x1b, 0xc8, 0xaa, 0x00, 0x00, 0x00,
}
