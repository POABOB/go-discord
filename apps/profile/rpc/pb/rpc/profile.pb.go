// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: apps/profile/rpc/profile.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUniqueProfileOrCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageUrl string `protobuf:"bytes,3,opt,name=ImageUrl,proto3" json:"ImageUrl,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *GetUniqueProfileOrCreateReq) Reset() {
	*x = GetUniqueProfileOrCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_profile_rpc_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUniqueProfileOrCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUniqueProfileOrCreateReq) ProtoMessage() {}

func (x *GetUniqueProfileOrCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_profile_rpc_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUniqueProfileOrCreateReq.ProtoReflect.Descriptor instead.
func (*GetUniqueProfileOrCreateReq) Descriptor() ([]byte, []int) {
	return file_apps_profile_rpc_profile_proto_rawDescGZIP(), []int{0}
}

func (x *GetUniqueProfileOrCreateReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUniqueProfileOrCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUniqueProfileOrCreateReq) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *GetUniqueProfileOrCreateReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetUniqueProfileOrCreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageUrl string `protobuf:"bytes,4,opt,name=ImageUrl,proto3" json:"ImageUrl,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *GetUniqueProfileOrCreateRes) Reset() {
	*x = GetUniqueProfileOrCreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_profile_rpc_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUniqueProfileOrCreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUniqueProfileOrCreateRes) ProtoMessage() {}

func (x *GetUniqueProfileOrCreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_apps_profile_rpc_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUniqueProfileOrCreateRes.ProtoReflect.Descriptor instead.
func (*GetUniqueProfileOrCreateRes) Descriptor() ([]byte, []int) {
	return file_apps_profile_rpc_profile_proto_rawDescGZIP(), []int{1}
}

func (x *GetUniqueProfileOrCreateRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetUniqueProfileOrCreateRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUniqueProfileOrCreateRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUniqueProfileOrCreateRes) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *GetUniqueProfileOrCreateRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PatchProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageUrl string `protobuf:"bytes,4,opt,name=ImageUrl,proto3" json:"ImageUrl,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *PatchProfileReq) Reset() {
	*x = PatchProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_profile_rpc_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchProfileReq) ProtoMessage() {}

func (x *PatchProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_profile_rpc_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchProfileReq.ProtoReflect.Descriptor instead.
func (*PatchProfileReq) Descriptor() ([]byte, []int) {
	return file_apps_profile_rpc_profile_proto_rawDescGZIP(), []int{2}
}

func (x *PatchProfileReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PatchProfileReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PatchProfileReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PatchProfileReq) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *PatchProfileReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DeleteProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DeleteProfileReq) Reset() {
	*x = DeleteProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_profile_rpc_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileReq) ProtoMessage() {}

func (x *DeleteProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_profile_rpc_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileReq.ProtoReflect.Descriptor instead.
func (*DeleteProfileReq) Descriptor() ([]byte, []int) {
	return file_apps_profile_rpc_profile_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteProfileReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmptyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty *emptypb.Empty `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *EmptyRes) Reset() {
	*x = EmptyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_profile_rpc_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRes) ProtoMessage() {}

func (x *EmptyRes) ProtoReflect() protoreflect.Message {
	mi := &file_apps_profile_rpc_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRes.ProtoReflect.Descriptor instead.
func (*EmptyRes) Descriptor() ([]byte, []int) {
	return file_apps_profile_rpc_profile_proto_rawDescGZIP(), []int{4}
}

func (x *EmptyRes) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

var File_apps_profile_rpc_profile_proto protoreflect.FileDescriptor

var file_apps_profile_rpc_profile_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x50, 0x43, 0x1a, 0x0e, 0x70, 0x62,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x8b, 0x01, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x7f, 0x0a, 0x0f, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x08,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xfb, 0x01, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x6c,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x50, 0x43, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x50, 0x43,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0c,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x50, 0x43, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x50, 0x43, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x12,
	0x43, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x50, 0x43, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x50, 0x43, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_profile_rpc_profile_proto_rawDescOnce sync.Once
	file_apps_profile_rpc_profile_proto_rawDescData = file_apps_profile_rpc_profile_proto_rawDesc
)

func file_apps_profile_rpc_profile_proto_rawDescGZIP() []byte {
	file_apps_profile_rpc_profile_proto_rawDescOnce.Do(func() {
		file_apps_profile_rpc_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_profile_rpc_profile_proto_rawDescData)
	})
	return file_apps_profile_rpc_profile_proto_rawDescData
}

var file_apps_profile_rpc_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_profile_rpc_profile_proto_goTypes = []interface{}{
	(*GetUniqueProfileOrCreateReq)(nil), // 0: profileRPC.GetUniqueProfileOrCreateReq
	(*GetUniqueProfileOrCreateRes)(nil), // 1: profileRPC.GetUniqueProfileOrCreateRes
	(*PatchProfileReq)(nil),             // 2: profileRPC.PatchProfileReq
	(*DeleteProfileReq)(nil),            // 3: profileRPC.DeleteProfileReq
	(*EmptyRes)(nil),                    // 4: profileRPC.EmptyRes
	(*emptypb.Empty)(nil),               // 5: google.protobuf.Empty
}
var file_apps_profile_rpc_profile_proto_depIdxs = []int32{
	5, // 0: profileRPC.EmptyRes.empty:type_name -> google.protobuf.Empty
	0, // 1: profileRPC.Rpc.GetUniqueProfileOrCreate:input_type -> profileRPC.GetUniqueProfileOrCreateReq
	2, // 2: profileRPC.Rpc.PatchProfile:input_type -> profileRPC.PatchProfileReq
	3, // 3: profileRPC.Rpc.DeleteProfile:input_type -> profileRPC.DeleteProfileReq
	1, // 4: profileRPC.Rpc.GetUniqueProfileOrCreate:output_type -> profileRPC.GetUniqueProfileOrCreateRes
	4, // 5: profileRPC.Rpc.PatchProfile:output_type -> profileRPC.EmptyRes
	4, // 6: profileRPC.Rpc.DeleteProfile:output_type -> profileRPC.EmptyRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_profile_rpc_profile_proto_init() }
func file_apps_profile_rpc_profile_proto_init() {
	if File_apps_profile_rpc_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_profile_rpc_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUniqueProfileOrCreateReq); i {
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
		file_apps_profile_rpc_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUniqueProfileOrCreateRes); i {
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
		file_apps_profile_rpc_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchProfileReq); i {
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
		file_apps_profile_rpc_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProfileReq); i {
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
		file_apps_profile_rpc_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRes); i {
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
			RawDescriptor: file_apps_profile_rpc_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_profile_rpc_profile_proto_goTypes,
		DependencyIndexes: file_apps_profile_rpc_profile_proto_depIdxs,
		MessageInfos:      file_apps_profile_rpc_profile_proto_msgTypes,
	}.Build()
	File_apps_profile_rpc_profile_proto = out.File
	file_apps_profile_rpc_profile_proto_rawDesc = nil
	file_apps_profile_rpc_profile_proto_goTypes = nil
	file_apps_profile_rpc_profile_proto_depIdxs = nil
}
