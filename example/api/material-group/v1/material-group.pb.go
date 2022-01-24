// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/material-group/v1/material-group.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SaveMaterialGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrgId  string                  `protobuf:"bytes,2,opt,name=orgId,proto3" json:"orgId,omitempty"`
	UserId *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"` // optional 如果没有，则表示是企业话术组
	Name   string                  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Type   *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Scope  *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
	Order  int64                   `protobuf:"varint,7,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *SaveMaterialGroupRequest) Reset() {
	*x = SaveMaterialGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_material_group_v1_material_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveMaterialGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMaterialGroupRequest) ProtoMessage() {}

func (x *SaveMaterialGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_material_group_v1_material_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveMaterialGroupRequest.ProtoReflect.Descriptor instead.
func (*SaveMaterialGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_material_group_v1_material_group_proto_rawDescGZIP(), []int{0}
}

func (x *SaveMaterialGroupRequest) GetId() *wrapperspb.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SaveMaterialGroupRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *SaveMaterialGroupRequest) GetUserId() *wrapperspb.StringValue {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *SaveMaterialGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SaveMaterialGroupRequest) GetType() *wrapperspb.StringValue {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *SaveMaterialGroupRequest) GetScope() *wrapperspb.StringValue {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *SaveMaterialGroupRequest) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

type MaterialGroupModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type  string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Order int64  `protobuf:"varint,4,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *MaterialGroupModel) Reset() {
	*x = MaterialGroupModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_material_group_v1_material_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialGroupModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialGroupModel) ProtoMessage() {}

func (x *MaterialGroupModel) ProtoReflect() protoreflect.Message {
	mi := &file_api_material_group_v1_material_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialGroupModel.ProtoReflect.Descriptor instead.
func (*MaterialGroupModel) Descriptor() ([]byte, []int) {
	return file_api_material_group_v1_material_group_proto_rawDescGZIP(), []int{1}
}

func (x *MaterialGroupModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaterialGroupModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MaterialGroupModel) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MaterialGroupModel) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

type MaterialGroupModelList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*MaterialGroupModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MaterialGroupModelList) Reset() {
	*x = MaterialGroupModelList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_material_group_v1_material_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialGroupModelList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialGroupModelList) ProtoMessage() {}

func (x *MaterialGroupModelList) ProtoReflect() protoreflect.Message {
	mi := &file_api_material_group_v1_material_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialGroupModelList.ProtoReflect.Descriptor instead.
func (*MaterialGroupModelList) Descriptor() ([]byte, []int) {
	return file_api_material_group_v1_material_group_proto_rawDescGZIP(), []int{2}
}

func (x *MaterialGroupModelList) GetData() []*MaterialGroupModel {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_material_group_v1_material_group_proto protoreflect.FileDescriptor

var file_api_material_group_v1_material_group_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2d, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a, 0x18, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x12, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x16, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x63, 0x0a, 0x0d, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x52, 0x0a, 0x04, 0x53, 0x61,
	0x76, 0x65, 0x12, 0x27, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x3a,
	0x5a, 0x38, 0x6a, 0x75, 0x7a, 0x69, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6d, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_material_group_v1_material_group_proto_rawDescOnce sync.Once
	file_api_material_group_v1_material_group_proto_rawDescData = file_api_material_group_v1_material_group_proto_rawDesc
)

func file_api_material_group_v1_material_group_proto_rawDescGZIP() []byte {
	file_api_material_group_v1_material_group_proto_rawDescOnce.Do(func() {
		file_api_material_group_v1_material_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_material_group_v1_material_group_proto_rawDescData)
	})
	return file_api_material_group_v1_material_group_proto_rawDescData
}

var file_api_material_group_v1_material_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_material_group_v1_material_group_proto_goTypes = []interface{}{
	(*SaveMaterialGroupRequest)(nil), // 0: materialGroup.SaveMaterialGroupRequest
	(*MaterialGroupModel)(nil),       // 1: materialGroup.MaterialGroupModel
	(*MaterialGroupModelList)(nil),   // 2: materialGroup.MaterialGroupModelList
	(*wrapperspb.StringValue)(nil),   // 3: google.protobuf.StringValue
}
var file_api_material_group_v1_material_group_proto_depIdxs = []int32{
	3, // 0: materialGroup.SaveMaterialGroupRequest.id:type_name -> google.protobuf.StringValue
	3, // 1: materialGroup.SaveMaterialGroupRequest.userId:type_name -> google.protobuf.StringValue
	3, // 2: materialGroup.SaveMaterialGroupRequest.type:type_name -> google.protobuf.StringValue
	3, // 3: materialGroup.SaveMaterialGroupRequest.scope:type_name -> google.protobuf.StringValue
	1, // 4: materialGroup.MaterialGroupModelList.data:type_name -> materialGroup.MaterialGroupModel
	0, // 5: materialGroup.MaterialGroup.Save:input_type -> materialGroup.SaveMaterialGroupRequest
	1, // 6: materialGroup.MaterialGroup.Save:output_type -> materialGroup.MaterialGroupModel
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_material_group_v1_material_group_proto_init() }
func file_api_material_group_v1_material_group_proto_init() {
	if File_api_material_group_v1_material_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_material_group_v1_material_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveMaterialGroupRequest); i {
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
		file_api_material_group_v1_material_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialGroupModel); i {
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
		file_api_material_group_v1_material_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterialGroupModelList); i {
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
			RawDescriptor: file_api_material_group_v1_material_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_material_group_v1_material_group_proto_goTypes,
		DependencyIndexes: file_api_material_group_v1_material_group_proto_depIdxs,
		MessageInfos:      file_api_material_group_v1_material_group_proto_msgTypes,
	}.Build()
	File_api_material_group_v1_material_group_proto = out.File
	file_api_material_group_v1_material_group_proto_rawDesc = nil
	file_api_material_group_v1_material_group_proto_goTypes = nil
	file_api_material_group_v1_material_group_proto_depIdxs = nil
}