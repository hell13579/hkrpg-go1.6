// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: MsgType.proto

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

type MsgType int32

const (
	MsgType_MSG_TYPE_NONE        MsgType = 0
	MsgType_MSG_TYPE_CUSTOM_TEXT MsgType = 1
	MsgType_MSG_TYPE_EMOJI       MsgType = 2
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "MSG_TYPE_NONE",
		1: "MSG_TYPE_CUSTOM_TEXT",
		2: "MSG_TYPE_EMOJI",
	}
	MsgType_value = map[string]int32{
		"MSG_TYPE_NONE":        0,
		"MSG_TYPE_CUSTOM_TEXT": 1,
		"MSG_TYPE_EMOJI":       2,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_MsgType_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_MsgType_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_MsgType_proto_rawDescGZIP(), []int{0}
}

var File_MsgType_proto protoreflect.FileDescriptor

var file_MsgType_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x4a, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d, 0x4f, 0x4a, 0x49,
	0x10, 0x02, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MsgType_proto_rawDescOnce sync.Once
	file_MsgType_proto_rawDescData = file_MsgType_proto_rawDesc
)

func file_MsgType_proto_rawDescGZIP() []byte {
	file_MsgType_proto_rawDescOnce.Do(func() {
		file_MsgType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MsgType_proto_rawDescData)
	})
	return file_MsgType_proto_rawDescData
}

var file_MsgType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MsgType_proto_goTypes = []interface{}{
	(MsgType)(0), // 0: proto.MsgType
}
var file_MsgType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MsgType_proto_init() }
func file_MsgType_proto_init() {
	if File_MsgType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MsgType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MsgType_proto_goTypes,
		DependencyIndexes: file_MsgType_proto_depIdxs,
		EnumInfos:         file_MsgType_proto_enumTypes,
	}.Build()
	File_MsgType_proto = out.File
	file_MsgType_proto_rawDesc = nil
	file_MsgType_proto_goTypes = nil
	file_MsgType_proto_depIdxs = nil
}
