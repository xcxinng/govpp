// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.6.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/plugins/cdp.api.json

// Package cdp contains generated bindings for API file cdp.api.
//
// Contents:
//   2 messages
//
package cdp

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "cdp"
	APIVersion = "1.0.0"
	VersionCrc = 0x8cfa825e
)

// CdpEnableDisable defines message 'cdp_enable_disable'.
type CdpEnableDisable struct {
	EnableDisable bool `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
}

func (m *CdpEnableDisable) Reset()               { *m = CdpEnableDisable{} }
func (*CdpEnableDisable) GetMessageName() string { return "cdp_enable_disable" }
func (*CdpEnableDisable) GetCrcString() string   { return "2e7b47df" }
func (*CdpEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CdpEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.EnableDisable
	return size
}
func (m *CdpEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.EnableDisable)
	return buf.Bytes(), nil
}
func (m *CdpEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	return nil
}

// CdpEnableDisableReply defines message 'cdp_enable_disable_reply'.
type CdpEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *CdpEnableDisableReply) Reset()               { *m = CdpEnableDisableReply{} }
func (*CdpEnableDisableReply) GetMessageName() string { return "cdp_enable_disable_reply" }
func (*CdpEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*CdpEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CdpEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *CdpEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *CdpEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_cdp_binapi_init() }
func file_cdp_binapi_init() {
	api.RegisterMessage((*CdpEnableDisable)(nil), "cdp_enable_disable_2e7b47df")
	api.RegisterMessage((*CdpEnableDisableReply)(nil), "cdp_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CdpEnableDisable)(nil),
		(*CdpEnableDisableReply)(nil),
	}
}
