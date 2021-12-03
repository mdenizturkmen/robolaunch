// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: launch/launch.proto

package launchpb

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password     string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email        string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Organization string `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_launch_launch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_launch_launch_proto_msgTypes[0]
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
	return file_launch_launch_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

type UserCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_launch_launch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_launch_launch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_launch_launch_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreateRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserDeleteRequest) Reset() {
	*x = UserDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_launch_launch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteRequest) ProtoMessage() {}

func (x *UserDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_launch_launch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteRequest.ProtoReflect.Descriptor instead.
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return file_launch_launch_proto_rawDescGZIP(), []int{2}
}

func (x *UserDeleteRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk bool  `protobuf:"varint,1,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
	User *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_launch_launch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_launch_launch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_launch_launch_proto_rawDescGZIP(), []int{3}
}

func (x *UserResponse) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *UserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type Launch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RobotType string `protobuf:"bytes,3,opt,name=robot_type,json=robotType,proto3" json:"robot_type,omitempty"`
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Launch) Reset() {
	*x = Launch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_launch_launch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Launch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Launch) ProtoMessage() {}

func (x *Launch) ProtoReflect() protoreflect.Message {
	mi := &file_launch_launch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Launch.ProtoReflect.Descriptor instead.
func (*Launch) Descriptor() ([]byte, []int) {
	return file_launch_launch_proto_rawDescGZIP(), []int{4}
}

func (x *Launch) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Launch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Launch) GetRobotType() string {
	if x != nil {
		return x.RobotType
	}
	return ""
}

func (x *Launch) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type LaunchCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Launch *Launch `protobuf:"bytes,1,opt,name=launch,proto3" json:"launch,omitempty"`
}

func (x *LaunchCreateRequest) Reset() {
	*x = LaunchCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_launch_launch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchCreateRequest) ProtoMessage() {}

func (x *LaunchCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_launch_launch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchCreateRequest.ProtoReflect.Descriptor instead.
func (*LaunchCreateRequest) Descriptor() ([]byte, []int) {
	return file_launch_launch_proto_rawDescGZIP(), []int{5}
}

func (x *LaunchCreateRequest) GetLaunch() *Launch {
	if x != nil {
		return x.Launch
	}
	return nil
}

type LaunchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *LaunchDeleteRequest) Reset() {
	*x = LaunchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_launch_launch_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchDeleteRequest) ProtoMessage() {}

func (x *LaunchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_launch_launch_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchDeleteRequest.ProtoReflect.Descriptor instead.
func (*LaunchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_launch_launch_proto_rawDescGZIP(), []int{6}
}

func (x *LaunchDeleteRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LaunchDeleteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LaunchDeleteRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type LaunchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk   bool    `protobuf:"varint,1,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
	Launch *Launch `protobuf:"bytes,2,opt,name=launch,proto3" json:"launch,omitempty"`
}

func (x *LaunchResponse) Reset() {
	*x = LaunchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_launch_launch_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchResponse) ProtoMessage() {}

func (x *LaunchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_launch_launch_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchResponse.ProtoReflect.Descriptor instead.
func (*LaunchResponse) Descriptor() ([]byte, []int) {
	return file_launch_launch_proto_rawDescGZIP(), []int{7}
}

func (x *LaunchResponse) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *LaunchResponse) GetLaunch() *Launch {
	if x != nil {
		return x.Launch
	}
	return nil
}

var File_launch_launch_proto protoreflect.FileDescriptor

var file_launch_launch_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x22, 0x78, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2f,
	0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x45, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x69, 0x73, 0x4f, 0x6b, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x75, 0x0a, 0x06, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x3d, 0x0a,
	0x13, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x22, 0x63, 0x0a, 0x13,
	0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x4d, 0x0a, 0x0e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x12, 0x26, 0x0a, 0x06, 0x6c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x75, 0x6e, 0x63,
	0x68, 0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x32, 0x97, 0x02, 0x0a, 0x0d, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x19, 0x2e, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x19, 0x2e, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x12, 0x1b, 0x2e, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x61, 0x75, 0x6e, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c,
	0x61, 0x75, 0x6e, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x4c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x6c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_launch_launch_proto_rawDescOnce sync.Once
	file_launch_launch_proto_rawDescData = file_launch_launch_proto_rawDesc
)

func file_launch_launch_proto_rawDescGZIP() []byte {
	file_launch_launch_proto_rawDescOnce.Do(func() {
		file_launch_launch_proto_rawDescData = protoimpl.X.CompressGZIP(file_launch_launch_proto_rawDescData)
	})
	return file_launch_launch_proto_rawDescData
}

var file_launch_launch_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_launch_launch_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: launch.User
	(*UserCreateRequest)(nil),   // 1: launch.UserCreateRequest
	(*UserDeleteRequest)(nil),   // 2: launch.UserDeleteRequest
	(*UserResponse)(nil),        // 3: launch.UserResponse
	(*Launch)(nil),              // 4: launch.Launch
	(*LaunchCreateRequest)(nil), // 5: launch.LaunchCreateRequest
	(*LaunchDeleteRequest)(nil), // 6: launch.LaunchDeleteRequest
	(*LaunchResponse)(nil),      // 7: launch.LaunchResponse
}
var file_launch_launch_proto_depIdxs = []int32{
	0, // 0: launch.UserCreateRequest.user:type_name -> launch.User
	0, // 1: launch.UserResponse.user:type_name -> launch.User
	4, // 2: launch.LaunchCreateRequest.launch:type_name -> launch.Launch
	4, // 3: launch.LaunchResponse.launch:type_name -> launch.Launch
	1, // 4: launch.LaunchService.CreateUser:input_type -> launch.UserCreateRequest
	2, // 5: launch.LaunchService.DeleteUser:input_type -> launch.UserDeleteRequest
	5, // 6: launch.LaunchService.CreateLaunch:input_type -> launch.LaunchCreateRequest
	6, // 7: launch.LaunchService.DeleteLaunch:input_type -> launch.LaunchDeleteRequest
	3, // 8: launch.LaunchService.CreateUser:output_type -> launch.UserResponse
	3, // 9: launch.LaunchService.DeleteUser:output_type -> launch.UserResponse
	7, // 10: launch.LaunchService.CreateLaunch:output_type -> launch.LaunchResponse
	7, // 11: launch.LaunchService.DeleteLaunch:output_type -> launch.LaunchResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_launch_launch_proto_init() }
func file_launch_launch_proto_init() {
	if File_launch_launch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_launch_launch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_launch_launch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateRequest); i {
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
		file_launch_launch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteRequest); i {
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
		file_launch_launch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_launch_launch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Launch); i {
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
		file_launch_launch_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchCreateRequest); i {
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
		file_launch_launch_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchDeleteRequest); i {
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
		file_launch_launch_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchResponse); i {
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
			RawDescriptor: file_launch_launch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_launch_launch_proto_goTypes,
		DependencyIndexes: file_launch_launch_proto_depIdxs,
		MessageInfos:      file_launch_launch_proto_msgTypes,
	}.Build()
	File_launch_launch_proto = out.File
	file_launch_launch_proto_rawDesc = nil
	file_launch_launch_proto_goTypes = nil
	file_launch_launch_proto_depIdxs = nil
}
