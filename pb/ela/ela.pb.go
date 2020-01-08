// SPDX-License-Identifier: Apache-2.0
// Copyright © 2019 Intel Corporation

package ela

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type TrafficTarget_TargetAction int32

const (
	TrafficTarget_ACCEPT TrafficTarget_TargetAction = 0
	TrafficTarget_REJECT TrafficTarget_TargetAction = 1
	TrafficTarget_DROP   TrafficTarget_TargetAction = 2
)

var TrafficTarget_TargetAction_name = map[int32]string{
	0: "ACCEPT",
	1: "REJECT",
	2: "DROP",
}

var TrafficTarget_TargetAction_value = map[string]int32{
	"ACCEPT": 0,
	"REJECT": 1,
	"DROP":   2,
}

func (x TrafficTarget_TargetAction) String() string {
	return proto.EnumName(TrafficTarget_TargetAction_name, int32(x))
}

func (TrafficTarget_TargetAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{6, 0}
}

type NetworkInterface_InterfaceDriver int32

const (
	NetworkInterface_KERNEL    NetworkInterface_InterfaceDriver = 0
	NetworkInterface_USERSPACE NetworkInterface_InterfaceDriver = 1
)

var NetworkInterface_InterfaceDriver_name = map[int32]string{
	0: "KERNEL",
	1: "USERSPACE",
}

var NetworkInterface_InterfaceDriver_value = map[string]int32{
	"KERNEL":    0,
	"USERSPACE": 1,
}

func (x NetworkInterface_InterfaceDriver) String() string {
	return proto.EnumName(NetworkInterface_InterfaceDriver_name, int32(x))
}

func (NetworkInterface_InterfaceDriver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{9, 0}
}

type NetworkInterface_InterfaceType int32

const (
	NetworkInterface_NONE          NetworkInterface_InterfaceType = 0
	NetworkInterface_UPSTREAM      NetworkInterface_InterfaceType = 1
	NetworkInterface_DOWNSTREAM    NetworkInterface_InterfaceType = 2
	NetworkInterface_BIDIRECTIONAL NetworkInterface_InterfaceType = 3
	NetworkInterface_BREAKOUT      NetworkInterface_InterfaceType = 4
)

var NetworkInterface_InterfaceType_name = map[int32]string{
	0: "NONE",
	1: "UPSTREAM",
	2: "DOWNSTREAM",
	3: "BIDIRECTIONAL",
	4: "BREAKOUT",
}

var NetworkInterface_InterfaceType_value = map[string]int32{
	"NONE":          0,
	"UPSTREAM":      1,
	"DOWNSTREAM":    2,
	"BIDIRECTIONAL": 3,
	"BREAKOUT":      4,
}

func (x NetworkInterface_InterfaceType) String() string {
	return proto.EnumName(NetworkInterface_InterfaceType_name, int32(x))
}

func (NetworkInterface_InterfaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{9, 1}
}

type NetworkSetting_Status int32

const (
	NetworkSetting_NONE   NetworkSetting_Status = 0
	NetworkSetting_STATIC NetworkSetting_Status = 1
	NetworkSetting_DHCPv4 NetworkSetting_Status = 2
	NetworkSetting_DHCPv6 NetworkSetting_Status = 3
	NetworkSetting_SLAAC  NetworkSetting_Status = 4
)

var NetworkSetting_Status_name = map[int32]string{
	0: "NONE",
	1: "STATIC",
	2: "DHCPv4",
	3: "DHCPv6",
	4: "SLAAC",
}

var NetworkSetting_Status_value = map[string]int32{
	"NONE":   0,
	"STATIC": 1,
	"DHCPv4": 2,
	"DHCPv6": 3,
	"SLAAC":  4,
}

func (x NetworkSetting_Status) String() string {
	return proto.EnumName(NetworkSetting_Status_name, int32(x))
}

func (NetworkSetting_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{13, 0}
}

// TrafficPolicy is a policy that defines a set of traffic rules for the
// identified component (i.e. an application, an interface, etc.).
//
// A policy engine applies these rules, using the context of the identified
// component in order to send traffic to a target. The policy engine acts as
// a man-in-the-middle. It may modify the packets in order to facilitate the
// traffic flow. Examples of a policy engine are DPDK, VPP or iptables
// applications.
type TrafficPolicy struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TrafficRules         []*TrafficRule `protobuf:"bytes,2,rep,name=traffic_rules,json=trafficRules,proto3" json:"traffic_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TrafficPolicy) Reset()         { *m = TrafficPolicy{} }
func (m *TrafficPolicy) String() string { return proto.CompactTextString(m) }
func (*TrafficPolicy) ProtoMessage()    {}
func (*TrafficPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{0}
}

func (m *TrafficPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficPolicy.Unmarshal(m, b)
}
func (m *TrafficPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficPolicy.Marshal(b, m, deterministic)
}
func (m *TrafficPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficPolicy.Merge(m, src)
}
func (m *TrafficPolicy) XXX_Size() int {
	return xxx_messageInfo_TrafficPolicy.Size(m)
}
func (m *TrafficPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficPolicy proto.InternalMessageInfo

func (m *TrafficPolicy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TrafficPolicy) GetTrafficRules() []*TrafficRule {
	if m != nil {
		return m.TrafficRules
	}
	return nil
}

// TrafficRule defines a single traffic rule. The traffic selectors are used in
// order to construct both a rule that must be matched as well as what action
// to take on the traffic if the rule is matched.
//
// Since this is generic, the receiver of this rule must validate if the
// information provided by the caller is sufficient enough to construct a
// policy of a particular type.
//
// A single rule only allows one of each traffic selector to be specified.
// However, if a system supports advanced networking rules, a traffic rule can
// specify a subnet mask or a range in order to create a more dynamic rule.
//
// For example, a rule with a source selector of 10.20.30.0/24 could match all
// source traffic in that subnet block.
type TrafficRule struct {
	Description          string           `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Priority             uint32           `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Source               *TrafficSelector `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Destination          *TrafficSelector `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Target               *TrafficTarget   `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TrafficRule) Reset()         { *m = TrafficRule{} }
func (m *TrafficRule) String() string { return proto.CompactTextString(m) }
func (*TrafficRule) ProtoMessage()    {}
func (*TrafficRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{1}
}

func (m *TrafficRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficRule.Unmarshal(m, b)
}
func (m *TrafficRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficRule.Marshal(b, m, deterministic)
}
func (m *TrafficRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficRule.Merge(m, src)
}
func (m *TrafficRule) XXX_Size() int {
	return xxx_messageInfo_TrafficRule.Size(m)
}
func (m *TrafficRule) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficRule.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficRule proto.InternalMessageInfo

func (m *TrafficRule) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TrafficRule) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TrafficRule) GetSource() *TrafficSelector {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *TrafficRule) GetDestination() *TrafficSelector {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *TrafficRule) GetTarget() *TrafficTarget {
	if m != nil {
		return m.Target
	}
	return nil
}

