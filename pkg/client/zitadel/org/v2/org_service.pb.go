// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: zitadel/org/v2/org_service.proto

package org

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v2 "github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/object/v2"
	_ "github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/protoc/v2"
	v22 "github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/user/v2"
	v21 "github.com/zitadel/zitadel/pkg/grpc/org/v2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddOrganizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Admins []*AddOrganizationRequest_Admin `protobuf:"bytes,2,rep,name=admins,proto3" json:"admins,omitempty"`
}

func (x *AddOrganizationRequest) Reset() {
	*x = AddOrganizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_org_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrganizationRequest) ProtoMessage() {}

func (x *AddOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_org_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrganizationRequest.ProtoReflect.Descriptor instead.
func (*AddOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_org_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddOrganizationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddOrganizationRequest) GetAdmins() []*AddOrganizationRequest_Admin {
	if x != nil {
		return x.Admins
	}
	return nil
}

type AddOrganizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details        *v2.Details                             `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	OrganizationId string                                  `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	CreatedAdmins  []*AddOrganizationResponse_CreatedAdmin `protobuf:"bytes,3,rep,name=created_admins,json=createdAdmins,proto3" json:"created_admins,omitempty"`
}

func (x *AddOrganizationResponse) Reset() {
	*x = AddOrganizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_org_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrganizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrganizationResponse) ProtoMessage() {}

func (x *AddOrganizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_org_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrganizationResponse.ProtoReflect.Descriptor instead.
func (*AddOrganizationResponse) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_org_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddOrganizationResponse) GetDetails() *v2.Details {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *AddOrganizationResponse) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *AddOrganizationResponse) GetCreatedAdmins() []*AddOrganizationResponse_CreatedAdmin {
	if x != nil {
		return x.CreatedAdmins
	}
	return nil
}

type ListOrganizationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list limitations and ordering
	Query *v2.ListQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// the field the result is sorted
	SortingColumn v21.OrganizationFieldName `protobuf:"varint,2,opt,name=sorting_column,json=sortingColumn,proto3,enum=zitadel.org.v2.OrganizationFieldName" json:"sorting_column,omitempty"`
	// criteria the client is looking for
	Queries []*v21.SearchQuery `protobuf:"bytes,3,rep,name=queries,proto3" json:"queries,omitempty"`
}

func (x *ListOrganizationsRequest) Reset() {
	*x = ListOrganizationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_org_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrganizationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationsRequest) ProtoMessage() {}

func (x *ListOrganizationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_org_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationsRequest.ProtoReflect.Descriptor instead.
func (*ListOrganizationsRequest) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_org_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListOrganizationsRequest) GetQuery() *v2.ListQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ListOrganizationsRequest) GetSortingColumn() v21.OrganizationFieldName {
	if x != nil {
		return x.SortingColumn
	}
	return v21.OrganizationFieldName(0)
}

func (x *ListOrganizationsRequest) GetQueries() []*v21.SearchQuery {
	if x != nil {
		return x.Queries
	}
	return nil
}

type ListOrganizationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details       *v2.ListDetails           `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	SortingColumn v21.OrganizationFieldName `protobuf:"varint,2,opt,name=sorting_column,json=sortingColumn,proto3,enum=zitadel.org.v2.OrganizationFieldName" json:"sorting_column,omitempty"`
	Result        []*v21.Organization       `protobuf:"bytes,3,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ListOrganizationsResponse) Reset() {
	*x = ListOrganizationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_org_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrganizationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrganizationsResponse) ProtoMessage() {}

func (x *ListOrganizationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_org_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrganizationsResponse.ProtoReflect.Descriptor instead.
func (*ListOrganizationsResponse) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_org_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListOrganizationsResponse) GetDetails() *v2.ListDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ListOrganizationsResponse) GetSortingColumn() v21.OrganizationFieldName {
	if x != nil {
		return x.SortingColumn
	}
	return v21.OrganizationFieldName(0)
}

func (x *ListOrganizationsResponse) GetResult() []*v21.Organization {
	if x != nil {
		return x.Result
	}
	return nil
}

type AddOrganizationRequest_Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UserType:
	//
	//	*AddOrganizationRequest_Admin_UserId
	//	*AddOrganizationRequest_Admin_Human
	UserType isAddOrganizationRequest_Admin_UserType `protobuf_oneof:"user_type"`
	// specify Org Member Roles for the provided user (default is ORG_OWNER if roles are empty)
	Roles []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *AddOrganizationRequest_Admin) Reset() {
	*x = AddOrganizationRequest_Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_org_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrganizationRequest_Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrganizationRequest_Admin) ProtoMessage() {}

