// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/hello/v1/service.proto

package hellopb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWeatherRequest) Reset() {
	*x = GetWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherRequest) ProtoMessage() {}

func (x *GetWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherRequest.ProtoReflect.Descriptor instead.
func (*GetWeatherRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{0}
}

type GetWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Weather:
	//	*GetWeatherResponse_Raining
	//	*GetWeatherResponse_Cloudy_
	Weather isGetWeatherResponse_Weather `protobuf_oneof:"weather"`
}

func (x *GetWeatherResponse) Reset() {
	*x = GetWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherResponse) ProtoMessage() {}

func (x *GetWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherResponse.ProtoReflect.Descriptor instead.
func (*GetWeatherResponse) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{1}
}

func (m *GetWeatherResponse) GetWeather() isGetWeatherResponse_Weather {
	if m != nil {
		return m.Weather
	}
	return nil
}

func (x *GetWeatherResponse) GetRaining() *GetWeatherResponse_Rainning {
	if x, ok := x.GetWeather().(*GetWeatherResponse_Raining); ok {
		return x.Raining
	}
	return nil
}

func (x *GetWeatherResponse) GetCloudy() *GetWeatherResponse_Cloudy {
	if x, ok := x.GetWeather().(*GetWeatherResponse_Cloudy_); ok {
		return x.Cloudy
	}
	return nil
}

type isGetWeatherResponse_Weather interface {
	isGetWeatherResponse_Weather()
}

type GetWeatherResponse_Raining struct {
	Raining *GetWeatherResponse_Rainning `protobuf:"bytes,1,opt,name=raining,proto3,oneof"`
}

type GetWeatherResponse_Cloudy_ struct {
	Cloudy *GetWeatherResponse_Cloudy `protobuf:"bytes,2,opt,name=cloudy,proto3,oneof"`
}

func (*GetWeatherResponse_Raining) isGetWeatherResponse_Weather() {}

func (*GetWeatherResponse_Cloudy_) isGetWeatherResponse_Weather() {}

type GetTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTimeRequest) Reset() {
	*x = GetTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimeRequest) ProtoMessage() {}

func (x *GetTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimeRequest.ProtoReflect.Descriptor instead.
func (*GetTimeRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{2}
}

type GetTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Now *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=now,proto3" json:"now,omitempty"`
}

func (x *GetTimeResponse) Reset() {
	*x = GetTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimeResponse) ProtoMessage() {}

