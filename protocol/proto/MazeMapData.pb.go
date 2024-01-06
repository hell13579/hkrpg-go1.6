// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: MazeMapData.proto

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

type MazeMapData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnlockedChestList    []*MazeChest `protobuf:"bytes,11,rep,name=unlocked_chest_list,json=unlockedChestList,proto3" json:"unlocked_chest_list,omitempty"`
	MazeGroupList        []*MazeGroup `protobuf:"bytes,1,rep,name=maze_group_list,json=mazeGroupList,proto3" json:"maze_group_list,omitempty"`
	Retcode              uint32       `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MazePropList         []*MazeProp  `protobuf:"bytes,4,rep,name=maze_prop_list,json=mazePropList,proto3" json:"maze_prop_list,omitempty"`
	LightenSectionList   []uint32     `protobuf:"varint,2,rep,packed,name=lighten_section_list,json=lightenSectionList,proto3" json:"lighten_section_list,omitempty"`
	UnlockedTeleportList []uint32     `protobuf:"varint,7,rep,packed,name=unlocked_teleport_list,json=unlockedTeleportList,proto3" json:"unlocked_teleport_list,omitempty"`
	CurMapEntryId        uint32       `protobuf:"varint,10,opt,name=cur_map_entry_id,json=curMapEntryId,proto3" json:"cur_map_entry_id,omitempty"`
	EntryId              uint32       `protobuf:"varint,13,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *MazeMapData) Reset() {
	*x = MazeMapData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MazeMapData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MazeMapData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MazeMapData) ProtoMessage() {}

func (x *MazeMapData) ProtoReflect() protoreflect.Message {
	mi := &file_MazeMapData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MazeMapData.ProtoReflect.Descriptor instead.
func (*MazeMapData) Descriptor() ([]byte, []int) {
	return file_MazeMapData_proto_rawDescGZIP(), []int{0}
}

func (x *MazeMapData) GetUnlockedChestList() []*MazeChest {
	if x != nil {
		return x.UnlockedChestList
	}
	return nil
}

func (x *MazeMapData) GetMazeGroupList() []*MazeGroup {
	if x != nil {
		return x.MazeGroupList
	}
	return nil
}

func (x *MazeMapData) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *MazeMapData) GetMazePropList() []*MazeProp {
	if x != nil {
		return x.MazePropList
	}
	return nil
}

func (x *MazeMapData) GetLightenSectionList() []uint32 {
	if x != nil {
		return x.LightenSectionList
	}
	return nil
}

func (x *MazeMapData) GetUnlockedTeleportList() []uint32 {
	if x != nil {
		return x.UnlockedTeleportList
	}
	return nil
}

func (x *MazeMapData) GetCurMapEntryId() uint32 {
	if x != nil {
		return x.CurMapEntryId
	}
	return 0
}

func (x *MazeMapData) GetEntryId() uint32 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

var File_MazeMapData_proto protoreflect.FileDescriptor

var file_MazeMapData_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x61, 0x7a, 0x65, 0x4d, 0x61, 0x70, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x4d, 0x61, 0x7a, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x4d, 0x61, 0x7a,
	0x65, 0x43, 0x68, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x4d, 0x61,
	0x7a, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a,
	0x0b, 0x4d, 0x61, 0x7a, 0x65, 0x4d, 0x61, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x13,
	0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x61, 0x7a, 0x65, 0x43, 0x68, 0x65, 0x73, 0x74, 0x52, 0x11, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x43, 0x68, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x0f, 0x6d, 0x61, 0x7a, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x61, 0x7a, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d, 0x6d, 0x61, 0x7a, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x6d, 0x61, 0x7a, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x7a, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x0c, 0x6d, 0x61, 0x7a,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x6e,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x14, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x75, 0x72,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MazeMapData_proto_rawDescOnce sync.Once
	file_MazeMapData_proto_rawDescData = file_MazeMapData_proto_rawDesc
)

func file_MazeMapData_proto_rawDescGZIP() []byte {
	file_MazeMapData_proto_rawDescOnce.Do(func() {
		file_MazeMapData_proto_rawDescData = protoimpl.X.CompressGZIP(file_MazeMapData_proto_rawDescData)
	})
	return file_MazeMapData_proto_rawDescData
}

var file_MazeMapData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MazeMapData_proto_goTypes = []interface{}{
	(*MazeMapData)(nil), // 0: proto.MazeMapData
	(*MazeChest)(nil),   // 1: proto.MazeChest
	(*MazeGroup)(nil),   // 2: proto.MazeGroup
	(*MazeProp)(nil),    // 3: proto.MazeProp
}
var file_MazeMapData_proto_depIdxs = []int32{
	1, // 0: proto.MazeMapData.unlocked_chest_list:type_name -> proto.MazeChest
	2, // 1: proto.MazeMapData.maze_group_list:type_name -> proto.MazeGroup
	3, // 2: proto.MazeMapData.maze_prop_list:type_name -> proto.MazeProp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_MazeMapData_proto_init() }
func file_MazeMapData_proto_init() {
	if File_MazeMapData_proto != nil {
		return
	}
	file_MazeGroup_proto_init()
	file_MazeChest_proto_init()
	file_MazeProp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MazeMapData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MazeMapData); i {
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
			RawDescriptor: file_MazeMapData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MazeMapData_proto_goTypes,
		DependencyIndexes: file_MazeMapData_proto_depIdxs,
		MessageInfos:      file_MazeMapData_proto_msgTypes,
	}.Build()
	File_MazeMapData_proto = out.File
	file_MazeMapData_proto_rawDesc = nil
	file_MazeMapData_proto_goTypes = nil
	file_MazeMapData_proto_depIdxs = nil
}
