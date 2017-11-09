// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steamdatagram_auth_messages.proto

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CMsgSteamDatagramCertificate_EKeyType int32

const (
	CMsgSteamDatagramCertificate_INVALID CMsgSteamDatagramCertificate_EKeyType = 0
	CMsgSteamDatagramCertificate_ED25519 CMsgSteamDatagramCertificate_EKeyType = 1
)

var CMsgSteamDatagramCertificate_EKeyType_name = map[int32]string{
	0: "INVALID",
	1: "ED25519",
}
var CMsgSteamDatagramCertificate_EKeyType_value = map[string]int32{
	"INVALID": 0,
	"ED25519": 1,
}

func (x CMsgSteamDatagramCertificate_EKeyType) Enum() *CMsgSteamDatagramCertificate_EKeyType {
	p := new(CMsgSteamDatagramCertificate_EKeyType)
	*p = x
	return p
}
func (x CMsgSteamDatagramCertificate_EKeyType) String() string {
	return proto.EnumName(CMsgSteamDatagramCertificate_EKeyType_name, int32(x))
}
func (x *CMsgSteamDatagramCertificate_EKeyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CMsgSteamDatagramCertificate_EKeyType_value, data, "CMsgSteamDatagramCertificate_EKeyType")
	if err != nil {
		return err
	}
	*x = CMsgSteamDatagramCertificate_EKeyType(value)
	return nil
}
func (CMsgSteamDatagramCertificate_EKeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor36, []int{2, 0}
}

