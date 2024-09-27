// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: zitadel/user/v2beta/idp.proto

package user

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LDAPCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LDAPCredentials) Reset() {
	*x = LDAPCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LDAPCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LDAPCredentials) ProtoMessage() {}

func (x *LDAPCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LDAPCredentials.ProtoReflect.Descriptor instead.
func (*LDAPCredentials) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_idp_proto_rawDescGZIP(), []int{0}
}

func (x *LDAPCredentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LDAPCredentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RedirectURLs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessUrl string `protobuf:"bytes,1,opt,name=success_url,json=successUrl,proto3" json:"success_url,omitempty"`
	FailureUrl string `protobuf:"bytes,2,opt,name=failure_url,json=failureUrl,proto3" json:"failure_url,omitempty"`
}

func (x *RedirectURLs) Reset() {
	*x = RedirectURLs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectURLs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectURLs) ProtoMessage() {}

func (x *RedirectURLs) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectURLs.ProtoReflect.Descriptor instead.
func (*RedirectURLs) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_idp_proto_rawDescGZIP(), []int{1}
}

func (x *RedirectURLs) GetSuccessUrl() string {
	if x != nil {
		return x.SuccessUrl
	}
	return ""
}

func (x *RedirectURLs) GetFailureUrl() string {
	if x != nil {
		return x.FailureUrl
	}
	return ""
}

type IDPIntent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdpIntentId    string `protobuf:"bytes,1,opt,name=idp_intent_id,json=idpIntentId,proto3" json:"idp_intent_id,omitempty"`
	IdpIntentToken string `protobuf:"bytes,2,opt,name=idp_intent_token,json=idpIntentToken,proto3" json:"idp_intent_token,omitempty"`
	UserId         string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *IDPIntent) Reset() {
	*x = IDPIntent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPIntent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPIntent) ProtoMessage() {}

func (x *IDPIntent) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPIntent.ProtoReflect.Descriptor instead.
func (*IDPIntent) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_idp_proto_rawDescGZIP(), []int{2}
}

func (x *IDPIntent) GetIdpIntentId() string {
	if x != nil {
		return x.IdpIntentId
	}
	return ""
}

func (x *IDPIntent) GetIdpIntentToken() string {
	if x != nil {
		return x.IdpIntentToken
	}
	return ""
}

func (x *IDPIntent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type IDPInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Access:
	//
	//	*IDPInformation_Oauth
	//	*IDPInformation_Ldap
	//	*IDPInformation_Saml
	Access         isIDPInformation_Access `protobuf_oneof:"access"`
	IdpId          string                  `protobuf:"bytes,2,opt,name=idp_id,json=idpId,proto3" json:"idp_id,omitempty"`
	UserId         string                  `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName       string                  `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RawInformation *structpb.Struct        `protobuf:"bytes,5,opt,name=raw_information,json=rawInformation,proto3" json:"raw_information,omitempty"`
}

func (x *IDPInformation) Reset() {
	*x = IDPInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPInformation) ProtoMessage() {}

func (x *IDPInformation) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPInformation.ProtoReflect.Descriptor instead.
func (*IDPInformation) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_idp_proto_rawDescGZIP(), []int{3}
}

func (m *IDPInformation) GetAccess() isIDPInformation_Access {
	if m != nil {
		return m.Access
	}
	return nil
}

func (x *IDPInformation) GetOauth() *IDPOAuthAccessInformation {
	if x, ok := x.GetAccess().(*IDPInformation_Oauth); ok {
		return x.Oauth
	}
	return nil
}

func (x *IDPInformation) GetLdap() *IDPLDAPAccessInformation {
	if x, ok := x.GetAccess().(*IDPInformation_Ldap); ok {
		return x.Ldap
	}
	return nil
}

func (x *IDPInformation) GetSaml() *IDPSAMLAccessInformation {
	if x, ok := x.GetAccess().(*IDPInformation_Saml); ok {
		return x.Saml
	}
	return nil
}

func (x *IDPInformation) GetIdpId() string {
	if x != nil {
		return x.IdpId
	}
	return ""
}