func (x *GetTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimeResponse.ProtoReflect.Descriptor instead.
func (*GetTimeResponse) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTimeResponse) GetNow() *timestamppb.Timestamp {
	if x != nil {
		return x.Now
	}
	return nil
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *HelloRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HelloMsg string `protobuf:"bytes,1,opt,name=hello_msg,json=helloMsg,proto3" json:"hello_msg,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *HelloResponse) GetHelloMsg() string {
	if x != nil {
		return x.HelloMsg
	}
	return ""
}

type HiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *HiRequest) Reset() {
	*x = HiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiRequest) ProtoMessage() {}

func (x *HiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiRequest.ProtoReflect.Descriptor instead.
func (*HiRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *HiRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type HiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HiMsg string `protobuf:"bytes,1,opt,name=hi_msg,json=hiMsg,proto3" json:"hi_msg,omitempty"`
}

func (x *HiResponse) Reset() {
	*x = HiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiResponse) ProtoMessage() {}

func (x *HiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiResponse.ProtoReflect.Descriptor instead.
func (*HiResponse) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *HiResponse) GetHiMsg() string {
	if x != nil {
		return x.HiMsg
	}
	return ""
}

type GetWeatherResponse_Rainning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWeatherResponse_Rainning) Reset() {
	*x = GetWeatherResponse_Rainning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherResponse_Rainning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherResponse_Rainning) ProtoMessage() {}

func (x *GetWeatherResponse_Rainning) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherResponse_Rainning.ProtoReflect.Descriptor instead.
func (*GetWeatherResponse_Rainning) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{1, 0}
}

type GetWeatherResponse_Cloudy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWeatherResponse_Cloudy) Reset() {
	*x = GetWeatherResponse_Cloudy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherResponse_Cloudy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherResponse_Cloudy) ProtoMessage() {}

func (x *GetWeatherResponse_Cloudy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherResponse_Cloudy.ProtoReflect.Descriptor instead.
func (*GetWeatherResponse_Cloudy) Descriptor() ([]byte, []int) {
	return file_proto_hello_v1_service_proto_rawDescGZIP(), []int{1, 1}
}

var File_proto_hello_v1_service_proto protoreflect.FileDescriptor

var file_proto_hello_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc3,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x61, 0x69, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x07, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x43,
	0x0a, 0x06, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x79, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x79, 0x1a, 0x0a, 0x0a, 0x08, 0x52, 0x61, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x1a,
	0x08, 0x0a, 0x06, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x6e, 0x6f, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x03, 0x6e, 0x6f, 0x77, 0x22, 0x38, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x2c, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x73, 0x67, 0x22,
	0x35, 0x0a, 0x09, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x0a, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x68, 0x69, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x69, 0x4d, 0x73, 0x67, 0x32, 0x91, 0x03, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x05,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x4e, 0x0a, 0x02, 0x48, 0x69, 0x12, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x69, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x68, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75,
	0x64, 0x75, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x67, 0x65, 0x6e, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_v1_service_proto_rawDescOnce sync.Once
	file_proto_hello_v1_service_proto_rawDescData = file_proto_hello_v1_service_proto_rawDesc
)

func file_proto_hello_v1_service_proto_rawDescGZIP() []byte {
	file_proto_hello_v1_service_proto_rawDescOnce.Do(func() {
		file_proto_hello_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_v1_service_proto_rawDescData)
	})
	return file_proto_hello_v1_service_proto_rawDescData
}

var file_proto_hello_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_hello_v1_service_proto_goTypes = []interface{}{
	(*GetWeatherRequest)(nil),           // 0: proto.hello.v1.GetWeatherRequest
	(*GetWeatherResponse)(nil),          // 1: proto.hello.v1.GetWeatherResponse
	(*GetTimeRequest)(nil),              // 2: proto.hello.v1.GetTimeRequest
	(*GetTimeResponse)(nil),             // 3: proto.hello.v1.GetTimeResponse
	(*HelloRequest)(nil),                // 4: proto.hello.v1.HelloRequest
	(*HelloResponse)(nil),               // 5: proto.hello.v1.HelloResponse
	(*HiRequest)(nil),                   // 6: proto.hello.v1.HiRequest
	(*HiResponse)(nil),                  // 7: proto.hello.v1.HiResponse
	(*GetWeatherResponse_Rainning)(nil), // 8: proto.hello.v1.GetWeatherResponse.Rainning
	(*GetWeatherResponse_Cloudy)(nil),   // 9: proto.hello.v1.GetWeatherResponse.Cloudy
	(*timestamppb.Timestamp)(nil),       // 10: google.protobuf.Timestamp
	(*User)(nil),                        // 11: proto.hello.v1.User
}
var file_proto_hello_v1_service_proto_depIdxs = []int32{
	8,  // 0: proto.hello.v1.GetWeatherResponse.raining:type_name -> proto.hello.v1.GetWeatherResponse.Rainning
	9,  // 1: proto.hello.v1.GetWeatherResponse.cloudy:type_name -> proto.hello.v1.GetWeatherResponse.Cloudy
	10, // 2: proto.hello.v1.GetTimeResponse.now:type_name -> google.protobuf.Timestamp
	11, // 3: proto.hello.v1.HelloRequest.user:type_name -> proto.hello.v1.User
	11, // 4: proto.hello.v1.HiRequest.user:type_name -> proto.hello.v1.User
	4,  // 5: proto.hello.v1.HelloService.Hello:input_type -> proto.hello.v1.HelloRequest
	6,  // 6: proto.hello.v1.HelloService.Hi:input_type -> proto.hello.v1.HiRequest
	2,  // 7: proto.hello.v1.HelloService.GetTime:input_type -> proto.hello.v1.GetTimeRequest
	0,  // 8: proto.hello.v1.HelloService.GetWeather:input_type -> proto.hello.v1.GetWeatherRequest
	5,  // 9: proto.hello.v1.HelloService.Hello:output_type -> proto.hello.v1.HelloResponse
	7,  // 10: proto.hello.v1.HelloService.Hi:output_type -> proto.hello.v1.HiResponse
	3,  // 11: proto.hello.v1.HelloService.GetTime:output_type -> proto.hello.v1.GetTimeResponse
	1,  // 12: proto.hello.v1.HelloService.GetWeather:output_type -> proto.hello.v1.GetWeatherResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_hello_v1_service_proto_init() }
func file_proto_hello_v1_service_proto_init() {
	if File_proto_hello_v1_service_proto != nil {
		return
	}
	file_proto_hello_v1_hello_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherRequest); i {
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
		file_proto_hello_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherResponse); i {
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
		file_proto_hello_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimeRequest); i {
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
		file_proto_hello_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimeResponse); i {
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
		file_proto_hello_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_proto_hello_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_proto_hello_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiRequest); i {
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
		file_proto_hello_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiResponse); i {
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
		file_proto_hello_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherResponse_Rainning); i {
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
		file_proto_hello_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherResponse_Cloudy); i {
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
	file_proto_hello_v1_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GetWeatherResponse_Raining)(nil),
		(*GetWeatherResponse_Cloudy_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_hello_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_v1_service_proto_goTypes,
		DependencyIndexes: file_proto_hello_v1_service_proto_depIdxs,
		MessageInfos:      file_proto_hello_v1_service_proto_msgTypes,
	}.Build()
	File_proto_hello_v1_service_proto = out.File
	file_proto_hello_v1_service_proto_rawDesc = nil
	file_proto_hello_v1_service_proto_goTypes = nil
	file_proto_hello_v1_service_proto_depIdxs = nil
}
