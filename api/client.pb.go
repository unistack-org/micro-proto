// Copyright 2018 Google LLC
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
// 	protoc        v4.25.3
// source: api/client.proto

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

var file_api_client_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         65535657,
		Name:          "micro.api.default_host",
		Tag:           "bytes,65535657,opt,name=default_host",
		Filename:      "api/client.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         65535658,
		Name:          "micro.api.oauth_scopes",
		Tag:           "bytes,65535658,opt,name=oauth_scopes",
		Filename:      "api/client.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// The hostname for this service.
	// This should be specified with no prefix or protocol.
	//
	// Example:
	//
	//   service Foo {
	//     option (google.api.default_host) = "foo.googleapi.com";
	//     ...
	//   }
	//
	// optional string default_host = 65535657;
	E_DefaultHost = &file_api_client_proto_extTypes[0]
	// OAuth scopes needed for the client.
	//
	// Example:
	//
	//   service Foo {
	//     option (google.api.oauth_scopes) = \
	//       "https://www.googleapis.com/auth/cloud-platform";
	//     ...
	//   }
	//
	// If there is more than one scope, use a comma-separated string:
	//
	// Example:
	//
	//   service Foo {
	//     option (google.api.oauth_scopes) = \
	//       "https://www.googleapis.com/auth/cloud-platform,"
	//       "https://www.googleapis.com/auth/monitoring";
	//     ...
	//   }
	//
	// optional string oauth_scopes = 65535658;
	E_OauthScopes = &file_api_client_proto_extTypes[1]
)

var File_api_client_proto protoreflect.FileDescriptor

var file_api_client_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a,
	0x45, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xa9, 0xfd, 0x9f, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x3a, 0x45, 0x0a, 0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaa, 0xfd, 0x9f, 0x1f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x42, 0x4e, 0x0a,
	0x09, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0b, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2d, 0x6f,
	0x72, 0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x04, 0x4d, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_client_proto_goTypes = []interface{}{
	(*descriptorpb.ServiceOptions)(nil), // 0: google.protobuf.ServiceOptions
}
var file_api_client_proto_depIdxs = []int32{
	0, // 0: micro.api.default_host:extendee -> google.protobuf.ServiceOptions
	0, // 1: micro.api.oauth_scopes:extendee -> google.protobuf.ServiceOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_client_proto_init() }
func file_api_client_proto_init() {
	if File_api_client_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_api_client_proto_goTypes,
		DependencyIndexes: file_api_client_proto_depIdxs,
		ExtensionInfos:    file_api_client_proto_extTypes,
	}.Build()
	File_api_client_proto = out.File
	file_api_client_proto_rawDesc = nil
	file_api_client_proto_goTypes = nil
	file_api_client_proto_depIdxs = nil
}
