// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: SceneEntityMoveCsReq.proto

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

type SceneEntityMoveCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId          uint32          `protobuf:"varint,11,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	EntityMotionList []*EntityMotion `protobuf:"bytes,8,rep,name=entity_motion_list,json=entityMotionList,proto3" json:"entity_motion_list,omitempty"`
}

func (x *SceneEntityMoveCsReq) Reset() {
	*x = SceneEntityMoveCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEntityMoveCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityMoveCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityMoveCsReq) ProtoMessage() {}

func (x *SceneEntityMoveCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEntityMoveCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityMoveCsReq.ProtoReflect.Descriptor instead.
func (*SceneEntityMoveCsReq) Descriptor() ([]byte, []int) {
	return file_SceneEntityMoveCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityMoveCsReq) GetEntryId() uint32 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *SceneEntityMoveCsReq) GetEntityMotionList() []*EntityMotion {
	if x != nil {
		return x.EntityMotionList
	}
	return nil
}

var File_SceneEntityMoveCsReq_proto protoreflect.FileDescriptor

var file_SceneEntityMoveCsReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76,
	0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x14, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x12, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneEntityMoveCsReq_proto_rawDescOnce sync.Once
	file_SceneEntityMoveCsReq_proto_rawDescData = file_SceneEntityMoveCsReq_proto_rawDesc
)

func file_SceneEntityMoveCsReq_proto_rawDescGZIP() []byte {
	file_SceneEntityMoveCsReq_proto_rawDescOnce.Do(func() {
		file_SceneEntityMoveCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEntityMoveCsReq_proto_rawDescData)
	})
	return file_SceneEntityMoveCsReq_proto_rawDescData
}

var file_SceneEntityMoveCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEntityMoveCsReq_proto_goTypes = []interface{}{
	(*SceneEntityMoveCsReq)(nil), // 0: proto.SceneEntityMoveCsReq
	(*EntityMotion)(nil),         // 1: proto.EntityMotion
}
var file_SceneEntityMoveCsReq_proto_depIdxs = []int32{
	1, // 0: proto.SceneEntityMoveCsReq.entity_motion_list:type_name -> proto.EntityMotion
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneEntityMoveCsReq_proto_init() }
func file_SceneEntityMoveCsReq_proto_init() {
	if File_SceneEntityMoveCsReq_proto != nil {
		return
	}
	file_EntityMotion_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneEntityMoveCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityMoveCsReq); i {
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
			RawDescriptor: file_SceneEntityMoveCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEntityMoveCsReq_proto_goTypes,
		DependencyIndexes: file_SceneEntityMoveCsReq_proto_depIdxs,
		MessageInfos:      file_SceneEntityMoveCsReq_proto_msgTypes,
	}.Build()
	File_SceneEntityMoveCsReq_proto = out.File
	file_SceneEntityMoveCsReq_proto_rawDesc = nil
	file_SceneEntityMoveCsReq_proto_goTypes = nil
	file_SceneEntityMoveCsReq_proto_depIdxs = nil
}
