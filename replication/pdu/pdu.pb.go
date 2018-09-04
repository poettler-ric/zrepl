// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pdu.proto

package pdu

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

type FilesystemVersion_VersionType int32

const (
	FilesystemVersion_Snapshot FilesystemVersion_VersionType = 0
	FilesystemVersion_Bookmark FilesystemVersion_VersionType = 1
)

var FilesystemVersion_VersionType_name = map[int32]string{
	0: "Snapshot",
	1: "Bookmark",
}
var FilesystemVersion_VersionType_value = map[string]int32{
	"Snapshot": 0,
	"Bookmark": 1,
}

func (x FilesystemVersion_VersionType) String() string {
	return proto.EnumName(FilesystemVersion_VersionType_name, int32(x))
}
func (FilesystemVersion_VersionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{5, 0}
}

type SnapshotReplicationStatusReq_Op int32

const (
	SnapshotReplicationStatusReq_Get           SnapshotReplicationStatusReq_Op = 0
	SnapshotReplicationStatusReq_SetReplicated SnapshotReplicationStatusReq_Op = 1
)

var SnapshotReplicationStatusReq_Op_name = map[int32]string{
	0: "Get",
	1: "SetReplicated",
}
var SnapshotReplicationStatusReq_Op_value = map[string]int32{
	"Get":           0,
	"SetReplicated": 1,
}

func (x SnapshotReplicationStatusReq_Op) String() string {
	return proto.EnumName(SnapshotReplicationStatusReq_Op_name, int32(x))
}
func (SnapshotReplicationStatusReq_Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{14, 0}
}

type SnapshotReplicationStatusRes_Status int32

const (
	SnapshotReplicationStatusRes_Nonexistent   SnapshotReplicationStatusRes_Status = 0
	SnapshotReplicationStatusRes_NotReplicated SnapshotReplicationStatusRes_Status = 1
	SnapshotReplicationStatusRes_Replicated    SnapshotReplicationStatusRes_Status = 2
)

var SnapshotReplicationStatusRes_Status_name = map[int32]string{
	0: "Nonexistent",
	1: "NotReplicated",
	2: "Replicated",
}
var SnapshotReplicationStatusRes_Status_value = map[string]int32{
	"Nonexistent":   0,
	"NotReplicated": 1,
	"Replicated":    2,
}

func (x SnapshotReplicationStatusRes_Status) String() string {
	return proto.EnumName(SnapshotReplicationStatusRes_Status_name, int32(x))
}
func (SnapshotReplicationStatusRes_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{15, 0}
}

type ListFilesystemReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFilesystemReq) Reset()         { *m = ListFilesystemReq{} }
func (m *ListFilesystemReq) String() string { return proto.CompactTextString(m) }
func (*ListFilesystemReq) ProtoMessage()    {}
func (*ListFilesystemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{0}
}
func (m *ListFilesystemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilesystemReq.Unmarshal(m, b)
}
func (m *ListFilesystemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilesystemReq.Marshal(b, m, deterministic)
}
func (dst *ListFilesystemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilesystemReq.Merge(dst, src)
}
func (m *ListFilesystemReq) XXX_Size() int {
	return xxx_messageInfo_ListFilesystemReq.Size(m)
}
func (m *ListFilesystemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilesystemReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilesystemReq proto.InternalMessageInfo

type ListFilesystemRes struct {
	Filesystems          []*Filesystem `protobuf:"bytes,1,rep,name=Filesystems,proto3" json:"Filesystems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListFilesystemRes) Reset()         { *m = ListFilesystemRes{} }
func (m *ListFilesystemRes) String() string { return proto.CompactTextString(m) }
func (*ListFilesystemRes) ProtoMessage()    {}
func (*ListFilesystemRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{1}
}
func (m *ListFilesystemRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilesystemRes.Unmarshal(m, b)
}
func (m *ListFilesystemRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilesystemRes.Marshal(b, m, deterministic)
}
func (dst *ListFilesystemRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilesystemRes.Merge(dst, src)
}
func (m *ListFilesystemRes) XXX_Size() int {
	return xxx_messageInfo_ListFilesystemRes.Size(m)
}
func (m *ListFilesystemRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilesystemRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilesystemRes proto.InternalMessageInfo

func (m *ListFilesystemRes) GetFilesystems() []*Filesystem {
	if m != nil {
		return m.Filesystems
	}
	return nil
}

type Filesystem struct {
	Path                 string   `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	ResumeToken          string   `protobuf:"bytes,2,opt,name=ResumeToken,proto3" json:"ResumeToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filesystem) Reset()         { *m = Filesystem{} }
func (m *Filesystem) String() string { return proto.CompactTextString(m) }
func (*Filesystem) ProtoMessage()    {}
func (*Filesystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{2}
}
func (m *Filesystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filesystem.Unmarshal(m, b)
}
func (m *Filesystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filesystem.Marshal(b, m, deterministic)
}
func (dst *Filesystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filesystem.Merge(dst, src)
}
func (m *Filesystem) XXX_Size() int {
	return xxx_messageInfo_Filesystem.Size(m)
}
func (m *Filesystem) XXX_DiscardUnknown() {
	xxx_messageInfo_Filesystem.DiscardUnknown(m)
}

var xxx_messageInfo_Filesystem proto.InternalMessageInfo

func (m *Filesystem) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Filesystem) GetResumeToken() string {
	if m != nil {
		return m.ResumeToken
	}
	return ""
}

type ListFilesystemVersionsReq struct {
	Filesystem           string   `protobuf:"bytes,1,opt,name=Filesystem,proto3" json:"Filesystem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFilesystemVersionsReq) Reset()         { *m = ListFilesystemVersionsReq{} }
func (m *ListFilesystemVersionsReq) String() string { return proto.CompactTextString(m) }
func (*ListFilesystemVersionsReq) ProtoMessage()    {}
func (*ListFilesystemVersionsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{3}
}
func (m *ListFilesystemVersionsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilesystemVersionsReq.Unmarshal(m, b)
}
func (m *ListFilesystemVersionsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilesystemVersionsReq.Marshal(b, m, deterministic)
}
func (dst *ListFilesystemVersionsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilesystemVersionsReq.Merge(dst, src)
}
func (m *ListFilesystemVersionsReq) XXX_Size() int {
	return xxx_messageInfo_ListFilesystemVersionsReq.Size(m)
}
func (m *ListFilesystemVersionsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilesystemVersionsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilesystemVersionsReq proto.InternalMessageInfo

func (m *ListFilesystemVersionsReq) GetFilesystem() string {
	if m != nil {
		return m.Filesystem
	}
	return ""
}

type ListFilesystemVersionsRes struct {
	Versions             []*FilesystemVersion `protobuf:"bytes,1,rep,name=Versions,proto3" json:"Versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListFilesystemVersionsRes) Reset()         { *m = ListFilesystemVersionsRes{} }
func (m *ListFilesystemVersionsRes) String() string { return proto.CompactTextString(m) }
func (*ListFilesystemVersionsRes) ProtoMessage()    {}
func (*ListFilesystemVersionsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{4}
}
func (m *ListFilesystemVersionsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilesystemVersionsRes.Unmarshal(m, b)
}
func (m *ListFilesystemVersionsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilesystemVersionsRes.Marshal(b, m, deterministic)
}
func (dst *ListFilesystemVersionsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilesystemVersionsRes.Merge(dst, src)
}
func (m *ListFilesystemVersionsRes) XXX_Size() int {
	return xxx_messageInfo_ListFilesystemVersionsRes.Size(m)
}
func (m *ListFilesystemVersionsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilesystemVersionsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilesystemVersionsRes proto.InternalMessageInfo

func (m *ListFilesystemVersionsRes) GetVersions() []*FilesystemVersion {
	if m != nil {
		return m.Versions
	}
	return nil
}

type FilesystemVersion struct {
	Type                 FilesystemVersion_VersionType `protobuf:"varint,1,opt,name=Type,proto3,enum=pdu.FilesystemVersion_VersionType" json:"Type,omitempty"`
	Name                 string                        `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Guid                 uint64                        `protobuf:"varint,3,opt,name=Guid,proto3" json:"Guid,omitempty"`
	CreateTXG            uint64                        `protobuf:"varint,4,opt,name=CreateTXG,proto3" json:"CreateTXG,omitempty"`
	Creation             string                        `protobuf:"bytes,5,opt,name=Creation,proto3" json:"Creation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FilesystemVersion) Reset()         { *m = FilesystemVersion{} }
func (m *FilesystemVersion) String() string { return proto.CompactTextString(m) }
func (*FilesystemVersion) ProtoMessage()    {}
func (*FilesystemVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{5}
}
func (m *FilesystemVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilesystemVersion.Unmarshal(m, b)
}
func (m *FilesystemVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilesystemVersion.Marshal(b, m, deterministic)
}
func (dst *FilesystemVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilesystemVersion.Merge(dst, src)
}
func (m *FilesystemVersion) XXX_Size() int {
	return xxx_messageInfo_FilesystemVersion.Size(m)
}
func (m *FilesystemVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_FilesystemVersion.DiscardUnknown(m)
}

var xxx_messageInfo_FilesystemVersion proto.InternalMessageInfo

func (m *FilesystemVersion) GetType() FilesystemVersion_VersionType {
	if m != nil {
		return m.Type
	}
	return FilesystemVersion_Snapshot
}

func (m *FilesystemVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FilesystemVersion) GetGuid() uint64 {
	if m != nil {
		return m.Guid
	}
	return 0
}

func (m *FilesystemVersion) GetCreateTXG() uint64 {
	if m != nil {
		return m.CreateTXG
	}
	return 0
}

func (m *FilesystemVersion) GetCreation() string {
	if m != nil {
		return m.Creation
	}
	return ""
}

type SendReq struct {
	Filesystem string `protobuf:"bytes,1,opt,name=Filesystem,proto3" json:"Filesystem,omitempty"`
	From       string `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	// May be empty / null to request a full transfer of From
	To string `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`
	// If ResumeToken is not empty, the resume token that CAN be tried for 'zfs send' by the sender.
	// The sender MUST indicate in SendRes.UsedResumeToken
	// If it does not work, the sender SHOULD clear the resume token on their side
	// and use From and To instead
	// If ResumeToken is not empty, the GUIDs of From and To
	// MUST correspond to those encoded in the ResumeToken.
	// Otherwise, the Sender MUST return an error.
	ResumeToken          string   `protobuf:"bytes,4,opt,name=ResumeToken,proto3" json:"ResumeToken,omitempty"`
	Compress             bool     `protobuf:"varint,5,opt,name=Compress,proto3" json:"Compress,omitempty"`
	Dedup                bool     `protobuf:"varint,6,opt,name=Dedup,proto3" json:"Dedup,omitempty"`
	DryRun               bool     `protobuf:"varint,7,opt,name=DryRun,proto3" json:"DryRun,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendReq) Reset()         { *m = SendReq{} }
func (m *SendReq) String() string { return proto.CompactTextString(m) }
func (*SendReq) ProtoMessage()    {}
func (*SendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{6}
}
func (m *SendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendReq.Unmarshal(m, b)
}
func (m *SendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendReq.Marshal(b, m, deterministic)
}
func (dst *SendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendReq.Merge(dst, src)
}
func (m *SendReq) XXX_Size() int {
	return xxx_messageInfo_SendReq.Size(m)
}
func (m *SendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendReq proto.InternalMessageInfo

func (m *SendReq) GetFilesystem() string {
	if m != nil {
		return m.Filesystem
	}
	return ""
}

func (m *SendReq) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendReq) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendReq) GetResumeToken() string {
	if m != nil {
		return m.ResumeToken
	}
	return ""
}

func (m *SendReq) GetCompress() bool {
	if m != nil {
		return m.Compress
	}
	return false
}

func (m *SendReq) GetDedup() bool {
	if m != nil {
		return m.Dedup
	}
	return false
}

func (m *SendReq) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

type Property struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{7}
}
func (m *Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Property.Unmarshal(m, b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Property.Marshal(b, m, deterministic)
}
func (dst *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(dst, src)
}
func (m *Property) XXX_Size() int {
	return xxx_messageInfo_Property.Size(m)
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

func (m *Property) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Property) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SendRes struct {
	// Whether the resume token provided in the request has been used or not.
	UsedResumeToken bool `protobuf:"varint,1,opt,name=UsedResumeToken,proto3" json:"UsedResumeToken,omitempty"`
	// Expected stream size determined by dry run, not exact
	ExpectedSize         int64       `protobuf:"varint,2,opt,name=ExpectedSize,proto3" json:"ExpectedSize,omitempty"`
	Properties           []*Property `protobuf:"bytes,3,rep,name=Properties,proto3" json:"Properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SendRes) Reset()         { *m = SendRes{} }
func (m *SendRes) String() string { return proto.CompactTextString(m) }
func (*SendRes) ProtoMessage()    {}
func (*SendRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{8}
}
func (m *SendRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRes.Unmarshal(m, b)
}
func (m *SendRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRes.Marshal(b, m, deterministic)
}
func (dst *SendRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRes.Merge(dst, src)
}
func (m *SendRes) XXX_Size() int {
	return xxx_messageInfo_SendRes.Size(m)
}
func (m *SendRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRes.DiscardUnknown(m)
}

var xxx_messageInfo_SendRes proto.InternalMessageInfo

func (m *SendRes) GetUsedResumeToken() bool {
	if m != nil {
		return m.UsedResumeToken
	}
	return false
}

func (m *SendRes) GetExpectedSize() int64 {
	if m != nil {
		return m.ExpectedSize
	}
	return 0
}

func (m *SendRes) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

type ReceiveReq struct {
	Filesystem string `protobuf:"bytes,1,opt,name=Filesystem,proto3" json:"Filesystem,omitempty"`
	// If true, the receiver should clear the resume token before perfoming the zfs recv of the stream in the request
	ClearResumeToken     bool     `protobuf:"varint,2,opt,name=ClearResumeToken,proto3" json:"ClearResumeToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveReq) Reset()         { *m = ReceiveReq{} }
func (m *ReceiveReq) String() string { return proto.CompactTextString(m) }
func (*ReceiveReq) ProtoMessage()    {}
func (*ReceiveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{9}
}
func (m *ReceiveReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveReq.Unmarshal(m, b)
}
func (m *ReceiveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveReq.Marshal(b, m, deterministic)
}
func (dst *ReceiveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveReq.Merge(dst, src)
}
func (m *ReceiveReq) XXX_Size() int {
	return xxx_messageInfo_ReceiveReq.Size(m)
}
func (m *ReceiveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveReq proto.InternalMessageInfo

func (m *ReceiveReq) GetFilesystem() string {
	if m != nil {
		return m.Filesystem
	}
	return ""
}

func (m *ReceiveReq) GetClearResumeToken() bool {
	if m != nil {
		return m.ClearResumeToken
	}
	return false
}

type ReceiveRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveRes) Reset()         { *m = ReceiveRes{} }
func (m *ReceiveRes) String() string { return proto.CompactTextString(m) }
func (*ReceiveRes) ProtoMessage()    {}
func (*ReceiveRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{10}
}
func (m *ReceiveRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveRes.Unmarshal(m, b)
}
func (m *ReceiveRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveRes.Marshal(b, m, deterministic)
}
func (dst *ReceiveRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveRes.Merge(dst, src)
}
func (m *ReceiveRes) XXX_Size() int {
	return xxx_messageInfo_ReceiveRes.Size(m)
}
func (m *ReceiveRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveRes proto.InternalMessageInfo

type DestroySnapshotsReq struct {
	Filesystem string `protobuf:"bytes,1,opt,name=Filesystem,proto3" json:"Filesystem,omitempty"`
	// Path to filesystem, snapshot or bookmark to be destroyed
	Snapshots            []*FilesystemVersion `protobuf:"bytes,2,rep,name=Snapshots,proto3" json:"Snapshots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DestroySnapshotsReq) Reset()         { *m = DestroySnapshotsReq{} }
func (m *DestroySnapshotsReq) String() string { return proto.CompactTextString(m) }
func (*DestroySnapshotsReq) ProtoMessage()    {}
func (*DestroySnapshotsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{11}
}
func (m *DestroySnapshotsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroySnapshotsReq.Unmarshal(m, b)
}
func (m *DestroySnapshotsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroySnapshotsReq.Marshal(b, m, deterministic)
}
func (dst *DestroySnapshotsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroySnapshotsReq.Merge(dst, src)
}
func (m *DestroySnapshotsReq) XXX_Size() int {
	return xxx_messageInfo_DestroySnapshotsReq.Size(m)
}
func (m *DestroySnapshotsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroySnapshotsReq.DiscardUnknown(m)
}

var xxx_messageInfo_DestroySnapshotsReq proto.InternalMessageInfo

func (m *DestroySnapshotsReq) GetFilesystem() string {
	if m != nil {
		return m.Filesystem
	}
	return ""
}

func (m *DestroySnapshotsReq) GetSnapshots() []*FilesystemVersion {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type DestroySnapshotRes struct {
	Snapshot             *FilesystemVersion `protobuf:"bytes,1,opt,name=Snapshot,proto3" json:"Snapshot,omitempty"`
	Error                string             `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DestroySnapshotRes) Reset()         { *m = DestroySnapshotRes{} }
func (m *DestroySnapshotRes) String() string { return proto.CompactTextString(m) }
func (*DestroySnapshotRes) ProtoMessage()    {}
func (*DestroySnapshotRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{12}
}
func (m *DestroySnapshotRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroySnapshotRes.Unmarshal(m, b)
}
func (m *DestroySnapshotRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroySnapshotRes.Marshal(b, m, deterministic)
}
func (dst *DestroySnapshotRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroySnapshotRes.Merge(dst, src)
}
func (m *DestroySnapshotRes) XXX_Size() int {
	return xxx_messageInfo_DestroySnapshotRes.Size(m)
}
func (m *DestroySnapshotRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroySnapshotRes.DiscardUnknown(m)
}

var xxx_messageInfo_DestroySnapshotRes proto.InternalMessageInfo

func (m *DestroySnapshotRes) GetSnapshot() *FilesystemVersion {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

func (m *DestroySnapshotRes) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DestroySnapshotsRes struct {
	Results              []*DestroySnapshotRes `protobuf:"bytes,1,rep,name=Results,proto3" json:"Results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DestroySnapshotsRes) Reset()         { *m = DestroySnapshotsRes{} }
func (m *DestroySnapshotsRes) String() string { return proto.CompactTextString(m) }
func (*DestroySnapshotsRes) ProtoMessage()    {}
func (*DestroySnapshotsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{13}
}
func (m *DestroySnapshotsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroySnapshotsRes.Unmarshal(m, b)
}
func (m *DestroySnapshotsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroySnapshotsRes.Marshal(b, m, deterministic)
}
func (dst *DestroySnapshotsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroySnapshotsRes.Merge(dst, src)
}
func (m *DestroySnapshotsRes) XXX_Size() int {
	return xxx_messageInfo_DestroySnapshotsRes.Size(m)
}
func (m *DestroySnapshotsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroySnapshotsRes.DiscardUnknown(m)
}

var xxx_messageInfo_DestroySnapshotsRes proto.InternalMessageInfo

func (m *DestroySnapshotsRes) GetResults() []*DestroySnapshotRes {
	if m != nil {
		return m.Results
	}
	return nil
}

type SnapshotReplicationStatusReq struct {
	Filesystem           string                          `protobuf:"bytes,1,opt,name=Filesystem,proto3" json:"Filesystem,omitempty"`
	Snapshot             string                          `protobuf:"bytes,2,opt,name=Snapshot,proto3" json:"Snapshot,omitempty"`
	Op                   SnapshotReplicationStatusReq_Op `protobuf:"varint,3,opt,name=op,proto3,enum=pdu.SnapshotReplicationStatusReq_Op" json:"op,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SnapshotReplicationStatusReq) Reset()         { *m = SnapshotReplicationStatusReq{} }
func (m *SnapshotReplicationStatusReq) String() string { return proto.CompactTextString(m) }
func (*SnapshotReplicationStatusReq) ProtoMessage()    {}
func (*SnapshotReplicationStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{14}
}
func (m *SnapshotReplicationStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotReplicationStatusReq.Unmarshal(m, b)
}
func (m *SnapshotReplicationStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotReplicationStatusReq.Marshal(b, m, deterministic)
}
func (dst *SnapshotReplicationStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotReplicationStatusReq.Merge(dst, src)
}
func (m *SnapshotReplicationStatusReq) XXX_Size() int {
	return xxx_messageInfo_SnapshotReplicationStatusReq.Size(m)
}
func (m *SnapshotReplicationStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotReplicationStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotReplicationStatusReq proto.InternalMessageInfo

func (m *SnapshotReplicationStatusReq) GetFilesystem() string {
	if m != nil {
		return m.Filesystem
	}
	return ""
}

func (m *SnapshotReplicationStatusReq) GetSnapshot() string {
	if m != nil {
		return m.Snapshot
	}
	return ""
}

func (m *SnapshotReplicationStatusReq) GetOp() SnapshotReplicationStatusReq_Op {
	if m != nil {
		return m.Op
	}
	return SnapshotReplicationStatusReq_Get
}

type SnapshotReplicationStatusRes struct {
	Status               SnapshotReplicationStatusRes_Status `protobuf:"varint,1,opt,name=status,proto3,enum=pdu.SnapshotReplicationStatusRes_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *SnapshotReplicationStatusRes) Reset()         { *m = SnapshotReplicationStatusRes{} }
func (m *SnapshotReplicationStatusRes) String() string { return proto.CompactTextString(m) }
func (*SnapshotReplicationStatusRes) ProtoMessage()    {}
func (*SnapshotReplicationStatusRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_pdu_b3a98b3542e9fb4e, []int{15}
}
func (m *SnapshotReplicationStatusRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotReplicationStatusRes.Unmarshal(m, b)
}
func (m *SnapshotReplicationStatusRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotReplicationStatusRes.Marshal(b, m, deterministic)
}
func (dst *SnapshotReplicationStatusRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotReplicationStatusRes.Merge(dst, src)
}
func (m *SnapshotReplicationStatusRes) XXX_Size() int {
	return xxx_messageInfo_SnapshotReplicationStatusRes.Size(m)
}
func (m *SnapshotReplicationStatusRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotReplicationStatusRes.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotReplicationStatusRes proto.InternalMessageInfo

func (m *SnapshotReplicationStatusRes) GetStatus() SnapshotReplicationStatusRes_Status {
	if m != nil {
		return m.Status
	}
	return SnapshotReplicationStatusRes_Nonexistent
}

func init() {
	proto.RegisterType((*ListFilesystemReq)(nil), "pdu.ListFilesystemReq")
	proto.RegisterType((*ListFilesystemRes)(nil), "pdu.ListFilesystemRes")
	proto.RegisterType((*Filesystem)(nil), "pdu.Filesystem")
	proto.RegisterType((*ListFilesystemVersionsReq)(nil), "pdu.ListFilesystemVersionsReq")
	proto.RegisterType((*ListFilesystemVersionsRes)(nil), "pdu.ListFilesystemVersionsRes")
	proto.RegisterType((*FilesystemVersion)(nil), "pdu.FilesystemVersion")
	proto.RegisterType((*SendReq)(nil), "pdu.SendReq")
	proto.RegisterType((*Property)(nil), "pdu.Property")
	proto.RegisterType((*SendRes)(nil), "pdu.SendRes")
	proto.RegisterType((*ReceiveReq)(nil), "pdu.ReceiveReq")
	proto.RegisterType((*ReceiveRes)(nil), "pdu.ReceiveRes")
	proto.RegisterType((*DestroySnapshotsReq)(nil), "pdu.DestroySnapshotsReq")
	proto.RegisterType((*DestroySnapshotRes)(nil), "pdu.DestroySnapshotRes")
	proto.RegisterType((*DestroySnapshotsRes)(nil), "pdu.DestroySnapshotsRes")
	proto.RegisterType((*SnapshotReplicationStatusReq)(nil), "pdu.SnapshotReplicationStatusReq")
	proto.RegisterType((*SnapshotReplicationStatusRes)(nil), "pdu.SnapshotReplicationStatusRes")
	proto.RegisterEnum("pdu.FilesystemVersion_VersionType", FilesystemVersion_VersionType_name, FilesystemVersion_VersionType_value)
	proto.RegisterEnum("pdu.SnapshotReplicationStatusReq_Op", SnapshotReplicationStatusReq_Op_name, SnapshotReplicationStatusReq_Op_value)
	proto.RegisterEnum("pdu.SnapshotReplicationStatusRes_Status", SnapshotReplicationStatusRes_Status_name, SnapshotReplicationStatusRes_Status_value)
}

func init() { proto.RegisterFile("pdu.proto", fileDescriptor_pdu_b3a98b3542e9fb4e) }

var fileDescriptor_pdu_b3a98b3542e9fb4e = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0xd9, 0x49, 0xf3, 0x67, 0xd2, 0xa6, 0xe9, 0xb6, 0xea, 0xcf, 0x54, 0x15, 0x8a, 0x56,
	0x1c, 0x02, 0x12, 0x91, 0x08, 0x15, 0x17, 0x38, 0xa0, 0xfe, 0xe5, 0x80, 0xda, 0x6a, 0x13, 0xaa,
	0x9e, 0x90, 0x4c, 0x3d, 0x52, 0xad, 0x24, 0xde, 0xed, 0xee, 0x1a, 0x35, 0x3c, 0x00, 0x8f, 0xc1,
	0x43, 0x70, 0xe3, 0x4d, 0x78, 0x1c, 0xe4, 0x89, 0xed, 0xb8, 0x49, 0x09, 0x39, 0x65, 0xbe, 0x6f,
	0x66, 0x67, 0xbe, 0x99, 0xdd, 0x71, 0xa0, 0xae, 0x82, 0xb8, 0xab, 0xb4, 0xb4, 0x92, 0x95, 0x54,
	0x10, 0xf3, 0x6d, 0xd8, 0xfa, 0x18, 0x1a, 0x7b, 0x1a, 0x8e, 0xd0, 0x4c, 0x8c, 0xc5, 0xb1, 0xc0,
	0x3b, 0x7e, 0xba, 0x48, 0x1a, 0xf6, 0x0a, 0x1a, 0x33, 0xc2, 0x78, 0x4e, 0xbb, 0xd4, 0x69, 0xf4,
	0x36, 0xbb, 0x49, 0xbe, 0x42, 0x60, 0x31, 0x86, 0x1f, 0x02, 0xcc, 0x20, 0x63, 0x50, 0xbe, 0xf4,
	0xed, 0xad, 0xe7, 0xb4, 0x9d, 0x4e, 0x5d, 0x90, 0xcd, 0xda, 0xd0, 0x10, 0x68, 0xe2, 0x31, 0x0e,
	0xe4, 0x10, 0x23, 0xcf, 0x25, 0x57, 0x91, 0xe2, 0x6f, 0xe1, 0xc9, 0x43, 0x2d, 0x57, 0xa8, 0x4d,
	0x28, 0x23, 0x23, 0xf0, 0x8e, 0x3d, 0x2d, 0x16, 0x48, 0x13, 0x17, 0x18, 0x7e, 0xf1, 0xf7, 0xc3,
	0x86, 0xf5, 0xa0, 0x96, 0xc1, 0xb4, 0x9b, 0xdd, 0xb9, 0x6e, 0x52, 0xb7, 0xc8, 0xe3, 0xf8, 0x6f,
	0x07, 0xb6, 0x16, 0xfc, 0xec, 0x0d, 0x94, 0x07, 0x13, 0x85, 0x24, 0xa0, 0xd9, 0xe3, 0x8f, 0x67,
	0xe9, 0xa6, 0xbf, 0x49, 0xa4, 0xa0, 0xf8, 0x64, 0x22, 0xe7, 0xfe, 0x18, 0xd3, 0xb6, 0xc9, 0x4e,
	0xb8, 0xb3, 0x38, 0x0c, 0xbc, 0x52, 0xdb, 0xe9, 0x94, 0x05, 0xd9, 0x6c, 0x1f, 0xea, 0x47, 0x1a,
	0x7d, 0x8b, 0x83, 0xeb, 0x33, 0xaf, 0x4c, 0x8e, 0x19, 0xc1, 0xf6, 0xa0, 0x46, 0x20, 0x94, 0x91,
	0xb7, 0x46, 0x99, 0x72, 0xcc, 0x9f, 0x43, 0xa3, 0x50, 0x96, 0xad, 0x43, 0xad, 0x1f, 0xf9, 0xca,
	0xdc, 0x4a, 0xdb, 0xfa, 0x2f, 0x41, 0x87, 0x52, 0x0e, 0xc7, 0xbe, 0x1e, 0xb6, 0x1c, 0xfe, 0xcb,
	0x81, 0x6a, 0x1f, 0xa3, 0x60, 0x85, 0xb9, 0x26, 0x22, 0x4f, 0xb5, 0x1c, 0x67, 0xc2, 0x13, 0x9b,
	0x35, 0xc1, 0x1d, 0x48, 0x92, 0x5d, 0x17, 0xee, 0x40, 0xce, 0x5f, 0x6d, 0x79, 0xe1, 0x6a, 0x49,
	0xb8, 0x1c, 0x2b, 0x8d, 0xc6, 0x90, 0xf0, 0x9a, 0xc8, 0x31, 0xdb, 0x81, 0xb5, 0x63, 0x0c, 0x62,
	0xe5, 0x55, 0xc8, 0x31, 0x05, 0x6c, 0x17, 0x2a, 0xc7, 0x7a, 0x22, 0xe2, 0xc8, 0xab, 0x12, 0x9d,
	0x22, 0x7e, 0x00, 0xb5, 0x4b, 0x2d, 0x15, 0x6a, 0x3b, 0xc9, 0x87, 0xea, 0x14, 0x86, 0xba, 0x03,
	0x6b, 0x57, 0xfe, 0x28, 0xce, 0x26, 0x3d, 0x05, 0xfc, 0x7b, 0xde, 0xb1, 0x61, 0x1d, 0xd8, 0xfc,
	0x64, 0x30, 0x28, 0x2a, 0x76, 0xa8, 0xc4, 0x3c, 0xcd, 0x38, 0xac, 0x9f, 0xdc, 0x2b, 0xbc, 0xb1,
	0x18, 0xf4, 0xc3, 0x6f, 0xd3, 0x94, 0x25, 0xf1, 0x80, 0x63, 0x2f, 0x01, 0x52, 0x3d, 0x21, 0x1a,
	0xaf, 0x44, 0x8f, 0x6b, 0x83, 0x9e, 0x45, 0x26, 0x53, 0x14, 0x02, 0xf8, 0x35, 0x80, 0xc0, 0x1b,
	0x0c, 0xbf, 0xe2, 0x2a, 0xc3, 0x7f, 0x01, 0xad, 0xa3, 0x11, 0xfa, 0x7a, 0x7e, 0x71, 0x6a, 0x62,
	0x81, 0xe7, 0xeb, 0x85, 0xcc, 0x86, 0x0f, 0x61, 0xfb, 0x18, 0x8d, 0xd5, 0x72, 0x92, 0xbd, 0x82,
	0x55, 0xb6, 0x88, 0x1d, 0x40, 0x3d, 0x8f, 0xf7, 0xdc, 0xa5, 0x9b, 0x32, 0x0b, 0xe4, 0x9f, 0x81,
	0xcd, 0x15, 0x4b, 0x97, 0x2e, 0x83, 0x54, 0x69, 0xc9, 0xd2, 0x65, 0x71, 0xc9, 0xed, 0x9d, 0x68,
	0x2d, 0x75, 0x76, 0x7b, 0x04, 0xf8, 0x87, 0xc7, 0x9a, 0x49, 0x3e, 0x53, 0xd5, 0x64, 0x00, 0x23,
	0x9b, 0x2d, 0xf5, 0xff, 0x94, 0x7f, 0x51, 0x8a, 0xc8, 0xe2, 0xf8, 0x4f, 0x07, 0xf6, 0x67, 0x0e,
	0x35, 0x0a, 0x6f, 0x68, 0x79, 0xfa, 0xd6, 0xb7, 0xf1, 0x4a, 0x03, 0xda, 0x2b, 0x34, 0x35, 0xd5,
	0x38, 0x13, 0x7f, 0x00, 0xae, 0x54, 0xb4, 0x16, 0xcd, 0xde, 0x33, 0x92, 0xb2, 0xac, 0x54, 0xf7,
	0x42, 0x09, 0x57, 0x2a, 0xde, 0x06, 0xf7, 0x42, 0xb1, 0x2a, 0x94, 0xce, 0x30, 0xd9, 0xd4, 0x2d,
	0xd8, 0xe8, 0x63, 0x7e, 0x00, 0x83, 0x96, 0xc3, 0x7f, 0x2c, 0x17, 0x6d, 0xd8, 0x7b, 0xa8, 0x18,
	0x02, 0xe9, 0x67, 0xa9, 0xf3, 0xaf, 0xe2, 0xa6, 0x9b, 0x5a, 0xe9, 0x39, 0xfe, 0x0e, 0x2a, 0x53,
	0x86, 0x6d, 0x42, 0xe3, 0x5c, 0x46, 0x78, 0x1f, 0x1a, 0x8b, 0x51, 0x2a, 0xe8, 0x5c, 0x3e, 0x10,
	0xc4, 0x9a, 0xc9, 0x53, 0xcb, 0xb1, 0xfb, 0xa5, 0x42, 0xff, 0x32, 0xaf, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x3a, 0x00, 0x8b, 0x6e, 0x72, 0x06, 0x00, 0x00,
}
