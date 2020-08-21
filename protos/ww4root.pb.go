// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: ww4root.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Type int32

const (
	Type_NONE  Type = 0
	Type_BUILD Type = 1
	Type_ARMY  Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "NONE",
		1: "BUILD",
		2: "ARMY",
	}
	Type_value = map[string]int32{
		"NONE":  0,
		"BUILD": 1,
		"ARMY":  2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_ww4root_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_ww4root_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_ww4root_proto_rawDescGZIP(), []int{0}
}

type SubType int32

const (
	SubType_CityHall        SubType = 0
	SubType_GoldMine        SubType = 1
	SubType_Hospital        SubType = 16
	SubType_ARMY_TYPE_LEVEL SubType = 3 //士兵的类型
)

// Enum value maps for SubType.
var (
	SubType_name = map[int32]string{
		0:  "CityHall",
		1:  "GoldMine",
		16: "Hospital",
		3:  "ARMY_TYPE_LEVEL",
	}
	SubType_value = map[string]int32{
		"CityHall":        0,
		"GoldMine":        1,
		"Hospital":        16,
		"ARMY_TYPE_LEVEL": 3,
	}
)

func (x SubType) Enum() *SubType {
	p := new(SubType)
	*p = x
	return p
}

func (x SubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubType) Descriptor() protoreflect.EnumDescriptor {
	return file_ww4root_proto_enumTypes[1].Descriptor()
}

func (SubType) Type() protoreflect.EnumType {
	return &file_ww4root_proto_enumTypes[1]
}

func (x SubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubType.Descriptor instead.
func (SubType) EnumDescriptor() ([]byte, []int) {
	return file_ww4root_proto_rawDescGZIP(), []int{1}
}

type NormalResponseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *NormalResponseResult) Reset() {
	*x = NormalResponseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ww4root_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalResponseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalResponseResult) ProtoMessage() {}

func (x *NormalResponseResult) ProtoReflect() protoreflect.Message {
	mi := &file_ww4root_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalResponseResult.ProtoReflect.Descriptor instead.
func (*NormalResponseResult) Descriptor() ([]byte, []int) {
	return file_ww4root_proto_rawDescGZIP(), []int{0}
}

func (x *NormalResponseResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *NormalResponseResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Type     `protobuf:"varint,1,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	Data []string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"` // [param1,param2]
	Ext  string   `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ww4root_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_ww4root_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_ww4root_proto_rawDescGZIP(), []int{1}
}

func (x *Param) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_NONE
}

func (x *Param) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Param) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type ResponseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*Data `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Ext  string  `protobuf:"bytes,4,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *ResponseResult) Reset() {
	*x = ResponseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ww4root_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseResult) ProtoMessage() {}

func (x *ResponseResult) ProtoReflect() protoreflect.Message {
	mi := &file_ww4root_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseResult.ProtoReflect.Descriptor instead.
func (*ResponseResult) Descriptor() ([]byte, []int) {
	return file_ww4root_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseResult) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ResponseResult) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ResponseResult) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    Type     `protobuf:"varint,1,opt,name=type,proto3,enum=Type" json:"type,omitempty"` // data 类型 building army
	SubType string   `protobuf:"bytes,2,opt,name=subType,proto3" json:"subType,omitempty"`      // 具体类型 cityhall army的类型
	Id      string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`                // buildingid or ArmyId
	Params  []string `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty"`        // 自定义返回参数 [topos/count,lv]
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ww4root_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_ww4root_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_ww4root_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_NONE
}

func (x *Data) GetSubType() string {
	if x != nil {
		return x.SubType
	}
	return ""
}

func (x *Data) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Data) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_ww4root_proto protoreflect.FileDescriptor

var file_ww4root_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x77, 0x34, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3c, 0x0a, 0x14, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x48, 0x0a,
	0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x22, 0x63, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x19, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x22, 0x63, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x2a, 0x25, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x41, 0x52, 0x4d, 0x59, 0x10, 0x02, 0x2a, 0x48, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x48, 0x61, 0x6c, 0x6c, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x6f, 0x6c, 0x64, 0x4d, 0x69, 0x6e, 0x65, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x10, 0x10, 0x12, 0x13, 0x0a,
	0x0f, 0x41, 0x52, 0x4d, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x10, 0x03, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ww4root_proto_rawDescOnce sync.Once
	file_ww4root_proto_rawDescData = file_ww4root_proto_rawDesc
)

func file_ww4root_proto_rawDescGZIP() []byte {
	file_ww4root_proto_rawDescOnce.Do(func() {
		file_ww4root_proto_rawDescData = protoimpl.X.CompressGZIP(file_ww4root_proto_rawDescData)
	})
	return file_ww4root_proto_rawDescData
}

var file_ww4root_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ww4root_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ww4root_proto_goTypes = []interface{}{
	(Type)(0),                    // 0: Type
	(SubType)(0),                 // 1: SubType
	(*NormalResponseResult)(nil), // 2: NormalResponseResult
	(*Param)(nil),                // 3: Param
	(*ResponseResult)(nil),       // 4: ResponseResult
	(*Data)(nil),                 // 5: Data
}
var file_ww4root_proto_depIdxs = []int32{
	0, // 0: Param.type:type_name -> Type
	5, // 1: ResponseResult.data:type_name -> Data
	0, // 2: Data.type:type_name -> Type
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ww4root_proto_init() }
func file_ww4root_proto_init() {
	if File_ww4root_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ww4root_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalResponseResult); i {
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
		file_ww4root_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
		file_ww4root_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseResult); i {
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
		file_ww4root_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
			RawDescriptor: file_ww4root_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ww4root_proto_goTypes,
		DependencyIndexes: file_ww4root_proto_depIdxs,
		EnumInfos:         file_ww4root_proto_enumTypes,
		MessageInfos:      file_ww4root_proto_msgTypes,
	}.Build()
	File_ww4root_proto = out.File
	file_ww4root_proto_rawDesc = nil
	file_ww4root_proto_goTypes = nil
	file_ww4root_proto_depIdxs = nil
}
