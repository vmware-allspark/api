// Copyright 2018 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: envoy/config/filter/http/alpn/v2alpha1/config.proto

// $title: ALPN filter for overriding ALPN for upstream TLS connections.

package v2alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Upstream protocols
type FilterConfig_Protocol int32

const (
	FilterConfig_HTTP10 FilterConfig_Protocol = 0
	FilterConfig_HTTP11 FilterConfig_Protocol = 1
	FilterConfig_HTTP2  FilterConfig_Protocol = 2
)

// Enum value maps for FilterConfig_Protocol.
var (
	FilterConfig_Protocol_name = map[int32]string{
		0: "HTTP10",
		1: "HTTP11",
		2: "HTTP2",
	}
	FilterConfig_Protocol_value = map[string]int32{
		"HTTP10": 0,
		"HTTP11": 1,
		"HTTP2":  2,
	}
)

func (x FilterConfig_Protocol) Enum() *FilterConfig_Protocol {
	p := new(FilterConfig_Protocol)
	*p = x
	return p
}

func (x FilterConfig_Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterConfig_Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_enumTypes[0].Descriptor()
}

func (FilterConfig_Protocol) Type() protoreflect.EnumType {
	return &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_enumTypes[0]
}

func (x FilterConfig_Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterConfig_Protocol.Descriptor instead.
func (FilterConfig_Protocol) EnumDescriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescGZIP(), []int{0, 0}
}

// FilterConfig is the config for Istio-specific filter.
type FilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map from upstream protocol to list of ALPN
	AlpnOverride []*FilterConfig_AlpnOverride `protobuf:"bytes,1,rep,name=alpn_override,json=alpnOverride,proto3" json:"alpn_override,omitempty"`
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig.ProtoReflect.Descriptor instead.
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfig) GetAlpnOverride() []*FilterConfig_AlpnOverride {
	if x != nil {
		return x.AlpnOverride
	}
	return nil
}

type FilterConfig_AlpnOverride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Upstream protocol
	UpstreamProtocol FilterConfig_Protocol `protobuf:"varint,1,opt,name=upstream_protocol,json=upstreamProtocol,proto3,enum=istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig_Protocol" json:"upstream_protocol,omitempty"`
	// A list of ALPN that will override the ALPN for upstream TLS connections.
	AlpnOverride []string `protobuf:"bytes,2,rep,name=alpn_override,json=alpnOverride,proto3" json:"alpn_override,omitempty"`
}

func (x *FilterConfig_AlpnOverride) Reset() {
	*x = FilterConfig_AlpnOverride{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig_AlpnOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig_AlpnOverride) ProtoMessage() {}

func (x *FilterConfig_AlpnOverride) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig_AlpnOverride.ProtoReflect.Descriptor instead.
func (*FilterConfig_AlpnOverride) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FilterConfig_AlpnOverride) GetUpstreamProtocol() FilterConfig_Protocol {
	if x != nil {
		return x.UpstreamProtocol
	}
	return FilterConfig_HTTP10
}

func (x *FilterConfig_AlpnOverride) GetAlpnOverride() []string {
	if x != nil {
		return x.AlpnOverride
	}
	return nil
}

var File_envoy_config_filter_http_alpn_v2alpha1_config_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x6c, 0x70, 0x6e, 0x2f,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x6c, 0x70, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x22, 0xd3, 0x02, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x6c, 0x0a, 0x0d, 0x61, 0x6c, 0x70, 0x6e, 0x5f, 0x6f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x6c, 0x70,
	0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x6c, 0x70, 0x6e, 0x4f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x52, 0x0c, 0x61, 0x6c, 0x70, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x1a, 0xa5, 0x01, 0x0a, 0x0c, 0x41, 0x6c, 0x70, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x12, 0x70, 0x0a, 0x11, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x43,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x61, 0x6c, 0x70, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x52, 0x10, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x70, 0x6e, 0x5f, 0x6f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c,
	0x70, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x22, 0x2d, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x54, 0x54, 0x50, 0x31, 0x30,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x54, 0x54, 0x50, 0x31, 0x31, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x32, 0x10, 0x02, 0x42, 0x35, 0x5a, 0x33, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x61, 0x6c, 0x70, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescData = file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDesc
)

func file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescData)
	})
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescData
}

var file_envoy_config_filter_http_alpn_v2alpha1_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_filter_http_alpn_v2alpha1_config_proto_goTypes = []interface{}{
	(FilterConfig_Protocol)(0),        // 0: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.Protocol
	(*FilterConfig)(nil),              // 1: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig
	(*FilterConfig_AlpnOverride)(nil), // 2: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride
}
var file_envoy_config_filter_http_alpn_v2alpha1_config_proto_depIdxs = []int32{
	2, // 0: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.alpn_override:type_name -> istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride
	0, // 1: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride.upstream_protocol:type_name -> istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.Protocol
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_alpn_v2alpha1_config_proto_init() }
func file_envoy_config_filter_http_alpn_v2alpha1_config_proto_init() {
	if File_envoy_config_filter_http_alpn_v2alpha1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig_AlpnOverride); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_alpn_v2alpha1_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_alpn_v2alpha1_config_proto_depIdxs,
		EnumInfos:         file_envoy_config_filter_http_alpn_v2alpha1_config_proto_enumTypes,
		MessageInfos:      file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_alpn_v2alpha1_config_proto = out.File
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDesc = nil
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_goTypes = nil
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_depIdxs = nil
}
