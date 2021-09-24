// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: Notify.proto

package notify

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "pb/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotifyType int32

const (
	NotifyType_NotifyType_None          NotifyType = 0 //无
	NotifyType_NotifyType_GtwEvent      NotifyType = 1 //OpenD运行事件通知
	NotifyType_NotifyType_ProgramStatus NotifyType = 2 //程序状态
	NotifyType_NotifyType_ConnStatus    NotifyType = 3 //连接状态
	NotifyType_NotifyType_QotRight      NotifyType = 4 //行情权限
	NotifyType_NotifyType_APILevel      NotifyType = 5 //用户等级，已在2.10版本之后废弃
	NotifyType_NotifyType_APIQuota      NotifyType = 6 //API额度
)

// Enum value maps for NotifyType.
var (
	NotifyType_name = map[int32]string{
		0: "NotifyType_None",
		1: "NotifyType_GtwEvent",
		2: "NotifyType_ProgramStatus",
		3: "NotifyType_ConnStatus",
		4: "NotifyType_QotRight",
		5: "NotifyType_APILevel",
		6: "NotifyType_APIQuota",
	}
	NotifyType_value = map[string]int32{
		"NotifyType_None":          0,
		"NotifyType_GtwEvent":      1,
		"NotifyType_ProgramStatus": 2,
		"NotifyType_ConnStatus":    3,
		"NotifyType_QotRight":      4,
		"NotifyType_APILevel":      5,
		"NotifyType_APIQuota":      6,
	}
)

func (x NotifyType) Enum() *NotifyType {
	p := new(NotifyType)
	*p = x
	return p
}

func (x NotifyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifyType) Descriptor() protoreflect.EnumDescriptor {
	return file_Notify_proto_enumTypes[0].Descriptor()
}

func (NotifyType) Type() protoreflect.EnumType {
	return &file_Notify_proto_enumTypes[0]
}

func (x NotifyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *NotifyType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = NotifyType(num)
	return nil
}

// Deprecated: Use NotifyType.Descriptor instead.
func (NotifyType) EnumDescriptor() ([]byte, []int) {
	return file_Notify_proto_rawDescGZIP(), []int{0}
}

type GtwEventType int32

const (
	GtwEventType_GtwEventType_None                GtwEventType = 0  //正常无错
	GtwEventType_GtwEventType_LocalCfgLoadFailed  GtwEventType = 1  //加载本地配置失败
	GtwEventType_GtwEventType_APISvrRunFailed     GtwEventType = 2  //服务器启动失败
	GtwEventType_GtwEventType_ForceUpdate         GtwEventType = 3  //客户端版本过低
	GtwEventType_GtwEventType_LoginFailed         GtwEventType = 4  //登录失败
	GtwEventType_GtwEventType_UnAgreeDisclaimer   GtwEventType = 5  //未同意免责声明
	GtwEventType_GtwEventType_NetCfgMissing       GtwEventType = 6  //缺少必要网络配置信息;例如控制订阅额度 //已优化，不会再出现该情况
	GtwEventType_GtwEventType_KickedOut           GtwEventType = 7  //牛牛帐号在别处登录
	GtwEventType_GtwEventType_LoginPwdChanged     GtwEventType = 8  //登录密码被修改
	GtwEventType_GtwEventType_BanLogin            GtwEventType = 9  //用户被禁止登录
	GtwEventType_GtwEventType_NeedPicVerifyCode   GtwEventType = 10 //需要图形验证码
	GtwEventType_GtwEventType_NeedPhoneVerifyCode GtwEventType = 11 //需要手机验证码
	GtwEventType_GtwEventType_AppDataNotExist     GtwEventType = 12 //程序自带数据不存在
	GtwEventType_GtwEventType_NessaryDataMissing  GtwEventType = 13 //缺少必要数据
	GtwEventType_GtwEventType_TradePwdChanged     GtwEventType = 14 //交易密码被修改
	GtwEventType_GtwEventType_EnableDeviceLock    GtwEventType = 15 //启用设备锁
)

