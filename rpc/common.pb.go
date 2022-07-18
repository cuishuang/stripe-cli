// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: common.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StripeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the object.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The Stripe API version used to render `data`. Note: This property is populated only for events
	// on or after October 31, 2014.
	ApiVersion string `protobuf:"bytes,2,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Object containing data associated with the event.
	Data *structpb.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// Information on the API request that instigated the event.
	Request *StripeEvent_Request `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	// Description of the event (e.g., invoice.created or charge.refunded).
	Type string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	// *CONNECT ONLY* The connected account that originated the event.
	Account string `protobuf:"bytes,6,opt,name=account,proto3" json:"account,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `protobuf:"varint,7,opt,name=created,proto3" json:"created,omitempty"`
	// Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode bool `protobuf:"varint,8,opt,name=livemode,proto3" json:"livemode,omitempty"`
	// Number of webhooks that have yet to be successfully delivered (i.e., to return a 20x response)
	// to the URLs you’ve specified.
	PendingWebhooks int64 `protobuf:"varint,9,opt,name=pending_webhooks,json=pendingWebhooks,proto3" json:"pending_webhooks,omitempty"`
}

func (x *StripeEvent) Reset() {
	*x = StripeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StripeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StripeEvent) ProtoMessage() {}

func (x *StripeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StripeEvent.ProtoReflect.Descriptor instead.
func (*StripeEvent) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *StripeEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StripeEvent) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *StripeEvent) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StripeEvent) GetRequest() *StripeEvent_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *StripeEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StripeEvent) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *StripeEvent) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *StripeEvent) GetLivemode() bool {
	if x != nil {
		return x.Livemode
	}
	return false
}

func (x *StripeEvent) GetPendingWebhooks() int64 {
	if x != nil {
		return x.PendingWebhooks
	}
	return 0
}

type StripeEvent_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the API request that caused the event. If null, the event was automatic (e.g., Stripe’s
	// automatic subscription handling). Request logs are available in the dashboard, but currently
	// not in the API.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The idempotency key transmitted during the request, if any. Note: This property is populated
	// only for events on or after May 23, 2017.
	IdempotencyKey string `protobuf:"bytes,2,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
}

func (x *StripeEvent_Request) Reset() {
	*x = StripeEvent_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StripeEvent_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StripeEvent_Request) ProtoMessage() {}

func (x *StripeEvent_Request) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StripeEvent_Request.ProtoReflect.Descriptor instead.
func (*StripeEvent_Request) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StripeEvent_Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StripeEvent_Request) GetIdempotencyKey() string {
	if x != nil {
		return x.IdempotencyKey
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x72, 0x70, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf2, 0x02, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x32, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x69, 0x76, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6c,
	0x69, 0x76, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x73, 0x1a, 0x42, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65,
	0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_proto_goTypes = []interface{}{
	(*StripeEvent)(nil),         // 0: rpc.StripeEvent
	(*StripeEvent_Request)(nil), // 1: rpc.StripeEvent.Request
	(*structpb.Struct)(nil),     // 2: google.protobuf.Struct
}
var file_common_proto_depIdxs = []int32{
	2, // 0: rpc.StripeEvent.data:type_name -> google.protobuf.Struct
	1, // 1: rpc.StripeEvent.request:type_name -> rpc.StripeEvent.Request
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StripeEvent); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StripeEvent_Request); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
