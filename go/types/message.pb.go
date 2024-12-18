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
	}
	Action_value = map[string]int32{
		"Invalid_Action": 0,
		"POST":           1,
		"PUT":            2,
		"PATCH":          3,
		"DELETE":         4,
		"GET":            5,
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

// Message is to transmit a piece of data, securely, from one process to one or more processes via
// Publish/Subscribe or Reuest/Reply
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source uuid identifier of the message sender
	SourceUuid string `protobuf:"bytes,1,opt,name=source_uuid,json=sourceUuid,proto3" json:"source_uuid,omitempty"`
	// The source uuid identifier of the message sender switch
	SourceSwitchUuid string `protobuf:"bytes,2,opt,name=source_switch_uuid,json=sourceSwitchUuid,proto3" json:"source_switch_uuid,omitempty"`
	// The destination uuid/or multicast group name of the message receiver.
	Destination string `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	// To uniquely identify the source message, the sender process maintain a sequence number.
	Sequence int32 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Priority of this Message for the handler callback.
	Priority Priority `protobuf:"varint,5,opt,name=priority,proto3,enum=types.Priority" json:"priority,omitempty"`
	// The protobuf marshaled data, encoded to base64 string
	Data string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// The type name of the serialized data
	Type string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	// Action to invoke with this data on the service point
	Action Action `protobuf:"varint,8,opt,name=action,proto3,enum=types.Action" json:"action,omitempty"`
	// timeout when request/reply
	Timeout int64 `protobuf:"varint,9,opt,name=timeout,proto3" json:"timeout,omitempty"`
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

func (x *Message) GetSourceSwitchUuid() string {
	if x != nil {
		return x.SourceSwitchUuid
	}
	return ""
}

func (x *Message) GetDestination() string {
	if x != nil {
		return x.Destination
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

// A config for the messages limitation and data size
type MessagingConfig struct {
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
	SwitchPort uint32 `protobuf:"varint,4,opt,name=switch_port,json=switchPort,proto3" json:"switch_port,omitempty"`
	// Send Status Info
	SendStateInfo bool `protobuf:"varint,5,opt,name=send_state_info,json=sendStateInfo,proto3" json:"send_state_info,omitempty"`
	// Send Status Info Interval in seconds
	SendStateIntervalSeconds int64 `protobuf:"varint,6,opt,name=send_state_interval_seconds,json=sendStateIntervalSeconds,proto3" json:"send_state_interval_seconds,omitempty"`
	// The Edge Uuid
	Local_Uuid string `protobuf:"bytes,7,opt,name=local_Uuid,json=localUuid,proto3" json:"local_Uuid,omitempty"`
	// The zSideUuid
	RemoteUuid string `protobuf:"bytes,8,opt,name=remote_uuid,json=remoteUuid,proto3" json:"remote_uuid,omitempty"`
	// The local/remote address, depending on if this edge is the initiator of the connection
	Address string `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	// Is this a switch side
	IsSwitchSide bool `protobuf:"varint,10,opt,name=is_switch_side,json=isSwitchSide,proto3" json:"is_switch_side,omitempty"`
	// Is the adjacent a switch
	IsAdjacentASwitch bool `protobuf:"varint,11,opt,name=is_adjacent_a_switch,json=isAdjacentASwitch,proto3" json:"is_adjacent_a_switch,omitempty"`
}

func (x *MessagingConfig) Reset() {
	*x = MessagingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagingConfig) ProtoMessage() {}

func (x *MessagingConfig) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MessagingConfig.ProtoReflect.Descriptor instead.
func (*MessagingConfig) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *MessagingConfig) GetMaxDataSize() uint64 {
	if x != nil {
		return x.MaxDataSize
	}
	return 0
}

func (x *MessagingConfig) GetTxQueueSize() uint64 {
	if x != nil {
		return x.TxQueueSize
	}
	return 0
}

func (x *MessagingConfig) GetRxQueueSize() uint64 {
	if x != nil {
		return x.RxQueueSize
	}
	return 0
}

func (x *MessagingConfig) GetSwitchPort() uint32 {
	if x != nil {
		return x.SwitchPort
	}
	return 0
}

func (x *MessagingConfig) GetSendStateInfo() bool {
	if x != nil {
		return x.SendStateInfo
	}
	return false
}

func (x *MessagingConfig) GetSendStateIntervalSeconds() int64 {
	if x != nil {
		return x.SendStateIntervalSeconds
	}
	return 0
}

func (x *MessagingConfig) GetLocal_Uuid() string {
	if x != nil {
		return x.Local_Uuid
	}
	return ""
}

func (x *MessagingConfig) GetRemoteUuid() string {
	if x != nil {
		return x.RemoteUuid
	}
	return ""
}

func (x *MessagingConfig) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MessagingConfig) GetIsSwitchSide() bool {
	if x != nil {
		return x.IsSwitchSide
	}
	return false
}

func (x *MessagingConfig) GetIsAdjacentASwitch() bool {
	if x != nil {
		return x.IsAdjacentASwitch
	}
	return false
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xac, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xb6, 0x03, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x78,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x6d, 0x61, 0x78, 0x44, 0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a,
	0x0d, 0x74, 0x78, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x78, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x78, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x78, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x73, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d,
	0x0a, 0x1b, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x18, 0x73, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x55, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x73, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x69, 0x73, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x53, 0x69, 0x64, 0x65, 0x12, 0x2f, 0x0a,
	0x14, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x5f, 0x73,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x41,
	0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x41, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2a, 0x4f,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x5f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x05, 0x2a,
	0x4a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x06, 0x0a, 0x02, 0x50,
	0x30, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x31, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x50,
	0x32, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x33, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x50,
	0x34, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x35, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x50,
	0x36, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x37, 0x10, 0x07, 0x42, 0x24, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42,
	0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_proto_goTypes = []interface{}{
	(Action)(0),             // 0: types.Action
	(Priority)(0),           // 1: types.Priority
	(*Message)(nil),         // 2: types.Message
	(*MessagingConfig)(nil), // 3: types.MessagingConfig
}
var file_message_proto_depIdxs = []int32{
	1, // 0: types.Message.priority:type_name -> types.Priority
	0, // 1: types.Message.action:type_name -> types.Action
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
			switch v := v.(*MessagingConfig); i {
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
			NumEnums:      2,
			NumMessages:   2,
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
