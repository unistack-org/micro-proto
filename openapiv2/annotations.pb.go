// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: openapiv2/annotations.proto

package openapiv2

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

var file_openapiv2_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*Document)(nil),
		Field:         847940,
		Name:          "micro.openapiv2.openapiv2_swagger",
		Tag:           "bytes,847940,opt,name=openapiv2_swagger",
		Filename:      "openapiv2/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Operation)(nil),
		Field:         847940,
		Name:          "micro.openapiv2.openapiv2_operation",
		Tag:           "bytes,847940,opt,name=openapiv2_operation",
		Filename:      "openapiv2/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Schema)(nil),
		Field:         847940,
		Name:          "micro.openapiv2.openapiv2_schema",
		Tag:           "bytes,847940,opt,name=openapiv2_schema",
		Filename:      "openapiv2/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*Tag)(nil),
		Field:         847940,
		Name:          "micro.openapiv2.openapiv2_tag",
		Tag:           "bytes,847940,opt,name=openapiv2_tag",
		Filename:      "openapiv2/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Schema)(nil),
		Field:         847940,
		Name:          "micro.openapiv2.openapiv2_field",
		Tag:           "bytes,847940,opt,name=openapiv2_field",
		Filename:      "openapiv2/annotations.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.Document openapiv2_swagger = 847940;
	E_Openapiv2Swagger = &file_openapiv2_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.Operation openapiv2_operation = 847940;
	E_Openapiv2Operation = &file_openapiv2_annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.Schema openapiv2_schema = 847940;
	E_Openapiv2Schema = &file_openapiv2_annotations_proto_extTypes[2]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.Tag openapiv2_tag = 847940;
	E_Openapiv2Tag = &file_openapiv2_annotations_proto_extTypes[3]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.Schema openapiv2_field = 847940;
	E_Openapiv2Field = &file_openapiv2_annotations_proto_extTypes[4]
)

var File_openapiv2_annotations_proto protoreflect.FileDescriptor

var file_openapiv2_annotations_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x66, 0x0a, 0x11, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x5f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc4,
	0xe0, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x10, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x53, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x3a, 0x6d, 0x0a, 0x13, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc4, 0xe0, 0x33, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x65, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc4, 0xe0, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3a, 0x5c, 0x0a, 0x0d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x5f, 0x74, 0x61, 0x67, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc4, 0xe0, 0x33, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x54, 0x61, 0x67, 0x3a, 0x61, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc4, 0xe0, 0x33, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x6f,
	0x2e, 0x75, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x3b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_openapiv2_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.FileOptions)(nil),    // 0: google.protobuf.FileOptions
	(*descriptorpb.MethodOptions)(nil),  // 1: google.protobuf.MethodOptions
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptorpb.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
	(*descriptorpb.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
	(*Document)(nil),                    // 5: micro.openapiv2.Document
	(*Operation)(nil),                   // 6: micro.openapiv2.Operation
	(*Schema)(nil),                      // 7: micro.openapiv2.Schema
	(*Tag)(nil),                         // 8: micro.openapiv2.Tag
}
var file_openapiv2_annotations_proto_depIdxs = []int32{
	0,  // 0: micro.openapiv2.openapiv2_swagger:extendee -> google.protobuf.FileOptions
	1,  // 1: micro.openapiv2.openapiv2_operation:extendee -> google.protobuf.MethodOptions
	2,  // 2: micro.openapiv2.openapiv2_schema:extendee -> google.protobuf.MessageOptions
	3,  // 3: micro.openapiv2.openapiv2_tag:extendee -> google.protobuf.ServiceOptions
	4,  // 4: micro.openapiv2.openapiv2_field:extendee -> google.protobuf.FieldOptions
	5,  // 5: micro.openapiv2.openapiv2_swagger:type_name -> micro.openapiv2.Document
	6,  // 6: micro.openapiv2.openapiv2_operation:type_name -> micro.openapiv2.Operation
	7,  // 7: micro.openapiv2.openapiv2_schema:type_name -> micro.openapiv2.Schema
	8,  // 8: micro.openapiv2.openapiv2_tag:type_name -> micro.openapiv2.Tag
	7,  // 9: micro.openapiv2.openapiv2_field:type_name -> micro.openapiv2.Schema
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	5,  // [5:10] is the sub-list for extension type_name
	0,  // [0:5] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_openapiv2_annotations_proto_init() }
func file_openapiv2_annotations_proto_init() {
	if File_openapiv2_annotations_proto != nil {
		return
	}
	file_openapiv2_openapiv2_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_openapiv2_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_openapiv2_annotations_proto_goTypes,
		DependencyIndexes: file_openapiv2_annotations_proto_depIdxs,
		ExtensionInfos:    file_openapiv2_annotations_proto_extTypes,
	}.Build()
	File_openapiv2_annotations_proto = out.File
	file_openapiv2_annotations_proto_rawDesc = nil
	file_openapiv2_annotations_proto_goTypes = nil
	file_openapiv2_annotations_proto_depIdxs = nil
}
