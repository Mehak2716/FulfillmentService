// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: fulfillment.proto

package fulfillment

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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XCordinate float32 `protobuf:"fixed32,1,opt,name=xCordinate,proto3" json:"xCordinate,omitempty"`
	YCordinate float32 `protobuf:"fixed32,2,opt,name=yCordinate,proto3" json:"yCordinate,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetXCordinate() float32 {
	if x != nil {
		return x.XCordinate
	}
	return 0
}

func (x *Location) GetYCordinate() float32 {
	if x != nil {
		return x.YCordinate
	}
	return 0
}

type RegisterDeliveryPartnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string    `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password     string    `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Location     *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Availability string    `protobuf:"bytes,4,opt,name=availability,proto3" json:"availability,omitempty"`
}

func (x *RegisterDeliveryPartnerRequest) Reset() {
	*x = RegisterDeliveryPartnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDeliveryPartnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeliveryPartnerRequest) ProtoMessage() {}

func (x *RegisterDeliveryPartnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeliveryPartnerRequest.ProtoReflect.Descriptor instead.
func (*RegisterDeliveryPartnerRequest) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterDeliveryPartnerRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterDeliveryPartnerRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterDeliveryPartnerRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *RegisterDeliveryPartnerRequest) GetAvailability() string {
	if x != nil {
		return x.Availability
	}
	return ""
}

type DeliveryPartner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username     string    `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Location     *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Availability string    `protobuf:"bytes,4,opt,name=availability,proto3" json:"availability,omitempty"`
}

func (x *DeliveryPartner) Reset() {
	*x = DeliveryPartner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryPartner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryPartner) ProtoMessage() {}

func (x *DeliveryPartner) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryPartner.ProtoReflect.Descriptor instead.
func (*DeliveryPartner) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryPartner) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeliveryPartner) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DeliveryPartner) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *DeliveryPartner) GetAvailability() string {
	if x != nil {
		return x.Availability
	}
	return ""
}

type DeliveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId          int64     `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	PickUpLocation   *Location `protobuf:"bytes,2,opt,name=pickUpLocation,proto3" json:"pickUpLocation,omitempty"`
	DeliveryLocation *Location `protobuf:"bytes,3,opt,name=deliveryLocation,proto3" json:"deliveryLocation,omitempty"`
}

func (x *DeliveryRequest) Reset() {
	*x = DeliveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryRequest) ProtoMessage() {}

func (x *DeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryRequest.ProtoReflect.Descriptor instead.
func (*DeliveryRequest) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{3}
}

func (x *DeliveryRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DeliveryRequest) GetPickUpLocation() *Location {
	if x != nil {
		return x.PickUpLocation
	}
	return nil
}

func (x *DeliveryRequest) GetDeliveryLocation() *Location {
	if x != nil {
		return x.DeliveryLocation
	}
	return nil
}

type DeliveredRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *DeliveredRequest) Reset() {
	*x = DeliveredRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveredRequest) ProtoMessage() {}

func (x *DeliveredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveredRequest.ProtoReflect.Descriptor instead.
func (*DeliveredRequest) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{4}
}

func (x *DeliveredRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeliveredResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId          int64                  `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	DeliveryParterId int64                  `protobuf:"varint,2,opt,name=deliveryParterId,proto3" json:"deliveryParterId,omitempty"`
	DeliveryLocation *Location              `protobuf:"bytes,3,opt,name=deliveryLocation,proto3" json:"deliveryLocation,omitempty"`
	DeliveredAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deliveredAt,proto3" json:"deliveredAt,omitempty"`
}

func (x *DeliveredResponse) Reset() {
	*x = DeliveredResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveredResponse) ProtoMessage() {}

func (x *DeliveredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveredResponse.ProtoReflect.Descriptor instead.
func (*DeliveredResponse) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{6}
}

func (x *DeliveredResponse) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DeliveredResponse) GetDeliveryParterId() int64 {
	if x != nil {
		return x.DeliveryParterId
	}
	return 0
}

