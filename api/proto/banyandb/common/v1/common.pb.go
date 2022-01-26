// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: banyandb/common/v1/common.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Catalog int32

const (
	Catalog_CATALOG_UNSPECIFIED Catalog = 0
	Catalog_CATALOG_STREAM      Catalog = 1
	Catalog_CATALOG_MEASURE     Catalog = 2
)

// Enum value maps for Catalog.
var (
	Catalog_name = map[int32]string{
		0: "CATALOG_UNSPECIFIED",
		1: "CATALOG_STREAM",
		2: "CATALOG_MEASURE",
	}
	Catalog_value = map[string]int32{
		"CATALOG_UNSPECIFIED": 0,
		"CATALOG_STREAM":      1,
		"CATALOG_MEASURE":     2,
	}
)

func (x Catalog) Enum() *Catalog {
	p := new(Catalog)
	*p = x
	return p
}

func (x Catalog) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Catalog) Descriptor() protoreflect.EnumDescriptor {
	return file_banyandb_common_v1_common_proto_enumTypes[0].Descriptor()
}

func (Catalog) Type() protoreflect.EnumType {
	return &file_banyandb_common_v1_common_proto_enumTypes[0]
}

func (x Catalog) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Catalog.Descriptor instead.
func (Catalog) EnumDescriptor() ([]byte, []int) {
	return file_banyandb_common_v1_common_proto_rawDescGZIP(), []int{0}
}

// Metadata is for multi-tenant, multi-model use
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// group contains a set of options, like retention policy, max
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// name of the entity
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id   uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// readonly. create_revision is the revision of last creation on this key.
	CreateRevision int64 `protobuf:"varint,4,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
	// readonly. mod_revision is the revision of last modification on this key.
	ModRevision int64 `protobuf:"varint,5,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_banyandb_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Metadata) GetCreateRevision() int64 {
	if x != nil {
		return x.CreateRevision
	}
	return 0
}

func (x *Metadata) GetModRevision() int64 {
	if x != nil {
		return x.ModRevision
	}
	return 0
}

// Group is an internal object for Group management
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the group
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// updated_at indicates when resources of the group are updated
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banyandb_common_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_banyandb_common_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_banyandb_common_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_banyandb_common_v1_common_proto protoreflect.FileDescriptor

var file_banyandb_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x5f, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x6f,
	0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x05, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x2a, 0x4b, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x41, 0x54, 0x41, 0x4c, 0x4f, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x54, 0x41, 0x4c, 0x4f, 0x47,
	0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x41, 0x54,
	0x41, 0x4c, 0x4f, 0x47, 0x5f, 0x4d, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x10, 0x02, 0x42, 0x6e,
	0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79,
	0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x64,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x79,
	0x61, 0x6e, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banyandb_common_v1_common_proto_rawDescOnce sync.Once
	file_banyandb_common_v1_common_proto_rawDescData = file_banyandb_common_v1_common_proto_rawDesc
)

func file_banyandb_common_v1_common_proto_rawDescGZIP() []byte {
	file_banyandb_common_v1_common_proto_rawDescOnce.Do(func() {
		file_banyandb_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_banyandb_common_v1_common_proto_rawDescData)
	})
	return file_banyandb_common_v1_common_proto_rawDescData
}

var file_banyandb_common_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_banyandb_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_banyandb_common_v1_common_proto_goTypes = []interface{}{
	(Catalog)(0),                  // 0: banyandb.common.v1.Catalog
	(*Metadata)(nil),              // 1: banyandb.common.v1.Metadata
	(*Group)(nil),                 // 2: banyandb.common.v1.Group
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_banyandb_common_v1_common_proto_depIdxs = []int32{
	3, // 0: banyandb.common.v1.Group.updated_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_banyandb_common_v1_common_proto_init() }
func file_banyandb_common_v1_common_proto_init() {
	if File_banyandb_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banyandb_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_banyandb_common_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
			RawDescriptor: file_banyandb_common_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banyandb_common_v1_common_proto_goTypes,
		DependencyIndexes: file_banyandb_common_v1_common_proto_depIdxs,
		EnumInfos:         file_banyandb_common_v1_common_proto_enumTypes,
		MessageInfos:      file_banyandb_common_v1_common_proto_msgTypes,
	}.Build()
	File_banyandb_common_v1_common_proto = out.File
	file_banyandb_common_v1_common_proto_rawDesc = nil
	file_banyandb_common_v1_common_proto_goTypes = nil
	file_banyandb_common_v1_common_proto_depIdxs = nil
}
