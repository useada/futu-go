// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: Qot_GetPriceReminder.proto

package qotgetpricereminder

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "pb/common"
	qotcommon "pb/qotcommon"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 提醒信息列表
type PriceReminderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      *int64   `protobuf:"varint,1,req,name=key" json:"key,omitempty"`           // 每个提醒的唯一标识
	Type     *int32   `protobuf:"varint,2,req,name=type" json:"type,omitempty"`         // Qot_Common::PriceReminderType 提醒类型
	Value    *float64 `protobuf:"fixed64,3,req,name=value" json:"value,omitempty"`      // 提醒参数值
	Note     *string  `protobuf:"bytes,4,req,name=note" json:"note,omitempty"`          // 备注仅支持 20 个以内的中文字符
	Freq     *int32   `protobuf:"varint,5,req,name=freq" json:"freq,omitempty"`         // Qot_Common::PriceReminderFreq 提醒频率类型
	IsEnable *bool    `protobuf:"varint,6,req,name=isEnable" json:"isEnable,omitempty"` // 该提醒设置是否生效。false不生效，true生效
}

func (x *PriceReminderItem) Reset() {
	*x = PriceReminderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetPriceReminder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceReminderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceReminderItem) ProtoMessage() {}

func (x *PriceReminderItem) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetPriceReminder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceReminderItem.ProtoReflect.Descriptor instead.
func (*PriceReminderItem) Descriptor() ([]byte, []int) {
	return file_Qot_GetPriceReminder_proto_rawDescGZIP(), []int{0}
}

func (x *PriceReminderItem) GetKey() int64 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *PriceReminderItem) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *PriceReminderItem) GetValue() float64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

func (x *PriceReminderItem) GetNote() string {
	if x != nil && x.Note != nil {
		return *x.Note
	}
	return ""
}

func (x *PriceReminderItem) GetFreq() int32 {
	if x != nil && x.Freq != nil {
		return *x.Freq
	}
	return 0
}

func (x *PriceReminderItem) GetIsEnable() bool {
	if x != nil && x.IsEnable != nil {
		return *x.IsEnable
	}
	return false
}

type PriceReminder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security *qotcommon.Security  `protobuf:"bytes,1,req,name=security" json:"security,omitempty"` // 股票
	ItemList []*PriceReminderItem `protobuf:"bytes,2,rep,name=itemList" json:"itemList,omitempty"` // 提醒信息列表
}

func (x *PriceReminder) Reset() {
	*x = PriceReminder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetPriceReminder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceReminder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceReminder) ProtoMessage() {}

func (x *PriceReminder) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetPriceReminder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceReminder.ProtoReflect.Descriptor instead.
func (*PriceReminder) Descriptor() ([]byte, []int) {
	return file_Qot_GetPriceReminder_proto_rawDescGZIP(), []int{1}
}

func (x *PriceReminder) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *PriceReminder) GetItemList() []*PriceReminderItem {
	if x != nil {
		return x.ItemList
	}
	return nil
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security *qotcommon.Security `protobuf:"bytes,1,opt,name=security" json:"security,omitempty"` // 查询股票下的到价提醒项，security和market二选一，都存在的情况下security优先。
	Market   *int32              `protobuf:"varint,2,opt,name=market" json:"market,omitempty"`    //Qot_Common::QotMarket 市场，查询市场下的到价提醒项，不区分沪深
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetPriceReminder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetPriceReminder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Qot_GetPriceReminder_proto_rawDescGZIP(), []int{2}
}

