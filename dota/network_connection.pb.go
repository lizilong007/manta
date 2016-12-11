// Code generated by protoc-gen-go.
// source: network_connection.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ENetworkDisconnectionReason int32

const (
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_INVALID                     ENetworkDisconnectionReason = 0
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SHUTDOWN                    ENetworkDisconnectionReason = 1
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DISCONNECT_BY_USER          ENetworkDisconnectionReason = 2
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DISCONNECT_BY_SERVER        ENetworkDisconnectionReason = 3
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOST                        ENetworkDisconnectionReason = 4
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_OVERFLOW                    ENetworkDisconnectionReason = 5
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_BANNED                ENetworkDisconnectionReason = 6
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_INUSE                 ENetworkDisconnectionReason = 7
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_TICKET                ENetworkDisconnectionReason = 8
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_LOGON                 ENetworkDisconnectionReason = 9
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_AUTHCANCELLED         ENetworkDisconnectionReason = 10
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_AUTHALREADYUSED       ENetworkDisconnectionReason = 11
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_AUTHINVALID           ENetworkDisconnectionReason = 12
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_VACBANSTATE           ENetworkDisconnectionReason = 13
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_LOGGED_IN_ELSEWHERE   ENetworkDisconnectionReason = 14
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_VAC_CHECK_TIMEDOUT    ENetworkDisconnectionReason = 15
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_DROPPED               ENetworkDisconnectionReason = 16
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_OWNERSHIP             ENetworkDisconnectionReason = 17
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVERINFO_OVERFLOW         ENetworkDisconnectionReason = 18
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_TICKMSG_OVERFLOW            ENetworkDisconnectionReason = 19
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STRINGTABLEMSG_OVERFLOW     ENetworkDisconnectionReason = 20
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DELTAENTMSG_OVERFLOW        ENetworkDisconnectionReason = 21
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_TEMPENTMSG_OVERFLOW         ENetworkDisconnectionReason = 22
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SOUNDSMSG_OVERFLOW          ENetworkDisconnectionReason = 23
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SNAPSHOTOVERFLOW            ENetworkDisconnectionReason = 24
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SNAPSHOTERROR               ENetworkDisconnectionReason = 25
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_RELIABLEOVERFLOW            ENetworkDisconnectionReason = 26
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BADDELTATICK                ENetworkDisconnectionReason = 27
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_NOMORESPLITS                ENetworkDisconnectionReason = 28
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_TIMEDOUT                    ENetworkDisconnectionReason = 29
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DISCONNECTED                ENetworkDisconnectionReason = 30
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LEAVINGSPLIT                ENetworkDisconnectionReason = 31
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DIFFERENTCLASSTABLES        ENetworkDisconnectionReason = 32
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BADRELAYPASSWORD            ENetworkDisconnectionReason = 33
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BADSPECTATORPASSWORD        ENetworkDisconnectionReason = 34
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVRESTRICTED              ENetworkDisconnectionReason = 35
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_NOSPECTATORS                ENetworkDisconnectionReason = 36
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVUNAVAILABLE             ENetworkDisconnectionReason = 37
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVSTOP                    ENetworkDisconnectionReason = 38
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_KICKED                      ENetworkDisconnectionReason = 39
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BANADDED                    ENetworkDisconnectionReason = 40
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_KICKBANADDED                ENetworkDisconnectionReason = 41
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HLTVDIRECT                  ENetworkDisconnectionReason = 42
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_PURESERVER_CLIENTEXTRA      ENetworkDisconnectionReason = 43
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_PURESERVER_MISMATCH         ENetworkDisconnectionReason = 44
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_USERCMD                     ENetworkDisconnectionReason = 45
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REJECTED_BY_GAME            ENetworkDisconnectionReason = 46
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_MESSAGE_PARSE_ERROR         ENetworkDisconnectionReason = 47
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_INVALID_MESSAGE_ERROR       ENetworkDisconnectionReason = 48
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_BAD_SERVER_PASSWORD         ENetworkDisconnectionReason = 49
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_DIRECT_CONNECT_RESERVATION  ENetworkDisconnectionReason = 50
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CONNECTION_FAILURE          ENetworkDisconnectionReason = 51
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_NO_PEER_GROUP_HANDLERS      ENetworkDisconnectionReason = 52
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_RECONNECTION                ENetworkDisconnectionReason = 53
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOOPSHUTDOWN                ENetworkDisconnectionReason = 54
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOOPDEACTIVATE              ENetworkDisconnectionReason = 55
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_HOST_ENDGAME                ENetworkDisconnectionReason = 56
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_LOOP_LEVELLOAD_ACTIVATE     ENetworkDisconnectionReason = 57
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CREATE_SERVER_FAILED        ENetworkDisconnectionReason = 58
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_EXITING                     ENetworkDisconnectionReason = 59
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REQUEST_HOSTSTATE_IDLE      ENetworkDisconnectionReason = 60
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REQUEST_HOSTSTATE_HLTVRELAY ENetworkDisconnectionReason = 61
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_CONSISTENCY_FAIL     ENetworkDisconnectionReason = 62
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_UNABLE_TO_CRC_MAP    ENetworkDisconnectionReason = 63
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_NO_MAP               ENetworkDisconnectionReason = 64
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CLIENT_DIFFERENT_MAP        ENetworkDisconnectionReason = 65
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVER_REQUIRES_STEAM       ENetworkDisconnectionReason = 66
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_DENY_MISC             ENetworkDisconnectionReason = 67
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_STEAM_DENY_BAD_ANTI_CHEAT   ENetworkDisconnectionReason = 68
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVER_SHUTDOWN             ENetworkDisconnectionReason = 69
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SPLITPACKET_SEND_OVERFLOW   ENetworkDisconnectionReason = 70
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_REPLAY_INCOMPATIBLE         ENetworkDisconnectionReason = 71
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_CONNECT_REQUEST_TIMEDOUT    ENetworkDisconnectionReason = 72
	ENetworkDisconnectionReason_NETWORK_DISCONNECT_SERVER_INCOMPATIBLE         ENetworkDisconnectionReason = 73
)

