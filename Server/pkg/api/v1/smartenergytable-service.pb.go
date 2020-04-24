// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smartenergytable-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Token struct {
	RoomUser             *RoomUser `protobuf:"bytes,1,opt,name=room_user,json=roomUser,proto3" json:"room_user,omitempty"`
	ObjectIndex          int32     `protobuf:"varint,2,opt,name=object_index,json=objectIndex,proto3" json:"object_index,omitempty"`
	Position             *Vector3  `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	ObjectId             string    `protobuf:"bytes,4,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{1}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetRoomUser() *RoomUser {
	if m != nil {
		return m.RoomUser
	}
	return nil
}

func (m *Token) GetObjectIndex() int32 {
	if m != nil {
		return m.ObjectIndex
	}
	return 0
}

func (m *Token) GetPosition() *Vector3 {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Token) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

type Vector3 struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector3) Reset()         { *m = Vector3{} }
func (m *Vector3) String() string { return proto.CompactTextString(m) }
func (*Vector3) ProtoMessage()    {}
func (*Vector3) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{2}
}

func (m *Vector3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector3.Unmarshal(m, b)
}
func (m *Vector3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector3.Marshal(b, m, deterministic)
}
func (m *Vector3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector3.Merge(m, src)
}
func (m *Vector3) XXX_Size() int {
	return xxx_messageInfo_Vector3.Size(m)
}
func (m *Vector3) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector3.DiscardUnknown(m)
}

var xxx_messageInfo_Vector3 proto.InternalMessageInfo

func (m *Vector3) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector3) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vector3) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

type Room struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Objects              []*Token `protobuf:"bytes,2,rep,name=objects,proto3" json:"objects,omitempty"`
	SceneId              int32    `protobuf:"varint,3,opt,name=sceneId,proto3" json:"sceneId,omitempty"`
	UserIds              []string `protobuf:"bytes,4,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{3}
}

func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (m *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(m, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Room) GetObjects() []*Token {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *Room) GetSceneId() int32 {
	if m != nil {
		return m.SceneId
	}
	return 0
}

func (m *Room) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type RoomUser struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomUser) Reset()         { *m = RoomUser{} }
func (m *RoomUser) String() string { return proto.CompactTextString(m) }
func (*RoomUser) ProtoMessage()    {}
func (*RoomUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{4}
}

func (m *RoomUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomUser.Unmarshal(m, b)
}
func (m *RoomUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomUser.Marshal(b, m, deterministic)
}
func (m *RoomUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomUser.Merge(m, src)
}
func (m *RoomUser) XXX_Size() int {
	return xxx_messageInfo_RoomUser.Size(m)
}
func (m *RoomUser) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomUser.DiscardUnknown(m)
}

var xxx_messageInfo_RoomUser proto.InternalMessageInfo

func (m *RoomUser) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RoomUser) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type UserPosition struct {
	RoomUser             *RoomUser `protobuf:"bytes,1,opt,name=room_user,json=roomUser,proto3" json:"room_user,omitempty"`
	NewPosition          *Vector3  `protobuf:"bytes,2,opt,name=new_position,json=newPosition,proto3" json:"new_position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserPosition) Reset()         { *m = UserPosition{} }
func (m *UserPosition) String() string { return proto.CompactTextString(m) }
func (*UserPosition) ProtoMessage()    {}
func (*UserPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{5}
}

func (m *UserPosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPosition.Unmarshal(m, b)
}
func (m *UserPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPosition.Marshal(b, m, deterministic)
}
func (m *UserPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPosition.Merge(m, src)
}
func (m *UserPosition) XXX_Size() int {
	return xxx_messageInfo_UserPosition.Size(m)
}
func (m *UserPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPosition.DiscardUnknown(m)
}

var xxx_messageInfo_UserPosition proto.InternalMessageInfo

func (m *UserPosition) GetRoomUser() *RoomUser {
	if m != nil {
		return m.RoomUser
	}
	return nil
}

