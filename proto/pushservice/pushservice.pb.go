// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pushservice.proto

package pushservice

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LsNode struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Asn                  int32    `protobuf:"varint,3,opt,name=asn,proto3" json:"asn,omitempty"`
	RouterIp             string   `protobuf:"bytes,4,opt,name=router_ip,json=routerIp,proto3" json:"router_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LsNode) Reset()         { *m = LsNode{} }
func (m *LsNode) String() string { return proto.CompactTextString(m) }
func (*LsNode) ProtoMessage()    {}
func (*LsNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac404c4a13b2757, []int{0}
}

func (m *LsNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LsNode.Unmarshal(m, b)
}
func (m *LsNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LsNode.Marshal(b, m, deterministic)
}
func (m *LsNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LsNode.Merge(m, src)
}
func (m *LsNode) XXX_Size() int {
	return xxx_messageInfo_LsNode.Size(m)
}
func (m *LsNode) XXX_DiscardUnknown() {
	xxx_messageInfo_LsNode.DiscardUnknown(m)
}

var xxx_messageInfo_LsNode proto.InternalMessageInfo

func (m *LsNode) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LsNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LsNode) GetAsn() int32 {
	if m != nil {
		return m.Asn
	}
	return 0
}

func (m *LsNode) GetRouterIp() string {
	if m != nil {
		return m.RouterIp
	}
	return ""
}

type LsNodeSubscription struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LsNodeSubscription) Reset()         { *m = LsNodeSubscription{} }
func (m *LsNodeSubscription) String() string { return proto.CompactTextString(m) }
func (*LsNodeSubscription) ProtoMessage()    {}
func (*LsNodeSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac404c4a13b2757, []int{1}
}

func (m *LsNodeSubscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LsNodeSubscription.Unmarshal(m, b)
}
func (m *LsNodeSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LsNodeSubscription.Marshal(b, m, deterministic)
}
func (m *LsNodeSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LsNodeSubscription.Merge(m, src)
}
func (m *LsNodeSubscription) XXX_Size() int {
	return xxx_messageInfo_LsNodeSubscription.Size(m)
}
func (m *LsNodeSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_LsNodeSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_LsNodeSubscription proto.InternalMessageInfo

func (m *LsNodeSubscription) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type LsNodeEvent struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	LsNode               *LsNode  `protobuf:"bytes,3,opt,name=lsNode,proto3" json:"lsNode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LsNodeEvent) Reset()         { *m = LsNodeEvent{} }
func (m *LsNodeEvent) String() string { return proto.CompactTextString(m) }
func (*LsNodeEvent) ProtoMessage()    {}
func (*LsNodeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac404c4a13b2757, []int{2}
}

func (m *LsNodeEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LsNodeEvent.Unmarshal(m, b)
}
func (m *LsNodeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LsNodeEvent.Marshal(b, m, deterministic)
}
func (m *LsNodeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LsNodeEvent.Merge(m, src)
}
func (m *LsNodeEvent) XXX_Size() int {
	return xxx_messageInfo_LsNodeEvent.Size(m)
}
func (m *LsNodeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LsNodeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LsNodeEvent proto.InternalMessageInfo

func (m *LsNodeEvent) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *LsNodeEvent) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LsNodeEvent) GetLsNode() *LsNode {
	if m != nil {
		return m.LsNode
	}
	return nil
}

type LsLink struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	RouterIp             string   `protobuf:"bytes,2,opt,name=router_ip,json=routerIp,proto3" json:"router_ip,omitempty"`
	PeerIp               string   `protobuf:"bytes,3,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	LocalLinkIp          string   `protobuf:"bytes,4,opt,name=localLink_ip,json=localLinkIp,proto3" json:"localLink_ip,omitempty"`
	RemoteLinkIp         string   `protobuf:"bytes,5,opt,name=remoteLink_ip,json=remoteLinkIp,proto3" json:"remoteLink_ip,omitempty"`
	IgpMetric            int32    `protobuf:"varint,6,opt,name=igp_metric,json=igpMetric,proto3" json:"igp_metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LsLink) Reset()         { *m = LsLink{} }
func (m *LsLink) String() string { return proto.CompactTextString(m) }
func (*LsLink) ProtoMessage()    {}
func (*LsLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac404c4a13b2757, []int{3}
}

func (m *LsLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LsLink.Unmarshal(m, b)
}
func (m *LsLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LsLink.Marshal(b, m, deterministic)
}
func (m *LsLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LsLink.Merge(m, src)
}
func (m *LsLink) XXX_Size() int {
	return xxx_messageInfo_LsLink.Size(m)
}
func (m *LsLink) XXX_DiscardUnknown() {
	xxx_messageInfo_LsLink.DiscardUnknown(m)
}

