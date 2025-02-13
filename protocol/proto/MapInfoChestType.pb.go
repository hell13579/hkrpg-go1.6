// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: MapInfoChestType.proto

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

type MapInfoChestType int32

const (
	MapInfoChestType_MAP_INFO_CHEST_TYPE_NONE      MapInfoChestType = 0
	MapInfoChestType_MAP_INFO_CHEST_TYPE_NORMAL    MapInfoChestType = 101
	MapInfoChestType_MAP_INFO_CHEST_TYPE_CHALLENGE MapInfoChestType = 102
	MapInfoChestType_MAP_INFO_CHEST_TYPE_PUZZLE    MapInfoChestType = 104
)

// Enum value maps for MapInfoChestType.
var (
	MapInfoChestType_name = map[int32]string{
		0:   "MAP_INFO_CHEST_TYPE_NONE",
		101: "MAP_INFO_CHEST_TYPE_NORMAL",
		102: "MAP_INFO_CHEST_TYPE_CHALLENGE",
		104: "MAP_INFO_CHEST_TYPE_PUZZLE",
	}
	MapInfoChestType_value = map[string]int32{
		"MAP_INFO_CHEST_TYPE_NONE":      0,
		"MAP_INFO_CHEST_TYPE_NORMAL":    101,
		"MAP_INFO_CHEST_TYPE_CHALLENGE": 102,
		"MAP_INFO_CHEST_TYPE_PUZZLE":    104,
	}
)

func (x MapInfoChestType) Enum() *MapInfoChestType {
	p := new(MapInfoChestType)
	*p = x
	return p
}

func (x MapInfoChestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MapInfoChestType) Descriptor() protoreflect.EnumDescriptor {
	return file_MapInfoChestType_proto_enumTypes[0].Descriptor()
}

func (MapInfoChestType) Type() protoreflect.EnumType {
	return &file_MapInfoChestType_proto_enumTypes[0]
}

func (x MapInfoChestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MapInfoChestType.Descriptor instead.
func (MapInfoChestType) EnumDescriptor() ([]byte, []int) {
	return file_MapInfoChestType_proto_rawDescGZIP(), []int{0}
}

var File_MapInfoChestType_proto protoreflect.FileDescriptor

var file_MapInfoChestType_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x65, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0x93, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x65, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x41, 0x50, 0x5f, 0x49, 0x4e, 0x46, 0x4f,
	0x5f, 0x43, 0x48, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x41, 0x50, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x43,
	0x48, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c,
	0x10, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x41, 0x50, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x43,
	0x48, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45,
	0x4e, 0x47, 0x45, 0x10, 0x66, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x41, 0x50, 0x5f, 0x49, 0x4e, 0x46,
	0x4f, 0x5f, 0x43, 0x48, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x55, 0x5a,
	0x5a, 0x4c, 0x45, 0x10, 0x68, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MapInfoChestType_proto_rawDescOnce sync.Once
	file_MapInfoChestType_proto_rawDescData = file_MapInfoChestType_proto_rawDesc
)

func file_MapInfoChestType_proto_rawDescGZIP() []byte {
	file_MapInfoChestType_proto_rawDescOnce.Do(func() {
		file_MapInfoChestType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MapInfoChestType_proto_rawDescData)
	})
	return file_MapInfoChestType_proto_rawDescData
}

var file_MapInfoChestType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MapInfoChestType_proto_goTypes = []interface{}{
	(MapInfoChestType)(0), // 0: proto.MapInfoChestType
}
var file_MapInfoChestType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MapInfoChestType_proto_init() }
func file_MapInfoChestType_proto_init() {
	if File_MapInfoChestType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MapInfoChestType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MapInfoChestType_proto_goTypes,
		DependencyIndexes: file_MapInfoChestType_proto_depIdxs,
		EnumInfos:         file_MapInfoChestType_proto_enumTypes,
	}.Build()
	File_MapInfoChestType_proto = out.File
	file_MapInfoChestType_proto_rawDesc = nil
	file_MapInfoChestType_proto_goTypes = nil
	file_MapInfoChestType_proto_depIdxs = nil
}
