// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: EnterRogueMapRoomScRsp.proto

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

type EnterRogueMapRoomScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lineup    *LineupInfo `protobuf:"bytes,2,opt,name=lineup,proto3" json:"lineup,omitempty"`
	CurSiteId uint32      `protobuf:"varint,8,opt,name=cur_site_id,json=curSiteId,proto3" json:"cur_site_id,omitempty"`
	Retcode   uint32      `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Scene     *SceneInfo  `protobuf:"bytes,15,opt,name=scene,proto3" json:"scene,omitempty"`
}

func (x *EnterRogueMapRoomScRsp) Reset() {
	*x = EnterRogueMapRoomScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterRogueMapRoomScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterRogueMapRoomScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterRogueMapRoomScRsp) ProtoMessage() {}

func (x *EnterRogueMapRoomScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_EnterRogueMapRoomScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterRogueMapRoomScRsp.ProtoReflect.Descriptor instead.
func (*EnterRogueMapRoomScRsp) Descriptor() ([]byte, []int) {
	return file_EnterRogueMapRoomScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *EnterRogueMapRoomScRsp) GetLineup() *LineupInfo {
	if x != nil {
		return x.Lineup
	}
	return nil
}

func (x *EnterRogueMapRoomScRsp) GetCurSiteId() uint32 {
	if x != nil {
		return x.CurSiteId
	}
	return 0
}

func (x *EnterRogueMapRoomScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *EnterRogueMapRoomScRsp) GetScene() *SceneInfo {
	if x != nil {
		return x.Scene
	}
	return nil
}

var File_EnterRogueMapRoomScRsp_proto protoreflect.FileDescriptor

var file_EnterRogueMapRoomScRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x52,
	0x6f, 0x6f, 0x6d, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x16, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x12, 0x1e,
	0x0a, 0x0b, 0x63, 0x75, 0x72, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x75, 0x72, 0x53, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_EnterRogueMapRoomScRsp_proto_rawDescOnce sync.Once
	file_EnterRogueMapRoomScRsp_proto_rawDescData = file_EnterRogueMapRoomScRsp_proto_rawDesc
)

func file_EnterRogueMapRoomScRsp_proto_rawDescGZIP() []byte {
	file_EnterRogueMapRoomScRsp_proto_rawDescOnce.Do(func() {
		file_EnterRogueMapRoomScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterRogueMapRoomScRsp_proto_rawDescData)
	})
	return file_EnterRogueMapRoomScRsp_proto_rawDescData
}

var file_EnterRogueMapRoomScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterRogueMapRoomScRsp_proto_goTypes = []interface{}{
	(*EnterRogueMapRoomScRsp)(nil), // 0: proto.EnterRogueMapRoomScRsp
	(*LineupInfo)(nil),             // 1: proto.LineupInfo
	(*SceneInfo)(nil),              // 2: proto.SceneInfo
}
var file_EnterRogueMapRoomScRsp_proto_depIdxs = []int32{
	1, // 0: proto.EnterRogueMapRoomScRsp.lineup:type_name -> proto.LineupInfo
	2, // 1: proto.EnterRogueMapRoomScRsp.scene:type_name -> proto.SceneInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EnterRogueMapRoomScRsp_proto_init() }
func file_EnterRogueMapRoomScRsp_proto_init() {
	if File_EnterRogueMapRoomScRsp_proto != nil {
		return
	}
	file_LineupInfo_proto_init()
	file_SceneInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EnterRogueMapRoomScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterRogueMapRoomScRsp); i {
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
			RawDescriptor: file_EnterRogueMapRoomScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterRogueMapRoomScRsp_proto_goTypes,
		DependencyIndexes: file_EnterRogueMapRoomScRsp_proto_depIdxs,
		MessageInfos:      file_EnterRogueMapRoomScRsp_proto_msgTypes,
	}.Build()
	File_EnterRogueMapRoomScRsp_proto = out.File
	file_EnterRogueMapRoomScRsp_proto_rawDesc = nil
	file_EnterRogueMapRoomScRsp_proto_goTypes = nil
	file_EnterRogueMapRoomScRsp_proto_depIdxs = nil
}
