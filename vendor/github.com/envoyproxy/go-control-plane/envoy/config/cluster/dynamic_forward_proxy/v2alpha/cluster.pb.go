// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/dynamic_forward_proxy/v2alpha/cluster.proto

package envoy_config_cluster_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2alpha "github.com/envoyproxy/go-control-plane/envoy/config/common/dynamic_forward_proxy/v2alpha"
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

type ClusterConfig struct {
	DnsCacheConfig       *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClusterConfig) Reset()         { *m = ClusterConfig{} }
func (m *ClusterConfig) String() string { return proto.CompactTextString(m) }
func (*ClusterConfig) ProtoMessage()    {}
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb9d327e132c11, []int{0}
}

func (m *ClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterConfig.Unmarshal(m, b)
}
func (m *ClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterConfig.Marshal(b, m, deterministic)
}
func (m *ClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterConfig.Merge(m, src)
}
func (m *ClusterConfig) XXX_Size() int {
	return xxx_messageInfo_ClusterConfig.Size(m)
}
func (m *ClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterConfig proto.InternalMessageInfo

func (m *ClusterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterConfig)(nil), "envoy.config.cluster.dynamic_forward_proxy.v2alpha.ClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/dynamic_forward_proxy/v2alpha/cluster.proto", fileDescriptor_faeb9d327e132c11)
}

var fileDescriptor_faeb9d327e132c11 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0x89, 0x30, 0x91, 0xfa, 0xc2, 0xd8, 0x45, 0xd9, 0x41, 0xc4, 0x93, 0xa7, 0x04, 0x3a,
	0x3f, 0xc0, 0xd6, 0xfa, 0x01, 0x86, 0x5f, 0xa0, 0x3c, 0x36, 0xd9, 0x16, 0x68, 0x9f, 0xa7, 0x24,
	0x69, 0x5d, 0xbf, 0x80, 0x07, 0x11, 0xbc, 0xfa, 0x39, 0x3d, 0x7a, 0x10, 0x69, 0xb3, 0x08, 0x05,
	0x5f, 0xd8, 0x2d, 0xe4, 0x1f, 0x7e, 0xff, 0x97, 0x44, 0x73, 0x85, 0x0d, 0xb5, 0x22, 0x27, 0x5c,
	0xe9, 0xb5, 0xc8, 0x8b, 0xda, 0x3a, 0x65, 0x84, 0x6c, 0x11, 0x4a, 0x9d, 0x67, 0x2b, 0x32, 0x8f,
	0x60, 0x64, 0x56, 0x19, 0xda, 0xb6, 0xa2, 0x89, 0xa1, 0xa8, 0x36, 0x10, 0x5e, 0xf1, 0xca, 0x90,
	0xa3, 0x49, 0xdc, 0x13, 0xb8, 0x27, 0xf0, 0xa0, 0xfd, 0x48, 0xe0, 0x3b, 0xc2, 0x74, 0x31, 0x74,
	0xa5, 0xb2, 0x24, 0xfc, 0xc7, 0x54, 0xa2, 0xcd, 0x72, 0xc8, 0x37, 0xca, 0xdb, 0x4e, 0x2f, 0x6b,
	0x59, 0x81, 0x00, 0x44, 0x72, 0xe0, 0x34, 0xa1, 0x15, 0xa5, 0x5e, 0x1b, 0x70, 0x41, 0x3f, 0x6f,
	0xa0, 0xd0, 0x12, 0x9c, 0x12, 0xe1, 0xe0, 0x85, 0xeb, 0x27, 0x16, 0x9d, 0xa6, 0x3e, 0x65, 0xda,
	0xfb, 0x4f, 0xea, 0x68, 0xfc, 0x4d, 0xcf, 0x7c, 0xa6, 0x0b, 0x76, 0xc5, 0x6e, 0x8e, 0xe3, 0x05,
	0x1f, 0x96, 0xeb, 0x83, 0xfe, 0xdd, 0x8d, 0xdf, 0xa1, 0x4d, 0x3b, 0x92, 0x87, 0x27, 0x47, 0x1f,
	0xc9, 0xe8, 0x99, 0x1d, 0x8c, 0xd9, 0xfd, 0x99, 0x1c, 0x2a, 0x2f, 0xec, 0xfd, 0xed, 0xf3, 0x75,
	0x74, 0x1b, 0x16, 0x54, 0x5b, 0xa7, 0xd0, 0x76, 0x55, 0xc2, 0x8a, 0xf6, 0x37, 0xab, 0x59, 0x34,
	0xd7, 0xe4, 0xb3, 0xf9, 0x9b, 0xfd, 0xff, 0x20, 0x39, 0xd9, 0x8d, 0xb0, 0xec, 0x56, 0x59, 0xb2,
	0x87, 0xc3, 0x7e, 0x9e, 0xd9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x63, 0x6d, 0x86, 0x12,
	0x02, 0x00, 0x00,
}
