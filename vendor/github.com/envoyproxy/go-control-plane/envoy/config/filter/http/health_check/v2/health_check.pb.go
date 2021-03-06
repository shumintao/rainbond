// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/health_check/v2/health_check.proto

package envoy_config_filter_http_health_check_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type HealthCheck struct {
	PassThroughMode              *wrappers.BoolValue       `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode,proto3" json:"pass_through_mode,omitempty"`
	CacheTime                    *duration.Duration        `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime,proto3" json:"cache_time,omitempty"`
	ClusterMinHealthyPercentages map[string]*_type.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages,proto3" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Headers                      []*route.HeaderMatcher    `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                  `json:"-"`
	XXX_unrecognized             []byte                    `json:"-"`
	XXX_sizecache                int32                     `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_75439d7b4d98e201, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPassThroughMode() *wrappers.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetCacheTime() *duration.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func (m *HealthCheck) GetClusterMinHealthyPercentages() map[string]*_type.Percent {
	if m != nil {
		return m.ClusterMinHealthyPercentages
	}
	return nil
}

func (m *HealthCheck) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck")
	proto.RegisterMapType((map[string]*_type.Percent)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck.ClusterMinHealthyPercentagesEntry")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/health_check/v2/health_check.proto", fileDescriptor_75439d7b4d98e201)
}

var fileDescriptor_75439d7b4d98e201 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3f, 0x6f, 0xd4, 0x30,
	0x14, 0x57, 0xee, 0x7a, 0xd0, 0xfa, 0x06, 0x8e, 0x30, 0x10, 0x4e, 0xa8, 0xb4, 0x4c, 0xd7, 0x01,
	0x5b, 0xba, 0x4a, 0xa8, 0xa2, 0xdb, 0x15, 0x24, 0x84, 0x74, 0xd2, 0x29, 0xaa, 0x60, 0x8c, 0x5c,
	0xe7, 0x5d, 0x62, 0x35, 0xb1, 0x2d, 0xe7, 0xe5, 0x68, 0x46, 0x56, 0x16, 0x56, 0x26, 0xbe, 0x1c,
	0xdf, 0x80, 0x91, 0x01, 0x21, 0xdb, 0x39, 0xd1, 0x8a, 0x7f, 0x5d, 0x22, 0x3b, 0xef, 0xf7, 0xe7,
	0xbd, 0xdf, 0x33, 0x39, 0x05, 0xb5, 0xd1, 0x1d, 0x13, 0x5a, 0xad, 0x65, 0xc1, 0xd6, 0xb2, 0x42,
	0xb0, 0xac, 0x44, 0x34, 0xac, 0x04, 0x5e, 0x61, 0x99, 0x89, 0x12, 0xc4, 0x25, 0xdb, 0xcc, 0x6f,
	0xdc, 0xa9, 0xb1, 0x1a, 0x75, 0x3c, 0xf3, 0x64, 0x1a, 0xc8, 0x34, 0x90, 0xa9, 0x23, 0xd3, 0x1b,
	0xe0, 0xcd, 0x7c, 0x7a, 0x14, 0x6c, 0xb8, 0x91, 0x4e, 0xca, 0xea, 0x16, 0x21, 0x7c, 0x33, 0xa1,
	0x6b, 0xa3, 0x15, 0x28, 0x6c, 0x82, 0xe8, 0x34, 0x09, 0x50, 0xec, 0x0c, 0x30, 0x03, 0x56, 0x80,
	0xc2, 0xbe, 0xb2, 0x5f, 0x68, 0x5d, 0x54, 0xc0, 0xfc, 0xed, 0xa2, 0x5d, 0xb3, 0xbc, 0xb5, 0x1c,
	0xa5, 0x56, 0x7f, 0xab, 0xbf, 0xb7, 0xdc, 0x18, 0xb0, 0x5b, 0xe5, 0xfd, 0x36, 0x37, 0x9c, 0x71,
	0xa5, 0x34, 0x7a, 0x5a, 0xc3, 0x6a, 0x59, 0x58, 0x8e, 0xd0, 0xd7, 0x1f, 0x6e, 0x78, 0x25, 0x73,
	0x8e, 0xc0, 0xb6, 0x87, 0x50, 0x78, 0xfa, 0x75, 0x48, 0xc6, 0xaf, 0xfd, 0x44, 0x67, 0x6e, 0xa0,
	0x78, 0x45, 0xee, 0x1b, 0xde, 0x34, 0x19, 0x96, 0x56, 0xb7, 0x45, 0x99, 0xd5, 0x3a, 0x87, 0x24,
	0x3a, 0x88, 0x66, 0xe3, 0xf9, 0x94, 0x86, 0x26, 0xe8, 0xb6, 0x09, 0xba, 0xd0, 0xba, 0x7a, 0xcb,
	0xab, 0x16, 0x16, 0xbb, 0xdf, 0x17, 0xa3, 0x8f, 0xd1, 0x60, 0x12, 0xa5, 0xf7, 0x1c, 0xfd, 0x3c,
	0xb0, 0x97, 0x3a, 0x87, 0xf8, 0x84, 0x10, 0xc1, 0x45, 0x09, 0x19, 0xca, 0x1a, 0x92, 0xa1, 0x97,
	0x7a, 0xf4, 0x9b, 0xd4, 0xcb, 0x7e, 0xde, 0x74, 0xcf, 0x83, 0xcf, 0x65, 0x0d, 0xf1, 0x97, 0x88,
	0x3c, 0x11, 0x55, 0xdb, 0x20, 0xd8, 0xac, 0x96, 0x2a, 0x0b, 0xc9, 0x77, 0x59, 0x1f, 0x1d, 0x2f,
	0xa0, 0x49, 0x76, 0x0e, 0x86, 0xb3, 0xf1, 0xfc, 0x1d, 0xbd, 0xed, 0xba, 0xe8, 0xb5, 0x61, 0xe9,
	0x59, 0x10, 0x5f, 0x4a, 0x15, 0xfe, 0x76, 0xab, 0x5f, 0xca, 0xaf, 0x14, 0xda, 0x2e, 0x7d, 0x2c,
	0xfe, 0x01, 0x89, 0x4f, 0xc9, 0xdd, 0x12, 0x78, 0x0e, 0xb6, 0x49, 0x46, 0xbe, 0x8f, 0xc3, 0xbe,
	0x0f, 0x6e, 0xa4, 0xf3, 0xf2, 0xcf, 0xc0, 0x39, 0xe6, 0x60, 0x97, 0x1c, 0x45, 0x09, 0x36, 0xdd,
	0x32, 0xa6, 0x39, 0x39, 0xfc, 0xaf, 0x7f, 0x3c, 0x21, 0xc3, 0x4b, 0xe8, 0xfc, 0x02, 0xf6, 0x52,
	0x77, 0x8c, 0x8f, 0xc8, 0x68, 0xe3, 0x22, 0x4f, 0x06, 0x3e, 0xc9, 0x07, 0xbd, 0xa3, 0x7b, 0x53,
	0xb4, 0xa7, 0xa7, 0x01, 0xf1, 0x62, 0x70, 0x12, 0xbd, 0xd9, 0xd9, 0x1d, 0x4c, 0x86, 0x8b, 0x0f,
	0xd1, 0xb7, 0xcf, 0x3f, 0x3e, 0x8d, 0x58, 0xfc, 0x2c, 0xa0, 0xe1, 0x0a, 0x41, 0x35, 0xee, 0x9d,
	0xf4, 0x59, 0x35, 0x7f, 0x0a, 0xeb, 0x98, 0x3c, 0x97, 0x3a, 0xe8, 0x1b, 0xab, 0xaf, 0xba, 0x5b,
	0x87, 0xbc, 0x98, 0x5c, 0x4b, 0x79, 0xe5, 0x56, 0xbc, 0x8a, 0x2e, 0xee, 0xf8, 0x5d, 0x1f, 0xff,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x79, 0x93, 0xf1, 0x54, 0x97, 0x03, 0x00, 0x00,
}
