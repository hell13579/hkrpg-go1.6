// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: ComposeItemCsReq.proto

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

type ComposeItemCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count           uint32        `protobuf:"varint,12,opt,name=count,proto3" json:"count,omitempty"`
	ComposeItemList *ItemCostList `protobuf:"bytes,1,opt,name=compose_item_list,json=composeItemList,proto3" json:"compose_item_list,omitempty"`
	ComposeId       uint32        `protobuf:"varint,13,opt,name=compose_id,json=composeId,proto3" json:"compose_id,omitempty"`
}

func (x *ComposeItemCsReq) Reset() {
	*x = ComposeItemCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ComposeItemCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeItemCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeItemCsReq) ProtoMessage() {}

func (x *ComposeItemCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ComposeItemCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeItemCsReq.ProtoReflect.Descriptor instead.
func (*ComposeItemCsReq) Descriptor() ([]byte, []int) {
	return file_ComposeItemCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *ComposeItemCsReq) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ComposeItemCsReq) GetComposeItemList() *ItemCostList {
	if x != nil {
		return x.ComposeItemList
	}
	return nil
}

func (x *ComposeItemCsReq) GetComposeId() uint32 {
	if x != nil {
		return x.ComposeId
	}
	return 0
}

var File_ComposeItemCsReq_proto protoreflect.FileDescriptor

var file_ComposeItemCsReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3f,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x49, 0x64, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ComposeItemCsReq_proto_rawDescOnce sync.Once
	file_ComposeItemCsReq_proto_rawDescData = file_ComposeItemCsReq_proto_rawDesc
)

func file_ComposeItemCsReq_proto_rawDescGZIP() []byte {
	file_ComposeItemCsReq_proto_rawDescOnce.Do(func() {
		file_ComposeItemCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ComposeItemCsReq_proto_rawDescData)
	})
	return file_ComposeItemCsReq_proto_rawDescData
}

var file_ComposeItemCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ComposeItemCsReq_proto_goTypes = []interface{}{
	(*ComposeItemCsReq)(nil), // 0: proto.ComposeItemCsReq
	(*ItemCostList)(nil),     // 1: proto.ItemCostList
}
var file_ComposeItemCsReq_proto_depIdxs = []int32{
	1, // 0: proto.ComposeItemCsReq.compose_item_list:type_name -> proto.ItemCostList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ComposeItemCsReq_proto_init() }
func file_ComposeItemCsReq_proto_init() {
	if File_ComposeItemCsReq_proto != nil {
		return
	}
	file_ItemCostList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ComposeItemCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeItemCsReq); i {
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
			RawDescriptor: file_ComposeItemCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ComposeItemCsReq_proto_goTypes,
		DependencyIndexes: file_ComposeItemCsReq_proto_depIdxs,
		MessageInfos:      file_ComposeItemCsReq_proto_msgTypes,
	}.Build()
	File_ComposeItemCsReq_proto = out.File
	file_ComposeItemCsReq_proto_rawDesc = nil
	file_ComposeItemCsReq_proto_goTypes = nil
	file_ComposeItemCsReq_proto_depIdxs = nil
}
