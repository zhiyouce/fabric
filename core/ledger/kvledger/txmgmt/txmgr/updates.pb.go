//
//Copyright IBM Corp. All Rights Reserved.
//
//SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: updates.proto

package txmgr

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KVWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace    string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Collection   string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
	Key          []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	IsDelete     bool   `protobuf:"varint,4,opt,name=isDelete,proto3" json:"isDelete,omitempty"`
	Value        []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	VersionBytes []byte `protobuf:"bytes,6,opt,name=version_bytes,json=versionBytes,proto3" json:"version_bytes,omitempty"`
}

func (x *KVWrite) Reset() {
	*x = KVWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVWrite) ProtoMessage() {}

func (x *KVWrite) ProtoReflect() protoreflect.Message {
	mi := &file_updates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVWrite.ProtoReflect.Descriptor instead.
func (*KVWrite) Descriptor() ([]byte, []int) {
	return file_updates_proto_rawDescGZIP(), []int{0}
}

func (x *KVWrite) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KVWrite) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *KVWrite) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KVWrite) GetIsDelete() bool {
	if x != nil {
		return x.IsDelete
	}
	return false
}

func (x *KVWrite) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KVWrite) GetVersionBytes() []byte {
	if x != nil {
		return x.VersionBytes
	}
	return nil
}

type Updates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kvwrites []*KVWrite `protobuf:"bytes,1,rep,name=kvwrites,proto3" json:"kvwrites,omitempty"`
}

func (x *Updates) Reset() {
	*x = Updates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Updates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Updates) ProtoMessage() {}

func (x *Updates) ProtoReflect() protoreflect.Message {
	mi := &file_updates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Updates.ProtoReflect.Descriptor instead.
func (*Updates) Descriptor() ([]byte, []int) {
	return file_updates_proto_rawDescGZIP(), []int{1}
}

func (x *Updates) GetKvwrites() []*KVWrite {
	if x != nil {
		return x.Kvwrites
	}
	return nil
}

var File_updates_proto protoreflect.FileDescriptor

var file_updates_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x78, 0x6d, 0x67, 0x72, 0x22, 0xb0, 0x01, 0x0a, 0x07, 0x4b, 0x56, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x07, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x6b, 0x76, 0x77, 0x72, 0x69, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x78, 0x6d, 0x67, 0x72, 0x2e, 0x4b,
	0x56, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x08, 0x6b, 0x76, 0x77, 0x72, 0x69, 0x74, 0x65, 0x73,
	0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69,
	0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6b, 0x76,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x78, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x74, 0x78,
	0x6d, 0x67, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_updates_proto_rawDescOnce sync.Once
	file_updates_proto_rawDescData = file_updates_proto_rawDesc
)

func file_updates_proto_rawDescGZIP() []byte {
	file_updates_proto_rawDescOnce.Do(func() {
		file_updates_proto_rawDescData = protoimpl.X.CompressGZIP(file_updates_proto_rawDescData)
	})
	return file_updates_proto_rawDescData
}

var file_updates_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_updates_proto_goTypes = []any{
	(*KVWrite)(nil), // 0: txmgr.KVWrite
	(*Updates)(nil), // 1: txmgr.Updates
}
var file_updates_proto_depIdxs = []int32{
	0, // 0: txmgr.Updates.kvwrites:type_name -> txmgr.KVWrite
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_updates_proto_init() }
func file_updates_proto_init() {
	if File_updates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_updates_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KVWrite); i {
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
		file_updates_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Updates); i {
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
			RawDescriptor: file_updates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_updates_proto_goTypes,
		DependencyIndexes: file_updates_proto_depIdxs,
		MessageInfos:      file_updates_proto_msgTypes,
	}.Build()
	File_updates_proto = out.File
	file_updates_proto_rawDesc = nil
	file_updates_proto_goTypes = nil
	file_updates_proto_depIdxs = nil
}
