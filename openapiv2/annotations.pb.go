// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: openapiv2/annotations.proto

package openapiv2

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var file_openapiv2_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*Swagger)(nil),
		Field:         1042,
		Name:          "micro.openapiv2.openapiv2_swagger",
		Tag:           "bytes,1042,opt,name=openapiv2_swagger",
		Filename:      "openapiv2/annotations.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*Operation)(nil),
		Field:         1042,
		Name:          "micro.openapiv2.openapiv2_operation",
		Tag:           "bytes,1042,opt,name=openapiv2_operation",
		Filename:      "openapiv2/annotations.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*Schema)(nil),
		Field:         1042,
		Name:          "micro.openapiv2.openapiv2_schema",
		Tag:           "bytes,1042,opt,name=openapiv2_schema",
		Filename:      "openapiv2/annotations.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*Tag)(nil),
		Field:         1042,
		Name:          "micro.openapiv2.openapiv2_tag",
		Tag:           "bytes,1042,opt,name=openapiv2_tag",
		Filename:      "openapiv2/annotations.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*JSONSchema)(nil),
		Field:         1042,
		Name:          "micro.openapiv2.openapiv2_field",
		Tag:           "bytes,1042,opt,name=openapiv2_field",
		Filename:      "openapiv2/annotations.proto",
	},
}

// Extension fields to descriptor.FileOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.Swagger openapiv2_swagger = 1042;
	E_Openapiv2Swagger = &file_openapiv2_annotations_proto_extTypes[0]
)

// Extension fields to descriptor.MethodOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.Operation openapiv2_operation = 1042;
	E_Openapiv2Operation = &file_openapiv2_annotations_proto_extTypes[1]
)

// Extension fields to descriptor.MessageOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.Schema openapiv2_schema = 1042;
	E_Openapiv2Schema = &file_openapiv2_annotations_proto_extTypes[2]
)

// Extension fields to descriptor.ServiceOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.Tag openapiv2_tag = 1042;
	E_Openapiv2Tag = &file_openapiv2_annotations_proto_extTypes[3]
)

// Extension fields to descriptor.FieldOptions.
var (
	// ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project.
	//
	// All IDs are the same, as assigned. It is okay that they are the same, as they extend
	// different descriptor messages.
	//
	// optional micro.openapiv2.JSONSchema openapiv2_field = 1042;
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
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x64, 0x0a, 0x11, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x5f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x92,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x10, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x3a, 0x6c, 0x0a, 0x13, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x92, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x64, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x92, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x3a, 0x5b, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x5f, 0x74, 0x61, 0x67, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x92, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2e, 0x54, 0x61, 0x67, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x54,
	0x61, 0x67, 0x3a, 0x64, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x92, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2e, 0x4a, 0x53,
	0x4f, 0x4e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2d,
	0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x3b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_openapiv2_annotations_proto_goTypes = []interface{}{
	(*descriptor.FileOptions)(nil),    // 0: google.protobuf.FileOptions
	(*descriptor.MethodOptions)(nil),  // 1: google.protobuf.MethodOptions
	(*descriptor.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptor.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
	(*descriptor.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
	(*Swagger)(nil),                   // 5: micro.openapiv2.Swagger
	(*Operation)(nil),                 // 6: micro.openapiv2.Operation
	(*Schema)(nil),                    // 7: micro.openapiv2.Schema
	(*Tag)(nil),                       // 8: micro.openapiv2.Tag
	(*JSONSchema)(nil),                // 9: micro.openapiv2.JSONSchema
}
var file_openapiv2_annotations_proto_depIdxs = []int32{
	0,  // 0: micro.openapiv2.openapiv2_swagger:extendee -> google.protobuf.FileOptions
	1,  // 1: micro.openapiv2.openapiv2_operation:extendee -> google.protobuf.MethodOptions
	2,  // 2: micro.openapiv2.openapiv2_schema:extendee -> google.protobuf.MessageOptions
	3,  // 3: micro.openapiv2.openapiv2_tag:extendee -> google.protobuf.ServiceOptions
	4,  // 4: micro.openapiv2.openapiv2_field:extendee -> google.protobuf.FieldOptions
	5,  // 5: micro.openapiv2.openapiv2_swagger:type_name -> micro.openapiv2.Swagger
	6,  // 6: micro.openapiv2.openapiv2_operation:type_name -> micro.openapiv2.Operation
	7,  // 7: micro.openapiv2.openapiv2_schema:type_name -> micro.openapiv2.Schema
	8,  // 8: micro.openapiv2.openapiv2_tag:type_name -> micro.openapiv2.Tag
	9,  // 9: micro.openapiv2.openapiv2_field:type_name -> micro.openapiv2.JSONSchema
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
