// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: GetCurLineupDataScRsp.proto

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

type GetCurLineupDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode uint32      `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Lineup  *LineupInfo `protobuf:"bytes,2,opt,name=lineup,proto3" json:"lineup,omitempty"`
}

func (x *GetCurLineupDataScRsp) Reset() {
	*x = GetCurLineupDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetCurLineupDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurLineupDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurLineupDataScRsp) ProtoMessage() {}

func (x *GetCurLineupDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetCurLineupDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurLineupDataScRsp.ProtoReflect.Descriptor instead.
func (*GetCurLineupDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetCurLineupDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurLineupDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetCurLineupDataScRsp) GetLineup() *LineupInfo {
	if x != nil {
		return x.Lineup
	}
	return nil
}

var File_GetCurLineupDataScRsp_proto protoreflect.FileDescriptor

var file_GetCurLineupDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x69, 0x6e,
	0x65, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x69,
	0x6e, 0x65, 0x75, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetCurLineupDataScRsp_proto_rawDescOnce sync.Once
	file_GetCurLineupDataScRsp_proto_rawDescData = file_GetCurLineupDataScRsp_proto_rawDesc
)

func file_GetCurLineupDataScRsp_proto_rawDescGZIP() []byte {
	file_GetCurLineupDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetCurLineupDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetCurLineupDataScRsp_proto_rawDescData)
	})
	return file_GetCurLineupDataScRsp_proto_rawDescData
}

var file_GetCurLineupDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetCurLineupDataScRsp_proto_goTypes = []interface{}{
	(*GetCurLineupDataScRsp)(nil), // 0: proto.GetCurLineupDataScRsp
	(*LineupInfo)(nil),            // 1: proto.LineupInfo
}
var file_GetCurLineupDataScRsp_proto_depIdxs = []int32{
	1, // 0: proto.GetCurLineupDataScRsp.lineup:type_name -> proto.LineupInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetCurLineupDataScRsp_proto_init() }
func file_GetCurLineupDataScRsp_proto_init() {
	if File_GetCurLineupDataScRsp_proto != nil {
		return
	}
	file_LineupInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetCurLineupDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurLineupDataScRsp); i {
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
			RawDescriptor: file_GetCurLineupDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetCurLineupDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetCurLineupDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetCurLineupDataScRsp_proto_msgTypes,
	}.Build()
	File_GetCurLineupDataScRsp_proto = out.File
	file_GetCurLineupDataScRsp_proto_rawDesc = nil
	file_GetCurLineupDataScRsp_proto_goTypes = nil
	file_GetCurLineupDataScRsp_proto_depIdxs = nil
}
