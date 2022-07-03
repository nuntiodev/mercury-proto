// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: mercury.proto

package go_mercury

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image     string                 `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Email     string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConversationId string                 `protobuf:"bytes,2,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	UserId         string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SentAt         *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *Message) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Message) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

type Conversation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AdminId   string                 `protobuf:"bytes,2,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	Image     *string                `protobuf:"bytes,3,opt,name=image,proto3,oneof" json:"image,omitempty"`
	Users     []string               `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Conversation) Reset() {
	*x = Conversation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conversation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conversation) ProtoMessage() {}

func (x *Conversation) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conversation.ProtoReflect.Descriptor instead.
func (*Conversation) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{2}
}

func (x *Conversation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Conversation) GetAdminId() string {
	if x != nil {
		return x.AdminId
	}
	return ""
}

func (x *Conversation) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *Conversation) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Conversation) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type MercuryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace    string        `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	User         *User         `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Message      *Message      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Conversation *Conversation `protobuf:"bytes,4,opt,name=conversation,proto3" json:"conversation,omitempty"`
	From         int32         `protobuf:"varint,5,opt,name=from,proto3" json:"from,omitempty"`
	To           int32         `protobuf:"varint,6,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *MercuryRequest) Reset() {
	*x = MercuryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercuryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercuryRequest) ProtoMessage() {}

func (x *MercuryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercuryRequest.ProtoReflect.Descriptor instead.
func (*MercuryRequest) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{3}
}

func (x *MercuryRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *MercuryRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *MercuryRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *MercuryRequest) GetConversation() *Conversation {
	if x != nil {
		return x.Conversation
	}
	return nil
}

func (x *MercuryRequest) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *MercuryRequest) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

type MercuryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User          *User           `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users         []*User         `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Message       *Message        `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Messages      []*Message      `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
	Conversation  *Conversation   `protobuf:"bytes,5,opt,name=conversation,proto3" json:"conversation,omitempty"`
	Conversations []*Conversation `protobuf:"bytes,6,rep,name=conversations,proto3" json:"conversations,omitempty"`
}

func (x *MercuryResponse) Reset() {
	*x = MercuryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercuryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercuryResponse) ProtoMessage() {}

func (x *MercuryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercuryResponse.ProtoReflect.Descriptor instead.
func (*MercuryResponse) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{4}
}

func (x *MercuryResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *MercuryResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *MercuryResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *MercuryResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *MercuryResponse) GetConversation() *Conversation {
	if x != nil {
		return x.Conversation
	}
	return nil
}

func (x *MercuryResponse) GetConversations() []*Conversation {
	if x != nil {
		return x.Conversations
	}
	return nil
}

var File_mercury_proto protoreflect.FileDescriptor

var file_mercury_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0xdc, 0x01,
	0x0a, 0x0e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a,
	0x0f, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75,
	0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x93, 0x07, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x17,
	0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e,
	0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75,
	0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x4d, 0x65,
	0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x15, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x1a, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04,
	0x53, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17,
	0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x4d, 0x65,
	0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mercury_proto_rawDescOnce sync.Once
	file_mercury_proto_rawDescData = file_mercury_proto_rawDesc
)

func file_mercury_proto_rawDescGZIP() []byte {
	file_mercury_proto_rawDescOnce.Do(func() {
		file_mercury_proto_rawDescData = protoimpl.X.CompressGZIP(file_mercury_proto_rawDescData)
	})
	return file_mercury_proto_rawDescData
}

var file_mercury_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mercury_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: Mercury.User
	(*Message)(nil),               // 1: Mercury.Message
	(*Conversation)(nil),          // 2: Mercury.Conversation
	(*MercuryRequest)(nil),        // 3: Mercury.MercuryRequest
	(*MercuryResponse)(nil),       // 4: Mercury.MercuryResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_mercury_proto_depIdxs = []int32{
	5,  // 0: Mercury.User.created_at:type_name -> google.protobuf.Timestamp
	5,  // 1: Mercury.User.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 2: Mercury.Message.sent_at:type_name -> google.protobuf.Timestamp
	5,  // 3: Mercury.Conversation.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 4: Mercury.MercuryRequest.user:type_name -> Mercury.User
	1,  // 5: Mercury.MercuryRequest.message:type_name -> Mercury.Message
	2,  // 6: Mercury.MercuryRequest.conversation:type_name -> Mercury.Conversation
	0,  // 7: Mercury.MercuryResponse.user:type_name -> Mercury.User
	0,  // 8: Mercury.MercuryResponse.users:type_name -> Mercury.User
	1,  // 9: Mercury.MercuryResponse.message:type_name -> Mercury.Message
	1,  // 10: Mercury.MercuryResponse.messages:type_name -> Mercury.Message
	2,  // 11: Mercury.MercuryResponse.conversation:type_name -> Mercury.Conversation
	2,  // 12: Mercury.MercuryResponse.conversations:type_name -> Mercury.Conversation
	3,  // 13: Mercury.Service.Ping:input_type -> Mercury.MercuryRequest
	3,  // 14: Mercury.Service.CreateUser:input_type -> Mercury.MercuryRequest
	3,  // 15: Mercury.Service.GetUser:input_type -> Mercury.MercuryRequest
	3,  // 16: Mercury.Service.GetAllUsers:input_type -> Mercury.MercuryRequest
	3,  // 17: Mercury.Service.UpdateUser:input_type -> Mercury.MercuryRequest
	3,  // 18: Mercury.Service.DeleteUser:input_type -> Mercury.MercuryRequest
	3,  // 19: Mercury.Service.CreateConversation:input_type -> Mercury.MercuryRequest
	3,  // 20: Mercury.Service.AddUserToConversation:input_type -> Mercury.MercuryRequest
	3,  // 21: Mercury.Service.RemoveUserFromConversation:input_type -> Mercury.MercuryRequest
	3,  // 22: Mercury.Service.Send:input_type -> Mercury.MercuryRequest
	3,  // 23: Mercury.Service.ListConversations:input_type -> Mercury.MercuryRequest
	3,  // 24: Mercury.Service.DeleteConversation:input_type -> Mercury.MercuryRequest
	3,  // 25: Mercury.Service.Heartbeat:input_type -> Mercury.MercuryRequest
	4,  // 26: Mercury.Service.Ping:output_type -> Mercury.MercuryResponse
	4,  // 27: Mercury.Service.CreateUser:output_type -> Mercury.MercuryResponse
	4,  // 28: Mercury.Service.GetUser:output_type -> Mercury.MercuryResponse
	4,  // 29: Mercury.Service.GetAllUsers:output_type -> Mercury.MercuryResponse
	4,  // 30: Mercury.Service.UpdateUser:output_type -> Mercury.MercuryResponse
	4,  // 31: Mercury.Service.DeleteUser:output_type -> Mercury.MercuryResponse
	4,  // 32: Mercury.Service.CreateConversation:output_type -> Mercury.MercuryResponse
	4,  // 33: Mercury.Service.AddUserToConversation:output_type -> Mercury.MercuryResponse
	4,  // 34: Mercury.Service.RemoveUserFromConversation:output_type -> Mercury.MercuryResponse
	4,  // 35: Mercury.Service.Send:output_type -> Mercury.MercuryResponse
	4,  // 36: Mercury.Service.ListConversations:output_type -> Mercury.MercuryResponse
	4,  // 37: Mercury.Service.DeleteConversation:output_type -> Mercury.MercuryResponse
	4,  // 38: Mercury.Service.Heartbeat:output_type -> Mercury.MercuryResponse
	26, // [26:39] is the sub-list for method output_type
	13, // [13:26] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_mercury_proto_init() }
func file_mercury_proto_init() {
	if File_mercury_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mercury_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_mercury_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_mercury_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conversation); i {
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
		file_mercury_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercuryRequest); i {
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
		file_mercury_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercuryResponse); i {
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
	file_mercury_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mercury_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mercury_proto_goTypes,
		DependencyIndexes: file_mercury_proto_depIdxs,
		MessageInfos:      file_mercury_proto_msgTypes,
	}.Build()
	File_mercury_proto = out.File
	file_mercury_proto_rawDesc = nil
	file_mercury_proto_goTypes = nil
	file_mercury_proto_depIdxs = nil
}
