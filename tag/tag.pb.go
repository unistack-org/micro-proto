// Copyright 2020 Unistack LLC
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tag/tag.proto

package tag

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_tag_tag_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847939,
		Name:          "micro.tag.tags",
		Tag:           "bytes,847939,opt,name=tags",
		Filename:      "tag/tag.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847939,
		Name:          "micro.tag.oneof_tags",
		Tag:           "bytes,847939,opt,name=oneof_tags",
		Filename:      "tag/tag.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// Multiple Tags can be spcified.
	//
	// optional string tags = 847939;
	E_Tags = &file_tag_tag_proto_extTypes[0]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// Multiple Tags can be spcified.
	//
	// optional string oneof_tags = 847939;
	E_OneofTags = &file_tag_tag_proto_extTypes[1]
)

var File_tag_tag_proto protoreflect.FileDescriptor

var file_tag_tag_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x74, 0x61, 0x67, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x33, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xc3, 0xe0, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x3a, 0x3e, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc3,
	0xe0, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x54, 0x61, 0x67,
	0x73, 0x42, 0x4b, 0x0a, 0x09, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x74, 0x61, 0x67, 0x42, 0x0a,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x6f,
	0x2e, 0x75, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x61, 0x67,
	0x3b, 0x74, 0x61, 0x67, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x4d, 0x54, 0x41, 0x47, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tag_tag_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
	(*descriptorpb.OneofOptions)(nil), // 1: google.protobuf.OneofOptions
}
var file_tag_tag_proto_depIdxs = []int32{
	0, // 0: micro.tag.tags:extendee -> google.protobuf.FieldOptions
	1, // 1: micro.tag.oneof_tags:extendee -> google.protobuf.OneofOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tag_tag_proto_init() }
func file_tag_tag_proto_init() {
	if File_tag_tag_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tag_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_tag_tag_proto_goTypes,
		DependencyIndexes: file_tag_tag_proto_depIdxs,
		ExtensionInfos:    file_tag_tag_proto_extTypes,
	}.Build()
	File_tag_tag_proto = out.File
	file_tag_tag_proto_rawDesc = nil
	file_tag_tag_proto_goTypes = nil
	file_tag_tag_proto_depIdxs = nil
}
