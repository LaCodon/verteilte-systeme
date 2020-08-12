// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

package rpc

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

type VoteRequest struct {
	Term                 int32    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId          uint32   `protobuf:"varint,2,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
	LastLogIndex         int32    `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm          int32    `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteRequest) Reset()         { *m = VoteRequest{} }
func (m *VoteRequest) String() string { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()    {}
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{0}
}

func (m *VoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteRequest.Unmarshal(m, b)
}
func (m *VoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteRequest.Marshal(b, m, deterministic)
}
func (m *VoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRequest.Merge(m, src)
}
func (m *VoteRequest) XXX_Size() int {
	return xxx_messageInfo_VoteRequest.Size(m)
}
func (m *VoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRequest proto.InternalMessageInfo

func (m *VoteRequest) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *VoteRequest) GetCandidateId() uint32 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *VoteRequest) GetLastLogIndex() int32 {
	if m != nil {
		return m.LastLogIndex
	}
	return 0
}

func (m *VoteRequest) GetLastLogTerm() int32 {
	if m != nil {
		return m.LastLogTerm
	}
	return 0
}

type VoteResponse struct {
	Term                 int32    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted          bool     `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteResponse) Reset()         { *m = VoteResponse{} }
func (m *VoteResponse) String() string { return proto.CompactTextString(m) }
func (*VoteResponse) ProtoMessage()    {}
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{1}
}

func (m *VoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteResponse.Unmarshal(m, b)
}
func (m *VoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteResponse.Marshal(b, m, deterministic)
}
func (m *VoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteResponse.Merge(m, src)
}
func (m *VoteResponse) XXX_Size() int {
	return xxx_messageInfo_VoteResponse.Size(m)
}
func (m *VoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VoteResponse proto.InternalMessageInfo

func (m *VoteResponse) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *VoteResponse) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

type AppendEntriesRequest struct {
	Term                 int32       `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId             uint32      `protobuf:"varint,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	PrevLogIndex         int32       `protobuf:"varint,3,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	PrevLogTerm          int32       `protobuf:"varint,4,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	Entries              []*LogEntry `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	LeaderCommit         int32       `protobuf:"varint,6,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
	AllNodes             []string    `protobuf:"bytes,7,rep,name=allNodes,proto3" json:"allNodes,omitempty"`
	LeaderTarget         string      `protobuf:"bytes,8,opt,name=leaderTarget,proto3" json:"leaderTarget,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppendEntriesRequest) Reset()         { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()    {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{2}
}

func (m *AppendEntriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesRequest.Unmarshal(m, b)
}
func (m *AppendEntriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesRequest.Marshal(b, m, deterministic)
}
func (m *AppendEntriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesRequest.Merge(m, src)
}
func (m *AppendEntriesRequest) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesRequest.Size(m)
}
func (m *AppendEntriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesRequest proto.InternalMessageInfo

func (m *AppendEntriesRequest) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderId() uint32 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() int32 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() int32 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AppendEntriesRequest) GetLeaderCommit() int32 {
	if m != nil {
		return m.LeaderCommit
	}
	return 0
}

func (m *AppendEntriesRequest) GetAllNodes() []string {
	if m != nil {
		return m.AllNodes
	}
	return nil
}

func (m *AppendEntriesRequest) GetLeaderTarget() string {
	if m != nil {
		return m.LeaderTarget
	}
	return ""
}

