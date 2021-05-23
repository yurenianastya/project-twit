// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: proto/reactions/reactions.proto

package reactions

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

type Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TwitId      *ReactionUUID `protobuf:"bytes,1,opt,name=twitId,proto3" json:"twitId,omitempty"`
	LikeCounter int32         `protobuf:"varint,2,opt,name=likeCounter,proto3" json:"likeCounter,omitempty"`
}

func (x *Like) Reset() {
	*x = Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reactions_reactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Like) ProtoMessage() {}

func (x *Like) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reactions_reactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Like.ProtoReflect.Descriptor instead.
func (*Like) Descriptor() ([]byte, []int) {
	return file_proto_reactions_reactions_proto_rawDescGZIP(), []int{0}
}

func (x *Like) GetTwitId() *ReactionUUID {
	if x != nil {
		return x.TwitId
	}
	return nil
}

func (x *Like) GetLikeCounter() int32 {
	if x != nil {
		return x.LikeCounter
	}
	return 0
}

type Retwit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TwitId        *ReactionUUID `protobuf:"bytes,1,opt,name=twitId,proto3" json:"twitId,omitempty"`
	RetwitCounter int32         `protobuf:"varint,2,opt,name=retwitCounter,proto3" json:"retwitCounter,omitempty"`
}

func (x *Retwit) Reset() {
	*x = Retwit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reactions_reactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Retwit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retwit) ProtoMessage() {}

func (x *Retwit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reactions_reactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retwit.ProtoReflect.Descriptor instead.
func (*Retwit) Descriptor() ([]byte, []int) {
	return file_proto_reactions_reactions_proto_rawDescGZIP(), []int{1}
}

func (x *Retwit) GetTwitId() *ReactionUUID {
	if x != nil {
		return x.TwitId
	}
	return nil
}

func (x *Retwit) GetRetwitCounter() int32 {
	if x != nil {
		return x.RetwitCounter
	}
	return 0
}

type ReactionUUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ReactionUUID) Reset() {
	*x = ReactionUUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reactions_reactions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReactionUUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReactionUUID) ProtoMessage() {}

func (x *ReactionUUID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reactions_reactions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReactionUUID.ProtoReflect.Descriptor instead.
func (*ReactionUUID) Descriptor() ([]byte, []int) {
	return file_proto_reactions_reactions_proto_rawDescGZIP(), []int{2}
}

func (x *ReactionUUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_reactions_reactions_proto protoreflect.FileDescriptor

var file_proto_reactions_reactions_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x54, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12,
	0x2a, 0x0a, 0x06, 0x74, 0x77, 0x69, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x06, 0x74, 0x77, 0x69, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6c,
	0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x5a, 0x0a,
	0x06, 0x52, 0x65, 0x74, 0x77, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x74, 0x77, 0x69, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x06, 0x74, 0x77, 0x69,
	0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x77, 0x69, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x77,
	0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x0c, 0x72, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0x97, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x54, 0x77, 0x69, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12,
	0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x55, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x12,
	0x2a, 0x0a, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x77, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x1a,
	0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x75,
	0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x54, 0x77, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x0a, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x32, 0xa5, 0x01, 0x0a, 0x0d, 0x52, 0x65,
	0x74, 0x77, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0e, 0x67,
	0x65, 0x74, 0x54, 0x77, 0x69, 0x74, 0x52, 0x65, 0x74, 0x77, 0x69, 0x74, 0x73, 0x12, 0x12, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49,
	0x44, 0x1a, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x74, 0x77, 0x69, 0x74, 0x12,
	0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x77, 0x69, 0x74, 0x54, 0x77, 0x69, 0x74, 0x12, 0x12, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49,
	0x44, 0x1a, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x74, 0x77, 0x69, 0x74, 0x12,
	0x30, 0x0a, 0x0c, 0x75, 0x6e, 0x72, 0x65, 0x74, 0x77, 0x69, 0x74, 0x54, 0x77, 0x69, 0x74, 0x12,
	0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x55, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x74, 0x77, 0x69,
	0x74, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_reactions_reactions_proto_rawDescOnce sync.Once
	file_proto_reactions_reactions_proto_rawDescData = file_proto_reactions_reactions_proto_rawDesc
)

func file_proto_reactions_reactions_proto_rawDescGZIP() []byte {
	file_proto_reactions_reactions_proto_rawDescOnce.Do(func() {
		file_proto_reactions_reactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reactions_reactions_proto_rawDescData)
	})
	return file_proto_reactions_reactions_proto_rawDescData
}

var file_proto_reactions_reactions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_reactions_reactions_proto_goTypes = []interface{}{
	(*Like)(nil),         // 0: main.Like
	(*Retwit)(nil),       // 1: main.Retwit
	(*ReactionUUID)(nil), // 2: main.reactionUUID
}
var file_proto_reactions_reactions_proto_depIdxs = []int32{
	2, // 0: main.Like.twitId:type_name -> main.reactionUUID
	2, // 1: main.Retwit.twitId:type_name -> main.reactionUUID
	2, // 2: main.LikeService.getTwitLikes:input_type -> main.reactionUUID
	2, // 3: main.LikeService.likeTwit:input_type -> main.reactionUUID
	2, // 4: main.LikeService.unlikeTwit:input_type -> main.reactionUUID
	2, // 5: main.RetwitService.getTwitRetwits:input_type -> main.reactionUUID
	2, // 6: main.RetwitService.retwitTwit:input_type -> main.reactionUUID
	2, // 7: main.RetwitService.unretwitTwit:input_type -> main.reactionUUID
	0, // 8: main.LikeService.getTwitLikes:output_type -> main.Like
	0, // 9: main.LikeService.likeTwit:output_type -> main.Like
	0, // 10: main.LikeService.unlikeTwit:output_type -> main.Like
	1, // 11: main.RetwitService.getTwitRetwits:output_type -> main.Retwit
	1, // 12: main.RetwitService.retwitTwit:output_type -> main.Retwit
	1, // 13: main.RetwitService.unretwitTwit:output_type -> main.Retwit
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_reactions_reactions_proto_init() }
func file_proto_reactions_reactions_proto_init() {
	if File_proto_reactions_reactions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_reactions_reactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Like); i {
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
		file_proto_reactions_reactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Retwit); i {
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
		file_proto_reactions_reactions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReactionUUID); i {
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
			RawDescriptor: file_proto_reactions_reactions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_reactions_reactions_proto_goTypes,
		DependencyIndexes: file_proto_reactions_reactions_proto_depIdxs,
		MessageInfos:      file_proto_reactions_reactions_proto_msgTypes,
	}.Build()
	File_proto_reactions_reactions_proto = out.File
	file_proto_reactions_reactions_proto_rawDesc = nil
	file_proto_reactions_reactions_proto_goTypes = nil
	file_proto_reactions_reactions_proto_depIdxs = nil
}