func (m *UserPosition) GetNewPosition() *Vector3 {
	if m != nil {
		return m.NewPosition
	}
	return nil
}

type MasterSwitch struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MasterId             string   `protobuf:"bytes,2,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	NewMasterId          string   `protobuf:"bytes,3,opt,name=new_master_id,json=newMasterId,proto3" json:"new_master_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MasterSwitch) Reset()         { *m = MasterSwitch{} }
func (m *MasterSwitch) String() string { return proto.CompactTextString(m) }
func (*MasterSwitch) ProtoMessage()    {}
func (*MasterSwitch) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{6}
}

func (m *MasterSwitch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MasterSwitch.Unmarshal(m, b)
}
func (m *MasterSwitch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MasterSwitch.Marshal(b, m, deterministic)
}
func (m *MasterSwitch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MasterSwitch.Merge(m, src)
}
func (m *MasterSwitch) XXX_Size() int {
	return xxx_messageInfo_MasterSwitch.Size(m)
}
func (m *MasterSwitch) XXX_DiscardUnknown() {
	xxx_messageInfo_MasterSwitch.DiscardUnknown(m)
}

var xxx_messageInfo_MasterSwitch proto.InternalMessageInfo

func (m *MasterSwitch) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MasterSwitch) GetMasterId() string {
	if m != nil {
		return m.MasterId
	}
	return ""
}

func (m *MasterSwitch) GetNewMasterId() string {
	if m != nil {
		return m.NewMasterId
	}
	return ""
}

type Scene struct {
	RoomUser             *RoomUser `protobuf:"bytes,1,opt,name=room_user,json=roomUser,proto3" json:"room_user,omitempty"`
	SceneId              int32     `protobuf:"varint,2,opt,name=sceneId,proto3" json:"sceneId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Scene) Reset()         { *m = Scene{} }
func (m *Scene) String() string { return proto.CompactTextString(m) }
func (*Scene) ProtoMessage()    {}
func (*Scene) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{7}
}

func (m *Scene) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scene.Unmarshal(m, b)
}
func (m *Scene) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scene.Marshal(b, m, deterministic)
}
func (m *Scene) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scene.Merge(m, src)
}
func (m *Scene) XXX_Size() int {
	return xxx_messageInfo_Scene.Size(m)
}
func (m *Scene) XXX_DiscardUnknown() {
	xxx_messageInfo_Scene.DiscardUnknown(m)
}

var xxx_messageInfo_Scene proto.InternalMessageInfo

func (m *Scene) GetRoomUser() *RoomUser {
	if m != nil {
		return m.RoomUser
	}
	return nil
}

func (m *Scene) GetSceneId() int32 {
	if m != nil {
		return m.SceneId
	}
	return 0
}

type Update struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Room                 *Room    `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	Position             *Vector3 `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	IsMaster             bool     `protobuf:"varint,4,opt,name=is_master,json=isMaster,proto3" json:"is_master,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Update) Reset()         { *m = Update{} }
func (m *Update) String() string { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()    {}
func (*Update) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{8}
}

func (m *Update) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update.Unmarshal(m, b)
}
func (m *Update) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update.Marshal(b, m, deterministic)
}
func (m *Update) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update.Merge(m, src)
}
func (m *Update) XXX_Size() int {
	return xxx_messageInfo_Update.Size(m)
}
func (m *Update) XXX_DiscardUnknown() {
	xxx_messageInfo_Update.DiscardUnknown(m)
}

var xxx_messageInfo_Update proto.InternalMessageInfo

func (m *Update) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Update) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *Update) GetPosition() *Vector3 {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Update) GetIsMaster() bool {
	if m != nil {
		return m.IsMaster
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*Token)(nil), "Token")
	proto.RegisterType((*Vector3)(nil), "Vector3")
	proto.RegisterType((*Room)(nil), "Room")
	proto.RegisterType((*RoomUser)(nil), "RoomUser")
	proto.RegisterType((*UserPosition)(nil), "UserPosition")
	proto.RegisterType((*MasterSwitch)(nil), "MasterSwitch")
	proto.RegisterType((*Scene)(nil), "Scene")
	proto.RegisterType((*Update)(nil), "Update")
}

