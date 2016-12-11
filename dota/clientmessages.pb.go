// Code generated by protoc-gen-go.
// source: clientmessages.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EBaseClientMessages int32

const (
	EBaseClientMessages_CM_CustomGameEvent                  EBaseClientMessages = 280
	EBaseClientMessages_CM_ClientUIEvent                    EBaseClientMessages = 282
	EBaseClientMessages_CM_DevPaletteVisibilityChanged      EBaseClientMessages = 283
	EBaseClientMessages_CM_WorldUIControllerHasPanelChanged EBaseClientMessages = 284
	EBaseClientMessages_CM_RotateAnchor                     EBaseClientMessages = 285
	EBaseClientMessages_CM_MAX_BASE                         EBaseClientMessages = 300
)

var EBaseClientMessages_name = map[int32]string{
	280: "CM_CustomGameEvent",
	282: "CM_ClientUIEvent",
	283: "CM_DevPaletteVisibilityChanged",
	284: "CM_WorldUIControllerHasPanelChanged",
	285: "CM_RotateAnchor",
	300: "CM_MAX_BASE",
}
var EBaseClientMessages_value = map[string]int32{
	"CM_CustomGameEvent":                  280,
	"CM_ClientUIEvent":                    282,
	"CM_DevPaletteVisibilityChanged":      283,
	"CM_WorldUIControllerHasPanelChanged": 284,
	"CM_RotateAnchor":                     285,
	"CM_MAX_BASE":                         300,
}

func (x EBaseClientMessages) Enum() *EBaseClientMessages {
	p := new(EBaseClientMessages)
	*p = x
	return p
}
func (x EBaseClientMessages) String() string {
	return proto.EnumName(EBaseClientMessages_name, int32(x))
}
func (x *EBaseClientMessages) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EBaseClientMessages_value, data, "EBaseClientMessages")
	if err != nil {
		return err
	}
	*x = EBaseClientMessages(value)
	return nil
}
func (EBaseClientMessages) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type EClientUIEvent int32

const (
	EClientUIEvent_EClientUIEvent_Invalid        EClientUIEvent = 0
	EClientUIEvent_EClientUIEvent_DialogFinished EClientUIEvent = 1
	EClientUIEvent_EClientUIEvent_FireOutput     EClientUIEvent = 2
)

var EClientUIEvent_name = map[int32]string{
	0: "EClientUIEvent_Invalid",
	1: "EClientUIEvent_DialogFinished",
	2: "EClientUIEvent_FireOutput",
}
var EClientUIEvent_value = map[string]int32{
	"EClientUIEvent_Invalid":        0,
	"EClientUIEvent_DialogFinished": 1,
	"EClientUIEvent_FireOutput":     2,
}

