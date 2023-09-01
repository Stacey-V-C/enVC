// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: nvc.proto

package pb

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

type SQLtype int32

const (
	SQLtype_UNSPECIFIED SQLtype = 0
	SQLtype_CREATE      SQLtype = 1
	SQLtype_UPDATE      SQLtype = 2
	SQLtype_MIGRATE     SQLtype = 3
	SQLtype_DELETE      SQLtype = 4
)

// Enum value maps for SQLtype.
var (
	SQLtype_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "CREATE",
		2: "UPDATE",
		3: "MIGRATE",
		4: "DELETE",
	}
	SQLtype_value = map[string]int32{
		"UNSPECIFIED": 0,
		"CREATE":      1,
		"UPDATE":      2,
		"MIGRATE":     3,
		"DELETE":      4,
	}
)

func (x SQLtype) Enum() *SQLtype {
	p := new(SQLtype)
	*p = x
	return p
}

func (x SQLtype) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SQLtype) Descriptor() protoreflect.EnumDescriptor {
	return file_nvc_proto_enumTypes[0].Descriptor()
}

func (SQLtype) Type() protoreflect.EnumType {
	return &file_nvc_proto_enumTypes[0]
}

func (x SQLtype) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SQLtype.Descriptor instead.
func (SQLtype) EnumDescriptor() ([]byte, []int) {
	return file_nvc_proto_rawDescGZIP(), []int{0}
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_nvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_nvc_proto_rawDescGZIP(), []int{0}
}

func (x *Value) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type SQLAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      SQLtype  `protobuf:"varint,1,opt,name=type,proto3,enum=nvc.SQLtype" json:"type,omitempty"`
	Table     string   `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Columns   []string `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	Values    []*Value `protobuf:"bytes,12,rep,name=values,proto3" json:"values,omitempty"`
	DataModel string   `protobuf:"bytes,5,opt,name=dataModel,proto3" json:"dataModel,omitempty"` //   string valuesJson = 4;
}

func (x *SQLAction) Reset() {
	*x = SQLAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLAction) ProtoMessage() {}

func (x *SQLAction) ProtoReflect() protoreflect.Message {
	mi := &file_nvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLAction.ProtoReflect.Descriptor instead.
func (*SQLAction) Descriptor() ([]byte, []int) {
	return file_nvc_proto_rawDescGZIP(), []int{1}
}

func (x *SQLAction) GetType() SQLtype {
	if x != nil {
		return x.Type
	}
	return SQLtype_UNSPECIFIED
}

func (x *SQLAction) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *SQLAction) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *SQLAction) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *SQLAction) GetDataModel() string {
	if x != nil {
		return x.DataModel
	}
	return ""
}

type ChromaResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// might not even need to define this in message?
	// they allow sending status code but it might be good to have app logic failures too
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	//	int32 insertedCount = 2;
	//
	// Types that are assignable to Result:
	//
	//	*ChromaResult_InsertedCount
	//	*ChromaResult_Error
	Result isChromaResult_Result `protobuf_oneof:"result"`
}

func (x *ChromaResult) Reset() {
	*x = ChromaResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChromaResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChromaResult) ProtoMessage() {}

func (x *ChromaResult) ProtoReflect() protoreflect.Message {
	mi := &file_nvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChromaResult.ProtoReflect.Descriptor instead.
func (*ChromaResult) Descriptor() ([]byte, []int) {
	return file_nvc_proto_rawDescGZIP(), []int{2}
}

func (x *ChromaResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *ChromaResult) GetResult() isChromaResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ChromaResult) GetInsertedCount() int32 {
	if x, ok := x.GetResult().(*ChromaResult_InsertedCount); ok {
		return x.InsertedCount
	}
	return 0
}

func (x *ChromaResult) GetError() string {
	if x, ok := x.GetResult().(*ChromaResult_Error); ok {
		return x.Error
	}
	return ""
}

type isChromaResult_Result interface {
	isChromaResult_Result()
}

type ChromaResult_InsertedCount struct {
	InsertedCount int32 `protobuf:"varint,2,opt,name=insertedCount,proto3,oneof"`
}

type ChromaResult_Error struct {
	Error string `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*ChromaResult_InsertedCount) isChromaResult_Result() {}

func (*ChromaResult_Error) isChromaResult_Result() {}

var File_nvc_proto protoreflect.FileDescriptor

var file_nvc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x76, 0x63,
	0x22, 0x1f, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x22, 0x9f, 0x01, 0x0a, 0x09, 0x53, 0x51, 0x4c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x6e, 0x76, 0x63, 0x2e, 0x53, 0x51, 0x4c, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x12, 0x22, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x6e, 0x76, 0x63, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x72, 0x0a, 0x0c, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a,
	0x0d, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x4b, 0x0a, 0x07, 0x53, 0x51, 0x4c, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d,
	0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x04, 0x32, 0x4a, 0x0a, 0x13, 0x53, 0x51, 0x4c, 0x54, 0x6f, 0x43, 0x68, 0x72,
	0x6f, 0x6d, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x53, 0x51, 0x4c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x6e, 0x76,
	0x63, 0x2e, 0x53, 0x51, 0x4c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x6e, 0x76,
	0x63, 0x2e, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x65, 0x79, 0x2d, 0x76, 0x2d, 0x63, 0x2f, 0x6e, 0x76, 0x63, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nvc_proto_rawDescOnce sync.Once
	file_nvc_proto_rawDescData = file_nvc_proto_rawDesc
)

func file_nvc_proto_rawDescGZIP() []byte {
	file_nvc_proto_rawDescOnce.Do(func() {
		file_nvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_nvc_proto_rawDescData)
	})
	return file_nvc_proto_rawDescData
}

var file_nvc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nvc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nvc_proto_goTypes = []interface{}{
	(SQLtype)(0),         // 0: nvc.SQLtype
	(*Value)(nil),        // 1: nvc.Value
	(*SQLAction)(nil),    // 2: nvc.SQLAction
	(*ChromaResult)(nil), // 3: nvc.ChromaResult
}
var file_nvc_proto_depIdxs = []int32{
	0, // 0: nvc.SQLAction.type:type_name -> nvc.SQLtype
	1, // 1: nvc.SQLAction.values:type_name -> nvc.Value
	2, // 2: nvc.SQLToChromaListener.LogSQLAction:input_type -> nvc.SQLAction
	3, // 3: nvc.SQLToChromaListener.LogSQLAction:output_type -> nvc.ChromaResult
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nvc_proto_init() }
func file_nvc_proto_init() {
	if File_nvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_nvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLAction); i {
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
		file_nvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChromaResult); i {
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
	file_nvc_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ChromaResult_InsertedCount)(nil),
		(*ChromaResult_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nvc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nvc_proto_goTypes,
		DependencyIndexes: file_nvc_proto_depIdxs,
		EnumInfos:         file_nvc_proto_enumTypes,
		MessageInfos:      file_nvc_proto_msgTypes,
	}.Build()
	File_nvc_proto = out.File
	file_nvc_proto_rawDesc = nil
	file_nvc_proto_goTypes = nil
	file_nvc_proto_depIdxs = nil
}