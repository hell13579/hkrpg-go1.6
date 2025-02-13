// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: RogueRecordInfo.proto

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

type RogueRecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuffList         []*RogueBuff         `protobuf:"bytes,6,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	RogueMiracleList []uint32             `protobuf:"varint,8,rep,packed,name=rogue_miracle_list,json=rogueMiracleList,proto3" json:"rogue_miracle_list,omitempty"`
	AvatarList       []*RogueRecordAvatar `protobuf:"bytes,7,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
}

func (x *RogueRecordInfo) Reset() {
	*x = RogueRecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueRecordInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueRecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueRecordInfo) ProtoMessage() {}

func (x *RogueRecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueRecordInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueRecordInfo.ProtoReflect.Descriptor instead.
func (*RogueRecordInfo) Descriptor() ([]byte, []int) {
	return file_RogueRecordInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueRecordInfo) GetBuffList() []*RogueBuff {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *RogueRecordInfo) GetRogueMiracleList() []uint32 {
	if x != nil {
		return x.RogueMiracleList
	}
	return nil
}

func (x *RogueRecordInfo) GetAvatarList() []*RogueRecordAvatar {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

var File_RogueRecordInfo_proto protoreflect.FileDescriptor

var file_RogueRecordInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x09,
	0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66,
	0x66, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueRecordInfo_proto_rawDescOnce sync.Once
	file_RogueRecordInfo_proto_rawDescData = file_RogueRecordInfo_proto_rawDesc
)

func file_RogueRecordInfo_proto_rawDescGZIP() []byte {
	file_RogueRecordInfo_proto_rawDescOnce.Do(func() {
		file_RogueRecordInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueRecordInfo_proto_rawDescData)
	})
	return file_RogueRecordInfo_proto_rawDescData
}

var file_RogueRecordInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueRecordInfo_proto_goTypes = []interface{}{
	(*RogueRecordInfo)(nil),   // 0: proto.RogueRecordInfo
	(*RogueBuff)(nil),         // 1: proto.RogueBuff
	(*RogueRecordAvatar)(nil), // 2: proto.RogueRecordAvatar
}
var file_RogueRecordInfo_proto_depIdxs = []int32{
	1, // 0: proto.RogueRecordInfo.buff_list:type_name -> proto.RogueBuff
	2, // 1: proto.RogueRecordInfo.avatar_list:type_name -> proto.RogueRecordAvatar
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RogueRecordInfo_proto_init() }
func file_RogueRecordInfo_proto_init() {
	if File_RogueRecordInfo_proto != nil {
		return
	}
	file_RogueBuff_proto_init()
	file_RogueRecordAvatar_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueRecordInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueRecordInfo); i {
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
			RawDescriptor: file_RogueRecordInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueRecordInfo_proto_goTypes,
		DependencyIndexes: file_RogueRecordInfo_proto_depIdxs,
		MessageInfos:      file_RogueRecordInfo_proto_msgTypes,
	}.Build()
	File_RogueRecordInfo_proto = out.File
	file_RogueRecordInfo_proto_rawDesc = nil
	file_RogueRecordInfo_proto_goTypes = nil
	file_RogueRecordInfo_proto_depIdxs = nil
}
