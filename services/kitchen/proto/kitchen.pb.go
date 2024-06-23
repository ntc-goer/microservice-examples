// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: kitchen.proto

package kitchenproto

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

type DishItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DishId string `protobuf:"bytes,1,opt,name=dish_id,json=dishId,proto3" json:"dish_id,omitempty"`
	Dish   string `protobuf:"bytes,2,opt,name=dish,proto3" json:"dish,omitempty"`
	Total  int32  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *DishItem) Reset() {
	*x = DishItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DishItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DishItem) ProtoMessage() {}

func (x *DishItem) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DishItem.ProtoReflect.Descriptor instead.
func (*DishItem) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{0}
}

func (x *DishItem) GetDishId() string {
	if x != nil {
		return x.DishId
	}
	return ""
}

func (x *DishItem) GetDish() string {
	if x != nil {
		return x.Dish
	}
	return ""
}

func (x *DishItem) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type VerifyOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId string      `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Dishes  []*DishItem `protobuf:"bytes,2,rep,name=dishes,proto3" json:"dishes,omitempty"`
}

func (x *VerifyOrderRequest) Reset() {
	*x = VerifyOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOrderRequest) ProtoMessage() {}

func (x *VerifyOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOrderRequest.ProtoReflect.Descriptor instead.
func (*VerifyOrderRequest) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyOrderRequest) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *VerifyOrderRequest) GetDishes() []*DishItem {
	if x != nil {
		return x.Dishes
	}
	return nil
}

type VerifyOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk bool `protobuf:"varint,1,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
}

func (x *VerifyOrderResponse) Reset() {
	*x = VerifyOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOrderResponse) ProtoMessage() {}

func (x *VerifyOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOrderResponse.ProtoReflect.Descriptor instead.
func (*VerifyOrderResponse) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyOrderResponse) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

var File_kitchen_proto protoreflect.FileDescriptor

var file_kitchen_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x22, 0x4d, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x68,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x73, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x69, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x69, 0x73,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x5a, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x64, 0x69, 0x73, 0x68,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x64, 0x69, 0x73,
	0x68, 0x65, 0x73, 0x22, 0x2a, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73,
	0x5f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x32,
	0x5a, 0x0a, 0x0e, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x48, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1b, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x63, 0x2d, 0x67, 0x6f,
	0x65, 0x72, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kitchen_proto_rawDescOnce sync.Once
	file_kitchen_proto_rawDescData = file_kitchen_proto_rawDesc
)

func file_kitchen_proto_rawDescGZIP() []byte {
	file_kitchen_proto_rawDescOnce.Do(func() {
		file_kitchen_proto_rawDescData = protoimpl.X.CompressGZIP(file_kitchen_proto_rawDescData)
	})
	return file_kitchen_proto_rawDescData
}

var file_kitchen_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kitchen_proto_goTypes = []any{
	(*DishItem)(nil),            // 0: kitchen.DishItem
	(*VerifyOrderRequest)(nil),  // 1: kitchen.VerifyOrderRequest
	(*VerifyOrderResponse)(nil), // 2: kitchen.VerifyOrderResponse
}
var file_kitchen_proto_depIdxs = []int32{
	0, // 0: kitchen.VerifyOrderRequest.dishes:type_name -> kitchen.DishItem
	1, // 1: kitchen.KitchenService.VerifyOrder:input_type -> kitchen.VerifyOrderRequest
	2, // 2: kitchen.KitchenService.VerifyOrder:output_type -> kitchen.VerifyOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kitchen_proto_init() }
func file_kitchen_proto_init() {
	if File_kitchen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kitchen_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DishItem); i {
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
		file_kitchen_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyOrderRequest); i {
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
		file_kitchen_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyOrderResponse); i {
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
			RawDescriptor: file_kitchen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kitchen_proto_goTypes,
		DependencyIndexes: file_kitchen_proto_depIdxs,
		MessageInfos:      file_kitchen_proto_msgTypes,
	}.Build()
	File_kitchen_proto = out.File
	file_kitchen_proto_rawDesc = nil
	file_kitchen_proto_goTypes = nil
	file_kitchen_proto_depIdxs = nil
}