func (x *DeliveredResponse) GetDeliveryLocation() *Location {
	if x != nil {
		return x.DeliveryLocation
	}
	return nil
}

func (x *DeliveredResponse) GetDeliveredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveredAt
	}
	return nil
}

var File_fulfillment_proto protoreflect.FileDescriptor

var file_fulfillment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x78, 0x43, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x78, 0x43, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x79, 0x43, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x79, 0x43, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x22, 0xa3, 0x01, 0x0a, 0x1e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x88, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xce, 0x01,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x50, 0x61, 0x72, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x10, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3c, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x41, 0x74, 0x32, 0x68,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x22, 0x00, 0x32, 0x7e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x10, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12,
	0x10, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x12,
	0x11, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x66, 0x75, 0x6c,
	0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fulfillment_proto_rawDescOnce sync.Once
	file_fulfillment_proto_rawDescData = file_fulfillment_proto_rawDesc
)

func file_fulfillment_proto_rawDescGZIP() []byte {
	file_fulfillment_proto_rawDescOnce.Do(func() {
		file_fulfillment_proto_rawDescData = protoimpl.X.CompressGZIP(file_fulfillment_proto_rawDescData)
	})
	return file_fulfillment_proto_rawDescData
}

var file_fulfillment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_fulfillment_proto_goTypes = []interface{}{
	(*Location)(nil),                       // 0: Location
	(*RegisterDeliveryPartnerRequest)(nil), // 1: RegisterDeliveryPartnerRequest
	(*DeliveryPartner)(nil),                // 2: DeliveryPartner
	(*DeliveryRequest)(nil),                // 3: DeliveryRequest
	(*DeliveredRequest)(nil),               // 4: DeliveredRequest
	(*Response)(nil),                       // 5: Response
	(*DeliveredResponse)(nil),              // 6: DeliveredResponse
	(*timestamppb.Timestamp)(nil),          // 7: google.protobuf.Timestamp
}
var file_fulfillment_proto_depIdxs = []int32{
	0, // 0: RegisterDeliveryPartnerRequest.location:type_name -> Location
	0, // 1: DeliveryPartner.location:type_name -> Location
	0, // 2: DeliveryRequest.pickUpLocation:type_name -> Location
	0, // 3: DeliveryRequest.deliveryLocation:type_name -> Location
	0, // 4: DeliveredResponse.deliveryLocation:type_name -> Location
	7, // 5: DeliveredResponse.deliveredAt:type_name -> google.protobuf.Timestamp
	1, // 6: DeliveryPartnerService.RegisterDeliveryPartner:input_type -> RegisterDeliveryPartnerRequest
	3, // 7: DeliveryService.InitiateDelivery:input_type -> DeliveryRequest
	4, // 8: DeliveryService.MarkDelivered:input_type -> DeliveredRequest
	2, // 9: DeliveryPartnerService.RegisterDeliveryPartner:output_type -> DeliveryPartner
	5, // 10: DeliveryService.InitiateDelivery:output_type -> Response
	6, // 11: DeliveryService.MarkDelivered:output_type -> DeliveredResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_fulfillment_proto_init() }
func file_fulfillment_proto_init() {
	if File_fulfillment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fulfillment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_fulfillment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDeliveryPartnerRequest); i {
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
		file_fulfillment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryPartner); i {
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
		file_fulfillment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryRequest); i {
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
		file_fulfillment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveredRequest); i {
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
		file_fulfillment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_fulfillment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveredResponse); i {
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
			RawDescriptor: file_fulfillment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_fulfillment_proto_goTypes,
		DependencyIndexes: file_fulfillment_proto_depIdxs,
		MessageInfos:      file_fulfillment_proto_msgTypes,
	}.Build()
	File_fulfillment_proto = out.File
	file_fulfillment_proto_rawDesc = nil
	file_fulfillment_proto_goTypes = nil
	file_fulfillment_proto_depIdxs = nil
}
