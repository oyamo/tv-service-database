// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: proto/database/database.proto

package database

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description       string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Length            int32  `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	StreamUrl         string `protobuf:"bytes,4,opt,name=stream_url,json=streamUrl,proto3" json:"stream_url,omitempty"`
	VideoThumbnailUrl string `protobuf:"bytes,5,opt,name=video_thumbnail_url,json=videoThumbnailUrl,proto3" json:"video_thumbnail_url,omitempty"`
	VideoCoverUrl     string `protobuf:"bytes,6,opt,name=video_cover_url,json=videoCoverUrl,proto3" json:"video_cover_url,omitempty"`
	VideoViews        string `protobuf:"bytes,7,opt,name=video_views,json=videoViews,proto3" json:"video_views,omitempty"`
	VideoLikes        string `protobuf:"bytes,8,opt,name=video_likes,json=videoLikes,proto3" json:"video_likes,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_proto_database_database_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Video) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Video) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Video) GetStreamUrl() string {
	if x != nil {
		return x.StreamUrl
	}
	return ""
}

func (x *Video) GetVideoThumbnailUrl() string {
	if x != nil {
		return x.VideoThumbnailUrl
	}
	return ""
}

func (x *Video) GetVideoCoverUrl() string {
	if x != nil {
		return x.VideoCoverUrl
	}
	return ""
}

func (x *Video) GetVideoViews() string {
	if x != nil {
		return x.VideoViews
	}
	return ""
}

