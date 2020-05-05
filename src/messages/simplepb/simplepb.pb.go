// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: messages/simplepb/simplepb.proto

package simplepb

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

type WeekDay int32

const (
	WeekDay_UNKNOWN   WeekDay = 0
	WeekDay_MONDAY    WeekDay = 1
	WeekDay_TUESDAY   WeekDay = 2
	WeekDay_WEDNESDAY WeekDay = 3
	WeekDay_THURSDAY  WeekDay = 4
	WeekDay_FRIDAY    WeekDay = 5
	WeekDay_SATURDAY  WeekDay = 6
	WeekDay_SUNDAY    WeekDay = 7
)

// Enum value maps for WeekDay.
var (
	WeekDay_name = map[int32]string{
		0: "UNKNOWN",
		1: "MONDAY",
		2: "TUESDAY",
		3: "WEDNESDAY",
		4: "THURSDAY",
		5: "FRIDAY",
		6: "SATURDAY",
		7: "SUNDAY",
	}
	WeekDay_value = map[string]int32{
		"UNKNOWN":   0,
		"MONDAY":    1,
		"TUESDAY":   2,
		"WEDNESDAY": 3,
		"THURSDAY":  4,
		"FRIDAY":    5,
		"SATURDAY":  6,
		"SUNDAY":    7,
	}
)

func (x WeekDay) Enum() *WeekDay {
	p := new(WeekDay)
	*p = x
	return p
}

func (x WeekDay) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WeekDay) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_simplepb_simplepb_proto_enumTypes[0].Descriptor()
}

func (WeekDay) Type() protoreflect.EnumType {
	return &file_messages_simplepb_simplepb_proto_enumTypes[0]
}

func (x WeekDay) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WeekDay.Descriptor instead.
func (WeekDay) EnumDescriptor() ([]byte, []int) {
	return file_messages_simplepb_simplepb_proto_rawDescGZIP(), []int{0}
}

type SimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsSimple     bool    `protobuf:"varint,2,opt,name=is_simple,json=isSimple,proto3" json:"is_simple,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SampleList   []int32 `protobuf:"varint,4,rep,packed,name=sample_list,json=sampleList,proto3" json:"sample_list,omitempty"`
	DayOfTheWeek WeekDay `protobuf:"varint,5,opt,name=day_of_the_week,json=dayOfTheWeek,proto3,enum=simple.WeekDay" json:"day_of_the_week,omitempty"`
}

func (x *SimpleMessage) Reset() {
	*x = SimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_simplepb_simplepb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessage) ProtoMessage() {}

func (x *SimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_simplepb_simplepb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessage.ProtoReflect.Descriptor instead.
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return file_messages_simplepb_simplepb_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SimpleMessage) GetIsSimple() bool {
	if x != nil {
		return x.IsSimple
	}
	return false
}

func (x *SimpleMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimpleMessage) GetSampleList() []int32 {
	if x != nil {
		return x.SampleList
	}
	return nil
}

func (x *SimpleMessage) GetDayOfTheWeek() WeekDay {
	if x != nil {
		return x.DayOfTheWeek
	}
	return WeekDay_UNKNOWN
}

var File_messages_simplepb_simplepb_proto protoreflect.FileDescriptor

var file_messages_simplepb_simplepb_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x0f, 0x64, 0x61, 0x79, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x68, 0x65, 0x5f, 0x77, 0x65, 0x65,
	0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x52, 0x0c, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x54,
	0x68, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x2a, 0x72, 0x0a, 0x07, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61,
	0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x55,
	0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x45, 0x44, 0x4e, 0x45,
	0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x48, 0x55, 0x52, 0x53, 0x44,
	0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x49, 0x44, 0x41, 0x59, 0x10, 0x05,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59, 0x10, 0x06, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x07, 0x42, 0x13, 0x5a, 0x11, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_simplepb_simplepb_proto_rawDescOnce sync.Once
	file_messages_simplepb_simplepb_proto_rawDescData = file_messages_simplepb_simplepb_proto_rawDesc
)

func file_messages_simplepb_simplepb_proto_rawDescGZIP() []byte {
	file_messages_simplepb_simplepb_proto_rawDescOnce.Do(func() {
		file_messages_simplepb_simplepb_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_simplepb_simplepb_proto_rawDescData)
	})
	return file_messages_simplepb_simplepb_proto_rawDescData
}

var file_messages_simplepb_simplepb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_simplepb_simplepb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messages_simplepb_simplepb_proto_goTypes = []interface{}{
	(WeekDay)(0),          // 0: simple.WeekDay
	(*SimpleMessage)(nil), // 1: simple.SimpleMessage
}
var file_messages_simplepb_simplepb_proto_depIdxs = []int32{
	0, // 0: simple.SimpleMessage.day_of_the_week:type_name -> simple.WeekDay
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_simplepb_simplepb_proto_init() }
func file_messages_simplepb_simplepb_proto_init() {
	if File_messages_simplepb_simplepb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_simplepb_simplepb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessage); i {
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
			RawDescriptor: file_messages_simplepb_simplepb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_simplepb_simplepb_proto_goTypes,
		DependencyIndexes: file_messages_simplepb_simplepb_proto_depIdxs,
		EnumInfos:         file_messages_simplepb_simplepb_proto_enumTypes,
		MessageInfos:      file_messages_simplepb_simplepb_proto_msgTypes,
	}.Build()
	File_messages_simplepb_simplepb_proto = out.File
	file_messages_simplepb_simplepb_proto_rawDesc = nil
	file_messages_simplepb_simplepb_proto_goTypes = nil
	file_messages_simplepb_simplepb_proto_depIdxs = nil
}