var xxx_messageInfo_LsLink proto.InternalMessageInfo

func (m *LsLink) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LsLink) GetRouterIp() string {
	if m != nil {
		return m.RouterIp
	}
	return ""
}

func (m *LsLink) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

func (m *LsLink) GetLocalLinkIp() string {
	if m != nil {
		return m.LocalLinkIp
	}
	return ""
}

func (m *LsLink) GetRemoteLinkIp() string {
	if m != nil {
		return m.RemoteLinkIp
	}
	return ""
}

func (m *LsLink) GetIgpMetric() int32 {
	if m != nil {
		return m.IgpMetric
	}
	return 0
}

type LsLinkSubscription struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LsLinkSubscription) Reset()         { *m = LsLinkSubscription{} }
func (m *LsLinkSubscription) String() string { return proto.CompactTextString(m) }
func (*LsLinkSubscription) ProtoMessage()    {}
func (*LsLinkSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac404c4a13b2757, []int{4}
}

func (m *LsLinkSubscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LsLinkSubscription.Unmarshal(m, b)
}
func (m *LsLinkSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LsLinkSubscription.Marshal(b, m, deterministic)
}
func (m *LsLinkSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LsLinkSubscription.Merge(m, src)
}
func (m *LsLinkSubscription) XXX_Size() int {
	return xxx_messageInfo_LsLinkSubscription.Size(m)
}
func (m *LsLinkSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_LsLinkSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_LsLinkSubscription proto.InternalMessageInfo

func (m *LsLinkSubscription) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type LsLinkEvent struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	LsLink               *LsLink  `protobuf:"bytes,3,opt,name=lsLink,proto3" json:"lsLink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LsLinkEvent) Reset()         { *m = LsLinkEvent{} }
func (m *LsLinkEvent) String() string { return proto.CompactTextString(m) }
func (*LsLinkEvent) ProtoMessage()    {}
func (*LsLinkEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac404c4a13b2757, []int{5}
}

func (m *LsLinkEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LsLinkEvent.Unmarshal(m, b)
}
func (m *LsLinkEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LsLinkEvent.Marshal(b, m, deterministic)
}
func (m *LsLinkEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LsLinkEvent.Merge(m, src)
}
func (m *LsLinkEvent) XXX_Size() int {
	return xxx_messageInfo_LsLinkEvent.Size(m)
}
func (m *LsLinkEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LsLinkEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LsLinkEvent proto.InternalMessageInfo

func (m *LsLinkEvent) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *LsLinkEvent) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LsLinkEvent) GetLsLink() *LsLink {
	if m != nil {
		return m.LsLink
	}
	return nil
}

type DataRate struct {
	Ipv4Address          string   `protobuf:"bytes,1,opt,name=ipv4address,proto3" json:"ipv4address,omitempty"`
	DataRate             int64    `protobuf:"varint,2,opt,name=dataRate,proto3" json:"dataRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRate) Reset()         { *m = DataRate{} }
func (m *DataRate) String() string { return proto.CompactTextString(m) }
func (*DataRate) ProtoMessage()    {}
func (*DataRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac404c4a13b2757, []int{6}
}

func (m *DataRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRate.Unmarshal(m, b)
}
func (m *DataRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRate.Marshal(b, m, deterministic)
}
func (m *DataRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRate.Merge(m, src)
}
func (m *DataRate) XXX_Size() int {
	return xxx_messageInfo_DataRate.Size(m)
}
func (m *DataRate) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRate.DiscardUnknown(m)
}

var xxx_messageInfo_DataRate proto.InternalMessageInfo

func (m *DataRate) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *DataRate) GetDataRate() int64 {
	if m != nil {
		return m.DataRate
	}
	return 0
}

