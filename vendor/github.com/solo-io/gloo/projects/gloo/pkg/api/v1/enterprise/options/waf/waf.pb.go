// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/waf/waf.proto

package waf

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	waf "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/waf"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Settings struct {
	// Disable waf on this resource (if omitted defaults to false).
	// If a route/virtual host is configured with WAF, you must explicitly disable its WAF,
	// i.e., it will not inherit the disabled status of its parent
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Custom massage to display if an intervention occurs.
	CustomInterventionMessage string `protobuf:"bytes,2,opt,name=custom_intervention_message,json=customInterventionMessage,proto3" json:"custom_intervention_message,omitempty"`
	// Add OWASP core rule set
	// if nil will not be added
	CoreRuleSet *CoreRuleSet `protobuf:"bytes,3,opt,name=core_rule_set,json=coreRuleSet,proto3" json:"core_rule_set,omitempty"`
	// Custom rule sets rules to add
	RuleSets             []*waf.RuleSet `protobuf:"bytes,4,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_0151c80aefddd633, []int{0}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Settings) GetCustomInterventionMessage() string {
	if m != nil {
		return m.CustomInterventionMessage
	}
	return ""
}

func (m *Settings) GetCoreRuleSet() *CoreRuleSet {
	if m != nil {
		return m.CoreRuleSet
	}
	return nil
}

func (m *Settings) GetRuleSets() []*waf.RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

type CoreRuleSet struct {
	// Optional custom settings for the OWASP core rule set.
	// For an example on the configuration options see: https://github.com/SpiderLabs/owasp-modsecurity-crs/blob/v3.2/dev/crs-setup.conf.example
	// The same rules apply to these options as do to the `RuleSet`s. The file option is better if possible.
	//
	// Types that are valid to be assigned to CustomSettingsType:
	//	*CoreRuleSet_CustomSettingsString
	//	*CoreRuleSet_CustomSettingsFile
	CustomSettingsType   isCoreRuleSet_CustomSettingsType `protobuf_oneof:"CustomSettingsType"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CoreRuleSet) Reset()         { *m = CoreRuleSet{} }
func (m *CoreRuleSet) String() string { return proto.CompactTextString(m) }
func (*CoreRuleSet) ProtoMessage()    {}
func (*CoreRuleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0151c80aefddd633, []int{1}
}
func (m *CoreRuleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreRuleSet.Unmarshal(m, b)
}
func (m *CoreRuleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreRuleSet.Marshal(b, m, deterministic)
}
func (m *CoreRuleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreRuleSet.Merge(m, src)
}
func (m *CoreRuleSet) XXX_Size() int {
	return xxx_messageInfo_CoreRuleSet.Size(m)
}
func (m *CoreRuleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreRuleSet.DiscardUnknown(m)
}

var xxx_messageInfo_CoreRuleSet proto.InternalMessageInfo

type isCoreRuleSet_CustomSettingsType interface {
	isCoreRuleSet_CustomSettingsType()
	Equal(interface{}) bool
}

type CoreRuleSet_CustomSettingsString struct {
	CustomSettingsString string `protobuf:"bytes,2,opt,name=custom_settings_string,json=customSettingsString,proto3,oneof" json:"custom_settings_string,omitempty"`
}
type CoreRuleSet_CustomSettingsFile struct {
	CustomSettingsFile string `protobuf:"bytes,3,opt,name=custom_settings_file,json=customSettingsFile,proto3,oneof" json:"custom_settings_file,omitempty"`
}

func (*CoreRuleSet_CustomSettingsString) isCoreRuleSet_CustomSettingsType() {}
func (*CoreRuleSet_CustomSettingsFile) isCoreRuleSet_CustomSettingsType()   {}

func (m *CoreRuleSet) GetCustomSettingsType() isCoreRuleSet_CustomSettingsType {
	if m != nil {
		return m.CustomSettingsType
	}
	return nil
}

func (m *CoreRuleSet) GetCustomSettingsString() string {
	if x, ok := m.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsString); ok {
		return x.CustomSettingsString
	}
	return ""
}

func (m *CoreRuleSet) GetCustomSettingsFile() string {
	if x, ok := m.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsFile); ok {
		return x.CustomSettingsFile
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CoreRuleSet) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CoreRuleSet_CustomSettingsString)(nil),
		(*CoreRuleSet_CustomSettingsFile)(nil),
	}
}

func init() {
	proto.RegisterType((*Settings)(nil), "waf.options.gloo.solo.io.Settings")
	proto.RegisterType((*CoreRuleSet)(nil), "waf.options.gloo.solo.io.CoreRuleSet")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/waf/waf.proto", fileDescriptor_0151c80aefddd633)
}