type CMsgSteamDatagramRelayAuthTicket struct {
	TimeExpiry         *uint32                                        `protobuf:"fixed32,1,opt,name=time_expiry,json=timeExpiry" json:"time_expiry,omitempty"`
	AuthorizedSteamId  *uint64                                        `protobuf:"fixed64,2,opt,name=authorized_steam_id,json=authorizedSteamId" json:"authorized_steam_id,omitempty"`
	AuthorizedPublicIp *uint32                                        `protobuf:"fixed32,3,opt,name=authorized_public_ip,json=authorizedPublicIp" json:"authorized_public_ip,omitempty"`
	GameserverSteamId  *uint64                                        `protobuf:"fixed64,4,opt,name=gameserver_steam_id,json=gameserverSteamId" json:"gameserver_steam_id,omitempty"`
	GameserverNetId    *uint64                                        `protobuf:"fixed64,5,opt,name=gameserver_net_id,json=gameserverNetId" json:"gameserver_net_id,omitempty"`
	LegacySignature    []byte                                         `protobuf:"bytes,6,opt,name=legacy_signature,json=legacySignature" json:"legacy_signature,omitempty"`
	AppId              *uint32                                        `protobuf:"varint,7,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	GameserverPopId    *uint32                                        `protobuf:"fixed32,9,opt,name=gameserver_pop_id,json=gameserverPopId" json:"gameserver_pop_id,omitempty"`
	VirtualPort        *uint32                                        `protobuf:"varint,10,opt,name=virtual_port,json=virtualPort" json:"virtual_port,omitempty"`
	ExtraFields        []*CMsgSteamDatagramRelayAuthTicket_ExtraField `protobuf:"bytes,8,rep,name=extra_fields,json=extraFields" json:"extra_fields,omitempty"`
	XXX_unrecognized   []byte                                         `json:"-"`
}

func (m *CMsgSteamDatagramRelayAuthTicket) Reset()         { *m = CMsgSteamDatagramRelayAuthTicket{} }
func (m *CMsgSteamDatagramRelayAuthTicket) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramRelayAuthTicket) ProtoMessage()    {}
func (*CMsgSteamDatagramRelayAuthTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor36, []int{0}
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetTimeExpiry() uint32 {
	if m != nil && m.TimeExpiry != nil {
		return *m.TimeExpiry
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetAuthorizedSteamId() uint64 {
	if m != nil && m.AuthorizedSteamId != nil {
		return *m.AuthorizedSteamId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetAuthorizedPublicIp() uint32 {
	if m != nil && m.AuthorizedPublicIp != nil {
		return *m.AuthorizedPublicIp
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetGameserverSteamId() uint64 {
	if m != nil && m.GameserverSteamId != nil {
		return *m.GameserverSteamId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetGameserverNetId() uint64 {
	if m != nil && m.GameserverNetId != nil {
		return *m.GameserverNetId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetLegacySignature() []byte {
	if m != nil {
		return m.LegacySignature
	}
	return nil
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetGameserverPopId() uint32 {
	if m != nil && m.GameserverPopId != nil {
		return *m.GameserverPopId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetVirtualPort() uint32 {
	if m != nil && m.VirtualPort != nil {
		return *m.VirtualPort
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetExtraFields() []*CMsgSteamDatagramRelayAuthTicket_ExtraField {
	if m != nil {
		return m.ExtraFields
	}
	return nil
}

type CMsgSteamDatagramRelayAuthTicket_ExtraField struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StringValue      *string `protobuf:"bytes,2,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	Int64Value       *int64  `protobuf:"zigzag64,3,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	Fixed64Value     *uint64 `protobuf:"fixed64,5,opt,name=fixed64_value,json=fixed64Value" json:"fixed64_value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) Reset() {
	*m = CMsgSteamDatagramRelayAuthTicket_ExtraField{}
}
func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgSteamDatagramRelayAuthTicket_ExtraField) ProtoMessage() {}
func (*CMsgSteamDatagramRelayAuthTicket_ExtraField) Descriptor() ([]byte, []int) {
	return fileDescriptor36, []int{0, 0}
}

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetInt64Value() int64 {
	if m != nil && m.Int64Value != nil {
		return *m.Int64Value
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetFixed64Value() uint64 {
	if m != nil && m.Fixed64Value != nil {
		return *m.Fixed64Value
	}
	return 0
}

type CMsgSteamDatagramSignedRelayAuthTicket struct {
	ReservedDoNotUse *uint64 `protobuf:"fixed64,1,opt,name=reserved_do_not_use,json=reservedDoNotUse" json:"reserved_do_not_use,omitempty"`
	KeyId            *uint64 `protobuf:"fixed64,2,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	Ticket           []byte  `protobuf:"bytes,3,opt,name=ticket" json:"ticket,omitempty"`
	Signature        []byte  `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgSteamDatagramSignedRelayAuthTicket) Reset() {
	*m = CMsgSteamDatagramSignedRelayAuthTicket{}
}
func (m *CMsgSteamDatagramSignedRelayAuthTicket) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramSignedRelayAuthTicket) ProtoMessage()    {}
func (*CMsgSteamDatagramSignedRelayAuthTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor36, []int{1}
}

func (m *CMsgSteamDatagramSignedRelayAuthTicket) GetReservedDoNotUse() uint64 {
	if m != nil && m.ReservedDoNotUse != nil {
		return *m.ReservedDoNotUse
	}
	return 0
}

func (m *CMsgSteamDatagramSignedRelayAuthTicket) GetKeyId() uint64 {
	if m != nil && m.KeyId != nil {
		return *m.KeyId
	}
	return 0
}

func (m *CMsgSteamDatagramSignedRelayAuthTicket) GetTicket() []byte {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func (m *CMsgSteamDatagramSignedRelayAuthTicket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type CMsgSteamDatagramCertificate struct {
	KeyType                *CMsgSteamDatagramCertificate_EKeyType `protobuf:"varint,1,opt,name=key_type,json=keyType,enum=dota.CMsgSteamDatagramCertificate_EKeyType,def=0" json:"key_type,omitempty"`
	KeyData                []byte                                 `protobuf:"bytes,2,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
	SteamId                *uint64                                `protobuf:"fixed64,4,opt,name=steam_id,json=steamId" json:"steam_id,omitempty"`
	GameserverDatacenterId *uint32                                `protobuf:"fixed32,5,opt,name=gameserver_datacenter_id,json=gameserverDatacenterId" json:"gameserver_datacenter_id,omitempty"`
	TimeCreated            *uint32                                `protobuf:"fixed32,8,opt,name=time_created,json=timeCreated" json:"time_created,omitempty"`
	TimeExpiry             *uint32                                `protobuf:"fixed32,9,opt,name=time_expiry,json=timeExpiry" json:"time_expiry,omitempty"`
	AppId                  *uint32                                `protobuf:"varint,10,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	XXX_unrecognized       []byte                                 `json:"-"`
}

func (m *CMsgSteamDatagramCertificate) Reset()                    { *m = CMsgSteamDatagramCertificate{} }
func (m *CMsgSteamDatagramCertificate) String() string            { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramCertificate) ProtoMessage()               {}
func (*CMsgSteamDatagramCertificate) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{2} }

const Default_CMsgSteamDatagramCertificate_KeyType CMsgSteamDatagramCertificate_EKeyType = CMsgSteamDatagramCertificate_INVALID

func (m *CMsgSteamDatagramCertificate) GetKeyType() CMsgSteamDatagramCertificate_EKeyType {
	if m != nil && m.KeyType != nil {
		return *m.KeyType
	}
	return Default_CMsgSteamDatagramCertificate_KeyType
}

func (m *CMsgSteamDatagramCertificate) GetKeyData() []byte {
	if m != nil {
		return m.KeyData
	}
	return nil
}

func (m *CMsgSteamDatagramCertificate) GetSteamId() uint64 {
	if m != nil && m.SteamId != nil {
		return *m.SteamId
	}
	return 0
}

func (m *CMsgSteamDatagramCertificate) GetGameserverDatacenterId() uint32 {
	if m != nil && m.GameserverDatacenterId != nil {
		return *m.GameserverDatacenterId
	}
	return 0
}

func (m *CMsgSteamDatagramCertificate) GetTimeCreated() uint32 {
	if m != nil && m.TimeCreated != nil {
		return *m.TimeCreated
	}
	return 0
}

func (m *CMsgSteamDatagramCertificate) GetTimeExpiry() uint32 {
	if m != nil && m.TimeExpiry != nil {
		return *m.TimeExpiry
	}
	return 0
}

func (m *CMsgSteamDatagramCertificate) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

type CMsgSteamDatagramCertificateSigned struct {
	Cert             []byte  `protobuf:"bytes,4,opt,name=cert" json:"cert,omitempty"`
	CaKeyId          *uint64 `protobuf:"fixed64,5,opt,name=ca_key_id,json=caKeyId" json:"ca_key_id,omitempty"`
	CaSignature      []byte  `protobuf:"bytes,6,opt,name=ca_signature,json=caSignature" json:"ca_signature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgSteamDatagramCertificateSigned) Reset()         { *m = CMsgSteamDatagramCertificateSigned{} }
func (m *CMsgSteamDatagramCertificateSigned) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramCertificateSigned) ProtoMessage()    {}
func (*CMsgSteamDatagramCertificateSigned) Descriptor() ([]byte, []int) {
	return fileDescriptor36, []int{3}
}