func (x *AddOrganizationRequest_Admin) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_org_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrganizationRequest_Admin.ProtoReflect.Descriptor instead.
func (*AddOrganizationRequest_Admin) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_org_service_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AddOrganizationRequest_Admin) GetUserType() isAddOrganizationRequest_Admin_UserType {
	if m != nil {
		return m.UserType
	}
	return nil
}

func (x *AddOrganizationRequest_Admin) GetUserId() string {
	if x, ok := x.GetUserType().(*AddOrganizationRequest_Admin_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *AddOrganizationRequest_Admin) GetHuman() *v22.AddHumanUserRequest {
	if x, ok := x.GetUserType().(*AddOrganizationRequest_Admin_Human); ok {
		return x.Human
	}
	return nil
}

func (x *AddOrganizationRequest_Admin) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type isAddOrganizationRequest_Admin_UserType interface {
	isAddOrganizationRequest_Admin_UserType()
}

type AddOrganizationRequest_Admin_UserId struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3,oneof"`
}

type AddOrganizationRequest_Admin_Human struct {
	Human *v22.AddHumanUserRequest `protobuf:"bytes,2,opt,name=human,proto3,oneof"`
}

func (*AddOrganizationRequest_Admin_UserId) isAddOrganizationRequest_Admin_UserType() {}

func (*AddOrganizationRequest_Admin_Human) isAddOrganizationRequest_Admin_UserType() {}

type AddOrganizationResponse_CreatedAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	EmailCode *string `protobuf:"bytes,2,opt,name=email_code,json=emailCode,proto3,oneof" json:"email_code,omitempty"`
	PhoneCode *string `protobuf:"bytes,3,opt,name=phone_code,json=phoneCode,proto3,oneof" json:"phone_code,omitempty"`
}

func (x *AddOrganizationResponse_CreatedAdmin) Reset() {
	*x = AddOrganizationResponse_CreatedAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_org_v2_org_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrganizationResponse_CreatedAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrganizationResponse_CreatedAdmin) ProtoMessage() {}

func (x *AddOrganizationResponse_CreatedAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_org_v2_org_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrganizationResponse_CreatedAdmin.ProtoReflect.Descriptor instead.
func (*AddOrganizationResponse_CreatedAdmin) Descriptor() ([]byte, []int) {
	return file_zitadel_org_v2_org_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AddOrganizationResponse_CreatedAdmin) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddOrganizationResponse_CreatedAdmin) GetEmailCode() string {
	if x != nil && x.EmailCode != nil {
		return *x.EmailCode
	}
	return ""
}

func (x *AddOrganizationResponse_CreatedAdmin) GetPhoneCode() string {
	if x != nil && x.PhoneCode != nil {
		return *x.PhoneCode
	}
	return ""
}

var File_zitadel_org_v2_org_service_proto protoreflect.FileDescriptor

