// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/staff/v1/staff.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_STAFFNO_MISSING ErrorReason = 0
	ErrorReason_GETSTAFF_FAILED ErrorReason = 1
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "STAFFNO_MISSING",
		1: "GETSTAFF_FAILED",
	}
	ErrorReason_value = map[string]int32{
		"STAFFNO_MISSING": 0,
		"GETSTAFF_FAILED": 1,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_staff_v1_staff_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_staff_v1_staff_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_staff_v1_staff_proto_rawDescGZIP(), []int{0}
}

type GetStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staffno string `protobuf:"bytes,1,opt,name=staffno,proto3" json:"staffno,omitempty"`
}

func (x *GetStaffRequest) Reset() {
	*x = GetStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_v1_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaffRequest) ProtoMessage() {}

func (x *GetStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_v1_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaffRequest.ProtoReflect.Descriptor instead.
func (*GetStaffRequest) Descriptor() ([]byte, []int) {
	return file_api_staff_v1_staff_proto_rawDescGZIP(), []int{0}
}

func (x *GetStaffRequest) GetStaffno() string {
	if x != nil {
		return x.Staffno
	}
	return ""
}

type GetStaffReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	No     string `protobuf:"bytes,2,opt,name=no,proto3" json:"no,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Gender string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Deptid int64  `protobuf:"varint,5,opt,name=deptid,proto3" json:"deptid,omitempty"`
	Dutyid int64  `protobuf:"varint,6,opt,name=dutyid,proto3" json:"dutyid,omitempty"`
}

func (x *GetStaffReply) Reset() {
	*x = GetStaffReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_v1_staff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaffReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaffReply) ProtoMessage() {}

func (x *GetStaffReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_v1_staff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaffReply.ProtoReflect.Descriptor instead.
func (*GetStaffReply) Descriptor() ([]byte, []int) {
	return file_api_staff_v1_staff_proto_rawDescGZIP(), []int{1}
}

func (x *GetStaffReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetStaffReply) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

func (x *GetStaffReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetStaffReply) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *GetStaffReply) GetDeptid() int64 {
	if x != nil {
		return x.Deptid
	}
	return 0
}

func (x *GetStaffReply) GetDutyid() int64 {
	if x != nil {
		return x.Dutyid
	}
	return 0
}

var File_api_staff_v1_staff_proto protoreflect.FileDescriptor

var file_api_staff_v1_staff_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x6e, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x70, 0x74,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x65, 0x70, 0x74, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x75, 0x74, 0x79, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x64, 0x75, 0x74, 0x79, 0x69, 0x64, 0x2a, 0x49, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x46, 0x46,
	0x4e, 0x4f, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45,
	0x90, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x47, 0x45, 0x54, 0x53, 0x54, 0x41, 0x46, 0x46, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a, 0x04, 0xa0,
	0x45, 0xf4, 0x03, 0x32, 0x4f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x46, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x28, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x16, 0x67, 0x62, 0x6d, 0x65, 0x72, 0x70, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_staff_v1_staff_proto_rawDescOnce sync.Once
	file_api_staff_v1_staff_proto_rawDescData = file_api_staff_v1_staff_proto_rawDesc
)

func file_api_staff_v1_staff_proto_rawDescGZIP() []byte {
	file_api_staff_v1_staff_proto_rawDescOnce.Do(func() {
		file_api_staff_v1_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_staff_v1_staff_proto_rawDescData)
	})
	return file_api_staff_v1_staff_proto_rawDescData
}

var file_api_staff_v1_staff_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_staff_v1_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_staff_v1_staff_proto_goTypes = []interface{}{
	(ErrorReason)(0),        // 0: api.staff.v1.ErrorReason
	(*GetStaffRequest)(nil), // 1: api.staff.v1.GetStaffRequest
	(*GetStaffReply)(nil),   // 2: api.staff.v1.GetStaffReply
}
var file_api_staff_v1_staff_proto_depIdxs = []int32{
	1, // 0: api.staff.v1.Staff.GetStaff:input_type -> api.staff.v1.GetStaffRequest
	2, // 1: api.staff.v1.Staff.GetStaff:output_type -> api.staff.v1.GetStaffReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_staff_v1_staff_proto_init() }
func file_api_staff_v1_staff_proto_init() {
	if File_api_staff_v1_staff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_staff_v1_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaffRequest); i {
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
		file_api_staff_v1_staff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaffReply); i {
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
			RawDescriptor: file_api_staff_v1_staff_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_staff_v1_staff_proto_goTypes,
		DependencyIndexes: file_api_staff_v1_staff_proto_depIdxs,
		EnumInfos:         file_api_staff_v1_staff_proto_enumTypes,
		MessageInfos:      file_api_staff_v1_staff_proto_msgTypes,
	}.Build()
	File_api_staff_v1_staff_proto = out.File
	file_api_staff_v1_staff_proto_rawDesc = nil
	file_api_staff_v1_staff_proto_goTypes = nil
	file_api_staff_v1_staff_proto_depIdxs = nil
}