func (x *IDPInformation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IDPInformation) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *IDPInformation) GetRawInformation() *structpb.Struct {
	if x != nil {
		return x.RawInformation
	}
	return nil
}

type isIDPInformation_Access interface {
	isIDPInformation_Access()
}

type IDPInformation_Oauth struct {
	Oauth *IDPOAuthAccessInformation `protobuf:"bytes,1,opt,name=oauth,proto3,oneof"`
}

type IDPInformation_Ldap struct {
	Ldap *IDPLDAPAccessInformation `protobuf:"bytes,6,opt,name=ldap,proto3,oneof"`
}

type IDPInformation_Saml struct {
	Saml *IDPSAMLAccessInformation `protobuf:"bytes,7,opt,name=saml,proto3,oneof"`
}

func (*IDPInformation_Oauth) isIDPInformation_Access() {}

func (*IDPInformation_Ldap) isIDPInformation_Access() {}

func (*IDPInformation_Saml) isIDPInformation_Access() {}

type IDPOAuthAccessInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string  `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	IdToken     *string `protobuf:"bytes,2,opt,name=id_token,json=idToken,proto3,oneof" json:"id_token,omitempty"`
}

func (x *IDPOAuthAccessInformation) Reset() {
	*x = IDPOAuthAccessInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPOAuthAccessInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPOAuthAccessInformation) ProtoMessage() {}

func (x *IDPOAuthAccessInformation) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPOAuthAccessInformation.ProtoReflect.Descriptor instead.
func (*IDPOAuthAccessInformation) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_idp_proto_rawDescGZIP(), []int{4}
}

func (x *IDPOAuthAccessInformation) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *IDPOAuthAccessInformation) GetIdToken() string {
	if x != nil && x.IdToken != nil {
		return *x.IdToken
	}
	return ""
}

type IDPLDAPAccessInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes *structpb.Struct `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *IDPLDAPAccessInformation) Reset() {
	*x = IDPLDAPAccessInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPLDAPAccessInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPLDAPAccessInformation) ProtoMessage() {}

func (x *IDPLDAPAccessInformation) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPLDAPAccessInformation.ProtoReflect.Descriptor instead.
func (*IDPLDAPAccessInformation) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_idp_proto_rawDescGZIP(), []int{5}
}

func (x *IDPLDAPAccessInformation) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type IDPSAMLAccessInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assertion []byte `protobuf:"bytes,1,opt,name=assertion,proto3" json:"assertion,omitempty"`
}

func (x *IDPSAMLAccessInformation) Reset() {
	*x = IDPSAMLAccessInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPSAMLAccessInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPSAMLAccessInformation) ProtoMessage() {}

func (x *IDPSAMLAccessInformation) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPSAMLAccessInformation.ProtoReflect.Descriptor instead.
func (*IDPSAMLAccessInformation) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_idp_proto_rawDescGZIP(), []int{6}
}

func (x *IDPSAMLAccessInformation) GetAssertion() []byte {
	if x != nil {
		return x.Assertion
	}
	return nil
}

type IDPLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdpId    string `protobuf:"bytes,1,opt,name=idp_id,json=idpId,proto3" json:"idp_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *IDPLink) Reset() {
	*x = IDPLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPLink) ProtoMessage() {}

func (x *IDPLink) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_idp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPLink.ProtoReflect.Descriptor instead.
func (*IDPLink) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_idp_proto_rawDescGZIP(), []int{7}
}

func (x *IDPLink) GetIdpId() string {
	if x != nil {
		return x.IdpId
	}
	return ""
}

func (x *IDPLink) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IDPLink) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

var File_zitadel_user_v2beta_idp_proto protoreflect.FileDescriptor

var file_zitadel_user_v2beta_idp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a,
	0x0f, 0x4c, 0x44, 0x41, 0x50, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x63, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x47, 0x92, 0x41, 0x37, 0x32, 0x23, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20,
	0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x20, 0x4c, 0x44, 0x41, 0x50, 0x4a, 0x0a, 0x22, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0xfa, 0x42,
	0x0a, 0x72, 0x08, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x90, 0x01, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x65, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x49, 0x92, 0x41, 0x39, 0x32, 0x23, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x20, 0x4c, 0x44, 0x41,
	0x50, 0x4a, 0x0c, 0x22, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x31, 0x21, 0x22, 0x78,
	0xc8, 0x01, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x90,
	0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xd2, 0x02, 0x0a,
	0x0c, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x52, 0x4c, 0x73, 0x12, 0xa3, 0x01,
	0x0a, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x81, 0x01, 0x92, 0x41, 0x71, 0x32, 0x41, 0x55, 0x52, 0x4c, 0x20, 0x6f,
	0x6e, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x20, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x61, 0x20, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4a, 0x26, 0x22, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x69, 0x64, 0x70, 0x2f, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x78, 0x80, 0x10, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10,
	0x01, 0x18, 0x80, 0x10, 0x90, 0x01, 0x01, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x55, 0x72, 0x6c, 0x12, 0x9b, 0x01, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x7a, 0x92, 0x41, 0x6a, 0x32, 0x3d,
	0x55, 0x52, 0x4c, 0x20, 0x6f, 0x6e, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x61,
	0x20, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4a, 0x23, 0x22,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x69, 0x64, 0x70, 0x2f, 0x66, 0x61, 0x69,
	0x6c, 0x22, 0x78, 0x80, 0x10, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x01, 0x18,
	0x80, 0x10, 0x90, 0x01, 0x01, 0x52, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x55, 0x72,
	0x6c, 0x22, 0xc9, 0x02, 0x0a, 0x09, 0x49, 0x44, 0x50, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x59, 0x0a, 0x0d, 0x69, 0x64, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0x92, 0x41, 0x32, 0x32, 0x14, 0x49, 0x44, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x49, 0x44, 0x50, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x4a, 0x14, 0x22, 0x31, 0x36, 0x33, 0x38, 0x34, 0x30, 0x37, 0x37, 0x36, 0x38, 0x33, 0x35,
	0x34, 0x33, 0x32, 0x37, 0x30, 0x35, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0x52, 0x0b, 0x69,
	0x64, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x72, 0x0a, 0x10, 0x69, 0x64,
	0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x48, 0x92, 0x41, 0x45, 0x32, 0x17, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x49, 0x44, 0x50, 0x20, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x4a, 0x24, 0x22, 0x53, 0x4a, 0x4b, 0x4c, 0x33, 0x69, 0x6f, 0x49, 0x44, 0x70, 0x6f,
	0x33, 0x34, 0x32, 0x69, 0x6f, 0x71, 0x77, 0x39, 0x38, 0x66, 0x6a, 0x70, 0x33, 0x73, 0x64, 0x66,
	0x33, 0x32, 0x77, 0x61, 0x68, 0x62, 0x3d, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0x52, 0x0e,
	0x69, 0x64, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x6d,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x54, 0x92, 0x41, 0x51, 0x32, 0x36, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x5a, 0x49, 0x54, 0x41, 0x44, 0x45, 0x4c, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x66, 0x20,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x6c,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x20, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x4a, 0x14, 0x22, 0x31,
	0x36, 0x33, 0x38, 0x34, 0x30, 0x37, 0x37, 0x36, 0x38, 0x33, 0x35, 0x34, 0x33, 0x32, 0x33, 0x34,
	0x35, 0x22, 0x78, 0xc8, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xd8, 0x06,
	0x0a, 0x0e, 0x49, 0x44, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x8f, 0x01, 0x0a, 0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x49, 0x44, 0x50, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x47, 0x92, 0x41, 0x44, 0x32, 0x42, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x2f, 0x4f, 0x49, 0x44,
	0x43, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x28, 0x61, 0x6e, 0x64, 0x20, 0x69, 0x64,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x29, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64,
	0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x82, 0x01, 0x0a, 0x04, 0x6c, 0x64, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x49, 0x44, 0x50, 0x4c, 0x44, 0x41, 0x50, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x3d, 0x92, 0x41, 0x3a, 0x32, 0x38, 0x4c, 0x44, 0x41, 0x50, 0x20, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x20, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x20, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x04, 0x6c, 0x64, 0x61, 0x70, 0x12, 0x78, 0x0a, 0x04, 0x73, 0x61, 0x6d, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x49, 0x44, 0x50, 0x53,
	0x41, 0x4d, 0x4c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x33, 0x92, 0x41, 0x30, 0x32, 0x2e, 0x53, 0x41, 0x4d, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64,
	0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x73, 0x61, 0x6d,
	0x6c, 0x12, 0x5f, 0x0a, 0x06, 0x69, 0x64, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x48, 0x92, 0x41, 0x45, 0x32, 0x1b, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x4a, 0x26, 0x22, 0x64, 0x36, 0x35, 0x34, 0x65, 0x36, 0x62, 0x61, 0x2d, 0x37,
	0x30, 0x61, 0x33, 0x2d, 0x34, 0x38, 0x65, 0x66, 0x2d, 0x61, 0x39, 0x35, 0x64, 0x2d, 0x33, 0x37,
	0x63, 0x38, 0x64, 0x38, 0x61, 0x37, 0x39, 0x30, 0x31, 0x61, 0x22, 0x52, 0x05, 0x69, 0x64, 0x70,
	0x49, 0x64, 0x12, 0x65, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x4c, 0x92, 0x41, 0x49, 0x32, 0x27, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x4a, 0x1e, 0x22, 0x36, 0x35, 0x31, 0x36, 0x38, 0x34, 0x39, 0x38, 0x30, 0x34, 0x38, 0x39,
	0x30, 0x34, 0x36, 0x38, 0x30, 0x34, 0x38, 0x34, 0x36, 0x31, 0x34, 0x30, 0x33, 0x35, 0x31, 0x38,
	0x22, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x64, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x47, 0x92, 0x41,
	0x44, 0x32, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x4a, 0x13, 0x22, 0x75, 0x73, 0x65, 0x72, 0x40, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x7d, 0x0a, 0x0f, 0x72, 0x61, 0x77, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x42, 0x3b, 0x92, 0x41, 0x38, 0x32, 0x36, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0e,
	0x72, 0x61, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x6b, 0x0a, 0x19, 0x49, 0x44, 0x50, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x69, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x64, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x53, 0x0a, 0x18, 0x49, 0x44, 0x50, 0x4c, 0x44, 0x41, 0x50,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x18, 0x49, 0x44,
	0x50, 0x53, 0x41, 0x4d, 0x4c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe7, 0x02, 0x0a, 0x07, 0x49, 0x44, 0x50, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x6f, 0x0a, 0x06, 0x69, 0x64, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x58, 0x92, 0x41, 0x4b, 0x32, 0x1b, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x4a, 0x26, 0x22, 0x64, 0x36, 0x35, 0x34, 0x65, 0x36, 0x62, 0x61, 0x2d, 0x37, 0x30,
	0x61, 0x33, 0x2d, 0x34, 0x38, 0x65, 0x66, 0x2d, 0x61, 0x39, 0x35, 0x64, 0x2d, 0x33, 0x37, 0x63,
	0x38, 0x64, 0x38, 0x61, 0x37, 0x39, 0x30, 0x31, 0x61, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x52, 0x05, 0x69, 0x64, 0x70, 0x49,
	0x64, 0x12, 0x75, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x5c, 0x92, 0x41, 0x4f, 0x32, 0x27, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x4a, 0x1e, 0x22, 0x36, 0x35, 0x31, 0x36, 0x38, 0x34, 0x39, 0x38, 0x30, 0x34, 0x38, 0x39, 0x30,
	0x34, 0x36, 0x38, 0x30, 0x34, 0x38, 0x34, 0x36, 0x31, 0x34, 0x30, 0x33, 0x35, 0x31, 0x38, 0x22,
	0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x74, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x57, 0x92, 0x41, 0x4a,
	0x32, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4a,
	0x13, 0x22, 0x75, 0x73, 0x65, 0x72, 0x40, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x6f, 0x6d, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10,
	0x01, 0x18, 0xc8, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x74,
	0x61, 0x64, 0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_user_v2beta_idp_proto_rawDescOnce sync.Once
	file_zitadel_user_v2beta_idp_proto_rawDescData = file_zitadel_user_v2beta_idp_proto_rawDesc
)

func file_zitadel_user_v2beta_idp_proto_rawDescGZIP() []byte {
	file_zitadel_user_v2beta_idp_proto_rawDescOnce.Do(func() {
		file_zitadel_user_v2beta_idp_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_user_v2beta_idp_proto_rawDescData)
	})
	return file_zitadel_user_v2beta_idp_proto_rawDescData
}

var file_zitadel_user_v2beta_idp_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_zitadel_user_v2beta_idp_proto_goTypes = []interface{}{
	(*LDAPCredentials)(nil),           // 0: zitadel.user.v2beta.LDAPCredentials
	(*RedirectURLs)(nil),              // 1: zitadel.user.v2beta.RedirectURLs
	(*IDPIntent)(nil),                 // 2: zitadel.user.v2beta.IDPIntent
	(*IDPInformation)(nil),            // 3: zitadel.user.v2beta.IDPInformation
	(*IDPOAuthAccessInformation)(nil), // 4: zitadel.user.v2beta.IDPOAuthAccessInformation
	(*IDPLDAPAccessInformation)(nil),  // 5: zitadel.user.v2beta.IDPLDAPAccessInformation
	(*IDPSAMLAccessInformation)(nil),  // 6: zitadel.user.v2beta.IDPSAMLAccessInformation
	(*IDPLink)(nil),                   // 7: zitadel.user.v2beta.IDPLink
	(*structpb.Struct)(nil),           // 8: google.protobuf.Struct
}
var file_zitadel_user_v2beta_idp_proto_depIdxs = []int32{
	4, // 0: zitadel.user.v2beta.IDPInformation.oauth:type_name -> zitadel.user.v2beta.IDPOAuthAccessInformation
	5, // 1: zitadel.user.v2beta.IDPInformation.ldap:type_name -> zitadel.user.v2beta.IDPLDAPAccessInformation
	6, // 2: zitadel.user.v2beta.IDPInformation.saml:type_name -> zitadel.user.v2beta.IDPSAMLAccessInformation
	8, // 3: zitadel.user.v2beta.IDPInformation.raw_information:type_name -> google.protobuf.Struct
	8, // 4: zitadel.user.v2beta.IDPLDAPAccessInformation.attributes:type_name -> google.protobuf.Struct
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_zitadel_user_v2beta_idp_proto_init() }
func file_zitadel_user_v2beta_idp_proto_init() {
	if File_zitadel_user_v2beta_idp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_user_v2beta_idp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LDAPCredentials); i {
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
		file_zitadel_user_v2beta_idp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectURLs); i {
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
		file_zitadel_user_v2beta_idp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDPIntent); i {
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
		file_zitadel_user_v2beta_idp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDPInformation); i {
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
		file_zitadel_user_v2beta_idp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDPOAuthAccessInformation); i {
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
		file_zitadel_user_v2beta_idp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDPLDAPAccessInformation); i {
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
		file_zitadel_user_v2beta_idp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDPSAMLAccessInformation); i {
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
		file_zitadel_user_v2beta_idp_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDPLink); i {
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
	file_zitadel_user_v2beta_idp_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*IDPInformation_Oauth)(nil),
		(*IDPInformation_Ldap)(nil),
		(*IDPInformation_Saml)(nil),
	}
	file_zitadel_user_v2beta_idp_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zitadel_user_v2beta_idp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_user_v2beta_idp_proto_goTypes,
		DependencyIndexes: file_zitadel_user_v2beta_idp_proto_depIdxs,
		MessageInfos:      file_zitadel_user_v2beta_idp_proto_msgTypes,
	}.Build()
	File_zitadel_user_v2beta_idp_proto = out.File
	file_zitadel_user_v2beta_idp_proto_rawDesc = nil
	file_zitadel_user_v2beta_idp_proto_goTypes = nil
	file_zitadel_user_v2beta_idp_proto_depIdxs = nil
}
