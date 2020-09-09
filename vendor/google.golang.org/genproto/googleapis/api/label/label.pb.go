// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/api/label.proto

package label

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Value types that can be used as label values.
type LabelDescriptor_ValueType int32

const (
	// A variable-length string. This is the default.
	LabelDescriptor_STRING LabelDescriptor_ValueType = 0
	// Boolean; true or false.
	LabelDescriptor_BOOL LabelDescriptor_ValueType = 1
	// A 64-bit signed integer.
	LabelDescriptor_INT64 LabelDescriptor_ValueType = 2
)

// Enum value maps for LabelDescriptor_ValueType.
var (
	LabelDescriptor_ValueType_name = map[int32]string{
		0: "STRING",
		1: "BOOL",
		2: "INT64",
	}
	LabelDescriptor_ValueType_value = map[string]int32{
		"STRING": 0,
		"BOOL":   1,
		"INT64":  2,
	}
)

func (x LabelDescriptor_ValueType) Enum() *LabelDescriptor_ValueType {
	p := new(LabelDescriptor_ValueType)
	*p = x
	return p
}

func (x LabelDescriptor_ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LabelDescriptor_ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_api_label_proto_enumTypes[0].Descriptor()
}

func (LabelDescriptor_ValueType) Type() protoreflect.EnumType {
	return &file_google_api_label_proto_enumTypes[0]
}

func (x LabelDescriptor_ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LabelDescriptor_ValueType.Descriptor instead.
func (LabelDescriptor_ValueType) EnumDescriptor() ([]byte, []int) {
	return file_google_api_label_proto_rawDescGZIP(), []int{0, 0}
}

// A description of a label.
type LabelDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The label key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The type of data that can be assigned to the label.
	ValueType LabelDescriptor_ValueType `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=google.api.LabelDescriptor_ValueType" json:"value_type,omitempty"`
	// A human-readable description for the label.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *LabelDescriptor) Reset() {
	*x = LabelDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_label_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelDescriptor) ProtoMessage() {}

func (x *LabelDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_label_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelDescriptor.ProtoReflect.Descriptor instead.
func (*LabelDescriptor) Descriptor() ([]byte, []int) {
	return file_google_api_label_proto_rawDescGZIP(), []int{0}
}

func (x *LabelDescriptor) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LabelDescriptor) GetValueType() LabelDescriptor_ValueType {
	if x != nil {
		return x.ValueType
	}
	return LabelDescriptor_STRING
}

func (x *LabelDescriptor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_google_api_label_proto protoreflect.FileDescriptor

var file_google_api_label_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x22, 0xb9, 0x01, 0x0a, 0x0f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x0a, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42,
	0x4f, 0x4f, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x02,
	0x42, 0x5f, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x42, 0x0a, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x3b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x50,
	0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_label_proto_rawDescOnce sync.Once
	file_google_api_label_proto_rawDescData = file_google_api_label_proto_rawDesc
)

func file_google_api_label_proto_rawDescGZIP() []byte {
	file_google_api_label_proto_rawDescOnce.Do(func() {
		file_google_api_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_label_proto_rawDescData)
	})
	return file_google_api_label_proto_rawDescData
}

var file_google_api_label_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_api_label_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_api_label_proto_goTypes = []interface{}{
	(LabelDescriptor_ValueType)(0), // 0: google.api.LabelDescriptor.ValueType
	(*LabelDescriptor)(nil),        // 1: google.api.LabelDescriptor
}
var file_google_api_label_proto_depIdxs = []int32{
	0, // 0: google.api.LabelDescriptor.value_type:type_name -> google.api.LabelDescriptor.ValueType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_api_label_proto_init() }
func file_google_api_label_proto_init() {
	if File_google_api_label_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_label_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelDescriptor); i {
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
			RawDescriptor: file_google_api_label_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_label_proto_goTypes,
		DependencyIndexes: file_google_api_label_proto_depIdxs,
		EnumInfos:         file_google_api_label_proto_enumTypes,
		MessageInfos:      file_google_api_label_proto_msgTypes,
	}.Build()
	File_google_api_label_proto = out.File
	file_google_api_label_proto_rawDesc = nil
	file_google_api_label_proto_goTypes = nil
	file_google_api_label_proto_depIdxs = nil
}