var file_zitadel_org_v2_org_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x32,
	0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x76, 0x32, 0x1a, 0x1e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x7a, 0x69, 0x74,
	0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x7a, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f,
	0x6f, 0x72, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x32,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a,
	0x16, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0x92, 0x41, 0x11, 0x4a, 0x09, 0x22, 0x5a, 0x49, 0x54,
	0x41, 0x44, 0x45, 0x4c, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0xe0, 0x41, 0x02, 0xfa, 0x42,
	0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44,
	0x0a, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x64, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x06, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x73, 0x1a, 0x83, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x05, 0x68, 0x75, 0x6d,
	0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x75,
	0x6d, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x05, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x0b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe5, 0x02, 0x0a, 0x17, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x5b, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x73, 0x1a, 0x8d, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0a,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x22, 0x0a, 0x0a, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x4c, 0x0a, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x7a, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x12, 0x35, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0xd9, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x4c, 0x0a, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x0d, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x34,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0xdc, 0x03, 0x0a, 0x13, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7, 0x01, 0x0a,
	0x0f, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x43, 0x92, 0x41, 0x0d, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a,
	0x02, 0x4f, 0x4b, 0x8a, 0xb5, 0x18, 0x13, 0x0a, 0x0c, 0x0a, 0x0a, 0x6f, 0x72, 0x67, 0x2e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x03, 0x08, 0xc9, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x9a, 0x02, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x7a,
	0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xaf, 0x01, 0x92, 0x41, 0x73, 0x4a, 0x37, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x30,
	0x0a, 0x2e, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x6c, 0x6c, 0x20,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x4a, 0x38, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x31, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x8a, 0xb5, 0x18, 0x11, 0x0a, 0x0f,
	0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x5f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0xaa, 0x07, 0x92, 0x41, 0xf6, 0x06, 0x12, 0xd4, 0x01, 0x0a, 0x14, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x43, 0x54, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73,
	0x20, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x20, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x5a, 0x49, 0x54, 0x41, 0x44, 0x45, 0x4c, 0x20, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x22, 0x2e, 0x0a, 0x07, 0x5a, 0x49, 0x54, 0x41,
	0x44, 0x45, 0x4c, 0x12, 0x13, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x7a, 0x69, 0x74,
	0x61, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x1a, 0x0e, 0x68, 0x69, 0x40, 0x7a, 0x69, 0x74,
	0x61, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x42, 0x0a, 0x0a, 0x41, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x20, 0x32, 0x2e, 0x30, 0x12, 0x34, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x03, 0x32, 0x2e,
	0x30, 0x1a, 0x0e, 0x24, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x2d, 0x44, 0x4f, 0x4d, 0x41, 0x49,
	0x4e, 0x22, 0x01, 0x2f, 0x2a, 0x02, 0x02, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x32, 0x1a, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x77,
	0x65, 0x62, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3a, 0x1a, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x77,
	0x65, 0x62, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x6d, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12,
	0x66, 0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e,
	0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17,
	0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x49,
	0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65,
	0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x1b, 0x0a, 0x19,
	0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5a, 0xc2, 0x01, 0x0a, 0xbf, 0x01, 0x0a,
	0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0xb4, 0x01, 0x08, 0x03, 0x28, 0x04, 0x32, 0x21,
	0x24, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x2d, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x2f, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x3a, 0x1d, 0x24, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x2d, 0x44, 0x4f, 0x4d, 0x41, 0x49,
	0x4e, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x42, 0x6c, 0x0a, 0x10, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x06, 0x6f, 0x70,
	0x65, 0x6e, 0x69, 0x64, 0x0a, 0x58, 0x0a, 0x2a, 0x75, 0x72, 0x6e, 0x3a, 0x7a, 0x69, 0x74, 0x61,
	0x64, 0x65, 0x6c, 0x3a, 0x69, 0x61, 0x6d, 0x3a, 0x6f, 0x72, 0x67, 0x3a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x3a, 0x69, 0x64, 0x3a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x3a, 0x61,
	0x75, 0x64, 0x12, 0x2a, 0x75, 0x72, 0x6e, 0x3a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x3a,
	0x69, 0x61, 0x6d, 0x3a, 0x6f, 0x72, 0x67, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a,
	0x69, 0x64, 0x3a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x3a, 0x61, 0x75, 0x64, 0x62, 0x40,
	0x0a, 0x3e, 0x0a, 0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x34, 0x0a, 0x06, 0x6f, 0x70,
	0x65, 0x6e, 0x69, 0x64, 0x0a, 0x2a, 0x75, 0x72, 0x6e, 0x3a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x3a, 0x69, 0x61, 0x6d, 0x3a, 0x6f, 0x72, 0x67, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x3a, 0x69, 0x64, 0x3a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x3a, 0x61, 0x75, 0x64,
	0x72, 0x3e, 0x0a, 0x22, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x5a,
	0x49, 0x54, 0x41, 0x44, 0x45, 0x4c, 0x12, 0x18, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x63, 0x73,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x74,
	0x61, 0x64, 0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x6f, 0x72, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_org_v2_org_service_proto_rawDescOnce sync.Once
	file_zitadel_org_v2_org_service_proto_rawDescData = file_zitadel_org_v2_org_service_proto_rawDesc
)

func file_zitadel_org_v2_org_service_proto_rawDescGZIP() []byte {
	file_zitadel_org_v2_org_service_proto_rawDescOnce.Do(func() {
		file_zitadel_org_v2_org_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_org_v2_org_service_proto_rawDescData)
	})
	return file_zitadel_org_v2_org_service_proto_rawDescData
}

