// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: SceneCastSkillScRsp.proto

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

type SceneCastSkillScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttackedGroupId uint32           `protobuf:"varint,5,opt,name=attacked_group_id,json=attackedGroupId,proto3" json:"attacked_group_id,omitempty"`
	Retcode         uint32           `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BattleInfo      *SceneBattleInfo `protobuf:"bytes,11,opt,name=battle_info,json=battleInfo,proto3" json:"battle_info,omitempty"`
}

func (x *SceneCastSkillScRsp) Reset() {
	*x = SceneCastSkillScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneCastSkillScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneCastSkillScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneCastSkillScRsp) ProtoMessage() {}

func (x *SceneCastSkillScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SceneCastSkillScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneCastSkillScRsp.ProtoReflect.Descriptor instead.
func (*SceneCastSkillScRsp) Descriptor() ([]byte, []int) {
	return file_SceneCastSkillScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SceneCastSkillScRsp) GetAttackedGroupId() uint32 {
	if x != nil {
		return x.AttackedGroupId
	}
	return 0
}

func (x *SceneCastSkillScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SceneCastSkillScRsp) GetBattleInfo() *SceneBattleInfo {
	if x != nil {
		return x.BattleInfo
	}
	return nil
}

var File_SceneCastSkillScRsp_proto protoreflect.FileDescriptor

var file_SceneCastSkillScRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x61, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x43, 0x61, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x63, 0x52, 0x73,
	0x70, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_SceneCastSkillScRsp_proto_rawDescOnce sync.Once
	file_SceneCastSkillScRsp_proto_rawDescData = file_SceneCastSkillScRsp_proto_rawDesc
)

func file_SceneCastSkillScRsp_proto_rawDescGZIP() []byte {
	file_SceneCastSkillScRsp_proto_rawDescOnce.Do(func() {
		file_SceneCastSkillScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneCastSkillScRsp_proto_rawDescData)
	})
	return file_SceneCastSkillScRsp_proto_rawDescData
}

var file_SceneCastSkillScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneCastSkillScRsp_proto_goTypes = []interface{}{
	(*SceneCastSkillScRsp)(nil), // 0: proto.SceneCastSkillScRsp
	(*SceneBattleInfo)(nil),     // 1: proto.SceneBattleInfo
}
var file_SceneCastSkillScRsp_proto_depIdxs = []int32{
	1, // 0: proto.SceneCastSkillScRsp.battle_info:type_name -> proto.SceneBattleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneCastSkillScRsp_proto_init() }
func file_SceneCastSkillScRsp_proto_init() {
	if File_SceneCastSkillScRsp_proto != nil {
		return
	}
	file_SceneBattleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneCastSkillScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneCastSkillScRsp); i {
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
			RawDescriptor: file_SceneCastSkillScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneCastSkillScRsp_proto_goTypes,
		DependencyIndexes: file_SceneCastSkillScRsp_proto_depIdxs,
		MessageInfos:      file_SceneCastSkillScRsp_proto_msgTypes,
	}.Build()
	File_SceneCastSkillScRsp_proto = out.File
	file_SceneCastSkillScRsp_proto_rawDesc = nil
	file_SceneCastSkillScRsp_proto_goTypes = nil
	file_SceneCastSkillScRsp_proto_depIdxs = nil
}
