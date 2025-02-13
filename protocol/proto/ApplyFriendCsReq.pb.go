// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: ApplyFriendCsReq.proto

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

type ApplyFriendCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    uint32            `protobuf:"varint,11,opt,name=uid,proto3" json:"uid,omitempty"`
	Source FriendApplySource `protobuf:"varint,12,opt,name=source,proto3,enum=proto.FriendApplySource" json:"source,omitempty"`
}

func (x *ApplyFriendCsReq) Reset() {
	*x = ApplyFriendCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApplyFriendCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyFriendCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyFriendCsReq) ProtoMessage() {}

func (x *ApplyFriendCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ApplyFriendCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyFriendCsReq.ProtoReflect.Descriptor instead.
func (*ApplyFriendCsReq) Descriptor() ([]byte, []int) {
	return file_ApplyFriendCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *ApplyFriendCsReq) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ApplyFriendCsReq) GetSource() FriendApplySource {
	if x != nil {
		return x.Source
	}
	return FriendApplySource_FRIEND_APPLY_SOURCE_NONE
}

var File_ApplyFriendCsReq_proto protoreflect.FileDescriptor

var file_ApplyFriendCsReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x30,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ApplyFriendCsReq_proto_rawDescOnce sync.Once
	file_ApplyFriendCsReq_proto_rawDescData = file_ApplyFriendCsReq_proto_rawDesc
)

func file_ApplyFriendCsReq_proto_rawDescGZIP() []byte {
	file_ApplyFriendCsReq_proto_rawDescOnce.Do(func() {
		file_ApplyFriendCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ApplyFriendCsReq_proto_rawDescData)
	})
	return file_ApplyFriendCsReq_proto_rawDescData
}

var file_ApplyFriendCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ApplyFriendCsReq_proto_goTypes = []interface{}{
	(*ApplyFriendCsReq)(nil), // 0: proto.ApplyFriendCsReq
	(FriendApplySource)(0),   // 1: proto.FriendApplySource
}
var file_ApplyFriendCsReq_proto_depIdxs = []int32{
	1, // 0: proto.ApplyFriendCsReq.source:type_name -> proto.FriendApplySource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ApplyFriendCsReq_proto_init() }
func file_ApplyFriendCsReq_proto_init() {
	if File_ApplyFriendCsReq_proto != nil {
		return
	}
	file_FriendApplySource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ApplyFriendCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyFriendCsReq); i {
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
			RawDescriptor: file_ApplyFriendCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ApplyFriendCsReq_proto_goTypes,
		DependencyIndexes: file_ApplyFriendCsReq_proto_depIdxs,
		MessageInfos:      file_ApplyFriendCsReq_proto_msgTypes,
	}.Build()
	File_ApplyFriendCsReq_proto = out.File
	file_ApplyFriendCsReq_proto_rawDesc = nil
	file_ApplyFriendCsReq_proto_goTypes = nil
	file_ApplyFriendCsReq_proto_depIdxs = nil
}
