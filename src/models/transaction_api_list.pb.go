// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: transaction_api_list.proto

package models

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

type TransactionAPIList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address"`
	ToAddress   string `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address"`
	Value       string `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
	Timestamp   string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp"`
	Hash        string `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash"`
	BlockNumber uint64 `protobuf:"varint,6,opt,name=block_number,json=blockNumber,proto3" json:"block_number"`
	Fee         uint64 `protobuf:"varint,7,opt,name=fee,proto3" json:"fee"`
}

func (x *TransactionAPIList) Reset() {
	*x = TransactionAPIList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_api_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionAPIList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionAPIList) ProtoMessage() {}

func (x *TransactionAPIList) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_api_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionAPIList.ProtoReflect.Descriptor instead.
func (*TransactionAPIList) Descriptor() ([]byte, []int) {
	return file_transaction_api_list_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionAPIList) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *TransactionAPIList) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *TransactionAPIList) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TransactionAPIList) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *TransactionAPIList) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *TransactionAPIList) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *TransactionAPIList) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

var File_transaction_api_list_proto protoreflect.FileDescriptor

var file_transaction_api_list_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x22, 0xd3, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x50, 0x49, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_api_list_proto_rawDescOnce sync.Once
	file_transaction_api_list_proto_rawDescData = file_transaction_api_list_proto_rawDesc
)

func file_transaction_api_list_proto_rawDescGZIP() []byte {
	file_transaction_api_list_proto_rawDescOnce.Do(func() {
		file_transaction_api_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_api_list_proto_rawDescData)
	})
	return file_transaction_api_list_proto_rawDescData
}

var file_transaction_api_list_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transaction_api_list_proto_goTypes = []interface{}{
	(*TransactionAPIList)(nil), // 0: models.TransactionAPIList
}
var file_transaction_api_list_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transaction_api_list_proto_init() }
func file_transaction_api_list_proto_init() {
	if File_transaction_api_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_api_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionAPIList); i {
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
			RawDescriptor: file_transaction_api_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_api_list_proto_goTypes,
		DependencyIndexes: file_transaction_api_list_proto_depIdxs,
		MessageInfos:      file_transaction_api_list_proto_msgTypes,
	}.Build()
	File_transaction_api_list_proto = out.File
	file_transaction_api_list_proto_rawDesc = nil
	file_transaction_api_list_proto_goTypes = nil
	file_transaction_api_list_proto_depIdxs = nil
}
