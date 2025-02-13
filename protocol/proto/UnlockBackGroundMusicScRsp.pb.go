// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: UnlockBackGroundMusicScRsp.proto

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

type UnlockBackGroundMusicScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MusicList   []*UnlockedMusic `protobuf:"bytes,6,rep,name=music_list,json=musicList,proto3" json:"music_list,omitempty"`
	UnlockedIds []uint32         `protobuf:"varint,3,rep,packed,name=unlocked_ids,json=unlockedIds,proto3" json:"unlocked_ids,omitempty"`
	Retcode     uint32           `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *UnlockBackGroundMusicScRsp) Reset() {
	*x = UnlockBackGroundMusicScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnlockBackGroundMusicScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockBackGroundMusicScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockBackGroundMusicScRsp) ProtoMessage() {}

func (x *UnlockBackGroundMusicScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_UnlockBackGroundMusicScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockBackGroundMusicScRsp.ProtoReflect.Descriptor instead.
func (*UnlockBackGroundMusicScRsp) Descriptor() ([]byte, []int) {
	return file_UnlockBackGroundMusicScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *UnlockBackGroundMusicScRsp) GetMusicList() []*UnlockedMusic {
	if x != nil {
		return x.MusicList
	}
	return nil
}

func (x *UnlockBackGroundMusicScRsp) GetUnlockedIds() []uint32 {
	if x != nil {
		return x.UnlockedIds
	}
	return nil
}

func (x *UnlockBackGroundMusicScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_UnlockBackGroundMusicScRsp_proto protoreflect.FileDescriptor

var file_UnlockBackGroundMusicScRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x47, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e,
	0x01, 0x0a, 0x1a, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x47, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x33, 0x0a,
	0x0a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x52, 0x09, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_UnlockBackGroundMusicScRsp_proto_rawDescOnce sync.Once
	file_UnlockBackGroundMusicScRsp_proto_rawDescData = file_UnlockBackGroundMusicScRsp_proto_rawDesc
)

func file_UnlockBackGroundMusicScRsp_proto_rawDescGZIP() []byte {
	file_UnlockBackGroundMusicScRsp_proto_rawDescOnce.Do(func() {
		file_UnlockBackGroundMusicScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_UnlockBackGroundMusicScRsp_proto_rawDescData)
	})
	return file_UnlockBackGroundMusicScRsp_proto_rawDescData
}

var file_UnlockBackGroundMusicScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UnlockBackGroundMusicScRsp_proto_goTypes = []interface{}{
	(*UnlockBackGroundMusicScRsp)(nil), // 0: proto.UnlockBackGroundMusicScRsp
	(*UnlockedMusic)(nil),              // 1: proto.UnlockedMusic
}
var file_UnlockBackGroundMusicScRsp_proto_depIdxs = []int32{
	1, // 0: proto.UnlockBackGroundMusicScRsp.music_list:type_name -> proto.UnlockedMusic
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UnlockBackGroundMusicScRsp_proto_init() }
func file_UnlockBackGroundMusicScRsp_proto_init() {
	if File_UnlockBackGroundMusicScRsp_proto != nil {
		return
	}
	file_UnlockedMusic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UnlockBackGroundMusicScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockBackGroundMusicScRsp); i {
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
			RawDescriptor: file_UnlockBackGroundMusicScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UnlockBackGroundMusicScRsp_proto_goTypes,
		DependencyIndexes: file_UnlockBackGroundMusicScRsp_proto_depIdxs,
		MessageInfos:      file_UnlockBackGroundMusicScRsp_proto_msgTypes,
	}.Build()
	File_UnlockBackGroundMusicScRsp_proto = out.File
	file_UnlockBackGroundMusicScRsp_proto_rawDesc = nil
	file_UnlockBackGroundMusicScRsp_proto_goTypes = nil
	file_UnlockBackGroundMusicScRsp_proto_depIdxs = nil
}