// Enum value maps for GtwEventType.
var (
	GtwEventType_name = map[int32]string{
		0:  "GtwEventType_None",
		1:  "GtwEventType_LocalCfgLoadFailed",
		2:  "GtwEventType_APISvrRunFailed",
		3:  "GtwEventType_ForceUpdate",
		4:  "GtwEventType_LoginFailed",
		5:  "GtwEventType_UnAgreeDisclaimer",
		6:  "GtwEventType_NetCfgMissing",
		7:  "GtwEventType_KickedOut",
		8:  "GtwEventType_LoginPwdChanged",
		9:  "GtwEventType_BanLogin",
		10: "GtwEventType_NeedPicVerifyCode",
		11: "GtwEventType_NeedPhoneVerifyCode",
		12: "GtwEventType_AppDataNotExist",
		13: "GtwEventType_NessaryDataMissing",
		14: "GtwEventType_TradePwdChanged",
		15: "GtwEventType_EnableDeviceLock",
	}
	GtwEventType_value = map[string]int32{
		"GtwEventType_None":                0,
		"GtwEventType_LocalCfgLoadFailed":  1,
		"GtwEventType_APISvrRunFailed":     2,
		"GtwEventType_ForceUpdate":         3,
		"GtwEventType_LoginFailed":         4,
		"GtwEventType_UnAgreeDisclaimer":   5,
		"GtwEventType_NetCfgMissing":       6,
		"GtwEventType_KickedOut":           7,
		"GtwEventType_LoginPwdChanged":     8,
		"GtwEventType_BanLogin":            9,
		"GtwEventType_NeedPicVerifyCode":   10,
		"GtwEventType_NeedPhoneVerifyCode": 11,
		"GtwEventType_AppDataNotExist":     12,
		"GtwEventType_NessaryDataMissing":  13,
		"GtwEventType_TradePwdChanged":     14,
		"GtwEventType_EnableDeviceLock":    15,
	}
)

func (x GtwEventType) Enum() *GtwEventType {
	p := new(GtwEventType)
	*p = x
	return p
}

func (x GtwEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GtwEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_Notify_proto_enumTypes[1].Descriptor()
}

func (GtwEventType) Type() protoreflect.EnumType {
	return &file_Notify_proto_enumTypes[1]
}

func (x GtwEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GtwEventType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GtwEventType(num)
	return nil
}

// Deprecated: Use GtwEventType.Descriptor instead.
func (GtwEventType) EnumDescriptor() ([]byte, []int) {
	return file_Notify_proto_rawDescGZIP(), []int{1}
}

type GtwEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType *int32  `protobuf:"varint,1,req,name=eventType" json:"eventType,omitempty"` //GtwEventType,事件类型
	Desc      *string `protobuf:"bytes,2,req,name=desc" json:"desc,omitempty"`            //事件描述
}

func (x *GtwEvent) Reset() {
	*x = GtwEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GtwEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GtwEvent) ProtoMessage() {}

func (x *GtwEvent) ProtoReflect() protoreflect.Message {
	mi := &file_Notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GtwEvent.ProtoReflect.Descriptor instead.
func (*GtwEvent) Descriptor() ([]byte, []int) {
	return file_Notify_proto_rawDescGZIP(), []int{0}
}

func (x *GtwEvent) GetEventType() int32 {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return 0
}

func (x *GtwEvent) GetDesc() string {
	if x != nil && x.Desc != nil {
		return *x.Desc
	}
	return ""
}

type ProgramStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProgramStatus *common.ProgramStatus `protobuf:"bytes,1,req,name=programStatus" json:"programStatus,omitempty"` //当前程序状态
}

func (x *ProgramStatus) Reset() {
	*x = ProgramStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgramStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramStatus) ProtoMessage() {}

