// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: fulfillment.proto

package fulfillment

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

	Username string    `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string    `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Location *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
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

type DeliveryPartnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string    `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Location *Location `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *DeliveryPartnerResponse) Reset() {
	*x = DeliveryPartnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryPartnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryPartnerResponse) ProtoMessage() {}

func (x *DeliveryPartnerResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeliveryPartnerResponse.ProtoReflect.Descriptor instead.
func (*DeliveryPartnerResponse) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryPartnerResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeliveryPartnerResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DeliveryPartnerResponse) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_fulfillment_proto protoreflect.FileDescriptor

var file_fulfillment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x78, 0x43, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x78, 0x43, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x79, 0x43, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x79, 0x43, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x22,
	0x7f, 0x0a, 0x1e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x6c, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x65,
	0x0a, 0x0b, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x56, 0x0a,
	0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_fulfillment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_fulfillment_proto_goTypes = []interface{}{
	(*Location)(nil),                       // 0: Location
	(*RegisterDeliveryPartnerRequest)(nil), // 1: RegisterDeliveryPartnerRequest
	(*DeliveryPartnerResponse)(nil),        // 2: DeliveryPartnerResponse
}
var file_fulfillment_proto_depIdxs = []int32{
	0, // 0: RegisterDeliveryPartnerRequest.location:type_name -> Location
	0, // 1: DeliveryPartnerResponse.location:type_name -> Location
	1, // 2: Fulfillment.RegisterDeliveryPartner:input_type -> RegisterDeliveryPartnerRequest
	2, // 3: Fulfillment.RegisterDeliveryPartner:output_type -> DeliveryPartnerResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
			switch v := v.(*DeliveryPartnerResponse); i {
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
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
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
