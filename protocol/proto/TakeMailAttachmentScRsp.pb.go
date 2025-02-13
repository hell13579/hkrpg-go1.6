// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: TakeMailAttachmentScRsp.proto

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

type TakeMailAttachmentScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccMailIdList []uint32  `protobuf:"varint,4,rep,packed,name=succ_mail_id_list,json=succMailIdList,proto3" json:"succ_mail_id_list,omitempty"`
	Attachment     *ItemList `protobuf:"bytes,12,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Retcode        uint32    `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *TakeMailAttachmentScRsp) Reset() {
	*x = TakeMailAttachmentScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeMailAttachmentScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeMailAttachmentScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeMailAttachmentScRsp) ProtoMessage() {}

func (x *TakeMailAttachmentScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TakeMailAttachmentScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeMailAttachmentScRsp.ProtoReflect.Descriptor instead.
func (*TakeMailAttachmentScRsp) Descriptor() ([]byte, []int) {
	return file_TakeMailAttachmentScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TakeMailAttachmentScRsp) GetSuccMailIdList() []uint32 {
	if x != nil {
		return x.SuccMailIdList
	}
	return nil
}

func (x *TakeMailAttachmentScRsp) GetAttachment() *ItemList {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *TakeMailAttachmentScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_TakeMailAttachmentScRsp_proto protoreflect.FileDescriptor

var file_TakeMailAttachmentScRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x54, 0x61, 0x6b, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x17, 0x54, 0x61, 0x6b, 0x65, 0x4d,
	0x61, 0x69, 0x6c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x12, 0x29, 0x0a, 0x11, 0x73, 0x75, 0x63, 0x63, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x73,
	0x75, 0x63, 0x63, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TakeMailAttachmentScRsp_proto_rawDescOnce sync.Once
	file_TakeMailAttachmentScRsp_proto_rawDescData = file_TakeMailAttachmentScRsp_proto_rawDesc
)

func file_TakeMailAttachmentScRsp_proto_rawDescGZIP() []byte {
	file_TakeMailAttachmentScRsp_proto_rawDescOnce.Do(func() {
		file_TakeMailAttachmentScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeMailAttachmentScRsp_proto_rawDescData)
	})
	return file_TakeMailAttachmentScRsp_proto_rawDescData
}

var file_TakeMailAttachmentScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeMailAttachmentScRsp_proto_goTypes = []interface{}{
	(*TakeMailAttachmentScRsp)(nil), // 0: proto.TakeMailAttachmentScRsp
	(*ItemList)(nil),                // 1: proto.ItemList
}
var file_TakeMailAttachmentScRsp_proto_depIdxs = []int32{
	1, // 0: proto.TakeMailAttachmentScRsp.attachment:type_name -> proto.ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TakeMailAttachmentScRsp_proto_init() }
func file_TakeMailAttachmentScRsp_proto_init() {
	if File_TakeMailAttachmentScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TakeMailAttachmentScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeMailAttachmentScRsp); i {
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
			RawDescriptor: file_TakeMailAttachmentScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeMailAttachmentScRsp_proto_goTypes,
		DependencyIndexes: file_TakeMailAttachmentScRsp_proto_depIdxs,
		MessageInfos:      file_TakeMailAttachmentScRsp_proto_msgTypes,
	}.Build()
	File_TakeMailAttachmentScRsp_proto = out.File
	file_TakeMailAttachmentScRsp_proto_rawDesc = nil
	file_TakeMailAttachmentScRsp_proto_goTypes = nil
	file_TakeMailAttachmentScRsp_proto_depIdxs = nil
}