type DataRateEvent struct {
	Action               string    `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Key                  string    `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	DataRate             *DataRate `protobuf:"bytes,3,opt,name=dataRate,proto3" json:"dataRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DataRateEvent) Reset()         { *m = DataRateEvent{} }
func (m *DataRateEvent) String() string { return proto.CompactTextString(m) }
func (*DataRateEvent) ProtoMessage()    {}
func (*DataRateEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac404c4a13b2757, []int{7}
}

func (m *DataRateEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRateEvent.Unmarshal(m, b)
}
func (m *DataRateEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRateEvent.Marshal(b, m, deterministic)
}
func (m *DataRateEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRateEvent.Merge(m, src)
}
func (m *DataRateEvent) XXX_Size() int {
	return xxx_messageInfo_DataRateEvent.Size(m)
}
func (m *DataRateEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRateEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DataRateEvent proto.InternalMessageInfo

func (m *DataRateEvent) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DataRateEvent) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataRateEvent) GetDataRate() *DataRate {
	if m != nil {
		return m.DataRate
	}
	return nil
}

type DataRateSubscription struct {
	Ipv4Addresses        []string `protobuf:"bytes,1,rep,name=ipv4addresses,proto3" json:"ipv4addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRateSubscription) Reset()         { *m = DataRateSubscription{} }
func (m *DataRateSubscription) String() string { return proto.CompactTextString(m) }
func (*DataRateSubscription) ProtoMessage()    {}
func (*DataRateSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac404c4a13b2757, []int{8}
}

func (m *DataRateSubscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRateSubscription.Unmarshal(m, b)
}
func (m *DataRateSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRateSubscription.Marshal(b, m, deterministic)
}
func (m *DataRateSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRateSubscription.Merge(m, src)
}
func (m *DataRateSubscription) XXX_Size() int {
	return xxx_messageInfo_DataRateSubscription.Size(m)
}
func (m *DataRateSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRateSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_DataRateSubscription proto.InternalMessageInfo

func (m *DataRateSubscription) GetIpv4Addresses() []string {
	if m != nil {
		return m.Ipv4Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*LsNode)(nil), "pushservice.LsNode")
	proto.RegisterType((*LsNodeSubscription)(nil), "pushservice.LsNodeSubscription")
	proto.RegisterType((*LsNodeEvent)(nil), "pushservice.LsNodeEvent")
	proto.RegisterType((*LsLink)(nil), "pushservice.LsLink")
	proto.RegisterType((*LsLinkSubscription)(nil), "pushservice.LsLinkSubscription")
	proto.RegisterType((*LsLinkEvent)(nil), "pushservice.LsLinkEvent")
	proto.RegisterType((*DataRate)(nil), "pushservice.DataRate")
	proto.RegisterType((*DataRateEvent)(nil), "pushservice.DataRateEvent")
	proto.RegisterType((*DataRateSubscription)(nil), "pushservice.DataRateSubscription")
}

func init() { proto.RegisterFile("pushservice.proto", fileDescriptor_dac404c4a13b2757) }

var fileDescriptor_dac404c4a13b2757 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x61, 0x8b, 0xd3, 0x30,
	0x1c, 0xc6, 0xed, 0x76, 0x57, 0xb7, 0x7f, 0x6f, 0xa0, 0x7f, 0x4f, 0x2d, 0x13, 0x71, 0x57, 0x7d,
	0x31, 0x90, 0x5b, 0xf5, 0xf4, 0xa5, 0x22, 0x88, 0x82, 0x83, 0x53, 0xa4, 0xf3, 0x8d, 0x82, 0x1c,
	0x59, 0x1b, 0xb6, 0xb8, 0xae, 0x09, 0x4d, 0x36, 0xb8, 0xaf, 0xe1, 0x97, 0xf1, 0xeb, 0x49, 0x92,
	0x76, 0x4b, 0xbd, 0x1e, 0xa8, 0xef, 0x92, 0xa7, 0x4f, 0x7e, 0x79, 0xfe, 0xcf, 0xc2, 0xe0, 0xb6,
	0xd8, 0xc8, 0xa5, 0xa4, 0xe5, 0x96, 0xa5, 0x74, 0x22, 0x4a, 0xae, 0x38, 0x06, 0x8e, 0x14, 0x7d,
	0x07, 0xff, 0x5c, 0x7e, 0xe2, 0x19, 0xc5, 0x5b, 0xd0, 0x5d, 0xd1, 0xcb, 0xd0, 0x1b, 0x79, 0xe3,
	0x7e, 0xa2, 0x97, 0x88, 0x70, 0x50, 0x90, 0x35, 0x0d, 0x3b, 0x46, 0x32, 0x6b, 0xed, 0x22, 0xb2,
	0x08, 0xbb, 0x23, 0x6f, 0x7c, 0x98, 0xe8, 0x25, 0x3e, 0x80, 0x7e, 0xc9, 0x37, 0x8a, 0x96, 0x17,
	0x4c, 0x84, 0x07, 0xc6, 0xda, 0xb3, 0xc2, 0x54, 0x44, 0x63, 0x40, 0x8b, 0x9f, 0x6d, 0xe6, 0x32,
	0x2d, 0x99, 0x50, 0x8c, 0x17, 0x1a, 0xbc, 0xa2, 0x97, 0x32, 0xf4, 0x46, 0x5d, 0x0d, 0xd6, 0xeb,
	0x28, 0x83, 0xc0, 0x3a, 0xdf, 0x6f, 0x69, 0xa1, 0xf0, 0x1e, 0xf8, 0x24, 0xd5, 0xe6, 0x2a, 0x50,
	0xb5, 0xab, 0x53, 0x76, 0xf6, 0x29, 0x9f, 0x82, 0x9f, 0x9b, 0x83, 0x26, 0x54, 0x70, 0x76, 0x67,
	0xe2, 0x8e, 0x6c, 0x99, 0x49, 0x65, 0x89, 0x7e, 0x79, 0x7a, 0xde, 0x73, 0x56, 0xac, 0x5a, 0xe6,
	0x6d, 0x4c, 0xd2, 0x69, 0x4e, 0x82, 0xf7, 0xe1, 0xa6, 0xa0, 0xf6, 0x53, 0xd7, 0x26, 0xd2, 0xdb,
	0xa9, 0xc0, 0x13, 0x38, 0xca, 0x79, 0x4a, 0x72, 0x0d, 0xdd, 0x57, 0x10, 0xec, 0xb4, 0xa9, 0xc0,
	0xc7, 0x30, 0x28, 0xe9, 0x9a, 0x2b, 0x5a, 0x7b, 0x0e, 0x8d, 0xe7, 0x68, 0x2f, 0x4e, 0x05, 0x3e,
	0x04, 0x60, 0x0b, 0x71, 0xb1, 0xa6, 0xaa, 0x64, 0x69, 0xe8, 0x9b, 0x82, 0xfb, 0x6c, 0x21, 0x3e,
	0x1a, 0xc1, 0x36, 0xa9, 0xad, 0x7f, 0xd7, 0xa4, 0x76, 0xfe, 0x57, 0x93, 0xfa, 0xe0, 0x35, 0x4d,
	0xea, 0x4f, 0x49, 0x65, 0x89, 0x3e, 0x40, 0xef, 0x1d, 0x51, 0x24, 0x21, 0x8a, 0xe2, 0x08, 0x02,
	0x26, 0xb6, 0x2f, 0x49, 0x96, 0x95, 0x54, 0xca, 0xea, 0x1e, 0x57, 0xc2, 0x21, 0xf4, 0xb2, 0xca,
	0x6d, 0x6e, 0xec, 0x26, 0xbb, 0x7d, 0x94, 0xc3, 0xa0, 0x26, 0xfd, 0x6b, 0xe2, 0xe7, 0x0e, 0xd6,
	0x66, 0xbe, 0xdb, 0xc8, 0x5c, 0x73, 0x9d, 0xdb, 0x5e, 0xc1, 0x71, 0xad, 0x36, 0x9a, 0x7c, 0x02,
	0x03, 0x27, 0x30, 0xad, 0x2b, 0x6d, 0x8a, 0x67, 0x3f, 0x3b, 0x10, 0x7c, 0xde, 0xc8, 0xe5, 0xcc,
	0x5e, 0x80, 0x5f, 0xe1, 0xb8, 0xa2, 0xcc, 0xe9, 0x17, 0x5e, 0x83, 0x25, 0x9e, 0xb4, 0xc6, 0x70,
	0x2f, 0x1c, 0x0e, 0x5b, 0x2d, 0xa6, 0x81, 0xe8, 0xc6, 0x33, 0x0f, 0x67, 0x80, 0x0e, 0xda, 0xbe,
	0x63, 0x89, 0x8f, 0x5a, 0x5e, 0x77, 0x03, 0x1b, 0xb6, 0x18, 0xae, 0x87, 0xea, 0x9f, 0xf2, 0x2a,
	0xf4, 0xcf, 0x67, 0x76, 0x05, 0xba, 0x7b, 0x5d, 0x1a, 0xfa, 0xf6, 0xcd, 0xb7, 0xd7, 0x0b, 0xa6,
	0x72, 0x32, 0x9f, 0x70, 0xa9, 0x26, 0xe9, 0x32, 0x66, 0x85, 0x8c, 0x7f, 0x90, 0x9c, 0x08, 0x5a,
	0xf0, 0x53, 0x22, 0x58, 0xac, 0x4f, 0x9f, 0x56, 0xc7, 0x63, 0xf3, 0x2f, 0x14, 0x3b, 0xc0, 0xb9,
	0x6f, 0xa4, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xb4, 0xe2, 0x48, 0xad, 0x04, 0x00,
	0x00,
}
