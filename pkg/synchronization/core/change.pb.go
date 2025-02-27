// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: synchronization/core/change.proto

package core

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

// Change encodes a change to an entry hierarchy. Change objects should be
// considered immutable and must not be modified.
type Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path is the path of the root of the change (relative to the
	// synchronization root).
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Old represents the old filesystem hierarchy at the change path. It may be
	// nil if no content previously existed.
	Old *Entry `protobuf:"bytes,2,opt,name=old,proto3" json:"old,omitempty"`
	// New represents the new filesystem hierarchy at the change path. It may be
	// nil if content has been deleted.
	New *Entry `protobuf:"bytes,3,opt,name=new,proto3" json:"new,omitempty"`
}

func (x *Change) Reset() {
	*x = Change{}
	mi := &file_synchronization_core_change_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Change) ProtoMessage() {}

func (x *Change) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_core_change_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Change.ProtoReflect.Descriptor instead.
func (*Change) Descriptor() ([]byte, []int) {
	return file_synchronization_core_change_proto_rawDescGZIP(), []int{0}
}

func (x *Change) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Change) GetOld() *Entry {
	if x != nil {
		return x.Old
	}
	return nil
}

func (x *Change) GetNew() *Entry {
	if x != nil {
		return x.New
	}
	return nil
}

var File_synchronization_core_change_proto protoreflect.FileDescriptor

var file_synchronization_core_change_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x68,
	0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x06, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x03, 0x6f, 0x6c, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x03, 0x6f, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f,
	0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_core_change_proto_rawDescOnce sync.Once
	file_synchronization_core_change_proto_rawDescData = file_synchronization_core_change_proto_rawDesc
)

func file_synchronization_core_change_proto_rawDescGZIP() []byte {
	file_synchronization_core_change_proto_rawDescOnce.Do(func() {
		file_synchronization_core_change_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_core_change_proto_rawDescData)
	})
	return file_synchronization_core_change_proto_rawDescData
}

var file_synchronization_core_change_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_synchronization_core_change_proto_goTypes = []any{
	(*Change)(nil), // 0: core.Change
	(*Entry)(nil),  // 1: core.Entry
}
var file_synchronization_core_change_proto_depIdxs = []int32{
	1, // 0: core.Change.old:type_name -> core.Entry
	1, // 1: core.Change.new:type_name -> core.Entry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_synchronization_core_change_proto_init() }
func file_synchronization_core_change_proto_init() {
	if File_synchronization_core_change_proto != nil {
		return
	}
	file_synchronization_core_entry_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synchronization_core_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_core_change_proto_goTypes,
		DependencyIndexes: file_synchronization_core_change_proto_depIdxs,
		MessageInfos:      file_synchronization_core_change_proto_msgTypes,
	}.Build()
	File_synchronization_core_change_proto = out.File
	file_synchronization_core_change_proto_rawDesc = nil
	file_synchronization_core_change_proto_goTypes = nil
	file_synchronization_core_change_proto_depIdxs = nil
}
