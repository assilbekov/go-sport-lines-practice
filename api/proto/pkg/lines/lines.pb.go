// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: api/proto/lines.proto

package lines

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

type SportLinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sports   []string `protobuf:"bytes,1,rep,name=sports,proto3" json:"sports,omitempty"`
	Interval int32    `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"` // interval in seconds
}

func (x *SportLinesRequest) Reset() {
	*x = SportLinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_lines_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportLinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportLinesRequest) ProtoMessage() {}

func (x *SportLinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_lines_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportLinesRequest.ProtoReflect.Descriptor instead.
func (*SportLinesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_lines_proto_rawDescGZIP(), []int{0}
}

func (x *SportLinesRequest) GetSports() []string {
	if x != nil {
		return x.Sports
	}
	return nil
}

func (x *SportLinesRequest) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type SportLinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines map[string]float64 `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *SportLinesResponse) Reset() {
	*x = SportLinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_lines_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportLinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportLinesResponse) ProtoMessage() {}

func (x *SportLinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_lines_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportLinesResponse.ProtoReflect.Descriptor instead.
func (*SportLinesResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_lines_proto_rawDescGZIP(), []int{1}
}

func (x *SportLinesResponse) GetLines() map[string]float64 {
	if x != nil {
		return x.Lines
	}
	return nil
}

var File_api_proto_lines_proto protoreflect.FileDescriptor

var file_api_proto_lines_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x47,
	0x0a, 0x11, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x4c, 0x69,
	0x6e, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x62, 0x0a, 0x0c, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x4f, 0x6e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x18, 0x2e,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_lines_proto_rawDescOnce sync.Once
	file_api_proto_lines_proto_rawDescData = file_api_proto_lines_proto_rawDesc
)

func file_api_proto_lines_proto_rawDescGZIP() []byte {
	file_api_proto_lines_proto_rawDescOnce.Do(func() {
		file_api_proto_lines_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_lines_proto_rawDescData)
	})
	return file_api_proto_lines_proto_rawDescData
}

var file_api_proto_lines_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_lines_proto_goTypes = []interface{}{
	(*SportLinesRequest)(nil),  // 0: lines.SportLinesRequest
	(*SportLinesResponse)(nil), // 1: lines.SportLinesResponse
	nil,                        // 2: lines.SportLinesResponse.LinesEntry
}
var file_api_proto_lines_proto_depIdxs = []int32{
	2, // 0: lines.SportLinesResponse.lines:type_name -> lines.SportLinesResponse.LinesEntry
	0, // 1: lines.LinesService.SubscribeOnSportLines:input_type -> lines.SportLinesRequest
	1, // 2: lines.LinesService.SubscribeOnSportLines:output_type -> lines.SportLinesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_lines_proto_init() }
func file_api_proto_lines_proto_init() {
	if File_api_proto_lines_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_lines_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportLinesRequest); i {
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
		file_api_proto_lines_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportLinesResponse); i {
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
			RawDescriptor: file_api_proto_lines_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_lines_proto_goTypes,
		DependencyIndexes: file_api_proto_lines_proto_depIdxs,
		MessageInfos:      file_api_proto_lines_proto_msgTypes,
	}.Build()
	File_api_proto_lines_proto = out.File
	file_api_proto_lines_proto_rawDesc = nil
	file_api_proto_lines_proto_goTypes = nil
	file_api_proto_lines_proto_depIdxs = nil
}
