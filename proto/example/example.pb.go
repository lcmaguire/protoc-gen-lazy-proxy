// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: example/example.proto

package example

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

type ExtraRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ExtraRequest) Reset() {
	*x = ExtraRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtraRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtraRequest) ProtoMessage() {}

func (x *ExtraRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtraRequest.ProtoReflect.Descriptor instead.
func (*ExtraRequest) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *ExtraRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ExtraResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ExtraResponse) Reset() {
	*x = ExtraResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtraResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtraResponse) ProtoMessage() {}

func (x *ExtraResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtraResponse.ProtoReflect.Descriptor instead.
func (*ExtraResponse) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{1}
}

func (x *ExtraResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_example_example_proto protoreflect.FileDescriptor

var file_example_example_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x22, 0x22, 0x0a, 0x0c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x72, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x46, 0x0a, 0x0c, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x12, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x63, 0x6d, 0x61, 0x67, 0x75, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x61, 0x7a, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_example_proto_rawDescOnce sync.Once
	file_example_example_proto_rawDescData = file_example_example_proto_rawDesc
)

func file_example_example_proto_rawDescGZIP() []byte {
	file_example_example_proto_rawDescOnce.Do(func() {
		file_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_example_proto_rawDescData)
	})
	return file_example_example_proto_rawDescData
}

var file_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_example_proto_goTypes = []interface{}{
	(*ExtraRequest)(nil),  // 0: example.ExtraRequest
	(*ExtraResponse)(nil), // 1: example.ExtraResponse
}
var file_example_example_proto_depIdxs = []int32{
	0, // 0: example.ExtraService.Extra:input_type -> example.ExtraRequest
	1, // 1: example.ExtraService.Extra:output_type -> example.ExtraResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_example_proto_init() }
func file_example_example_proto_init() {
	if File_example_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtraRequest); i {
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
		file_example_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtraResponse); i {
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
			RawDescriptor: file_example_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_example_proto_goTypes,
		DependencyIndexes: file_example_example_proto_depIdxs,
		MessageInfos:      file_example_example_proto_msgTypes,
	}.Build()
	File_example_example_proto = out.File
	file_example_example_proto_rawDesc = nil
	file_example_example_proto_goTypes = nil
	file_example_example_proto_depIdxs = nil
}
