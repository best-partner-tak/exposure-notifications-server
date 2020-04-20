// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: pkg/pb/federation.proto

package pb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type DiagnosisStatus int32

const (
	DiagnosisStatus_unknown           DiagnosisStatus = 0
	DiagnosisStatus_positive_verified DiagnosisStatus = 1
	DiagnosisStatus_self_reported     DiagnosisStatus = 2
)

// Enum value maps for DiagnosisStatus.
var (
	DiagnosisStatus_name = map[int32]string{
		0: "unknown",
		1: "positive_verified",
		2: "self_reported",
	}
	DiagnosisStatus_value = map[string]int32{
		"unknown":           0,
		"positive_verified": 1,
		"self_reported":     2,
	}
)

func (x DiagnosisStatus) Enum() *DiagnosisStatus {
	p := new(DiagnosisStatus)
	*p = x
	return p
}

func (x DiagnosisStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiagnosisStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_federation_proto_enumTypes[0].Descriptor()
}

func (DiagnosisStatus) Type() protoreflect.EnumType {
	return &file_pkg_pb_federation_proto_enumTypes[0]
}

func (x DiagnosisStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiagnosisStatus.Descriptor instead.
func (DiagnosisStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_federation_proto_rawDescGZIP(), []int{0}
}

type FederationFetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionIdentifiers             []string `protobuf:"bytes,1,rep,name=regionIdentifiers,proto3" json:"regionIdentifiers,omitempty"`
	ExcludeRegionIdentifiers      []string `protobuf:"bytes,2,rep,name=excludeRegionIdentifiers,proto3" json:"excludeRegionIdentifiers,omitempty"`
	LastFetchResponseKeyTimestamp int64    `protobuf:"varint,3,opt,name=lastFetchResponseKeyTimestamp,proto3" json:"lastFetchResponseKeyTimestamp,omitempty"`
	// regionIdentifiers, excludeRegionIdentifiers, lastFetchResponseKeyTimestamp must be stable to send a fetchToken.
	FetchToken string `protobuf:"bytes,4,opt,name=fetchToken,proto3" json:"fetchToken,omitempty"`
}

func (x *FederationFetchRequest) Reset() {
	*x = FederationFetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_federation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationFetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationFetchRequest) ProtoMessage() {}

