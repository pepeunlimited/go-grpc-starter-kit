// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: resources/todo.proto

package resources

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

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsDone    bool   `protobuf:"varint,5,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_resources_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_resources_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Todo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Todo) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

var File_resources_todo_proto protoreflect.FileDescriptor

var file_resources_todo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x70, 0x65, 0x70, 0x65, 0x75, 0x6e, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x64, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6b, 0x69, 0x74, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x70, 0x65,
	0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_todo_proto_rawDescOnce sync.Once
	file_resources_todo_proto_rawDescData = file_resources_todo_proto_rawDesc
)

func file_resources_todo_proto_rawDescGZIP() []byte {
	file_resources_todo_proto_rawDescOnce.Do(func() {
		file_resources_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_todo_proto_rawDescData)
	})
	return file_resources_todo_proto_rawDescData
}

var file_resources_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_todo_proto_goTypes = []interface{}{
	(*Todo)(nil), // 0: pepeunlimited.grpckit.resources.Todo
}
var file_resources_todo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resources_todo_proto_init() }
func file_resources_todo_proto_init() {
	if File_resources_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
			RawDescriptor: file_resources_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_todo_proto_goTypes,
		DependencyIndexes: file_resources_todo_proto_depIdxs,
		MessageInfos:      file_resources_todo_proto_msgTypes,
	}.Build()
	File_resources_todo_proto = out.File
	file_resources_todo_proto_rawDesc = nil
	file_resources_todo_proto_goTypes = nil
	file_resources_todo_proto_depIdxs = nil
}