var fileDescriptor_0151c80aefddd633 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8a, 0x13, 0x41,
	0x10, 0x75, 0x76, 0x45, 0xb2, 0x1d, 0xbc, 0x34, 0x41, 0xc6, 0x08, 0x12, 0x16, 0x84, 0x5c, 0xec,
	0xd6, 0x08, 0x1e, 0x3d, 0xec, 0x82, 0xba, 0x87, 0x1c, 0x9c, 0x78, 0xf2, 0x32, 0x4c, 0x26, 0x95,
	0xde, 0xd6, 0x9e, 0xae, 0xa6, 0xab, 0x66, 0x34, 0x37, 0xff, 0xc0, 0xdf, 0xf0, 0xbb, 0xfc, 0x12,
	0xe9, 0x99, 0x44, 0x77, 0x15, 0x25, 0x87, 0x81, 0xf7, 0xaa, 0xea, 0x55, 0xcd, 0xab, 0x2e, 0xb1,
	0x34, 0x96, 0xaf, 0xdb, 0xb5, 0xaa, 0xb1, 0xd1, 0x84, 0x0e, 0x9f, 0x5a, 0xd4, 0xc6, 0x21, 0xea,
	0x10, 0xf1, 0x23, 0xd4, 0x4c, 0x03, 0xab, 0x82, 0xd5, 0xdd, 0x73, 0x0d, 0x9e, 0x21, 0x86, 0x68,
	0x09, 0x34, 0x06, 0xb6, 0xe8, 0x49, 0x7f, 0xae, 0xb6, 0xe9, 0x53, 0x21, 0x22, 0xa3, 0xcc, 0x13,
	0xdc, 0xa7, 0x54, 0x52, 0xaa, 0xd4, 0x54, 0x59, 0x9c, 0xbe, 0x3b, 0x7e, 0x10, 0x7c, 0x61, 0x88,
	0xbe, 0x72, 0x1a, 0x7c, 0x87, 0xbb, 0x9e, 0x7a, 0xfa, 0x7b, 0xd8, 0x74, 0x62, 0xd0, 0x60, 0x0f,
	0x75, 0x42, 0x43, 0xf4, 0xfc, 0xeb, 0x89, 0x18, 0xad, 0x80, 0xd9, 0x7a, 0x43, 0x72, 0x2a, 0x46,
	0x1b, 0x4b, 0xd5, 0xda, 0xc1, 0x26, 0xcf, 0x66, 0xd9, 0x7c, 0x54, 0xfc, 0xe2, 0xf2, 0x95, 0x78,
	0x54, 0xb7, 0xc4, 0xd8, 0x94, 0x36, 0xd9, 0xea, 0xc0, 0xa7, 0xdf, 0x2e, 0x1b, 0x20, 0xaa, 0x0c,
	0xe4, 0x27, 0xb3, 0x6c, 0x7e, 0x56, 0x3c, 0x1c, 0x4a, 0xae, 0x6e, 0x54, 0x2c, 0x87, 0x02, 0x79,
	0x25, 0xee, 0xd7, 0x18, 0xa1, 0x8c, 0xad, 0x83, 0x92, 0x80, 0xf3, 0xd3, 0x59, 0x36, 0x1f, 0x2f,
	0x9e, 0xa8, 0x7f, 0xed, 0x40, 0x5d, 0x62, 0x84, 0xa2, 0x75, 0xb0, 0x02, 0x2e, 0xc6, 0xf5, 0x6f,
	0x22, 0x97, 0xe2, 0xec, 0xd0, 0x85, 0xf2, 0xbb, 0xb3, 0xd3, 0xf9, 0x78, 0xf1, 0x4c, 0xf5, 0xee,
	0x55, 0x8d, 0x7e, 0x6b, 0x8d, 0xda, 0x5a, 0xc7, 0x10, 0xd5, 0x35, 0x73, 0x50, 0x0d, 0x6e, 0x08,
	0xea, 0x36, 0x5a, 0xde, 0xa9, 0x6e, 0xa1, 0x0e, 0x1d, 0x47, 0x71, 0x00, 0x74, 0xfe, 0x2d, 0x13,
	0xe3, 0x1b, 0xb3, 0xe4, 0x4b, 0xf1, 0x60, 0xef, 0x94, 0xf6, 0x8b, 0x29, 0x89, 0xa3, 0xf5, 0x66,
	0x30, 0xf9, 0xf6, 0x4e, 0x31, 0x19, 0xf2, 0x87, 0xbd, 0xad, 0xfa, 0xac, 0x5c, 0x88, 0xc9, 0x9f,
	0xba, 0xad, 0x75, 0xd0, 0x1b, 0x4d, 0x2a, 0x79, 0x5b, 0xf5, 0xda, 0x3a, 0xb8, 0x98, 0x08, 0x79,
	0x79, 0x2b, 0xfa, 0x7e, 0x17, 0xe0, 0x62, 0xf9, 0xfd, 0xc7, 0xe3, 0xec, 0xc3, 0x9b, 0xe3, 0x6e,
	0x20, 0x7c, 0x32, 0xff, 0x3f, 0xb8, 0xf5, 0xbd, 0xfe, 0xa9, 0x5f, 0xfc, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0x6b, 0x9d, 0xc5, 0xbe, 0x02, 0x00, 0x00,
}

func (this *Settings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Disabled != that1.Disabled {
		return false
	}
	if this.CustomInterventionMessage != that1.CustomInterventionMessage {
		return false
	}
	if !this.CoreRuleSet.Equal(that1.CoreRuleSet) {
		return false
	}
	if len(this.RuleSets) != len(that1.RuleSets) {
		return false
	}
	for i := range this.RuleSets {
		if !this.RuleSets[i].Equal(that1.RuleSets[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CoreRuleSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet)
	if !ok {
		that2, ok := that.(CoreRuleSet)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.CustomSettingsType == nil {
		if this.CustomSettingsType != nil {
			return false
		}
	} else if this.CustomSettingsType == nil {
		return false
	} else if !this.CustomSettingsType.Equal(that1.CustomSettingsType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CoreRuleSet_CustomSettingsString) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet_CustomSettingsString)
	if !ok {
		that2, ok := that.(CoreRuleSet_CustomSettingsString)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CustomSettingsString != that1.CustomSettingsString {
		return false
	}
	return true
}
func (this *CoreRuleSet_CustomSettingsFile) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet_CustomSettingsFile)
	if !ok {
		that2, ok := that.(CoreRuleSet_CustomSettingsFile)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CustomSettingsFile != that1.CustomSettingsFile {
		return false
	}
	return true
}
