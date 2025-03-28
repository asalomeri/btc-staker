// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/transaction.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TrackedTransaction struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// index of tracked transaction in database, first tracked transaction has index 1
	TrackedTransactionIdx uint64 `protobuf:"varint,1,opt,name=tracked_transaction_idx,json=trackedTransactionIdx,proto3" json:"tracked_transaction_idx,omitempty"`
	StakingTransaction    []byte `protobuf:"bytes,2,opt,name=staking_transaction,json=stakingTransaction,proto3" json:"staking_transaction,omitempty"`
	StakerAddress         string `protobuf:"bytes,3,opt,name=staker_address,json=stakerAddress,proto3" json:"staker_address,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *TrackedTransaction) Reset() {
	*x = TrackedTransaction{}
	mi := &file_proto_transaction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackedTransaction) ProtoMessage() {}

func (x *TrackedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackedTransaction.ProtoReflect.Descriptor instead.
func (*TrackedTransaction) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TrackedTransaction) GetTrackedTransactionIdx() uint64 {
	if x != nil {
		return x.TrackedTransactionIdx
	}
	return 0
}

func (x *TrackedTransaction) GetStakingTransaction() []byte {
	if x != nil {
		return x.StakingTransaction
	}
	return nil
}

func (x *TrackedTransaction) GetStakerAddress() string {
	if x != nil {
		return x.StakerAddress
	}
	return ""
}

var File_proto_transaction_proto protoreflect.FileDescriptor

var file_proto_transaction_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa4, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x78, 0x12,
	0x2f, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x62, 0x79, 0x6c, 0x6f, 0x6e, 0x6c, 0x61, 0x62,
	0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x62, 0x74, 0x63, 0x2d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_transaction_proto_rawDescOnce sync.Once
	file_proto_transaction_proto_rawDescData []byte
)

func file_proto_transaction_proto_rawDescGZIP() []byte {
	file_proto_transaction_proto_rawDescOnce.Do(func() {
		file_proto_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_transaction_proto_rawDesc), len(file_proto_transaction_proto_rawDesc)))
	})
	return file_proto_transaction_proto_rawDescData
}

var file_proto_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_transaction_proto_goTypes = []any{
	(*TrackedTransaction)(nil), // 0: proto.TrackedTransaction
}
var file_proto_transaction_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_transaction_proto_init() }
func file_proto_transaction_proto_init() {
	if File_proto_transaction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_transaction_proto_rawDesc), len(file_proto_transaction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_transaction_proto_goTypes,
		DependencyIndexes: file_proto_transaction_proto_depIdxs,
		MessageInfos:      file_proto_transaction_proto_msgTypes,
	}.Build()
	File_proto_transaction_proto = out.File
	file_proto_transaction_proto_goTypes = nil
	file_proto_transaction_proto_depIdxs = nil
}
