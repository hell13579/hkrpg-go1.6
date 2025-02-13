// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: SceneCastSkill.proto

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

type SceneCastSkill int32

const (
	SceneCastSkill_SCENE_CAST_SKILL_NONE                       SceneCastSkill = 0
	SceneCastSkill_SCENE_CAST_SKILL_PROJECTILE_HIT             SceneCastSkill = 1
	SceneCastSkill_SCENE_CAST_SKILL_PROJECTILE_LIFETIME_FINISH SceneCastSkill = 2
)

// Enum value maps for SceneCastSkill.
var (
	SceneCastSkill_name = map[int32]string{
		0: "SCENE_CAST_SKILL_NONE",
		1: "SCENE_CAST_SKILL_PROJECTILE_HIT",
		2: "SCENE_CAST_SKILL_PROJECTILE_LIFETIME_FINISH",
	}
	SceneCastSkill_value = map[string]int32{
		"SCENE_CAST_SKILL_NONE":                       0,
		"SCENE_CAST_SKILL_PROJECTILE_HIT":             1,
		"SCENE_CAST_SKILL_PROJECTILE_LIFETIME_FINISH": 2,
	}
)

func (x SceneCastSkill) Enum() *SceneCastSkill {
	p := new(SceneCastSkill)
	*p = x
	return p
}

func (x SceneCastSkill) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneCastSkill) Descriptor() protoreflect.EnumDescriptor {
	return file_SceneCastSkill_proto_enumTypes[0].Descriptor()
}

func (SceneCastSkill) Type() protoreflect.EnumType {
	return &file_SceneCastSkill_proto_enumTypes[0]
}

func (x SceneCastSkill) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneCastSkill.Descriptor instead.
func (SceneCastSkill) EnumDescriptor() ([]byte, []int) {
	return file_SceneCastSkill_proto_rawDescGZIP(), []int{0}
}

var File_SceneCastSkill_proto protoreflect.FileDescriptor

var file_SceneCastSkill_proto_rawDesc = []byte{
	0x0a, 0x14, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x61, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x81, 0x01,
	0x0a, 0x0e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x61, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x12, 0x19, 0x0a, 0x15, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x53,
	0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x53,
	0x43, 0x45, 0x4e, 0x45, 0x5f, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f,
	0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x49, 0x4c, 0x45, 0x5f, 0x48, 0x49, 0x54, 0x10, 0x01,
	0x12, 0x2f, 0x0a, 0x2b, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x53,
	0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x49, 0x4c, 0x45, 0x5f,
	0x4c, 0x49, 0x46, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x10,
	0x02, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneCastSkill_proto_rawDescOnce sync.Once
	file_SceneCastSkill_proto_rawDescData = file_SceneCastSkill_proto_rawDesc
)

func file_SceneCastSkill_proto_rawDescGZIP() []byte {
	file_SceneCastSkill_proto_rawDescOnce.Do(func() {
		file_SceneCastSkill_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneCastSkill_proto_rawDescData)
	})
	return file_SceneCastSkill_proto_rawDescData
}

var file_SceneCastSkill_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SceneCastSkill_proto_goTypes = []interface{}{
	(SceneCastSkill)(0), // 0: proto.SceneCastSkill
}
var file_SceneCastSkill_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneCastSkill_proto_init() }
func file_SceneCastSkill_proto_init() {
	if File_SceneCastSkill_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SceneCastSkill_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneCastSkill_proto_goTypes,
		DependencyIndexes: file_SceneCastSkill_proto_depIdxs,
		EnumInfos:         file_SceneCastSkill_proto_enumTypes,
	}.Build()
	File_SceneCastSkill_proto = out.File
	file_SceneCastSkill_proto_rawDesc = nil
	file_SceneCastSkill_proto_goTypes = nil
	file_SceneCastSkill_proto_depIdxs = nil
}
