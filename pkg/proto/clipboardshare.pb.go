// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/clipboardshare.proto

package proto

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

type UpdateClipboardMsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *UpdateClipboardMsgRequest) Reset() {
	*x = UpdateClipboardMsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clipboardshare_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClipboardMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClipboardMsgRequest) ProtoMessage() {}

func (x *UpdateClipboardMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clipboardshare_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClipboardMsgRequest.ProtoReflect.Descriptor instead.
func (*UpdateClipboardMsgRequest) Descriptor() ([]byte, []int) {
	return file_proto_clipboardshare_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateClipboardMsgRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UpdateClipboardMsgRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type UpdateClipboardMsgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *UpdateClipboardMsgResponse) Reset() {
	*x = UpdateClipboardMsgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clipboardshare_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClipboardMsgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClipboardMsgResponse) ProtoMessage() {}

func (x *UpdateClipboardMsgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clipboardshare_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClipboardMsgResponse.ProtoReflect.Descriptor instead.
func (*UpdateClipboardMsgResponse) Descriptor() ([]byte, []int) {
	return file_proto_clipboardshare_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateClipboardMsgResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetClipboardMsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClipboardMsgRequest) Reset() {
	*x = GetClipboardMsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clipboardshare_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClipboardMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClipboardMsgRequest) ProtoMessage() {}

func (x *GetClipboardMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clipboardshare_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClipboardMsgRequest.ProtoReflect.Descriptor instead.
func (*GetClipboardMsgRequest) Descriptor() ([]byte, []int) {
	return file_proto_clipboardshare_proto_rawDescGZIP(), []int{2}
}

type GetClipboardMsgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data           []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	LastUpdateTime int64  `protobuf:"varint,2,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime,omitempty"`
}

func (x *GetClipboardMsgResponse) Reset() {
	*x = GetClipboardMsgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clipboardshare_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClipboardMsgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClipboardMsgResponse) ProtoMessage() {}

func (x *GetClipboardMsgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clipboardshare_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClipboardMsgResponse.ProtoReflect.Descriptor instead.
func (*GetClipboardMsgResponse) Descriptor() ([]byte, []int) {
	return file_proto_clipboardshare_proto_rawDescGZIP(), []int{3}
}

func (x *GetClipboardMsgResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetClipboardMsgResponse) GetLastUpdateTime() int64 {
	if x != nil {
		return x.LastUpdateTime
	}
	return 0
}

var File_proto_clipboardshare_proto protoreflect.FileDescriptor

var file_proto_clipboardshare_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x30, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x55,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xbc, 0x01, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x5b, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x4d, 0x73, 0x67, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x79, 0x64, 0x6f, 0x76, 0x65, 0x65, 0x2f, 0x63, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_clipboardshare_proto_rawDescOnce sync.Once
	file_proto_clipboardshare_proto_rawDescData = file_proto_clipboardshare_proto_rawDesc
)

func file_proto_clipboardshare_proto_rawDescGZIP() []byte {
	file_proto_clipboardshare_proto_rawDescOnce.Do(func() {
		file_proto_clipboardshare_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_clipboardshare_proto_rawDescData)
	})
	return file_proto_clipboardshare_proto_rawDescData
}

var file_proto_clipboardshare_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_clipboardshare_proto_goTypes = []interface{}{
	(*UpdateClipboardMsgRequest)(nil),  // 0: proto.UpdateClipboardMsgRequest
	(*UpdateClipboardMsgResponse)(nil), // 1: proto.UpdateClipboardMsgResponse
	(*GetClipboardMsgRequest)(nil),     // 2: proto.GetClipboardMsgRequest
	(*GetClipboardMsgResponse)(nil),    // 3: proto.GetClipboardMsgResponse
}
var file_proto_clipboardshare_proto_depIdxs = []int32{
	0, // 0: proto.Clipboard.UpdateClipboardMsg:input_type -> proto.UpdateClipboardMsgRequest
	2, // 1: proto.Clipboard.GetClipboardMsg:input_type -> proto.GetClipboardMsgRequest
	1, // 2: proto.Clipboard.UpdateClipboardMsg:output_type -> proto.UpdateClipboardMsgResponse
	3, // 3: proto.Clipboard.GetClipboardMsg:output_type -> proto.GetClipboardMsgResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_clipboardshare_proto_init() }
func file_proto_clipboardshare_proto_init() {
	if File_proto_clipboardshare_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_clipboardshare_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClipboardMsgRequest); i {
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
		file_proto_clipboardshare_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClipboardMsgResponse); i {
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
		file_proto_clipboardshare_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClipboardMsgRequest); i {
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
		file_proto_clipboardshare_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClipboardMsgResponse); i {
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
			RawDescriptor: file_proto_clipboardshare_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_clipboardshare_proto_goTypes,
		DependencyIndexes: file_proto_clipboardshare_proto_depIdxs,
		MessageInfos:      file_proto_clipboardshare_proto_msgTypes,
	}.Build()
	File_proto_clipboardshare_proto = out.File
	file_proto_clipboardshare_proto_rawDesc = nil
	file_proto_clipboardshare_proto_goTypes = nil
	file_proto_clipboardshare_proto_depIdxs = nil
}
