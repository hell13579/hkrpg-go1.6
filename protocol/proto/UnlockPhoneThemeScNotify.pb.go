// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: UnlockPhoneThemeScNotify.proto

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

type UnlockPhoneThemeScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThemeId uint32 `protobuf:"varint,15,opt,name=theme_id,json=themeId,proto3" json:"theme_id,omitempty"`
}

func (x *UnlockPhoneThemeScNotify) Reset() {
	*x = UnlockPhoneThemeScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnlockPhoneThemeScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockPhoneThemeScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockPhoneThemeScNotify) ProtoMessage() {}

func (x *UnlockPhoneThemeScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_UnlockPhoneThemeScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockPhoneThemeScNotify.ProtoReflect.Descriptor instead.
func (*UnlockPhoneThemeScNotify) Descriptor() ([]byte, []int) {
	return file_UnlockPhoneThemeScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *UnlockPhoneThemeScNotify) GetThemeId() uint32 {
	if x != nil {
		return x.ThemeId
	}
	return 0
}

var File_UnlockPhoneThemeScNotify_proto protoreflect.FileDescriptor

var file_UnlockPhoneThemeScNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x68, 0x65,
	0x6d, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x18, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x64, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_UnlockPhoneThemeScNotify_proto_rawDescOnce sync.Once
	file_UnlockPhoneThemeScNotify_proto_rawDescData = file_UnlockPhoneThemeScNotify_proto_rawDesc
)

func file_UnlockPhoneThemeScNotify_proto_rawDescGZIP() []byte {
	file_UnlockPhoneThemeScNotify_proto_rawDescOnce.Do(func() {
		file_UnlockPhoneThemeScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_UnlockPhoneThemeScNotify_proto_rawDescData)
	})
	return file_UnlockPhoneThemeScNotify_proto_rawDescData
}

var file_UnlockPhoneThemeScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UnlockPhoneThemeScNotify_proto_goTypes = []interface{}{
	(*UnlockPhoneThemeScNotify)(nil), // 0: proto.UnlockPhoneThemeScNotify
}
var file_UnlockPhoneThemeScNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_UnlockPhoneThemeScNotify_proto_init() }
func file_UnlockPhoneThemeScNotify_proto_init() {
	if File_UnlockPhoneThemeScNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UnlockPhoneThemeScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockPhoneThemeScNotify); i {
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
			RawDescriptor: file_UnlockPhoneThemeScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UnlockPhoneThemeScNotify_proto_goTypes,
		DependencyIndexes: file_UnlockPhoneThemeScNotify_proto_depIdxs,
		MessageInfos:      file_UnlockPhoneThemeScNotify_proto_msgTypes,
	}.Build()
	File_UnlockPhoneThemeScNotify_proto = out.File
	file_UnlockPhoneThemeScNotify_proto_rawDesc = nil
	file_UnlockPhoneThemeScNotify_proto_goTypes = nil
	file_UnlockPhoneThemeScNotify_proto_depIdxs = nil
}
