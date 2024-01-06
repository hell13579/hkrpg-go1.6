// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: ScenePlaneEventScNotify.proto

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

type ScenePlaneEventScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetItemList *ItemList `protobuf:"bytes,8,opt,name=get_item_list,json=getItemList,proto3" json:"get_item_list,omitempty"`
}

func (x *ScenePlaneEventScNotify) Reset() {
	*x = ScenePlaneEventScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ScenePlaneEventScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScenePlaneEventScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScenePlaneEventScNotify) ProtoMessage() {}

func (x *ScenePlaneEventScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ScenePlaneEventScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScenePlaneEventScNotify.ProtoReflect.Descriptor instead.
func (*ScenePlaneEventScNotify) Descriptor() ([]byte, []int) {
	return file_ScenePlaneEventScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ScenePlaneEventScNotify) GetGetItemList() *ItemList {
	if x != nil {
		return x.GetItemList
	}
	return nil
}

var File_ScenePlaneEventScNotify_proto protoreflect.FileDescriptor

var file_ScenePlaneEventScNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x17, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x33, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x67, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ScenePlaneEventScNotify_proto_rawDescOnce sync.Once
	file_ScenePlaneEventScNotify_proto_rawDescData = file_ScenePlaneEventScNotify_proto_rawDesc
)

func file_ScenePlaneEventScNotify_proto_rawDescGZIP() []byte {
	file_ScenePlaneEventScNotify_proto_rawDescOnce.Do(func() {
		file_ScenePlaneEventScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ScenePlaneEventScNotify_proto_rawDescData)
	})
	return file_ScenePlaneEventScNotify_proto_rawDescData
}

var file_ScenePlaneEventScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ScenePlaneEventScNotify_proto_goTypes = []interface{}{
	(*ScenePlaneEventScNotify)(nil), // 0: proto.ScenePlaneEventScNotify
	(*ItemList)(nil),                // 1: proto.ItemList
}
var file_ScenePlaneEventScNotify_proto_depIdxs = []int32{
	1, // 0: proto.ScenePlaneEventScNotify.get_item_list:type_name -> proto.ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ScenePlaneEventScNotify_proto_init() }
func file_ScenePlaneEventScNotify_proto_init() {
	if File_ScenePlaneEventScNotify_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ScenePlaneEventScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScenePlaneEventScNotify); i {
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
			RawDescriptor: file_ScenePlaneEventScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ScenePlaneEventScNotify_proto_goTypes,
		DependencyIndexes: file_ScenePlaneEventScNotify_proto_depIdxs,
		MessageInfos:      file_ScenePlaneEventScNotify_proto_msgTypes,
	}.Build()
	File_ScenePlaneEventScNotify_proto = out.File
	file_ScenePlaneEventScNotify_proto_rawDesc = nil
	file_ScenePlaneEventScNotify_proto_goTypes = nil
	file_ScenePlaneEventScNotify_proto_depIdxs = nil
}
