// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: exchange.proto

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

type GetMarketsCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetMarketsCryptoRequest) Reset() {
	*x = GetMarketsCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarketsCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarketsCryptoRequest) ProtoMessage() {}

func (x *GetMarketsCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarketsCryptoRequest.ProtoReflect.Descriptor instead.
func (*GetMarketsCryptoRequest) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *GetMarketsCryptoRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetMarketsCryptoRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type MarketsCryptoResponseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MarketsCryptoResponseItem) Reset() {
	*x = MarketsCryptoResponseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketsCryptoResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketsCryptoResponseItem) ProtoMessage() {}

func (x *MarketsCryptoResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketsCryptoResponseItem.ProtoReflect.Descriptor instead.
func (*MarketsCryptoResponseItem) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *MarketsCryptoResponseItem) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MarketsCryptoResponseItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetMarketsCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*MarketsCryptoResponseItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetMarketsCryptoResponse) Reset() {
	*x = GetMarketsCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarketsCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarketsCryptoResponse) ProtoMessage() {}

func (x *GetMarketsCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarketsCryptoResponse.ProtoReflect.Descriptor instead.
func (*GetMarketsCryptoResponse) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *GetMarketsCryptoResponse) GetResult() []*MarketsCryptoResponseItem {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetMarketsStocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country  string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Page     int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetMarketsStocksRequest) Reset() {
	*x = GetMarketsStocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarketsStocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarketsStocksRequest) ProtoMessage() {}

func (x *GetMarketsStocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarketsStocksRequest.ProtoReflect.Descriptor instead.
func (*GetMarketsStocksRequest) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *GetMarketsStocksRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetMarketsStocksRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetMarketsStocksRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetMarketsStocksRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetMarketsStocksRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type MarketsStocksResponseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Mic     string `protobuf:"bytes,2,opt,name=mic,proto3" json:"mic,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MarketsStocksResponseItem) Reset() {
	*x = MarketsStocksResponseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketsStocksResponseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketsStocksResponseItem) ProtoMessage() {}

func (x *MarketsStocksResponseItem) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketsStocksResponseItem.ProtoReflect.Descriptor instead.
func (*MarketsStocksResponseItem) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{4}
}

func (x *MarketsStocksResponseItem) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MarketsStocksResponseItem) GetMic() string {
	if x != nil {
		return x.Mic
	}
	return ""
}

func (x *MarketsStocksResponseItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetMarketsStocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*MarketsStocksResponseItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetMarketsStocksResponse) Reset() {
	*x = GetMarketsStocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarketsStocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarketsStocksResponse) ProtoMessage() {}

func (x *GetMarketsStocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarketsStocksResponse.ProtoReflect.Descriptor instead.
func (*GetMarketsStocksResponse) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{5}
}

func (x *GetMarketsStocksResponse) GetResult() []*MarketsStocksResponseItem {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_exchange_proto protoreflect.FileDescriptor

var file_exchange_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x43, 0x0a, 0x19, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x5b, 0x0a, 0x19, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xc7, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x20,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x73, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x73, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x20, 0x2e, 0x66, 0x69, 0x6e, 0x61,
	0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x69,
	0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exchange_proto_rawDescOnce sync.Once
	file_exchange_proto_rawDescData = file_exchange_proto_rawDesc
)

func file_exchange_proto_rawDescGZIP() []byte {
	file_exchange_proto_rawDescOnce.Do(func() {
		file_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchange_proto_rawDescData)
	})
	return file_exchange_proto_rawDescData
}

var file_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_exchange_proto_goTypes = []interface{}{
	(*GetMarketsCryptoRequest)(nil),   // 0: finazon.GetMarketsCryptoRequest
	(*MarketsCryptoResponseItem)(nil), // 1: finazon.MarketsCryptoResponseItem
	(*GetMarketsCryptoResponse)(nil),  // 2: finazon.GetMarketsCryptoResponse
	(*GetMarketsStocksRequest)(nil),   // 3: finazon.GetMarketsStocksRequest
	(*MarketsStocksResponseItem)(nil), // 4: finazon.MarketsStocksResponseItem
	(*GetMarketsStocksResponse)(nil),  // 5: finazon.GetMarketsStocksResponse
}
var file_exchange_proto_depIdxs = []int32{
	1, // 0: finazon.GetMarketsCryptoResponse.result:type_name -> finazon.MarketsCryptoResponseItem
	4, // 1: finazon.GetMarketsStocksResponse.result:type_name -> finazon.MarketsStocksResponseItem
	0, // 2: finazon.ExchangeService.GetMarketsCrypto:input_type -> finazon.GetMarketsCryptoRequest
	3, // 3: finazon.ExchangeService.GetMarketsStocks:input_type -> finazon.GetMarketsStocksRequest
	2, // 4: finazon.ExchangeService.GetMarketsCrypto:output_type -> finazon.GetMarketsCryptoResponse
	5, // 5: finazon.ExchangeService.GetMarketsStocks:output_type -> finazon.GetMarketsStocksResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_exchange_proto_init() }
func file_exchange_proto_init() {
	if File_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarketsCryptoRequest); i {
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
		file_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketsCryptoResponseItem); i {
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
		file_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarketsCryptoResponse); i {
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
		file_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarketsStocksRequest); i {
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
		file_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketsStocksResponseItem); i {
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
		file_exchange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarketsStocksResponse); i {
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
			RawDescriptor: file_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exchange_proto_goTypes,
		DependencyIndexes: file_exchange_proto_depIdxs,
		MessageInfos:      file_exchange_proto_msgTypes,
	}.Build()
	File_exchange_proto = out.File
	file_exchange_proto_rawDesc = nil
	file_exchange_proto_goTypes = nil
	file_exchange_proto_depIdxs = nil
}
