// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: EntityBuffChangeInfo.proto

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

type EntityBuffChangeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddBuffInfo  *BuffInfo `protobuf:"bytes,12,opt,name=add_buff_info,json=addBuffInfo,proto3" json:"add_buff_info,omitempty"`
	RemoveBuffId uint32    `protobuf:"varint,4,opt,name=remove_buff_id,json=removeBuffId,proto3" json:"remove_buff_id,omitempty"`
	EntityId     uint32    `protobuf:"varint,9,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *EntityBuffChangeInfo) Reset() {
	*x = EntityBuffChangeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityBuffChangeInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityBuffChangeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityBuffChangeInfo) ProtoMessage() {}

func (x *EntityBuffChangeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_EntityBuffChangeInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityBuffChangeInfo.ProtoReflect.Descriptor instead.
func (*EntityBuffChangeInfo) Descriptor() ([]byte, []int) {
	return file_EntityBuffChangeInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EntityBuffChangeInfo) GetAddBuffInfo() *BuffInfo {
	if x != nil {
		return x.AddBuffInfo
	}
	return nil
}

func (x *EntityBuffChangeInfo) GetRemoveBuffId() uint32 {
	if x != nil {
		return x.RemoveBuffId
	}
	return 0
}

func (x *EntityBuffChangeInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_EntityBuffChangeInfo_proto protoreflect.FileDescriptor

var file_EntityBuffChangeInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x75, 0x66, 0x66, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x75,
	0x66, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x0d,
	0x61, 0x64, 0x64, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x66, 0x66,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityBuffChangeInfo_proto_rawDescOnce sync.Once
	file_EntityBuffChangeInfo_proto_rawDescData = file_EntityBuffChangeInfo_proto_rawDesc
)

func file_EntityBuffChangeInfo_proto_rawDescGZIP() []byte {
	file_EntityBuffChangeInfo_proto_rawDescOnce.Do(func() {
		file_EntityBuffChangeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityBuffChangeInfo_proto_rawDescData)
	})
	return file_EntityBuffChangeInfo_proto_rawDescData
}

var file_EntityBuffChangeInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityBuffChangeInfo_proto_goTypes = []interface{}{
	(*EntityBuffChangeInfo)(nil), // 0: proto.EntityBuffChangeInfo
	(*BuffInfo)(nil),             // 1: proto.BuffInfo
}
var file_EntityBuffChangeInfo_proto_depIdxs = []int32{
	1, // 0: proto.EntityBuffChangeInfo.add_buff_info:type_name -> proto.BuffInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EntityBuffChangeInfo_proto_init() }
func file_EntityBuffChangeInfo_proto_init() {
	if File_EntityBuffChangeInfo_proto != nil {
		return
	}
	file_BuffInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityBuffChangeInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityBuffChangeInfo); i {
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
			RawDescriptor: file_EntityBuffChangeInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityBuffChangeInfo_proto_goTypes,
		DependencyIndexes: file_EntityBuffChangeInfo_proto_depIdxs,
		MessageInfos:      file_EntityBuffChangeInfo_proto_msgTypes,
	}.Build()
	File_EntityBuffChangeInfo_proto = out.File
	file_EntityBuffChangeInfo_proto_rawDesc = nil
	file_EntityBuffChangeInfo_proto_goTypes = nil
	file_EntityBuffChangeInfo_proto_depIdxs = nil
}
