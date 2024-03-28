// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: user.proto

package handlers

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

type PostFullLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*PostFullLinkResponse_ShortLink
	//	*PostFullLinkResponse_Message
	Result isPostFullLinkResponse_Result `protobuf_oneof:"result"`
}

func (x *PostFullLinkResponse) Reset() {
	*x = PostFullLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFullLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFullLinkResponse) ProtoMessage() {}

func (x *PostFullLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFullLinkResponse.ProtoReflect.Descriptor instead.
func (*PostFullLinkResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (m *PostFullLinkResponse) GetResult() isPostFullLinkResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *PostFullLinkResponse) GetShortLink() *ShortLinkResponse {
	if x, ok := x.GetResult().(*PostFullLinkResponse_ShortLink); ok {
		return x.ShortLink
	}
	return nil
}

func (x *PostFullLinkResponse) GetMessage() *MessageResponse {
	if x, ok := x.GetResult().(*PostFullLinkResponse_Message); ok {
		return x.Message
	}
	return nil
}

type isPostFullLinkResponse_Result interface {
	isPostFullLinkResponse_Result()
}

type PostFullLinkResponse_ShortLink struct {
	ShortLink *ShortLinkResponse `protobuf:"bytes,1,opt,name=shortLink,proto3,oneof"`
}

type PostFullLinkResponse_Message struct {
	Message *MessageResponse `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

func (*PostFullLinkResponse_ShortLink) isPostFullLinkResponse_Result() {}

func (*PostFullLinkResponse_Message) isPostFullLinkResponse_Result() {}

type LongLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongLink string `protobuf:"bytes,1,opt,name=longLink,proto3" json:"longLink,omitempty"`
}

func (x *LongLinkRequest) Reset() {
	*x = LongLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongLinkRequest) ProtoMessage() {}

func (x *LongLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongLinkRequest.ProtoReflect.Descriptor instead.
func (*LongLinkRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *LongLinkRequest) GetLongLink() string {
	if x != nil {
		return x.LongLink
	}
	return ""
}

type FullLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullLink string `protobuf:"bytes,1,opt,name=fullLink,proto3" json:"fullLink,omitempty"`
}

func (x *FullLinkResponse) Reset() {
	*x = FullLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullLinkResponse) ProtoMessage() {}

func (x *FullLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullLinkResponse.ProtoReflect.Descriptor instead.
func (*FullLinkResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *FullLinkResponse) GetFullLink() string {
	if x != nil {
		return x.FullLink
	}
	return ""
}

type GetFullLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*GetFullLinkResponse_FullLink
	//	*GetFullLinkResponse_Message
	Result isGetFullLinkResponse_Result `protobuf_oneof:"result"`
}

func (x *GetFullLinkResponse) Reset() {
	*x = GetFullLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullLinkResponse) ProtoMessage() {}

func (x *GetFullLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullLinkResponse.ProtoReflect.Descriptor instead.
func (*GetFullLinkResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (m *GetFullLinkResponse) GetResult() isGetFullLinkResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *GetFullLinkResponse) GetFullLink() *FullLinkResponse {
	if x, ok := x.GetResult().(*GetFullLinkResponse_FullLink); ok {
		return x.FullLink
	}
	return nil
}

func (x *GetFullLinkResponse) GetMessage() *MessageResponse {
	if x, ok := x.GetResult().(*GetFullLinkResponse_Message); ok {
		return x.Message
	}
	return nil
}

type isGetFullLinkResponse_Result interface {
	isGetFullLinkResponse_Result()
}

type GetFullLinkResponse_FullLink struct {
	FullLink *FullLinkResponse `protobuf:"bytes,1,opt,name=fullLink,proto3,oneof"`
}

type GetFullLinkResponse_Message struct {
	Message *MessageResponse `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

func (*GetFullLinkResponse_FullLink) isGetFullLinkResponse_Result() {}

func (*GetFullLinkResponse_Message) isGetFullLinkResponse_Result() {}

type ShortLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortLink string `protobuf:"bytes,1,opt,name=shortLink,proto3" json:"shortLink,omitempty"`
}

func (x *ShortLinkRequest) Reset() {
	*x = ShortLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortLinkRequest) ProtoMessage() {}

func (x *ShortLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortLinkRequest.ProtoReflect.Descriptor instead.
func (*ShortLinkRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *ShortLinkRequest) GetShortLink() string {
	if x != nil {
		return x.ShortLink
	}
	return ""
}

type ShortLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortLink string `protobuf:"bytes,1,opt,name=shortLink,proto3" json:"shortLink,omitempty"`
}

func (x *ShortLinkResponse) Reset() {
	*x = ShortLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortLinkResponse) ProtoMessage() {}

func (x *ShortLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortLinkResponse.ProtoReflect.Descriptor instead.
func (*ShortLinkResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *ShortLinkResponse) GetShortLink() string {
	if x != nil {
		return x.ShortLink
	}
	return ""
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *MessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x75, 0x6c, 0x6c,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x2d, 0x0a, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x22, 0x2e, 0x0a, 0x10, 0x66, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4c,
	0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4c,
	0x69, 0x6e, 0x6b, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x66, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x30, 0x0a, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x22, 0x31, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x2b, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x97, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x16, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x6c, 0x6f, 0x6e, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x17, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11,
	0x6f, 0x7a, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_user_proto_goTypes = []interface{}{
	(*PostFullLinkResponse)(nil), // 0: links.PostFullLinkResponse
	(*LongLinkRequest)(nil),      // 1: links.longLinkRequest
	(*FullLinkResponse)(nil),     // 2: links.fullLinkResponse
	(*GetFullLinkResponse)(nil),  // 3: links.GetFullLinkResponse
	(*ShortLinkRequest)(nil),     // 4: links.shortLinkRequest
	(*ShortLinkResponse)(nil),    // 5: links.shortLinkResponse
	(*MessageResponse)(nil),      // 6: links.messageResponse
}
var file_user_proto_depIdxs = []int32{
	5, // 0: links.PostFullLinkResponse.shortLink:type_name -> links.shortLinkResponse
	6, // 1: links.PostFullLinkResponse.message:type_name -> links.messageResponse
	2, // 2: links.GetFullLinkResponse.fullLink:type_name -> links.fullLinkResponse
	6, // 3: links.GetFullLinkResponse.message:type_name -> links.messageResponse
	1, // 4: links.LinksService.PostFullLink:input_type -> links.longLinkRequest
	4, // 5: links.LinksService.GetFullLink:input_type -> links.shortLinkRequest
	0, // 6: links.LinksService.PostFullLink:output_type -> links.PostFullLinkResponse
	3, // 7: links.LinksService.GetFullLink:output_type -> links.GetFullLinkResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFullLinkResponse); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongLinkRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullLinkResponse); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullLinkResponse); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortLinkRequest); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortLinkResponse); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
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
	file_user_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PostFullLinkResponse_ShortLink)(nil),
		(*PostFullLinkResponse_Message)(nil),
	}
	file_user_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GetFullLinkResponse_FullLink)(nil),
		(*GetFullLinkResponse_Message)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