var ENetworkDisconnectionReason_name = map[int32]string{
	0:  "NETWORK_DISCONNECT_INVALID",
	1:  "NETWORK_DISCONNECT_SHUTDOWN",
	2:  "NETWORK_DISCONNECT_DISCONNECT_BY_USER",
	3:  "NETWORK_DISCONNECT_DISCONNECT_BY_SERVER",
	4:  "NETWORK_DISCONNECT_LOST",
	5:  "NETWORK_DISCONNECT_OVERFLOW",
	6:  "NETWORK_DISCONNECT_STEAM_BANNED",
	7:  "NETWORK_DISCONNECT_STEAM_INUSE",
	8:  "NETWORK_DISCONNECT_STEAM_TICKET",
	9:  "NETWORK_DISCONNECT_STEAM_LOGON",
	10: "NETWORK_DISCONNECT_STEAM_AUTHCANCELLED",
	11: "NETWORK_DISCONNECT_STEAM_AUTHALREADYUSED",
	12: "NETWORK_DISCONNECT_STEAM_AUTHINVALID",
	13: "NETWORK_DISCONNECT_STEAM_VACBANSTATE",
	14: "NETWORK_DISCONNECT_STEAM_LOGGED_IN_ELSEWHERE",
	15: "NETWORK_DISCONNECT_STEAM_VAC_CHECK_TIMEDOUT",
	16: "NETWORK_DISCONNECT_STEAM_DROPPED",
	17: "NETWORK_DISCONNECT_STEAM_OWNERSHIP",
	18: "NETWORK_DISCONNECT_SERVERINFO_OVERFLOW",
	19: "NETWORK_DISCONNECT_TICKMSG_OVERFLOW",
	20: "NETWORK_DISCONNECT_STRINGTABLEMSG_OVERFLOW",
	21: "NETWORK_DISCONNECT_DELTAENTMSG_OVERFLOW",
	22: "NETWORK_DISCONNECT_TEMPENTMSG_OVERFLOW",
	23: "NETWORK_DISCONNECT_SOUNDSMSG_OVERFLOW",
	24: "NETWORK_DISCONNECT_SNAPSHOTOVERFLOW",
	25: "NETWORK_DISCONNECT_SNAPSHOTERROR",
	26: "NETWORK_DISCONNECT_RELIABLEOVERFLOW",
	27: "NETWORK_DISCONNECT_BADDELTATICK",
	28: "NETWORK_DISCONNECT_NOMORESPLITS",
	29: "NETWORK_DISCONNECT_TIMEDOUT",
	30: "NETWORK_DISCONNECT_DISCONNECTED",
	31: "NETWORK_DISCONNECT_LEAVINGSPLIT",
	32: "NETWORK_DISCONNECT_DIFFERENTCLASSTABLES",
	33: "NETWORK_DISCONNECT_BADRELAYPASSWORD",
	34: "NETWORK_DISCONNECT_BADSPECTATORPASSWORD",
	35: "NETWORK_DISCONNECT_HLTVRESTRICTED",
	36: "NETWORK_DISCONNECT_NOSPECTATORS",
	37: "NETWORK_DISCONNECT_HLTVUNAVAILABLE",
	38: "NETWORK_DISCONNECT_HLTVSTOP",
	39: "NETWORK_DISCONNECT_KICKED",
	40: "NETWORK_DISCONNECT_BANADDED",
	41: "NETWORK_DISCONNECT_KICKBANADDED",
	42: "NETWORK_DISCONNECT_HLTVDIRECT",
	43: "NETWORK_DISCONNECT_PURESERVER_CLIENTEXTRA",
	44: "NETWORK_DISCONNECT_PURESERVER_MISMATCH",
	45: "NETWORK_DISCONNECT_USERCMD",
	46: "NETWORK_DISCONNECT_REJECTED_BY_GAME",
	47: "NETWORK_DISCONNECT_MESSAGE_PARSE_ERROR",
	48: "NETWORK_DISCONNECT_INVALID_MESSAGE_ERROR",
	49: "NETWORK_DISCONNECT_BAD_SERVER_PASSWORD",
	50: "NETWORK_DISCONNECT_DIRECT_CONNECT_RESERVATION",
	51: "NETWORK_DISCONNECT_CONNECTION_FAILURE",
	52: "NETWORK_DISCONNECT_NO_PEER_GROUP_HANDLERS",
	53: "NETWORK_DISCONNECT_RECONNECTION",
	54: "NETWORK_DISCONNECT_LOOPSHUTDOWN",
	55: "NETWORK_DISCONNECT_LOOPDEACTIVATE",
	56: "NETWORK_DISCONNECT_HOST_ENDGAME",
	57: "NETWORK_DISCONNECT_LOOP_LEVELLOAD_ACTIVATE",
	58: "NETWORK_DISCONNECT_CREATE_SERVER_FAILED",
	59: "NETWORK_DISCONNECT_EXITING",
	60: "NETWORK_DISCONNECT_REQUEST_HOSTSTATE_IDLE",
	61: "NETWORK_DISCONNECT_REQUEST_HOSTSTATE_HLTVRELAY",
	62: "NETWORK_DISCONNECT_CLIENT_CONSISTENCY_FAIL",
	63: "NETWORK_DISCONNECT_CLIENT_UNABLE_TO_CRC_MAP",
	64: "NETWORK_DISCONNECT_CLIENT_NO_MAP",
	65: "NETWORK_DISCONNECT_CLIENT_DIFFERENT_MAP",
	66: "NETWORK_DISCONNECT_SERVER_REQUIRES_STEAM",
	67: "NETWORK_DISCONNECT_STEAM_DENY_MISC",
	68: "NETWORK_DISCONNECT_STEAM_DENY_BAD_ANTI_CHEAT",
	69: "NETWORK_DISCONNECT_SERVER_SHUTDOWN",
	70: "NETWORK_DISCONNECT_SPLITPACKET_SEND_OVERFLOW",
	71: "NETWORK_DISCONNECT_REPLAY_INCOMPATIBLE",
	72: "NETWORK_DISCONNECT_CONNECT_REQUEST_TIMEDOUT",
	73: "NETWORK_DISCONNECT_SERVER_INCOMPATIBLE",
}
var ENetworkDisconnectionReason_value = map[string]int32{
	"NETWORK_DISCONNECT_INVALID":                     0,
	"NETWORK_DISCONNECT_SHUTDOWN":                    1,
	"NETWORK_DISCONNECT_DISCONNECT_BY_USER":          2,
	"NETWORK_DISCONNECT_DISCONNECT_BY_SERVER":        3,
	"NETWORK_DISCONNECT_LOST":                        4,
	"NETWORK_DISCONNECT_OVERFLOW":                    5,
	"NETWORK_DISCONNECT_STEAM_BANNED":                6,
	"NETWORK_DISCONNECT_STEAM_INUSE":                 7,
	"NETWORK_DISCONNECT_STEAM_TICKET":                8,
	"NETWORK_DISCONNECT_STEAM_LOGON":                 9,
	"NETWORK_DISCONNECT_STEAM_AUTHCANCELLED":         10,
	"NETWORK_DISCONNECT_STEAM_AUTHALREADYUSED":       11,
	"NETWORK_DISCONNECT_STEAM_AUTHINVALID":           12,
	"NETWORK_DISCONNECT_STEAM_VACBANSTATE":           13,
	"NETWORK_DISCONNECT_STEAM_LOGGED_IN_ELSEWHERE":   14,
	"NETWORK_DISCONNECT_STEAM_VAC_CHECK_TIMEDOUT":    15,
	"NETWORK_DISCONNECT_STEAM_DROPPED":               16,
	"NETWORK_DISCONNECT_STEAM_OWNERSHIP":             17,
	"NETWORK_DISCONNECT_SERVERINFO_OVERFLOW":         18,
	"NETWORK_DISCONNECT_TICKMSG_OVERFLOW":            19,
	"NETWORK_DISCONNECT_STRINGTABLEMSG_OVERFLOW":     20,
	"NETWORK_DISCONNECT_DELTAENTMSG_OVERFLOW":        21,
	"NETWORK_DISCONNECT_TEMPENTMSG_OVERFLOW":         22,
	"NETWORK_DISCONNECT_SOUNDSMSG_OVERFLOW":          23,
	"NETWORK_DISCONNECT_SNAPSHOTOVERFLOW":            24,
	"NETWORK_DISCONNECT_SNAPSHOTERROR":               25,
	"NETWORK_DISCONNECT_RELIABLEOVERFLOW":            26,
	"NETWORK_DISCONNECT_BADDELTATICK":                27,
	"NETWORK_DISCONNECT_NOMORESPLITS":                28,
	"NETWORK_DISCONNECT_TIMEDOUT":                    29,
	"NETWORK_DISCONNECT_DISCONNECTED":                30,
	"NETWORK_DISCONNECT_LEAVINGSPLIT":                31,
	"NETWORK_DISCONNECT_DIFFERENTCLASSTABLES":        32,
	"NETWORK_DISCONNECT_BADRELAYPASSWORD":            33,
	"NETWORK_DISCONNECT_BADSPECTATORPASSWORD":        34,
	"NETWORK_DISCONNECT_HLTVRESTRICTED":              35,
	"NETWORK_DISCONNECT_NOSPECTATORS":                36,
	"NETWORK_DISCONNECT_HLTVUNAVAILABLE":             37,
	"NETWORK_DISCONNECT_HLTVSTOP":                    38,
	"NETWORK_DISCONNECT_KICKED":                      39,
	"NETWORK_DISCONNECT_BANADDED":                    40,
	"NETWORK_DISCONNECT_KICKBANADDED":                41,
	"NETWORK_DISCONNECT_HLTVDIRECT":                  42,
	"NETWORK_DISCONNECT_PURESERVER_CLIENTEXTRA":      43,
	"NETWORK_DISCONNECT_PURESERVER_MISMATCH":         44,
	"NETWORK_DISCONNECT_USERCMD":                     45,
	"NETWORK_DISCONNECT_REJECTED_BY_GAME":            46,
	"NETWORK_DISCONNECT_MESSAGE_PARSE_ERROR":         47,
	"NETWORK_DISCONNECT_INVALID_MESSAGE_ERROR":       48,
	"NETWORK_DISCONNECT_BAD_SERVER_PASSWORD":         49,
	"NETWORK_DISCONNECT_DIRECT_CONNECT_RESERVATION":  50,
	"NETWORK_DISCONNECT_CONNECTION_FAILURE":          51,
	"NETWORK_DISCONNECT_NO_PEER_GROUP_HANDLERS":      52,
	"NETWORK_DISCONNECT_RECONNECTION":                53,
	"NETWORK_DISCONNECT_LOOPSHUTDOWN":                54,
	"NETWORK_DISCONNECT_LOOPDEACTIVATE":              55,
	"NETWORK_DISCONNECT_HOST_ENDGAME":                56,
	"NETWORK_DISCONNECT_LOOP_LEVELLOAD_ACTIVATE":     57,
	"NETWORK_DISCONNECT_CREATE_SERVER_FAILED":        58,
	"NETWORK_DISCONNECT_EXITING":                     59,
	"NETWORK_DISCONNECT_REQUEST_HOSTSTATE_IDLE":      60,
	"NETWORK_DISCONNECT_REQUEST_HOSTSTATE_HLTVRELAY": 61,
	"NETWORK_DISCONNECT_CLIENT_CONSISTENCY_FAIL":     62,
	"NETWORK_DISCONNECT_CLIENT_UNABLE_TO_CRC_MAP":    63,
	"NETWORK_DISCONNECT_CLIENT_NO_MAP":               64,
	"NETWORK_DISCONNECT_CLIENT_DIFFERENT_MAP":        65,
	"NETWORK_DISCONNECT_SERVER_REQUIRES_STEAM":       66,
	"NETWORK_DISCONNECT_STEAM_DENY_MISC":             67,
	"NETWORK_DISCONNECT_STEAM_DENY_BAD_ANTI_CHEAT":   68,
	"NETWORK_DISCONNECT_SERVER_SHUTDOWN":             69,
	"NETWORK_DISCONNECT_SPLITPACKET_SEND_OVERFLOW":   70,
	"NETWORK_DISCONNECT_REPLAY_INCOMPATIBLE":         71,
	"NETWORK_DISCONNECT_CONNECT_REQUEST_TIMEDOUT":    72,
	"NETWORK_DISCONNECT_SERVER_INCOMPATIBLE":         73,
}