type AppendEntriesResponse struct {
	Term                 int32    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEntriesResponse) Reset()         { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()    {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{3}
}

func (m *AppendEntriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesResponse.Unmarshal(m, b)
}
func (m *AppendEntriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesResponse.Marshal(b, m, deterministic)
}
func (m *AppendEntriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesResponse.Merge(m, src)
}
func (m *AppendEntriesResponse) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesResponse.Size(m)
}
func (m *AppendEntriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesResponse proto.InternalMessageInfo

func (m *AppendEntriesResponse) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type LogEntry struct {
	Index                int32    `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Term                 int32    `protobuf:"varint,2,opt,name=Term,proto3" json:"Term,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Action               int32    `protobuf:"varint,4,opt,name=Action,proto3" json:"Action,omitempty"`
	Value                string   `protobuf:"bytes,5,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{4}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *LogEntry) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *LogEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LogEntry) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *LogEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type NodeRegisterRequest struct {
	NodeId               uint32   `protobuf:"varint,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	ConnectionData       string   `protobuf:"bytes,2,opt,name=connectionData,proto3" json:"connectionData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRegisterRequest) Reset()         { *m = NodeRegisterRequest{} }
func (m *NodeRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*NodeRegisterRequest) ProtoMessage()    {}
func (*NodeRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{5}
}

func (m *NodeRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRegisterRequest.Unmarshal(m, b)
}
func (m *NodeRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRegisterRequest.Marshal(b, m, deterministic)
}
func (m *NodeRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRegisterRequest.Merge(m, src)
}
func (m *NodeRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_NodeRegisterRequest.Size(m)
}
func (m *NodeRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRegisterRequest proto.InternalMessageInfo

func (m *NodeRegisterRequest) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeRegisterRequest) GetConnectionData() string {
	if m != nil {
		return m.ConnectionData
	}
	return ""
}

type NodeRegisterResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	RedirectTarget       string   `protobuf:"bytes,2,opt,name=redirectTarget,proto3" json:"redirectTarget,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRegisterResponse) Reset()         { *m = NodeRegisterResponse{} }
func (m *NodeRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*NodeRegisterResponse) ProtoMessage()    {}
func (*NodeRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{6}
}

func (m *NodeRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRegisterResponse.Unmarshal(m, b)
}
func (m *NodeRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRegisterResponse.Marshal(b, m, deterministic)
}
func (m *NodeRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRegisterResponse.Merge(m, src)
}
func (m *NodeRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_NodeRegisterResponse.Size(m)
}
func (m *NodeRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRegisterResponse proto.InternalMessageInfo

func (m *NodeRegisterResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *NodeRegisterResponse) GetRedirectTarget() string {
	if m != nil {
		return m.RedirectTarget
	}
	return ""
}

type UserRequest struct {
	RequestCode          int32    `protobuf:"varint,1,opt,name=requestCode,proto3" json:"requestCode,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{7}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetRequestCode() int32 {
	if m != nil {
		return m.RequestCode
	}
	return 0
}

func (m *UserRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UserRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UserResponse struct {
	ResponseCode         int32             `protobuf:"varint,1,opt,name=responseCode,proto3" json:"responseCode,omitempty"`
	RedirectTo           string            `protobuf:"bytes,2,opt,name=redirectTo,proto3" json:"redirectTo,omitempty"`
	Data                 map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{8}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetResponseCode() int32 {
	if m != nil {
		return m.ResponseCode
	}
	return 0
}

func (m *UserResponse) GetRedirectTo() string {
	if m != nil {
		return m.RedirectTo
	}
	return ""
}

func (m *UserResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*VoteRequest)(nil), "smkvs.VoteRequest")
	proto.RegisterType((*VoteResponse)(nil), "smkvs.VoteResponse")
	proto.RegisterType((*AppendEntriesRequest)(nil), "smkvs.AppendEntriesRequest")
	proto.RegisterType((*AppendEntriesResponse)(nil), "smkvs.AppendEntriesResponse")
	proto.RegisterType((*LogEntry)(nil), "smkvs.LogEntry")
	proto.RegisterType((*NodeRegisterRequest)(nil), "smkvs.NodeRegisterRequest")
	proto.RegisterType((*NodeRegisterResponse)(nil), "smkvs.NodeRegisterResponse")
	proto.RegisterType((*UserRequest)(nil), "smkvs.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "smkvs.UserResponse")
	proto.RegisterMapType((map[string]string)(nil), "smkvs.UserResponse.DataEntry")
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor_c33392ef2c1961ba) }