func (x *ProgramStatus) ProtoReflect() protoreflect.Message {
	mi := &file_Notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgramStatus.ProtoReflect.Descriptor instead.
func (*ProgramStatus) Descriptor() ([]byte, []int) {
	return file_Notify_proto_rawDescGZIP(), []int{1}
}

func (x *ProgramStatus) GetProgramStatus() *common.ProgramStatus {
	if x != nil {
		return x.ProgramStatus
	}
	return nil
}

type ConnectStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QotLogined *bool `protobuf:"varint,1,req,name=qotLogined" json:"qotLogined,omitempty"` //是否登陆行情服务器
	TrdLogined *bool `protobuf:"varint,2,req,name=trdLogined" json:"trdLogined,omitempty"` //是否登陆交易服务器
}

func (x *ConnectStatus) Reset() {
	*x = ConnectStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Notify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectStatus) ProtoMessage() {}

func (x *ConnectStatus) ProtoReflect() protoreflect.Message {
	mi := &file_Notify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectStatus.ProtoReflect.Descriptor instead.
func (*ConnectStatus) Descriptor() ([]byte, []int) {
	return file_Notify_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectStatus) GetQotLogined() bool {
	if x != nil && x.QotLogined != nil {
		return *x.QotLogined
	}
	return false
}

func (x *ConnectStatus) GetTrdLogined() bool {
	if x != nil && x.TrdLogined != nil {
		return *x.TrdLogined
	}
	return false
}

type QotRight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HkQotRight          *int32 `protobuf:"varint,4,req,name=hkQotRight" json:"hkQotRight,omitempty"`                   //港股行情权限, Qot_Common.QotRight
	UsQotRight          *int32 `protobuf:"varint,5,req,name=usQotRight" json:"usQotRight,omitempty"`                   //美股行情权限, Qot_Common.QotRight
	CnQotRight          *int32 `protobuf:"varint,6,req,name=cnQotRight" json:"cnQotRight,omitempty"`                   //A股行情权限, Qot_Common.QotRight
	HkOptionQotRight    *int32 `protobuf:"varint,7,opt,name=hkOptionQotRight" json:"hkOptionQotRight,omitempty"`       //港股期权行情权限, Qot_Common.QotRight
	HasUSOptionQotRight *bool  `protobuf:"varint,8,opt,name=hasUSOptionQotRight" json:"hasUSOptionQotRight,omitempty"` //是否有美股期权行情权限
	HkFutureQotRight    *int32 `protobuf:"varint,9,opt,name=hkFutureQotRight" json:"hkFutureQotRight,omitempty"`       //港股期货行情权限, Qot_Common.QotRight
	UsFutureQotRight    *int32 `protobuf:"varint,10,opt,name=usFutureQotRight" json:"usFutureQotRight,omitempty"`      //美股期货行情权限, Qot_Common.QotRight
	UsOptionQotRight    *int32 `protobuf:"varint,11,opt,name=usOptionQotRight" json:"usOptionQotRight,omitempty"`      //美股期权行情权限, Qot_Common.QotRight
	UsIndexQotRight     *int32 `protobuf:"varint,12,opt,name=usIndexQotRight" json:"usIndexQotRight,omitempty"`        //美股指数行情权限, Qot_Common.QotRight
	UsOtcQotRight       *int32 `protobuf:"varint,13,opt,name=usOtcQotRight" json:"usOtcQotRight,omitempty"`            //美股OTC市场行情权限, Qot_Common.QotRight
}

func (x *QotRight) Reset() {
	*x = QotRight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Notify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QotRight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QotRight) ProtoMessage() {}