func (x *C2S) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *C2S) GetMarket() int32 {
	if x != nil && x.Market != nil {
		return *x.Market
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PriceReminderList []*PriceReminder `protobuf:"bytes,1,rep,name=priceReminderList" json:"priceReminderList,omitempty"` //到价提醒
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetPriceReminder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetPriceReminder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Qot_GetPriceReminder_proto_rawDescGZIP(), []int{3}
}

func (x *S2C) GetPriceReminderList() []*PriceReminder {
	if x != nil {
		return x.PriceReminderList
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetPriceReminder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetPriceReminder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Qot_GetPriceReminder_proto_rawDescGZIP(), []int{4}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetPriceReminder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetPriceReminder_proto_msgTypes[5]
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
	return file_Qot_GetPriceReminder_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Qot_GetPriceReminder_proto protoreflect.FileDescriptor

var file_Qot_GetPriceReminder_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x51, 0x6f,
	0x74, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x65, 0x71,
	0x18, 0x05, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x02, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51,
	0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x08,
	0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d,
	0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x4f, 0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x22, 0x58, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12, 0x51, 0x0a, 0x11, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x11, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x32, 0x53, 0x52,
	0x03, 0x63, 0x32, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63,
	0x42, 0x2d, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x16, 0x70, 0x62, 0x2f, 0x71, 0x6f, 0x74, 0x67,
	0x65, 0x74, 0x70, 0x72, 0x69, 0x63, 0x65, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
}

var (
	file_Qot_GetPriceReminder_proto_rawDescOnce sync.Once
	file_Qot_GetPriceReminder_proto_rawDescData = file_Qot_GetPriceReminder_proto_rawDesc
)

func file_Qot_GetPriceReminder_proto_rawDescGZIP() []byte {
	file_Qot_GetPriceReminder_proto_rawDescOnce.Do(func() {
		file_Qot_GetPriceReminder_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetPriceReminder_proto_rawDescData)
	})
	return file_Qot_GetPriceReminder_proto_rawDescData
}

var file_Qot_GetPriceReminder_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Qot_GetPriceReminder_proto_goTypes = []interface{}{
	(*PriceReminderItem)(nil),  // 0: Qot_GetPriceReminder.PriceReminderItem
	(*PriceReminder)(nil),      // 1: Qot_GetPriceReminder.PriceReminder
	(*C2S)(nil),                // 2: Qot_GetPriceReminder.C2S
	(*S2C)(nil),                // 3: Qot_GetPriceReminder.S2C
	(*Request)(nil),            // 4: Qot_GetPriceReminder.Request
	(*Response)(nil),           // 5: Qot_GetPriceReminder.Response
	(*qotcommon.Security)(nil), // 6: Qot_Common.Security
}
var file_Qot_GetPriceReminder_proto_depIdxs = []int32{
	6, // 0: Qot_GetPriceReminder.PriceReminder.security:type_name -> Qot_Common.Security
	0, // 1: Qot_GetPriceReminder.PriceReminder.itemList:type_name -> Qot_GetPriceReminder.PriceReminderItem
	6, // 2: Qot_GetPriceReminder.C2S.security:type_name -> Qot_Common.Security
	1, // 3: Qot_GetPriceReminder.S2C.priceReminderList:type_name -> Qot_GetPriceReminder.PriceReminder
	2, // 4: Qot_GetPriceReminder.Request.c2s:type_name -> Qot_GetPriceReminder.C2S
	3, // 5: Qot_GetPriceReminder.Response.s2c:type_name -> Qot_GetPriceReminder.S2C
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_Qot_GetPriceReminder_proto_init() }
func file_Qot_GetPriceReminder_proto_init() {
	if File_Qot_GetPriceReminder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetPriceReminder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceReminderItem); i {
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
		file_Qot_GetPriceReminder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceReminder); i {
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
		file_Qot_GetPriceReminder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_Qot_GetPriceReminder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_Qot_GetPriceReminder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Qot_GetPriceReminder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Qot_GetPriceReminder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetPriceReminder_proto_goTypes,
		DependencyIndexes: file_Qot_GetPriceReminder_proto_depIdxs,
		MessageInfos:      file_Qot_GetPriceReminder_proto_msgTypes,
	}.Build()
	File_Qot_GetPriceReminder_proto = out.File
	file_Qot_GetPriceReminder_proto_rawDesc = nil
	file_Qot_GetPriceReminder_proto_goTypes = nil
	file_Qot_GetPriceReminder_proto_depIdxs = nil
}
