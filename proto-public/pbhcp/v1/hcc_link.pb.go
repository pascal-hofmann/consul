// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbhcp/v1/hcc_link.proto

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

type HCCLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId   string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ClientId     string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *HCCLink) Reset() {
	*x = HCCLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbhcp_v1_hcc_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HCCLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HCCLink) ProtoMessage() {}

func (x *HCCLink) ProtoReflect() protoreflect.Message {
	mi := &file_pbhcp_v1_hcc_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HCCLink.ProtoReflect.Descriptor instead.
func (*HCCLink) Descriptor() ([]byte, []int) {
	return file_pbhcp_v1_hcc_link_proto_rawDescGZIP(), []int{0}
}

func (x *HCCLink) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *HCCLink) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *HCCLink) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

var File_pbhcp_v1_hcc_link_proto protoreflect.FileDescriptor

var file_pbhcp_v1_hcc_link_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x68, 0x63, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x63, 0x63, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x74, 0x0a, 0x07, 0x48, 0x43, 0x43, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x3a, 0x06,
	0xa2, 0x93, 0x04, 0x02, 0x08, 0x01, 0x42, 0xe3, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x68, 0x63, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x48, 0x63, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2f, 0x70, 0x62, 0x68, 0x63, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x63, 0x70, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x48, 0x43, 0x48, 0xaa, 0x02, 0x17, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x48, 0x63, 0x70, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x17, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x5c, 0x48, 0x63, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x48, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x48, 0x63, 0x70,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1a, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x48, 0x63, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbhcp_v1_hcc_link_proto_rawDescOnce sync.Once
	file_pbhcp_v1_hcc_link_proto_rawDescData = file_pbhcp_v1_hcc_link_proto_rawDesc
)

func file_pbhcp_v1_hcc_link_proto_rawDescGZIP() []byte {
	file_pbhcp_v1_hcc_link_proto_rawDescOnce.Do(func() {
		file_pbhcp_v1_hcc_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbhcp_v1_hcc_link_proto_rawDescData)
	})
	return file_pbhcp_v1_hcc_link_proto_rawDescData
}

var file_pbhcp_v1_hcc_link_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pbhcp_v1_hcc_link_proto_goTypes = []interface{}{
	(*HCCLink)(nil), // 0: hashicorp.consul.hcp.v1.HCCLink
}
var file_pbhcp_v1_hcc_link_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbhcp_v1_hcc_link_proto_init() }
func file_pbhcp_v1_hcc_link_proto_init() {
	if File_pbhcp_v1_hcc_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbhcp_v1_hcc_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HCCLink); i {
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
			RawDescriptor: file_pbhcp_v1_hcc_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbhcp_v1_hcc_link_proto_goTypes,
		DependencyIndexes: file_pbhcp_v1_hcc_link_proto_depIdxs,
		MessageInfos:      file_pbhcp_v1_hcc_link_proto_msgTypes,
	}.Build()
	File_pbhcp_v1_hcc_link_proto = out.File
	file_pbhcp_v1_hcc_link_proto_rawDesc = nil
	file_pbhcp_v1_hcc_link_proto_goTypes = nil
	file_pbhcp_v1_hcc_link_proto_depIdxs = nil
}
