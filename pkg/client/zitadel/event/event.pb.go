// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: zitadel/event.proto

package event

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	message "github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/message"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Editor    *Editor    `protobuf:"bytes,1,opt,name=editor,proto3" json:"editor,omitempty"`
	Aggregate *Aggregate `protobuf:"bytes,2,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	Sequence  uint64     `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// The timestamp the event occurred
	CreationDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	Payload      *structpb.Struct       `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Type         *EventType             `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_zitadel_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEditor() *Editor {
	if x != nil {
		return x.Editor
	}
	return nil
}

func (x *Event) GetAggregate() *Aggregate {
	if x != nil {
		return x.Aggregate
	}
	return nil
}

func (x *Event) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Event) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

func (x *Event) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Event) GetType() *EventType {
	if x != nil {
		return x.Type
	}
	return nil
}

type Editor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Service     string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *Editor) Reset() {
	*x = Editor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Editor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Editor) ProtoMessage() {}

func (x *Editor) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Editor.ProtoReflect.Descriptor instead.
func (*Editor) Descriptor() ([]byte, []int) {
	return file_zitadel_event_proto_rawDescGZIP(), []int{1}
}

func (x *Editor) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Editor) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Editor) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type Aggregate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          *AggregateType `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ResourceOwner string         `protobuf:"bytes,3,opt,name=resource_owner,json=resourceOwner,proto3" json:"resource_owner,omitempty"`
}

func (x *Aggregate) Reset() {
	*x = Aggregate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aggregate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aggregate) ProtoMessage() {}

func (x *Aggregate) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aggregate.ProtoReflect.Descriptor instead.
func (*Aggregate) Descriptor() ([]byte, []int) {
	return file_zitadel_event_proto_rawDescGZIP(), []int{2}
}

func (x *Aggregate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Aggregate) GetType() *AggregateType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Aggregate) GetResourceOwner() string {
	if x != nil {
		return x.ResourceOwner
	}
	return ""
}

type EventType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string                    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Localized *message.LocalizedMessage `protobuf:"bytes,2,opt,name=localized,proto3" json:"localized,omitempty"`
}

func (x *EventType) Reset() {
	*x = EventType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventType) ProtoMessage() {}

func (x *EventType) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventType.ProtoReflect.Descriptor instead.
func (*EventType) Descriptor() ([]byte, []int) {
	return file_zitadel_event_proto_rawDescGZIP(), []int{3}
}

func (x *EventType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EventType) GetLocalized() *message.LocalizedMessage {
	if x != nil {
		return x.Localized
	}
	return nil
}

type AggregateType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string                    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Localized *message.LocalizedMessage `protobuf:"bytes,2,opt,name=localized,proto3" json:"localized,omitempty"`
}

func (x *AggregateType) Reset() {
	*x = AggregateType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateType) ProtoMessage() {}

func (x *AggregateType) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateType.ProtoReflect.Descriptor instead.
func (*AggregateType) Descriptor() ([]byte, []int) {
	return file_zitadel_event_proto_rawDescGZIP(), []int{4}
}

func (x *AggregateType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AggregateType) GetLocalized() *message.LocalizedMessage {
	if x != nil {
		return x.Localized
	}
	return nil
}

var File_zitadel_event_proto protoreflect.FileDescriptor

var file_zitadel_event_proto_rawDesc = []byte{
	0x0a, 0x13, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x02,
	0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x6f,
	0x72, 0x52, 0x06, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x7a,
	0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x63, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x22, 0x92, 0x41, 0x1f, 0x4a, 0x1d, 0x22, 0x32, 0x30, 0x31, 0x39, 0x2d,
	0x30, 0x34, 0x2d, 0x30, 0x31, 0x54, 0x30, 0x38, 0x3a, 0x34, 0x35, 0x3a, 0x30, 0x30, 0x2e, 0x30,
	0x30, 0x30, 0x30, 0x30, 0x30, 0x5a, 0x22, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5e, 0x0a, 0x06, 0x45, 0x64, 0x69,
	0x74, 0x6f, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x77, 0x0a, 0x09, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x22, 0x5b, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x22,
	0x5f, 0x0a, 0x0d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a,
	0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_event_proto_rawDescOnce sync.Once
	file_zitadel_event_proto_rawDescData = file_zitadel_event_proto_rawDesc
)

func file_zitadel_event_proto_rawDescGZIP() []byte {
	file_zitadel_event_proto_rawDescOnce.Do(func() {
		file_zitadel_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_event_proto_rawDescData)
	})
	return file_zitadel_event_proto_rawDescData
}

var file_zitadel_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_zitadel_event_proto_goTypes = []interface{}{
	(*Event)(nil),                    // 0: zitadel.event.v1.Event
	(*Editor)(nil),                   // 1: zitadel.event.v1.Editor
	(*Aggregate)(nil),                // 2: zitadel.event.v1.Aggregate
	(*EventType)(nil),                // 3: zitadel.event.v1.EventType
	(*AggregateType)(nil),            // 4: zitadel.event.v1.AggregateType
	(*timestamppb.Timestamp)(nil),    // 5: google.protobuf.Timestamp
	(*structpb.Struct)(nil),          // 6: google.protobuf.Struct
	(*message.LocalizedMessage)(nil), // 7: zitadel.v1.LocalizedMessage
}
var file_zitadel_event_proto_depIdxs = []int32{
	1, // 0: zitadel.event.v1.Event.editor:type_name -> zitadel.event.v1.Editor
	2, // 1: zitadel.event.v1.Event.aggregate:type_name -> zitadel.event.v1.Aggregate
	5, // 2: zitadel.event.v1.Event.creation_date:type_name -> google.protobuf.Timestamp
	6, // 3: zitadel.event.v1.Event.payload:type_name -> google.protobuf.Struct
	3, // 4: zitadel.event.v1.Event.type:type_name -> zitadel.event.v1.EventType
	4, // 5: zitadel.event.v1.Aggregate.type:type_name -> zitadel.event.v1.AggregateType
	7, // 6: zitadel.event.v1.EventType.localized:type_name -> zitadel.v1.LocalizedMessage
	7, // 7: zitadel.event.v1.AggregateType.localized:type_name -> zitadel.v1.LocalizedMessage
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_zitadel_event_proto_init() }
func file_zitadel_event_proto_init() {
	if File_zitadel_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_zitadel_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Editor); i {
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
		file_zitadel_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aggregate); i {
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
		file_zitadel_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventType); i {
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
		file_zitadel_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateType); i {
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
			RawDescriptor: file_zitadel_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_event_proto_goTypes,
		DependencyIndexes: file_zitadel_event_proto_depIdxs,
		MessageInfos:      file_zitadel_event_proto_msgTypes,
	}.Build()
	File_zitadel_event_proto = out.File
	file_zitadel_event_proto_rawDesc = nil
	file_zitadel_event_proto_goTypes = nil
	file_zitadel_event_proto_depIdxs = nil
}
