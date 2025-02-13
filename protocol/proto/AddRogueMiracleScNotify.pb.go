// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: AddRogueMiracleScNotify.proto

package proto

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

type AddRogueMiracleScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source       RogueMiracleSource `protobuf:"varint,15,opt,name=source,proto3,enum=proto.RogueMiracleSource" json:"source,omitempty"`
	RogueMiracle *RogueMiracle      `protobuf:"bytes,4,opt,name=rogue_miracle,json=rogueMiracle,proto3" json:"rogue_miracle,omitempty"`
}

func (x *AddRogueMiracleScNotify) Reset() {
	*x = AddRogueMiracleScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddRogueMiracleScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRogueMiracleScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRogueMiracleScNotify) ProtoMessage() {}

func (x *AddRogueMiracleScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AddRogueMiracleScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRogueMiracleScNotify.ProtoReflect.Descriptor instead.
func (*AddRogueMiracleScNotify) Descriptor() ([]byte, []int) {
	return file_AddRogueMiracleScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AddRogueMiracleScNotify) GetSource() RogueMiracleSource {
	if x != nil {
		return x.Source
	}
	return RogueMiracleSource_ROGUE_MIRACLE_SOURCE_TYPE_NONE
}

func (x *AddRogueMiracleScNotify) GetRogueMiracle() *RogueMiracle {
	if x != nil {
		return x.RogueMiracle
	}
	return nil
}

var File_AddRogueMiracleScNotify_proto protoreflect.FileDescriptor

var file_AddRogueMiracleScNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x31, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52,
	0x0c, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AddRogueMiracleScNotify_proto_rawDescOnce sync.Once
	file_AddRogueMiracleScNotify_proto_rawDescData = file_AddRogueMiracleScNotify_proto_rawDesc
)

func file_AddRogueMiracleScNotify_proto_rawDescGZIP() []byte {
	file_AddRogueMiracleScNotify_proto_rawDescOnce.Do(func() {
		file_AddRogueMiracleScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AddRogueMiracleScNotify_proto_rawDescData)
	})
	return file_AddRogueMiracleScNotify_proto_rawDescData
}

var file_AddRogueMiracleScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AddRogueMiracleScNotify_proto_goTypes = []interface{}{
	(*AddRogueMiracleScNotify)(nil), // 0: proto.AddRogueMiracleScNotify
	(RogueMiracleSource)(0),         // 1: proto.RogueMiracleSource
	(*RogueMiracle)(nil),            // 2: proto.RogueMiracle
}
var file_AddRogueMiracleScNotify_proto_depIdxs = []int32{
	1, // 0: proto.AddRogueMiracleScNotify.source:type_name -> proto.RogueMiracleSource
	2, // 1: proto.AddRogueMiracleScNotify.rogue_miracle:type_name -> proto.RogueMiracle
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AddRogueMiracleScNotify_proto_init() }
func file_AddRogueMiracleScNotify_proto_init() {
	if File_AddRogueMiracleScNotify_proto != nil {
		return
	}
	file_RogueMiracleSource_proto_init()
	file_RogueMiracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AddRogueMiracleScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRogueMiracleScNotify); i {
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
			RawDescriptor: file_AddRogueMiracleScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AddRogueMiracleScNotify_proto_goTypes,
		DependencyIndexes: file_AddRogueMiracleScNotify_proto_depIdxs,
		MessageInfos:      file_AddRogueMiracleScNotify_proto_msgTypes,
	}.Build()
	File_AddRogueMiracleScNotify_proto = out.File
	file_AddRogueMiracleScNotify_proto_rawDesc = nil
	file_AddRogueMiracleScNotify_proto_goTypes = nil
	file_AddRogueMiracleScNotify_proto_depIdxs = nil
}
