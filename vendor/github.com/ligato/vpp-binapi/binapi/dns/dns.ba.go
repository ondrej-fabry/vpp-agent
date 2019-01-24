// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: vppapi/dns.api.json

/*
 Package dns is a generated from VPP binary API module 'dns'.

 It contains following objects:
	  8 messages
	  4 services

*/
package dns

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// VlAPIVersion represents version of the binary API module.
const VlAPIVersion = 0x55f2dff8

// Services represents VPP binary API services:
//
//	"services": {
//	    "dns_resolve_name": {
//	        "reply": "dns_resolve_name_reply"
//	    },
//	    "dns_name_server_add_del": {
//	        "reply": "dns_name_server_add_del_reply"
//	    },
//	    "dns_resolve_ip": {
//	        "reply": "dns_resolve_ip_reply"
//	    },
//	    "dns_enable_disable": {
//	        "reply": "dns_enable_disable_reply"
//	    }
//	},
//
type Services interface {
	DNSEnableDisable(*DNSEnableDisable) (*DNSEnableDisableReply, error)
	DNSNameServerAddDel(*DNSNameServerAddDel) (*DNSNameServerAddDelReply, error)
	DNSResolveIP(*DNSResolveIP) (*DNSResolveIPReply, error)
	DNSResolveName(*DNSResolveName) (*DNSResolveNameReply, error)
}

/* Messages */

// DNSEnableDisable represents VPP binary API message 'dns_enable_disable':
//
//	"dns_enable_disable",
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
//	    "u8",
//	    "enable"
//	],
//	{
//	    "crc": "0x8050327d"
//	}
//
type DNSEnableDisable struct {
	Enable uint8
}

func (*DNSEnableDisable) GetMessageName() string {
	return "dns_enable_disable"
}
func (*DNSEnableDisable) GetCrcString() string {
	return "8050327d"
}
func (*DNSEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DNSEnableDisableReply represents VPP binary API message 'dns_enable_disable_reply':
//
//	"dns_enable_disable_reply",
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
type DNSEnableDisableReply struct {
	Retval int32
}

func (*DNSEnableDisableReply) GetMessageName() string {
	return "dns_enable_disable_reply"
}
func (*DNSEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DNSEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DNSNameServerAddDel represents VPP binary API message 'dns_name_server_add_del':
//
//	"dns_name_server_add_del",
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
//	    "u8",
//	    "is_ip6"
//	],
//	[
//	    "u8",
//	    "is_add"
//	],
//	[
//	    "u8",
//	    "server_address",
//	    16
//	],
//	{
//	    "crc": "0x3bb05d8c"
//	}
//
type DNSNameServerAddDel struct {
	IsIP6         uint8
	IsAdd         uint8
	ServerAddress []byte `struc:"[16]byte"`
}

func (*DNSNameServerAddDel) GetMessageName() string {
	return "dns_name_server_add_del"
}
func (*DNSNameServerAddDel) GetCrcString() string {
	return "3bb05d8c"
}
func (*DNSNameServerAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DNSNameServerAddDelReply represents VPP binary API message 'dns_name_server_add_del_reply':
//
//	"dns_name_server_add_del_reply",
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
type DNSNameServerAddDelReply struct {
	Retval int32
}

func (*DNSNameServerAddDelReply) GetMessageName() string {
	return "dns_name_server_add_del_reply"
}
func (*DNSNameServerAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DNSNameServerAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DNSResolveName represents VPP binary API message 'dns_resolve_name':
//
//	"dns_resolve_name",
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
//	    "u8",
//	    "name",
//	    256
//	],
//	{
//	    "crc": "0xc6566676"
//	}
//
type DNSResolveName struct {
	Name []byte `struc:"[256]byte"`
}

func (*DNSResolveName) GetMessageName() string {
	return "dns_resolve_name"
}
func (*DNSResolveName) GetCrcString() string {
	return "c6566676"
}
func (*DNSResolveName) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DNSResolveNameReply represents VPP binary API message 'dns_resolve_name_reply':
//
//	"dns_resolve_name_reply",
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
//	[
//	    "u8",
//	    "ip4_set"
//	],
//	[
//	    "u8",
//	    "ip6_set"
//	],
//	[
//	    "u8",
//	    "ip4_address",
//	    4
//	],
//	[
//	    "u8",
//	    "ip6_address",
//	    16
//	],
//	{
//	    "crc": "0xc2d758c3"
//	}
//
type DNSResolveNameReply struct {
	Retval     int32
	IP4Set     uint8
	IP6Set     uint8
	IP4Address []byte `struc:"[4]byte"`
	IP6Address []byte `struc:"[16]byte"`
}

func (*DNSResolveNameReply) GetMessageName() string {
	return "dns_resolve_name_reply"
}
func (*DNSResolveNameReply) GetCrcString() string {
	return "c2d758c3"
}
func (*DNSResolveNameReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DNSResolveIP represents VPP binary API message 'dns_resolve_ip':
//
//	"dns_resolve_ip",
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
//	    "u8",
//	    "is_ip6"
//	],
//	[
//	    "u8",
//	    "address",
//	    16
//	],
//	{
//	    "crc": "0xae96a1a3"
//	}
//
type DNSResolveIP struct {
	IsIP6   uint8
	Address []byte `struc:"[16]byte"`
}

func (*DNSResolveIP) GetMessageName() string {
	return "dns_resolve_ip"
}
func (*DNSResolveIP) GetCrcString() string {
	return "ae96a1a3"
}
func (*DNSResolveIP) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DNSResolveIPReply represents VPP binary API message 'dns_resolve_ip_reply':
//
//	"dns_resolve_ip_reply",
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
//	[
//	    "u8",
//	    "name",
//	    256
//	],
//	{
//	    "crc": "0x49ed78d6"
//	}
//
type DNSResolveIPReply struct {
	Retval int32
	Name   []byte `struc:"[256]byte"`
}

func (*DNSResolveIPReply) GetMessageName() string {
	return "dns_resolve_ip_reply"
}
func (*DNSResolveIPReply) GetCrcString() string {
	return "49ed78d6"
}
func (*DNSResolveIPReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*DNSEnableDisable)(nil), "dns.DNSEnableDisable")
	api.RegisterMessage((*DNSEnableDisableReply)(nil), "dns.DNSEnableDisableReply")
	api.RegisterMessage((*DNSNameServerAddDel)(nil), "dns.DNSNameServerAddDel")
	api.RegisterMessage((*DNSNameServerAddDelReply)(nil), "dns.DNSNameServerAddDelReply")
	api.RegisterMessage((*DNSResolveName)(nil), "dns.DNSResolveName")
	api.RegisterMessage((*DNSResolveNameReply)(nil), "dns.DNSResolveNameReply")
	api.RegisterMessage((*DNSResolveIP)(nil), "dns.DNSResolveIP")
	api.RegisterMessage((*DNSResolveIPReply)(nil), "dns.DNSResolveIPReply")
}
