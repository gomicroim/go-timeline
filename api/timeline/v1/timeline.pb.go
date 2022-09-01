// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/timeline/v1/timeline.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 发送消息请求(c2c)
type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To      string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendMessageRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence int64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *SendMessageReply) Reset() {
	*x = SendMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReply) ProtoMessage() {}

func (x *SendMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReply.ProtoReflect.Descriptor instead.
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageReply) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

// 发送消息请求(c2g)
type SendGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName    string   `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	GroupMembers []string `protobuf:"bytes,2,rep,name=group_members,json=groupMembers,proto3" json:"group_members,omitempty"`
	Message      string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendGroupRequest) Reset() {
	*x = SendGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendGroupRequest) ProtoMessage() {}

func (x *SendGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendGroupRequest.ProtoReflect.Descriptor instead.
func (*SendGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{2}
}

func (x *SendGroupRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *SendGroupRequest) GetGroupMembers() []string {
	if x != nil {
		return x.GroupMembers
	}
	return nil
}

func (x *SendGroupRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailedMembers []string `protobuf:"bytes,1,rep,name=failed_members,json=failedMembers,proto3" json:"failed_members,omitempty"`
}

func (x *SendGroupReply) Reset() {
	*x = SendGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendGroupReply) ProtoMessage() {}

func (x *SendGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendGroupReply.ProtoReflect.Descriptor instead.
func (*SendGroupReply) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{3}
}

func (x *SendGroupReply) GetFailedMembers() []string {
	if x != nil {
		return x.FailedMembers
	}
	return nil
}

// 同步消息请求
type SyncMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member   string `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
	LastRead int64  `protobuf:"varint,2,opt,name=last_read,json=lastRead,proto3" json:"last_read,omitempty"`
	Count    int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SyncMessageRequest) Reset() {
	*x = SyncMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMessageRequest) ProtoMessage() {}

func (x *SyncMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMessageRequest.ProtoReflect.Descriptor instead.
func (*SyncMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{4}
}

func (x *SyncMessageRequest) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *SyncMessageRequest) GetLastRead() int64 {
	if x != nil {
		return x.LastRead
	}
	return 0
}

