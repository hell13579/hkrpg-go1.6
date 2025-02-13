// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: NewMailScNotify.proto

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

type NewMailScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailIdList []uint32 `protobuf:"varint,10,rep,packed,name=mail_id_list,json=mailIdList,proto3" json:"mail_id_list,omitempty"`
}

func (x *NewMailScNotify) Reset() {
	*x = NewMailScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NewMailScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMailScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMailScNotify) ProtoMessage() {}

func (x *NewMailScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_NewMailScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMailScNotify.ProtoReflect.Descriptor instead.
func (*NewMailScNotify) Descriptor() ([]byte, []int) {
	return file_NewMailScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *NewMailScNotify) GetMailIdList() []uint32 {
	if x != nil {
		return x.MailIdList
	}
	return nil
}

var File_NewMailScNotify_proto protoreflect.FileDescriptor

var file_NewMailScNotify_proto_rawDesc = []byte{
	0x0a, 0x15, 0x4e, 0x65, 0x77, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33,
	0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NewMailScNotify_proto_rawDescOnce sync.Once
	file_NewMailScNotify_proto_rawDescData = file_NewMailScNotify_proto_rawDesc
)

func file_NewMailScNotify_proto_rawDescGZIP() []byte {
	file_NewMailScNotify_proto_rawDescOnce.Do(func() {
		file_NewMailScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_NewMailScNotify_proto_rawDescData)
	})
	return file_NewMailScNotify_proto_rawDescData
}

var file_NewMailScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NewMailScNotify_proto_goTypes = []interface{}{
	(*NewMailScNotify)(nil), // 0: proto.NewMailScNotify
}
var file_NewMailScNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_NewMailScNotify_proto_init() }
func file_NewMailScNotify_proto_init() {
	if File_NewMailScNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NewMailScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMailScNotify); i {
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
			RawDescriptor: file_NewMailScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NewMailScNotify_proto_goTypes,
		DependencyIndexes: file_NewMailScNotify_proto_depIdxs,
		MessageInfos:      file_NewMailScNotify_proto_msgTypes,
	}.Build()
	File_NewMailScNotify_proto = out.File
	file_NewMailScNotify_proto_rawDesc = nil
	file_NewMailScNotify_proto_goTypes = nil
	file_NewMailScNotify_proto_depIdxs = nil
}