// TrafficSelector defines the parameters for a traffic selector in a
// TrafficRule. If a filter is empty, the selector does not evaluate it. The
// receiver can select traffic by using the filters as it is examining a packet
// or payload. They must filter using the OSI stack from layer 7 to layer 1.
// For example, if a MAC and IP are provided, the selector must first evaluate
// the IP (layer 3) before the MAC (layer 2).
//
// If a TrafficSelector has only the MAC filter specified, the selector is
// created only for that filter. However, if the selector contains a GTP and IP
// filter, the selector is created on both and the traffic must match both
// filters.
type TrafficSelector struct {
	Description          string     `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Macs                 *MACFilter `protobuf:"bytes,2,opt,name=macs,proto3" json:"macs,omitempty"`
	Ip                   *IPFilter  `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Gtp                  *GTPFilter `protobuf:"bytes,4,opt,name=gtp,proto3" json:"gtp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TrafficSelector) Reset()         { *m = TrafficSelector{} }
func (m *TrafficSelector) String() string { return proto.CompactTextString(m) }
func (*TrafficSelector) ProtoMessage()    {}
func (*TrafficSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{2}
}

func (m *TrafficSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSelector.Unmarshal(m, b)
}
func (m *TrafficSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSelector.Marshal(b, m, deterministic)
}
func (m *TrafficSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSelector.Merge(m, src)
}
func (m *TrafficSelector) XXX_Size() int {
	return xxx_messageInfo_TrafficSelector.Size(m)
}
func (m *TrafficSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSelector.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSelector proto.InternalMessageInfo

func (m *TrafficSelector) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TrafficSelector) GetMacs() *MACFilter {
	if m != nil {
		return m.Macs
	}
	return nil
}

func (m *TrafficSelector) GetIp() *IPFilter {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *TrafficSelector) GetGtp() *GTPFilter {
	if m != nil {
		return m.Gtp
	}
	return nil
}

// MACFilter specifies properties related to MAC filters. Some implementations
// may not support multiple MAC addresses.
type MACFilter struct {
	MacAddresses         []string `protobuf:"bytes,1,rep,name=mac_addresses,json=macAddresses,proto3" json:"mac_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MACFilter) Reset()         { *m = MACFilter{} }
func (m *MACFilter) String() string { return proto.CompactTextString(m) }
func (*MACFilter) ProtoMessage()    {}
func (*MACFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{3}
}

func (m *MACFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MACFilter.Unmarshal(m, b)
}
func (m *MACFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MACFilter.Marshal(b, m, deterministic)
}
func (m *MACFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MACFilter.Merge(m, src)
}
func (m *MACFilter) XXX_Size() int {
	return xxx_messageInfo_MACFilter.Size(m)
}
func (m *MACFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_MACFilter.DiscardUnknown(m)
}

var xxx_messageInfo_MACFilter proto.InternalMessageInfo

func (m *MACFilter) GetMacAddresses() []string {
	if m != nil {
		return m.MacAddresses
	}
	return nil
}

// IPFilter specifies properties related to IP filters. Some implementations
// may not support multiple IP address (subnets) or have IPv6 support.
//
// If a caller wishes to define a single port, begin_port and end_port should
// be the same. For example, if the port is 3306, begin_port is 3306 and
// end_port is 3306. It is invalid to provide a begin_port that is greater than
// the end_port.
//
// Leaving the address and mask fields empty implies that all possible IP
// addresses are in the filter. Leaving these primitive datatypes empty
// defaults to the type's zero-value (as is the norm in protobuf). The
// following describes the behavior depending on how the fields are populated:
//  ___________________________________________________________________________
// |      Address      |      Mask      |               Result                |
// |   Zero-value ("") | Zero-value (0) |  All IPv4 (and IPv6, if supported)  |
// |    "0.0.0.0"      | Zero-value (0) |  All IPv4 only                      |
// |       "::"        | Zero-value (0) |  All IPv6 only (if supported)       |
// |    "1.2.3.4"      | Zero-value (0) |  Invalid                            |
// |   Zero-value ("") |       24       |  Invalid                            |
//  ___________________________________________________________________________
type IPFilter struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Mask                 uint32   `protobuf:"varint,2,opt,name=mask,proto3" json:"mask,omitempty"`
	BeginPort            uint32   `protobuf:"varint,3,opt,name=begin_port,json=beginPort,proto3" json:"begin_port,omitempty"`
	EndPort              uint32   `protobuf:"varint,4,opt,name=end_port,json=endPort,proto3" json:"end_port,omitempty"`
	Protocol             string   `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPFilter) Reset()         { *m = IPFilter{} }
func (m *IPFilter) String() string { return proto.CompactTextString(m) }
func (*IPFilter) ProtoMessage()    {}
func (*IPFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{4}
}

func (m *IPFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPFilter.Unmarshal(m, b)
}
func (m *IPFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPFilter.Marshal(b, m, deterministic)
}
func (m *IPFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPFilter.Merge(m, src)
}
func (m *IPFilter) XXX_Size() int {
	return xxx_messageInfo_IPFilter.Size(m)
}
func (m *IPFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_IPFilter.DiscardUnknown(m)
}

var xxx_messageInfo_IPFilter proto.InternalMessageInfo

func (m *IPFilter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPFilter) GetMask() uint32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *IPFilter) GetBeginPort() uint32 {
	if m != nil {
		return m.BeginPort
	}
	return 0
}

func (m *IPFilter) GetEndPort() uint32 {
	if m != nil {
		return m.EndPort
	}
	return 0
}

func (m *IPFilter) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

// GTPFilter specifies properties related to GTP filters. Some implementations
// may not support multiple addresses or multiple IMSIs.
type GTPFilter struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Mask                 uint32   `protobuf:"varint,2,opt,name=mask,proto3" json:"mask,omitempty"`
	Imsis                []string `protobuf:"bytes,3,rep,name=imsis,proto3" json:"imsis,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GTPFilter) Reset()         { *m = GTPFilter{} }
func (m *GTPFilter) String() string { return proto.CompactTextString(m) }
func (*GTPFilter) ProtoMessage()    {}
func (*GTPFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{5}
}

func (m *GTPFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GTPFilter.Unmarshal(m, b)
}
func (m *GTPFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GTPFilter.Marshal(b, m, deterministic)
}
func (m *GTPFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GTPFilter.Merge(m, src)
}
func (m *GTPFilter) XXX_Size() int {
	return xxx_messageInfo_GTPFilter.Size(m)
}
func (m *GTPFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_GTPFilter.DiscardUnknown(m)
}

var xxx_messageInfo_GTPFilter proto.InternalMessageInfo

func (m *GTPFilter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GTPFilter) GetMask() uint32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *GTPFilter) GetImsis() []string {
	if m != nil {
		return m.Imsis
	}
	return nil
}

// TrafficTarget defines the parameters for a traffic target in a TrafficRule.
// The action indicates what target action to perform. If a modify field is
// empty, the target does not perform that type of modification.
//
// For example, if the target should modify the MAC address, then it should be
// provided in the message. The modifiers are currently only applicable if the
// interface is trying to modify the traffic, such as is the case with a
// breakout interface.
type TrafficTarget struct {
	Description          string                     `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Action               TrafficTarget_TargetAction `protobuf:"varint,2,opt,name=action,proto3,enum=openness.ela.TrafficTarget_TargetAction" json:"action,omitempty"`
	Mac                  *MACModifier               `protobuf:"bytes,3,opt,name=mac,proto3" json:"mac,omitempty"`
	Ip                   *IPModifier                `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TrafficTarget) Reset()         { *m = TrafficTarget{} }
func (m *TrafficTarget) String() string { return proto.CompactTextString(m) }
func (*TrafficTarget) ProtoMessage()    {}
func (*TrafficTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{6}
}