func (m *CMsgSteamDatagramCertificateSigned) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *CMsgSteamDatagramCertificateSigned) GetCaKeyId() uint64 {
	if m != nil && m.CaKeyId != nil {
		return *m.CaKeyId
	}
	return 0
}

func (m *CMsgSteamDatagramCertificateSigned) GetCaSignature() []byte {
	if m != nil {
		return m.CaSignature
	}
	return nil
}

func init() {
	proto.RegisterType((*CMsgSteamDatagramRelayAuthTicket)(nil), "dota.CMsgSteamDatagramRelayAuthTicket")
	proto.RegisterType((*CMsgSteamDatagramRelayAuthTicket_ExtraField)(nil), "dota.CMsgSteamDatagramRelayAuthTicket.ExtraField")
	proto.RegisterType((*CMsgSteamDatagramSignedRelayAuthTicket)(nil), "dota.CMsgSteamDatagramSignedRelayAuthTicket")
	proto.RegisterType((*CMsgSteamDatagramCertificate)(nil), "dota.CMsgSteamDatagramCertificate")
	proto.RegisterType((*CMsgSteamDatagramCertificateSigned)(nil), "dota.CMsgSteamDatagramCertificateSigned")
	proto.RegisterEnum("dota.CMsgSteamDatagramCertificate_EKeyType", CMsgSteamDatagramCertificate_EKeyType_name, CMsgSteamDatagramCertificate_EKeyType_value)
}

func init() { proto.RegisterFile("steamdatagram_auth_messages.proto", fileDescriptor36) }