func (x *SyncMessageRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TimelineEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence int64  `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TimelineEntry) Reset() {
	*x = TimelineEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimelineEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimelineEntry) ProtoMessage() {}

func (x *TimelineEntry) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimelineEntry.ProtoReflect.Descriptor instead.
func (*TimelineEntry) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{5}
}

func (x *TimelineEntry) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *TimelineEntry) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SyncMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 最新的消息序号
	LatestSeq int64 `protobuf:"varint,1,opt,name=latest_seq,json=latestSeq,proto3" json:"latest_seq,omitempty"`
	// entry_set 中最后的序号值
	EntrySetLastSeq int64            `protobuf:"varint,2,opt,name=entry_set_last_seq,json=entrySetLastSeq,proto3" json:"entry_set_last_seq,omitempty"`
	EntrySet        []*TimelineEntry `protobuf:"bytes,3,rep,name=entry_set,json=entrySet,proto3" json:"entry_set,omitempty"`
}

func (x *SyncMessageReply) Reset() {
	*x = SyncMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMessageReply) ProtoMessage() {}

func (x *SyncMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMessageReply.ProtoReflect.Descriptor instead.
func (*SyncMessageReply) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{6}
}

func (x *SyncMessageReply) GetLatestSeq() int64 {
	if x != nil {
		return x.LatestSeq
	}
	return 0
}

func (x *SyncMessageReply) GetEntrySetLastSeq() int64 {
	if x != nil {
		return x.EntrySetLastSeq
	}
	return 0
}

func (x *SyncMessageReply) GetEntrySet() []*TimelineEntry {
	if x != nil {
		return x.EntrySet
	}
	return nil
}

// 查询单聊历史消息请求
type GetSingleHistoryMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	LastRead int64  `protobuf:"varint,3,opt,name=last_read,json=lastRead,proto3" json:"last_read,omitempty"`
	Count    int32  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetSingleHistoryMessageRequest) Reset() {
	*x = GetSingleHistoryMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleHistoryMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleHistoryMessageRequest) ProtoMessage() {}

func (x *GetSingleHistoryMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleHistoryMessageRequest.ProtoReflect.Descriptor instead.
func (*GetSingleHistoryMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{7}
}

func (x *GetSingleHistoryMessageRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GetSingleHistoryMessageRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *GetSingleHistoryMessageRequest) GetLastRead() int64 {
	if x != nil {
		return x.LastRead
	}
	return 0
}

func (x *GetSingleHistoryMessageRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetSingleHistoryMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntrySet []*TimelineEntry `protobuf:"bytes,1,rep,name=entry_set,json=entrySet,proto3" json:"entry_set,omitempty"`
}

func (x *GetSingleHistoryMessageReply) Reset() {
	*x = GetSingleHistoryMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleHistoryMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleHistoryMessageReply) ProtoMessage() {}

func (x *GetSingleHistoryMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleHistoryMessageReply.ProtoReflect.Descriptor instead.
func (*GetSingleHistoryMessageReply) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{8}
}

func (x *GetSingleHistoryMessageReply) GetEntrySet() []*TimelineEntry {
	if x != nil {
		return x.EntrySet
	}
	return nil
}

// 查询群历史消息请求
type GetGroupHistoryMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group    string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	LastRead int64  `protobuf:"varint,3,opt,name=last_read,json=lastRead,proto3" json:"last_read,omitempty"`
	Count    int32  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetGroupHistoryMessageRequest) Reset() {
	*x = GetGroupHistoryMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupHistoryMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupHistoryMessageRequest) ProtoMessage() {}

func (x *GetGroupHistoryMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupHistoryMessageRequest.ProtoReflect.Descriptor instead.
func (*GetGroupHistoryMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{9}
}

func (x *GetGroupHistoryMessageRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GetGroupHistoryMessageRequest) GetLastRead() int64 {
	if x != nil {
		return x.LastRead
	}
	return 0
}

func (x *GetGroupHistoryMessageRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetGroupHistoryMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntrySet []*TimelineEntry `protobuf:"bytes,1,rep,name=entry_set,json=entrySet,proto3" json:"entry_set,omitempty"`
}

func (x *GetGroupHistoryMessageReply) Reset() {
	*x = GetGroupHistoryMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_timeline_v1_timeline_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupHistoryMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupHistoryMessageReply) ProtoMessage() {}

func (x *GetGroupHistoryMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_timeline_v1_timeline_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupHistoryMessageReply.ProtoReflect.Descriptor instead.
func (*GetGroupHistoryMessageReply) Descriptor() ([]byte, []int) {
	return file_api_timeline_v1_timeline_proto_rawDescGZIP(), []int{10}
}

func (x *GetGroupHistoryMessageReply) GetEntrySet() []*TimelineEntry {
	if x != nil {
		return x.EntrySet
	}
	return nil
}

var File_api_timeline_v1_timeline_proto protoreflect.FileDescriptor

var file_api_timeline_v1_timeline_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x12, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x2e, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22,
	0x70, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x37, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x5f, 0x0a, 0x12, 0x53, 0x79,
	0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x0d, 0x54,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x53, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x12, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x74, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74,
	0x53, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x22, 0x77, 0x0a, 0x1e,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x22, 0x68,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74,
	0x32, 0xfc, 0x04, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x61, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0x67, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x22, 0x13, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x6e,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x73,
	0x79, 0x6e, 0x63, 0x12, 0x9f, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2b, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12,
	0x24, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x7b, 0x66, 0x72, 0x6f, 0x6d, 0x7d,
	0x2f, 0x7b, 0x74, 0x6f, 0x7d, 0x12, 0x97, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2a, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0x42,
	0x25, 0x5a, 0x23, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_timeline_v1_timeline_proto_rawDescOnce sync.Once
	file_api_timeline_v1_timeline_proto_rawDescData = file_api_timeline_v1_timeline_proto_rawDesc
)

func file_api_timeline_v1_timeline_proto_rawDescGZIP() []byte {
	file_api_timeline_v1_timeline_proto_rawDescOnce.Do(func() {
		file_api_timeline_v1_timeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_timeline_v1_timeline_proto_rawDescData)
	})
	return file_api_timeline_v1_timeline_proto_rawDescData
}

var file_api_timeline_v1_timeline_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_timeline_v1_timeline_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil),             // 0: timeline.v1.SendMessageRequest
	(*SendMessageReply)(nil),               // 1: timeline.v1.SendMessageReply
	(*SendGroupRequest)(nil),               // 2: timeline.v1.SendGroupRequest
	(*SendGroupReply)(nil),                 // 3: timeline.v1.SendGroupReply
	(*SyncMessageRequest)(nil),             // 4: timeline.v1.SyncMessageRequest
	(*TimelineEntry)(nil),                  // 5: timeline.v1.TimelineEntry
	(*SyncMessageReply)(nil),               // 6: timeline.v1.SyncMessageReply
	(*GetSingleHistoryMessageRequest)(nil), // 7: timeline.v1.GetSingleHistoryMessageRequest
	(*GetSingleHistoryMessageReply)(nil),   // 8: timeline.v1.GetSingleHistoryMessageReply
	(*GetGroupHistoryMessageRequest)(nil),  // 9: timeline.v1.GetGroupHistoryMessageRequest
	(*GetGroupHistoryMessageReply)(nil),    // 10: timeline.v1.GetGroupHistoryMessageReply
}
var file_api_timeline_v1_timeline_proto_depIdxs = []int32{
	5,  // 0: timeline.v1.SyncMessageReply.entry_set:type_name -> timeline.v1.TimelineEntry
	5,  // 1: timeline.v1.GetSingleHistoryMessageReply.entry_set:type_name -> timeline.v1.TimelineEntry
	5,  // 2: timeline.v1.GetGroupHistoryMessageReply.entry_set:type_name -> timeline.v1.TimelineEntry
	0,  // 3: timeline.v1.Timeline.Send:input_type -> timeline.v1.SendMessageRequest
	2,  // 4: timeline.v1.Timeline.SendGroup:input_type -> timeline.v1.SendGroupRequest
	4,  // 5: timeline.v1.Timeline.GetSyncMessage:input_type -> timeline.v1.SyncMessageRequest
	7,  // 6: timeline.v1.Timeline.GetSingleHistoryMessage:input_type -> timeline.v1.GetSingleHistoryMessageRequest
	9,  // 7: timeline.v1.Timeline.GetGroupHistoryMessage:input_type -> timeline.v1.GetGroupHistoryMessageRequest
	1,  // 8: timeline.v1.Timeline.Send:output_type -> timeline.v1.SendMessageReply
	3,  // 9: timeline.v1.Timeline.SendGroup:output_type -> timeline.v1.SendGroupReply
	6,  // 10: timeline.v1.Timeline.GetSyncMessage:output_type -> timeline.v1.SyncMessageReply
	8,  // 11: timeline.v1.Timeline.GetSingleHistoryMessage:output_type -> timeline.v1.GetSingleHistoryMessageReply
	10, // 12: timeline.v1.Timeline.GetGroupHistoryMessage:output_type -> timeline.v1.GetGroupHistoryMessageReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_timeline_v1_timeline_proto_init() }
func file_api_timeline_v1_timeline_proto_init() {
	if File_api_timeline_v1_timeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_timeline_v1_timeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReply); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendGroupRequest); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendGroupReply); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMessageRequest); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimelineEntry); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMessageReply); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSingleHistoryMessageRequest); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSingleHistoryMessageReply); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupHistoryMessageRequest); i {
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
		file_api_timeline_v1_timeline_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupHistoryMessageReply); i {
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
			RawDescriptor: file_api_timeline_v1_timeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_timeline_v1_timeline_proto_goTypes,
		DependencyIndexes: file_api_timeline_v1_timeline_proto_depIdxs,
		MessageInfos:      file_api_timeline_v1_timeline_proto_msgTypes,
	}.Build()
	File_api_timeline_v1_timeline_proto = out.File
	file_api_timeline_v1_timeline_proto_rawDesc = nil
	file_api_timeline_v1_timeline_proto_goTypes = nil
	file_api_timeline_v1_timeline_proto_depIdxs = nil
}