func (m *TrafficTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTarget.Unmarshal(m, b)
}
func (m *TrafficTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTarget.Marshal(b, m, deterministic)
}
func (m *TrafficTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTarget.Merge(m, src)
}
func (m *TrafficTarget) XXX_Size() int {
	return xxx_messageInfo_TrafficTarget.Size(m)
}
func (m *TrafficTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTarget.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTarget proto.InternalMessageInfo

func (m *TrafficTarget) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TrafficTarget) GetAction() TrafficTarget_TargetAction {
	if m != nil {
		return m.Action
	}
	return TrafficTarget_ACCEPT
}

func (m *TrafficTarget) GetMac() *MACModifier {
	if m != nil {
		return m.Mac
	}
	return nil
}

func (m *TrafficTarget) GetIp() *IPModifier {
	if m != nil {
		return m.Ip
	}
	return nil
}

// MACModifier defines the MAC properties that should be modified.
type MACModifier struct {
	MacAddress           string   `protobuf:"bytes,1,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MACModifier) Reset()         { *m = MACModifier{} }
func (m *MACModifier) String() string { return proto.CompactTextString(m) }
func (*MACModifier) ProtoMessage()    {}
func (*MACModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{7}
}

func (m *MACModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MACModifier.Unmarshal(m, b)
}
func (m *MACModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MACModifier.Marshal(b, m, deterministic)
}
func (m *MACModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MACModifier.Merge(m, src)
}
func (m *MACModifier) XXX_Size() int {
	return xxx_messageInfo_MACModifier.Size(m)
}
func (m *MACModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MACModifier.DiscardUnknown(m)
}

var xxx_messageInfo_MACModifier proto.InternalMessageInfo

func (m *MACModifier) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

// IPModifier defines the IP properties that should be modified
type IPModifier struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPModifier) Reset()         { *m = IPModifier{} }
func (m *IPModifier) String() string { return proto.CompactTextString(m) }
func (*IPModifier) ProtoMessage()    {}
func (*IPModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{8}
}

func (m *IPModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPModifier.Unmarshal(m, b)
}
func (m *IPModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPModifier.Marshal(b, m, deterministic)
}
func (m *IPModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPModifier.Merge(m, src)
}
func (m *IPModifier) XXX_Size() int {
	return xxx_messageInfo_IPModifier.Size(m)
}
func (m *IPModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_IPModifier.DiscardUnknown(m)
}

var xxx_messageInfo_IPModifier proto.InternalMessageInfo

func (m *IPModifier) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPModifier) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// NetworkInterface defines a network interface available on the host.
// Interfaces are typically kernel interfaces by default, and can be changed if
// the caller wishes to do so.
//
// The interface's type assists the policy engine in determining what types of
// traffic the interface can expect to be handling, and is mainly here for
// support of legacy implementations (which may require the field is updated in
// order to work properly).
//
// An interface can belong to multiple zones, which can be useful for when
// the amount of actual interfaces is limited.
type NetworkInterface struct {
	Id          string                           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string                           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Driver      NetworkInterface_InterfaceDriver `protobuf:"varint,3,opt,name=driver,proto3,enum=openness.ela.NetworkInterface_InterfaceDriver" json:"driver,omitempty"`
	Type        NetworkInterface_InterfaceType   `protobuf:"varint,4,opt,name=type,proto3,enum=openness.ela.NetworkInterface_InterfaceType" json:"type,omitempty"`
	MacAddress  string                           `protobuf:"bytes,5,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Vlan        uint32                           `protobuf:"varint,6,opt,name=vlan,proto3" json:"vlan,omitempty"`
	Zones       []string                         `protobuf:"bytes,7,rep,name=zones,proto3" json:"zones,omitempty"`
	// (LEGACY) The fallback interface for this interface. This only exists for
	// legacy dataplane implementations. In future implementations, a traffic
	// policy should be used to yield the same results. Using this is not
	// advisable as it belongs in the traffic policy and exposes a fallback
	// behavior that can be seen as insecure.
	FallbackInterface    string   `protobuf:"bytes,8,opt,name=fallback_interface,json=fallbackInterface,proto3" json:"fallback_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInterface) Reset()         { *m = NetworkInterface{} }
func (m *NetworkInterface) String() string { return proto.CompactTextString(m) }
func (*NetworkInterface) ProtoMessage()    {}
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{9}
}

func (m *NetworkInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterface.Unmarshal(m, b)
}
func (m *NetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterface.Marshal(b, m, deterministic)
}
func (m *NetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterface.Merge(m, src)
}
func (m *NetworkInterface) XXX_Size() int {
	return xxx_messageInfo_NetworkInterface.Size(m)
}
func (m *NetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterface proto.InternalMessageInfo

func (m *NetworkInterface) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkInterface) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NetworkInterface) GetDriver() NetworkInterface_InterfaceDriver {
	if m != nil {
		return m.Driver
	}
	return NetworkInterface_KERNEL
}

func (m *NetworkInterface) GetType() NetworkInterface_InterfaceType {
	if m != nil {
		return m.Type
	}
	return NetworkInterface_NONE
}

func (m *NetworkInterface) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *NetworkInterface) GetVlan() uint32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *NetworkInterface) GetZones() []string {
	if m != nil {
		return m.Zones
	}
	return nil
}

func (m *NetworkInterface) GetFallbackInterface() string {
	if m != nil {
		return m.FallbackInterface
	}
	return ""
}

type NetworkInterfaces struct {
	NetworkInterfaces    []*NetworkInterface `protobuf:"bytes,1,rep,name=network_interfaces,json=networkInterfaces,proto3" json:"network_interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NetworkInterfaces) Reset()         { *m = NetworkInterfaces{} }
func (m *NetworkInterfaces) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaces) ProtoMessage()    {}
func (*NetworkInterfaces) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{10}
}

func (m *NetworkInterfaces) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterfaces.Unmarshal(m, b)
}
func (m *NetworkInterfaces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterfaces.Marshal(b, m, deterministic)
}
func (m *NetworkInterfaces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterfaces.Merge(m, src)
}
func (m *NetworkInterfaces) XXX_Size() int {
	return xxx_messageInfo_NetworkInterfaces.Size(m)
}
func (m *NetworkInterfaces) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterfaces.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterfaces proto.InternalMessageInfo

func (m *NetworkInterfaces) GetNetworkInterfaces() []*NetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

// NetworkZone defines a network zone. A zone is effectively a label that
// isolates network traffic within an appliance. It allows for further rules
// to be made surrounding the zone and interfaces that are assigned to it.
type NetworkZone struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkZone) Reset()         { *m = NetworkZone{} }
func (m *NetworkZone) String() string { return proto.CompactTextString(m) }
func (*NetworkZone) ProtoMessage()    {}
func (*NetworkZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{11}
}

func (m *NetworkZone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkZone.Unmarshal(m, b)
}
func (m *NetworkZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkZone.Marshal(b, m, deterministic)
}
func (m *NetworkZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkZone.Merge(m, src)
}
func (m *NetworkZone) XXX_Size() int {
	return xxx_messageInfo_NetworkZone.Size(m)
}
func (m *NetworkZone) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkZone.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkZone proto.InternalMessageInfo

