// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: provider.proto

package provider

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

// The request message containing the provider from which the cipher needs to fetched.
type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{0}
}

func (x *Provider) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

// The response message containing apikey
type APIKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apikey string `protobuf:"bytes,1,opt,name=apikey,proto3" json:"apikey,omitempty"`
}

func (x *APIKey) Reset() {
	*x = APIKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIKey) ProtoMessage() {}

func (x *APIKey) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIKey.ProtoReflect.Descriptor instead.
func (*APIKey) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{1}
}

func (x *APIKey) GetApikey() string {
	if x != nil {
		return x.Apikey
	}
	return ""
}

var File_provider_proto protoreflect.FileDescriptor

var file_provider_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70,
	0x69, 0x6b, 0x65, 0x79, 0x32, 0x45, 0x0a, 0x0e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49,
	0x4b, 0x65, 0x79, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x42, 0x74, 0x0a, 0x1a, 0x69,
	0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x62, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x12, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x01, 0x5a,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x62, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6c, 0x63, 0x68, 0x65, 0x6d, 0x79, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x2f, 0x61, 0x72, 0x6d, 0x61, 0x64, 0x61, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provider_proto_rawDescOnce sync.Once
	file_provider_proto_rawDescData = file_provider_proto_rawDesc
)

func file_provider_proto_rawDescGZIP() []byte {
	file_provider_proto_rawDescOnce.Do(func() {
		file_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_provider_proto_rawDescData)
	})
	return file_provider_proto_rawDescData
}

var file_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_provider_proto_goTypes = []interface{}{
	(*Provider)(nil), // 0: provider.Provider
	(*APIKey)(nil),   // 1: provider.APIKey
}
var file_provider_proto_depIdxs = []int32{
	0, // 0: provider.APIKeyProvider.GetAPIKey:input_type -> provider.Provider
	1, // 1: provider.APIKeyProvider.GetAPIKey:output_type -> provider.APIKey
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_provider_proto_init() }
func file_provider_proto_init() {
	if File_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
		file_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIKey); i {
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
			RawDescriptor: file_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provider_proto_goTypes,
		DependencyIndexes: file_provider_proto_depIdxs,
		MessageInfos:      file_provider_proto_msgTypes,
	}.Build()
	File_provider_proto = out.File
	file_provider_proto_rawDesc = nil
	file_provider_proto_goTypes = nil
	file_provider_proto_depIdxs = nil
}
