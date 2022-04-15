// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: enums/withdrawal_status.proto

package enums

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

type WithdrawalStatus int32

const (
	WithdrawalStatus_WITHDRAWAL_STATUS_UNSPECIFIED WithdrawalStatus = 0
	WithdrawalStatus_WITHDRAWAL_STATUS_CREATED     WithdrawalStatus = 1
	WithdrawalStatus_WITHDRAWAL_STATUS_RUNNING     WithdrawalStatus = 2
	WithdrawalStatus_WITHDRAWAL_STATUS_FAILED      WithdrawalStatus = 3
	WithdrawalStatus_WITHDRAWAL_STATUS_CANCELED    WithdrawalStatus = 4
	WithdrawalStatus_WITHDRAWAL_STATUS_COMPLETED   WithdrawalStatus = 5
)

// Enum value maps for WithdrawalStatus.
var (
	WithdrawalStatus_name = map[int32]string{
		0: "WITHDRAWAL_STATUS_UNSPECIFIED",
		1: "WITHDRAWAL_STATUS_CREATED",
		2: "WITHDRAWAL_STATUS_RUNNING",
		3: "WITHDRAWAL_STATUS_FAILED",
		4: "WITHDRAWAL_STATUS_CANCELED",
		5: "WITHDRAWAL_STATUS_COMPLETED",
	}
	WithdrawalStatus_value = map[string]int32{
		"WITHDRAWAL_STATUS_UNSPECIFIED": 0,
		"WITHDRAWAL_STATUS_CREATED":     1,
		"WITHDRAWAL_STATUS_RUNNING":     2,
		"WITHDRAWAL_STATUS_FAILED":      3,
		"WITHDRAWAL_STATUS_CANCELED":    4,
		"WITHDRAWAL_STATUS_COMPLETED":   5,
	}
)

func (x WithdrawalStatus) Enum() *WithdrawalStatus {
	p := new(WithdrawalStatus)
	*p = x
	return p
}

func (x WithdrawalStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WithdrawalStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_withdrawal_status_proto_enumTypes[0].Descriptor()
}

func (WithdrawalStatus) Type() protoreflect.EnumType {
	return &file_enums_withdrawal_status_proto_enumTypes[0]
}

func (x WithdrawalStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WithdrawalStatus.Descriptor instead.
func (WithdrawalStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_withdrawal_status_proto_rawDescGZIP(), []int{0}
}

var File_enums_withdrawal_status_proto protoreflect.FileDescriptor

var file_enums_withdrawal_status_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x70, 0x65, 0x70, 0x65, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x6b, 0x69, 0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0xd2, 0x01, 0x0a,
	0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x21, 0x0a, 0x1d, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57,
	0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41,
	0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x1e, 0x0a, 0x1a, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04,
	0x12, 0x1f, 0x0a, 0x1b, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x05, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x65, 0x70, 0x65, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x6b, 0x69,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_withdrawal_status_proto_rawDescOnce sync.Once
	file_enums_withdrawal_status_proto_rawDescData = file_enums_withdrawal_status_proto_rawDesc
)

func file_enums_withdrawal_status_proto_rawDescGZIP() []byte {
	file_enums_withdrawal_status_proto_rawDescOnce.Do(func() {
		file_enums_withdrawal_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_withdrawal_status_proto_rawDescData)
	})
	return file_enums_withdrawal_status_proto_rawDescData
}

var file_enums_withdrawal_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_withdrawal_status_proto_goTypes = []interface{}{
	(WithdrawalStatus)(0), // 0: pepeunlimited.grpckit.enums.WithdrawalStatus
}
var file_enums_withdrawal_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_withdrawal_status_proto_init() }
func file_enums_withdrawal_status_proto_init() {
	if File_enums_withdrawal_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_withdrawal_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_withdrawal_status_proto_goTypes,
		DependencyIndexes: file_enums_withdrawal_status_proto_depIdxs,
		EnumInfos:         file_enums_withdrawal_status_proto_enumTypes,
	}.Build()
	File_enums_withdrawal_status_proto = out.File
	file_enums_withdrawal_status_proto_rawDesc = nil
	file_enums_withdrawal_status_proto_goTypes = nil
	file_enums_withdrawal_status_proto_depIdxs = nil
}
