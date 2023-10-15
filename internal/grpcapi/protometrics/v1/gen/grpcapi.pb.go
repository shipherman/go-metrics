// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.8
// source: internal/grpcapi/protometrics/grpcapi.proto

package protometrics

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Delta uint32 `protobuf:"varint,2,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_internal_grpcapi_protometrics_grpcapi_proto_rawDescGZIP(), []int{0}
}

func (x *Counter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Counter) GetDelta() uint32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

type Gauge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Gauge) Reset() {
	*x = Gauge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gauge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gauge) ProtoMessage() {}

func (x *Gauge) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gauge.ProtoReflect.Descriptor instead.
func (*Gauge) Descriptor() ([]byte, []int) {
	return file_internal_grpcapi_protometrics_grpcapi_proto_rawDescGZIP(), []int{1}
}

func (x *Gauge) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Gauge) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type GaugeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gauge *Gauge `protobuf:"bytes,1,opt,name=gauge,proto3" json:"gauge,omitempty"`
}

func (x *GaugeRequest) Reset() {
	*x = GaugeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaugeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaugeRequest) ProtoMessage() {}

func (x *GaugeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaugeRequest.ProtoReflect.Descriptor instead.
func (*GaugeRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpcapi_protometrics_grpcapi_proto_rawDescGZIP(), []int{2}
}

func (x *GaugeRequest) GetGauge() *Gauge {
	if x != nil {
		return x.Gauge
	}
	return nil
}

type GaugeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gauge *Gauge `protobuf:"bytes,1,opt,name=gauge,proto3" json:"gauge,omitempty"`
}

func (x *GaugeResponse) Reset() {
	*x = GaugeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaugeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaugeResponse) ProtoMessage() {}

func (x *GaugeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaugeResponse.ProtoReflect.Descriptor instead.
func (*GaugeResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpcapi_protometrics_grpcapi_proto_rawDescGZIP(), []int{3}
}

func (x *GaugeResponse) GetGauge() *Gauge {
	if x != nil {
		return x.Gauge
	}
	return nil
}

type CounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter *Counter `protobuf:"bytes,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *CounterRequest) Reset() {
	*x = CounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterRequest) ProtoMessage() {}

func (x *CounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterRequest.ProtoReflect.Descriptor instead.
func (*CounterRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpcapi_protometrics_grpcapi_proto_rawDescGZIP(), []int{4}
}

func (x *CounterRequest) GetCounter() *Counter {
	if x != nil {
		return x.Counter
	}
	return nil
}

type CounterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter *Counter `protobuf:"bytes,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *CounterResponse) Reset() {
	*x = CounterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterResponse) ProtoMessage() {}

func (x *CounterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterResponse.ProtoReflect.Descriptor instead.
func (*CounterResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpcapi_protometrics_grpcapi_proto_rawDescGZIP(), []int{5}
}

func (x *CounterResponse) GetCounter() *Counter {
	if x != nil {
		return x.Counter
	}
	return nil
}

var File_internal_grpcapi_protometrics_grpcapi_proto protoreflect.FileDescriptor

var file_internal_grpcapi_protometrics_grpcapi_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67,
	0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3c, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x22, 0x31,
	0x0a, 0x05, 0x47, 0x61, 0x75, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x41, 0x0a, 0x0c, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x05, 0x67,
	0x61, 0x75, 0x67, 0x65, 0x22, 0x42, 0x0a, 0x0d, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x61, 0x75, 0x67,
	0x65, 0x52, 0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x32,
	0xf6, 0x02, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x61, 0x75, 0x67, 0x65, 0x12, 0x22,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x61, 0x75, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x61,
	0x75, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x24, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_grpcapi_protometrics_grpcapi_proto_rawDescOnce sync.Once
	file_internal_grpcapi_protometrics_grpcapi_proto_rawDescData = file_internal_grpcapi_protometrics_grpcapi_proto_rawDesc
)

func file_internal_grpcapi_protometrics_grpcapi_proto_rawDescGZIP() []byte {
	file_internal_grpcapi_protometrics_grpcapi_proto_rawDescOnce.Do(func() {
		file_internal_grpcapi_protometrics_grpcapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpcapi_protometrics_grpcapi_proto_rawDescData)
	})
	return file_internal_grpcapi_protometrics_grpcapi_proto_rawDescData
}

var file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_grpcapi_protometrics_grpcapi_proto_goTypes = []interface{}{
	(*Counter)(nil),         // 0: grpcapi.protometrics.Counter
	(*Gauge)(nil),           // 1: grpcapi.protometrics.Gauge
	(*GaugeRequest)(nil),    // 2: grpcapi.protometrics.GaugeRequest
	(*GaugeResponse)(nil),   // 3: grpcapi.protometrics.GaugeResponse
	(*CounterRequest)(nil),  // 4: grpcapi.protometrics.CounterRequest
	(*CounterResponse)(nil), // 5: grpcapi.protometrics.CounterResponse
}
var file_internal_grpcapi_protometrics_grpcapi_proto_depIdxs = []int32{
	1, // 0: grpcapi.protometrics.GaugeRequest.gauge:type_name -> grpcapi.protometrics.Gauge
	1, // 1: grpcapi.protometrics.GaugeResponse.gauge:type_name -> grpcapi.protometrics.Gauge
	0, // 2: grpcapi.protometrics.CounterRequest.counter:type_name -> grpcapi.protometrics.Counter
	0, // 3: grpcapi.protometrics.CounterResponse.counter:type_name -> grpcapi.protometrics.Counter
	2, // 4: grpcapi.protometrics.MetricsService.GetGauge:input_type -> grpcapi.protometrics.GaugeRequest
	2, // 5: grpcapi.protometrics.MetricsService.UpdateGauge:input_type -> grpcapi.protometrics.GaugeRequest
	4, // 6: grpcapi.protometrics.MetricsService.GetCounter:input_type -> grpcapi.protometrics.CounterRequest
	4, // 7: grpcapi.protometrics.MetricsService.UpdateCounter:input_type -> grpcapi.protometrics.CounterRequest
	3, // 8: grpcapi.protometrics.MetricsService.GetGauge:output_type -> grpcapi.protometrics.GaugeResponse
	3, // 9: grpcapi.protometrics.MetricsService.UpdateGauge:output_type -> grpcapi.protometrics.GaugeResponse
	5, // 10: grpcapi.protometrics.MetricsService.GetCounter:output_type -> grpcapi.protometrics.CounterResponse
	5, // 11: grpcapi.protometrics.MetricsService.UpdateCounter:output_type -> grpcapi.protometrics.CounterResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_grpcapi_protometrics_grpcapi_proto_init() }
func file_internal_grpcapi_protometrics_grpcapi_proto_init() {
	if File_internal_grpcapi_protometrics_grpcapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
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
		file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gauge); i {
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
		file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaugeRequest); i {
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
		file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaugeResponse); i {
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
		file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterRequest); i {
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
		file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterResponse); i {
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
			RawDescriptor: file_internal_grpcapi_protometrics_grpcapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpcapi_protometrics_grpcapi_proto_goTypes,
		DependencyIndexes: file_internal_grpcapi_protometrics_grpcapi_proto_depIdxs,
		MessageInfos:      file_internal_grpcapi_protometrics_grpcapi_proto_msgTypes,
	}.Build()
	File_internal_grpcapi_protometrics_grpcapi_proto = out.File
	file_internal_grpcapi_protometrics_grpcapi_proto_rawDesc = nil
	file_internal_grpcapi_protometrics_grpcapi_proto_goTypes = nil
	file_internal_grpcapi_protometrics_grpcapi_proto_depIdxs = nil
}
