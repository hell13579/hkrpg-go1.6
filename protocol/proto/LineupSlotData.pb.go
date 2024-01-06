// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: LineupSlotData.proto

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

type LineupSlotData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32     `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	AvatarType AvatarType `protobuf:"varint,2,opt,name=avatar_type,json=avatarType,proto3,enum=proto.AvatarType" json:"avatar_type,omitempty"`
	Slot       uint32     `protobuf:"varint,8,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *LineupSlotData) Reset() {
	*x = LineupSlotData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LineupSlotData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineupSlotData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineupSlotData) ProtoMessage() {}

func (x *LineupSlotData) ProtoReflect() protoreflect.Message {
	mi := &file_LineupSlotData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineupSlotData.ProtoReflect.Descriptor instead.
func (*LineupSlotData) Descriptor() ([]byte, []int) {
	return file_LineupSlotData_proto_rawDescGZIP(), []int{0}
}

func (x *LineupSlotData) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LineupSlotData) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *LineupSlotData) GetSlot() uint32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

var File_LineupSlotData_proto protoreflect.FileDescriptor

var file_LineupSlotData_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x53, 0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x68, 0x0a, 0x0e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x53, 0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x32, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LineupSlotData_proto_rawDescOnce sync.Once
	file_LineupSlotData_proto_rawDescData = file_LineupSlotData_proto_rawDesc
)

func file_LineupSlotData_proto_rawDescGZIP() []byte {
	file_LineupSlotData_proto_rawDescOnce.Do(func() {
		file_LineupSlotData_proto_rawDescData = protoimpl.X.CompressGZIP(file_LineupSlotData_proto_rawDescData)
	})
	return file_LineupSlotData_proto_rawDescData
}

var file_LineupSlotData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LineupSlotData_proto_goTypes = []interface{}{
	(*LineupSlotData)(nil), // 0: proto.LineupSlotData
	(AvatarType)(0),        // 1: proto.AvatarType
}
var file_LineupSlotData_proto_depIdxs = []int32{
	1, // 0: proto.LineupSlotData.avatar_type:type_name -> proto.AvatarType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LineupSlotData_proto_init() }
func file_LineupSlotData_proto_init() {
	if File_LineupSlotData_proto != nil {
		return
	}
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LineupSlotData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineupSlotData); i {
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
			RawDescriptor: file_LineupSlotData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LineupSlotData_proto_goTypes,
		DependencyIndexes: file_LineupSlotData_proto_depIdxs,
		MessageInfos:      file_LineupSlotData_proto_msgTypes,
	}.Build()
	File_LineupSlotData_proto = out.File
	file_LineupSlotData_proto_rawDesc = nil
	file_LineupSlotData_proto_goTypes = nil
	file_LineupSlotData_proto_depIdxs = nil
}
