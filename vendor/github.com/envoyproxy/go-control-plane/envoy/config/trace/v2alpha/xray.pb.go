// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2alpha/xray.proto

package envoy_config_trace_v2alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
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

type XRayConfig struct {
	DaemonEndpoint       *core.SocketAddress `protobuf:"bytes,1,opt,name=daemon_endpoint,json=daemonEndpoint,proto3" json:"daemon_endpoint,omitempty"`
	SegmentName          string              `protobuf:"bytes,2,opt,name=segment_name,json=segmentName,proto3" json:"segment_name,omitempty"`
	SamplingRuleManifest *core.DataSource    `protobuf:"bytes,3,opt,name=sampling_rule_manifest,json=samplingRuleManifest,proto3" json:"sampling_rule_manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *XRayConfig) Reset()         { *m = XRayConfig{} }
func (m *XRayConfig) String() string { return proto.CompactTextString(m) }
func (*XRayConfig) ProtoMessage()    {}
func (*XRayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c272343b96f21fb4, []int{0}
}

func (m *XRayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRayConfig.Unmarshal(m, b)
}
func (m *XRayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRayConfig.Marshal(b, m, deterministic)
}
func (m *XRayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRayConfig.Merge(m, src)
}
func (m *XRayConfig) XXX_Size() int {
	return xxx_messageInfo_XRayConfig.Size(m)
}
func (m *XRayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_XRayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_XRayConfig proto.InternalMessageInfo

func (m *XRayConfig) GetDaemonEndpoint() *core.SocketAddress {
	if m != nil {
		return m.DaemonEndpoint
	}
	return nil
}

func (m *XRayConfig) GetSegmentName() string {
	if m != nil {
		return m.SegmentName
	}
	return ""
}

func (m *XRayConfig) GetSamplingRuleManifest() *core.DataSource {
	if m != nil {
		return m.SamplingRuleManifest
	}
	return nil
}

func init() {
	proto.RegisterType((*XRayConfig)(nil), "envoy.config.trace.v2alpha.XRayConfig")
}

func init() {
	proto.RegisterFile("envoy/config/trace/v2alpha/xray.proto", fileDescriptor_c272343b96f21fb4)
}

var fileDescriptor_c272343b96f21fb4 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x82, 0xd0, 0xad, 0x28, 0x04, 0xd1, 0x10, 0x14, 0xa3, 0x20, 0xe4, 0xb4, 0x0b,
	0xf1, 0xe0, 0xd9, 0xaa, 0x07, 0x0f, 0x4a, 0x49, 0x2e, 0xbd, 0x85, 0x69, 0x32, 0x8d, 0x8b, 0xd9,
	0x0f, 0x76, 0x37, 0xa1, 0xf9, 0x95, 0xfe, 0x25, 0x69, 0x36, 0x3d, 0x69, 0x6f, 0xcb, 0xce, 0xf3,
	0xcc, 0x3b, 0xbc, 0xe4, 0x01, 0x65, 0xaf, 0x06, 0x56, 0x29, 0xb9, 0xe1, 0x0d, 0x73, 0x06, 0x2a,
	0x64, 0x7d, 0x06, 0xad, 0xfe, 0x02, 0xb6, 0x35, 0x30, 0x50, 0x6d, 0x94, 0x53, 0x61, 0x3c, 0x62,
	0xd4, 0x63, 0x74, 0xc4, 0xe8, 0x84, 0xc5, 0xb7, 0x7e, 0x05, 0x68, 0xce, 0xfa, 0x8c, 0x55, 0xca,
	0x20, 0x83, 0xba, 0x36, 0x68, 0xad, 0x97, 0xe3, 0xeb, 0xbf, 0xc0, 0x1a, 0x2c, 0x4e, 0xd3, 0xab,
	0x1e, 0x5a, 0x5e, 0x83, 0x43, 0xb6, 0x7f, 0xf8, 0xc1, 0xfd, 0x4f, 0x40, 0xc8, 0x2a, 0x87, 0xe1,
	0x65, 0x0c, 0x0d, 0xdf, 0xc9, 0x79, 0x0d, 0x28, 0x94, 0x2c, 0x51, 0xd6, 0x5a, 0x71, 0xe9, 0xa2,
	0x20, 0x09, 0xd2, 0x79, 0x96, 0x50, 0x7f, 0x1c, 0x68, 0x4e, 0xfb, 0x8c, 0xee, 0xf6, 0xd3, 0x42,
	0x55, 0xdf, 0xe8, 0x9e, 0xfd, 0x19, 0xf9, 0x99, 0x17, 0xdf, 0x26, 0x2f, 0xbc, 0x23, 0xa7, 0x16,
	0x1b, 0x81, 0xd2, 0x95, 0x12, 0x04, 0x46, 0x47, 0x49, 0x90, 0xce, 0xf2, 0xf9, 0xf4, 0xf7, 0x09,
	0x02, 0xc3, 0x82, 0x5c, 0x5a, 0x10, 0xba, 0xe5, 0xb2, 0x29, 0x4d, 0xd7, 0x62, 0x29, 0x40, 0xf2,
	0x0d, 0x5a, 0x17, 0x1d, 0x8f, 0xa1, 0x37, 0xff, 0x84, 0xbe, 0x82, 0x83, 0x42, 0x75, 0xa6, 0xc2,
	0xfc, 0x62, 0x2f, 0xe7, 0x5d, 0x8b, 0x1f, 0x93, 0xba, 0x78, 0x22, 0x29, 0x57, 0x5e, 0xd4, 0x46,
	0x6d, 0x07, 0x7a, 0xb8, 0xd5, 0xc5, 0x6c, 0x65, 0x60, 0x58, 0xee, 0x8a, 0x58, 0x06, 0xeb, 0x93,
	0xb1, 0x91, 0xc7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x4c, 0x00, 0x80, 0xae, 0x01, 0x00,
	0x00,
}
