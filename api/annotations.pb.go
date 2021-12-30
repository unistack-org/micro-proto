// Copyright (c) 2015, Google Inc.
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/annotations.proto

package api

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

var file_api_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*HttpRule)(nil),
		Field:         65535654,
		Name:          "micro.api.http",
		Tag:           "bytes,65535654,opt,name=http",
		Filename:      "api/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MicroMethod)(nil),
		Field:         65535655,
		Name:          "micro.api.micro_method",
		Tag:           "bytes,65535655,opt,name=micro_method",
		Filename:      "api/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*MicroService)(nil),
		Field:         65535655,
		Name:          "micro.api.micro_service",
		Tag:           "bytes,65535655,opt,name=micro_service",
		Filename:      "api/annotations.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// See `HttpRule`.
	//
	// optional micro.api.HttpRule http = 65535654;
	E_Http = &file_api_annotations_proto_extTypes[0]
	// See `MicroMethod`.
	//
	// optional micro.api.MicroMethod micro_method = 65535655;
	E_MicroMethod = &file_api_annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// See `MicroService`
	//
	// optional micro.api.MicroService micro_service = 65535655;
	E_MicroService = &file_api_annotations_proto_extTypes[2]
)

var File_api_annotations_proto protoreflect.FileDescriptor

var file_api_annotations_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x4a, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa6, 0xfd,
	0x9f, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x68, 0x74, 0x74,
	0x70, 0x3a, 0x5c, 0x0a, 0x0c, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xa7, 0xfd, 0x9f, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x0b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a,
	0x60, 0x0a, 0x0d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xa7, 0xfd, 0x9f, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x0c, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x4e, 0x0a, 0x09, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x10,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x26, 0x67, 0x6f, 0x2e, 0x75, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x04, 0x4d, 0x41, 0x50,
	0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.MethodOptions)(nil),  // 0: google.protobuf.MethodOptions
	(*descriptorpb.ServiceOptions)(nil), // 1: google.protobuf.ServiceOptions
	(*HttpRule)(nil),                    // 2: micro.api.HttpRule
	(*MicroMethod)(nil),                 // 3: micro.api.MicroMethod
	(*MicroService)(nil),                // 4: micro.api.MicroService
}
var file_api_annotations_proto_depIdxs = []int32{
	0, // 0: micro.api.http:extendee -> google.protobuf.MethodOptions
	0, // 1: micro.api.micro_method:extendee -> google.protobuf.MethodOptions
	1, // 2: micro.api.micro_service:extendee -> google.protobuf.ServiceOptions
	2, // 3: micro.api.http:type_name -> micro.api.HttpRule
	3, // 4: micro.api.micro_method:type_name -> micro.api.MicroMethod
	4, // 5: micro.api.micro_service:type_name -> micro.api.MicroService
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	3, // [3:6] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_annotations_proto_init() }
func file_api_annotations_proto_init() {
	if File_api_annotations_proto != nil {
		return
	}
	file_api_http_proto_init()
	file_api_micro_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_api_annotations_proto_goTypes,
		DependencyIndexes: file_api_annotations_proto_depIdxs,
		ExtensionInfos:    file_api_annotations_proto_extTypes,
	}.Build()
	File_api_annotations_proto = out.File
	file_api_annotations_proto_rawDesc = nil
	file_api_annotations_proto_goTypes = nil
	file_api_annotations_proto_depIdxs = nil
}
