// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: message.proto

package types

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

// The action the message will invoke on the handler
type Action int32

const (
	Action_Invalid_Action Action = 0
	Action_POST           Action = 1
	Action_PUT            Action = 2
	Action_PATCH          Action = 3
	Action_DELETE         Action = 4
	Action_GET            Action = 5
	Action_Reply          Action = 6
	Action_Notify         Action = 7
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "Invalid_Action",
		1: "POST",
		2: "PUT",
		3: "PATCH",
		4: "DELETE",
		5: "GET",
		6: "Reply",
		7: "Notify",
	}
	Action_value = map[string]int32{
		"Invalid_Action": 0,
		"POST":           1,
		"PUT":            2,
		"PATCH":          3,
		"DELETE":         4,
		"GET":            5,
		"Reply":          6,
		"Notify":         7,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

// The priority of the message
type Priority int32

const (
	Priority_P0 Priority = 0
	Priority_P1 Priority = 1
	Priority_P2 Priority = 2
	Priority_P3 Priority = 3
	Priority_P4 Priority = 4
	Priority_P5 Priority = 5
	Priority_P6 Priority = 6
	Priority_P7 Priority = 7
)

// Enum value maps for Priority.
var (
	Priority_name = map[int32]string{
		0: "P0",
		1: "P1",
		2: "P2",
		3: "P3",
		4: "P4",
		5: "P5",
		6: "P6",
		7: "P7",
	}
	Priority_value = map[string]int32{
		"P0": 0,
		"P1": 1,
		"P2": 2,
		"P3": 3,
		"P4": 4,
		"P5": 5,
		"P6": 6,
		"P7": 7,
	}
)

func (x Priority) Enum() *Priority {
	p := new(Priority)
	*p = x
	return p
}

func (x Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[1].Descriptor()
}

func (Priority) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[1]
}

func (x Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Priority.Descriptor instead.
func (Priority) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

// Cast mode
type CastMode int32

const (
	CastMode_Invalid_Cast_mode CastMode = 0
	CastMode_All               CastMode = 1
	CastMode_Single            CastMode = 2
	CastMode_Leader            CastMode = 3
)

// Enum value maps for CastMode.
var (
	CastMode_name = map[int32]string{
		0: "Invalid_Cast_mode",
		1: "All",
		2: "Single",
		3: "Leader",
	}
	CastMode_value = map[string]int32{
		"Invalid_Cast_mode": 0,
		"All":               1,
		"Single":            2,
		"Leader":            3,
	}
)

func (x CastMode) Enum() *CastMode {
	p := new(CastMode)
	*p = x
	return p
}

func (x CastMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CastMode) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[2].Descriptor()
}

func (CastMode) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[2]
}

func (x CastMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CastMode.Descriptor instead.
func (CastMode) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

type HealthState int32

const (
	HealthState_Invalid_State HealthState = 0
	HealthState_Up            HealthState = 1
	HealthState_Down          HealthState = 2
	HealthState_Unreachable   HealthState = 3
)

// Enum value maps for HealthState.
var (
	HealthState_name = map[int32]string{
		0: "Invalid_State",
		1: "Up",
		2: "Down",
		3: "Unreachable",
	}
	HealthState_value = map[string]int32{
		"Invalid_State": 0,
		"Up":            1,
		"Down":          2,
		"Unreachable":   3,
	}
)

func (x HealthState) Enum() *HealthState {
	p := new(HealthState)
	*p = x
	return p
}

func (x HealthState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthState) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[3].Descriptor()
}

func (HealthState) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[3]
}