func (x *QotRight) ProtoReflect() protoreflect.Message {
	mi := &file_Notify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QotRight.ProtoReflect.Descriptor instead.
func (*QotRight) Descriptor() ([]byte, []int) {
	return file_Notify_proto_rawDescGZIP(), []int{3}
}

func (x *QotRight) GetHkQotRight() int32 {
	if x != nil && x.HkQotRight != nil {
		return *x.HkQotRight
	}
	return 0
}

func (x *QotRight) GetUsQotRight() int32 {
	if x != nil && x.UsQotRight != nil {
		return *x.UsQotRight
	}
	return 0
}

func (x *QotRight) GetCnQotRight() int32 {
	if x != nil && x.CnQotRight != nil {
		return *x.CnQotRight
	}
	return 0
}

func (x *QotRight) GetHkOptionQotRight() int32 {
	if x != nil && x.HkOptionQotRight != nil {
		return *x.HkOptionQotRight
	}
	return 0
}

func (x *QotRight) GetHasUSOptionQotRight() bool {
	if x != nil && x.HasUSOptionQotRight != nil {
		return *x.HasUSOptionQotRight
	}
	return false
}

func (x *QotRight) GetHkFutureQotRight() int32 {
	if x != nil && x.HkFutureQotRight != nil {
		return *x.HkFutureQotRight
	}
	return 0
}

func (x *QotRight) GetUsFutureQotRight() int32 {
	if x != nil && x.UsFutureQotRight != nil {
		return *x.UsFutureQotRight
	}
	return 0
}

func (x *QotRight) GetUsOptionQotRight() int32 {
	if x != nil && x.UsOptionQotRight != nil {
		return *x.UsOptionQotRight
	}
	return 0
}

func (x *QotRight) GetUsIndexQotRight() int32 {
	if x != nil && x.UsIndexQotRight != nil {
		return *x.UsIndexQotRight
	}
	return 0
}

func (x *QotRight) GetUsOtcQotRight() int32 {
	if x != nil && x.UsOtcQotRight != nil {
		return *x.UsOtcQotRight
	}
	return 0
}

type APILevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiLevel *string `protobuf:"bytes,3,req,name=apiLevel" json:"apiLevel,omitempty"` //api用户等级描述，已在2.10版本之后废弃
}

func (x *APILevel) Reset() {
	*x = APILevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Notify_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APILevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APILevel) ProtoMessage() {}

func (x *APILevel) ProtoReflect() protoreflect.Message {
	mi := &file_Notify_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APILevel.ProtoReflect.Descriptor instead.
func (*APILevel) Descriptor() ([]byte, []int) {
	return file_Notify_proto_rawDescGZIP(), []int{4}
}

func (x *APILevel) GetApiLevel() string {
	if x != nil && x.ApiLevel != nil {
		return *x.ApiLevel
	}
	return ""
}

type APIQuota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubQuota       *int32 `protobuf:"varint,1,req,name=subQuota" json:"subQuota,omitempty"`             //订阅额度
	HistoryKLQuota *int32 `protobuf:"varint,2,req,name=historyKLQuota" json:"historyKLQuota,omitempty"` //历史K线额度
}

func (x *APIQuota) Reset() {
	*x = APIQuota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Notify_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIQuota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIQuota) ProtoMessage() {}

func (x *APIQuota) ProtoReflect() protoreflect.Message {
	mi := &file_Notify_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIQuota.ProtoReflect.Descriptor instead.
func (*APIQuota) Descriptor() ([]byte, []int) {
	return file_Notify_proto_rawDescGZIP(), []int{5}
}

func (x *APIQuota) GetSubQuota() int32 {
	if x != nil && x.SubQuota != nil {
		return *x.SubQuota
	}
	return 0
}

func (x *APIQuota) GetHistoryKLQuota() int32 {
	if x != nil && x.HistoryKLQuota != nil {
		return *x.HistoryKLQuota
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          *int32         `protobuf:"varint,1,req,name=type" json:"type,omitempty"`                  //通知类型
	Event         *GtwEvent      `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`                 //事件通息
	ProgramStatus *ProgramStatus `protobuf:"bytes,3,opt,name=programStatus" json:"programStatus,omitempty"` //程序状态
	ConnectStatus *ConnectStatus `protobuf:"bytes,4,opt,name=connectStatus" json:"connectStatus,omitempty"` //连接状态
	QotRight      *QotRight      `protobuf:"bytes,5,opt,name=qotRight" json:"qotRight,omitempty"`           //行情权限
	ApiLevel      *APILevel      `protobuf:"bytes,6,opt,name=apiLevel" json:"apiLevel,omitempty"`           //用户等级，已在2.10版本之后废弃
	ApiQuota      *APIQuota      `protobuf:"bytes,7,opt,name=apiQuota" json:"apiQuota,omitempty"`           //API额度
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Notify_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Notify_proto_msgTypes[6]
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
	return file_Notify_proto_rawDescGZIP(), []int{6}
}

func (x *S2C) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *S2C) GetEvent() *GtwEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *S2C) GetProgramStatus() *ProgramStatus {
	if x != nil {
		return x.ProgramStatus
	}
	return nil
}