var fileDescriptor_c33392ef2c1961ba = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcd, 0x6e, 0xd4, 0x3e,
	0x14, 0xc5, 0x9b, 0xcc, 0x67, 0x6e, 0xa6, 0xff, 0xfe, 0xe5, 0x96, 0x2a, 0x4a, 0x01, 0x45, 0x5e,
	0xa0, 0x61, 0x33, 0x15, 0x65, 0x41, 0x85, 0xd8, 0x94, 0xb6, 0x42, 0x23, 0x2a, 0x16, 0x51, 0x5b,
	0x10, 0xbb, 0x90, 0x5c, 0x8d, 0xa2, 0xce, 0xc4, 0xc1, 0x76, 0x47, 0xf4, 0x11, 0x78, 0x02, 0x16,
	0xbc, 0x0b, 0xcf, 0x86, 0xfc, 0x91, 0xd4, 0xa9, 0x86, 0xee, 0x7c, 0x4f, 0xec, 0xe3, 0x73, 0x7f,
	0xb9, 0x09, 0xec, 0xd6, 0x9c, 0x49, 0x76, 0x28, 0x90, 0xaf, 0xcb, 0x1c, 0x67, 0xba, 0x22, 0x03,
	0xb1, 0xba, 0x59, 0x0b, 0xfa, 0xd3, 0x83, 0xf0, 0x9a, 0x49, 0x4c, 0xf1, 0xfb, 0x2d, 0x0a, 0x49,
	0x08, 0xf4, 0x25, 0xf2, 0x55, 0xe4, 0x25, 0xde, 0x74, 0x90, 0xea, 0x35, 0x49, 0x20, 0xcc, 0xb3,
	0xaa, 0x28, 0x8b, 0x4c, 0xe2, 0xbc, 0x88, 0xfc, 0xc4, 0x9b, 0x6e, 0xa7, 0xae, 0x44, 0x28, 0x4c,
	0x96, 0x99, 0x90, 0x17, 0x6c, 0x31, 0xaf, 0x0a, 0xfc, 0x11, 0xf5, 0xf4, 0xe9, 0x8e, 0xa6, 0x5c,
	0x6c, 0x7d, 0xa9, 0x2e, 0xe8, 0xeb, 0x2d, 0xae, 0x44, 0xcf, 0x60, 0x62, 0xa2, 0x88, 0x9a, 0x55,
	0x02, 0xff, 0x95, 0x65, 0xcd, 0x24, 0x7e, 0xe0, 0x59, 0x25, 0xd1, 0x64, 0x19, 0xa7, 0xae, 0x44,
	0x7f, 0xfb, 0xb0, 0x77, 0x52, 0xd7, 0x58, 0x15, 0xe7, 0x95, 0xe4, 0x25, 0x8a, 0xc7, 0x5a, 0x8b,
	0x61, 0xbc, 0xc4, 0xac, 0x40, 0xde, 0xf6, 0xd5, 0xd6, 0xaa, 0xa9, 0x9a, 0xe3, 0xfa, 0x61, 0x53,
	0xae, 0xa6, 0xe2, 0xd8, 0xda, 0x6d, 0xca, 0x91, 0xc8, 0x4b, 0x18, 0xa1, 0xc9, 0x11, 0x0d, 0x92,
	0xde, 0x34, 0x3c, 0xda, 0x99, 0x69, 0xf2, 0xb3, 0x0b, 0xb6, 0x50, 0x01, 0xef, 0xd2, 0xe6, 0xb9,
	0xa6, 0xa8, 0x2f, 0x3f, 0x65, 0xab, 0x55, 0x29, 0xa3, 0xa1, 0xa5, 0xe8, 0x68, 0x2a, 0x70, 0xb6,
	0x5c, 0x7e, 0x62, 0x05, 0x8a, 0x68, 0x94, 0xf4, 0xa6, 0x41, 0xda, 0xd6, 0xf7, 0xe7, 0x2f, 0x33,
	0xbe, 0x40, 0x19, 0x8d, 0x13, 0x6f, 0x1a, 0xa4, 0x1d, 0x8d, 0x9e, 0xc3, 0x93, 0x07, 0x70, 0x1e,
	0x81, 0x1d, 0xc1, 0x48, 0xdc, 0xe6, 0x39, 0x0a, 0x61, 0x41, 0x37, 0x25, 0x95, 0x30, 0x6e, 0xf2,
	0x93, 0x3d, 0x18, 0x18, 0x40, 0xe6, 0xa8, 0x29, 0x94, 0x9f, 0x46, 0xe2, 0x1b, 0x3f, 0xcd, 0xe2,
	0x7f, 0xe8, 0x7d, 0xc4, 0x3b, 0x0d, 0x32, 0x48, 0xd5, 0x92, 0xec, 0xc3, 0xf0, 0x24, 0x97, 0x25,
	0xab, 0x2c, 0x3a, 0x5b, 0x29, 0xcf, 0xeb, 0x6c, 0x79, 0x8b, 0xd1, 0x40, 0xef, 0x35, 0x05, 0xbd,
	0x82, 0x5d, 0xd5, 0x69, 0x8a, 0x8b, 0x52, 0x48, 0xe4, 0xcd, 0x8b, 0xdd, 0x87, 0x61, 0xc5, 0x0a,
	0x35, 0x9a, 0x9e, 0x7e, 0x85, 0xb6, 0x22, 0x2f, 0xe0, 0xbf, 0x9c, 0x55, 0x15, 0x6a, 0xcb, 0xb3,
	0x4c, 0x66, 0x3a, 0x4c, 0x90, 0x3e, 0x50, 0xe9, 0x17, 0xd8, 0xeb, 0xda, 0x5a, 0x24, 0x4e, 0xfb,
	0x5e, 0xa7, 0x7d, 0xe5, 0xcc, 0xb1, 0x28, 0x39, 0xe6, 0xd2, 0xb2, 0xb6, 0xce, 0x5d, 0x95, 0x7e,
	0x86, 0xf0, 0x4a, 0xdc, 0x07, 0x4d, 0x20, 0xe4, 0x66, 0x79, 0xca, 0x0a, 0xb4, 0xbc, 0x5c, 0x49,
	0x11, 0xba, 0xc1, 0x3b, 0xeb, 0xa6, 0x96, 0x8a, 0xc4, 0x5a, 0x93, 0x30, 0xd4, 0x4c, 0x41, 0xff,
	0x78, 0x30, 0x31, 0xce, 0x36, 0x2b, 0x85, 0x09, 0xb7, 0x6b, 0xc7, 0xbb, 0xa3, 0x91, 0xe7, 0x00,
	0x6d, 0x3e, 0x66, 0xef, 0x70, 0x14, 0xf2, 0x0a, 0xfa, 0x85, 0xa2, 0xd4, 0xd3, 0x73, 0xfa, 0xcc,
	0xce, 0xa9, 0x7b, 0xcd, 0x4c, 0xf1, 0x32, 0x53, 0xab, 0xb7, 0xc6, 0x6f, 0x20, 0x68, 0xa5, 0x26,
	0xbc, 0xb7, 0x21, 0xbc, 0xef, 0x84, 0x7f, 0xeb, 0x1f, 0x7b, 0x47, 0xbf, 0x7c, 0xe8, 0x2b, 0xe8,
	0xe4, 0x18, 0x42, 0x8b, 0x47, 0x7d, 0xfb, 0x84, 0xd8, 0x5b, 0x9d, 0x7f, 0x52, 0xbc, 0xdb, 0xd1,
	0x4c, 0x12, 0xba, 0x45, 0x2e, 0x60, 0xbb, 0x33, 0xca, 0xe4, 0xc0, 0xee, 0xdb, 0xf4, 0xf5, 0xc7,
	0x4f, 0x37, 0x3f, 0x6c, 0xdd, 0xe6, 0x30, 0x69, 0x06, 0x40, 0xe7, 0x8a, 0xed, 0xfe, 0x0d, 0x03,
	0x17, 0x1f, 0x6c, 0x7c, 0xd6, 0x5a, 0xbd, 0x83, 0x1d, 0x05, 0x6d, 0x5e, 0x49, 0xe4, 0x99, 0x99,
	0x67, 0xd2, 0x81, 0xd9, 0x6d, 0xcb, 0x05, 0x4c, 0xb7, 0xde, 0x07, 0x5f, 0x47, 0xf5, 0xcd, 0xe2,
	0x90, 0xd7, 0xf9, 0xb7, 0xa1, 0xfe, 0x55, 0xbf, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x16,
	0x1c, 0x55, 0xc1, 0x05, 0x00, 0x00,
}