func init() { proto.RegisterFile("smartenergytable-service.proto", fileDescriptor_7be9bf9d675643a6) }

var fileDescriptor_7be9bf9d675643a6 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0x5e, 0xd2, 0x26, 0x71, 0x4e, 0xb2, 0xdf, 0x85, 0x6f, 0x96, 0x6d, 0xbf, 0xb1, 0x60, 0x4d,
	0x50, 0x09, 0x91, 0xc1, 0xfa, 0x04, 0x30, 0xed, 0xa2, 0x88, 0x49, 0xc8, 0xdd, 0xb8, 0xe0, 0xa6,
	0xca, 0x62, 0x6b, 0x33, 0x23, 0x71, 0x14, 0x87, 0x6e, 0xdd, 0x7b, 0xf0, 0x90, 0xbc, 0x05, 0xb2,
	0x9d, 0xac, 0x81, 0x09, 0xd4, 0xab, 0xd6, 0xe7, 0xcf, 0x77, 0xbe, 0xef, 0xeb, 0x39, 0x85, 0x67,
	0xaa, 0xcc, 0x9b, 0x96, 0x57, 0xbc, 0xb9, 0x5e, 0xb5, 0xf9, 0xd5, 0x37, 0xfe, 0x5a, 0xf1, 0x66,
	0x29, 0x0a, 0x9e, 0xd5, 0x8d, 0x6c, 0x25, 0x09, 0xc0, 0x3b, 0x2b, 0xeb, 0x76, 0x45, 0x7e, 0x38,
	0xe0, 0x5d, 0xc8, 0x5b, 0x5e, 0xe1, 0x17, 0x10, 0x36, 0x52, 0x96, 0x8b, 0xef, 0x8a, 0x37, 0x89,
	0x93, 0x3a, 0x93, 0xe8, 0x24, 0xcc, 0xa8, 0x94, 0xe5, 0xa5, 0xe2, 0x0d, 0x45, 0x4d, 0xf7, 0x0d,
	0x3f, 0x87, 0x58, 0x5e, 0x7d, 0xe5, 0x45, 0xbb, 0x10, 0x15, 0xe3, 0xf7, 0x89, 0x9b, 0x3a, 0x13,
	0x8f, 0x46, 0x36, 0x36, 0xd3, 0x21, 0x7c, 0x04, 0xa8, 0x96, 0x4a, 0xb4, 0x42, 0x56, 0xc9, 0xc8,
	0x20, 0xa1, 0xec, 0x33, 0x2f, 0x5a, 0xd9, 0x4c, 0xe9, 0x63, 0x06, 0xef, 0x43, 0xd8, 0x03, 0xb1,
	0x64, 0x9c, 0x3a, 0x93, 0x90, 0xa2, 0x0e, 0x85, 0x91, 0x29, 0x04, 0x5d, 0x07, 0x8e, 0xc1, 0xb9,
	0x37, 0x84, 0x5c, 0xea, 0xdc, 0xeb, 0xd7, 0xca, 0xcc, 0x74, 0xa9, 0xb3, 0xd2, 0xaf, 0x07, 0x33,
	0xc2, 0xa5, 0xce, 0x03, 0x29, 0x61, 0xac, 0x09, 0xe3, 0xff, 0xc0, 0x15, 0xcc, 0xb4, 0x84, 0xd4,
	0x15, 0x0c, 0xa7, 0x10, 0x58, 0x60, 0x95, 0xb8, 0xe9, 0x68, 0x12, 0x9d, 0xf8, 0x99, 0xd1, 0x4c,
	0xfb, 0x30, 0x4e, 0x20, 0x50, 0x05, 0xaf, 0xf8, 0x8c, 0x19, 0x34, 0x8f, 0xf6, 0x4f, 0xbc, 0x0b,
	0x48, 0x3b, 0xb2, 0x10, 0x4c, 0x25, 0xe3, 0x74, 0x34, 0x09, 0x69, 0xa0, 0xdf, 0x33, 0xa6, 0xc8,
	0x14, 0x50, 0xef, 0xcf, 0x93, 0x91, 0x3b, 0x10, 0x74, 0x6d, 0x86, 0x6c, 0x48, 0x7d, 0xdb, 0x45,
	0x0a, 0x88, 0x75, 0xc3, 0xa7, 0xde, 0x85, 0x4d, 0x6d, 0x7f, 0x05, 0x71, 0xc5, 0xef, 0x16, 0x8f,
	0xbe, 0xba, 0x7f, 0xf8, 0x1a, 0x55, 0xfc, 0xae, 0x07, 0x25, 0x0b, 0x88, 0xcf, 0x73, 0xd5, 0xf2,
	0x66, 0x7e, 0x27, 0xda, 0xe2, 0xe6, 0x09, 0xbb, 0x7d, 0x08, 0x4b, 0x93, 0x5f, 0xf3, 0x43, 0x36,
	0x30, 0x63, 0x98, 0xc0, 0xb6, 0x9e, 0xb4, 0x2e, 0x18, 0x99, 0x02, 0x3d, 0xe0, 0xbc, 0xab, 0x21,
	0x33, 0xf0, 0xe6, 0xda, 0xa0, 0x8d, 0xe9, 0x0f, 0x0c, 0x76, 0x7f, 0x33, 0x98, 0x2c, 0xc1, 0xbf,
	0xac, 0x59, 0xde, 0xf2, 0x27, 0x2c, 0x77, 0x61, 0xac, 0xfb, 0x3b, 0xa9, 0x9e, 0x81, 0xa5, 0x26,
	0xb4, 0xf9, 0x86, 0x09, 0xd5, 0x09, 0x31, 0x1b, 0x86, 0x28, 0x12, 0xca, 0x8a, 0x38, 0xf9, 0xe9,
	0xc2, 0xce, 0x5c, 0x5f, 0xc9, 0x99, 0xb9, 0x92, 0x0b, 0x7d, 0x25, 0x73, 0x7b, 0x24, 0xf8, 0x00,
	0xe0, 0xb4, 0xe1, 0x79, 0xcb, 0xcd, 0x3a, 0xf9, 0x99, 0xb9, 0x95, 0x3d, 0xcb, 0x80, 0x6c, 0x61,
	0x02, 0xe8, 0x83, 0x14, 0x95, 0x49, 0xae, 0xd5, 0xee, 0x05, 0x99, 0x15, 0x42, 0xb6, 0xde, 0x38,
	0x78, 0x1f, 0xd0, 0x3c, 0x5f, 0x5a, 0x00, 0xdb, 0xb8, 0xd7, 0xe1, 0x90, 0x2d, 0xfc, 0x3f, 0xa0,
	0x77, 0x8c, 0xd9, 0xbb, 0xeb, 0x76, 0x71, 0x90, 0x3d, 0x84, 0x88, 0xf2, 0x52, 0x2e, 0xf9, 0xdf,
	0x0a, 0x0e, 0x20, 0x3c, 0xff, 0x47, 0xfa, 0x10, 0xa2, 0xd3, 0x9b, 0xbc, 0xba, 0xe6, 0xf6, 0x27,
	0xf2, 0x33, 0xf3, 0x39, 0x28, 0x38, 0xb2, 0xfd, 0x9a, 0xb4, 0xc2, 0xdb, 0xd9, 0x70, 0x1f, 0x07,
	0x55, 0x29, 0x84, 0x1f, 0x79, 0x2f, 0x61, 0x20, 0x73, 0x5d, 0xf1, 0x12, 0x62, 0x3b, 0xc8, 0x5a,
	0x8a, 0xb7, 0xb3, 0xe1, 0xd6, 0xad, 0x0b, 0xdf, 0xc7, 0x5f, 0xa0, 0xbe, 0xbd, 0x3e, 0xce, 0x6b,
	0x71, 0xbc, 0x7c, 0x7b, 0xe5, 0x9b, 0xff, 0xa0, 0xe9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0xf9, 0x97, 0x8a, 0xa5, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SmartEnergyTableServiceClient is the client API for SmartEnergyTableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmartEnergyTableServiceClient interface {
	CreateRoom(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Room, error)
	JoinRoom(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (SmartEnergyTableService_JoinRoomClient, error)
	SaveRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (*Empty, error)
	AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error)
	RemoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error)
	MoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error)
	ChangeScene(ctx context.Context, in *Scene, opts ...grpc.CallOption) (*Empty, error)
	MoveUsers(ctx context.Context, in *UserPosition, opts ...grpc.CallOption) (*Empty, error)
	LeaveRoom(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (*Empty, error)
	ChangeMaster(ctx context.Context, in *MasterSwitch, opts ...grpc.CallOption) (*Empty, error)
}