var fileDescriptor36 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0x5d, 0xb6, 0xae, 0x69, 0x6f, 0xb2, 0x6f, 0xfd, 0x3c, 0x98, 0xc2, 0x34, 0x89, 0xae, 0x20,
	0x54, 0x40, 0x54, 0x6c, 0x62, 0x08, 0x78, 0x1b, 0x6b, 0x91, 0xaa, 0xc1, 0x98, 0xb2, 0xb1, 0x57,
	0xcb, 0x24, 0x77, 0x99, 0xd5, 0x36, 0x89, 0x1c, 0x67, 0x2c, 0x3c, 0xf1, 0xca, 0xff, 0xe0, 0x87,
	0xf1, 0x23, 0xf8, 0x01, 0xc8, 0x76, 0xda, 0x54, 0xad, 0x80, 0xb7, 0xe4, 0x9c, 0xe3, 0x7b, 0x7d,
	0x7d, 0xce, 0x85, 0xbd, 0x4c, 0x22, 0x9b, 0x84, 0x4c, 0xb2, 0x48, 0xb0, 0x09, 0x65, 0xb9, 0xbc,
	0xa6, 0x13, 0xcc, 0x32, 0x16, 0x61, 0xd6, 0x4b, 0x45, 0x22, 0x13, 0x52, 0x0b, 0x13, 0xc9, 0x3a,
	0xbf, 0x6a, 0xd0, 0x3e, 0xfe, 0x90, 0x45, 0xe7, 0x4a, 0xdf, 0x2f, 0xf5, 0x3e, 0x8e, 0x59, 0x71,
	0x94, 0xcb, 0xeb, 0x0b, 0x1e, 0x8c, 0x50, 0x92, 0xfb, 0xe0, 0x48, 0x3e, 0x41, 0x8a, 0xb7, 0x29,
	0x17, 0x85, 0x67, 0xb5, 0xad, 0xae, 0xed, 0x83, 0x82, 0x06, 0x1a, 0x21, 0x3d, 0xd8, 0x52, 0x2d,
	0x12, 0xc1, 0xbf, 0x62, 0x48, 0x75, 0x6f, 0xca, 0x43, 0x6f, 0xb5, 0x6d, 0x75, 0xeb, 0xfe, 0xff,
	0x15, 0xa5, 0xbb, 0x0c, 0x43, 0xf2, 0x1c, 0xee, 0xcc, 0xe9, 0xd3, 0xfc, 0xf3, 0x98, 0x07, 0x94,
	0xa7, 0xde, 0x9a, 0xae, 0x4c, 0x2a, 0xee, 0x4c, 0x53, 0xc3, 0x54, 0x75, 0x88, 0xd8, 0x04, 0x33,
	0x14, 0x37, 0x28, 0xaa, 0x0e, 0x35, 0xd3, 0xa1, 0xa2, 0xa6, 0x1d, 0x9e, 0xc0, 0x1c, 0x48, 0x63,
	0x94, 0x4a, 0xbd, 0xae, 0xd5, 0x9b, 0x15, 0x71, 0x8a, 0x72, 0x18, 0x92, 0xc7, 0xd0, 0x1a, 0x63,
	0xc4, 0x82, 0x82, 0x66, 0x3c, 0x8a, 0x99, 0xcc, 0x05, 0x7a, 0xf5, 0xb6, 0xd5, 0x75, 0xfd, 0x4d,
	0x83, 0x9f, 0x4f, 0x61, 0x72, 0x17, 0xea, 0x2c, 0x4d, 0x55, 0x2d, 0xbb, 0x6d, 0x75, 0x37, 0xfc,
	0x75, 0x96, 0xa6, 0x4b, 0xdd, 0xd2, 0x44, 0x2b, 0x9a, 0x7a, 0x98, 0xb9, 0x6e, 0x67, 0x89, 0xd2,
	0xee, 0x81, 0x7b, 0xc3, 0x85, 0xcc, 0xd9, 0x98, 0xa6, 0x89, 0x90, 0x1e, 0xe8, 0x42, 0x4e, 0x89,
	0x9d, 0x25, 0x42, 0x92, 0x0b, 0x70, 0xf1, 0x56, 0x0a, 0x46, 0xaf, 0x38, 0x8e, 0xc3, 0xcc, 0x6b,
	0xb4, 0xd7, 0xba, 0xce, 0xc1, 0x7e, 0x4f, 0x39, 0xd6, 0xfb, 0x97, 0x5b, 0xbd, 0x81, 0x3a, 0xfa,
	0x4e, 0x9d, 0xf4, 0x1d, 0x9c, 0x7d, 0x67, 0x3b, 0xdf, 0x2d, 0x80, 0x8a, 0x23, 0x04, 0x6a, 0x31,
	0x9b, 0xa0, 0x76, 0xb3, 0xe9, 0xeb, 0x6f, 0x75, 0xb7, 0x4c, 0x0a, 0x1e, 0x47, 0xf4, 0x86, 0x8d,
	0x73, 0xd4, 0x06, 0x36, 0x7d, 0xc7, 0x60, 0x97, 0x0a, 0x52, 0x59, 0xe0, 0xb1, 0x7c, 0xf9, 0xa2,
	0x54, 0x28, 0xc7, 0x88, 0x0f, 0x1a, 0x32, 0x82, 0x07, 0xb0, 0x71, 0xc5, 0x6f, 0x31, 0x9c, 0x49,
	0xcc, 0xab, 0xbb, 0x25, 0xa8, 0x45, 0x9d, 0x1f, 0x16, 0x3c, 0x5a, 0x1a, 0x44, 0x3d, 0x33, 0x86,
	0x8b, 0xe1, 0x7b, 0x06, 0x5b, 0xc2, 0x3c, 0x60, 0x48, 0xc3, 0x84, 0xc6, 0x89, 0xa4, 0x79, 0x66,
	0xae, 0x5d, 0xf7, 0x5b, 0x53, 0xaa, 0x9f, 0x9c, 0x26, 0xf2, 0x53, 0xa6, 0x1d, 0x1a, 0x61, 0x51,
	0xa5, 0x6f, 0x7d, 0x84, 0xc5, 0x30, 0x24, 0xdb, 0x50, 0x97, 0xba, 0x9e, 0xbe, 0xb1, 0xeb, 0x97,
	0x7f, 0x64, 0x17, 0x9a, 0x95, 0xe9, 0x35, 0x4d, 0x55, 0x40, 0xe7, 0xe7, 0x2a, 0xec, 0x2e, 0x5d,
	0xf3, 0x18, 0x85, 0xe4, 0x57, 0x3c, 0x60, 0x12, 0xc9, 0x47, 0x68, 0xa8, 0x6e, 0xb2, 0x48, 0xcd,
	0x8d, 0xfe, 0x3b, 0x78, 0xfa, 0x07, 0x97, 0xe6, 0x4e, 0xf5, 0x06, 0x27, 0x58, 0x5c, 0x14, 0x29,
	0xbe, 0xb1, 0x87, 0xa7, 0x97, 0x47, 0xef, 0x87, 0x7d, 0xdf, 0x1e, 0x19, 0x84, 0xdc, 0x33, 0x05,
	0xd5, 0xe6, 0xea, 0x01, 0x5c, 0x4d, 0xa9, 0x22, 0x8a, 0x5a, 0xc8, 0xbd, 0x9d, 0x95, 0x69, 0x7f,
	0x05, 0xde, 0x5c, 0xfe, 0xd4, 0xe1, 0x00, 0x63, 0x89, 0x62, 0x1a, 0x7a, 0xdb, 0xdf, 0xae, 0xf8,
	0xfe, 0x8c, 0x36, 0x69, 0xd4, 0xab, 0x1d, 0x08, 0x64, 0x12, 0x43, 0xaf, 0xa1, 0xd5, 0x7a, 0xdd,
	0x8f, 0x0d, 0xb4, 0xb8, 0xfd, 0xcd, 0xa5, 0xed, 0xaf, 0x96, 0x02, 0xe6, 0x96, 0xa2, 0xf3, 0x10,
	0x1a, 0xd3, 0x41, 0x89, 0x03, 0xd3, 0x51, 0x5b, 0x2b, 0xea, 0x67, 0xd0, 0x3f, 0x38, 0x3c, 0xdc,
	0x7f, 0xdd, 0xb2, 0x3a, 0x5f, 0xa0, 0xf3, 0xb7, 0xb7, 0x32, 0x99, 0x50, 0x61, 0x0d, 0x50, 0xc8,
	0xd2, 0x21, 0xfd, 0x4d, 0x76, 0xa0, 0x19, 0x30, 0x5a, 0x9a, 0x6d, 0x42, 0x66, 0x07, 0xec, 0x44,
	0xdb, 0xbd, 0x07, 0x6e, 0xc0, 0x96, 0xd6, 0xd9, 0x09, 0xd8, 0x6c, 0x95, 0xdf, 0xae, 0x7d, 0xb3,
	0x56, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x6a, 0xc0, 0x3f, 0x28, 0x05, 0x00, 0x00,
}
