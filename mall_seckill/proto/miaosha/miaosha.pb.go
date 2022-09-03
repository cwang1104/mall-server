// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: proto/miaosha/miaosha.proto

package pbMiaoSha

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

type MiaoshaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FrontUserEmail string `protobuf:"bytes,2,opt,name=front_user_email,json=frontUserEmail,proto3" json:"front_user_email,omitempty"`
	UserID         int32  `protobuf:"varint,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *MiaoshaRequest) Reset() {
	*x = MiaoshaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miaosha_miaosha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiaoshaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiaoshaRequest) ProtoMessage() {}

func (x *MiaoshaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miaosha_miaosha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiaoshaRequest.ProtoReflect.Descriptor instead.
func (*MiaoshaRequest) Descriptor() ([]byte, []int) {
	return file_proto_miaosha_miaosha_proto_rawDescGZIP(), []int{0}
}

func (x *MiaoshaRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MiaoshaRequest) GetFrontUserEmail() string {
	if x != nil {
		return x.FrontUserEmail
	}
	return ""
}

func (x *MiaoshaRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type MiaoShaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *MiaoShaResponse) Reset() {
	*x = MiaoShaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_miaosha_miaosha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiaoShaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiaoShaResponse) ProtoMessage() {}

func (x *MiaoShaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_miaosha_miaosha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiaoShaResponse.ProtoReflect.Descriptor instead.
func (*MiaoShaResponse) Descriptor() ([]byte, []int) {
	return file_proto_miaosha_miaosha_proto_rawDescGZIP(), []int{1}
}

func (x *MiaoShaResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MiaoShaResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_miaosha_miaosha_proto protoreflect.FileDescriptor

var file_proto_miaosha_miaosha_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x61, 0x6f, 0x73, 0x68, 0x61, 0x2f,
	0x6d, 0x69, 0x61, 0x6f, 0x73, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x22, 0x62, 0x0a, 0x0e, 0x4d, 0x69, 0x61, 0x6f, 0x73, 0x68,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x37, 0x0a, 0x0f, 0x4d, 0x69,
	0x61, 0x6f, 0x53, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x32, 0x4c, 0x0a, 0x07, 0x4d, 0x69, 0x61, 0x6f, 0x53, 0x68, 0x61, 0x12, 0x41,
	0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x4d, 0x69, 0x61, 0x6f, 0x53, 0x68, 0x61, 0x12, 0x17,
	0x2e, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x4d, 0x69, 0x61, 0x6f, 0x73, 0x68, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c,
	0x6c, 0x2e, 0x4d, 0x69, 0x61, 0x6f, 0x53, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x69, 0x61,
	0x6f, 0x73, 0x68, 0x61, 0x3b, 0x70, 0x62, 0x4d, 0x69, 0x61, 0x6f, 0x53, 0x68, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_miaosha_miaosha_proto_rawDescOnce sync.Once
	file_proto_miaosha_miaosha_proto_rawDescData = file_proto_miaosha_miaosha_proto_rawDesc
)

func file_proto_miaosha_miaosha_proto_rawDescGZIP() []byte {
	file_proto_miaosha_miaosha_proto_rawDescOnce.Do(func() {
		file_proto_miaosha_miaosha_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_miaosha_miaosha_proto_rawDescData)
	})
	return file_proto_miaosha_miaosha_proto_rawDescData
}

var file_proto_miaosha_miaosha_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_miaosha_miaosha_proto_goTypes = []interface{}{
	(*MiaoshaRequest)(nil),  // 0: seckill.MiaoshaRequest
	(*MiaoShaResponse)(nil), // 1: seckill.MiaoShaResponse
}
var file_proto_miaosha_miaosha_proto_depIdxs = []int32{
	0, // 0: seckill.MiaoSha.FrontMiaoSha:input_type -> seckill.MiaoshaRequest
	1, // 1: seckill.MiaoSha.FrontMiaoSha:output_type -> seckill.MiaoShaResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_miaosha_miaosha_proto_init() }
func file_proto_miaosha_miaosha_proto_init() {
	if File_proto_miaosha_miaosha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_miaosha_miaosha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiaoshaRequest); i {
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
		file_proto_miaosha_miaosha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiaoShaResponse); i {
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
			RawDescriptor: file_proto_miaosha_miaosha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_miaosha_miaosha_proto_goTypes,
		DependencyIndexes: file_proto_miaosha_miaosha_proto_depIdxs,
		MessageInfos:      file_proto_miaosha_miaosha_proto_msgTypes,
	}.Build()
	File_proto_miaosha_miaosha_proto = out.File
	file_proto_miaosha_miaosha_proto_rawDesc = nil
	file_proto_miaosha_miaosha_proto_goTypes = nil
	file_proto_miaosha_miaosha_proto_depIdxs = nil
}