type smartEnergyTableServiceClient struct {
	cc *grpc.ClientConn
}

func NewSmartEnergyTableServiceClient(cc *grpc.ClientConn) SmartEnergyTableServiceClient {
	return &smartEnergyTableServiceClient{cc}
}

func (c *smartEnergyTableServiceClient) CreateRoom(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/SmartEnergyTableService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartEnergyTableServiceClient) JoinRoom(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (SmartEnergyTableService_JoinRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SmartEnergyTableService_serviceDesc.Streams[0], "/SmartEnergyTableService/JoinRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartEnergyTableServiceJoinRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmartEnergyTableService_JoinRoomClient interface {
	Recv() (*Update, error)
	grpc.ClientStream
}

type smartEnergyTableServiceJoinRoomClient struct {
	grpc.ClientStream
}

func (x *smartEnergyTableServiceJoinRoomClient) Recv() (*Update, error) {
	m := new(Update)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smartEnergyTableServiceClient) SaveRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SmartEnergyTableService/SaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartEnergyTableServiceClient) AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SmartEnergyTableService/AddToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartEnergyTableServiceClient) RemoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SmartEnergyTableService/RemoveToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartEnergyTableServiceClient) MoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SmartEnergyTableService/MoveToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartEnergyTableServiceClient) ChangeScene(ctx context.Context, in *Scene, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SmartEnergyTableService/ChangeScene", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartEnergyTableServiceClient) MoveUsers(ctx context.Context, in *UserPosition, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SmartEnergyTableService/MoveUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartEnergyTableServiceClient) LeaveRoom(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SmartEnergyTableService/LeaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartEnergyTableServiceClient) ChangeMaster(ctx context.Context, in *MasterSwitch, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SmartEnergyTableService/ChangeMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartEnergyTableServiceServer is the server API for SmartEnergyTableService service.
type SmartEnergyTableServiceServer interface {
	CreateRoom(context.Context, *Empty) (*Room, error)
	JoinRoom(*RoomUser, SmartEnergyTableService_JoinRoomServer) error
	SaveRoom(context.Context, *Room) (*Empty, error)
	AddToken(context.Context, *Token) (*Empty, error)
	RemoveToken(context.Context, *Token) (*Empty, error)
	MoveToken(context.Context, *Token) (*Empty, error)
	ChangeScene(context.Context, *Scene) (*Empty, error)
	MoveUsers(context.Context, *UserPosition) (*Empty, error)
	LeaveRoom(context.Context, *RoomUser) (*Empty, error)
	ChangeMaster(context.Context, *MasterSwitch) (*Empty, error)
}

// UnimplementedSmartEnergyTableServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmartEnergyTableServiceServer struct {
}