func (x *S2C) GetConnectStatus() *ConnectStatus {
	if x != nil {
		return x.ConnectStatus
	}
	return nil
}

func (x *S2C) GetQotRight() *QotRight {
	if x != nil {
		return x.QotRight
	}
	return nil
}

func (x *S2C) GetApiLevel() *APILevel {
	if x != nil {
		return x.ApiLevel
	}
	return nil
}

func (x *S2C) GetApiQuota() *APIQuota {
	if x != nil {
		return x.ApiQuota
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
		mi := &file_Notify_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Notify_proto_msgTypes[7]
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
	return file_Notify_proto_rawDescGZIP(), []int{7}
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

var File_Notify_proto protoreflect.FileDescriptor

var file_Notify_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x08, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x22, 0x4c, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x4f, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x6f, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0a, 0x71, 0x6f, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65,
	0x64, 0x22, 0x9c, 0x03, 0x0a, 0x08, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x68, 0x6b, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0a, 0x68, 0x6b, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x73, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0a, 0x75, 0x73, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0a, 0x63, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2a,
	0x0a, 0x10, 0x68, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x68, 0x6b, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x68, 0x61,
	0x73, 0x55, 0x53, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x68, 0x61, 0x73, 0x55, 0x53, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x68, 0x6b, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x68, 0x6b, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65,
	0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x73, 0x46, 0x75,
	0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x75, 0x73, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x51, 0x6f, 0x74, 0x52,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x75, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x51, 0x6f, 0x74, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x75, 0x73, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73,
	0x4f, 0x74, 0x63, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x75, 0x73, 0x4f, 0x74, 0x63, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x26, 0x0a, 0x08, 0x41, 0x50, 0x49, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x70, 0x69, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x70, 0x69, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x4e, 0x0a, 0x08, 0x41, 0x50, 0x49, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x62, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x12, 0x26, 0x0a, 0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x4b, 0x4c, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x22, 0xc5, 0x02, 0x0a, 0x03, 0x53, 0x32, 0x43,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x47, 0x74, 0x77,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x71, 0x6f, 0x74, 0x52, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x52, 0x08, 0x71, 0x6f, 0x74, 0x52,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x41, 0x50, 0x49, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x41, 0x50,
	0x49, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x08, 0x61, 0x70, 0x69, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x22, 0x7b, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d,
	0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x2a, 0xbe, 0x01,
	0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x47, 0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x43, 0x6f, 0x6e, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x5f, 0x51, 0x6f, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x41, 0x50, 0x49, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x41, 0x50, 0x49, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x10, 0x06, 0x2a, 0x9b,
	0x04, 0x0a, 0x0c, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x66, 0x67, 0x4c,
	0x6f, 0x61, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x47,
	0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x41, 0x50, 0x49, 0x53,
	0x76, 0x72, 0x52, 0x75, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x12, 0x1c, 0x0a,
	0x18, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x46, 0x6f,
	0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x47,
	0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x47, 0x74, 0x77,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x41, 0x67, 0x72, 0x65,
	0x65, 0x44, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x72, 0x10, 0x05, 0x12, 0x1e, 0x0a,
	0x1a, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x65,
	0x74, 0x43, 0x66, 0x67, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x10, 0x06, 0x12, 0x1a, 0x0a,
	0x16, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4b, 0x69,
	0x63, 0x6b, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x10, 0x07, 0x12, 0x20, 0x0a, 0x1c, 0x47, 0x74, 0x77,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50,
	0x77, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x47,
	0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x42, 0x61, 0x6e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x10, 0x09, 0x12, 0x22, 0x0a, 0x1e, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x65, 0x65, 0x64, 0x50, 0x69, 0x63, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x10, 0x0a, 0x12, 0x24, 0x0a, 0x20, 0x47, 0x74,
	0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x65, 0x65, 0x64, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x10, 0x0b,
	0x12, 0x20, 0x0a, 0x1c, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x10, 0x0c, 0x12, 0x23, 0x0a, 0x1f, 0x47, 0x74, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x4e, 0x65, 0x73, 0x73, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x10, 0x0d, 0x12, 0x20, 0x0a, 0x1c, 0x47, 0x74, 0x77, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x77, 0x64,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x10, 0x0e, 0x12, 0x21, 0x0a, 0x1d, 0x47, 0x74, 0x77,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x10, 0x0f, 0x42, 0x20, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x62, 0x5a, 0x09, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
}

var (
	file_Notify_proto_rawDescOnce sync.Once
	file_Notify_proto_rawDescData = file_Notify_proto_rawDesc
)

func file_Notify_proto_rawDescGZIP() []byte {
	file_Notify_proto_rawDescOnce.Do(func() {
		file_Notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_Notify_proto_rawDescData)
	})
	return file_Notify_proto_rawDescData
}

