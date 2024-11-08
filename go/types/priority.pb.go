// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: priority.proto

package types

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

type Priority int32

const (
	Priority_P0 Priority = 0
	Priority_P1 Priority = 1
	Priority_P2 Priority = 2
	Priority_P3 Priority = 3
	Priority_P4 Priority = 4
	Priority_P5 Priority = 5
	Priority_P6 Priority = 6
	Priority_P7 Priority = 7
)

// Enum value maps for Priority.
var (
	Priority_name = map[int32]string{
		0: "P0",
		1: "P1",
		2: "P2",
		3: "P3",
		4: "P4",
		5: "P5",
		6: "P6",
		7: "P7",
	}
	Priority_value = map[string]int32{
		"P0": 0,
		"P1": 1,
		"P2": 2,
		"P3": 3,
		"P4": 4,
		"P5": 5,
		"P6": 6,
		"P7": 7,
	}
)

func (x Priority) Enum() *Priority {
	p := new(Priority)
	*p = x
	return p
}

func (x Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_priority_proto_enumTypes[0].Descriptor()
}

func (Priority) Type() protoreflect.EnumType {
	return &file_priority_proto_enumTypes[0]
}

func (x Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Priority.Descriptor instead.
func (Priority) EnumDescriptor() ([]byte, []int) {
	return file_priority_proto_rawDescGZIP(), []int{0}
}

var File_priority_proto protoreflect.FileDescriptor

var file_priority_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2a, 0x4a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x30, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x50,
	0x31, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x32, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x50,
	0x33, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x34, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x50,
	0x35, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x36, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x50,
	0x37, 0x10, 0x07, 0x42, 0x24, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01,
	0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_priority_proto_rawDescOnce sync.Once
	file_priority_proto_rawDescData = file_priority_proto_rawDesc
)

func file_priority_proto_rawDescGZIP() []byte {
	file_priority_proto_rawDescOnce.Do(func() {
		file_priority_proto_rawDescData = protoimpl.X.CompressGZIP(file_priority_proto_rawDescData)
	})
	return file_priority_proto_rawDescData
}

var file_priority_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_priority_proto_goTypes = []interface{}{
	(Priority)(0), // 0: types.Priority
}
var file_priority_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_priority_proto_init() }
func file_priority_proto_init() {
	if File_priority_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_priority_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_priority_proto_goTypes,
		DependencyIndexes: file_priority_proto_depIdxs,
		EnumInfos:         file_priority_proto_enumTypes,
	}.Build()
	File_priority_proto = out.File
	file_priority_proto_rawDesc = nil
	file_priority_proto_goTypes = nil
	file_priority_proto_depIdxs = nil
}