func (x ENetworkDisconnectionReason) Enum() *ENetworkDisconnectionReason {
	p := new(ENetworkDisconnectionReason)
	*p = x
	return p
}
func (x ENetworkDisconnectionReason) String() string {
	return proto.EnumName(ENetworkDisconnectionReason_name, int32(x))
}
func (x *ENetworkDisconnectionReason) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ENetworkDisconnectionReason_value, data, "ENetworkDisconnectionReason")
	if err != nil {
		return err
	}
	*x = ENetworkDisconnectionReason(value)
	return nil
}
func (ENetworkDisconnectionReason) EnumDescriptor() ([]byte, []int) { return fileDescriptor32, []int{0} }

var E_NetworkConnectionToken = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50500,
	Name:          "dota.network_connection_token",
	Tag:           "bytes,50500,opt,name=network_connection_token,json=networkConnectionToken",
}

func init() {
	proto.RegisterEnum("dota.ENetworkDisconnectionReason", ENetworkDisconnectionReason_name, ENetworkDisconnectionReason_value)
	proto.RegisterExtension(E_NetworkConnectionToken)
}

func init() { proto.RegisterFile("network_connection.proto", fileDescriptor32) }

var fileDescriptor32 = []byte{
	// 1935 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x98, 0x6b, 0x77, 0xdb, 0xb6,
	0x19, 0xc7, 0x9b, 0xa5, 0xbb, 0x94, 0xbb, 0x61, 0xec, 0xd6, 0x30, 0x74, 0xe3, 0x4b, 0x9c, 0x8b,
	0x9d, 0x8b, 0xd3, 0xee, 0xbe, 0xae, 0xed, 0x0a, 0x91, 0xb0, 0x84, 0x98, 0x02, 0x19, 0x00, 0x92,
	0xe3, 0x9d, 0xee, 0xf0, 0x30, 0x22, 0x6c, 0x73, 0x91, 0x49, 0x8d, 0xa4, 0xec, 0xfa, 0xdd, 0x5e,
	0xef, 0xcd, 0xbe, 0x03, 0x3f, 0xcb, 0x3e, 0x02, 0x3f, 0xd0, 0x0e, 0x48, 0x99, 0x8a, 0x7b, 0x00,
	0xc5, 0xaf, 0xac, 0x63, 0x41, 0x3f, 0xfe, 0x09, 0x3c, 0xcf, 0xff, 0x79, 0x1e, 0x18, 0x56, 0x2a,
	0xca, 0x8b, 0x2c, 0x7f, 0x1b, 0x4e, 0xb2, 0x34, 0x15, 0x93, 0x32, 0xc9, 0xd2, 0xbd, 0x59, 0x9e,
	0x95, 0x99, 0xf9, 0x61, 0x9c, 0x95, 0x91, 0xbd, 0x79, 0x92, 0x65, 0x27, 0x53, 0xf1, 0xa2, 0xf9,
	0xdf, 0x9b, 0xf9, 0xf1, 0x8b, 0x58, 0x14, 0x93, 0x3c, 0x99, 0x95, 0x59, 0xde, 0xae, 0x7b, 0xf2,
	0xdf, 0xe7, 0xc6, 0x1a, 0x22, 0x2d, 0xc5, 0x4d, 0x8a, 0x25, 0x87, 0x8a, 0xa8, 0xc8, 0x52, 0x73,
	0xdd, 0xb0, 0x09, 0xe2, 0x87, 0x3e, 0x3d, 0x08, 0x5d, 0xcc, 0x1c, 0x9f, 0x10, 0xe4, 0xf0, 0x10,
	0x93, 0x31, 0xf4, 0xb0, 0x0b, 0x3e, 0x30, 0x37, 0x8c, 0x35, 0xc5, 0xf7, 0x6c, 0x30, 0xe2, 0xae,
	0x7f, 0x48, 0xc0, 0x2d, 0x73, 0xdf, 0x78, 0xa8, 0x58, 0xf0, 0xce, 0xc7, 0xde, 0x51, 0x38, 0x62,
	0x88, 0x82, 0x1f, 0xd8, 0x6b, 0x55, 0x6d, 0xdd, 0xd9, 0xee, 0x47, 0x67, 0x62, 0x84, 0xc3, 0xa5,
	0x98, 0x70, 0x54, 0x88, 0xdc, 0x7c, 0x69, 0x3c, 0x7e, 0x2f, 0x87, 0x21, 0x3a, 0x46, 0x14, 0xdc,
	0xb6, 0xef, 0x55, 0xb5, 0x75, 0x57, 0x41, 0x62, 0x22, 0x3f, 0x17, 0xb9, 0xd9, 0x33, 0xee, 0x28,
	0x58, 0x9e, 0xcf, 0x38, 0xf8, 0xd0, 0x7e, 0x58, 0xd5, 0xd6, 0x96, 0xe2, 0xb7, 0x4e, 0xb7, 0x35,
	0x5e, 0x56, 0x94, 0xe6, 0x4b, 0xe5, 0x8b, 0xfb, 0x63, 0x44, 0xf7, 0x3d, 0xff, 0x10, 0xfc, 0xd0,
	0xde, 0xad, 0x6a, 0xeb, 0xe1, 0x4a, 0x8e, 0x7f, 0x2e, 0xf2, 0xe3, 0x69, 0x76, 0x61, 0x62, 0x63,
	0x43, 0xb5, 0x89, 0x1c, 0xc1, 0x61, 0xd8, 0x83, 0x84, 0x20, 0x17, 0xfc, 0xc8, 0x7e, 0x50, 0xd5,
	0xd6, 0xa6, 0xea, 0x9d, 0x4a, 0x11, 0x9d, 0x61, 0xb7, 0x17, 0xa5, 0xa9, 0x88, 0xcd, 0xbe, 0xb1,
	0xae, 0x45, 0x61, 0x32, 0x62, 0x08, 0xfc, 0xd8, 0xde, 0xae, 0x6a, 0x6b, 0x43, 0x4f, 0xc2, 0xe9,
	0xa8, 0x10, 0x66, 0x7f, 0x85, 0x26, 0x8e, 0x9d, 0x03, 0xc4, 0xc1, 0x4f, 0xec, 0xfb, 0x55, 0x6d,
	0xad, 0xeb, 0x48, 0x3c, 0x99, 0xbc, 0x15, 0xa5, 0x89, 0x56, 0x28, 0xf2, 0xfc, 0xbe, 0x4f, 0xc0,
	0x47, 0xf6, 0x56, 0x55, 0x5b, 0xf7, 0x74, 0x1c, 0x2f, 0x3b, 0xc9, 0x52, 0x73, 0x68, 0x3c, 0xd2,
	0x62, 0xe0, 0x88, 0x0f, 0x1c, 0x48, 0x1c, 0xe4, 0x79, 0xc8, 0x05, 0xc6, 0x4d, 0x70, 0xbe, 0xb1,
	0xb3, 0x12, 0x07, 0x3d, 0x8a, 0xa0, 0x7b, 0x34, 0x62, 0xc8, 0x05, 0x3f, 0xbd, 0x09, 0xf0, 0xc0,
	0x78, 0xb0, 0x12, 0x78, 0x95, 0x32, 0x3f, 0xbb, 0x09, 0x0c, 0xaf, 0x80, 0x8d, 0xa1, 0xd3, 0x83,
	0x84, 0x71, 0xc8, 0x11, 0xf8, 0xb9, 0xbd, 0x51, 0xd5, 0xd6, 0x9a, 0x0e, 0x36, 0x86, 0x8e, 0xc9,
	0x8c, 0x67, 0xab, 0xb6, 0xbf, 0x8f, 0xdc, 0x10, 0x93, 0x10, 0x79, 0x0c, 0x1d, 0x0e, 0x10, 0x45,
	0xe0, 0x17, 0xab, 0xf5, 0xb5, 0xc1, 0xc1, 0x8d, 0xa7, 0xab, 0xf4, 0x85, 0xce, 0x00, 0x39, 0x07,
	0x21, 0xc7, 0x43, 0xe4, 0xfa, 0x23, 0x0e, 0x7e, 0xb9, 0x3a, 0xe4, 0x78, 0x72, 0x26, 0xfc, 0x79,
	0x69, 0x62, 0x63, 0x53, 0x4b, 0x75, 0xa9, 0x1f, 0x04, 0xc8, 0x05, 0x60, 0x35, 0xca, 0xcd, 0xb3,
	0xd9, 0x4c, 0xc4, 0xe6, 0xd0, 0xb8, 0xaf, 0x45, 0xf9, 0x87, 0x04, 0x51, 0x36, 0xc0, 0x01, 0xf8,
	0x95, 0x36, 0xd9, 0x1b, 0x98, 0x7f, 0x91, 0x8a, 0xbc, 0x38, 0x4d, 0x66, 0xe6, 0x48, 0x1d, 0x7c,
	0x8d, 0xdd, 0x60, 0xb2, 0xef, 0x2f, 0xf3, 0xde, 0xd4, 0xe6, 0x7d, 0xeb, 0x3d, 0x38, 0x3d, 0xce,
	0xba, 0xbc, 0x3f, 0x30, 0xb6, 0x15, 0x58, 0x99, 0x5d, 0x43, 0xd6, 0x5f, 0x32, 0x3f, 0xd6, 0xe6,
	0x99, 0x4c, 0xb1, 0xa1, 0x28, 0x8a, 0xe8, 0x44, 0x98, 0x47, 0xc6, 0x13, 0xe5, 0x2b, 0x53, 0x4c,
	0xfa, 0x1c, 0xf6, 0x3c, 0x74, 0x8d, 0xf9, 0x6b, 0xbd, 0xce, 0x32, 0x4f, 0xd2, 0x13, 0x1e, 0xbd,
	0x99, 0x8a, 0x2b, 0x34, 0x53, 0x7b, 0x2f, 0xf2, 0x38, 0x44, 0x84, 0x5f, 0xe3, 0xfe, 0xc6, 0x7e,
	0x54, 0xd5, 0xd6, 0x7d, 0x05, 0xd7, 0x15, 0xd3, 0x32, 0x42, 0x69, 0x79, 0x05, 0x7d, 0xa5, 0xdc,
	0x53, 0x8e, 0x86, 0xc1, 0xf7, 0x99, 0x9f, 0x68, 0x8f, 0x89, 0x8b, 0xb3, 0xd9, 0x3b, 0x48, 0x5f,
	0x59, 0x6b, 0x98, 0x3f, 0x22, 0x2e, 0xbb, 0x46, 0xbc, 0xa3, 0x77, 0xd3, 0x6c, 0x9e, 0xc6, 0xc5,
	0x15, 0x30, 0x50, 0x1e, 0x10, 0x23, 0x30, 0x60, 0x03, 0x9f, 0x77, 0x38, 0xcb, 0x7e, 0x5c, 0xd5,
	0xd6, 0xb6, 0x0a, 0x97, 0x46, 0xb3, 0xe2, 0x34, 0x2b, 0xbb, 0x23, 0x7f, 0xa9, 0x8e, 0xf1, 0x05,
	0x11, 0x51, 0xea, 0x53, 0x70, 0x57, 0xaf, 0x6e, 0x81, 0x43, 0x79, 0x9e, 0xe5, 0x1a, 0x75, 0x14,
	0x79, 0x58, 0x1e, 0x77, 0xa7, 0xce, 0xd6, 0xaa, 0xa3, 0x62, 0x9a, 0xc8, 0x73, 0xee, 0xd4, 0x11,
	0xa5, 0xe9, 0xf7, 0xa0, 0xdb, 0x9c, 0xb5, 0x0c, 0x4c, 0xb0, 0xa6, 0x0d, 0x9c, 0x5e, 0x14, 0x3b,
	0xd3, 0x44, 0xa4, 0x65, 0x73, 0xd2, 0x32, 0x34, 0xcd, 0x81, 0x92, 0x47, 0xfc, 0xa1, 0x4f, 0x11,
	0x0b, 0x3c, 0xcc, 0x19, 0xf8, 0x54, 0x9b, 0xd0, 0x24, 0x1b, 0x66, 0xb9, 0x60, 0xb3, 0x69, 0x52,
	0x16, 0xe6, 0x37, 0xca, 0x72, 0xdb, 0x39, 0xcc, 0x3d, 0xad, 0x11, 0x4a, 0x73, 0x89, 0xa5, 0xbb,
	0xa8, 0xb5, 0x2c, 0x3f, 0x22, 0x17, 0xac, 0x6b, 0xb5, 0x2c, 0x3f, 0x8a, 0x58, 0x43, 0xf2, 0x10,
	0x1c, 0x63, 0xd2, 0x6f, 0x5e, 0x0b, 0x6c, 0x68, 0x49, 0x9e, 0x88, 0xce, 0x93, 0xf4, 0xa4, 0x79,
	0x2d, 0xf3, 0xb5, 0xa6, 0xa9, 0xd9, 0xdf, 0x47, 0x14, 0x11, 0xee, 0x78, 0x90, 0xb1, 0x26, 0x77,
	0x19, 0xd8, 0xb4, 0x9f, 0x56, 0xb5, 0xf5, 0x58, 0xa9, 0xed, 0xf8, 0x58, 0xe4, 0x22, 0x2d, 0x9d,
	0x69, 0x54, 0x14, 0x4d, 0xe2, 0x16, 0x9a, 0xd8, 0xe8, 0x41, 0x97, 0x22, 0x0f, 0x1e, 0x05, 0x90,
	0xb1, 0x43, 0x9f, 0xba, 0x60, 0x4b, 0x1b, 0x1b, 0xbd, 0x28, 0xa6, 0x62, 0x1a, 0x5d, 0x06, 0x51,
	0x51, 0x5c, 0x64, 0x79, 0xac, 0xd1, 0xda, 0x83, 0x2e, 0x0b, 0x90, 0xc3, 0x21, 0xf7, 0x69, 0x47,
	0xbd, 0xaf, 0xd5, 0xda, 0x8b, 0x62, 0x36, 0x13, 0x93, 0x32, 0x2a, 0xb3, 0xbc, 0x23, 0x7b, 0xc6,
	0x96, 0x82, 0x3c, 0xf0, 0xf8, 0x98, 0x22, 0xe9, 0x5f, 0xcd, 0xd9, 0x6c, 0x6b, 0x4d, 0xa0, 0x59,
	0x28, 0x8a, 0x32, 0x4f, 0x56, 0x9c, 0x0e, 0xf1, 0x3b, 0x99, 0x0c, 0x3c, 0x58, 0x11, 0x73, 0x9d,
	0xbc, 0xc2, 0x24, 0xca, 0x22, 0x22, 0x1f, 0x37, 0x22, 0x70, 0x0c, 0xb1, 0x27, 0x4f, 0x06, 0x3c,
	0xd4, 0x3a, 0x5e, 0xb3, 0x32, 0x8d, 0xce, 0xa3, 0x64, 0x2a, 0x0f, 0x45, 0x13, 0xc3, 0x72, 0x15,
	0xe3, 0x7e, 0x00, 0x1e, 0x69, 0x63, 0xb8, 0x59, 0x52, 0x66, 0x33, 0xf3, 0x4b, 0xe3, 0xae, 0x82,
	0x70, 0x20, 0xdb, 0x31, 0x17, 0x3c, 0xd6, 0xb6, 0xbd, 0x07, 0xb2, 0x13, 0x8b, 0x35, 0xcf, 0xef,
	0x41, 0x22, 0xf3, 0xdb, 0x05, 0x3b, 0xda, 0xe7, 0xf7, 0xa2, 0x14, 0xc6, 0xb1, 0x76, 0x6f, 0xe5,
	0xf3, 0x3b, 0xca, 0xae, 0x76, 0x6f, 0xa5, 0x8a, 0x8e, 0xe4, 0x1a, 0xf7, 0x34, 0x7b, 0xe1, 0x62,
	0x8a, 0x1c, 0x0e, 0x9e, 0x68, 0xfb, 0x90, 0x66, 0x51, 0x92, 0x8b, 0x49, 0x69, 0x7e, 0x6b, 0xec,
	0x2a, 0x28, 0xc1, 0x88, 0xa2, 0xb6, 0x36, 0x87, 0x8e, 0x87, 0x11, 0xe1, 0xe8, 0x35, 0xa7, 0x10,
	0x3c, 0xb5, 0x9f, 0x57, 0xb5, 0xb5, 0xab, 0x20, 0x06, 0xf3, 0x5c, 0xb4, 0xe5, 0x39, 0x6c, 0x1d,
	0x0c, 0x7d, 0x57, 0xe6, 0x91, 0x39, 0x56, 0x56, 0xa8, 0x77, 0xe8, 0x43, 0xcc, 0x86, 0x90, 0x3b,
	0x03, 0xf0, 0xcc, 0x7e, 0x52, 0xd5, 0xd6, 0xa3, 0xd5, 0xe8, 0x61, 0x52, 0x9c, 0x45, 0xe5, 0xe4,
	0xd4, 0xfc, 0x5a, 0x39, 0x53, 0xc9, 0x21, 0xc8, 0x19, 0xba, 0xe0, 0xb9, 0xbd, 0x5e, 0xd5, 0x96,
	0xad, 0x99, 0x83, 0x9c, 0xb3, 0xd8, 0x24, 0x1a, 0xdf, 0x7f, 0xd9, 0xb8, 0x98, 0x1c, 0x84, 0xfa,
	0x70, 0x88, 0xc0, 0x9e, 0x36, 0x63, 0xa8, 0xf8, 0x67, 0xe3, 0x64, 0xbd, 0x4b, 0xf9, 0x9d, 0xc9,
	0x95, 0xef, 0x39, 0x44, 0x8c, 0xc1, 0x3e, 0x0a, 0x03, 0x48, 0x19, 0x0a, 0xdb, 0xca, 0xf4, 0xc2,
	0xde, 0xa9, 0x6a, 0xeb, 0x81, 0x02, 0xb9, 0xa8, 0x98, 0x41, 0x94, 0x17, 0xa2, 0xad, 0x4e, 0xaf,
	0x95, 0x1d, 0xf6, 0xa2, 0x0d, 0xee, 0xe8, 0x2d, 0xf7, 0x33, 0xed, 0xfe, 0xe1, 0xf4, 0x3c, 0x9a,
	0x26, 0xf1, 0x02, 0xdf, 0x92, 0xd5, 0x7a, 0x7b, 0xd0, 0x5d, 0x74, 0x64, 0x61, 0x67, 0x44, 0x9f,
	0x6b, 0xf5, 0x4a, 0x23, 0x6a, 0x8e, 0xa5, 0x73, 0xa1, 0xcf, 0x8d, 0xe7, 0x4a, 0x2f, 0x96, 0xd1,
	0x18, 0x2e, 0x37, 0x59, 0x3e, 0x02, 0x72, 0xec, 0x13, 0xf0, 0x5b, 0x93, 0x29, 0xfb, 0x8d, 0xc5,
	0x5f, 0xec, 0x93, 0x70, 0x1f, 0x62, 0x6f, 0x44, 0x11, 0xf8, 0x9d, 0x56, 0xc7, 0x72, 0x1a, 0xdc,
	0x8f, 0x92, 0xe9, 0x3c, 0x97, 0x7d, 0xdc, 0xae, 0xd2, 0xbf, 0xc2, 0x00, 0x21, 0x1a, 0xf6, 0xa9,
	0x3f, 0x0a, 0xc2, 0x01, 0x24, 0xae, 0x87, 0x28, 0x03, 0xbf, 0xd7, 0x6e, 0x1c, 0xc9, 0x02, 0x21,
	0xf2, 0x7e, 0x9e, 0xcd, 0x67, 0x83, 0x28, 0x8d, 0xa7, 0x22, 0x2f, 0xcc, 0x6d, 0x65, 0xfa, 0x52,
	0xb4, 0x54, 0x0c, 0xfe, 0xa0, 0xab, 0x6e, 0xbe, 0x1f, 0x74, 0x53, 0xfd, 0x1f, 0xf5, 0xd5, 0x2d,
	0xcb, 0x66, 0xec, 0x74, 0x5e, 0xc6, 0xd9, 0x45, 0xaa, 0xf1, 0x75, 0x49, 0x72, 0x11, 0x74, 0x38,
	0x1e, 0xcb, 0x11, 0xe6, 0x4f, 0xda, 0x28, 0x95, 0x2c, 0x57, 0x44, 0x93, 0x32, 0x39, 0x8f, 0x4a,
	0xa1, 0xd1, 0x35, 0xf0, 0x19, 0x0f, 0x11, 0x71, 0x9b, 0x88, 0xff, 0xb3, 0x56, 0xd7, 0x20, 0x2b,
	0xca, 0x10, 0xa5, 0x71, 0x13, 0xef, 0xdf, 0x2a, 0x3b, 0x65, 0xa9, 0x2b, 0xf4, 0xd0, 0x18, 0x79,
	0x9e, 0x0f, 0xdd, 0xb0, 0x13, 0xf8, 0x17, 0xfb, 0x59, 0x55, 0x5b, 0x3b, 0x1a, 0x81, 0x9e, 0x38,
	0x17, 0x53, 0x2f, 0x8b, 0x62, 0x78, 0xa5, 0x73, 0xac, 0xac, 0x93, 0x0e, 0x45, 0x90, 0xa3, 0xab,
	0x00, 0x95, 0x71, 0x81, 0x5c, 0xf0, 0x85, 0xfe, 0x92, 0x20, 0x17, 0x51, 0xb9, 0x30, 0x0e, 0x19,
	0x18, 0x8d, 0x63, 0xaa, 0x5c, 0x03, 0xbd, 0xc6, 0x1c, 0x93, 0x3e, 0xf8, 0xab, 0xb6, 0x67, 0x44,
	0xdf, 0x25, 0x65, 0x92, 0x9e, 0xa0, 0xf4, 0x24, 0x49, 0x65, 0xae, 0xef, 0x2a, 0x43, 0xe0, 0xd5,
	0x08, 0xb1, 0x76, 0x37, 0x9b, 0xc9, 0x32, 0xc4, 0xae, 0x87, 0xc0, 0x97, 0x2b, 0x1c, 0xe4, 0x5f,
	0x73, 0x51, 0x94, 0xe1, 0x80, 0xe1, 0x78, 0x2a, 0x77, 0x74, 0xef, 0x46, 0xd4, 0xb6, 0xa6, 0x7b,
	0xf0, 0x08, 0x7c, 0xa5, 0xcd, 0x88, 0x0e, 0xdd, 0x94, 0xf5, 0x69, 0x74, 0x69, 0xbe, 0x52, 0x9e,
	0x57, 0x6b, 0xed, 0x32, 0xdb, 0x18, 0x66, 0x1c, 0x11, 0xe7, 0xa8, 0xd9, 0x56, 0xf0, 0xf5, 0xf5,
	0xc2, 0xd1, 0x1a, 0xba, 0x93, 0xa5, 0x45, 0x52, 0x94, 0x22, 0x9d, 0x5c, 0xca, 0xed, 0x34, 0x5f,
	0x29, 0x07, 0xd8, 0x05, 0x72, 0x44, 0x64, 0x61, 0x0f, 0xb9, 0x1f, 0x3a, 0xd4, 0x09, 0x87, 0x30,
	0x00, 0x7f, 0xb3, 0x37, 0xab, 0xda, 0xfa, 0xf4, 0x3a, 0x73, 0x94, 0xca, 0xaa, 0xce, 0x33, 0x87,
	0x3a, 0xc3, 0x68, 0x66, 0x7e, 0xa5, 0xec, 0xec, 0x17, 0x48, 0xe2, 0x37, 0x9c, 0x6f, 0xec, 0x3b,
	0x55, 0x6d, 0x7d, 0x7c, 0x9d, 0x43, 0x32, 0xf9, 0xf3, 0x03, 0x75, 0xd8, 0xb4, 0x3f, 0xef, 0x3a,
	0xc2, 0x86, 0x02, 0xaf, 0x57, 0x88, 0x45, 0xd3, 0x7d, 0xd5, 0x05, 0x4a, 0xd8, 0x50, 0x7d, 0xbb,
	0xd1, 0x06, 0x9f, 0x3c, 0x16, 0x4c, 0x11, 0x6b, 0xc7, 0x61, 0xd0, 0xbb, 0x5e, 0xf6, 0xdb, 0x98,
	0x93, 0x87, 0x90, 0xe4, 0xa2, 0x19, 0x83, 0x8b, 0x95, 0xd3, 0xb4, 0x8b, 0xc8, 0x91, 0x2c, 0x84,
	0x0e, 0x70, 0x56, 0x4f, 0xd3, 0xae, 0x48, 0x2f, 0x65, 0x0d, 0x9c, 0x98, 0xe1, 0x8a, 0x2b, 0x89,
	0x06, 0x27, 0xad, 0x1c, 0x12, 0x8e, 0x43, 0x67, 0x80, 0x20, 0x07, 0xae, 0xb6, 0x70, 0x2f, 0xc1,
	0xbd, 0x28, 0x86, 0x69, 0x99, 0x38, 0xa7, 0x22, 0x2a, 0x75, 0x7a, 0xdb, 0xd7, 0xef, 0x5c, 0x0c,
	0xe9, 0xf5, 0x36, 0x7b, 0xd0, 0xf9, 0xd8, 0x1b, 0xb5, 0x5e, 0xd9, 0xe8, 0x07, 0xd0, 0x39, 0x40,
	0x12, 0x4d, 0xdc, 0xe5, 0x74, 0xb9, 0x6f, 0x7f, 0x56, 0xd5, 0xd6, 0x33, 0x15, 0x58, 0x76, 0xfd,
	0xb3, 0x68, 0xf2, 0x56, 0xc8, 0x87, 0xa4, 0x71, 0xd8, 0x4d, 0x5e, 0xea, 0x1b, 0x06, 0x8a, 0x02,
	0x0f, 0x1e, 0x85, 0x98, 0x38, 0xfe, 0x30, 0x80, 0x1c, 0xcb, 0x7e, 0xb3, 0xaf, 0x35, 0x0d, 0x2a,
	0x66, 0xd3, 0xe8, 0x12, 0xa7, 0x93, 0xec, 0x6c, 0x16, 0x95, 0x89, 0x6c, 0x39, 0xff, 0xae, 0x8e,
	0xf3, 0xef, 0x25, 0x68, 0x37, 0x46, 0x0d, 0x6e, 0x70, 0x6b, 0xd9, 0x0c, 0x54, 0xd9, 0xbc, 0x5c,
	0x79, 0x29, 0x72, 0x5d, 0x32, 0x7e, 0xef, 0xa5, 0xc8, 0x52, 0xf2, 0x17, 0xff, 0x50, 0xdd, 0x6a,
	0x87, 0x65, 0xf6, 0x56, 0xa4, 0xe6, 0xd6, 0x5e, 0x7b, 0xa1, 0xbd, 0x77, 0x75, 0xa1, 0xbd, 0x87,
	0xd2, 0xf9, 0xd9, 0x38, 0x9a, 0xce, 0x85, 0x3f, 0x93, 0xeb, 0x0a, 0xeb, 0x7f, 0xff, 0xb9, 0xbd,
	0x79, 0x6b, 0xe7, 0x23, 0xfa, 0xc9, 0x02, 0xf2, 0x8e, 0x70, 0x89, 0xe8, 0xdd, 0xfe, 0xf7, 0xad,
	0x0f, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x48, 0x92, 0x2a, 0xda, 0x38, 0x17, 0x00, 0x00,
}
