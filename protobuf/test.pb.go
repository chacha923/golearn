// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Test
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}
var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Test struct {
	Label         *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type          *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps          []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	Optionalgroup *Test_OptionalGroup `protobuf:"group,4,opt,name=OptionalGroup,json=optionalgroup" json:"optionalgroup,omitempty"`
	// Types that are valid to be assigned to Union:
	//	*Test_Number
	//	*Test_Name
	Union            isTest_Union `protobuf_oneof:"union"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_Test_Type int32 = 77

type isTest_Union interface {
	isTest_Union()
}

type Test_Number struct {
	Number int32 `protobuf:"varint,6,opt,name=number,oneof"`
}
type Test_Name struct {
	Name string `protobuf:"bytes,7,opt,name=name,oneof"`
}

func (*Test_Number) isTest_Union() {}
func (*Test_Name) isTest_Union()   {}

func (m *Test) GetUnion() isTest_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *Test) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Test_Type
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func (m *Test) GetOptionalgroup() *Test_OptionalGroup {
	if m != nil {
		return m.Optionalgroup
	}
	return nil
}

func (m *Test) GetNumber() int32 {
	if x, ok := m.GetUnion().(*Test_Number); ok {
		return x.Number
	}
	return 0
}

func (m *Test) GetName() string {
	if x, ok := m.GetUnion().(*Test_Name); ok {
		return x.Name
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Test) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Test_OneofMarshaler, _Test_OneofUnmarshaler, _Test_OneofSizer, []interface{}{
		(*Test_Number)(nil),
		(*Test_Name)(nil),
	}
}

func _Test_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Test)
	// union
	switch x := m.Union.(type) {
	case *Test_Number:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Number))
	case *Test_Name:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case nil:
	default:
		return fmt.Errorf("Test.Union has unexpected type %T", x)
	}
	return nil
}

func _Test_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Test)
	switch tag {
	case 6: // union.number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &Test_Number{int32(x)}
		return true, err
	case 7: // union.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &Test_Name{x}
		return true, err
	default:
		return false, nil
	}
}

func _Test_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Test)
	// union
	switch x := m.Union.(type) {
	case *Test_Number:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Number))
	case *Test_Name:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Test_OptionalGroup struct {
	RequiredField    *string `protobuf:"bytes,5,req,name=RequiredField" json:"RequiredField,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Test_OptionalGroup) Reset()                    { *m = Test_OptionalGroup{} }
func (m *Test_OptionalGroup) String() string            { return proto.CompactTextString(m) }
func (*Test_OptionalGroup) ProtoMessage()               {}
func (*Test_OptionalGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Test_OptionalGroup) GetRequiredField() string {
	if m != nil && m.RequiredField != nil {
		return *m.RequiredField
	}
	return ""
}

func init() {
	proto.RegisterType((*Test)(nil), "protobuf.Test")
	proto.RegisterType((*Test_OptionalGroup)(nil), "protobuf.Test.OptionalGroup")
	proto.RegisterEnum("protobuf.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x46, 0x2f, 0xbb, 0xc9, 0x9d, 0x37, 0xb8, 0xa0, 0xc3, 0x21, 0x41, 0x2c, 0x82, 0x58, 0x04,
	0x8b, 0x2d, 0x04, 0x39, 0xb0, 0xbc, 0xe2, 0xbc, 0x6e, 0x21, 0x58, 0xd8, 0xee, 0x72, 0xa3, 0x2c,
	0xec, 0x25, 0x31, 0x9b, 0x14, 0xfe, 0x77, 0x0b, 0xd9, 0x78, 0x16, 0x5b, 0xcd, 0xbc, 0xe1, 0x9b,
	0x79, 0x0c, 0x40, 0xa4, 0x31, 0xd6, 0x3e, 0xb8, 0xe8, 0xf0, 0x22, 0x97, 0x2e, 0x7d, 0xdc, 0xff,
	0x30, 0xe0, 0x6f, 0x34, 0x46, 0xdc, 0x80, 0x18, 0xda, 0x8e, 0x06, 0xc9, 0x54, 0xa1, 0xd7, 0xe6,
	0x0f, 0xf0, 0x06, 0x78, 0xfc, 0xf6, 0x24, 0x0b, 0xc5, 0xb4, 0x78, 0x29, 0xb6, 0x5b, 0x93, 0x19,
	0x11, 0x78, 0x20, 0x3f, 0xca, 0x52, 0x95, 0xba, 0x34, 0xb9, 0xc7, 0x1d, 0x54, 0xce, 0xc7, 0xde,
	0xd9, 0x76, 0xf8, 0x0c, 0x2e, 0x79, 0xc9, 0x15, 0xd3, 0xf0, 0x74, 0x57, 0xff, 0xcb, 0xea, 0x49,
	0x54, 0x37, 0xe7, 0xcc, 0xeb, 0x94, 0x31, 0xf3, 0x15, 0x94, 0xb0, 0xb4, 0xe9, 0xd4, 0x51, 0x90,
	0xcb, 0xc9, 0x78, 0x58, 0x98, 0x33, 0xe3, 0x06, 0xb8, 0x6d, 0x4f, 0x24, 0x57, 0x8a, 0xe9, 0xf5,
	0x61, 0x61, 0x32, 0xdd, 0x3e, 0x43, 0x35, 0xbb, 0x87, 0x0f, 0x50, 0x19, 0xfa, 0x4a, 0x7d, 0xa0,
	0xe3, 0xbe, 0xa7, 0xe1, 0x28, 0x45, 0x7e, 0x67, 0x3e, 0xdc, 0xad, 0x40, 0x24, 0xdb, 0x3b, 0xfb,
	0x78, 0x09, 0xe5, 0xbe, 0x69, 0x50, 0x00, 0x7b, 0xbf, 0xba, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x0b, 0xd5, 0x1e, 0x15, 0x23, 0x01, 0x00, 0x00,
}
