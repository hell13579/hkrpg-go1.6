// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: RogueCurrentInfo.proto

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

type RogueCurrentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueAvatarInfo  *RogueAvatarInfo  `protobuf:"bytes,15,opt,name=rogue_avatar_info,json=rogueAvatarInfo,proto3" json:"rogue_avatar_info,omitempty"`
	RoomMap          *RogueMapInfo     `protobuf:"bytes,4,opt,name=room_map,json=roomMap,proto3" json:"room_map,omitempty"`
	Status           RogueStatus       `protobuf:"varint,14,opt,name=status,proto3,enum=proto.RogueStatus" json:"status,omitempty"`
	RogueBuffInfo    *RogueBuffInfo    `protobuf:"bytes,6,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	RogueMiracleInfo *RogueMiracleInfo `protobuf:"bytes,10,opt,name=rogue_miracle_info,json=rogueMiracleInfo,proto3" json:"rogue_miracle_info,omitempty"`
}

func (x *RogueCurrentInfo) Reset() {
	*x = RogueCurrentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueCurrentInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueCurrentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueCurrentInfo) ProtoMessage() {}

func (x *RogueCurrentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueCurrentInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueCurrentInfo.ProtoReflect.Descriptor instead.
func (*RogueCurrentInfo) Descriptor() ([]byte, []int) {
	return file_RogueCurrentInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueCurrentInfo) GetRogueAvatarInfo() *RogueAvatarInfo {
	if x != nil {
		return x.RogueAvatarInfo
	}
	return nil
}

func (x *RogueCurrentInfo) GetRoomMap() *RogueMapInfo {
	if x != nil {
		return x.RoomMap
	}
	return nil
}

func (x *RogueCurrentInfo) GetStatus() RogueStatus {
	if x != nil {
		return x.Status
	}
	return RogueStatus_ROGUE_STATUS_NONE
}

func (x *RogueCurrentInfo) GetRogueBuffInfo() *RogueBuffInfo {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *RogueCurrentInfo) GetRogueMiracleInfo() *RogueMiracleInfo {
	if x != nil {
		return x.RogueMiracleInfo
	}
	return nil
}

var File_RogueCurrentInfo_proto protoreflect.FileDescriptor

var file_RogueCurrentInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x02, 0x0a, 0x10, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x42, 0x0a, 0x11, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d,
	0x4d, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3c, 0x0a, 0x0f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a,
	0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueCurrentInfo_proto_rawDescOnce sync.Once
	file_RogueCurrentInfo_proto_rawDescData = file_RogueCurrentInfo_proto_rawDesc
)

func file_RogueCurrentInfo_proto_rawDescGZIP() []byte {
	file_RogueCurrentInfo_proto_rawDescOnce.Do(func() {
		file_RogueCurrentInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueCurrentInfo_proto_rawDescData)
	})
	return file_RogueCurrentInfo_proto_rawDescData
}

var file_RogueCurrentInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueCurrentInfo_proto_goTypes = []interface{}{
	(*RogueCurrentInfo)(nil), // 0: proto.RogueCurrentInfo
	(*RogueAvatarInfo)(nil),  // 1: proto.RogueAvatarInfo
	(*RogueMapInfo)(nil),     // 2: proto.RogueMapInfo
	(RogueStatus)(0),         // 3: proto.RogueStatus
	(*RogueBuffInfo)(nil),    // 4: proto.RogueBuffInfo
	(*RogueMiracleInfo)(nil), // 5: proto.RogueMiracleInfo
}
var file_RogueCurrentInfo_proto_depIdxs = []int32{
	1, // 0: proto.RogueCurrentInfo.rogue_avatar_info:type_name -> proto.RogueAvatarInfo
	2, // 1: proto.RogueCurrentInfo.room_map:type_name -> proto.RogueMapInfo
	3, // 2: proto.RogueCurrentInfo.status:type_name -> proto.RogueStatus
	4, // 3: proto.RogueCurrentInfo.rogue_buff_info:type_name -> proto.RogueBuffInfo
	5, // 4: proto.RogueCurrentInfo.rogue_miracle_info:type_name -> proto.RogueMiracleInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_RogueCurrentInfo_proto_init() }
func file_RogueCurrentInfo_proto_init() {
	if File_RogueCurrentInfo_proto != nil {
		return
	}
	file_RogueBuffInfo_proto_init()
	file_RogueAvatarInfo_proto_init()
	file_RogueMapInfo_proto_init()
	file_RogueMiracleInfo_proto_init()
	file_RogueStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueCurrentInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueCurrentInfo); i {
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
			RawDescriptor: file_RogueCurrentInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueCurrentInfo_proto_goTypes,
		DependencyIndexes: file_RogueCurrentInfo_proto_depIdxs,
		MessageInfos:      file_RogueCurrentInfo_proto_msgTypes,
	}.Build()
	File_RogueCurrentInfo_proto = out.File
	file_RogueCurrentInfo_proto_rawDesc = nil
	file_RogueCurrentInfo_proto_goTypes = nil
	file_RogueCurrentInfo_proto_depIdxs = nil
}