func (x HealthState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthState.Descriptor instead.
func (HealthState) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

// Message is to transmit a piece of data, securely, from one process to one or more processes via
// Publish/Subscribe or Reuest/Reply
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source uuid identifier of the message sender
	SourceUuid string `protobuf:"bytes,1,opt,name=source_uuid,json=sourceUuid,proto3" json:"source_uuid,omitempty"`
	// The source uuid identifier of the message sender switch
	SourceVnetUuid string `protobuf:"bytes,2,opt,name=source_vnet_uuid,json=sourceVnetUuid,proto3" json:"source_vnet_uuid,omitempty"`
	// The destination uuid identifier of the message, in case of a unicast
	Area int32 `protobuf:"varint,3,opt,name=area,proto3" json:"area,omitempty"`
	// The topic of the multicast
	Topic string `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	// To uniquely identify the source message, the sender process maintain a sequence number.
	Sequence int32 `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Priority of this Message for the handler callback.
	Priority Priority `protobuf:"varint,6,opt,name=priority,proto3,enum=types.Priority" json:"priority,omitempty"`
	// The protobuf marshaled data, encoded to base64 string
	Data string `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	// The type name of the serialized data
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	// Action to invoke with this data on the service point
	Action Action `protobuf:"varint,9,opt,name=action,proto3,enum=types.Action" json:"action,omitempty"`
	// timeout when request/reply
	Timeout int64 `protobuf:"varint,10,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// If this a fail notification, what is the failing info
	FailMsg string `protobuf:"bytes,11,opt,name=fail_msg,json=failMsg,proto3" json:"fail_msg,omitempty"`
	// If this message needs a reply message
	IsRequest bool `protobuf:"varint,12,opt,name=is_request,json=isRequest,proto3" json:"is_request,omitempty"`
	IsReply   bool `protobuf:"varint,13,opt,name=is_reply,json=isReply,proto3" json:"is_reply,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSourceUuid() string {
	if x != nil {
		return x.SourceUuid
	}
	return ""
}

func (x *Message) GetSourceVnetUuid() string {
	if x != nil {
		return x.SourceVnetUuid
	}
	return ""
}

func (x *Message) GetArea() int32 {
	if x != nil {
		return x.Area
	}
	return 0
}

func (x *Message) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Message) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Message) GetPriority() Priority {
	if x != nil {
		return x.Priority
	}
	return Priority_P0
}

func (x *Message) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_Invalid_Action
}

func (x *Message) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Message) GetFailMsg() string {
	if x != nil {
		return x.FailMsg
	}
	return ""
}

func (x *Message) GetIsRequest() bool {
	if x != nil {
		return x.IsRequest
	}
	return false
}

func (x *Message) GetIsReply() bool {
	if x != nil {
		return x.IsReply
	}
	return false
}

// A config for the messages limitation and data size
type VNicConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum data size for a message
	MaxDataSize uint64 `protobuf:"varint,1,opt,name=max_data_size,json=maxDataSize,proto3" json:"max_data_size,omitempty"`
	// The transmit message queue size
	TxQueueSize uint64 `protobuf:"varint,2,opt,name=tx_queue_size,json=txQueueSize,proto3" json:"tx_queue_size,omitempty"`
	// The received message queue size
	RxQueueSize uint64 `protobuf:"varint,3,opt,name=rx_queue_size,json=rxQueueSize,proto3" json:"rx_queue_size,omitempty"`
	// What is the switch port for this configuration
	VnetPort uint32 `protobuf:"varint,4,opt,name=vnet_port,json=vnetPort,proto3" json:"vnet_port,omitempty"`
	// The local uuid
	LocalUuid string `protobuf:"bytes,5,opt,name=local_uuid,json=localUuid,proto3" json:"local_uuid,omitempty"`
	// The remote uuid
	RemoteUuid string `protobuf:"bytes,6,opt,name=remote_uuid,json=remoteUuid,proto3" json:"remote_uuid,omitempty"`
	// The address of the connection initiator, regardless of the side.
	Address string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	// force external mode in case two switches are on the same machine
	// woth different ports
	ForceExternal bool `protobuf:"varint,8,opt,name=force_external,json=forceExternal,proto3" json:"force_external,omitempty"`
	// Custom local alias for this vnic
	LocalAlias string `protobuf:"bytes,9,opt,name=local_alias,json=localAlias,proto3" json:"local_alias,omitempty"`
	// Custom remote alias for this vnic
	RemoteAlias string `protobuf:"bytes,10,opt,name=remote_alias,json=remoteAlias,proto3" json:"remote_alias,omitempty"`
	// Service areas
	ServiceAreas *Areas `protobuf:"bytes,11,opt,name=service_areas,json=serviceAreas,proto3" json:"service_areas,omitempty"`
	Area         int32  `protobuf:"varint,12,opt,name=area,proto3" json:"area,omitempty"`
}

func (x *VNicConfig) Reset() {
	*x = VNicConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VNicConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VNicConfig) ProtoMessage() {}

func (x *VNicConfig) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VNicConfig.ProtoReflect.Descriptor instead.
func (*VNicConfig) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *VNicConfig) GetMaxDataSize() uint64 {
	if x != nil {
		return x.MaxDataSize
	}
	return 0
}

func (x *VNicConfig) GetTxQueueSize() uint64 {
	if x != nil {
		return x.TxQueueSize
	}
	return 0
}

func (x *VNicConfig) GetRxQueueSize() uint64 {
	if x != nil {
		return x.RxQueueSize
	}
	return 0
}

func (x *VNicConfig) GetVnetPort() uint32 {
	if x != nil {
		return x.VnetPort
	}
	return 0
}

func (x *VNicConfig) GetLocalUuid() string {
	if x != nil {
		return x.LocalUuid
	}
	return ""
}

func (x *VNicConfig) GetRemoteUuid() string {
	if x != nil {
		return x.RemoteUuid
	}
	return ""
}

func (x *VNicConfig) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *VNicConfig) GetForceExternal() bool {
	if x != nil {
		return x.ForceExternal
	}
	return false
}

func (x *VNicConfig) GetLocalAlias() string {
	if x != nil {
		return x.LocalAlias
	}
	return ""
}

func (x *VNicConfig) GetRemoteAlias() string {
	if x != nil {
		return x.RemoteAlias
	}
	return ""
}

func (x *VNicConfig) GetServiceAreas() *Areas {
	if x != nil {
		return x.ServiceAreas
	}
	return nil
}

func (x *VNicConfig) GetArea() int32 {
	if x != nil {
		return x.Area
	}
	return 0
}

type Areas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreasMap map[int32]*Area `protobuf:"bytes,11,rep,name=areas_map,json=areasMap,proto3" json:"areas_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Areas) Reset() {
	*x = Areas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Areas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Areas) ProtoMessage() {}

