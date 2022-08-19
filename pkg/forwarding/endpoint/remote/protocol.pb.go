// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: forwarding/endpoint/remote/protocol.proto

package remote

import (
	forwarding "github.com/mutagen-io/mutagen/pkg/forwarding"
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

// InitializeForwardingRequest is the initialization request sent to remote
// forwarding endpoints.
type InitializeForwardingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Field 1 is reserved for the session identifier.
	// Version is the session version.
	Version forwarding.Version `protobuf:"varint,2,opt,name=version,proto3,enum=forwarding.Version" json:"version,omitempty"`
	// Configuration is the session configuration.
	Configuration *forwarding.Configuration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Protocol is the protocol specification for the listener/dialer.
	Protocol string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Address it the bind address for a listener or dial address for a dialer.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Listener indicates whether this endpoint should function as a listener or
	// dialer for the associated address.
	Listener bool `protobuf:"varint,6,opt,name=listener,proto3" json:"listener,omitempty"`
}

func (x *InitializeForwardingRequest) Reset() {
	*x = InitializeForwardingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarding_endpoint_remote_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeForwardingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeForwardingRequest) ProtoMessage() {}

func (x *InitializeForwardingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forwarding_endpoint_remote_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeForwardingRequest.ProtoReflect.Descriptor instead.
func (*InitializeForwardingRequest) Descriptor() ([]byte, []int) {
	return file_forwarding_endpoint_remote_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeForwardingRequest) GetVersion() forwarding.Version {
	if x != nil {
		return x.Version
	}
	return forwarding.Version(0)
}

func (x *InitializeForwardingRequest) GetConfiguration() *forwarding.Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *InitializeForwardingRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *InitializeForwardingRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InitializeForwardingRequest) GetListener() bool {
	if x != nil {
		return x.Listener
	}
	return false
}

// InitializeForwardingResponse is the initialization response sent from remote
// forwarding endpoint.
type InitializeForwardingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error is any error that occurred during initialization.
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *InitializeForwardingResponse) Reset() {
	*x = InitializeForwardingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarding_endpoint_remote_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeForwardingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeForwardingResponse) ProtoMessage() {}

func (x *InitializeForwardingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forwarding_endpoint_remote_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeForwardingResponse.ProtoReflect.Descriptor instead.
func (*InitializeForwardingResponse) Descriptor() ([]byte, []int) {
	return file_forwarding_endpoint_remote_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *InitializeForwardingResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_forwarding_endpoint_remote_protocol_proto protoreflect.FileDescriptor

var file_forwarding_endpoint_remote_protocol_proto_rawDesc = []byte{
	0x0a, 0x29, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x1a, 0x1e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01,
	0x0a, 0x1b, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x22,
	0x34, 0x0a, 0x1c, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d,
	0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_forwarding_endpoint_remote_protocol_proto_rawDescOnce sync.Once
	file_forwarding_endpoint_remote_protocol_proto_rawDescData = file_forwarding_endpoint_remote_protocol_proto_rawDesc
)

func file_forwarding_endpoint_remote_protocol_proto_rawDescGZIP() []byte {
	file_forwarding_endpoint_remote_protocol_proto_rawDescOnce.Do(func() {
		file_forwarding_endpoint_remote_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_forwarding_endpoint_remote_protocol_proto_rawDescData)
	})
	return file_forwarding_endpoint_remote_protocol_proto_rawDescData
}

var file_forwarding_endpoint_remote_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_forwarding_endpoint_remote_protocol_proto_goTypes = []interface{}{
	(*InitializeForwardingRequest)(nil),  // 0: remote.InitializeForwardingRequest
	(*InitializeForwardingResponse)(nil), // 1: remote.InitializeForwardingResponse
	(forwarding.Version)(0),              // 2: forwarding.Version
	(*forwarding.Configuration)(nil),     // 3: forwarding.Configuration
}
var file_forwarding_endpoint_remote_protocol_proto_depIdxs = []int32{
	2, // 0: remote.InitializeForwardingRequest.version:type_name -> forwarding.Version
	3, // 1: remote.InitializeForwardingRequest.configuration:type_name -> forwarding.Configuration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_forwarding_endpoint_remote_protocol_proto_init() }
func file_forwarding_endpoint_remote_protocol_proto_init() {
	if File_forwarding_endpoint_remote_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_forwarding_endpoint_remote_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeForwardingRequest); i {
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
		file_forwarding_endpoint_remote_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeForwardingResponse); i {
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
			RawDescriptor: file_forwarding_endpoint_remote_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_forwarding_endpoint_remote_protocol_proto_goTypes,
		DependencyIndexes: file_forwarding_endpoint_remote_protocol_proto_depIdxs,
		MessageInfos:      file_forwarding_endpoint_remote_protocol_proto_msgTypes,
	}.Build()
	File_forwarding_endpoint_remote_protocol_proto = out.File
	file_forwarding_endpoint_remote_protocol_proto_rawDesc = nil
	file_forwarding_endpoint_remote_protocol_proto_goTypes = nil
	file_forwarding_endpoint_remote_protocol_proto_depIdxs = nil
}
