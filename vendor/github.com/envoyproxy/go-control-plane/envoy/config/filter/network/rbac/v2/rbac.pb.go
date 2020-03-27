// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/rbac/v2/rbac.proto

package envoy_config_filter_network_rbac_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v2"
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

type RBAC_EnforcementType int32

const (
	RBAC_ONE_TIME_ON_FIRST_BYTE RBAC_EnforcementType = 0
	RBAC_CONTINUOUS             RBAC_EnforcementType = 1
)

var RBAC_EnforcementType_name = map[int32]string{
	0: "ONE_TIME_ON_FIRST_BYTE",
	1: "CONTINUOUS",
}

var RBAC_EnforcementType_value = map[string]int32{
	"ONE_TIME_ON_FIRST_BYTE": 0,
	"CONTINUOUS":             1,
}

func (x RBAC_EnforcementType) String() string {
	return proto.EnumName(RBAC_EnforcementType_name, int32(x))
}

func (RBAC_EnforcementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ec60cc393c44598, []int{0, 0}
}

type RBAC struct {
	Rules                *v2.RBAC             `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	ShadowRules          *v2.RBAC             `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	StatPrefix           string               `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	EnforcementType      RBAC_EnforcementType `protobuf:"varint,4,opt,name=enforcement_type,json=enforcementType,proto3,enum=envoy.config.filter.network.rbac.v2.RBAC_EnforcementType" json:"enforcement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec60cc393c44598, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v2.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v2.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

func (m *RBAC) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RBAC) GetEnforcementType() RBAC_EnforcementType {
	if m != nil {
		return m.EnforcementType
	}
	return RBAC_ONE_TIME_ON_FIRST_BYTE
}

func init() {
	proto.RegisterEnum("envoy.config.filter.network.rbac.v2.RBAC_EnforcementType", RBAC_EnforcementType_name, RBAC_EnforcementType_value)
	proto.RegisterType((*RBAC)(nil), "envoy.config.filter.network.rbac.v2.RBAC")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/rbac/v2/rbac.proto", fileDescriptor_8ec60cc393c44598)
}

var fileDescriptor_8ec60cc393c44598 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x6e, 0xe2, 0x30,
	0x18, 0xc7, 0x2f, 0x1c, 0xdc, 0x09, 0x73, 0x02, 0x94, 0xe1, 0x0e, 0x65, 0xb8, 0x8b, 0xb8, 0x25,
	0xea, 0x60, 0xb7, 0x61, 0xea, 0xc0, 0xd0, 0xa0, 0x54, 0x62, 0x68, 0x82, 0x42, 0x18, 0x3a, 0x45,
	0x26, 0x71, 0x68, 0x54, 0xb0, 0x23, 0xc7, 0x04, 0xf2, 0x14, 0x5d, 0xfb, 0x44, 0x7d, 0xa0, 0x8e,
	0x1d, 0xaa, 0x2a, 0x71, 0xaa, 0x0a, 0x3a, 0x94, 0xc9, 0x96, 0xff, 0xbf, 0xef, 0xf3, 0xcf, 0xfe,
	0x00, 0x24, 0x34, 0x67, 0x05, 0x0a, 0x19, 0x8d, 0x93, 0x15, 0x8a, 0x93, 0xb5, 0x20, 0x1c, 0x51,
	0x22, 0x76, 0x8c, 0xdf, 0x23, 0xbe, 0xc4, 0x21, 0xca, 0xcd, 0x6a, 0x85, 0x29, 0x67, 0x82, 0xa9,
	0xff, 0x2b, 0x1e, 0x4a, 0x1e, 0x4a, 0x1e, 0xd6, 0x3c, 0xac, 0xb8, 0xdc, 0xd4, 0xfe, 0x1d, 0x34,
	0xfd, 0xdc, 0x45, 0xfb, 0xbb, 0x8d, 0x52, 0x8c, 0x30, 0xa5, 0x4c, 0x60, 0x91, 0x30, 0x9a, 0xa1,
	0x4d, 0xb2, 0xe2, 0x58, 0x90, 0x3a, 0xff, 0x93, 0xe3, 0x75, 0x12, 0x61, 0x41, 0xd0, 0xfb, 0x46,
	0x06, 0xc3, 0xa7, 0x06, 0x68, 0x7a, 0xd6, 0xd5, 0x44, 0x3d, 0x07, 0x2d, 0xbe, 0x5d, 0x93, 0x6c,
	0xa0, 0xe8, 0x8a, 0xd1, 0x31, 0x35, 0x78, 0xe0, 0x55, 0x8b, 0xc0, 0x12, 0xf5, 0x24, 0xa8, 0x8e,
	0xc1, 0xaf, 0xec, 0x0e, 0x47, 0x6c, 0x17, 0xc8, 0xc2, 0xc6, 0x97, 0x85, 0x1d, 0xc9, 0x7b, 0x55,
	0xb9, 0x01, 0x3a, 0x99, 0xc0, 0x22, 0x48, 0x39, 0x89, 0x93, 0xfd, 0xe0, 0xbb, 0xae, 0x18, 0x6d,
	0xeb, 0xe7, 0x8b, 0xd5, 0xe4, 0x0d, 0x5d, 0xf1, 0x40, 0x99, 0xcd, 0xaa, 0x48, 0x8d, 0x40, 0x9f,
	0xd0, 0x98, 0xf1, 0x90, 0x6c, 0x08, 0x15, 0x81, 0x28, 0x52, 0x32, 0x68, 0xea, 0x8a, 0xd1, 0x35,
	0x2f, 0xe1, 0x09, 0xbf, 0x57, 0xdd, 0x0d, 0xed, 0x8f, 0x0e, 0x7e, 0x91, 0x12, 0xaf, 0x47, 0x0e,
	0x0f, 0x86, 0x63, 0xd0, 0x3b, 0x62, 0x54, 0x0d, 0xfc, 0x76, 0x1d, 0x3b, 0xf0, 0xa7, 0x37, 0x76,
	0xe0, 0x3a, 0xc1, 0xf5, 0xd4, 0x9b, 0xfb, 0x81, 0x75, 0xeb, 0xdb, 0xfd, 0x6f, 0x6a, 0x17, 0x80,
	0x89, 0xeb, 0xf8, 0x53, 0x67, 0xe1, 0x2e, 0xe6, 0x7d, 0xc5, 0x4a, 0x9f, 0x1f, 0x5f, 0x1f, 0x5a,
	0x67, 0xaa, 0x21, 0x8d, 0xc8, 0x5e, 0x10, 0x9a, 0x95, 0x93, 0xa8, 0xad, 0xb2, 0x23, 0xad, 0x11,
	0xb8, 0x48, 0x98, 0xd4, 0x4f, 0x39, 0xdb, 0x17, 0xa7, 0xbc, 0xc4, 0x6a, 0x7b, 0x4b, 0x1c, 0xce,
	0xca, 0xc1, 0xcd, 0x94, 0xe5, 0x8f, 0x6a, 0x82, 0xa3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92,
	0xb3, 0xaf, 0x0f, 0x72, 0x02, 0x00, 0x00,
}
