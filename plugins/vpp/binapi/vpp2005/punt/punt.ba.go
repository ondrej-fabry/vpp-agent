// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/punt.api.json

/*
Package punt is a generated VPP binary API for 'punt' module.

It consists of:
	  5 enums
	  5 aliases
	 11 types
	  2 unions
	 10 messages
	  5 services
*/
package punt

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"

	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2005/ip_types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "punt"
	// APIVersion is the API version of this module.
	APIVersion = "2.2.1"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x31354154
)

type AddressFamily = ip_types.AddressFamily

type IPDscp = ip_types.IPDscp

type IPEcn = ip_types.IPEcn

type IPProto = ip_types.IPProto

// PuntType represents VPP binary API enum 'punt_type'.
type PuntType uint32

const (
	PUNT_API_TYPE_L4        PuntType = 1
	PUNT_API_TYPE_IP_PROTO  PuntType = 2
	PUNT_API_TYPE_EXCEPTION PuntType = 3
)

var PuntType_name = map[uint32]string{
	1: "PUNT_API_TYPE_L4",
	2: "PUNT_API_TYPE_IP_PROTO",
	3: "PUNT_API_TYPE_EXCEPTION",
}

var PuntType_value = map[string]uint32{
	"PUNT_API_TYPE_L4":        1,
	"PUNT_API_TYPE_IP_PROTO":  2,
	"PUNT_API_TYPE_EXCEPTION": 3,
}

func (x PuntType) String() string {
	s, ok := PuntType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

type AddressWithPrefix = ip_types.AddressWithPrefix

type IP4Address = ip_types.IP4Address

type IP4AddressWithPrefix = ip_types.IP4AddressWithPrefix

type IP6Address = ip_types.IP6Address

type IP6AddressWithPrefix = ip_types.IP6AddressWithPrefix

type Address = ip_types.Address

type IP4Prefix = ip_types.IP4Prefix

type IP6Prefix = ip_types.IP6Prefix

type Mprefix = ip_types.Mprefix

type Prefix = ip_types.Prefix

type PrefixMatcher = ip_types.PrefixMatcher

// Punt represents VPP binary API type 'punt'.
type Punt struct {
	Type PuntType
	Punt PuntUnion
}

func (*Punt) GetTypeName() string { return "punt" }

// PuntException represents VPP binary API type 'punt_exception'.
type PuntException struct {
	ID uint32
}

func (*PuntException) GetTypeName() string { return "punt_exception" }

// PuntIPProto represents VPP binary API type 'punt_ip_proto'.
type PuntIPProto struct {
	Af       AddressFamily
	Protocol IPProto
}

func (*PuntIPProto) GetTypeName() string { return "punt_ip_proto" }

// PuntL4 represents VPP binary API type 'punt_l4'.
type PuntL4 struct {
	Af       AddressFamily
	Protocol IPProto
	Port     uint16
}

func (*PuntL4) GetTypeName() string { return "punt_l4" }

// PuntReason represents VPP binary API type 'punt_reason'.
type PuntReason struct {
	ID          uint32
	XXX_NameLen uint32 `struc:"sizeof=Name"`
	Name        string
}

func (*PuntReason) GetTypeName() string { return "punt_reason" }

type AddressUnion = ip_types.AddressUnion

// PuntUnion represents VPP binary API union 'punt_union'.
type PuntUnion struct {
	XXX_UnionData [4]byte
}

func (*PuntUnion) GetTypeName() string { return "punt_union" }

func PuntUnionException(a PuntException) (u PuntUnion) {
	u.SetException(a)
	return
}
func (u *PuntUnion) SetException(a PuntException) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *PuntUnion) GetException() (a PuntException) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

func PuntUnionL4(a PuntL4) (u PuntUnion) {
	u.SetL4(a)
	return
}
func (u *PuntUnion) SetL4(a PuntL4) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *PuntUnion) GetL4() (a PuntL4) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

func PuntUnionIPProto(a PuntIPProto) (u PuntUnion) {
	u.SetIPProto(a)
	return
}
func (u *PuntUnion) SetIPProto(a PuntIPProto) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.XXX_UnionData[:], b.Bytes())
}
func (u *PuntUnion) GetIPProto() (a PuntIPProto) {
	var b = bytes.NewReader(u.XXX_UnionData[:])
	struc.Unpack(b, &a)
	return
}

