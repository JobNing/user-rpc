// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.3
// source: pb/role/role.proto

package role

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

type RoleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int64  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,20,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *RoleInfo) Reset() {
	*x = RoleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_role_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleInfo) ProtoMessage() {}

func (x *RoleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_role_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleInfo.ProtoReflect.Descriptor instead.
func (*RoleInfo) Descriptor() ([]byte, []int) {
	return file_pb_role_role_proto_rawDescGZIP(), []int{0}
}

func (x *RoleInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RoleInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In *RoleInfo `protobuf:"bytes,10,opt,name=In,proto3" json:"In,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_role_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_role_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_pb_role_role_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoleRequest) GetIn() *RoleInfo {
	if x != nil {
		return x.In
	}
	return nil
}

type CreateRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *RoleInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateRoleResponse) Reset() {
	*x = CreateRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_role_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleResponse) ProtoMessage() {}

func (x *CreateRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_role_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleResponse.ProtoReflect.Descriptor instead.
func (*CreateRoleResponse) Descriptor() ([]byte, []int) {
	return file_pb_role_role_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoleResponse) GetInfo() *RoleInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In *RoleInfo `protobuf:"bytes,10,opt,name=In,proto3" json:"In,omitempty"`
}

func (x *UpdateRoleRequest) Reset() {
	*x = UpdateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_role_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRequest) ProtoMessage() {}

func (x *UpdateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_role_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return file_pb_role_role_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRoleRequest) GetIn() *RoleInfo {
	if x != nil {
		return x.In
	}
	return nil
}

type UpdateRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *RoleInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateRoleResponse) Reset() {
	*x = UpdateRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_role_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleResponse) ProtoMessage() {}

func (x *UpdateRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_role_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleResponse.ProtoReflect.Descriptor instead.
func (*UpdateRoleResponse) Descriptor() ([]byte, []int) {
	return file_pb_role_role_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRoleResponse) GetInfo() *RoleInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetRoleRequest) Reset() {
	*x = GetRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_role_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleRequest) ProtoMessage() {}

func (x *GetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_role_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleRequest.ProtoReflect.Descriptor instead.
func (*GetRoleRequest) Descriptor() ([]byte, []int) {
	return file_pb_role_role_proto_rawDescGZIP(), []int{5}
}

func (x *GetRoleRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *RoleInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetRoleResponse) Reset() {
	*x = GetRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_role_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleResponse) ProtoMessage() {}

func (x *GetRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_role_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleResponse.ProtoReflect.Descriptor instead.
func (*GetRoleResponse) Descriptor() ([]byte, []int) {
	return file_pb_role_role_proto_rawDescGZIP(), []int{6}
}

func (x *GetRoleResponse) GetInfo() *RoleInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type SearchRoleByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,10,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *SearchRoleByNameRequest) Reset() {
	*x = SearchRoleByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_role_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRoleByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRoleByNameRequest) ProtoMessage() {}

func (x *SearchRoleByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_role_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRoleByNameRequest.ProtoReflect.Descriptor instead.
func (*SearchRoleByNameRequest) Descriptor() ([]byte, []int) {
	return file_pb_role_role_proto_rawDescGZIP(), []int{7}
}

func (x *SearchRoleByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SearchRoleByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*RoleInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *SearchRoleByNameResponse) Reset() {
	*x = SearchRoleByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_role_role_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRoleByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRoleByNameResponse) ProtoMessage() {}

func (x *SearchRoleByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_role_role_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRoleByNameResponse.ProtoReflect.Descriptor instead.
func (*SearchRoleByNameResponse) Descriptor() ([]byte, []int) {
	return file_pb_role_role_proto_rawDescGZIP(), []int{8}
}

func (x *SearchRoleByNameResponse) GetInfos() []*RoleInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_pb_role_role_proto protoreflect.FileDescriptor

var file_pb_role_role_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x2e, 0x0a, 0x08, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x02, 0x49, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x02, 0x49, 0x6e, 0x22,
	0x38, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x33, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x02, 0x49, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x02, 0x49, 0x6e, 0x22, 0x38,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x35, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x2d, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x40, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x32, 0x9b, 0x02, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x6f, 0x6c,
	0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_role_role_proto_rawDescOnce sync.Once
	file_pb_role_role_proto_rawDescData = file_pb_role_role_proto_rawDesc
)

func file_pb_role_role_proto_rawDescGZIP() []byte {
	file_pb_role_role_proto_rawDescOnce.Do(func() {
		file_pb_role_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_role_role_proto_rawDescData)
	})
	return file_pb_role_role_proto_rawDescData
}

var file_pb_role_role_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_role_role_proto_goTypes = []interface{}{
	(*RoleInfo)(nil),                 // 0: role.RoleInfo
	(*CreateRoleRequest)(nil),        // 1: role.CreateRoleRequest
	(*CreateRoleResponse)(nil),       // 2: role.CreateRoleResponse
	(*UpdateRoleRequest)(nil),        // 3: role.UpdateRoleRequest
	(*UpdateRoleResponse)(nil),       // 4: role.UpdateRoleResponse
	(*GetRoleRequest)(nil),           // 5: role.GetRoleRequest
	(*GetRoleResponse)(nil),          // 6: role.GetRoleResponse
	(*SearchRoleByNameRequest)(nil),  // 7: role.SearchRoleByNameRequest
	(*SearchRoleByNameResponse)(nil), // 8: role.SearchRoleByNameResponse
}
var file_pb_role_role_proto_depIdxs = []int32{
	0,  // 0: role.CreateRoleRequest.In:type_name -> role.RoleInfo
	0,  // 1: role.CreateRoleResponse.Info:type_name -> role.RoleInfo
	0,  // 2: role.UpdateRoleRequest.In:type_name -> role.RoleInfo
	0,  // 3: role.UpdateRoleResponse.Info:type_name -> role.RoleInfo
	0,  // 4: role.GetRoleResponse.Info:type_name -> role.RoleInfo
	0,  // 5: role.SearchRoleByNameResponse.Infos:type_name -> role.RoleInfo
	1,  // 6: role.Role.CreateRole:input_type -> role.CreateRoleRequest
	3,  // 7: role.Role.UpdateRole:input_type -> role.UpdateRoleRequest
	5,  // 8: role.Role.GetRole:input_type -> role.GetRoleRequest
	7,  // 9: role.Role.SearchRoleByName:input_type -> role.SearchRoleByNameRequest
	2,  // 10: role.Role.CreateRole:output_type -> role.CreateRoleResponse
	4,  // 11: role.Role.UpdateRole:output_type -> role.UpdateRoleResponse
	6,  // 12: role.Role.GetRole:output_type -> role.GetRoleResponse
	8,  // 13: role.Role.SearchRoleByName:output_type -> role.SearchRoleByNameResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pb_role_role_proto_init() }
func file_pb_role_role_proto_init() {
	if File_pb_role_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_role_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleInfo); i {
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
		file_pb_role_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_pb_role_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleResponse); i {
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
		file_pb_role_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleRequest); i {
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
		file_pb_role_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleResponse); i {
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
		file_pb_role_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleRequest); i {
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
		file_pb_role_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleResponse); i {
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
		file_pb_role_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRoleByNameRequest); i {
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
		file_pb_role_role_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRoleByNameResponse); i {
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
			RawDescriptor: file_pb_role_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_role_role_proto_goTypes,
		DependencyIndexes: file_pb_role_role_proto_depIdxs,
		MessageInfos:      file_pb_role_role_proto_msgTypes,
	}.Build()
	File_pb_role_role_proto = out.File
	file_pb_role_role_proto_rawDesc = nil
	file_pb_role_role_proto_goTypes = nil
	file_pb_role_role_proto_depIdxs = nil
}
