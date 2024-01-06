// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: SetSignatureCsReq.proto

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

type SetSignatureCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,15,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SetSignatureCsReq) Reset() {
	*x = SetSignatureCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetSignatureCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSignatureCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSignatureCsReq) ProtoMessage() {}

func (x *SetSignatureCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SetSignatureCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSignatureCsReq.ProtoReflect.Descriptor instead.
func (*SetSignatureCsReq) Descriptor() ([]byte, []int) {
	return file_SetSignatureCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SetSignatureCsReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_SetSignatureCsReq_proto protoreflect.FileDescriptor

var file_SetSignatureCsReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x53, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x73,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x31, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetSignatureCsReq_proto_rawDescOnce sync.Once
	file_SetSignatureCsReq_proto_rawDescData = file_SetSignatureCsReq_proto_rawDesc
)

func file_SetSignatureCsReq_proto_rawDescGZIP() []byte {
	file_SetSignatureCsReq_proto_rawDescOnce.Do(func() {
		file_SetSignatureCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetSignatureCsReq_proto_rawDescData)
	})
	return file_SetSignatureCsReq_proto_rawDescData
}

var file_SetSignatureCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetSignatureCsReq_proto_goTypes = []interface{}{
	(*SetSignatureCsReq)(nil), // 0: proto.SetSignatureCsReq
}
var file_SetSignatureCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetSignatureCsReq_proto_init() }
func file_SetSignatureCsReq_proto_init() {
	if File_SetSignatureCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetSignatureCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSignatureCsReq); i {
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
			RawDescriptor: file_SetSignatureCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetSignatureCsReq_proto_goTypes,
		DependencyIndexes: file_SetSignatureCsReq_proto_depIdxs,
		MessageInfos:      file_SetSignatureCsReq_proto_msgTypes,
	}.Build()
	File_SetSignatureCsReq_proto = out.File
	file_SetSignatureCsReq_proto_rawDesc = nil
	file_SetSignatureCsReq_proto_goTypes = nil
	file_SetSignatureCsReq_proto_depIdxs = nil
}
