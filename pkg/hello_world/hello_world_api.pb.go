// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: api/grpc/hello_world_api.proto

package hello_world

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Version struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Version) Reset() {
	*x = Version{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{0}
}

type Value struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Value) Reset() {
	*x = Value{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Value) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Values struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []*Value               `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Values) Reset() {
	*x = Values{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Values) ProtoMessage() {}

func (x *Values) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Values.ProtoReflect.Descriptor instead.
func (*Values) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{2}
}

func (x *Values) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type Get struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Get) Reset() {
	*x = Get{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get) ProtoMessage() {}

func (x *Get) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get.ProtoReflect.Descriptor instead.
func (*Get) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{3}
}

type Set struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Set) Reset() {
	*x = Set{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{4}
}

type Version_Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Version_Request) Reset() {
	*x = Version_Request{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Version_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version_Request) ProtoMessage() {}

func (x *Version_Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version_Request.ProtoReflect.Descriptor instead.
func (*Version_Request) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{0, 0}
}

type Version_Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Version_Response) Reset() {
	*x = Version_Response{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Version_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version_Response) ProtoMessage() {}

func (x *Version_Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version_Response.ProtoReflect.Descriptor instead.
func (*Version_Response) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Version_Response) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Get_Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Get_Request) Reset() {
	*x = Get_Request{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Get_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get_Request) ProtoMessage() {}

func (x *Get_Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get_Request.ProtoReflect.Descriptor instead.
func (*Get_Request) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Get_Request) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Set_Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Vals          *Values                `protobuf:"bytes,1,opt,name=vals,proto3" json:"vals,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Set_Request) Reset() {
	*x = Set_Request{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Set_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set_Request) ProtoMessage() {}

func (x *Set_Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set_Request.ProtoReflect.Descriptor instead.
func (*Set_Request) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Set_Request) GetVals() *Values {
	if x != nil {
		return x.Vals
	}
	return nil
}

type Set_Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Set_Response) Reset() {
	*x = Set_Response{}
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Set_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set_Response) ProtoMessage() {}

func (x *Set_Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_hello_world_api_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set_Response.ProtoReflect.Descriptor instead.
func (*Set_Response) Descriptor() ([]byte, []int) {
	return file_api_grpc_hello_world_api_proto_rawDescGZIP(), []int{4, 1}
}

var File_api_grpc_hello_world_api_proto protoreflect.FileDescriptor

var file_api_grpc_hello_world_api_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x61, 0x70,
	0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3a, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2f, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x06,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x22, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x1a, 0x1b, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x49, 0x0a, 0x03, 0x53, 0x65,
	0x74, 0x1a, 0x36, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04,
	0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x1a, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa3, 0x02, 0x0a, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x41, 0x50, 0x49, 0x12, 0x64, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x6b, 0x65,
	0x79, 0x7d, 0x12, 0x57, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01,
	0x2a, 0x22, 0x08, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_api_grpc_hello_world_api_proto_rawDescOnce sync.Once
	file_api_grpc_hello_world_api_proto_rawDescData []byte
)

func file_api_grpc_hello_world_api_proto_rawDescGZIP() []byte {
	file_api_grpc_hello_world_api_proto_rawDescOnce.Do(func() {
		file_api_grpc_hello_world_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_grpc_hello_world_api_proto_rawDesc), len(file_api_grpc_hello_world_api_proto_rawDesc)))
	})
	return file_api_grpc_hello_world_api_proto_rawDescData
}

var file_api_grpc_hello_world_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_grpc_hello_world_api_proto_goTypes = []any{
	(*Version)(nil),          // 0: hello_world_api.Version
	(*Value)(nil),            // 1: hello_world_api.Value
	(*Values)(nil),           // 2: hello_world_api.Values
	(*Get)(nil),              // 3: hello_world_api.Get
	(*Set)(nil),              // 4: hello_world_api.Set
	(*Version_Request)(nil),  // 5: hello_world_api.Version.Request
	(*Version_Response)(nil), // 6: hello_world_api.Version.Response
	(*Get_Request)(nil),      // 7: hello_world_api.Get.Request
	(*Set_Request)(nil),      // 8: hello_world_api.Set.Request
	(*Set_Response)(nil),     // 9: hello_world_api.Set.Response
}
var file_api_grpc_hello_world_api_proto_depIdxs = []int32{
	1, // 0: hello_world_api.Values.values:type_name -> hello_world_api.Value
	2, // 1: hello_world_api.Set.Request.vals:type_name -> hello_world_api.Values
	5, // 2: hello_world_api.helloWorldAPI.Version:input_type -> hello_world_api.Version.Request
	7, // 3: hello_world_api.helloWorldAPI.Get:input_type -> hello_world_api.Get.Request
	8, // 4: hello_world_api.helloWorldAPI.Set:input_type -> hello_world_api.Set.Request
	6, // 5: hello_world_api.helloWorldAPI.Version:output_type -> hello_world_api.Version.Response
	1, // 6: hello_world_api.helloWorldAPI.Get:output_type -> hello_world_api.Value
	9, // 7: hello_world_api.helloWorldAPI.Set:output_type -> hello_world_api.Set.Response
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_grpc_hello_world_api_proto_init() }
func file_api_grpc_hello_world_api_proto_init() {
	if File_api_grpc_hello_world_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_grpc_hello_world_api_proto_rawDesc), len(file_api_grpc_hello_world_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_hello_world_api_proto_goTypes,
		DependencyIndexes: file_api_grpc_hello_world_api_proto_depIdxs,
		MessageInfos:      file_api_grpc_hello_world_api_proto_msgTypes,
	}.Build()
	File_api_grpc_hello_world_api_proto = out.File
	file_api_grpc_hello_world_api_proto_goTypes = nil
	file_api_grpc_hello_world_api_proto_depIdxs = nil
}
