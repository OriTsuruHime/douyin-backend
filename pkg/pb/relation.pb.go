// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: relation.proto

package pb

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

type RelationActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ToUserId   uint64 `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`
	ActionType int32  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
}

func (x *RelationActionRequest) Reset() {
	*x = RelationActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionRequest) ProtoMessage() {}

func (x *RelationActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionRequest.ProtoReflect.Descriptor instead.
func (*RelationActionRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{0}
}

func (x *RelationActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RelationActionRequest) GetToUserId() uint64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *RelationActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type RelationActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *RelationActionResponse) Reset() {
	*x = RelationActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionResponse) ProtoMessage() {}

func (x *RelationActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionResponse.ProtoReflect.Descriptor instead.
func (*RelationActionResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{1}
}

func (x *RelationActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type GetFollowListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetFollowListRequest) Reset() {
	*x = GetFollowListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowListRequest) ProtoMessage() {}

func (x *GetFollowListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowListRequest.ProtoReflect.Descriptor instead.
func (*GetFollowListRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{2}
}

func (x *GetFollowListRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetFollowListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetFollowListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	UserList   []*User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *GetFollowListResponse) Reset() {
	*x = GetFollowListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowListResponse) ProtoMessage() {}

func (x *GetFollowListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowListResponse.ProtoReflect.Descriptor instead.
func (*GetFollowListResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{3}
}

func (x *GetFollowListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetFollowListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *GetFollowListResponse) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type GetFollowerListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetFollowerListRequest) Reset() {
	*x = GetFollowerListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowerListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowerListRequest) ProtoMessage() {}

func (x *GetFollowerListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowerListRequest.ProtoReflect.Descriptor instead.
func (*GetFollowerListRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{4}
}

func (x *GetFollowerListRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetFollowerListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetFollowerListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	UserList   []*User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *GetFollowerListResponse) Reset() {
	*x = GetFollowerListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowerListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowerListResponse) ProtoMessage() {}

func (x *GetFollowerListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowerListResponse.ProtoReflect.Descriptor instead.
func (*GetFollowerListResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{5}
}

func (x *GetFollowerListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetFollowerListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *GetFollowerListResponse) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type GetFriendListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetFriendListRequest) Reset() {
	*x = GetFriendListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendListRequest) ProtoMessage() {}

func (x *GetFriendListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendListRequest.ProtoReflect.Descriptor instead.
func (*GetFriendListRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{6}
}

func (x *GetFriendListRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetFriendListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetFriendListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	UserList   []*FriendUser `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *GetFriendListResponse) Reset() {
	*x = GetFriendListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendListResponse) ProtoMessage() {}

func (x *GetFriendListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendListResponse.ProtoReflect.Descriptor instead.
func (*GetFriendListResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{7}
}

func (x *GetFriendListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetFriendListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *GetFriendListResponse) GetUserList() []*FriendUser {
	if x != nil {
		return x.UserList
	}
	return nil
}

type FriendUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MsgType int32  `protobuf:"varint,3,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
}

func (x *FriendUser) Reset() {
	*x = FriendUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendUser) ProtoMessage() {}

func (x *FriendUser) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendUser.ProtoReflect.Descriptor instead.
func (*FriendUser) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{8}
}

func (x *FriendUser) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *FriendUser) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FriendUser) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

var File_relation_proto protoreflect.FileDescriptor

var file_relation_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x15,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74,
	0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x16, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x22, 0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7b, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x47, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x7d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x12, 0x28, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5c, 0x0a, 0x0a, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x32, 0xa2, 0x02, 0x0a, 0x0f, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_relation_proto_rawDescOnce sync.Once
	file_relation_proto_rawDescData = file_relation_proto_rawDesc
)

func file_relation_proto_rawDescGZIP() []byte {
	file_relation_proto_rawDescOnce.Do(func() {
		file_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_relation_proto_rawDescData)
	})
	return file_relation_proto_rawDescData
}

var file_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_relation_proto_goTypes = []interface{}{
	(*RelationActionRequest)(nil),   // 0: RelationActionRequest
	(*RelationActionResponse)(nil),  // 1: RelationActionResponse
	(*GetFollowListRequest)(nil),    // 2: GetFollowListRequest
	(*GetFollowListResponse)(nil),   // 3: GetFollowListResponse
	(*GetFollowerListRequest)(nil),  // 4: GetFollowerListRequest
	(*GetFollowerListResponse)(nil), // 5: GetFollowerListResponse
	(*GetFriendListRequest)(nil),    // 6: GetFriendListRequest
	(*GetFriendListResponse)(nil),   // 7: GetFriendListResponse
	(*FriendUser)(nil),              // 8: FriendUser
	(*User)(nil),                    // 9: User
}
var file_relation_proto_depIdxs = []int32{
	9, // 0: GetFollowListResponse.user_list:type_name -> User
	9, // 1: GetFollowerListResponse.user_list:type_name -> User
	8, // 2: GetFriendListResponse.user_list:type_name -> FriendUser
	9, // 3: FriendUser.user:type_name -> User
	0, // 4: RelationService.RelationAction:input_type -> RelationActionRequest
	2, // 5: RelationService.GetFollowList:input_type -> GetFollowListRequest
	4, // 6: RelationService.GetFollowerList:input_type -> GetFollowerListRequest
	6, // 7: RelationService.GetFriendList:input_type -> GetFriendListRequest
	1, // 8: RelationService.RelationAction:output_type -> RelationActionResponse
	3, // 9: RelationService.GetFollowList:output_type -> GetFollowListResponse
	5, // 10: RelationService.GetFollowerList:output_type -> GetFollowerListResponse
	7, // 11: RelationService.GetFriendList:output_type -> GetFriendListResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_relation_proto_init() }
func file_relation_proto_init() {
	if File_relation_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionRequest); i {
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
		file_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionResponse); i {
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
		file_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowListRequest); i {
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
		file_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowListResponse); i {
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
		file_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowerListRequest); i {
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
		file_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowerListResponse); i {
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
		file_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendListRequest); i {
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
		file_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendListResponse); i {
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
		file_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendUser); i {
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
			RawDescriptor: file_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relation_proto_goTypes,
		DependencyIndexes: file_relation_proto_depIdxs,
		MessageInfos:      file_relation_proto_msgTypes,
	}.Build()
	File_relation_proto = out.File
	file_relation_proto_rawDesc = nil
	file_relation_proto_goTypes = nil
	file_relation_proto_depIdxs = nil
}
