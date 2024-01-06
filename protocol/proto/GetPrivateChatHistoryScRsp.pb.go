// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: GetPrivateChatHistoryScRsp.proto

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

type GetPrivateChatHistoryScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderUid uint32  `protobuf:"varint,4,opt,name=sender_uid,json=senderUid,proto3" json:"sender_uid,omitempty"`
	ToUid     uint32  `protobuf:"varint,9,opt,name=to_uid,json=toUid,proto3" json:"to_uid,omitempty"`
	Retcode   uint32  `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ChatList  []*Chat `protobuf:"bytes,3,rep,name=chat_list,json=chatList,proto3" json:"chat_list,omitempty"`
}

func (x *GetPrivateChatHistoryScRsp) Reset() {
	*x = GetPrivateChatHistoryScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPrivateChatHistoryScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateChatHistoryScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateChatHistoryScRsp) ProtoMessage() {}

func (x *GetPrivateChatHistoryScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPrivateChatHistoryScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateChatHistoryScRsp.ProtoReflect.Descriptor instead.
func (*GetPrivateChatHistoryScRsp) Descriptor() ([]byte, []int) {
	return file_GetPrivateChatHistoryScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPrivateChatHistoryScRsp) GetSenderUid() uint32 {
	if x != nil {
		return x.SenderUid
	}
	return 0
}

func (x *GetPrivateChatHistoryScRsp) GetToUid() uint32 {
	if x != nil {
		return x.ToUid
	}
	return 0
}

func (x *GetPrivateChatHistoryScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetPrivateChatHistoryScRsp) GetChatList() []*Chat {
	if x != nil {
		return x.ChatList
	}
	return nil
}

var File_GetPrivateChatHistoryScRsp_proto protoreflect.FileDescriptor

var file_GetPrivateChatHistoryScRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x55, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_GetPrivateChatHistoryScRsp_proto_rawDescOnce sync.Once
	file_GetPrivateChatHistoryScRsp_proto_rawDescData = file_GetPrivateChatHistoryScRsp_proto_rawDesc
)

func file_GetPrivateChatHistoryScRsp_proto_rawDescGZIP() []byte {
	file_GetPrivateChatHistoryScRsp_proto_rawDescOnce.Do(func() {
		file_GetPrivateChatHistoryScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPrivateChatHistoryScRsp_proto_rawDescData)
	})
	return file_GetPrivateChatHistoryScRsp_proto_rawDescData
}

var file_GetPrivateChatHistoryScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPrivateChatHistoryScRsp_proto_goTypes = []interface{}{
	(*GetPrivateChatHistoryScRsp)(nil), // 0: proto.GetPrivateChatHistoryScRsp
	(*Chat)(nil),                       // 1: proto.Chat
}
var file_GetPrivateChatHistoryScRsp_proto_depIdxs = []int32{
	1, // 0: proto.GetPrivateChatHistoryScRsp.chat_list:type_name -> proto.Chat
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetPrivateChatHistoryScRsp_proto_init() }
func file_GetPrivateChatHistoryScRsp_proto_init() {
	if File_GetPrivateChatHistoryScRsp_proto != nil {
		return
	}
	file_Chat_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetPrivateChatHistoryScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrivateChatHistoryScRsp); i {
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
			RawDescriptor: file_GetPrivateChatHistoryScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPrivateChatHistoryScRsp_proto_goTypes,
		DependencyIndexes: file_GetPrivateChatHistoryScRsp_proto_depIdxs,
		MessageInfos:      file_GetPrivateChatHistoryScRsp_proto_msgTypes,
	}.Build()
	File_GetPrivateChatHistoryScRsp_proto = out.File
	file_GetPrivateChatHistoryScRsp_proto_rawDesc = nil
	file_GetPrivateChatHistoryScRsp_proto_goTypes = nil
	file_GetPrivateChatHistoryScRsp_proto_depIdxs = nil
}