func (x EClientUIEvent) Enum() *EClientUIEvent {
	p := new(EClientUIEvent)
	*p = x
	return p
}
func (x EClientUIEvent) String() string {
	return proto.EnumName(EClientUIEvent_name, int32(x))
}
func (x *EClientUIEvent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EClientUIEvent_value, data, "EClientUIEvent")
	if err != nil {
		return err
	}
	*x = EClientUIEvent(value)
	return nil
}
func (EClientUIEvent) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type CClientMsg_CustomGameEvent struct {
	EventName        *string `protobuf:"bytes,1,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CClientMsg_CustomGameEvent) Reset()                    { *m = CClientMsg_CustomGameEvent{} }
func (m *CClientMsg_CustomGameEvent) String() string            { return proto.CompactTextString(m) }
func (*CClientMsg_CustomGameEvent) ProtoMessage()               {}
func (*CClientMsg_CustomGameEvent) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CClientMsg_CustomGameEvent) GetEventName() string {
	if m != nil && m.EventName != nil {
		return *m.EventName
	}
	return ""
}

func (m *CClientMsg_CustomGameEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CClientMsg_ClientUIEvent struct {
	Event            *EClientUIEvent `protobuf:"varint,1,opt,name=event,enum=dota.EClientUIEvent,def=0" json:"event,omitempty"`
	EntEhandle       *uint32         `protobuf:"varint,2,opt,name=ent_ehandle,json=entEhandle" json:"ent_ehandle,omitempty"`
	ClientEhandle    *uint32         `protobuf:"varint,3,opt,name=client_ehandle,json=clientEhandle" json:"client_ehandle,omitempty"`
	Data1            *string         `protobuf:"bytes,4,opt,name=data1" json:"data1,omitempty"`
	Data2            *string         `protobuf:"bytes,5,opt,name=data2" json:"data2,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *CClientMsg_ClientUIEvent) Reset()                    { *m = CClientMsg_ClientUIEvent{} }
func (m *CClientMsg_ClientUIEvent) String() string            { return proto.CompactTextString(m) }
func (*CClientMsg_ClientUIEvent) ProtoMessage()               {}
func (*CClientMsg_ClientUIEvent) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

const Default_CClientMsg_ClientUIEvent_Event EClientUIEvent = EClientUIEvent_EClientUIEvent_Invalid

func (m *CClientMsg_ClientUIEvent) GetEvent() EClientUIEvent {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return Default_CClientMsg_ClientUIEvent_Event
}

func (m *CClientMsg_ClientUIEvent) GetEntEhandle() uint32 {
	if m != nil && m.EntEhandle != nil {
		return *m.EntEhandle
	}
	return 0
}

func (m *CClientMsg_ClientUIEvent) GetClientEhandle() uint32 {
	if m != nil && m.ClientEhandle != nil {
		return *m.ClientEhandle
	}
	return 0
}

func (m *CClientMsg_ClientUIEvent) GetData1() string {
	if m != nil && m.Data1 != nil {
		return *m.Data1
	}
	return ""
}

func (m *CClientMsg_ClientUIEvent) GetData2() string {
	if m != nil && m.Data2 != nil {
		return *m.Data2
	}
	return ""
}

type CClientMsg_DevPaletteVisibilityChangedEvent struct {
	Visible          *bool  `protobuf:"varint,1,opt,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CClientMsg_DevPaletteVisibilityChangedEvent) Reset() {
	*m = CClientMsg_DevPaletteVisibilityChangedEvent{}
}
func (m *CClientMsg_DevPaletteVisibilityChangedEvent) String() string {
	return proto.CompactTextString(m)
}
func (*CClientMsg_DevPaletteVisibilityChangedEvent) ProtoMessage() {}
func (*CClientMsg_DevPaletteVisibilityChangedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{2}
}

func (m *CClientMsg_DevPaletteVisibilityChangedEvent) GetVisible() bool {
	if m != nil && m.Visible != nil {
		return *m.Visible
	}
	return false
}

type CClientMsg_WorldUIControllerHasPanelChangedEvent struct {
	HasPanel         *bool   `protobuf:"varint,1,opt,name=has_panel,json=hasPanel" json:"has_panel,omitempty"`
	ClientEhandle    *uint32 `protobuf:"varint,2,opt,name=client_ehandle,json=clientEhandle" json:"client_ehandle,omitempty"`
	LiteralHandType  *uint32 `protobuf:"varint,3,opt,name=literal_hand_type,json=literalHandType" json:"literal_hand_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) Reset() {
	*m = CClientMsg_WorldUIControllerHasPanelChangedEvent{}
}
func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) String() string {
	return proto.CompactTextString(m)
}
func (*CClientMsg_WorldUIControllerHasPanelChangedEvent) ProtoMessage() {}
func (*CClientMsg_WorldUIControllerHasPanelChangedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{3}
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetHasPanel() bool {
	if m != nil && m.HasPanel != nil {
		return *m.HasPanel
	}
	return false
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetClientEhandle() uint32 {
	if m != nil && m.ClientEhandle != nil {
		return *m.ClientEhandle
	}
	return 0
}

func (m *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetLiteralHandType() uint32 {
	if m != nil && m.LiteralHandType != nil {
		return *m.LiteralHandType
	}
	return 0
}

type CClientMsg_RotateAnchor struct {
	Angle            *float32 `protobuf:"fixed32,1,opt,name=angle" json:"angle,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CClientMsg_RotateAnchor) Reset()                    { *m = CClientMsg_RotateAnchor{} }
func (m *CClientMsg_RotateAnchor) String() string            { return proto.CompactTextString(m) }
func (*CClientMsg_RotateAnchor) ProtoMessage()               {}
func (*CClientMsg_RotateAnchor) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *CClientMsg_RotateAnchor) GetAngle() float32 {
	if m != nil && m.Angle != nil {
		return *m.Angle
	}
	return 0
}

func init() {
	proto.RegisterType((*CClientMsg_CustomGameEvent)(nil), "dota.CClientMsg_CustomGameEvent")
	proto.RegisterType((*CClientMsg_ClientUIEvent)(nil), "dota.CClientMsg_ClientUIEvent")
	proto.RegisterType((*CClientMsg_DevPaletteVisibilityChangedEvent)(nil), "dota.CClientMsg_DevPaletteVisibilityChangedEvent")
	proto.RegisterType((*CClientMsg_WorldUIControllerHasPanelChangedEvent)(nil), "dota.CClientMsg_WorldUIControllerHasPanelChangedEvent")
	proto.RegisterType((*CClientMsg_RotateAnchor)(nil), "dota.CClientMsg_RotateAnchor")
	proto.RegisterEnum("dota.EBaseClientMessages", EBaseClientMessages_name, EBaseClientMessages_value)
	proto.RegisterEnum("dota.EClientUIEvent", EClientUIEvent_name, EClientUIEvent_value)
}

func init() { proto.RegisterFile("clientmessages.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x4d, 0x22, 0x9a, 0x29, 0x4d, 0x97, 0x25, 0x50, 0x53, 0x14, 0x28, 0xa9, 0x90, 0xa2,
	0x20, 0x05, 0xc8, 0x91, 0x5b, 0xe2, 0xa6, 0x4d, 0x0e, 0xa6, 0x95, 0xa1, 0xc0, 0xcd, 0x1a, 0xea,
	0x51, 0xbc, 0xd2, 0x66, 0x1d, 0x79, 0x37, 0x91, 0x72, 0xe3, 0x67, 0x20, 0x3e, 0x2e, 0x9c, 0xf9,
	0x09, 0xfc, 0x0a, 0x7e, 0x11, 0xf2, 0x47, 0x20, 0x8e, 0x0a, 0xbd, 0xed, 0xbc, 0xf7, 0xe6, 0x79,
	0xde, 0x78, 0xa0, 0x71, 0x29, 0x05, 0x29, 0x33, 0x25, 0xad, 0x71, 0x42, 0xba, 0x3b, 0x4b, 0x62,
	0x13, 0xf3, 0x4a, 0x18, 0x1b, 0x6c, 0x9d, 0xc1, 0x81, 0xeb, 0x66, 0xb4, 0xa7, 0x27, 0x81, 0x3b,
	0xd7, 0x26, 0x9e, 0x9e, 0xe2, 0x94, 0x86, 0x0b, 0x52, 0x86, 0x37, 0x01, 0x28, 0x7d, 0x04, 0x0a,
	0xa7, 0xe4, 0x58, 0x87, 0x56, 0xbb, 0xe6, 0xd7, 0x32, 0xe4, 0x15, 0x4e, 0x89, 0x73, 0xa8, 0x84,
	0x68, 0xd0, 0xb1, 0x0f, 0xad, 0xf6, 0x2d, 0x3f, 0x7b, 0xb7, 0x7e, 0x59, 0xe0, 0xac, 0x3b, 0x66,
	0xaf, 0x8b, 0x71, 0xee, 0x37, 0x80, 0x6a, 0xd6, 0x9d, 0x59, 0xd5, 0x7b, 0x8d, 0x6e, 0x3a, 0x43,
	0x77, 0x58, 0x12, 0xbd, 0xbc, 0x57, 0xae, 0x83, 0xb1, 0x5a, 0xa0, 0x14, 0xa1, 0x9f, 0xb7, 0xf2,
	0x47, 0xb0, 0x93, 0xa2, 0x14, 0xa1, 0x0a, 0x25, 0x65, 0xdf, 0xde, 0xf5, 0x81, 0x94, 0x19, 0xe6,
	0x08, 0x7f, 0x02, 0xf5, 0x3c, 0xf0, 0x1f, 0xcd, 0x8d, 0x4c, 0xb3, 0x9b, 0xa3, 0x2b, 0x59, 0x03,
	0xaa, 0xe9, 0xc0, 0x2f, 0x9c, 0x4a, 0x16, 0x2b, 0x2f, 0x56, 0x68, 0xcf, 0xa9, 0xfe, 0x45, 0x7b,
	0xad, 0x53, 0x78, 0xba, 0x96, 0xe9, 0x98, 0x16, 0xe7, 0x28, 0xc9, 0x18, 0x7a, 0x2b, 0xb4, 0xf8,
	0x20, 0xa4, 0x30, 0x4b, 0x37, 0x42, 0x35, 0xa1, 0x30, 0x8f, 0xe9, 0xc0, 0xcd, 0x45, 0xca, 0xc8,
	0x7c, 0x67, 0xdb, 0xfe, 0xaa, 0x6c, 0x7d, 0xb7, 0xe0, 0xf9, 0x9a, 0xd3, 0xbb, 0x38, 0x91, 0xe1,
	0xc5, 0xd8, 0x8d, 0x95, 0x49, 0x62, 0x29, 0x29, 0x19, 0xa1, 0x3e, 0x47, 0x45, 0xb2, 0x64, 0xf7,
	0x00, 0x6a, 0x11, 0xea, 0x60, 0x96, 0x12, 0x85, 0xe1, 0x76, 0x54, 0x08, 0xaf, 0x48, 0x6b, 0x5f,
	0x95, 0xb6, 0x03, 0xb7, 0xa5, 0x30, 0x94, 0xa0, 0x0c, 0x52, 0x24, 0x30, 0xcb, 0xd9, 0x6a, 0x2f,
	0x7b, 0x05, 0x31, 0x42, 0x15, 0xbe, 0x59, 0xce, 0xa8, 0xf5, 0x0c, 0xf6, 0xd7, 0x66, 0xf4, 0x63,
	0x83, 0x86, 0xfa, 0xea, 0x32, 0x8a, 0x93, 0x74, 0x3d, 0xa8, 0x26, 0x45, 0x2e, 0xdb, 0xcf, 0x8b,
	0xce, 0x4f, 0x0b, 0xee, 0x0c, 0x07, 0xa8, 0xa9, 0xe8, 0x2a, 0x0e, 0x8d, 0xef, 0x03, 0x77, 0xbd,
	0xcd, 0xa3, 0x62, 0x9f, 0x6c, 0x7e, 0x17, 0x58, 0x4a, 0xac, 0xff, 0x66, 0xf6, 0xd9, 0xe6, 0x47,
	0xf0, 0xd0, 0xf5, 0xfe, 0xb7, 0x5e, 0xf6, 0xc5, 0xe6, 0x6d, 0x38, 0x72, 0xbd, 0x6b, 0x37, 0xc7,
	0xbe, 0xda, 0xbc, 0x01, 0x7b, 0xae, 0x57, 0x9a, 0x9f, 0x7d, 0xb3, 0x39, 0x83, 0x1d, 0xd7, 0x0b,
	0xbc, 0xfe, 0xfb, 0x60, 0xd0, 0x7f, 0x3d, 0x64, 0x3f, 0xec, 0x8e, 0x82, 0x7a, 0xf9, 0xe4, 0xf8,
	0x01, 0xfc, 0xe3, 0x08, 0xd9, 0x16, 0x7f, 0x0c, 0xcd, 0x0d, 0xee, 0x58, 0xa0, 0x8c, 0x27, 0x27,
	0x42, 0x09, 0x1d, 0x51, 0xc8, 0x2c, 0xde, 0x84, 0xfb, 0x1b, 0x92, 0x13, 0x91, 0xd0, 0xd9, 0xdc,
	0xcc, 0xe6, 0x86, 0xd9, 0x83, 0xea, 0xc8, 0xfa, 0x68, 0x6d, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x39, 0x51, 0x7b, 0x90, 0x97, 0x03, 0x00, 0x00,
}