func (m *NetworkZone) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkZone) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type NetworkZones struct {
	NetworkZones         []*NetworkZone `protobuf:"bytes,1,rep,name=network_zones,json=networkZones,proto3" json:"network_zones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NetworkZones) Reset()         { *m = NetworkZones{} }
func (m *NetworkZones) String() string { return proto.CompactTextString(m) }
func (*NetworkZones) ProtoMessage()    {}
func (*NetworkZones) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{12}
}

func (m *NetworkZones) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkZones.Unmarshal(m, b)
}
func (m *NetworkZones) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkZones.Marshal(b, m, deterministic)
}
func (m *NetworkZones) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkZones.Merge(m, src)
}
func (m *NetworkZones) XXX_Size() int {
	return xxx_messageInfo_NetworkZones.Size(m)
}
func (m *NetworkZones) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkZones.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkZones proto.InternalMessageInfo

func (m *NetworkZones) GetNetworkZones() []*NetworkZone {
	if m != nil {
		return m.NetworkZones
	}
	return nil
}

// NetworkSetting defines a network setting. It can be included in an interface
// to configure it's IP properties.
type NetworkSetting struct {
	Status               NetworkSetting_Status `protobuf:"varint,1,opt,name=status,proto3,enum=openness.ela.NetworkSetting_Status" json:"status,omitempty"`
	Address              string                `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Mask                 uint32                `protobuf:"varint,3,opt,name=mask,proto3" json:"mask,omitempty"`
	Gateway              string                `protobuf:"bytes,4,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Dns                  []string              `protobuf:"bytes,5,rep,name=dns,proto3" json:"dns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NetworkSetting) Reset()         { *m = NetworkSetting{} }
func (m *NetworkSetting) String() string { return proto.CompactTextString(m) }
func (*NetworkSetting) ProtoMessage()    {}
func (*NetworkSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{13}
}

func (m *NetworkSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkSetting.Unmarshal(m, b)
}
func (m *NetworkSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkSetting.Marshal(b, m, deterministic)
}
func (m *NetworkSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkSetting.Merge(m, src)
}
func (m *NetworkSetting) XXX_Size() int {
	return xxx_messageInfo_NetworkSetting.Size(m)
}
func (m *NetworkSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkSetting.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkSetting proto.InternalMessageInfo

func (m *NetworkSetting) GetStatus() NetworkSetting_Status {
	if m != nil {
		return m.Status
	}
	return NetworkSetting_NONE
}

func (m *NetworkSetting) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NetworkSetting) GetMask() uint32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *NetworkSetting) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *NetworkSetting) GetDns() []string {
	if m != nil {
		return m.Dns
	}
	return nil
}

