// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pbhcp/v1/telemetry_state.proto

package hcpv1

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
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

type TelemetryState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId   string            `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ClientId     string            `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string            `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	Endpoint     string            `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Labels       map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metrics      *MetricsState     `protobuf:"bytes,6,opt,name=metrics,proto3" json:"metrics,omitempty"`
	// Auth Endpoint is the endpoint against which client id and secret are exchanged for tokens
	// to call the API. In all cases, except development, this is https://auth.idp.hashicorp.com.
	AuthEndpoint string `protobuf:"bytes,7,opt,name=auth_endpoint,json=authEndpoint,proto3" json:"auth_endpoint,omitempty"`
}

func (x *TelemetryState) Reset() {
	*x = TelemetryState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbhcp_v1_telemetry_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryState) ProtoMessage() {}

func (x *TelemetryState) ProtoReflect() protoreflect.Message {
	mi := &file_pbhcp_v1_telemetry_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryState.ProtoReflect.Descriptor instead.
func (*TelemetryState) Descriptor() ([]byte, []int) {
	return file_pbhcp_v1_telemetry_state_proto_rawDescGZIP(), []int{0}
}

func (x *TelemetryState) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *TelemetryState) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *TelemetryState) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *TelemetryState) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *TelemetryState) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TelemetryState) GetMetrics() *MetricsState {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *TelemetryState) GetAuthEndpoint() string {
	if x != nil {
		return x.AuthEndpoint
	}
	return ""
}

type MetricsState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint    string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	IncludeList []string `protobuf:"bytes,2,rep,name=include_list,json=includeList,proto3" json:"include_list,omitempty"`
	Disabled    bool     `protobuf:"varint,3,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *MetricsState) Reset() {
	*x = MetricsState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbhcp_v1_telemetry_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsState) ProtoMessage() {}

func (x *MetricsState) ProtoReflect() protoreflect.Message {
	mi := &file_pbhcp_v1_telemetry_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsState.ProtoReflect.Descriptor instead.
func (*MetricsState) Descriptor() ([]byte, []int) {
	return file_pbhcp_v1_telemetry_state_proto_rawDescGZIP(), []int{1}
}

func (x *MetricsState) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *MetricsState) GetIncludeList() []string {
	if x != nil {
		return x.IncludeList
	}
	return nil
}

func (x *MetricsState) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

var File_pbhcp_v1_telemetry_state_proto protoreflect.FileDescriptor

var file_pbhcp_v1_telemetry_state_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x62, 0x68, 0x63, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x03, 0x0a, 0x0e, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x75, 0x74, 0x68, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x01, 0x22,
	0x69, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0xea, 0x01, 0x0a, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x68, 0x63,
	0x70, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x63, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x48,
	0xaa, 0x02, 0x17, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x48, 0x63, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x48, 0x63,
	0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x48, 0x63, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a,
	0x48, 0x63, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbhcp_v1_telemetry_state_proto_rawDescOnce sync.Once
	file_pbhcp_v1_telemetry_state_proto_rawDescData = file_pbhcp_v1_telemetry_state_proto_rawDesc
)

func file_pbhcp_v1_telemetry_state_proto_rawDescGZIP() []byte {
	file_pbhcp_v1_telemetry_state_proto_rawDescOnce.Do(func() {
		file_pbhcp_v1_telemetry_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbhcp_v1_telemetry_state_proto_rawDescData)
	})
	return file_pbhcp_v1_telemetry_state_proto_rawDescData
}

var file_pbhcp_v1_telemetry_state_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pbhcp_v1_telemetry_state_proto_goTypes = []interface{}{
	(*TelemetryState)(nil), // 0: hashicorp.consul.hcp.v1.TelemetryState
	(*MetricsState)(nil),   // 1: hashicorp.consul.hcp.v1.MetricsState
	nil,                    // 2: hashicorp.consul.hcp.v1.TelemetryState.LabelsEntry
}
var file_pbhcp_v1_telemetry_state_proto_depIdxs = []int32{
	2, // 0: hashicorp.consul.hcp.v1.TelemetryState.labels:type_name -> hashicorp.consul.hcp.v1.TelemetryState.LabelsEntry
	1, // 1: hashicorp.consul.hcp.v1.TelemetryState.metrics:type_name -> hashicorp.consul.hcp.v1.MetricsState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pbhcp_v1_telemetry_state_proto_init() }
func file_pbhcp_v1_telemetry_state_proto_init() {
	if File_pbhcp_v1_telemetry_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbhcp_v1_telemetry_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryState); i {
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
		file_pbhcp_v1_telemetry_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsState); i {
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
			RawDescriptor: file_pbhcp_v1_telemetry_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbhcp_v1_telemetry_state_proto_goTypes,
		DependencyIndexes: file_pbhcp_v1_telemetry_state_proto_depIdxs,
		MessageInfos:      file_pbhcp_v1_telemetry_state_proto_msgTypes,
	}.Build()
	File_pbhcp_v1_telemetry_state_proto = out.File
	file_pbhcp_v1_telemetry_state_proto_rawDesc = nil
	file_pbhcp_v1_telemetry_state_proto_goTypes = nil
	file_pbhcp_v1_telemetry_state_proto_depIdxs = nil
}
