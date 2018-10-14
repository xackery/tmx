// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/out.proto

package pb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OutMap struct {
	Source     string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Width      int64  `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height     int64  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	TileWidth  int64  `protobuf:"varint,5,opt,name=tile_width,json=tileWidth,proto3" json:"tile_width,omitempty"`
	TileHeight int64  `protobuf:"varint,6,opt,name=tile_height,json=tileHeight,proto3" json:"tile_height,omitempty"`
	// repeated Property properties = 7;
	Layers               []*OutLayer `protobuf:"bytes,8,rep,name=layers,proto3" json:"layers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OutMap) Reset()         { *m = OutMap{} }
func (m *OutMap) String() string { return proto.CompactTextString(m) }
func (*OutMap) ProtoMessage()    {}
func (*OutMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_705866a120287d7e, []int{0}
}

func (m *OutMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutMap.Unmarshal(m, b)
}
func (m *OutMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutMap.Marshal(b, m, deterministic)
}
func (m *OutMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutMap.Merge(m, src)
}
func (m *OutMap) XXX_Size() int {
	return xxx_messageInfo_OutMap.Size(m)
}
func (m *OutMap) XXX_DiscardUnknown() {
	xxx_messageInfo_OutMap.DiscardUnknown(m)
}

var xxx_messageInfo_OutMap proto.InternalMessageInfo

func (m *OutMap) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *OutMap) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *OutMap) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *OutMap) GetTileWidth() int64 {
	if m != nil {
		return m.TileWidth
	}
	return 0
}

func (m *OutMap) GetTileHeight() int64 {
	if m != nil {
		return m.TileHeight
	}
	return 0
}

func (m *OutMap) GetLayers() []*OutLayer {
	if m != nil {
		return m.Layers
	}
	return nil
}

type OutLayer struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Opacity              float32    `protobuf:"fixed32,2,opt,name=opacity,proto3" json:"opacity,omitempty"`
	Tiles                []*OutTile `protobuf:"bytes,4,rep,name=tiles,proto3" json:"tiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OutLayer) Reset()         { *m = OutLayer{} }
func (m *OutLayer) String() string { return proto.CompactTextString(m) }
func (*OutLayer) ProtoMessage()    {}
func (*OutLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_705866a120287d7e, []int{1}
}

func (m *OutLayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutLayer.Unmarshal(m, b)
}
func (m *OutLayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutLayer.Marshal(b, m, deterministic)
}
func (m *OutLayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutLayer.Merge(m, src)
}
func (m *OutLayer) XXX_Size() int {
	return xxx_messageInfo_OutLayer.Size(m)
}
func (m *OutLayer) XXX_DiscardUnknown() {
	xxx_messageInfo_OutLayer.DiscardUnknown(m)
}

var xxx_messageInfo_OutLayer proto.InternalMessageInfo

func (m *OutLayer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OutLayer) GetOpacity() float32 {
	if m != nil {
		return m.Opacity
	}
	return 0
}

func (m *OutLayer) GetTiles() []*OutTile {
	if m != nil {
		return m.Tiles
	}
	return nil
}

type OutTile struct {
	Gid                  uint32   `protobuf:"fixed32,1,opt,name=gid,proto3" json:"gid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutTile) Reset()         { *m = OutTile{} }
func (m *OutTile) String() string { return proto.CompactTextString(m) }
func (*OutTile) ProtoMessage()    {}
func (*OutTile) Descriptor() ([]byte, []int) {
	return fileDescriptor_705866a120287d7e, []int{2}
}

func (m *OutTile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutTile.Unmarshal(m, b)
}
func (m *OutTile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutTile.Marshal(b, m, deterministic)
}
func (m *OutTile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutTile.Merge(m, src)
}
func (m *OutTile) XXX_Size() int {
	return xxx_messageInfo_OutTile.Size(m)
}
func (m *OutTile) XXX_DiscardUnknown() {
	xxx_messageInfo_OutTile.DiscardUnknown(m)
}

var xxx_messageInfo_OutTile proto.InternalMessageInfo

func (m *OutTile) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func init() {
	proto.RegisterType((*OutMap)(nil), "pb.OutMap")
	proto.RegisterType((*OutLayer)(nil), "pb.OutLayer")
	proto.RegisterType((*OutTile)(nil), "pb.OutTile")
}

func init() { proto.RegisterFile("pb/out.proto", fileDescriptor_705866a120287d7e) }

var fileDescriptor_705866a120287d7e = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xe9, 0xbf, 0x74, 0x77, 0xba, 0x07, 0x19, 0x44, 0x02, 0x22, 0xd6, 0xe2, 0xa1, 0xa7,
	0x0a, 0xfa, 0x25, 0x3c, 0x28, 0x0b, 0x41, 0xf0, 0xe0, 0x41, 0xda, 0xdd, 0xb0, 0x0d, 0x54, 0x13,
	0xda, 0x29, 0xb2, 0xdf, 0xcb, 0x0f, 0xb8, 0x64, 0x9a, 0xde, 0xe6, 0xbd, 0x97, 0xdf, 0x4b, 0x32,
	0xb0, 0x73, 0xdd, 0x93, 0x9d, 0xa9, 0x71, 0xa3, 0x25, 0x8b, 0xb1, 0xeb, 0xaa, 0xff, 0x08, 0xc4,
	0x7e, 0xa6, 0xf7, 0xd6, 0xe1, 0x0d, 0x88, 0xc9, 0xce, 0xe3, 0x41, 0xcb, 0xb8, 0x8c, 0xea, 0xad,
	0x0a, 0x0a, 0xaf, 0x21, 0xfb, 0x33, 0x47, 0xea, 0x65, 0x52, 0x46, 0x75, 0xa2, 0x16, 0xe1, 0x4f,
	0xf7, 0xda, 0x9c, 0x7a, 0x92, 0x29, 0xdb, 0x41, 0xe1, 0x1d, 0x00, 0x99, 0x41, 0x7f, 0x2f, 0x48,
	0xc6, 0xd9, 0xd6, 0x3b, 0x9f, 0x8c, 0xdd, 0x43, 0xc1, 0x71, 0x60, 0x05, 0xe7, 0x4c, 0xbc, 0x2e,
	0xfc, 0x23, 0x88, 0xa1, 0x3d, 0xeb, 0x71, 0x92, 0x9b, 0x32, 0xa9, 0x8b, 0xe7, 0x5d, 0xe3, 0xba,
	0x66, 0x3f, 0xd3, 0x9b, 0x37, 0x55, 0xc8, 0xaa, 0x2f, 0xd8, 0xac, 0x1e, 0x22, 0xa4, 0xbf, 0xed,
	0x8f, 0x96, 0x11, 0xbf, 0x9a, 0x67, 0x94, 0x90, 0x5b, 0xd7, 0x1e, 0x0c, 0x9d, 0xf9, 0x33, 0xb1,
	0x5a, 0x25, 0x3e, 0x40, 0xe6, 0x6f, 0x9b, 0x64, 0xca, 0xf5, 0x45, 0xa8, 0xff, 0x30, 0x83, 0x56,
	0x4b, 0x52, 0xdd, 0x42, 0x1e, 0x1c, 0xbc, 0x82, 0xe4, 0x64, 0x8e, 0x5c, 0x9d, 0x2b, 0x3f, 0x76,
	0x82, 0x77, 0xf7, 0x72, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xa9, 0x11, 0x4f, 0x4b, 0x01, 0x00,
	0x00,
}
