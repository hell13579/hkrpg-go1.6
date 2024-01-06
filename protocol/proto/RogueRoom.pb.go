// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: RogueRoom.proto

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

type RogueRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomStatus RogueRoomStatus `protobuf:"varint,2,opt,name=room_status,json=roomStatus,proto3,enum=proto.RogueRoomStatus" json:"room_status,omitempty"`
	SiteId     uint32          `protobuf:"varint,14,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
	RoomId     uint32          `protobuf:"varint,7,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *RogueRoom) Reset() {
	*x = RogueRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueRoom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueRoom) ProtoMessage() {}

func (x *RogueRoom) ProtoReflect() protoreflect.Message {
	mi := &file_RogueRoom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueRoom.ProtoReflect.Descriptor instead.
func (*RogueRoom) Descriptor() ([]byte, []int) {
	return file_RogueRoom_proto_rawDescGZIP(), []int{0}
}

func (x *RogueRoom) GetRoomStatus() RogueRoomStatus {
	if x != nil {
		return x.RoomStatus
	}
	return RogueRoomStatus_ROGUE_ROOM_STATUS_NONE
}

func (x *RogueRoom) GetSiteId() uint32 {
	if x != nil {
		return x.SiteId
	}
	return 0
}

func (x *RogueRoom) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

var File_RogueRoom_proto protoreflect.FileDescriptor

var file_RogueRoom_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x76, 0x0a, 0x09, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x37, 0x0a, 0x0b,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueRoom_proto_rawDescOnce sync.Once
	file_RogueRoom_proto_rawDescData = file_RogueRoom_proto_rawDesc
)

func file_RogueRoom_proto_rawDescGZIP() []byte {
	file_RogueRoom_proto_rawDescOnce.Do(func() {
		file_RogueRoom_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueRoom_proto_rawDescData)
	})
	return file_RogueRoom_proto_rawDescData
}

var file_RogueRoom_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueRoom_proto_goTypes = []interface{}{
	(*RogueRoom)(nil),    // 0: proto.RogueRoom
	(RogueRoomStatus)(0), // 1: proto.RogueRoomStatus
}
var file_RogueRoom_proto_depIdxs = []int32{
	1, // 0: proto.RogueRoom.room_status:type_name -> proto.RogueRoomStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueRoom_proto_init() }
func file_RogueRoom_proto_init() {
	if File_RogueRoom_proto != nil {
		return
	}
	file_RogueRoomStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueRoom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueRoom); i {
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
			RawDescriptor: file_RogueRoom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueRoom_proto_goTypes,
		DependencyIndexes: file_RogueRoom_proto_depIdxs,
		MessageInfos:      file_RogueRoom_proto_msgTypes,
	}.Build()
	File_RogueRoom_proto = out.File
	file_RogueRoom_proto_rawDesc = nil
	file_RogueRoom_proto_goTypes = nil
	file_RogueRoom_proto_depIdxs = nil
}