func (x *Video) GetVideoLikes() string {
	if x != nil {
		return x.VideoLikes
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_database_database_proto_rawDescGZIP(), []int{1}
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VideosCount          int32  `protobuf:"varint,3,opt,name=videos_count,json=videosCount,proto3" json:"videos_count,omitempty"`
	CategoryThumbnailUrl string `protobuf:"bytes,4,opt,name=category_thumbnail_url,json=categoryThumbnailUrl,proto3" json:"category_thumbnail_url,omitempty"`
	CategoryCoverUrl     string `protobuf:"bytes,5,opt,name=category_cover_url,json=categoryCoverUrl,proto3" json:"category_cover_url,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_proto_database_database_proto_rawDescGZIP(), []int{2}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetVideosCount() int32 {
	if x != nil {
		return x.VideosCount
	}
	return 0
}

func (x *Category) GetCategoryThumbnailUrl() string {
	if x != nil {
		return x.CategoryThumbnailUrl
	}
	return ""
}

func (x *Category) GetCategoryCoverUrl() string {
	if x != nil {
		return x.CategoryCoverUrl
	}
	return ""
}

type ResponseCategories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     int32       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Categories []*Category `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	Message    string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseCategories) Reset() {
	*x = ResponseCategories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCategories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCategories) ProtoMessage() {}

func (x *ResponseCategories) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCategories.ProtoReflect.Descriptor instead.
func (*ResponseCategories) Descriptor() ([]byte, []int) {
	return file_proto_database_database_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseCategories) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ResponseCategories) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *ResponseCategories) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResponseVideos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Video   []*Video `protobuf:"bytes,2,rep,name=video,proto3" json:"video,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseVideos) Reset() {
	*x = ResponseVideos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseVideos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseVideos) ProtoMessage() {}

func (x *ResponseVideos) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseVideos.ProtoReflect.Descriptor instead.
func (*ResponseVideos) Descriptor() ([]byte, []int) {
	return file_proto_database_database_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseVideos) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ResponseVideos) GetVideo() []*Video {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *ResponseVideos) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_database_database_proto protoreflect.FileDescriptor

var file_proto_database_database_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0x8a, 0x02, 0x0a, 0x05, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x13,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x56, 0x69, 0x65, 0x77, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xb5, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x2c,
	0x0a, 0x12, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x7a, 0x0a, 0x12,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x97, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x18, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x79, 0x61, 0x6d,
	0x6f, 0x68, 0x2d, 0x62, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x74, 0x76, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_database_database_proto_rawDescOnce sync.Once
	file_proto_database_database_proto_rawDescData = file_proto_database_database_proto_rawDesc
)

func file_proto_database_database_proto_rawDescGZIP() []byte {
	file_proto_database_database_proto_rawDescOnce.Do(func() {
		file_proto_database_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_database_database_proto_rawDescData)
	})
	return file_proto_database_database_proto_rawDescData
}

var file_proto_database_database_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_database_database_proto_goTypes = []interface{}{
	(*Video)(nil),              // 0: database.Video
	(*GetRequest)(nil),         // 1: database.GetRequest
	(*Category)(nil),           // 2: database.Category
	(*ResponseCategories)(nil), // 3: database.ResponseCategories
	(*ResponseVideos)(nil),     // 4: database.ResponseVideos
}
var file_proto_database_database_proto_depIdxs = []int32{
	2, // 0: database.ResponseCategories.categories:type_name -> database.Category
	0, // 1: database.ResponseVideos.video:type_name -> database.Video
	2, // 2: database.DataBaseService.GetAllVideos:input_type -> database.Category
	1, // 3: database.DataBaseService.GetAllCategories:input_type -> database.GetRequest
	4, // 4: database.DataBaseService.GetAllVideos:output_type -> database.ResponseVideos
	3, // 5: database.DataBaseService.GetAllCategories:output_type -> database.ResponseCategories
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_database_database_proto_init() }
func file_proto_database_database_proto_init() {
	if File_proto_database_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_database_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_proto_database_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_proto_database_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_proto_database_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCategories); i {
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
		file_proto_database_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseVideos); i {
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
			RawDescriptor: file_proto_database_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_database_database_proto_goTypes,
		DependencyIndexes: file_proto_database_database_proto_depIdxs,
		MessageInfos:      file_proto_database_database_proto_msgTypes,
	}.Build()
	File_proto_database_database_proto = out.File
	file_proto_database_database_proto_rawDesc = nil
	file_proto_database_database_proto_goTypes = nil
	file_proto_database_database_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataBaseServiceClient is the client API for DataBaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataBaseServiceClient interface {
	GetAllVideos(ctx context.Context, in *Category, opts ...grpc.CallOption) (*ResponseVideos, error)
	GetAllCategories(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ResponseCategories, error)
}

type dataBaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataBaseServiceClient(cc grpc.ClientConnInterface) DataBaseServiceClient {
	return &dataBaseServiceClient{cc}
}

func (c *dataBaseServiceClient) GetAllVideos(ctx context.Context, in *Category, opts ...grpc.CallOption) (*ResponseVideos, error) {
	out := new(ResponseVideos)
	err := c.cc.Invoke(ctx, "/database.DataBaseService/GetAllVideos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataBaseServiceClient) GetAllCategories(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ResponseCategories, error) {
	out := new(ResponseCategories)
	err := c.cc.Invoke(ctx, "/database.DataBaseService/GetAllCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataBaseServiceServer is the server API for DataBaseService service.
type DataBaseServiceServer interface {
	GetAllVideos(context.Context, *Category) (*ResponseVideos, error)
	GetAllCategories(context.Context, *GetRequest) (*ResponseCategories, error)
}

// UnimplementedDataBaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataBaseServiceServer struct {
}

func (*UnimplementedDataBaseServiceServer) GetAllVideos(context.Context, *Category) (*ResponseVideos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllVideos not implemented")
}
func (*UnimplementedDataBaseServiceServer) GetAllCategories(context.Context, *GetRequest) (*ResponseCategories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}

func RegisterDataBaseServiceServer(s *grpc.Server, srv DataBaseServiceServer) {
	s.RegisterService(&_DataBaseService_serviceDesc, srv)
}

func _DataBaseService_GetAllVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBaseServiceServer).GetAllVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/database.DataBaseService/GetAllVideos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBaseServiceServer).GetAllVideos(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataBaseService_GetAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataBaseServiceServer).GetAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/database.DataBaseService/GetAllCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataBaseServiceServer).GetAllCategories(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataBaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "database.DataBaseService",
	HandlerType: (*DataBaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllVideos",
			Handler:    _DataBaseService_GetAllVideos_Handler,
		},
		{
			MethodName: "GetAllCategories",
			Handler:    _DataBaseService_GetAllCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/database/database.proto",
}