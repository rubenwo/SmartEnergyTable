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

type Diff_Action int32

const (
	Diff_ADD    Diff_Action = 0
	Diff_DELETE Diff_Action = 1
	Diff_MOVE   Diff_Action = 2
)

var Diff_Action_name = map[int32]string{
	0: "ADD",
	1: "DELETE",
	2: "MOVE",
}

var Diff_Action_value = map[string]int32{
	"ADD":    0,
	"DELETE": 1,
	"MOVE":   2,
}

func (x Diff_Action) String() string {
	return proto.EnumName(Diff_Action_name, int32(x))
}

func (Diff_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{7, 0}
}

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
	Rotation             *Vector3  `protobuf:"bytes,5,opt,name=rotation,proto3" json:"rotation,omitempty"`
	Scale                float32   `protobuf:"fixed32,6,opt,name=scale,proto3" json:"scale,omitempty"`
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

func (m *Token) GetRotation() *Vector3 {
	if m != nil {
		return m.Rotation
	}
	return nil
}

func (m *Token) GetScale() float32 {
	if m != nil {
		return m.Scale
	}
	return 0
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
	return fileDescriptor_7be9bf9d675643a6, []int{3}
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
	return fileDescriptor_7be9bf9d675643a6, []int{4}
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
	return fileDescriptor_7be9bf9d675643a6, []int{5}
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
	return fileDescriptor_7be9bf9d675643a6, []int{6}
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

type Diff struct {
	Action               Diff_Action `protobuf:"varint,1,opt,name=action,proto3,enum=Diff_Action" json:"action,omitempty"`
	Token                *Token      `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Diff) Reset()         { *m = Diff{} }
func (m *Diff) String() string { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()    {}
func (*Diff) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{7}
}

func (m *Diff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Diff.Unmarshal(m, b)
}
func (m *Diff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Diff.Marshal(b, m, deterministic)
}
func (m *Diff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Diff.Merge(m, src)
}
func (m *Diff) XXX_Size() int {
	return xxx_messageInfo_Diff.Size(m)
}
func (m *Diff) XXX_DiscardUnknown() {
	xxx_messageInfo_Diff.DiscardUnknown(m)
}

var xxx_messageInfo_Diff proto.InternalMessageInfo

func (m *Diff) GetAction() Diff_Action {
	if m != nil {
		return m.Action
	}
	return Diff_ADD
}

func (m *Diff) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Patch struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	SceneId              int32    `protobuf:"varint,2,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	UserPosition         *Vector3 `protobuf:"bytes,3,opt,name=user_position,json=userPosition,proto3" json:"user_position,omitempty"`
	IsMaster             bool     `protobuf:"varint,4,opt,name=is_master,json=isMaster,proto3" json:"is_master,omitempty"`
	Diffs                []*Diff  `protobuf:"bytes,5,rep,name=diffs,proto3" json:"diffs,omitempty"`
	Objects              []*Token `protobuf:"bytes,6,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Patch) Reset()         { *m = Patch{} }
func (m *Patch) String() string { return proto.CompactTextString(m) }
func (*Patch) ProtoMessage()    {}
func (*Patch) Descriptor() ([]byte, []int) {
	return fileDescriptor_7be9bf9d675643a6, []int{8}
}

func (m *Patch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Patch.Unmarshal(m, b)
}
func (m *Patch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Patch.Marshal(b, m, deterministic)
}
func (m *Patch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Patch.Merge(m, src)
}
func (m *Patch) XXX_Size() int {
	return xxx_messageInfo_Patch.Size(m)
}
func (m *Patch) XXX_DiscardUnknown() {
	xxx_messageInfo_Patch.DiscardUnknown(m)
}

var xxx_messageInfo_Patch proto.InternalMessageInfo

func (m *Patch) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *Patch) GetSceneId() int32 {
	if m != nil {
		return m.SceneId
	}
	return 0
}

func (m *Patch) GetUserPosition() *Vector3 {
	if m != nil {
		return m.UserPosition
	}
	return nil
}

func (m *Patch) GetIsMaster() bool {
	if m != nil {
		return m.IsMaster
	}
	return false
}

func (m *Patch) GetDiffs() []*Diff {
	if m != nil {
		return m.Diffs
	}
	return nil
}

func (m *Patch) GetObjects() []*Token {
	if m != nil {
		return m.Objects
	}
	return nil
}

func init() {
	proto.RegisterEnum("Diff_Action", Diff_Action_name, Diff_Action_value)
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*Token)(nil), "Token")
	proto.RegisterType((*Vector3)(nil), "Vector3")
	proto.RegisterType((*RoomUser)(nil), "RoomUser")
	proto.RegisterType((*UserPosition)(nil), "UserPosition")
	proto.RegisterType((*MasterSwitch)(nil), "MasterSwitch")
	proto.RegisterType((*Scene)(nil), "Scene")
	proto.RegisterType((*Diff)(nil), "Diff")
	proto.RegisterType((*Patch)(nil), "Patch")
}

func init() { proto.RegisterFile("smartenergytable-service.proto", fileDescriptor_7be9bf9d675643a6) }

var fileDescriptor_7be9bf9d675643a6 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x4e, 0xdb, 0x30,
	0x14, 0x6e, 0x52, 0xf2, 0x77, 0x9a, 0x22, 0x64, 0x4d, 0x22, 0xa3, 0x6c, 0x04, 0x0b, 0x8d, 0x4a,
	0x13, 0x61, 0x6b, 0x9f, 0x80, 0xd1, 0x5e, 0x74, 0xa2, 0x1a, 0x72, 0x19, 0x17, 0xbb, 0xa9, 0x42,
	0xe2, 0x82, 0x07, 0x8d, 0xab, 0x38, 0x14, 0xca, 0x03, 0xee, 0x7e, 0x6f, 0xb2, 0x47, 0x98, 0x6c,
	0x27, 0xb4, 0x02, 0x81, 0xb8, 0x4a, 0xce, 0x39, 0x9f, 0x8f, 0xbf, 0xef, 0xfc, 0x18, 0x3e, 0x8a,
	0x69, 0x9c, 0x17, 0x34, 0xa3, 0xf9, 0xe5, 0xa2, 0x88, 0x2f, 0x6e, 0xe8, 0x81, 0xa0, 0xf9, 0x9c,
	0x25, 0x34, 0x9a, 0xe5, 0xbc, 0xe0, 0xd8, 0x01, 0xab, 0x3f, 0x9d, 0x15, 0x0b, 0xfc, 0xd7, 0x00,
	0xeb, 0x8c, 0x5f, 0xd3, 0x0c, 0x7d, 0x02, 0x2f, 0xe7, 0x7c, 0x3a, 0xbe, 0x15, 0x34, 0x0f, 0x8c,
	0xd0, 0x68, 0x37, 0x3a, 0x5e, 0x44, 0x38, 0x9f, 0xfe, 0x14, 0x34, 0x27, 0x6e, 0x5e, 0xfe, 0xa1,
	0x5d, 0xf0, 0xf9, 0xc5, 0x6f, 0x9a, 0x14, 0x63, 0x96, 0xa5, 0xf4, 0x3e, 0x30, 0x43, 0xa3, 0x6d,
	0x91, 0x86, 0xf6, 0x0d, 0xa4, 0x0b, 0xed, 0x81, 0x3b, 0xe3, 0x82, 0x15, 0x8c, 0x67, 0x41, 0x5d,
	0x65, 0x72, 0xa3, 0x73, 0x9a, 0x14, 0x3c, 0xef, 0x92, 0xc7, 0x08, 0x6a, 0x81, 0x57, 0x25, 0x4a,
	0x83, 0xb5, 0xd0, 0x68, 0x7b, 0xc4, 0x2d, 0xb3, 0xa4, 0x32, 0x45, 0xce, 0x8b, 0x58, 0xa5, 0xb0,
	0x9e, 0xa6, 0xa8, 0x22, 0xe8, 0x1d, 0x58, 0x22, 0x89, 0x6f, 0x68, 0x60, 0x87, 0x46, 0xdb, 0x24,
	0xda, 0xc0, 0x5d, 0x70, 0x4a, 0x28, 0xf2, 0xc1, 0xb8, 0x57, 0x62, 0x4c, 0x62, 0xdc, 0x4b, 0x6b,
	0xa1, 0xf8, 0x9a, 0xc4, 0x58, 0x48, 0xeb, 0x41, 0xd1, 0x33, 0x89, 0xf1, 0x80, 0xbb, 0xe0, 0x56,
	0x62, 0xd1, 0x3a, 0x98, 0x2c, 0x55, 0xc7, 0x3c, 0x62, 0xb2, 0x14, 0x6d, 0x82, 0x23, 0xab, 0x22,
	0x79, 0x9a, 0xca, 0x69, 0x4b, 0x73, 0x90, 0xe2, 0x04, 0x7c, 0x79, 0xe0, 0xb4, 0x92, 0xf4, 0xd6,
	0x1a, 0x7e, 0x06, 0x3f, 0xa3, 0x77, 0xe3, 0xc7, 0x22, 0x99, 0x4f, 0x14, 0x36, 0x32, 0x7a, 0x57,
	0x25, 0xc5, 0x63, 0xf0, 0x87, 0xb1, 0x28, 0x68, 0x3e, 0xba, 0x63, 0x45, 0x72, 0xf5, 0x8c, 0x5d,
	0x0b, 0xbc, 0xa9, 0x8a, 0x2f, 0xf9, 0xb9, 0xda, 0x31, 0x48, 0x11, 0x86, 0xa6, 0xbc, 0x69, 0x09,
	0xa8, 0x2b, 0x80, 0xbc, 0x60, 0x58, 0x62, 0xf0, 0x00, 0xac, 0x51, 0x42, 0x33, 0xfa, 0x66, 0xfa,
	0x01, 0x38, 0x42, 0x1e, 0x18, 0xa4, 0x65, 0xf7, 0x2b, 0x13, 0x0b, 0x58, 0xeb, 0xb1, 0xc9, 0x04,
	0xed, 0x81, 0x1d, 0x27, 0x4a, 0x9a, 0x4c, 0xb3, 0xde, 0xf1, 0x23, 0xe9, 0x8e, 0x8e, 0x94, 0x8f,
	0x94, 0x31, 0xb4, 0x0d, 0x56, 0x21, 0x67, 0xaf, 0xd4, 0x6f, 0x47, 0x6a, 0x12, 0x89, 0x76, 0xe2,
	0x7d, 0xb0, 0x35, 0x1e, 0x39, 0x50, 0x3f, 0xea, 0xf5, 0x36, 0x6a, 0x08, 0xc0, 0xee, 0xf5, 0x4f,
	0xfa, 0x67, 0xfd, 0x0d, 0x03, 0xb9, 0xb0, 0x36, 0xfc, 0x71, 0xde, 0xdf, 0x30, 0xf1, 0x1f, 0x03,
	0xac, 0xd3, 0x58, 0x96, 0x66, 0x13, 0x1c, 0x25, 0xe0, 0xb1, 0x3e, 0xb6, 0x34, 0x07, 0x29, 0x7a,
	0x0f, 0xae, 0xa2, 0x58, 0x95, 0x68, 0x49, 0x19, 0x1d, 0x40, 0x53, 0x35, 0xf7, 0xc5, 0x89, 0xf5,
	0x6f, 0x57, 0x5b, 0xdc, 0x02, 0x8f, 0x89, 0xb2, 0x9e, 0x6a, 0x6a, 0x5d, 0xe2, 0x32, 0xa1, 0x6b,
	0x89, 0x5a, 0x60, 0xa5, 0x6c, 0x32, 0x11, 0x81, 0x15, 0xd6, 0xdb, 0x8d, 0x8e, 0xa5, 0x54, 0x13,
	0xed, 0x43, 0x21, 0x38, 0x7a, 0xbc, 0x45, 0x60, 0xab, 0x70, 0xa5, 0xb7, 0x72, 0x77, 0xfe, 0x99,
	0xb0, 0x39, 0x92, 0x8b, 0xdb, 0x57, 0x8b, 0x7b, 0x26, 0x17, 0x77, 0xa4, 0xf7, 0x16, 0xed, 0x02,
	0x1c, 0xe7, 0x34, 0x2e, 0xa8, 0xec, 0x07, 0xb2, 0x23, 0xb5, 0xbe, 0x5b, 0xcb, 0xf6, 0xe0, 0x1a,
	0xda, 0x05, 0xf7, 0x3b, 0x67, 0x99, 0x02, 0x2c, 0x03, 0x5b, 0x76, 0xa4, 0x8a, 0x83, 0x6b, 0x5f,
	0x0c, 0xb4, 0x03, 0xee, 0x28, 0x9e, 0xd3, 0xe7, 0x10, 0xfd, 0x1a, 0xd4, 0xd0, 0x36, 0xb8, 0x47,
	0x69, 0xaa, 0x5f, 0x84, 0x92, 0xdf, 0x4a, 0x74, 0x07, 0x1a, 0x84, 0x4e, 0xf9, 0x9c, 0xbe, 0x04,
	0xf8, 0x00, 0xde, 0xf0, 0x95, 0xf0, 0x0e, 0x34, 0x8e, 0xaf, 0xe2, 0xec, 0x92, 0xea, 0x79, 0xb3,
	0x23, 0xf5, 0x5d, 0x01, 0xec, 0xe9, 0xf3, 0x92, 0x94, 0x40, 0xcd, 0x68, 0x75, 0xb9, 0x56, 0x50,
	0x21, 0x78, 0x27, 0xf4, 0x55, 0x19, 0xfb, 0xe0, 0xeb, 0x8b, 0xca, 0xc6, 0x34, 0xa3, 0xd5, 0x15,
	0x5a, 0x02, 0xbf, 0xf9, 0xbf, 0x60, 0x76, 0x7d, 0x79, 0x18, 0xcf, 0xd8, 0xe1, 0xfc, 0xeb, 0x85,
	0xad, 0x5e, 0xc7, 0xee, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x90, 0x3d, 0x12, 0x3f, 0x05,
	0x00, 0x00,
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
	CreateRoom(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RoomUser, error)
	JoinRoom(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (SmartEnergyTableService_JoinRoomClient, error)
	SaveRoom(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (*Empty, error)
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

func (c *smartEnergyTableServiceClient) CreateRoom(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RoomUser, error) {
	out := new(RoomUser)
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
	Recv() (*Patch, error)
	grpc.ClientStream
}

type smartEnergyTableServiceJoinRoomClient struct {
	grpc.ClientStream
}

func (x *smartEnergyTableServiceJoinRoomClient) Recv() (*Patch, error) {
	m := new(Patch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smartEnergyTableServiceClient) SaveRoom(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (*Empty, error) {
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
	CreateRoom(context.Context, *Empty) (*RoomUser, error)
	JoinRoom(*RoomUser, SmartEnergyTableService_JoinRoomServer) error
	SaveRoom(context.Context, *RoomUser) (*Empty, error)
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

func (*UnimplementedSmartEnergyTableServiceServer) CreateRoom(ctx context.Context, req *Empty) (*RoomUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) JoinRoom(req *RoomUser, srv SmartEnergyTableService_JoinRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (*UnimplementedSmartEnergyTableServiceServer) SaveRoom(ctx context.Context, req *RoomUser) (*Empty, error) {
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
	Send(*Patch) error
	grpc.ServerStream
}

type smartEnergyTableServiceJoinRoomServer struct {
	grpc.ServerStream
}

func (x *smartEnergyTableServiceJoinRoomServer) Send(m *Patch) error {
	return x.ServerStream.SendMsg(m)
}

func _SmartEnergyTableService_SaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomUser)
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
		return srv.(SmartEnergyTableServiceServer).SaveRoom(ctx, req.(*RoomUser))
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
