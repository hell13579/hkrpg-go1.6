// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: PlayerBasicInfo.proto

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

type PlayerBasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname   string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Level      uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Exp        uint32 `protobuf:"varint,3,opt,name=exp,proto3" json:"exp,omitempty"`
	Stamina    uint32 `protobuf:"varint,4,opt,name=stamina,proto3" json:"stamina,omitempty"`
	Mcoin      uint32 `protobuf:"varint,5,opt,name=mcoin,proto3" json:"mcoin,omitempty"`
	Hcoin      uint32 `protobuf:"varint,6,opt,name=hcoin,proto3" json:"hcoin,omitempty"`
	Scoin      uint32 `protobuf:"varint,7,opt,name=scoin,proto3" json:"scoin,omitempty"`
	WorldLevel uint32 `protobuf:"varint,8,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
}

func (x *PlayerBasicInfo) Reset() {
	*x = PlayerBasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerBasicInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerBasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerBasicInfo) ProtoMessage() {}

func (x *PlayerBasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerBasicInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerBasicInfo.ProtoReflect.Descriptor instead.
func (*PlayerBasicInfo) Descriptor() ([]byte, []int) {
	return file_PlayerBasicInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerBasicInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PlayerBasicInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *PlayerBasicInfo) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *PlayerBasicInfo) GetStamina() uint32 {
	if x != nil {
		return x.Stamina
	}
	return 0
}

func (x *PlayerBasicInfo) GetMcoin() uint32 {
	if x != nil {
		return x.Mcoin
	}
	return 0
}

func (x *PlayerBasicInfo) GetHcoin() uint32 {
	if x != nil {
		return x.Hcoin
	}
	return 0
}

func (x *PlayerBasicInfo) GetScoin() uint32 {
	if x != nil {
		return x.Scoin
	}
	return 0
}

func (x *PlayerBasicInfo) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

var File_PlayerBasicInfo_proto protoreflect.FileDescriptor

var file_PlayerBasicInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2,
	0x01, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x6d, 0x69, 0x6e,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6d, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x63, 0x6f, 0x69, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x68, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerBasicInfo_proto_rawDescOnce sync.Once
	file_PlayerBasicInfo_proto_rawDescData = file_PlayerBasicInfo_proto_rawDesc
)

func file_PlayerBasicInfo_proto_rawDescGZIP() []byte {
	file_PlayerBasicInfo_proto_rawDescOnce.Do(func() {
		file_PlayerBasicInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerBasicInfo_proto_rawDescData)
	})
	return file_PlayerBasicInfo_proto_rawDescData
}

var file_PlayerBasicInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerBasicInfo_proto_goTypes = []interface{}{
	(*PlayerBasicInfo)(nil), // 0: proto.PlayerBasicInfo
}
var file_PlayerBasicInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerBasicInfo_proto_init() }
func file_PlayerBasicInfo_proto_init() {
	if File_PlayerBasicInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerBasicInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerBasicInfo); i {
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
			RawDescriptor: file_PlayerBasicInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerBasicInfo_proto_goTypes,
		DependencyIndexes: file_PlayerBasicInfo_proto_depIdxs,
		MessageInfos:      file_PlayerBasicInfo_proto_msgTypes,
	}.Build()
	File_PlayerBasicInfo_proto = out.File
	file_PlayerBasicInfo_proto_rawDesc = nil
	file_PlayerBasicInfo_proto_goTypes = nil
	file_PlayerBasicInfo_proto_depIdxs = nil
}