func (x *Areas) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Areas.ProtoReflect.Descriptor instead.
func (*Areas) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *Areas) GetAreasMap() map[int32]*Area {
	if x != nil {
		return x.AreasMap
	}
	return nil
}

type Area struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32             `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Topics map[string]*Addrs `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Area) Reset() {
	*x = Area{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Area) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Area) ProtoMessage() {}

func (x *Area) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Area.ProtoReflect.Descriptor instead.
func (*Area) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *Area) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Area) GetTopics() map[string]*Addrs {
	if x != nil {
		return x.Topics
	}
	return nil
}

type Addrs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuids map[string]bool `protobuf:"bytes,1,rep,name=uuids,proto3" json:"uuids,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Addrs) Reset() {
	*x = Addrs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addrs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addrs) ProtoMessage() {}

func (x *Addrs) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addrs.ProtoReflect.Descriptor instead.
func (*Addrs) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *Addrs) GetUuids() map[string]bool {
	if x != nil {
		return x.Uuids
	}
	return nil
}

type HealthPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AUuid        string      `protobuf:"bytes,1,opt,name=a_uuid,json=aUuid,proto3" json:"a_uuid,omitempty"`
	ZUuid        string      `protobuf:"bytes,2,opt,name=z_uuid,json=zUuid,proto3" json:"z_uuid,omitempty"`
	Alias        string      `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	ServiceAreas *Areas      `protobuf:"bytes,4,opt,name=serviceAreas,proto3" json:"serviceAreas,omitempty"`
	Status       HealthState `protobuf:"varint,5,opt,name=status,proto3,enum=types.HealthState" json:"status,omitempty"`
	Started      int64       `protobuf:"varint,6,opt,name=started,proto3" json:"started,omitempty"`
	LastMsg      int64       `protobuf:"varint,7,opt,name=last_msg,json=lastMsg,proto3" json:"last_msg,omitempty"`
}

func (x *HealthPoint) Reset() {
	*x = HealthPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthPoint) ProtoMessage() {}

func (x *HealthPoint) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthPoint.ProtoReflect.Descriptor instead.
func (*HealthPoint) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *HealthPoint) GetAUuid() string {
	if x != nil {
		return x.AUuid
	}
	return ""
}

func (x *HealthPoint) GetZUuid() string {
	if x != nil {
		return x.ZUuid
	}
	return ""
}

func (x *HealthPoint) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *HealthPoint) GetServiceAreas() *Areas {
	if x != nil {
		return x.ServiceAreas
	}
	return nil
}

func (x *HealthPoint) GetStatus() HealthState {
	if x != nil {
		return x.Status
	}
	return HealthState_Invalid_State
}

