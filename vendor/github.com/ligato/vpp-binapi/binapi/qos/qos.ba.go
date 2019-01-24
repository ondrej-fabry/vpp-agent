// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/qos.api.json

/*
 Package qos is a generated from VPP binary API module 'qos'.

 It contains following objects:
	  8 messages
	  1 type
	  1 enum
	  4 services

*/
package qos

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0xa23036b4

// Services represents VPP binary API services:
//
//	"services": {
//	    "qos_record_enable_disable": {
//	        "reply": "qos_record_enable_disable_reply"
//	    },
//	    "qos_mark_enable_disable": {
//	        "reply": "qos_mark_enable_disable_reply"
//	    },
//	    "qos_egress_map_delete": {
//	        "reply": "qos_egress_map_delete_reply"
//	    },
//	    "qos_egress_map_update": {
//	        "reply": "qos_egress_map_update_reply"
//	    }
//	},
//
type Services interface {
	QosEgressMapDelete(*QosEgressMapDelete) (*QosEgressMapDeleteReply, error)
	QosEgressMapUpdate(*QosEgressMapUpdate) (*QosEgressMapUpdateReply, error)
	QosMarkEnableDisable(*QosMarkEnableDisable) (*QosMarkEnableDisableReply, error)
	QosRecordEnableDisable(*QosRecordEnableDisable) (*QosRecordEnableDisableReply, error)
}

/* Enums */

// QosSource represents VPP binary API enum 'qos_source':
//
//	"qos_source",
//	[
//	    "QOS_API_SOURCE_EXT",
//	    0
//	],
//	[
//	    "QOS_API_SOURCE_VLAN",
//	    1
//	],
//	[
//	    "QOS_API_SOURCE_MPLS",
//	    2
//	],
//	[
//	    "QOS_API_SOURCE_IP",
//	    3
//	],
//	{
//	    "enumtype": "u32"
//	}
//
type QosSource uint32

const (
	QOS_API_SOURCE_EXT  QosSource = 0
	QOS_API_SOURCE_VLAN QosSource = 1
	QOS_API_SOURCE_MPLS QosSource = 2
	QOS_API_SOURCE_IP   QosSource = 3
)

/* Types */

// QosEgressMapRow represents VPP binary API type 'qos_egress_map_row':
//
//	"qos_egress_map_row",
//	[
//	    "u8",
//	    "outputs",
//	    256
//	],
//	{
//	    "crc": "0xd3bbaed6"
//	}
//
type QosEgressMapRow struct {
	Outputs []byte `struc:"[256]byte"`
}

func (*QosEgressMapRow) GetTypeName() string {
	return "qos_egress_map_row"
}
func (*QosEgressMapRow) GetCrcString() string {
	return "d3bbaed6"
}

/* Messages */

// QosRecordEnableDisable represents VPP binary API message 'qos_record_enable_disable':
//
//	"qos_record_enable_disable",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "vl_api_qos_source_t",
//	    "input_source"
//	],
//	[
//	    "u8",
//	    "enable"
//	],
//	{
//	    "crc": "0xf768050f"
//	}
//
type QosRecordEnableDisable struct {
	SwIfIndex   uint32
	InputSource QosSource
	Enable      uint8
}

func (*QosRecordEnableDisable) GetMessageName() string {
	return "qos_record_enable_disable"
}
func (*QosRecordEnableDisable) GetCrcString() string {
	return "f768050f"
}
func (*QosRecordEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// QosRecordEnableDisableReply represents VPP binary API message 'qos_record_enable_disable_reply':
//
//	"qos_record_enable_disable_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type QosRecordEnableDisableReply struct {
	Retval int32
}

