// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: MotionInfo.proto

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

type MotionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos *Vector `protobuf:"bytes,12,opt,name=pos,proto3" json:"pos,omitempty"`
	Rot *Vector `protobuf:"bytes,9,opt,name=rot,proto3" json:"rot,omitempty"`
}

func (x *MotionInfo) Reset() {
	*x = MotionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MotionInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MotionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MotionInfo) ProtoMessage() {}

func (x *MotionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_MotionInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MotionInfo.ProtoReflect.Descriptor instead.
func (*MotionInfo) Descriptor() ([]byte, []int) {
	return file_MotionInfo_proto_rawDescGZIP(), []int{0}
}

func (x *MotionInfo) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *MotionInfo) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

var File_MotionInfo_proto protoreflect.FileDescriptor

var file_MotionInfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x0a, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x72, 0x6f, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MotionInfo_proto_rawDescOnce sync.Once
	file_MotionInfo_proto_rawDescData = file_MotionInfo_proto_rawDesc
)

func file_MotionInfo_proto_rawDescGZIP() []byte {
	file_MotionInfo_proto_rawDescOnce.Do(func() {
		file_MotionInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_MotionInfo_proto_rawDescData)
	})
	return file_MotionInfo_proto_rawDescData
}

var file_MotionInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MotionInfo_proto_goTypes = []interface{}{
	(*MotionInfo)(nil), // 0: proto.MotionInfo
	(*Vector)(nil),     // 1: proto.Vector
}
var file_MotionInfo_proto_depIdxs = []int32{
	1, // 0: proto.MotionInfo.pos:type_name -> proto.Vector
	1, // 1: proto.MotionInfo.rot:type_name -> proto.Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MotionInfo_proto_init() }
func file_MotionInfo_proto_init() {
	if File_MotionInfo_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MotionInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MotionInfo); i {
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
			RawDescriptor: file_MotionInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MotionInfo_proto_goTypes,
		DependencyIndexes: file_MotionInfo_proto_depIdxs,
		MessageInfos:      file_MotionInfo_proto_msgTypes,
	}.Build()
	File_MotionInfo_proto = out.File
	file_MotionInfo_proto_rawDesc = nil
	file_MotionInfo_proto_goTypes = nil
	file_MotionInfo_proto_depIdxs = nil
}
