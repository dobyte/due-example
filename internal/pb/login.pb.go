// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: login.proto

package pb

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

type RegisterCode_Code int32

const (
	RegisterCode_Ok            RegisterCode_Code = 0 // 注册成功
	RegisterCode_Failed        RegisterCode_Code = 1 // 注册失败
	RegisterCode_AccountExists RegisterCode_Code = 2 // 账号已存在
)

// Enum value maps for RegisterCode_Code.
var (
	RegisterCode_Code_name = map[int32]string{
		0: "Ok",
		1: "Failed",
		2: "AccountExists",
	}
	RegisterCode_Code_value = map[string]int32{
		"Ok":            0,
		"Failed":        1,
		"AccountExists": 2,
	}
)

func (x RegisterCode_Code) Enum() *RegisterCode_Code {
	p := new(RegisterCode_Code)
	*p = x
	return p
}

func (x RegisterCode_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterCode_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_login_proto_enumTypes[0].Descriptor()
}

func (RegisterCode_Code) Type() protoreflect.EnumType {
	return &file_login_proto_enumTypes[0]
}

func (x RegisterCode_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterCode_Code.Descriptor instead.
func (RegisterCode_Code) EnumDescriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0, 0}
}

type LoginCode_Code int32

const (
	LoginCode_Ok                         LoginCode_Code = 0 // 登录成功
	LoginCode_Failed                     LoginCode_Code = 1 // 登录失败
	LoginCode_IncorrectAccountOrPassword LoginCode_Code = 2 // 账号或密码错误
)

// Enum value maps for LoginCode_Code.
var (
	LoginCode_Code_name = map[int32]string{
		0: "Ok",
		1: "Failed",
		2: "IncorrectAccountOrPassword",
	}
	LoginCode_Code_value = map[string]int32{
		"Ok":                         0,
		"Failed":                     1,
		"IncorrectAccountOrPassword": 2,
	}
)

func (x LoginCode_Code) Enum() *LoginCode_Code {
	p := new(LoginCode_Code)
	*p = x
	return p
}

func (x LoginCode_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginCode_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_login_proto_enumTypes[1].Descriptor()
}

func (LoginCode_Code) Type() protoreflect.EnumType {
	return &file_login_proto_enumTypes[1]
}

func (x LoginCode_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginCode_Code.Descriptor instead.
func (LoginCode_Code) EnumDescriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1, 0}
}

type RegisterCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterCode) Reset() {
	*x = RegisterCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCode) ProtoMessage() {}

func (x *RegisterCode) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCode.ProtoReflect.Descriptor instead.
func (*RegisterCode) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

type LoginCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginCode) Reset() {
	*x = LoginCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginCode) ProtoMessage() {}

func (x *LoginCode) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginCode.ProtoReflect.Descriptor instead.
func (*LoginCode) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

// 注册请求
type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`   // 账号
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"` // 密码
	Nickname string `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty"` // 昵称
	Age      int32  `protobuf:"varint,4,opt,name=Age,proto3" json:"Age,omitempty"`          // 年龄
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RegisterReq) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

// 注册响应
type RegisterRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    RegisterCode_Code `protobuf:"varint,1,opt,name=Code,proto3,enum=pb.RegisterCode_Code" json:"Code,omitempty"` // 返回码
	ID      int32             `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`                               // 用户ID
	Account string            `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`                      // 账号
}

func (x *RegisterRes) Reset() {
	*x = RegisterRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRes) ProtoMessage() {}

func (x *RegisterRes) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRes.ProtoReflect.Descriptor instead.
func (*RegisterRes) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRes) GetCode() RegisterCode_Code {
	if x != nil {
		return x.Code
	}
	return RegisterCode_Ok
}

func (x *RegisterRes) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RegisterRes) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

// 登录请求
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`   // 账号
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"` // 密码
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 登录响应
type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    LoginCode_Code `protobuf:"varint,1,opt,name=Code,proto3,enum=pb.LoginCode_Code" json:"Code,omitempty"` // 返回码
	ID      int32          `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`                            // 用户ID
	Account string         `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`                   // 账号
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{5}
}

func (x *LoginRes) GetCode() LoginCode_Code {
	if x != nil {
		return x.Code
	}
	return LoginCode_Ok
}

func (x *LoginRes) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *LoginRes) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x3d, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x2d, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x10, 0x02,
	0x22, 0x47, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x6e, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x72, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x02, 0x22, 0x71, 0x0a, 0x0b, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x67, 0x65, 0x22, 0x62, 0x0a, 0x0b,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x40, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x5c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_login_proto_goTypes = []interface{}{
	(RegisterCode_Code)(0), // 0: pb.RegisterCode.Code
	(LoginCode_Code)(0),    // 1: pb.LoginCode.Code
	(*RegisterCode)(nil),   // 2: pb.RegisterCode
	(*LoginCode)(nil),      // 3: pb.LoginCode
	(*RegisterReq)(nil),    // 4: pb.RegisterReq
	(*RegisterRes)(nil),    // 5: pb.RegisterRes
	(*LoginReq)(nil),       // 6: pb.LoginReq
	(*LoginRes)(nil),       // 7: pb.LoginRes
}
var file_login_proto_depIdxs = []int32{
	0, // 0: pb.RegisterRes.Code:type_name -> pb.RegisterCode.Code
	1, // 1: pb.LoginRes.Code:type_name -> pb.LoginCode.Code
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterCode); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginCode); i {
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
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRes); i {
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
		file_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		EnumInfos:         file_login_proto_enumTypes,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