// DNSForwarders represents the upstream DNS forwarders, which may be used when
// the DNS services is performing a recursive lookup. Forwarders should be
// utilized when more advanced DNS usage is desired.
type DNSForwarders struct {
	IpAddresses          []string `protobuf:"bytes,1,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNSForwarders) Reset()         { *m = DNSForwarders{} }
func (m *DNSForwarders) String() string { return proto.CompactTextString(m) }
func (*DNSForwarders) ProtoMessage()    {}
func (*DNSForwarders) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{14}
}

func (m *DNSForwarders) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNSForwarders.Unmarshal(m, b)
}
func (m *DNSForwarders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNSForwarders.Marshal(b, m, deterministic)
}
func (m *DNSForwarders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNSForwarders.Merge(m, src)
}
func (m *DNSForwarders) XXX_Size() int {
	return xxx_messageInfo_DNSForwarders.Size(m)
}
func (m *DNSForwarders) XXX_DiscardUnknown() {
	xxx_messageInfo_DNSForwarders.DiscardUnknown(m)
}

var xxx_messageInfo_DNSForwarders proto.InternalMessageInfo

func (m *DNSForwarders) GetIpAddresses() []string {
	if m != nil {
		return m.IpAddresses
	}
	return nil
}

// DNSARecordSet contains one or more values for a name, which is a fully
// qualified domain name (FQDN). The values are typically either an ID for
// the record (such as an application ID or a VNF ID) or an IP address.
type DNSARecordSet struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNSARecordSet) Reset()         { *m = DNSARecordSet{} }
func (m *DNSARecordSet) String() string { return proto.CompactTextString(m) }
func (*DNSARecordSet) ProtoMessage()    {}
func (*DNSARecordSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{15}
}

func (m *DNSARecordSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNSARecordSet.Unmarshal(m, b)
}
func (m *DNSARecordSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNSARecordSet.Marshal(b, m, deterministic)
}
func (m *DNSARecordSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNSARecordSet.Merge(m, src)
}
func (m *DNSARecordSet) XXX_Size() int {
	return xxx_messageInfo_DNSARecordSet.Size(m)
}
func (m *DNSARecordSet) XXX_DiscardUnknown() {
	xxx_messageInfo_DNSARecordSet.DiscardUnknown(m)
}

var xxx_messageInfo_DNSARecordSet proto.InternalMessageInfo

func (m *DNSARecordSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DNSARecordSet) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type InterfaceID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceID) Reset()         { *m = InterfaceID{} }
func (m *InterfaceID) String() string { return proto.CompactTextString(m) }
func (*InterfaceID) ProtoMessage()    {}
func (*InterfaceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{16}
}

func (m *InterfaceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceID.Unmarshal(m, b)
}
func (m *InterfaceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceID.Marshal(b, m, deterministic)
}
func (m *InterfaceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceID.Merge(m, src)
}
func (m *InterfaceID) XXX_Size() int {
	return xxx_messageInfo_InterfaceID.Size(m)
}
func (m *InterfaceID) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceID.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceID proto.InternalMessageInfo

func (m *InterfaceID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ZoneID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZoneID) Reset()         { *m = ZoneID{} }
func (m *ZoneID) String() string { return proto.CompactTextString(m) }
func (*ZoneID) ProtoMessage()    {}
func (*ZoneID) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb26205266db6e19, []int{17}
}

func (m *ZoneID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZoneID.Unmarshal(m, b)
}
func (m *ZoneID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZoneID.Marshal(b, m, deterministic)
}
func (m *ZoneID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZoneID.Merge(m, src)
}
func (m *ZoneID) XXX_Size() int {
	return xxx_messageInfo_ZoneID.Size(m)
}
func (m *ZoneID) XXX_DiscardUnknown() {
	xxx_messageInfo_ZoneID.DiscardUnknown(m)
}

var xxx_messageInfo_ZoneID proto.InternalMessageInfo

func (m *ZoneID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("openness.ela.TrafficTarget_TargetAction", TrafficTarget_TargetAction_name, TrafficTarget_TargetAction_value)
	proto.RegisterEnum("openness.ela.NetworkInterface_InterfaceDriver", NetworkInterface_InterfaceDriver_name, NetworkInterface_InterfaceDriver_value)
	proto.RegisterEnum("openness.ela.NetworkInterface_InterfaceType", NetworkInterface_InterfaceType_name, NetworkInterface_InterfaceType_value)
	proto.RegisterEnum("openness.ela.NetworkSetting_Status", NetworkSetting_Status_name, NetworkSetting_Status_value)
	proto.RegisterType((*TrafficPolicy)(nil), "openness.ela.TrafficPolicy")
	proto.RegisterType((*TrafficRule)(nil), "openness.ela.TrafficRule")
	proto.RegisterType((*TrafficSelector)(nil), "openness.ela.TrafficSelector")
	proto.RegisterType((*MACFilter)(nil), "openness.ela.MACFilter")
	proto.RegisterType((*IPFilter)(nil), "openness.ela.IPFilter")
	proto.RegisterType((*GTPFilter)(nil), "openness.ela.GTPFilter")
	proto.RegisterType((*TrafficTarget)(nil), "openness.ela.TrafficTarget")
	proto.RegisterType((*MACModifier)(nil), "openness.ela.MACModifier")
	proto.RegisterType((*IPModifier)(nil), "openness.ela.IPModifier")
	proto.RegisterType((*NetworkInterface)(nil), "openness.ela.NetworkInterface")
	proto.RegisterType((*NetworkInterfaces)(nil), "openness.ela.NetworkInterfaces")
	proto.RegisterType((*NetworkZone)(nil), "openness.ela.NetworkZone")
	proto.RegisterType((*NetworkZones)(nil), "openness.ela.NetworkZones")
	proto.RegisterType((*NetworkSetting)(nil), "openness.ela.NetworkSetting")
	proto.RegisterType((*DNSForwarders)(nil), "openness.ela.DNSForwarders")
	proto.RegisterType((*DNSARecordSet)(nil), "openness.ela.DNSARecordSet")
	proto.RegisterType((*InterfaceID)(nil), "openness.ela.InterfaceID")
	proto.RegisterType((*ZoneID)(nil), "openness.ela.ZoneID")
}

func init() { proto.RegisterFile("ela.proto", fileDescriptor_eb26205266db6e19) }

var fileDescriptor_eb26205266db6e19 = []byte{
	// 1281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x17, 0x15, 0x49, 0x99, 0xb6, 0xae, 0x7e, 0x42, 0x0f, 0x02, 0x87, 0x56, 0x90, 0x2f, 0xfe, 0x18,
	0xa0, 0x50, 0x91, 0x84, 0x29, 0x94, 0xb6, 0x28, 0x9a, 0xe6, 0x87, 0xfa, 0xb1, 0xab, 0x26, 0x96,
	0x85, 0xa1, 0xdc, 0x06, 0xd9, 0x18, 0x34, 0x39, 0x16, 0x06, 0xa6, 0x48, 0x82, 0x1c, 0x39, 0x70,
	0xb7, 0xdd, 0xf7, 0x19, 0xba, 0xec, 0xbe, 0x2f, 0xd2, 0xe7, 0xe8, 0x13, 0x74, 0xd1, 0x45, 0xc1,
	0x21, 0x45, 0x51, 0xb4, 0x2d, 0xbb, 0x6e, 0x57, 0x9a, 0xd1, 0x9c, 0x7b, 0xee, 0xdc, 0x73, 0xee,
	0x0c, 0x07, 0x2a, 0xc4, 0xb5, 0xf4, 0x20, 0xf4, 0x99, 0x8f, 0x6a, 0x7e, 0x40, 0x3c, 0x8f, 0x44,
	0x91, 0x4e, 0x5c, 0xab, 0x79, 0x7f, 0xe2, 0xfb, 0x13, 0x97, 0x3c, 0xe3, 0x6b, 0xc7, 0xb3, 0x93,
	0x67, 0x64, 0x1a, 0xb0, 0xf3, 0x04, 0xaa, 0x1d, 0x41, 0x7d, 0x1c, 0x5a, 0x27, 0x27, 0xd4, 0x1e,
	0xf9, 0x2e, 0xb5, 0xcf, 0x51, 0x03, 0x44, 0xea, 0xa8, 0xc2, 0x8e, 0xd0, 0xaa, 0x60, 0x91, 0x3a,
	0xe8, 0x15, 0xd4, 0x59, 0x02, 0x38, 0x0a, 0x67, 0x2e, 0x89, 0x54, 0x71, 0x47, 0x6a, 0x55, 0xdb,
	0xdb, 0x7a, 0x3e, 0x87, 0x9e, 0x72, 0xe0, 0x99, 0x4b, 0x70, 0x8d, 0x2d, 0x26, 0x91, 0xf6, 0xa7,
	0x00, 0xd5, 0xdc, 0x2a, 0xda, 0x81, 0xaa, 0x43, 0x22, 0x3b, 0xa4, 0x01, 0xa3, 0xbe, 0x97, 0x26,
	0xca, 0xff, 0x85, 0x9a, 0xb0, 0x11, 0x84, 0xd4, 0x0f, 0x29, 0x3b, 0x57, 0xc5, 0x1d, 0xa1, 0x55,
	0xc7, 0xd9, 0x1c, 0x7d, 0x01, 0x72, 0xe4, 0xcf, 0x42, 0x9b, 0xa8, 0xd2, 0x8e, 0xd0, 0xaa, 0xb6,
	0x1f, 0x5c, 0xba, 0x0d, 0x93, 0xb8, 0xc4, 0x66, 0x7e, 0x88, 0x53, 0x30, 0x7a, 0xcd, 0x93, 0x32,
	0xea, 0x59, 0x3c, 0x69, 0xf9, 0x26, 0xb1, 0xf9, 0x08, 0xf4, 0x1c, 0x64, 0x66, 0x85, 0x13, 0xc2,
	0xd4, 0x35, 0x1e, 0x7b, 0xff, 0xd2, 0xd8, 0x31, 0x87, 0xe0, 0x14, 0xaa, 0xfd, 0x26, 0xc0, 0x9d,
	0x02, 0xeb, 0x0d, 0xca, 0x7f, 0x0c, 0xe5, 0xa9, 0x65, 0x47, 0xbc, 0xf4, 0x6a, 0xfb, 0xde, 0x72,
	0xa2, 0x7d, 0xa3, 0xbb, 0x4b, 0x5d, 0x46, 0x42, 0xcc, 0x41, 0xe8, 0x13, 0x10, 0x69, 0x90, 0x6a,
	0xb1, 0xb5, 0x0c, 0x1d, 0x8c, 0x52, 0xa4, 0x48, 0x03, 0xf4, 0x29, 0x48, 0x13, 0x16, 0xa4, 0x85,
	0x17, 0x38, 0xf7, 0xc6, 0x73, 0x64, 0x8c, 0xd1, 0x3e, 0x83, 0x4a, 0x96, 0x05, 0x3d, 0x82, 0xfa,
	0xd4, 0xb2, 0x8f, 0x2c, 0xc7, 0x09, 0x49, 0x14, 0x91, 0x48, 0x15, 0x76, 0xa4, 0x56, 0x05, 0xd7,
	0xa6, 0x96, 0x6d, 0xcc, 0xff, 0xd3, 0x7e, 0x16, 0x60, 0x63, 0x9e, 0x0d, 0xa9, 0xb0, 0x9e, 0xa2,
	0xd3, 0xe2, 0xe6, 0x53, 0x84, 0xe2, 0xc2, 0xa2, 0xd3, 0xd4, 0x53, 0x3e, 0x46, 0x0f, 0x00, 0x8e,
	0xc9, 0x84, 0x7a, 0x47, 0x81, 0x1f, 0x32, 0x5e, 0x47, 0x1d, 0x57, 0xf8, 0x3f, 0x23, 0x3f, 0x64,
	0x68, 0x1b, 0x36, 0x88, 0xe7, 0x24, 0x8b, 0x65, 0xbe, 0xb8, 0x4e, 0x3c, 0x87, 0x2f, 0xf1, 0x2e,
	0xf1, 0x99, 0x6f, 0xfb, 0x2e, 0xf7, 0xa4, 0x82, 0xb3, 0xb9, 0x76, 0x00, 0x95, 0xac, 0xa8, 0x7f,
	0xb8, 0xa1, 0xbb, 0xb0, 0x46, 0xa7, 0x11, 0x8d, 0x54, 0x89, 0x17, 0x9a, 0x4c, 0xb4, 0xbf, 0x84,
	0xec, 0x98, 0x24, 0x1e, 0xdf, 0xc0, 0xc7, 0x37, 0x20, 0x5b, 0x36, 0x5f, 0x8c, 0xf9, 0x1b, 0xed,
	0xd6, 0x8a, 0x96, 0xd1, 0x93, 0x1f, 0x83, 0xe3, 0x71, 0x1a, 0x87, 0x1e, 0x83, 0x34, 0xb5, 0xec,
	0xd4, 0xdd, 0xed, 0x0b, 0x8d, 0xb0, 0xef, 0x3b, 0xf4, 0x84, 0xc6, 0xb6, 0x4d, 0x2d, 0x1b, 0xb5,
	0x78, 0x27, 0x24, 0x06, 0xab, 0xc5, 0x4e, 0xc8, 0xa0, 0x22, 0x8d, 0x0d, 0xae, 0xe5, 0xd3, 0x21,
	0x00, 0xd9, 0xe8, 0x76, 0xfb, 0xa3, 0xb1, 0x52, 0x8a, 0xc7, 0xb8, 0xff, 0x5d, 0xbf, 0x3b, 0x56,
	0x04, 0xb4, 0x01, 0xe5, 0x1e, 0x3e, 0x18, 0x29, 0xa2, 0xa6, 0x43, 0x35, 0x97, 0x0f, 0x3d, 0x84,
	0x6a, 0xae, 0x29, 0xd2, 0xda, 0x61, 0xd1, 0x12, 0xda, 0xd7, 0x00, 0x8b, 0x9c, 0xab, 0x0d, 0xe0,
	0xd6, 0xa6, 0x06, 0xc4, 0x63, 0xed, 0x77, 0x09, 0x94, 0x21, 0x61, 0x1f, 0xfd, 0xf0, 0x74, 0xe0,
	0x31, 0x12, 0x9e, 0x58, 0x36, 0xb9, 0x70, 0x29, 0x15, 0xd4, 0x17, 0x2f, 0xaa, 0xbf, 0x0b, 0xb2,
	0x13, 0xd2, 0x33, 0x12, 0x72, 0xf9, 0x1a, 0x6d, 0x7d, 0x59, 0x92, 0x62, 0x06, 0x3d, 0x1b, 0xf5,
	0x78, 0x14, 0x4e, 0xa3, 0xd1, 0x1b, 0x28, 0xb3, 0xf3, 0x80, 0x70, 0x61, 0x1b, 0xed, 0x27, 0x37,
	0x65, 0x19, 0x9f, 0x07, 0x04, 0xf3, 0xc8, 0xa2, 0x5a, 0x6b, 0x45, 0xb5, 0x62, 0x15, 0xce, 0x5c,
	0xcb, 0x53, 0xe5, 0x44, 0x85, 0x78, 0x1c, 0xb7, 0xe1, 0x8f, 0xbe, 0x47, 0x22, 0x75, 0x3d, 0x69,
	0x43, 0x3e, 0x41, 0x4f, 0x01, 0x9d, 0x58, 0xae, 0x7b, 0x6c, 0xd9, 0xa7, 0x47, 0x74, 0x9e, 0x4a,
	0xdd, 0xe0, 0x8c, 0x9b, 0xf3, 0x95, 0x6c, 0x0f, 0xda, 0x13, 0xb8, 0x53, 0x28, 0x2b, 0xf6, 0xf7,
	0x6d, 0x1f, 0x0f, 0xfb, 0xef, 0x94, 0x12, 0xaa, 0x43, 0xe5, 0xd0, 0xec, 0x63, 0x73, 0x64, 0x74,
	0xfb, 0x8a, 0xa0, 0xbd, 0x87, 0xfa, 0xd2, 0xf6, 0x63, 0xff, 0x87, 0x07, 0xc3, 0xbe, 0x52, 0x42,
	0x35, 0xd8, 0x38, 0x1c, 0x99, 0x63, 0xdc, 0x37, 0xf6, 0x15, 0x01, 0x35, 0x00, 0x7a, 0x07, 0x3f,
	0x0c, 0xd3, 0xb9, 0x88, 0x36, 0xa1, 0xde, 0x19, 0xf4, 0x06, 0xb8, 0xdf, 0x1d, 0x0f, 0x0e, 0x86,
	0xc6, 0x3b, 0x45, 0x8a, 0x03, 0x3a, 0xb8, 0x6f, 0xbc, 0x3d, 0x38, 0x1c, 0x2b, 0x65, 0xed, 0x18,
	0x36, 0x8b, 0x4a, 0x45, 0x68, 0x1f, 0x90, 0x97, 0xfc, 0xb9, 0x28, 0x25, 0xb9, 0x5e, 0xaa, 0xed,
	0xff, 0xad, 0x96, 0x19, 0x6f, 0x7a, 0x45, 0x3a, 0xed, 0x35, 0x54, 0x53, 0xd8, 0x07, 0xdf, 0xbb,
	0x45, 0xc3, 0x68, 0x43, 0xa8, 0xe5, 0x08, 0xa2, 0xf8, 0xbb, 0x37, 0xdf, 0x5f, 0xe2, 0x84, 0x70,
	0xd9, 0x77, 0x2f, 0x17, 0x82, 0x6b, 0x5e, 0x2e, 0x5e, 0xfb, 0x43, 0x80, 0x46, 0xba, 0x6a, 0x12,
	0xc6, 0xa8, 0x37, 0x41, 0x2f, 0x40, 0x8e, 0x98, 0xc5, 0x66, 0xc9, 0x39, 0x68, 0xb4, 0x1f, 0x5d,
	0xca, 0x95, 0xa2, 0x75, 0x93, 0x43, 0x71, 0x1a, 0x92, 0x3f, 0x45, 0xe2, 0xe5, 0xd7, 0x98, 0x94,
	0xbb, 0xc6, 0x54, 0x58, 0x9f, 0x58, 0x8c, 0x7c, 0xb4, 0xce, 0x79, 0xe7, 0x56, 0xf0, 0x7c, 0x8a,
	0x14, 0x90, 0x1c, 0x2f, 0x6e, 0xc3, 0xb8, 0xaf, 0xe2, 0xa1, 0x66, 0x80, 0x9c, 0xe4, 0xca, 0x39,
	0x0e, 0x20, 0x9b, 0x63, 0x63, 0x3c, 0xe8, 0x2a, 0x42, 0x3c, 0xee, 0x7d, 0xdb, 0x1d, 0x9d, 0x7d,
	0xae, 0x88, 0xd9, 0xf8, 0x4b, 0x45, 0x42, 0x15, 0x58, 0x33, 0xdf, 0x19, 0x46, 0x57, 0x29, 0x6b,
	0x6d, 0xa8, 0xf7, 0x86, 0xe6, 0xae, 0x1f, 0x7e, 0xb4, 0x42, 0x87, 0x84, 0x11, 0xfa, 0x3f, 0xd4,
	0x68, 0x70, 0xe1, 0xb3, 0x51, 0xa5, 0xc1, 0xe2, 0xab, 0xf1, 0x82, 0xc7, 0x18, 0x98, 0xd8, 0x7e,
	0xe8, 0x98, 0x84, 0xc5, 0x75, 0x78, 0xd6, 0x94, 0xa4, 0xae, 0xf1, 0x31, 0xda, 0x02, 0xf9, 0xcc,
	0x72, 0x67, 0xe9, 0xb3, 0xa3, 0x82, 0xd3, 0x99, 0xf6, 0x00, 0xaa, 0x99, 0xf9, 0x83, 0x5e, 0xd1,
	0x6e, 0x4d, 0x05, 0x39, 0x76, 0xe1, 0xe2, 0x4a, 0xfb, 0x57, 0x11, 0x94, 0x2c, 0xd2, 0x24, 0xe1,
	0x19, 0xb5, 0x09, 0xea, 0x80, 0x7c, 0x18, 0x38, 0x16, 0x23, 0xe8, 0x9a, 0xce, 0x6b, 0x6e, 0xe9,
	0xc9, 0x63, 0x4a, 0x9f, 0x3f, 0xa6, 0xf4, 0x7e, 0xfc, 0x98, 0xd2, 0x4a, 0x68, 0x0f, 0xa0, 0x33,
	0x73, 0x4f, 0x53, 0x9e, 0x87, 0xab, 0x79, 0xa2, 0x15, 0x44, 0x5d, 0x90, 0xf7, 0x08, 0x33, 0x5c,
	0x17, 0x5d, 0x81, 0x69, 0x5e, 0x47, 0xae, 0x95, 0x50, 0x07, 0xa4, 0x3d, 0xc2, 0x50, 0xa1, 0x5b,
	0x73, 0x92, 0x35, 0xaf, 0xa9, 0x54, 0x2b, 0xb5, 0x7f, 0x92, 0xa0, 0x1a, 0xab, 0x38, 0x57, 0xe9,
	0x25, 0xc8, 0xdd, 0x90, 0xc4, 0xd5, 0x5d, 0x7d, 0x08, 0x56, 0xd4, 0xf5, 0x32, 0x13, 0xf9, 0x56,
	0xe1, 0x9d, 0x25, 0x7d, 0x9b, 0x57, 0x52, 0xac, 0x92, 0xf6, 0xd5, 0xb5, 0xd2, 0xae, 0xe0, 0xd5,
	0x4a, 0xe8, 0xab, 0x44, 0xd5, 0xbb, 0xcb, 0xa0, 0xa4, 0xd3, 0x9a, 0x57, 0x57, 0xc5, 0x23, 0xe5,
	0x1e, 0x71, 0x09, 0x23, 0x57, 0x04, 0x5f, 0xb9, 0xe7, 0xf6, 0x7b, 0x50, 0x8d, 0x20, 0x70, 0xa9,
	0xcd, 0x1f, 0xa2, 0xc9, 0x23, 0x7d, 0xee, 0xc8, 0x37, 0x20, 0xc5, 0x07, 0xe7, 0xf2, 0xc7, 0x68,
	0x02, 0x5d, 0xc1, 0xfc, 0x3d, 0x6c, 0x65, 0x76, 0xff, 0x97, 0xbc, 0xbf, 0x88, 0x00, 0xbd, 0xa1,
	0xb9, 0x68, 0x9b, 0xb2, 0x49, 0x98, 0x51, 0x64, 0x5b, 0x3a, 0xfb, 0x2b, 0x3c, 0x7b, 0x03, 0xeb,
	0x89, 0x72, 0xb7, 0x66, 0xd8, 0x85, 0xba, 0x49, 0x58, 0xee, 0x72, 0xba, 0xc8, 0xb3, 0x58, 0x5c,
	0xc1, 0x33, 0x00, 0x25, 0xd9, 0xc9, 0xbf, 0xa6, 0xea, 0x6c, 0x7f, 0xb8, 0x67, 0xfb, 0x0e, 0xd1,
	0xa3, 0xa9, 0x15, 0xb2, 0xa7, 0xc4, 0x99, 0x10, 0xdd, 0xf6, 0xa7, 0xcf, 0x88, 0x6b, 0x1d, 0xcb,
	0x1c, 0xfc, 0xfc, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x1a, 0xb9, 0xc8, 0xcf, 0x0d, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InterfaceServiceClient is the client API for InterfaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InterfaceServiceClient interface {
	Update(ctx context.Context, in *NetworkInterface, opts ...grpc.CallOption) (*empty.Empty, error)
	BulkUpdate(ctx context.Context, in *NetworkInterfaces, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NetworkInterfaces, error)
	Get(ctx context.Context, in *InterfaceID, opts ...grpc.CallOption) (*NetworkInterface, error)
}

type interfaceServiceClient struct {
	cc *grpc.ClientConn
}

func NewInterfaceServiceClient(cc *grpc.ClientConn) InterfaceServiceClient {
	return &interfaceServiceClient{cc}
}

func (c *interfaceServiceClient) Update(ctx context.Context, in *NetworkInterface, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.InterfaceService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceServiceClient) BulkUpdate(ctx context.Context, in *NetworkInterfaces, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.InterfaceService/BulkUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceServiceClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NetworkInterfaces, error) {
	out := new(NetworkInterfaces)
	err := c.cc.Invoke(ctx, "/openness.ela.InterfaceService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceServiceClient) Get(ctx context.Context, in *InterfaceID, opts ...grpc.CallOption) (*NetworkInterface, error) {
	out := new(NetworkInterface)
	err := c.cc.Invoke(ctx, "/openness.ela.InterfaceService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterfaceServiceServer is the server API for InterfaceService service.
type InterfaceServiceServer interface {
	Update(context.Context, *NetworkInterface) (*empty.Empty, error)
	BulkUpdate(context.Context, *NetworkInterfaces) (*empty.Empty, error)
	GetAll(context.Context, *empty.Empty) (*NetworkInterfaces, error)
	Get(context.Context, *InterfaceID) (*NetworkInterface, error)
}

func RegisterInterfaceServiceServer(s *grpc.Server, srv InterfaceServiceServer) {
	s.RegisterService(&_InterfaceService_serviceDesc, srv)
}

func _InterfaceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInterface)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.InterfaceService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServiceServer).Update(ctx, req.(*NetworkInterface))
	}
	return interceptor(ctx, in, info, handler)
}

func _InterfaceService_BulkUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInterfaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServiceServer).BulkUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.InterfaceService/BulkUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServiceServer).BulkUpdate(ctx, req.(*NetworkInterfaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _InterfaceService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.InterfaceService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServiceServer).GetAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InterfaceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterfaceID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.InterfaceService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServiceServer).Get(ctx, req.(*InterfaceID))
	}
	return interceptor(ctx, in, info, handler)
}

var _InterfaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openness.ela.InterfaceService",
	HandlerType: (*InterfaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _InterfaceService_Update_Handler,
		},
		{
			MethodName: "BulkUpdate",
			Handler:    _InterfaceService_BulkUpdate_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _InterfaceService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _InterfaceService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ela.proto",
}

// ZoneServiceClient is the client API for ZoneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ZoneServiceClient interface {
	Create(ctx context.Context, in *NetworkZone, opts ...grpc.CallOption) (*empty.Empty, error)
	Update(ctx context.Context, in *NetworkZone, opts ...grpc.CallOption) (*empty.Empty, error)
	BulkUpdate(ctx context.Context, in *NetworkZones, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NetworkZones, error)
	Get(ctx context.Context, in *ZoneID, opts ...grpc.CallOption) (*NetworkZone, error)
	Delete(ctx context.Context, in *ZoneID, opts ...grpc.CallOption) (*empty.Empty, error)
}

type zoneServiceClient struct {
	cc *grpc.ClientConn
}

func NewZoneServiceClient(cc *grpc.ClientConn) ZoneServiceClient {
	return &zoneServiceClient{cc}
}

func (c *zoneServiceClient) Create(ctx context.Context, in *NetworkZone, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.ZoneService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) Update(ctx context.Context, in *NetworkZone, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.ZoneService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) BulkUpdate(ctx context.Context, in *NetworkZones, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.ZoneService/BulkUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NetworkZones, error) {
	out := new(NetworkZones)
	err := c.cc.Invoke(ctx, "/openness.ela.ZoneService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) Get(ctx context.Context, in *ZoneID, opts ...grpc.CallOption) (*NetworkZone, error) {
	out := new(NetworkZone)
	err := c.cc.Invoke(ctx, "/openness.ela.ZoneService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) Delete(ctx context.Context, in *ZoneID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.ZoneService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZoneServiceServer is the server API for ZoneService service.
type ZoneServiceServer interface {
	Create(context.Context, *NetworkZone) (*empty.Empty, error)
	Update(context.Context, *NetworkZone) (*empty.Empty, error)
	BulkUpdate(context.Context, *NetworkZones) (*empty.Empty, error)
	GetAll(context.Context, *empty.Empty) (*NetworkZones, error)
	Get(context.Context, *ZoneID) (*NetworkZone, error)
	Delete(context.Context, *ZoneID) (*empty.Empty, error)
}

func RegisterZoneServiceServer(s *grpc.Server, srv ZoneServiceServer) {
	s.RegisterService(&_ZoneService_serviceDesc, srv)
}

func _ZoneService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.ZoneService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).Create(ctx, req.(*NetworkZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.ZoneService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).Update(ctx, req.(*NetworkZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_BulkUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkZones)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).BulkUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.ZoneService/BulkUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).BulkUpdate(ctx, req.(*NetworkZones))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.ZoneService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).GetAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.ZoneService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).Get(ctx, req.(*ZoneID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.ZoneService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).Delete(ctx, req.(*ZoneID))
	}
	return interceptor(ctx, in, info, handler)
}

var _ZoneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openness.ela.ZoneService",
	HandlerType: (*ZoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ZoneService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ZoneService_Update_Handler,
		},
		{
			MethodName: "BulkUpdate",
			Handler:    _ZoneService_BulkUpdate_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ZoneService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ZoneService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ZoneService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ela.proto",
}

// ApplicationPolicyServiceClient is the client API for ApplicationPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationPolicyServiceClient interface {
	Set(ctx context.Context, in *TrafficPolicy, opts ...grpc.CallOption) (*empty.Empty, error)
}

type applicationPolicyServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationPolicyServiceClient(cc *grpc.ClientConn) ApplicationPolicyServiceClient {
	return &applicationPolicyServiceClient{cc}
}

func (c *applicationPolicyServiceClient) Set(ctx context.Context, in *TrafficPolicy, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.ApplicationPolicyService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationPolicyServiceServer is the server API for ApplicationPolicyService service.
type ApplicationPolicyServiceServer interface {
	Set(context.Context, *TrafficPolicy) (*empty.Empty, error)
}

func RegisterApplicationPolicyServiceServer(s *grpc.Server, srv ApplicationPolicyServiceServer) {
	s.RegisterService(&_ApplicationPolicyService_serviceDesc, srv)
}

func _ApplicationPolicyService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrafficPolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationPolicyServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.ApplicationPolicyService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationPolicyServiceServer).Set(ctx, req.(*TrafficPolicy))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationPolicyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openness.ela.ApplicationPolicyService",
	HandlerType: (*ApplicationPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _ApplicationPolicyService_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ela.proto",
}

// InterfacePolicyServiceClient is the client API for InterfacePolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InterfacePolicyServiceClient interface {
	Set(ctx context.Context, in *TrafficPolicy, opts ...grpc.CallOption) (*empty.Empty, error)
}

type interfacePolicyServiceClient struct {
	cc *grpc.ClientConn
}

func NewInterfacePolicyServiceClient(cc *grpc.ClientConn) InterfacePolicyServiceClient {
	return &interfacePolicyServiceClient{cc}
}

func (c *interfacePolicyServiceClient) Set(ctx context.Context, in *TrafficPolicy, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.InterfacePolicyService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterfacePolicyServiceServer is the server API for InterfacePolicyService service.
type InterfacePolicyServiceServer interface {
	Set(context.Context, *TrafficPolicy) (*empty.Empty, error)
}

func RegisterInterfacePolicyServiceServer(s *grpc.Server, srv InterfacePolicyServiceServer) {
	s.RegisterService(&_InterfacePolicyService_serviceDesc, srv)
}

func _InterfacePolicyService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrafficPolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfacePolicyServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.InterfacePolicyService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfacePolicyServiceServer).Set(ctx, req.(*TrafficPolicy))
	}
	return interceptor(ctx, in, info, handler)
}

var _InterfacePolicyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openness.ela.InterfacePolicyService",
	HandlerType: (*InterfacePolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _InterfacePolicyService_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ela.proto",
}

// DNSServiceClient is the client API for DNSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DNSServiceClient interface {
	SetA(ctx context.Context, in *DNSARecordSet, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteA(ctx context.Context, in *DNSARecordSet, opts ...grpc.CallOption) (*empty.Empty, error)
	SetForwarders(ctx context.Context, in *DNSForwarders, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteForwarders(ctx context.Context, in *DNSForwarders, opts ...grpc.CallOption) (*empty.Empty, error)
}

type dNSServiceClient struct {
	cc *grpc.ClientConn
}

func NewDNSServiceClient(cc *grpc.ClientConn) DNSServiceClient {
	return &dNSServiceClient{cc}
}

func (c *dNSServiceClient) SetA(ctx context.Context, in *DNSARecordSet, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.DNSService/SetA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSServiceClient) DeleteA(ctx context.Context, in *DNSARecordSet, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.DNSService/DeleteA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSServiceClient) SetForwarders(ctx context.Context, in *DNSForwarders, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.DNSService/SetForwarders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSServiceClient) DeleteForwarders(ctx context.Context, in *DNSForwarders, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/openness.ela.DNSService/DeleteForwarders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSServiceServer is the server API for DNSService service.
type DNSServiceServer interface {
	SetA(context.Context, *DNSARecordSet) (*empty.Empty, error)
	DeleteA(context.Context, *DNSARecordSet) (*empty.Empty, error)
	SetForwarders(context.Context, *DNSForwarders) (*empty.Empty, error)
	DeleteForwarders(context.Context, *DNSForwarders) (*empty.Empty, error)
}

func RegisterDNSServiceServer(s *grpc.Server, srv DNSServiceServer) {
	s.RegisterService(&_DNSService_serviceDesc, srv)
}

func _DNSService_SetA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DNSARecordSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).SetA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.DNSService/SetA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).SetA(ctx, req.(*DNSARecordSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSService_DeleteA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DNSARecordSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).DeleteA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.DNSService/DeleteA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).DeleteA(ctx, req.(*DNSARecordSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSService_SetForwarders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DNSForwarders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).SetForwarders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.DNSService/SetForwarders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).SetForwarders(ctx, req.(*DNSForwarders))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSService_DeleteForwarders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DNSForwarders)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).DeleteForwarders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openness.ela.DNSService/DeleteForwarders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).DeleteForwarders(ctx, req.(*DNSForwarders))
	}
	return interceptor(ctx, in, info, handler)
}

var _DNSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openness.ela.DNSService",
	HandlerType: (*DNSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetA",
			Handler:    _DNSService_SetA_Handler,
		},
		{
			MethodName: "DeleteA",
			Handler:    _DNSService_DeleteA_Handler,
		},
		{
			MethodName: "SetForwarders",
			Handler:    _DNSService_SetForwarders_Handler,
		},
		{
			MethodName: "DeleteForwarders",
			Handler:    _DNSService_DeleteForwarders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ela.proto",
}
