// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: common/ack.proto

package common

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

type Ack_STATUS int32

const (
	Ack_OK    Ack_STATUS = 0
	Ack_ERROR Ack_STATUS = 1
)

// Enum value maps for Ack_STATUS.
var (
	Ack_STATUS_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	Ack_STATUS_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x Ack_STATUS) Enum() *Ack_STATUS {
	p := new(Ack_STATUS)
	*p = x
	return p
}

func (x Ack_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ack_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_common_ack_proto_enumTypes[0].Descriptor()
}

func (Ack_STATUS) Type() protoreflect.EnumType {
	return &file_common_ack_proto_enumTypes[0]
}

func (x Ack_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ack_STATUS.Descriptor instead.
func (Ack_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_common_ack_proto_rawDescGZIP(), []int{0, 0}
}

// This message respresents "ACKs" sent between relay-relay,
// relay-driver and relay-network
type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    Ack_STATUS `protobuf:"varint,2,opt,name=status,proto3,enum=common.ack.Ack_STATUS" json:"status,omitempty"`
	RequestId string     `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// an error can have an associated string
	// this is the best way to represent this in protobuf
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_common_ack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_common_ack_proto_rawDescGZIP(), []int{0}
}

func (x *Ack) GetStatus() Ack_STATUS {
	if x != nil {
		return x.Status
	}
	return Ack_OK
}

func (x *Ack) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Ack) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_common_ack_proto protoreflect.FileDescriptor

var file_common_ack_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x22, 0x8b,
	0x01, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x61, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x2e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x1b, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x42, 0x6f, 0x0a, 0x1c,
	0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x5a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72,
	0x2d, 0x64, 0x6c, 0x74, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_ack_proto_rawDescOnce sync.Once
	file_common_ack_proto_rawDescData = file_common_ack_proto_rawDesc
)

func file_common_ack_proto_rawDescGZIP() []byte {
	file_common_ack_proto_rawDescOnce.Do(func() {
		file_common_ack_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_ack_proto_rawDescData)
	})
	return file_common_ack_proto_rawDescData
}

var file_common_ack_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_ack_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_ack_proto_goTypes = []interface{}{
	(Ack_STATUS)(0), // 0: common.ack.Ack.STATUS
	(*Ack)(nil),     // 1: common.ack.Ack
}
var file_common_ack_proto_depIdxs = []int32{
	0, // 0: common.ack.Ack.status:type_name -> common.ack.Ack.STATUS
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_ack_proto_init() }
func file_common_ack_proto_init() {
	if File_common_ack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_ack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
			RawDescriptor: file_common_ack_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_ack_proto_goTypes,
		DependencyIndexes: file_common_ack_proto_depIdxs,
		EnumInfos:         file_common_ack_proto_enumTypes,
		MessageInfos:      file_common_ack_proto_msgTypes,
	}.Build()
	File_common_ack_proto = out.File
	file_common_ack_proto_rawDesc = nil
	file_common_ack_proto_goTypes = nil
	file_common_ack_proto_depIdxs = nil
}