var file_Notify_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Notify_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_Notify_proto_goTypes = []interface{}{
	(NotifyType)(0),              // 0: Notify.NotifyType
	(GtwEventType)(0),            // 1: Notify.GtwEventType
	(*GtwEvent)(nil),             // 2: Notify.GtwEvent
	(*ProgramStatus)(nil),        // 3: Notify.ProgramStatus
	(*ConnectStatus)(nil),        // 4: Notify.ConnectStatus
	(*QotRight)(nil),             // 5: Notify.QotRight
	(*APILevel)(nil),             // 6: Notify.APILevel
	(*APIQuota)(nil),             // 7: Notify.APIQuota
	(*S2C)(nil),                  // 8: Notify.S2C
	(*Response)(nil),             // 9: Notify.Response
	(*common.ProgramStatus)(nil), // 10: Common.ProgramStatus
}
var file_Notify_proto_depIdxs = []int32{
	10, // 0: Notify.ProgramStatus.programStatus:type_name -> Common.ProgramStatus
	2,  // 1: Notify.S2C.event:type_name -> Notify.GtwEvent
	3,  // 2: Notify.S2C.programStatus:type_name -> Notify.ProgramStatus
	4,  // 3: Notify.S2C.connectStatus:type_name -> Notify.ConnectStatus
	5,  // 4: Notify.S2C.qotRight:type_name -> Notify.QotRight
	6,  // 5: Notify.S2C.apiLevel:type_name -> Notify.APILevel
	7,  // 6: Notify.S2C.apiQuota:type_name -> Notify.APIQuota
	8,  // 7: Notify.Response.s2c:type_name -> Notify.S2C
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_Notify_proto_init() }
func file_Notify_proto_init() {
	if File_Notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GtwEvent); i {
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
		file_Notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgramStatus); i {
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
		file_Notify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectStatus); i {
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
		file_Notify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QotRight); i {
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
		file_Notify_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APILevel); i {
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
		file_Notify_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIQuota); i {
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
		file_Notify_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_Notify_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_Notify_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Notify_proto_goTypes,
		DependencyIndexes: file_Notify_proto_depIdxs,
		EnumInfos:         file_Notify_proto_enumTypes,
		MessageInfos:      file_Notify_proto_msgTypes,
	}.Build()
	File_Notify_proto = out.File
	file_Notify_proto_rawDesc = nil
	file_Notify_proto_goTypes = nil
	file_Notify_proto_depIdxs = nil
}