func (*QosRecordEnableDisableReply) GetMessageName() string {
	return "qos_record_enable_disable_reply"
}
func (*QosRecordEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*QosRecordEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// QosEgressMapUpdate represents VPP binary API message 'qos_egress_map_update':
//
//	"qos_egress_map_update",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "map_id"
//	],
//	[
//	    "vl_api_qos_egress_map_row_t",
//	    "rows",
//	    4
//	],
//	{
//	    "crc": "0x5d5c3cad"
//	}
//
type QosEgressMapUpdate struct {
	MapID uint32
	Rows  []QosEgressMapRow `struc:"[4]QosEgressMapRow"`
}

func (*QosEgressMapUpdate) GetMessageName() string {
	return "qos_egress_map_update"
}
func (*QosEgressMapUpdate) GetCrcString() string {
	return "5d5c3cad"
}
func (*QosEgressMapUpdate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// QosEgressMapUpdateReply represents VPP binary API message 'qos_egress_map_update_reply':
//
//	"qos_egress_map_update_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type QosEgressMapUpdateReply struct {
	Retval int32
}

func (*QosEgressMapUpdateReply) GetMessageName() string {
	return "qos_egress_map_update_reply"
}
func (*QosEgressMapUpdateReply) GetCrcString() string {
	return "e8d4e804"
}
func (*QosEgressMapUpdateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// QosEgressMapDelete represents VPP binary API message 'qos_egress_map_delete':
//
//	"qos_egress_map_delete",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "map_id"
//	],
//	{
//	    "crc": "0xdaab68c1"
//	}
//
type QosEgressMapDelete struct {
	MapID uint32
}

func (*QosEgressMapDelete) GetMessageName() string {
	return "qos_egress_map_delete"
}
func (*QosEgressMapDelete) GetCrcString() string {
	return "daab68c1"
}
func (*QosEgressMapDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// QosEgressMapDeleteReply represents VPP binary API message 'qos_egress_map_delete_reply':
//
//	"qos_egress_map_delete_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type QosEgressMapDeleteReply struct {
	Retval int32
}

func (*QosEgressMapDeleteReply) GetMessageName() string {
	return "qos_egress_map_delete_reply"
}
func (*QosEgressMapDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*QosEgressMapDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// QosMarkEnableDisable represents VPP binary API message 'qos_mark_enable_disable':
//
//	"qos_mark_enable_disable",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u32",
//	    "map_id"
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "vl_api_qos_source_t",
//	    "output_source"
//	],
//	[
//	    "u8",
//	    "enable"
//	],
//	{
//	    "crc": "0x3990ab06"
//	}
//
type QosMarkEnableDisable struct {
	MapID        uint32
	SwIfIndex    uint32
	OutputSource QosSource
	Enable       uint8
}

func (*QosMarkEnableDisable) GetMessageName() string {
	return "qos_mark_enable_disable"
}
func (*QosMarkEnableDisable) GetCrcString() string {
	return "3990ab06"
}
func (*QosMarkEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// QosMarkEnableDisableReply represents VPP binary API message 'qos_mark_enable_disable_reply':
//
//	"qos_mark_enable_disable_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type QosMarkEnableDisableReply struct {
	Retval int32
}

func (*QosMarkEnableDisableReply) GetMessageName() string {
	return "qos_mark_enable_disable_reply"
}
func (*QosMarkEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*QosMarkEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*QosRecordEnableDisable)(nil), "qos.QosRecordEnableDisable")
	api.RegisterMessage((*QosRecordEnableDisableReply)(nil), "qos.QosRecordEnableDisableReply")
	api.RegisterMessage((*QosEgressMapUpdate)(nil), "qos.QosEgressMapUpdate")
	api.RegisterMessage((*QosEgressMapUpdateReply)(nil), "qos.QosEgressMapUpdateReply")
	api.RegisterMessage((*QosEgressMapDelete)(nil), "qos.QosEgressMapDelete")
	api.RegisterMessage((*QosEgressMapDeleteReply)(nil), "qos.QosEgressMapDeleteReply")
	api.RegisterMessage((*QosMarkEnableDisable)(nil), "qos.QosMarkEnableDisable")
	api.RegisterMessage((*QosMarkEnableDisableReply)(nil), "qos.QosMarkEnableDisableReply")
}
