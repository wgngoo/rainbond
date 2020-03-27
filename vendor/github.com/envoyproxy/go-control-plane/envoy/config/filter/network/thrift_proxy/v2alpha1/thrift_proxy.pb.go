// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto

package envoy_config_filter_network_thrift_proxy_v2alpha1

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type TransportType int32

const (
	TransportType_AUTO_TRANSPORT TransportType = 0
	TransportType_FRAMED         TransportType = 1
	TransportType_UNFRAMED       TransportType = 2
	TransportType_HEADER         TransportType = 3
)

var TransportType_name = map[int32]string{
	0: "AUTO_TRANSPORT",
	1: "FRAMED",
	2: "UNFRAMED",
	3: "HEADER",
}

var TransportType_value = map[string]int32{
	"AUTO_TRANSPORT": 0,
	"FRAMED":         1,
	"UNFRAMED":       2,
	"HEADER":         3,
}

func (x TransportType) String() string {
	return proto.EnumName(TransportType_name, int32(x))
}

func (TransportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{0}
}

type ProtocolType int32

const (
	ProtocolType_AUTO_PROTOCOL ProtocolType = 0
	ProtocolType_BINARY        ProtocolType = 1
	ProtocolType_LAX_BINARY    ProtocolType = 2
	ProtocolType_COMPACT       ProtocolType = 3
	ProtocolType_TWITTER       ProtocolType = 4
)

var ProtocolType_name = map[int32]string{
	0: "AUTO_PROTOCOL",
	1: "BINARY",
	2: "LAX_BINARY",
	3: "COMPACT",
	4: "TWITTER",
}

var ProtocolType_value = map[string]int32{
	"AUTO_PROTOCOL": 0,
	"BINARY":        1,
	"LAX_BINARY":    2,
	"COMPACT":       3,
	"TWITTER":       4,
}

func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}

func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{1}
}

type ThriftProxy struct {
	Transport            TransportType       `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType" json:"transport,omitempty"`
	Protocol             ProtocolType        `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType" json:"protocol,omitempty"`
	StatPrefix           string              `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	RouteConfig          *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	ThriftFilters        []*ThriftFilter     `protobuf:"bytes,5,rep,name=thrift_filters,json=thriftFilters,proto3" json:"thrift_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ThriftProxy) Reset()         { *m = ThriftProxy{} }
func (m *ThriftProxy) String() string { return proto.CompactTextString(m) }
func (*ThriftProxy) ProtoMessage()    {}
func (*ThriftProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{0}
}

func (m *ThriftProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProxy.Unmarshal(m, b)
}
func (m *ThriftProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProxy.Marshal(b, m, deterministic)
}
func (m *ThriftProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProxy.Merge(m, src)
}
func (m *ThriftProxy) XXX_Size() int {
	return xxx_messageInfo_ThriftProxy.Size(m)
}
func (m *ThriftProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProxy proto.InternalMessageInfo

func (m *ThriftProxy) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProxy) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func (m *ThriftProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ThriftProxy) GetRouteConfig() *RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *ThriftProxy) GetThriftFilters() []*ThriftFilter {
	if m != nil {
		return m.ThriftFilters
	}
	return nil
}

type ThriftFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ThriftFilter_Config
	//	*ThriftFilter_TypedConfig
	ConfigType           isThriftFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ThriftFilter) Reset()         { *m = ThriftFilter{} }
func (m *ThriftFilter) String() string { return proto.CompactTextString(m) }
func (*ThriftFilter) ProtoMessage()    {}
func (*ThriftFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{1}
}

func (m *ThriftFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftFilter.Unmarshal(m, b)
}
func (m *ThriftFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftFilter.Marshal(b, m, deterministic)
}
func (m *ThriftFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftFilter.Merge(m, src)
}
func (m *ThriftFilter) XXX_Size() int {
	return xxx_messageInfo_ThriftFilter.Size(m)
}
func (m *ThriftFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftFilter proto.InternalMessageInfo

func (m *ThriftFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isThriftFilter_ConfigType interface {
	isThriftFilter_ConfigType()
}

type ThriftFilter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type ThriftFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ThriftFilter_Config) isThriftFilter_ConfigType() {}

func (*ThriftFilter_TypedConfig) isThriftFilter_ConfigType() {}

func (m *ThriftFilter) GetConfigType() isThriftFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *ThriftFilter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ThriftFilter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ThriftFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ThriftFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ThriftFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ThriftFilter_Config)(nil),
		(*ThriftFilter_TypedConfig)(nil),
	}
}