func (*UnimplementedSmartEnergyTableServiceServer) CreateRoom(ctx context.Context, req *Empty) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) JoinRoom(req *RoomUser, srv SmartEnergyTableService_JoinRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) SaveRoom(ctx context.Context, req *Room) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveRoom not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) AddToken(ctx context.Context, req *Token) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToken not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) RemoveToken(ctx context.Context, req *Token) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveToken not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) MoveToken(ctx context.Context, req *Token) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveToken not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) ChangeScene(ctx context.Context, req *Scene) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeScene not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) MoveUsers(ctx context.Context, req *UserPosition) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveUsers not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) LeaveRoom(ctx context.Context, req *RoomUser) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) ChangeMaster(ctx context.Context, req *MasterSwitch) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeMaster not implemented")
}

func RegisterSmartEnergyTableServiceServer(s *grpc.Server, srv SmartEnergyTableServiceServer) {
	s.RegisterService(&_SmartEnergyTableService_serviceDesc, srv)
}

func _SmartEnergyTableService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartEnergyTableServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartEnergyTableService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartEnergyTableServiceServer).CreateRoom(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartEnergyTableService_JoinRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RoomUser)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmartEnergyTableServiceServer).JoinRoom(m, &smartEnergyTableServiceJoinRoomServer{stream})
}