func (x *HealthPoint) GetStarted() int64 {
	if x != nil {
		return x.Started
	}
	return 0
}

func (x *HealthPoint) GetLastMsg() int64 {
	if x != nil {
		return x.LastMsg
	}
	return 0
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x85, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x6e,
	0x65, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x6e, 0x65, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61, 0x72, 0x65,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x61, 0x69,
	0x6c, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x69,
	0x6c, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xa1,
	0x03, 0x0a, 0x0a, 0x56, 0x4e, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a,
	0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x44, 0x61, 0x74, 0x61, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x78, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x78, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x78, 0x5f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x78,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6e, 0x65,
	0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x76, 0x6e,
	0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x31, 0x0a, 0x0d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x73,
	0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x22, 0x8a, 0x01, 0x0a, 0x05, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12, 0x37, 0x0a, 0x09,
	0x61, 0x72, 0x65, 0x61, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x41, 0x72,
	0x65, 0x61, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x61, 0x72, 0x65,
	0x61, 0x73, 0x4d, 0x61, 0x70, 0x1a, 0x48, 0x0a, 0x0d, 0x41, 0x72, 0x65, 0x61, 0x73, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x41, 0x72, 0x65, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x98, 0x01, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x2f, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x2e, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x1a, 0x47, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x70, 0x0a, 0x05, 0x41, 0x64,
	0x64, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x75, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x73,
	0x2e, 0x55, 0x75, 0x69, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x75, 0x75, 0x69,
	0x64, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x55, 0x75, 0x69, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe4, 0x01, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x7a, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x7a, 0x55, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x12, 0x30, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x72, 0x65, 0x61, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41,
	0x72, 0x65, 0x61, 0x73, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x72, 0x65,
	0x61, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x2a, 0x66, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x0e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50,
	0x55, 0x54, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x47,
	0x45, 0x54, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x10, 0x06, 0x12,
	0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x10, 0x07, 0x2a, 0x4a, 0x0a, 0x08, 0x50,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x30, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x31, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x32, 0x10, 0x02, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x33, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x34, 0x10, 0x04, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x35, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x36, 0x10, 0x06, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x37, 0x10, 0x07, 0x2a, 0x42, 0x0a, 0x08, 0x43, 0x61, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x43,
	0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x6c,
	0x6c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x10, 0x03, 0x2a, 0x43, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x06, 0x0a,
	0x02, 0x55, 0x70, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x6f, 0x77, 0x6e, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x03,
	0x42, 0x24, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x42, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x07, 0x2e,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_message_proto_goTypes = []interface{}{
	(Action)(0),         // 0: types.Action
	(Priority)(0),       // 1: types.Priority
	(CastMode)(0),       // 2: types.CastMode
	(HealthState)(0),    // 3: types.HealthState
	(*Message)(nil),     // 4: types.Message
	(*VNicConfig)(nil),  // 5: types.VNicConfig
	(*Areas)(nil),       // 6: types.Areas
	(*Area)(nil),        // 7: types.Area
	(*Addrs)(nil),       // 8: types.Addrs
	(*HealthPoint)(nil), // 9: types.HealthPoint
	nil,                 // 10: types.Areas.AreasMapEntry
	nil,                 // 11: types.Area.TopicsEntry
	nil,                 // 12: types.Addrs.UuidsEntry
}
var file_message_proto_depIdxs = []int32{
	1,  // 0: types.Message.priority:type_name -> types.Priority
	0,  // 1: types.Message.action:type_name -> types.Action
	6,  // 2: types.VNicConfig.service_areas:type_name -> types.Areas
	10, // 3: types.Areas.areas_map:type_name -> types.Areas.AreasMapEntry
	11, // 4: types.Area.topics:type_name -> types.Area.TopicsEntry
	12, // 5: types.Addrs.uuids:type_name -> types.Addrs.UuidsEntry
	6,  // 6: types.HealthPoint.serviceAreas:type_name -> types.Areas
	3,  // 7: types.HealthPoint.status:type_name -> types.HealthState
	7,  // 8: types.Areas.AreasMapEntry.value:type_name -> types.Area
	8,  // 9: types.Area.TopicsEntry.value:type_name -> types.Addrs
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VNicConfig); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Areas); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Area); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addrs); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthPoint); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
