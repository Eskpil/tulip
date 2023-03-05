// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/api/api.proto

package api

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

type GenericEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Driver         string `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	DeviceId       string `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	EntityMetadata []byte `protobuf:"bytes,4,opt,name=entity_metadata,json=entityMetadata,proto3" json:"entity_metadata,omitempty"`
	DriverMetadata []byte `protobuf:"bytes,5,opt,name=driver_metadata,json=driverMetadata,proto3" json:"driver_metadata,omitempty"`
	Name           string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Kind           string `protobuf:"bytes,7,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *GenericEntityRequest) Reset() {
	*x = GenericEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericEntityRequest) ProtoMessage() {}

func (x *GenericEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericEntityRequest.ProtoReflect.Descriptor instead.
func (*GenericEntityRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *GenericEntityRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenericEntityRequest) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *GenericEntityRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *GenericEntityRequest) GetEntityMetadata() []byte {
	if x != nil {
		return x.EntityMetadata
	}
	return nil
}

func (x *GenericEntityRequest) GetDriverMetadata() []byte {
	if x != nil {
		return x.DriverMetadata
	}
	return nil
}

func (x *GenericEntityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenericEntityRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type CreateEntityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CreateEntityResponse) Reset() {
	*x = CreateEntityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntityResponse) ProtoMessage() {}

func (x *CreateEntityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntityResponse.ProtoReflect.Descriptor instead.
func (*CreateEntityResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEntityResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type UpdateEntityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *UpdateEntityResponse) Reset() {
	*x = UpdateEntityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityResponse) ProtoMessage() {}

func (x *UpdateEntityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityResponse.ProtoReflect.Descriptor instead.
func (*UpdateEntityResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateEntityResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type EntityExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EntityExistsRequest) Reset() {
	*x = EntityExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityExistsRequest) ProtoMessage() {}

func (x *EntityExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityExistsRequest.ProtoReflect.Descriptor instead.
func (*EntityExistsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *EntityExistsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EntityExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *EntityExistsResponse) Reset() {
	*x = EntityExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityExistsResponse) ProtoMessage() {}

func (x *EntityExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityExistsResponse.ProtoReflect.Descriptor instead.
func (*EntityExistsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *EntityExistsResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type AppendEntityHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId string `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	State    string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *AppendEntityHistoryRequest) Reset() {
	*x = AppendEntityHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntityHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntityHistoryRequest) ProtoMessage() {}

func (x *AppendEntityHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntityHistoryRequest.ProtoReflect.Descriptor instead.
func (*AppendEntityHistoryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *AppendEntityHistoryRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *AppendEntityHistoryRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type AppendEntityHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *AppendEntityHistoryResponse) Reset() {
	*x = AppendEntityHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntityHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntityHistoryResponse) ProtoMessage() {}

func (x *AppendEntityHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntityHistoryResponse.ProtoReflect.Descriptor instead.
func (*AppendEntityHistoryResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *AppendEntityHistoryResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_pkg_api_api_proto protoreflect.FileDescriptor

var file_pkg_api_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x22, 0x26, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x25, 0x0a, 0x13, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x4f, 0x0a, 0x1a, 0x41, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2d, 0x0a, 0x1b, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0x90, 0x02, 0x0a, 0x03, 0x61,
	0x70, 0x69, 0x12, 0x3b, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x12, 0x14, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3c, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x15, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x15, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x13, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x1b, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x6b, 0x70,
	0x69, 0x6c, 0x2f, 0x74, 0x75, 0x6c, 0x69, 0x70, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_api_proto_rawDescOnce sync.Once
	file_pkg_api_api_proto_rawDescData = file_pkg_api_api_proto_rawDesc
)

func file_pkg_api_api_proto_rawDescGZIP() []byte {
	file_pkg_api_api_proto_rawDescOnce.Do(func() {
		file_pkg_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_api_proto_rawDescData)
	})
	return file_pkg_api_api_proto_rawDescData
}

var file_pkg_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_api_api_proto_goTypes = []interface{}{
	(*GenericEntityRequest)(nil),        // 0: GenericEntityRequest
	(*CreateEntityResponse)(nil),        // 1: CreateEntityResponse
	(*UpdateEntityResponse)(nil),        // 2: UpdateEntityResponse
	(*EntityExistsRequest)(nil),         // 3: EntityExistsRequest
	(*EntityExistsResponse)(nil),        // 4: EntityExistsResponse
	(*AppendEntityHistoryRequest)(nil),  // 5: AppendEntityHistoryRequest
	(*AppendEntityHistoryResponse)(nil), // 6: AppendEntityHistoryResponse
}
var file_pkg_api_api_proto_depIdxs = []int32{
	3, // 0: api.EntityExists:input_type -> EntityExistsRequest
	0, // 1: api.CreateEntity:input_type -> GenericEntityRequest
	0, // 2: api.UpdateEntity:input_type -> GenericEntityRequest
	5, // 3: api.AppendEntityHistory:input_type -> AppendEntityHistoryRequest
	4, // 4: api.EntityExists:output_type -> EntityExistsResponse
	1, // 5: api.CreateEntity:output_type -> CreateEntityResponse
	2, // 6: api.UpdateEntity:output_type -> UpdateEntityResponse
	6, // 7: api.AppendEntityHistory:output_type -> AppendEntityHistoryResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_api_proto_init() }
func file_pkg_api_api_proto_init() {
	if File_pkg_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericEntityRequest); i {
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
		file_pkg_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntityResponse); i {
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
		file_pkg_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityResponse); i {
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
		file_pkg_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityExistsRequest); i {
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
		file_pkg_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityExistsResponse); i {
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
		file_pkg_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntityHistoryRequest); i {
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
		file_pkg_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntityHistoryResponse); i {
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
			RawDescriptor: file_pkg_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_api_proto_goTypes,
		DependencyIndexes: file_pkg_api_api_proto_depIdxs,
		MessageInfos:      file_pkg_api_api_proto_msgTypes,
	}.Build()
	File_pkg_api_api_proto = out.File
	file_pkg_api_api_proto_rawDesc = nil
	file_pkg_api_api_proto_goTypes = nil
	file_pkg_api_api_proto_depIdxs = nil
}