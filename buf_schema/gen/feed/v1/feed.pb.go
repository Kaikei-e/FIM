// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: feed/v1/feed.proto

package feedv1

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

type GetFeedRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FeedUrl       string                 `protobuf:"bytes,1,opt,name=feed_url,json=feedUrl,proto3" json:"feed_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFeedRequest) Reset() {
	*x = GetFeedRequest{}
	mi := &file_feed_v1_feed_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedRequest) ProtoMessage() {}

func (x *GetFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feed_v1_feed_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedRequest.ProtoReflect.Descriptor instead.
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return file_feed_v1_feed_proto_rawDescGZIP(), []int{0}
}

func (x *GetFeedRequest) GetFeedUrl() string {
	if x != nil {
		return x.FeedUrl
	}
	return ""
}

type GetFeedResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Author        string                 `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	PublishedAt   string                 `protobuf:"bytes,5,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFeedResponse) Reset() {
	*x = GetFeedResponse{}
	mi := &file_feed_v1_feed_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedResponse) ProtoMessage() {}

func (x *GetFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feed_v1_feed_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedResponse.ProtoReflect.Descriptor instead.
func (*GetFeedResponse) Descriptor() ([]byte, []int) {
	return file_feed_v1_feed_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeedResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetFeedResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetFeedResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetFeedResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetFeedResponse) GetPublishedAt() string {
	if x != nil {
		return x.PublishedAt
	}
	return ""
}

var File_feed_v1_feed_proto protoreflect.FileDescriptor

var file_feed_v1_feed_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x2b, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x66, 0x65, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x96, 0x01, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x41, 0x74, 0x32, 0x4d, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x12, 0x17, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4b, 0x61, 0x69, 0x6b, 0x65, 0x69, 0x2d, 0x65, 0x2f, 0x46, 0x49, 0x4d, 0x2f, 0x62, 0x75,
	0x66, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x66, 0x65, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x65, 0x65, 0x64, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feed_v1_feed_proto_rawDescOnce sync.Once
	file_feed_v1_feed_proto_rawDescData = file_feed_v1_feed_proto_rawDesc
)

func file_feed_v1_feed_proto_rawDescGZIP() []byte {
	file_feed_v1_feed_proto_rawDescOnce.Do(func() {
		file_feed_v1_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_feed_v1_feed_proto_rawDescData)
	})
	return file_feed_v1_feed_proto_rawDescData
}

var file_feed_v1_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_feed_v1_feed_proto_goTypes = []any{
	(*GetFeedRequest)(nil),  // 0: feed.v1.GetFeedRequest
	(*GetFeedResponse)(nil), // 1: feed.v1.GetFeedResponse
}
var file_feed_v1_feed_proto_depIdxs = []int32{
	0, // 0: feed.v1.FeedService.GetFeed:input_type -> feed.v1.GetFeedRequest
	1, // 1: feed.v1.FeedService.GetFeed:output_type -> feed.v1.GetFeedResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feed_v1_feed_proto_init() }
func file_feed_v1_feed_proto_init() {
	if File_feed_v1_feed_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feed_v1_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feed_v1_feed_proto_goTypes,
		DependencyIndexes: file_feed_v1_feed_proto_depIdxs,
		MessageInfos:      file_feed_v1_feed_proto_msgTypes,
	}.Build()
	File_feed_v1_feed_proto = out.File
	file_feed_v1_feed_proto_rawDesc = nil
	file_feed_v1_feed_proto_goTypes = nil
	file_feed_v1_feed_proto_depIdxs = nil
}
