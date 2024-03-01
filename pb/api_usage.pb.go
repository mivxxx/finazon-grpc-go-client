// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: api_usage.proto

package finazon_grpc_go_client

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

type ApiUsageItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product            string    `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	ApiCalls           *ApiCalls `protobuf:"bytes,2,opt,name=apiCalls,proto3" json:"apiCalls,omitempty"`
	HistoricalApiCalls *ApiCalls `protobuf:"bytes,3,opt,name=historicalApiCalls,proto3" json:"historicalApiCalls,omitempty"`
}

func (x *ApiUsageItem) Reset() {
	*x = ApiUsageItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_usage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiUsageItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiUsageItem) ProtoMessage() {}

func (x *ApiUsageItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_usage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiUsageItem.ProtoReflect.Descriptor instead.
func (*ApiUsageItem) Descriptor() ([]byte, []int) {
	return file_api_usage_proto_rawDescGZIP(), []int{0}
}

func (x *ApiUsageItem) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *ApiUsageItem) GetApiCalls() *ApiCalls {
	if x != nil {
		return x.ApiCalls
	}
	return nil
}

func (x *ApiUsageItem) GetHistoricalApiCalls() *ApiCalls {
	if x != nil {
		return x.HistoricalApiCalls
	}
	return nil
}

type ApiCalls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Usage int32 `protobuf:"varint,2,opt,name=usage,proto3" json:"usage,omitempty"`
}

func (x *ApiCalls) Reset() {
	*x = ApiCalls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_usage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiCalls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiCalls) ProtoMessage() {}

func (x *ApiCalls) ProtoReflect() protoreflect.Message {
	mi := &file_api_usage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiCalls.ProtoReflect.Descriptor instead.
func (*ApiCalls) Descriptor() ([]byte, []int) {
	return file_api_usage_proto_rawDescGZIP(), []int{1}
}

func (x *ApiCalls) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ApiCalls) GetUsage() int32 {
	if x != nil {
		return x.Usage
	}
	return 0
}

type GetApiUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product string `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *GetApiUsageRequest) Reset() {
	*x = GetApiUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_usage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiUsageRequest) ProtoMessage() {}

func (x *GetApiUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_usage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiUsageRequest.ProtoReflect.Descriptor instead.
func (*GetApiUsageRequest) Descriptor() ([]byte, []int) {
	return file_api_usage_proto_rawDescGZIP(), []int{2}
}

func (x *GetApiUsageRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

type GetApiUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*ApiUsageItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetApiUsageResponse) Reset() {
	*x = GetApiUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_usage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiUsageResponse) ProtoMessage() {}

func (x *GetApiUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_usage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiUsageResponse.ProtoReflect.Descriptor instead.
func (*GetApiUsageResponse) Descriptor() ([]byte, []int) {
	return file_api_usage_proto_rawDescGZIP(), []int{3}
}

func (x *GetApiUsageResponse) GetResult() []*ApiUsageItem {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_api_usage_proto protoreflect.FileDescriptor

var file_api_usage_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x41,
	0x70, 0x69, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x43, 0x61, 0x6c, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x69, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x52, 0x08, 0x61, 0x70, 0x69, 0x43,
	0x61, 0x6c, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x12, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63,
	0x61, 0x6c, 0x41, 0x70, 0x69, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x43, 0x61,
	0x6c, 0x6c, 0x73, 0x52, 0x12, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x41,
	0x70, 0x69, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x22, 0x36, 0x0a, 0x08, 0x41, 0x70, 0x69, 0x43, 0x61,
	0x6c, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0x44, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x69, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x5d, 0x0a, 0x0f, 0x41, 0x70, 0x69, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x69, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x69, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_usage_proto_rawDescOnce sync.Once
	file_api_usage_proto_rawDescData = file_api_usage_proto_rawDesc
)

func file_api_usage_proto_rawDescGZIP() []byte {
	file_api_usage_proto_rawDescOnce.Do(func() {
		file_api_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_usage_proto_rawDescData)
	})
	return file_api_usage_proto_rawDescData
}

var file_api_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_usage_proto_goTypes = []interface{}{
	(*ApiUsageItem)(nil),        // 0: finazon.ApiUsageItem
	(*ApiCalls)(nil),            // 1: finazon.ApiCalls
	(*GetApiUsageRequest)(nil),  // 2: finazon.GetApiUsageRequest
	(*GetApiUsageResponse)(nil), // 3: finazon.GetApiUsageResponse
}
var file_api_usage_proto_depIdxs = []int32{
	1, // 0: finazon.ApiUsageItem.apiCalls:type_name -> finazon.ApiCalls
	1, // 1: finazon.ApiUsageItem.historicalApiCalls:type_name -> finazon.ApiCalls
	0, // 2: finazon.GetApiUsageResponse.result:type_name -> finazon.ApiUsageItem
	2, // 3: finazon.ApiUsageService.GetApiUsage:input_type -> finazon.GetApiUsageRequest
	3, // 4: finazon.ApiUsageService.GetApiUsage:output_type -> finazon.GetApiUsageResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_usage_proto_init() }
func file_api_usage_proto_init() {
	if File_api_usage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_usage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiUsageItem); i {
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
		file_api_usage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiCalls); i {
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
		file_api_usage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApiUsageRequest); i {
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
		file_api_usage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApiUsageResponse); i {
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
			RawDescriptor: file_api_usage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_usage_proto_goTypes,
		DependencyIndexes: file_api_usage_proto_depIdxs,
		MessageInfos:      file_api_usage_proto_msgTypes,
	}.Build()
	File_api_usage_proto = out.File
	file_api_usage_proto_rawDesc = nil
	file_api_usage_proto_goTypes = nil
	file_api_usage_proto_depIdxs = nil
}
