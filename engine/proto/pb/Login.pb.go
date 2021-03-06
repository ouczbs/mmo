// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: Login.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LoginAccountCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginAccountCmd) Reset() {
	*x = LoginAccountCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccountCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccountCmd) ProtoMessage() {}

func (x *LoginAccountCmd) ProtoReflect() protoreflect.Message {
	mi := &file_Login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccountCmd.ProtoReflect.Descriptor instead.
func (*LoginAccountCmd) Descriptor() ([]byte, []int) {
	return file_Login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginAccountCmd) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginAccountCmd) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginAccountAckCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     string  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Success    bool    `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	RoleIdList []int32 `protobuf:"varint,3,rep,packed,name=role_id_list,json=roleIdList,proto3" json:"role_id_list,omitempty"`
}

func (x *LoginAccountAckCmd) Reset() {
	*x = LoginAccountAckCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccountAckCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccountAckCmd) ProtoMessage() {}

func (x *LoginAccountAckCmd) ProtoReflect() protoreflect.Message {
	mi := &file_Login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccountAckCmd.ProtoReflect.Descriptor instead.
func (*LoginAccountAckCmd) Descriptor() ([]byte, []int) {
	return file_Login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginAccountAckCmd) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *LoginAccountAckCmd) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoginAccountAckCmd) GetRoleIdList() []int32 {
	if x != nil {
		return x.RoleIdList
	}
	return nil
}

var File_Login_proto protoreflect.FileDescriptor

var file_Login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x47, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x6d, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x68, 0x0a, 0x12, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x6b, 0x43, 0x6d, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Login_proto_rawDescOnce sync.Once
	file_Login_proto_rawDescData = file_Login_proto_rawDesc
)

func file_Login_proto_rawDescGZIP() []byte {
	file_Login_proto_rawDescOnce.Do(func() {
		file_Login_proto_rawDescData = protoimpl.X.CompressGZIP(file_Login_proto_rawDescData)
	})
	return file_Login_proto_rawDescData
}

var file_Login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Login_proto_goTypes = []interface{}{
	(*LoginAccountCmd)(nil),    // 0: pb.LoginAccountCmd
	(*LoginAccountAckCmd)(nil), // 1: pb.LoginAccountAckCmd
}
var file_Login_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Login_proto_init() }
func file_Login_proto_init() {
	if File_Login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccountCmd); i {
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
		file_Login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccountAckCmd); i {
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
			RawDescriptor: file_Login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Login_proto_goTypes,
		DependencyIndexes: file_Login_proto_depIdxs,
		MessageInfos:      file_Login_proto_msgTypes,
	}.Build()
	File_Login_proto = out.File
	file_Login_proto_rawDesc = nil
	file_Login_proto_goTypes = nil
	file_Login_proto_depIdxs = nil
}