func (x *FederationFetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_federation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationFetchRequest.ProtoReflect.Descriptor instead.
func (*FederationFetchRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_federation_proto_rawDescGZIP(), []int{0}
}

func (x *FederationFetchRequest) GetRegionIdentifiers() []string {
	if x != nil {
		return x.RegionIdentifiers
	}
	return nil
}

func (x *FederationFetchRequest) GetExcludeRegionIdentifiers() []string {
	if x != nil {
		return x.ExcludeRegionIdentifiers
	}
	return nil
}

func (x *FederationFetchRequest) GetLastFetchResponseKeyTimestamp() int64 {
	if x != nil {
		return x.LastFetchResponseKeyTimestamp
	}
	return 0
}

func (x *FederationFetchRequest) GetFetchToken() string {
	if x != nil {
		return x.FetchToken
	}
	return ""
}

type FederationFetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response                  []*ContactTracingResponse `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
	PartialResponse           bool                      `protobuf:"varint,2,opt,name=partialResponse,proto3" json:"partialResponse,omitempty"`                     // required
	NextFetchToken            string                    `protobuf:"bytes,3,opt,name=nextFetchToken,proto3" json:"nextFetchToken,omitempty"`                        // nextFetchToken will be present if partialResponse==true
	FetchResponseKeyTimestamp int64                     `protobuf:"varint,4,opt,name=fetchResponseKeyTimestamp,proto3" json:"fetchResponseKeyTimestamp,omitempty"` // required
}

func (x *FederationFetchResponse) Reset() {
	*x = FederationFetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_federation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationFetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationFetchResponse) ProtoMessage() {}

func (x *FederationFetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_federation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationFetchResponse.ProtoReflect.Descriptor instead.
func (*FederationFetchResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_federation_proto_rawDescGZIP(), []int{1}
}

func (x *FederationFetchResponse) GetResponse() []*ContactTracingResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *FederationFetchResponse) GetPartialResponse() bool {
	if x != nil {
		return x.PartialResponse
	}
	return false
}

func (x *FederationFetchResponse) GetNextFetchToken() string {
	if x != nil {
		return x.NextFetchToken
	}
	return ""
}

func (x *FederationFetchResponse) GetFetchResponseKeyTimestamp() int64 {
	if x != nil {
		return x.FetchResponseKeyTimestamp
	}
	return 0
}

type ContactTracingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactTracingInfo []*ContactTracingInfo `protobuf:"bytes,1,rep,name=contactTracingInfo,proto3" json:"contactTracingInfo,omitempty"`
	RegionIdentifiers  []string              `protobuf:"bytes,2,rep,name=regionIdentifiers,proto3" json:"regionIdentifiers,omitempty"`
}

func (x *ContactTracingResponse) Reset() {
	*x = ContactTracingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_federation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactTracingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactTracingResponse) ProtoMessage() {}

func (x *ContactTracingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_federation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactTracingResponse.ProtoReflect.Descriptor instead.
func (*ContactTracingResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_federation_proto_rawDescGZIP(), []int{2}
}

func (x *ContactTracingResponse) GetContactTracingInfo() []*ContactTracingInfo {
	if x != nil {
		return x.ContactTracingInfo
	}
	return nil
}

func (x *ContactTracingResponse) GetRegionIdentifiers() []string {
	if x != nil {
		return x.RegionIdentifiers
	}
	return nil
}

type ContactTracingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiagnosisStatus DiagnosisStatus `protobuf:"varint,1,opt,name=diagnosisStatus,proto3,enum=DiagnosisStatus" json:"diagnosisStatus,omitempty"` // required
	DiagnosisKeys   []*DiagnosisKey `protobuf:"bytes,2,rep,name=diagnosisKeys,proto3" json:"diagnosisKeys,omitempty"`
}

func (x *ContactTracingInfo) Reset() {
	*x = ContactTracingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_federation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactTracingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactTracingInfo) ProtoMessage() {}

func (x *ContactTracingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_federation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactTracingInfo.ProtoReflect.Descriptor instead.
func (*ContactTracingInfo) Descriptor() ([]byte, []int) {
	return file_pkg_pb_federation_proto_rawDescGZIP(), []int{3}
}

func (x *ContactTracingInfo) GetDiagnosisStatus() DiagnosisStatus {
	if x != nil {
		return x.DiagnosisStatus
	}
	return DiagnosisStatus_unknown
}

func (x *ContactTracingInfo) GetDiagnosisKeys() []*DiagnosisKey {
	if x != nil {
		return x.DiagnosisKeys
	}
	return nil
}

type DiagnosisKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiagnosisKey   []byte `protobuf:"bytes,1,opt,name=diagnosisKey,proto3" json:"diagnosisKey,omitempty"`      // required
	IntervalNumber int64  `protobuf:"varint,2,opt,name=intervalNumber,proto3" json:"intervalNumber,omitempty"` // required
	IntervalCount  int64  `protobuf:"varint,3,opt,name=intervalCount,proto3" json:"intervalCount,omitempty"`   // required
}

func (x *DiagnosisKey) Reset() {
	*x = DiagnosisKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_federation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiagnosisKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiagnosisKey) ProtoMessage() {}

func (x *DiagnosisKey) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_federation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiagnosisKey.ProtoReflect.Descriptor instead.
func (*DiagnosisKey) Descriptor() ([]byte, []int) {
	return file_pkg_pb_federation_proto_rawDescGZIP(), []int{4}
}

func (x *DiagnosisKey) GetDiagnosisKey() []byte {
	if x != nil {
		return x.DiagnosisKey
	}
	return nil
}

func (x *DiagnosisKey) GetIntervalNumber() int64 {
	if x != nil {
		return x.IntervalNumber
	}
	return 0
}

func (x *DiagnosisKey) GetIntervalCount() int64 {
	if x != nil {
		return x.IntervalCount
	}
	return 0
}

var File_pkg_pb_federation_proto protoreflect.FileDescriptor

var file_pkg_pb_federation_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x16, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x44,
	0x0a, 0x1d, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xde, 0x01, 0x0a, 0x17, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3c, 0x0a, 0x19, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x19, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x8b, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x0f, 0x64, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0d, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x64, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0c,
	0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c,
	0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x4b, 0x65, 0x79,
	0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x48,
	0x0a, 0x0f, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x10, 0x02, 0x32, 0x4a, 0x0a, 0x0a, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12,
	0x17, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_federation_proto_rawDescOnce sync.Once
	file_pkg_pb_federation_proto_rawDescData = file_pkg_pb_federation_proto_rawDesc
)

func file_pkg_pb_federation_proto_rawDescGZIP() []byte {
	file_pkg_pb_federation_proto_rawDescOnce.Do(func() {
		file_pkg_pb_federation_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_federation_proto_rawDescData)
	})
	return file_pkg_pb_federation_proto_rawDescData
}

var file_pkg_pb_federation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_pb_federation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_pb_federation_proto_goTypes = []interface{}{
	(DiagnosisStatus)(0),            // 0: DiagnosisStatus
	(*FederationFetchRequest)(nil),  // 1: FederationFetchRequest
	(*FederationFetchResponse)(nil), // 2: FederationFetchResponse
	(*ContactTracingResponse)(nil),  // 3: ContactTracingResponse
	(*ContactTracingInfo)(nil),      // 4: ContactTracingInfo
	(*DiagnosisKey)(nil),            // 5: DiagnosisKey
}
var file_pkg_pb_federation_proto_depIdxs = []int32{
	3, // 0: FederationFetchResponse.response:type_name -> ContactTracingResponse
	4, // 1: ContactTracingResponse.contactTracingInfo:type_name -> ContactTracingInfo
	0, // 2: ContactTracingInfo.diagnosisStatus:type_name -> DiagnosisStatus
	5, // 3: ContactTracingInfo.diagnosisKeys:type_name -> DiagnosisKey
	1, // 4: Federation.Fetch:input_type -> FederationFetchRequest
	2, // 5: Federation.Fetch:output_type -> FederationFetchResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_federation_proto_init() }
func file_pkg_pb_federation_proto_init() {
	if File_pkg_pb_federation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_federation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederationFetchRequest); i {
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
		file_pkg_pb_federation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederationFetchResponse); i {
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
		file_pkg_pb_federation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactTracingResponse); i {
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
		file_pkg_pb_federation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactTracingInfo); i {
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
		file_pkg_pb_federation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiagnosisKey); i {
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
			RawDescriptor: file_pkg_pb_federation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_federation_proto_goTypes,
		DependencyIndexes: file_pkg_pb_federation_proto_depIdxs,
		EnumInfos:         file_pkg_pb_federation_proto_enumTypes,
		MessageInfos:      file_pkg_pb_federation_proto_msgTypes,
	}.Build()
	File_pkg_pb_federation_proto = out.File
	file_pkg_pb_federation_proto_rawDesc = nil
	file_pkg_pb_federation_proto_goTypes = nil
	file_pkg_pb_federation_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FederationClient is the client API for Federation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FederationClient interface {
	Fetch(ctx context.Context, in *FederationFetchRequest, opts ...grpc.CallOption) (*FederationFetchResponse, error)
}

type federationClient struct {
	cc grpc.ClientConnInterface
}

func NewFederationClient(cc grpc.ClientConnInterface) FederationClient {
	return &federationClient{cc}
}

func (c *federationClient) Fetch(ctx context.Context, in *FederationFetchRequest, opts ...grpc.CallOption) (*FederationFetchResponse, error) {
	out := new(FederationFetchResponse)
	err := c.cc.Invoke(ctx, "/Federation/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederationServer is the server API for Federation service.
type FederationServer interface {
	Fetch(context.Context, *FederationFetchRequest) (*FederationFetchResponse, error)
}

// UnimplementedFederationServer can be embedded to have forward compatible implementations.
type UnimplementedFederationServer struct {
}

func (*UnimplementedFederationServer) Fetch(context.Context, *FederationFetchRequest) (*FederationFetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}

func RegisterFederationServer(s *grpc.Server, srv FederationServer) {
	s.RegisterService(&_Federation_serviceDesc, srv)
}

func _Federation_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FederationFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Federation/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServer).Fetch(ctx, req.(*FederationFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Federation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Federation",
	HandlerType: (*FederationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Federation_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/federation.proto",
}
