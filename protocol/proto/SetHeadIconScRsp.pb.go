// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: SetHeadIconScRsp.proto

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

type SetHeadIconScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentHeadIconId uint32 `protobuf:"varint,11,opt,name=current_head_icon_id,json=currentHeadIconId,proto3" json:"current_head_icon_id,omitempty"`
	Retcode           uint32 `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SetHeadIconScRsp) Reset() {
	*x = SetHeadIconScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetHeadIconScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHeadIconScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHeadIconScRsp) ProtoMessage() {}

func (x *SetHeadIconScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetHeadIconScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHeadIconScRsp.ProtoReflect.Descriptor instead.
func (*SetHeadIconScRsp) Descriptor() ([]byte, []int) {
	return file_SetHeadIconScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetHeadIconScRsp) GetCurrentHeadIconId() uint32 {
	if x != nil {
		return x.CurrentHeadIconId
	}
	return 0
}

func (x *SetHeadIconScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SetHeadIconScRsp_proto protoreflect.FileDescriptor

var file_SetHeadIconScRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x53, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5d, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x64, 0x49, 0x63,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_SetHeadIconScRsp_proto_rawDescOnce sync.Once
	file_SetHeadIconScRsp_proto_rawDescData = file_SetHeadIconScRsp_proto_rawDesc
)

func file_SetHeadIconScRsp_proto_rawDescGZIP() []byte {
	file_SetHeadIconScRsp_proto_rawDescOnce.Do(func() {
		file_SetHeadIconScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetHeadIconScRsp_proto_rawDescData)
	})
	return file_SetHeadIconScRsp_proto_rawDescData
}

var file_SetHeadIconScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetHeadIconScRsp_proto_goTypes = []interface{}{
	(*SetHeadIconScRsp)(nil), // 0: proto.SetHeadIconScRsp
}
var file_SetHeadIconScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetHeadIconScRsp_proto_init() }
func file_SetHeadIconScRsp_proto_init() {
	if File_SetHeadIconScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetHeadIconScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHeadIconScRsp); i {
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
			RawDescriptor: file_SetHeadIconScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetHeadIconScRsp_proto_goTypes,
		DependencyIndexes: file_SetHeadIconScRsp_proto_depIdxs,
		MessageInfos:      file_SetHeadIconScRsp_proto_msgTypes,
	}.Build()
	File_SetHeadIconScRsp_proto = out.File
	file_SetHeadIconScRsp_proto_rawDesc = nil
	file_SetHeadIconScRsp_proto_goTypes = nil
	file_SetHeadIconScRsp_proto_depIdxs = nil
}