// PuntReasonDetails represents VPP binary API message 'punt_reason_details'.
type PuntReasonDetails struct {
	Reason PuntReason
}

func (m *PuntReasonDetails) Reset()                        { *m = PuntReasonDetails{} }
func (*PuntReasonDetails) GetMessageName() string          { return "punt_reason_details" }
func (*PuntReasonDetails) GetCrcString() string            { return "2c9d4a40" }
func (*PuntReasonDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// PuntReasonDump represents VPP binary API message 'punt_reason_dump'.
type PuntReasonDump struct {
	Reason PuntReason
}

func (m *PuntReasonDump) Reset()                        { *m = PuntReasonDump{} }
func (*PuntReasonDump) GetMessageName() string          { return "punt_reason_dump" }
func (*PuntReasonDump) GetCrcString() string            { return "5c0dd4fe" }
func (*PuntReasonDump) GetMessageType() api.MessageType { return api.RequestMessage }

// PuntSocketDeregister represents VPP binary API message 'punt_socket_deregister'.
type PuntSocketDeregister struct {
	Punt Punt
}

func (m *PuntSocketDeregister) Reset()                        { *m = PuntSocketDeregister{} }
func (*PuntSocketDeregister) GetMessageName() string          { return "punt_socket_deregister" }
func (*PuntSocketDeregister) GetCrcString() string            { return "98a444f4" }
func (*PuntSocketDeregister) GetMessageType() api.MessageType { return api.RequestMessage }

// PuntSocketDeregisterReply represents VPP binary API message 'punt_socket_deregister_reply'.
type PuntSocketDeregisterReply struct {
	Retval int32
}

func (m *PuntSocketDeregisterReply) Reset()                        { *m = PuntSocketDeregisterReply{} }
func (*PuntSocketDeregisterReply) GetMessageName() string          { return "punt_socket_deregister_reply" }
func (*PuntSocketDeregisterReply) GetCrcString() string            { return "e8d4e804" }
func (*PuntSocketDeregisterReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// PuntSocketDetails represents VPP binary API message 'punt_socket_details'.
type PuntSocketDetails struct {
	Punt     Punt
	Pathname string `struc:"[108]byte"`
}

func (m *PuntSocketDetails) Reset()                        { *m = PuntSocketDetails{} }
func (*PuntSocketDetails) GetMessageName() string          { return "punt_socket_details" }
func (*PuntSocketDetails) GetCrcString() string            { return "1de0ce75" }
func (*PuntSocketDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// PuntSocketDump represents VPP binary API message 'punt_socket_dump'.
type PuntSocketDump struct {
	Type PuntType
}

func (m *PuntSocketDump) Reset()                        { *m = PuntSocketDump{} }
func (*PuntSocketDump) GetMessageName() string          { return "punt_socket_dump" }
func (*PuntSocketDump) GetCrcString() string            { return "52974935" }
func (*PuntSocketDump) GetMessageType() api.MessageType { return api.RequestMessage }

// PuntSocketRegister represents VPP binary API message 'punt_socket_register'.
type PuntSocketRegister struct {
	HeaderVersion uint32
	Punt          Punt
	Pathname      string `struc:"[108]byte"`
}

func (m *PuntSocketRegister) Reset()                        { *m = PuntSocketRegister{} }
func (*PuntSocketRegister) GetMessageName() string          { return "punt_socket_register" }
func (*PuntSocketRegister) GetCrcString() string            { return "c8cd10fa" }
func (*PuntSocketRegister) GetMessageType() api.MessageType { return api.RequestMessage }

// PuntSocketRegisterReply represents VPP binary API message 'punt_socket_register_reply'.
type PuntSocketRegisterReply struct {
	Retval   int32
	Pathname string `struc:"[108]byte"`
}

func (m *PuntSocketRegisterReply) Reset()                        { *m = PuntSocketRegisterReply{} }
func (*PuntSocketRegisterReply) GetMessageName() string          { return "punt_socket_register_reply" }
func (*PuntSocketRegisterReply) GetCrcString() string            { return "bd30ae90" }
func (*PuntSocketRegisterReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SetPunt represents VPP binary API message 'set_punt'.
type SetPunt struct {
	IsAdd bool
	Punt  Punt
}

func (m *SetPunt) Reset()                        { *m = SetPunt{} }
func (*SetPunt) GetMessageName() string          { return "set_punt" }
func (*SetPunt) GetCrcString() string            { return "83799618" }
func (*SetPunt) GetMessageType() api.MessageType { return api.RequestMessage }

// SetPuntReply represents VPP binary API message 'set_punt_reply'.
type SetPuntReply struct {
	Retval int32
}

func (m *SetPuntReply) Reset()                        { *m = SetPuntReply{} }
func (*SetPuntReply) GetMessageName() string          { return "set_punt_reply" }
func (*SetPuntReply) GetCrcString() string            { return "e8d4e804" }
func (*SetPuntReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*PuntReasonDetails)(nil), "punt.PuntReasonDetails")
	api.RegisterMessage((*PuntReasonDump)(nil), "punt.PuntReasonDump")
	api.RegisterMessage((*PuntSocketDeregister)(nil), "punt.PuntSocketDeregister")
	api.RegisterMessage((*PuntSocketDeregisterReply)(nil), "punt.PuntSocketDeregisterReply")
	api.RegisterMessage((*PuntSocketDetails)(nil), "punt.PuntSocketDetails")
	api.RegisterMessage((*PuntSocketDump)(nil), "punt.PuntSocketDump")
	api.RegisterMessage((*PuntSocketRegister)(nil), "punt.PuntSocketRegister")
	api.RegisterMessage((*PuntSocketRegisterReply)(nil), "punt.PuntSocketRegisterReply")
	api.RegisterMessage((*SetPunt)(nil), "punt.SetPunt")
	api.RegisterMessage((*SetPuntReply)(nil), "punt.SetPuntReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PuntReasonDetails)(nil),
		(*PuntReasonDump)(nil),
		(*PuntSocketDeregister)(nil),
		(*PuntSocketDeregisterReply)(nil),
		(*PuntSocketDetails)(nil),
		(*PuntSocketDump)(nil),
		(*PuntSocketRegister)(nil),
		(*PuntSocketRegisterReply)(nil),
		(*SetPunt)(nil),
		(*SetPuntReply)(nil),
	}
}

// RPCService represents RPC service API for punt module.
type RPCService interface {
	DumpPuntReason(ctx context.Context, in *PuntReasonDump) (RPCService_DumpPuntReasonClient, error)
	DumpPuntSocket(ctx context.Context, in *PuntSocketDump) (RPCService_DumpPuntSocketClient, error)
	PuntSocketDeregister(ctx context.Context, in *PuntSocketDeregister) (*PuntSocketDeregisterReply, error)
	PuntSocketRegister(ctx context.Context, in *PuntSocketRegister) (*PuntSocketRegisterReply, error)
	SetPunt(ctx context.Context, in *SetPunt) (*SetPuntReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpPuntReason(ctx context.Context, in *PuntReasonDump) (RPCService_DumpPuntReasonClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpPuntReasonClient{stream}
	return x, nil
}

type RPCService_DumpPuntReasonClient interface {
	Recv() (*PuntReasonDetails, error)
}

type serviceClient_DumpPuntReasonClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpPuntReasonClient) Recv() (*PuntReasonDetails, error) {
	m := new(PuntReasonDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpPuntSocket(ctx context.Context, in *PuntSocketDump) (RPCService_DumpPuntSocketClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpPuntSocketClient{stream}
	return x, nil
}

type RPCService_DumpPuntSocketClient interface {
	Recv() (*PuntSocketDetails, error)
}

type serviceClient_DumpPuntSocketClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpPuntSocketClient) Recv() (*PuntSocketDetails, error) {
	m := new(PuntSocketDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) PuntSocketDeregister(ctx context.Context, in *PuntSocketDeregister) (*PuntSocketDeregisterReply, error) {
	out := new(PuntSocketDeregisterReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PuntSocketRegister(ctx context.Context, in *PuntSocketRegister) (*PuntSocketRegisterReply, error) {
	out := new(PuntSocketRegisterReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SetPunt(ctx context.Context, in *SetPunt) (*SetPuntReply, error) {
	out := new(SetPuntReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack