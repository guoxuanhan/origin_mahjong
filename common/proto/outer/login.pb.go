// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: outerproto/login.proto

package outer

import (
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

// 登录状态
type LoginStatus int32

const (
	LoginStatus_Loginning LoginStatus = 0
	LoginStatus_Logined   LoginStatus = 1
	LoginStatus_Logouting LoginStatus = 2
	LoginStatus_Logouted  LoginStatus = 3
)

// Enum value maps for LoginStatus.
var (
	LoginStatus_name = map[int32]string{
		0: "Loginning",
		1: "Logined",
		2: "Logouting",
		3: "Logouted",
	}
	LoginStatus_value = map[string]int32{
		"Loginning": 0,
		"Logined":   1,
		"Logouting": 2,
		"Logouted":  3,
	}
)

func (x LoginStatus) Enum() *LoginStatus {
	p := new(LoginStatus)
	*p = x
	return p
}

func (x LoginStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_outerproto_login_proto_enumTypes[0].Descriptor()
}

func (LoginStatus) Type() protoreflect.EnumType {
	return &file_outerproto_login_proto_enumTypes[0]
}

func (x LoginStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginStatus.Descriptor instead.
func (LoginStatus) EnumDescriptor() ([]byte, []int) {
	return file_outerproto_login_proto_rawDescGZIP(), []int{0}
}

// 登录类型
type LoginType int32

const (
	LoginType_Guest   LoginType = 0
	LoginType_Account LoginType = 1
	LoginType_Max     LoginType = 2
)

// Enum value maps for LoginType.
var (
	LoginType_name = map[int32]string{
		0: "Guest",
		1: "Account",
		2: "Max",
	}
	LoginType_value = map[string]int32{
		"Guest":   0,
		"Account": 1,
		"Max":     2,
	}
)

func (x LoginType) Enum() *LoginType {
	p := new(LoginType)
	*p = x
	return p
}

func (x LoginType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginType) Descriptor() protoreflect.EnumDescriptor {
	return file_outerproto_login_proto_enumTypes[1].Descriptor()
}

func (LoginType) Type() protoreflect.EnumType {
	return &file_outerproto_login_proto_enumTypes[1]
}

func (x LoginType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginType.Descriptor instead.
func (LoginType) EnumDescriptor() ([]byte, []int) {
	return file_outerproto_login_proto_rawDescGZIP(), []int{1}
}

// client->login 请求登录
type C2L_LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LoginType     LoginType              `protobuf:"varint,1,opt,name=LoginType,proto3,enum=outer.LoginType" json:"LoginType,omitempty"`
	Account       string                 `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *C2L_LoginRequest) Reset() {
	*x = C2L_LoginRequest{}
	mi := &file_outerproto_login_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *C2L_LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2L_LoginRequest) ProtoMessage() {}

func (x *C2L_LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_outerproto_login_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2L_LoginRequest.ProtoReflect.Descriptor instead.
func (*C2L_LoginRequest) Descriptor() ([]byte, []int) {
	return file_outerproto_login_proto_rawDescGZIP(), []int{0}
}

func (x *C2L_LoginRequest) GetLoginType() LoginType {
	if x != nil {
		return x.LoginType
	}
	return LoginType_Guest
}

func (x *C2L_LoginRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *C2L_LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// login->client 请求登录
type L2C_LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         int32                  `protobuf:"varint,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Token         string                 `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *L2C_LoginResponse) Reset() {
	*x = L2C_LoginResponse{}
	mi := &file_outerproto_login_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *L2C_LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L2C_LoginResponse) ProtoMessage() {}

func (x *L2C_LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_outerproto_login_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L2C_LoginResponse.ProtoReflect.Descriptor instead.
func (*L2C_LoginResponse) Descriptor() ([]byte, []int) {
	return file_outerproto_login_proto_rawDescGZIP(), []int{1}
}

func (x *L2C_LoginResponse) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *L2C_LoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *L2C_LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_outerproto_login_proto protoreflect.FileDescriptor

var file_outerproto_login_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x22,
	0x78, 0x0a, 0x10, 0x43, 0x32, 0x4c, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x59, 0x0a, 0x11, 0x4c, 0x32, 0x43,
	0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x46, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x10, 0x03, 0x2a, 0x2c, 0x0a, 0x09,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x61, 0x78, 0x10, 0x02, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_outerproto_login_proto_rawDescOnce sync.Once
	file_outerproto_login_proto_rawDescData []byte
)

func file_outerproto_login_proto_rawDescGZIP() []byte {
	file_outerproto_login_proto_rawDescOnce.Do(func() {
		file_outerproto_login_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_outerproto_login_proto_rawDesc), len(file_outerproto_login_proto_rawDesc)))
	})
	return file_outerproto_login_proto_rawDescData
}

var file_outerproto_login_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_outerproto_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_outerproto_login_proto_goTypes = []any{
	(LoginStatus)(0),          // 0: outer.LoginStatus
	(LoginType)(0),            // 1: outer.LoginType
	(*C2L_LoginRequest)(nil),  // 2: outer.C2L_LoginRequest
	(*L2C_LoginResponse)(nil), // 3: outer.L2C_LoginResponse
}
var file_outerproto_login_proto_depIdxs = []int32{
	1, // 0: outer.C2L_LoginRequest.LoginType:type_name -> outer.LoginType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_outerproto_login_proto_init() }
func file_outerproto_login_proto_init() {
	if File_outerproto_login_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_outerproto_login_proto_rawDesc), len(file_outerproto_login_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_outerproto_login_proto_goTypes,
		DependencyIndexes: file_outerproto_login_proto_depIdxs,
		EnumInfos:         file_outerproto_login_proto_enumTypes,
		MessageInfos:      file_outerproto_login_proto_msgTypes,
	}.Build()
	File_outerproto_login_proto = out.File
	file_outerproto_login_proto_goTypes = nil
	file_outerproto_login_proto_depIdxs = nil
}