type SmartEnergyTableService_JoinRoomServer interface {
	Send(*Update) error
	grpc.ServerStream
}

type smartEnergyTableServiceJoinRoomServer struct {
	grpc.ServerStream
}

func (x *smartEnergyTableServiceJoinRoomServer) Send(m *Update) error {
	return x.ServerStream.SendMsg(m)
}

func _SmartEnergyTableService_SaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Room)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartEnergyTableServiceServer).SaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartEnergyTableService/SaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartEnergyTableServiceServer).SaveRoom(ctx, req.(*Room))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartEnergyTableService_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartEnergyTableServiceServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartEnergyTableService/AddToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartEnergyTableServiceServer).AddToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartEnergyTableService_RemoveToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartEnergyTableServiceServer).RemoveToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartEnergyTableService/RemoveToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartEnergyTableServiceServer).RemoveToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartEnergyTableService_MoveToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartEnergyTableServiceServer).MoveToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartEnergyTableService/MoveToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartEnergyTableServiceServer).MoveToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartEnergyTableService_ChangeScene_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Scene)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartEnergyTableServiceServer).ChangeScene(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartEnergyTableService/ChangeScene",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartEnergyTableServiceServer).ChangeScene(ctx, req.(*Scene))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartEnergyTableService_MoveUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartEnergyTableServiceServer).MoveUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartEnergyTableService/MoveUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartEnergyTableServiceServer).MoveUsers(ctx, req.(*UserPosition))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartEnergyTableService_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartEnergyTableServiceServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartEnergyTableService/LeaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartEnergyTableServiceServer).LeaveRoom(ctx, req.(*RoomUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartEnergyTableService_ChangeMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterSwitch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartEnergyTableServiceServer).ChangeMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartEnergyTableService/ChangeMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartEnergyTableServiceServer).ChangeMaster(ctx, req.(*MasterSwitch))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmartEnergyTableService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SmartEnergyTableService",
	HandlerType: (*SmartEnergyTableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _SmartEnergyTableService_CreateRoom_Handler,
		},
		{
			MethodName: "SaveRoom",
			Handler:    _SmartEnergyTableService_SaveRoom_Handler,
		},
		{
			MethodName: "AddToken",
			Handler:    _SmartEnergyTableService_AddToken_Handler,
		},
		{
			MethodName: "RemoveToken",
			Handler:    _SmartEnergyTableService_RemoveToken_Handler,
		},
		{
			MethodName: "MoveToken",
			Handler:    _SmartEnergyTableService_MoveToken_Handler,
		},
		{
			MethodName: "ChangeScene",
			Handler:    _SmartEnergyTableService_ChangeScene_Handler,
		},
		{
			MethodName: "MoveUsers",
			Handler:    _SmartEnergyTableService_MoveUsers_Handler,
		},
		{
			MethodName: "LeaveRoom",
			Handler:    _SmartEnergyTableService_LeaveRoom_Handler,
		},
		{
			MethodName: "ChangeMaster",
			Handler:    _SmartEnergyTableService_ChangeMaster_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinRoom",
			Handler:       _SmartEnergyTableService_JoinRoom_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "smartenergytable-service.proto",
}
