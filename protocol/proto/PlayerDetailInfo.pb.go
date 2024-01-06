// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: PlayerDetailInfo.proto

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

type PlayerDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldLevel        uint32       `protobuf:"varint,10,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	Signature         string       `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	PlatformType      PlatformType `protobuf:"varint,4,opt,name=platform_type,json=platformType,proto3,enum=proto.PlatformType" json:"platform_type,omitempty"`
	Level             uint32       `protobuf:"varint,9,opt,name=level,proto3" json:"level,omitempty"`
	HeadIcon          uint32       `protobuf:"varint,7,opt,name=head_icon,json=headIcon,proto3" json:"head_icon,omitempty"`
	Uid               uint32       `protobuf:"varint,13,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname          string       `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
	DisplayAvatarInfo string       `protobuf:"bytes,14,opt,name=display_avatar_info,json=displayAvatarInfo,proto3" json:"display_avatar_info,omitempty"`
	RecordInfo        string       `protobuf:"bytes,15,opt,name=record_info,json=recordInfo,proto3" json:"record_info,omitempty"` //DisplayAvatarDetailInfo display_avatar_info = 14;
}

func (x *PlayerDetailInfo) Reset() {
	*x = PlayerDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerDetailInfo) ProtoMessage() {}

func (x *PlayerDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerDetailInfo.ProtoReflect.Descriptor instead.
func (*PlayerDetailInfo) Descriptor() ([]byte, []int) {
	return file_PlayerDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerDetailInfo) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *PlayerDetailInfo) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PlayerDetailInfo) GetPlatformType() PlatformType {
	if x != nil {
		return x.PlatformType
	}
	return PlatformType_EDITOR
}

func (x *PlayerDetailInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *PlayerDetailInfo) GetHeadIcon() uint32 {
	if x != nil {
		return x.HeadIcon
	}
	return 0
}

func (x *PlayerDetailInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PlayerDetailInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PlayerDetailInfo) GetDisplayAvatarInfo() string {
	if x != nil {
		return x.DisplayAvatarInfo
	}
	return ""
}

func (x *PlayerDetailInfo) GetRecordInfo() string {
	if x != nil {
		return x.RecordInfo
	}
	return ""
}

var File_PlayerDetailInfo_proto protoreflect.FileDescriptor

var file_PlayerDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x02, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x64, 0x5f,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64,
	0x49, 0x63, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerDetailInfo_proto_rawDescOnce sync.Once
	file_PlayerDetailInfo_proto_rawDescData = file_PlayerDetailInfo_proto_rawDesc
)

func file_PlayerDetailInfo_proto_rawDescGZIP() []byte {
	file_PlayerDetailInfo_proto_rawDescOnce.Do(func() {
		file_PlayerDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerDetailInfo_proto_rawDescData)
	})
	return file_PlayerDetailInfo_proto_rawDescData
}

var file_PlayerDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerDetailInfo_proto_goTypes = []interface{}{
	(*PlayerDetailInfo)(nil), // 0: proto.PlayerDetailInfo
	(PlatformType)(0),        // 1: proto.PlatformType
}
var file_PlayerDetailInfo_proto_depIdxs = []int32{
	1, // 0: proto.PlayerDetailInfo.platform_type:type_name -> proto.PlatformType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerDetailInfo_proto_init() }
func file_PlayerDetailInfo_proto_init() {
	if File_PlayerDetailInfo_proto != nil {
		return
	}
	file_PlatformType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerDetailInfo); i {
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
			RawDescriptor: file_PlayerDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerDetailInfo_proto_goTypes,
		DependencyIndexes: file_PlayerDetailInfo_proto_depIdxs,
		MessageInfos:      file_PlayerDetailInfo_proto_msgTypes,
	}.Build()
	File_PlayerDetailInfo_proto = out.File
	file_PlayerDetailInfo_proto_rawDesc = nil
	file_PlayerDetailInfo_proto_goTypes = nil
	file_PlayerDetailInfo_proto_depIdxs = nil
}