type ThriftProtocolOptions struct {
	Transport            TransportType `protobuf:"varint,1,opt,name=transport,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType" json:"transport,omitempty"`
	Protocol             ProtocolType  `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ThriftProtocolOptions) Reset()         { *m = ThriftProtocolOptions{} }
func (m *ThriftProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*ThriftProtocolOptions) ProtoMessage()    {}
func (*ThriftProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{2}
}

func (m *ThriftProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProtocolOptions.Unmarshal(m, b)
}
func (m *ThriftProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProtocolOptions.Marshal(b, m, deterministic)
}
func (m *ThriftProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProtocolOptions.Merge(m, src)
}
func (m *ThriftProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_ThriftProtocolOptions.Size(m)
}
func (m *ThriftProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProtocolOptions proto.InternalMessageInfo

func (m *ThriftProtocolOptions) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProtocolOptions) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func init() {
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.TransportType", TransportType_name, TransportType_value)
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterType((*ThriftProxy)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProxy")
	proto.RegisterType((*ThriftFilter)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftFilter")
	proto.RegisterType((*ThriftProtocolOptions)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto", fileDescriptor_e8fab7646d88fc90)
}

var fileDescriptor_e8fab7646d88fc90 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x3d, 0x4f, 0xdb, 0x40,
	0x18, 0xe6, 0x9c, 0xf0, 0xf5, 0x3a, 0x89, 0xdc, 0x53, 0x2b, 0x5c, 0x5a, 0x55, 0x11, 0x53, 0xc4,
	0x60, 0x97, 0xa0, 0x0e, 0x1d, 0x2a, 0x6a, 0x87, 0x20, 0x90, 0x00, 0x5b, 0x87, 0x11, 0xed, 0x14,
	0x1d, 0x70, 0x4e, 0xac, 0x06, 0x9f, 0x75, 0xbe, 0x50, 0xb2, 0x76, 0xaf, 0xba, 0xf6, 0x47, 0xf4,
	0x17, 0xf5, 0x17, 0x74, 0xee, 0xc8, 0x50, 0x55, 0xbe, 0x73, 0x20, 0xa1, 0xea, 0x00, 0x52, 0xbb,
	0xf9, 0xfd, 0x7a, 0x9e, 0xf7, 0x9e, 0xf7, 0x91, 0x61, 0x9b, 0xa5, 0x97, 0x7c, 0xec, 0x9e, 0xf1,
	0x34, 0x4e, 0xfa, 0x6e, 0x9c, 0x0c, 0x25, 0x13, 0x6e, 0xca, 0xe4, 0x47, 0x2e, 0x3e, 0xb8, 0x72,
	0x20, 0x92, 0x58, 0xf6, 0x32, 0xc1, 0xaf, 0xc6, 0xee, 0x65, 0x9b, 0x0e, 0xb3, 0x01, 0xdd, 0x98,
	0xc9, 0x3a, 0x99, 0xe0, 0x92, 0xe3, 0x0d, 0x85, 0xe2, 0x68, 0x14, 0x47, 0xa3, 0x38, 0x25, 0x8a,
	0x33, 0xd3, 0x3f, 0x41, 0x59, 0x7d, 0x73, 0x7f, 0x62, 0xc1, 0x47, 0x92, 0x69, 0xc6, 0xd5, 0xa7,
	0x7d, 0xce, 0xfb, 0x43, 0xe6, 0xaa, 0xe8, 0x74, 0x14, 0xbb, 0x34, 0x2d, 0x97, 0x59, 0x7d, 0x7e,
	0xb7, 0x94, 0x4b, 0x31, 0x3a, 0x93, 0x65, 0xf5, 0xc5, 0xe8, 0x3c, 0xa3, 0x2e, 0x4d, 0x53, 0x2e,
	0xa9, 0x4c, 0x78, 0x9a, 0xbb, 0x17, 0x49, 0x5f, 0xd0, 0x1b, 0xe0, 0x95, 0x4b, 0x3a, 0x4c, 0xce,
	0xa9, 0x64, 0xee, 0xe4, 0x43, 0x17, 0xd6, 0xbe, 0x57, 0xc0, 0x8c, 0xd4, 0x5e, 0x61, 0xb1, 0x16,
	0x1e, 0xc0, 0xb2, 0x14, 0x34, 0xcd, 0x33, 0x2e, 0xa4, 0x6d, 0x34, 0x51, 0xab, 0xd1, 0x7e, 0xeb,
	0xdc, 0x5b, 0x07, 0x27, 0x9a, 0x60, 0x44, 0xe3, 0x8c, 0xf9, 0x4b, 0xd7, 0xfe, 0xfc, 0x27, 0x64,
	0x58, 0x88, 0xdc, 0x82, 0x63, 0x06, 0x4b, 0x6a, 0x85, 0x33, 0x3e, 0xb4, 0x2b, 0x8a, 0x68, 0xeb,
	0x01, 0x44, 0x61, 0x09, 0x71, 0x87, 0xe7, 0x06, 0x1a, 0xb7, 0xc0, 0xcc, 0x25, 0x2d, 0xe6, 0x58,
	0x9c, 0x5c, 0xd9, 0xa8, 0x89, 0x5a, 0xcb, 0xfe, 0xe2, 0xb5, 0x5f, 0x15, 0x46, 0x13, 0x11, 0x28,
	0x6a, 0xa1, 0x2a, 0xe1, 0x01, 0xd4, 0xd4, 0x2d, 0x7a, 0x9a, 0xdf, 0xae, 0x36, 0x51, 0xcb, 0x6c,
	0x77, 0x1f, 0xb0, 0x14, 0x29, 0x60, 0x3a, 0x6a, 0x60, 0x24, 0xd4, 0x3d, 0x88, 0x29, 0x6e, 0x73,
	0x38, 0x86, 0x46, 0x39, 0xa8, 0xe1, 0x72, 0x7b, 0xbe, 0x59, 0x69, 0x99, 0x0f, 0x12, 0x40, 0x1f,
	0x6f, 0x47, 0xb5, 0x92, 0xba, 0x9c, 0x8a, 0xf2, 0xb5, 0x6f, 0x08, 0x6a, 0xd3, 0x75, 0xfc, 0x0c,
	0xaa, 0x29, 0xbd, 0x60, 0x77, 0x55, 0x50, 0x49, 0xfc, 0x0a, 0x16, 0xca, 0x97, 0x1b, 0xea, 0xe5,
	0x2b, 0x8e, 0xb6, 0x9c, 0x33, 0xb1, 0x9c, 0x73, 0xa4, 0x2c, 0xe7, 0x1b, 0x36, 0xda, 0x9d, 0x23,
	0x65, 0x33, 0x7e, 0x0d, 0x35, 0x39, 0xce, 0xd8, 0xf9, 0x44, 0xb6, 0x8a, 0x1a, 0x7e, 0xfc, 0xc7,
	0xb0, 0x97, 0x8e, 0x77, 0xe7, 0x88, 0xa9, 0x7a, 0xb5, 0x0e, 0x7e, 0x1d, 0x4c, 0x3d, 0xd4, 0x2b,
	0xb2, 0x6b, 0x3f, 0x10, 0x3c, 0xb9, 0xf1, 0xa2, 0xba, 0x5e, 0x90, 0x29, 0x33, 0xcf, 0xba, 0x12,
	0xfd, 0x2f, 0x57, 0x1a, 0xff, 0xcc, 0x95, 0xeb, 0x7b, 0x50, 0x9f, 0x59, 0x06, 0x63, 0x68, 0x78,
	0xc7, 0x51, 0xd0, 0x8b, 0x88, 0x77, 0x78, 0x14, 0x06, 0x24, 0xb2, 0xe6, 0x30, 0xc0, 0xc2, 0x0e,
	0xf1, 0x0e, 0xba, 0xdb, 0x16, 0xc2, 0x35, 0x58, 0x3a, 0x3e, 0x2c, 0x23, 0xa3, 0xa8, 0xec, 0x76,
	0xbd, 0xed, 0x2e, 0xb1, 0x2a, 0xeb, 0x27, 0x50, 0x9b, 0xa6, 0xc3, 0x8f, 0xa0, 0xae, 0x90, 0x42,
	0x12, 0x44, 0x41, 0x27, 0xd8, 0xd7, 0x40, 0xfe, 0xde, 0xa1, 0x47, 0xde, 0x5b, 0x08, 0x37, 0x00,
	0xf6, 0xbd, 0x77, 0xbd, 0x32, 0x36, 0xb0, 0x09, 0x8b, 0x9d, 0xe0, 0x20, 0xf4, 0x3a, 0x91, 0x55,
	0x29, 0x82, 0xe8, 0x64, 0x2f, 0x8a, 0xba, 0xc4, 0xaa, 0xfa, 0x9f, 0xd1, 0xcf, 0xaf, 0xbf, 0xbe,
	0xcc, 0xb7, 0xf1, 0x4b, 0x2d, 0x00, 0xbb, 0x92, 0x2c, 0xcd, 0x8b, 0x7b, 0x94, 0x22, 0xe4, 0x7f,
	0x51, 0x61, 0x13, 0xb6, 0x12, 0xae, 0x55, 0xd3, 0x99, 0x7b, 0x0b, 0xe8, 0x5b, 0x53, 0xff, 0x24,
	0xf5, 0xba, 0x10, 0x9d, 0x2e, 0x28, 0xf5, 0x36, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x34, 0xec,
	0x99, 0xe0, 0xdd, 0x05, 0x00, 0x00,
}