var file_zitadel_org_v2_org_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_zitadel_org_v2_org_service_proto_goTypes = []interface{}{
	(*AddOrganizationRequest)(nil),               // 0: zitadel.org.v2.AddOrganizationRequest
	(*AddOrganizationResponse)(nil),              // 1: zitadel.org.v2.AddOrganizationResponse
	(*ListOrganizationsRequest)(nil),             // 2: zitadel.org.v2.ListOrganizationsRequest
	(*ListOrganizationsResponse)(nil),            // 3: zitadel.org.v2.ListOrganizationsResponse
	(*AddOrganizationRequest_Admin)(nil),         // 4: zitadel.org.v2.AddOrganizationRequest.Admin
	(*AddOrganizationResponse_CreatedAdmin)(nil), // 5: zitadel.org.v2.AddOrganizationResponse.CreatedAdmin
	(*v2.Details)(nil),                           // 6: zitadel.object.v2.Details
	(*v2.ListQuery)(nil),                         // 7: zitadel.object.v2.ListQuery
	(v21.OrganizationFieldName)(0),               // 8: zitadel.org.v2.OrganizationFieldName
	(*v21.SearchQuery)(nil),                      // 9: zitadel.org.v2.SearchQuery
	(*v2.ListDetails)(nil),                       // 10: zitadel.object.v2.ListDetails
	(*v21.Organization)(nil),                     // 11: zitadel.org.v2.Organization
	(*v22.AddHumanUserRequest)(nil),              // 12: zitadel.user.v2.AddHumanUserRequest
}
var file_zitadel_org_v2_org_service_proto_depIdxs = []int32{
	4,  // 0: zitadel.org.v2.AddOrganizationRequest.admins:type_name -> zitadel.org.v2.AddOrganizationRequest.Admin
	6,  // 1: zitadel.org.v2.AddOrganizationResponse.details:type_name -> zitadel.object.v2.Details
	5,  // 2: zitadel.org.v2.AddOrganizationResponse.created_admins:type_name -> zitadel.org.v2.AddOrganizationResponse.CreatedAdmin
	7,  // 3: zitadel.org.v2.ListOrganizationsRequest.query:type_name -> zitadel.object.v2.ListQuery
	8,  // 4: zitadel.org.v2.ListOrganizationsRequest.sorting_column:type_name -> zitadel.org.v2.OrganizationFieldName
	9,  // 5: zitadel.org.v2.ListOrganizationsRequest.queries:type_name -> zitadel.org.v2.SearchQuery
	10, // 6: zitadel.org.v2.ListOrganizationsResponse.details:type_name -> zitadel.object.v2.ListDetails
	8,  // 7: zitadel.org.v2.ListOrganizationsResponse.sorting_column:type_name -> zitadel.org.v2.OrganizationFieldName
	11, // 8: zitadel.org.v2.ListOrganizationsResponse.result:type_name -> zitadel.org.v2.Organization
	12, // 9: zitadel.org.v2.AddOrganizationRequest.Admin.human:type_name -> zitadel.user.v2.AddHumanUserRequest
	0,  // 10: zitadel.org.v2.OrganizationService.AddOrganization:input_type -> zitadel.org.v2.AddOrganizationRequest
	2,  // 11: zitadel.org.v2.OrganizationService.ListOrganizations:input_type -> zitadel.org.v2.ListOrganizationsRequest
	1,  // 12: zitadel.org.v2.OrganizationService.AddOrganization:output_type -> zitadel.org.v2.AddOrganizationResponse
	3,  // 13: zitadel.org.v2.OrganizationService.ListOrganizations:output_type -> zitadel.org.v2.ListOrganizationsResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_zitadel_org_v2_org_service_proto_init() }
func file_zitadel_org_v2_org_service_proto_init() {
	if File_zitadel_org_v2_org_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_org_v2_org_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrganizationRequest); i {
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
		file_zitadel_org_v2_org_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrganizationResponse); i {
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
		file_zitadel_org_v2_org_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrganizationsRequest); i {
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
		file_zitadel_org_v2_org_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrganizationsResponse); i {
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
		file_zitadel_org_v2_org_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrganizationRequest_Admin); i {
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
		file_zitadel_org_v2_org_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrganizationResponse_CreatedAdmin); i {
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
	file_zitadel_org_v2_org_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*AddOrganizationRequest_Admin_UserId)(nil),
		(*AddOrganizationRequest_Admin_Human)(nil),
	}
	file_zitadel_org_v2_org_service_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zitadel_org_v2_org_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zitadel_org_v2_org_service_proto_goTypes,
		DependencyIndexes: file_zitadel_org_v2_org_service_proto_depIdxs,
		MessageInfos:      file_zitadel_org_v2_org_service_proto_msgTypes,
	}.Build()
	File_zitadel_org_v2_org_service_proto = out.File
	file_zitadel_org_v2_org_service_proto_rawDesc = nil
	file_zitadel_org_v2_org_service_proto_goTypes = nil
	file_zitadel_org_v2_org_service_proto_depIdxs = nil
}
