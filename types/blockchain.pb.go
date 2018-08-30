// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package types // import "github.com/aergoio/aergo/types"

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

type TxType int32

const (
	TxType_NORMAL     TxType = 0
	TxType_GOVERNANCE TxType = 1
)

var TxType_name = map[int32]string{
	0: "NORMAL",
	1: "GOVERNANCE",
}
var TxType_value = map[string]int32{
	"NORMAL":     0,
	"GOVERNANCE": 1,
}

func (x TxType) String() string {
	return proto.EnumName(TxType_name, int32(x))
}
func (TxType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{0}
}

type Block struct {
	Hash                 []byte       `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Header               *BlockHeader `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body                 *BlockBody   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetBody() *BlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type BlockHeader struct {
	PrevBlockHash        []byte   `protobuf:"bytes,1,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	BlockNo              uint64   `protobuf:"varint,2,opt,name=blockNo,proto3" json:"blockNo,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BlocksRootHash       []byte   `protobuf:"bytes,4,opt,name=blocksRootHash,proto3" json:"blocksRootHash,omitempty"`
	TxsRootHash          []byte   `protobuf:"bytes,5,opt,name=txsRootHash,proto3" json:"txsRootHash,omitempty"`
	Confirms             uint64   `protobuf:"varint,6,opt,name=confirms,proto3" json:"confirms,omitempty"`
	PubKey               []byte   `protobuf:"bytes,7,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Sign                 []byte   `protobuf:"bytes,8,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{1}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *BlockHeader) GetBlockNo() uint64 {
	if m != nil {
		return m.BlockNo
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetBlocksRootHash() []byte {
	if m != nil {
		return m.BlocksRootHash
	}
	return nil
}

func (m *BlockHeader) GetTxsRootHash() []byte {
	if m != nil {
		return m.TxsRootHash
	}
	return nil
}

func (m *BlockHeader) GetConfirms() uint64 {
	if m != nil {
		return m.Confirms
	}
	return 0
}

func (m *BlockHeader) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *BlockHeader) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type BlockBody struct {
	Txs                  []*Tx    `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockBody) Reset()         { *m = BlockBody{} }
func (m *BlockBody) String() string { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()    {}
func (*BlockBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{2}
}
func (m *BlockBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockBody.Unmarshal(m, b)
}
func (m *BlockBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockBody.Marshal(b, m, deterministic)
}
func (dst *BlockBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockBody.Merge(dst, src)
}
func (m *BlockBody) XXX_Size() int {
	return xxx_messageInfo_BlockBody.Size(m)
}
func (m *BlockBody) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockBody.DiscardUnknown(m)
}

var xxx_messageInfo_BlockBody proto.InternalMessageInfo

func (m *BlockBody) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

type TxList struct {
	Txs                  []*Tx    `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxList) Reset()         { *m = TxList{} }
func (m *TxList) String() string { return proto.CompactTextString(m) }
func (*TxList) ProtoMessage()    {}
func (*TxList) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{3}
}
func (m *TxList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxList.Unmarshal(m, b)
}
func (m *TxList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxList.Marshal(b, m, deterministic)
}
func (dst *TxList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxList.Merge(dst, src)
}
func (m *TxList) XXX_Size() int {
	return xxx_messageInfo_TxList.Size(m)
}
func (m *TxList) XXX_DiscardUnknown() {
	xxx_messageInfo_TxList.DiscardUnknown(m)
}

var xxx_messageInfo_TxList proto.InternalMessageInfo

func (m *TxList) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Tx struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Body                 *TxBody  `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{4}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tx.Unmarshal(m, b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
}
func (dst *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(dst, src)
}
func (m *Tx) XXX_Size() int {
	return xxx_messageInfo_Tx.Size(m)
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Tx) GetBody() *TxBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type TxBody struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Account              []byte   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Recipient            []byte   `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount               uint64   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Limit                uint64   `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Price                uint64   `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	Sign                 []byte   `protobuf:"bytes,8,opt,name=sign,proto3" json:"sign,omitempty"`
	Type                 TxType   `protobuf:"varint,9,opt,name=type,proto3,enum=types.TxType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxBody) Reset()         { *m = TxBody{} }
func (m *TxBody) String() string { return proto.CompactTextString(m) }
func (*TxBody) ProtoMessage()    {}
func (*TxBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{5}
}
func (m *TxBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxBody.Unmarshal(m, b)
}
func (m *TxBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxBody.Marshal(b, m, deterministic)
}
func (dst *TxBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxBody.Merge(dst, src)
}
func (m *TxBody) XXX_Size() int {
	return xxx_messageInfo_TxBody.Size(m)
}
func (m *TxBody) XXX_DiscardUnknown() {
	xxx_messageInfo_TxBody.DiscardUnknown(m)
}

var xxx_messageInfo_TxBody proto.InternalMessageInfo

func (m *TxBody) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TxBody) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *TxBody) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *TxBody) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxBody) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TxBody) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TxBody) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TxBody) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *TxBody) GetType() TxType {
	if m != nil {
		return m.Type
	}
	return TxType_NORMAL
}

type TxIdx struct {
	BlockHash            []byte   `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Idx                  int32    `protobuf:"varint,2,opt,name=idx,proto3" json:"idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxIdx) Reset()         { *m = TxIdx{} }
func (m *TxIdx) String() string { return proto.CompactTextString(m) }
func (*TxIdx) ProtoMessage()    {}
func (*TxIdx) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{6}
}
func (m *TxIdx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxIdx.Unmarshal(m, b)
}
func (m *TxIdx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxIdx.Marshal(b, m, deterministic)
}
func (dst *TxIdx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxIdx.Merge(dst, src)
}
func (m *TxIdx) XXX_Size() int {
	return xxx_messageInfo_TxIdx.Size(m)
}
func (m *TxIdx) XXX_DiscardUnknown() {
	xxx_messageInfo_TxIdx.DiscardUnknown(m)
}

var xxx_messageInfo_TxIdx proto.InternalMessageInfo

func (m *TxIdx) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *TxIdx) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

type TxInBlock struct {
	TxIdx                *TxIdx   `protobuf:"bytes,1,opt,name=txIdx,proto3" json:"txIdx,omitempty"`
	Tx                   *Tx      `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxInBlock) Reset()         { *m = TxInBlock{} }
func (m *TxInBlock) String() string { return proto.CompactTextString(m) }
func (*TxInBlock) ProtoMessage()    {}
func (*TxInBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{7}
}
func (m *TxInBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxInBlock.Unmarshal(m, b)
}
func (m *TxInBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxInBlock.Marshal(b, m, deterministic)
}
func (dst *TxInBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxInBlock.Merge(dst, src)
}
func (m *TxInBlock) XXX_Size() int {
	return xxx_messageInfo_TxInBlock.Size(m)
}
func (m *TxInBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_TxInBlock.DiscardUnknown(m)
}

var xxx_messageInfo_TxInBlock proto.InternalMessageInfo

func (m *TxInBlock) GetTxIdx() *TxIdx {
	if m != nil {
		return m.TxIdx
	}
	return nil
}

func (m *TxInBlock) GetTx() *Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

type State struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Balance              uint64   `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	CodeHash             []byte   `protobuf:"bytes,3,opt,name=codeHash,proto3" json:"codeHash,omitempty"`
	StorageRoot          []byte   `protobuf:"bytes,4,opt,name=storageRoot,proto3" json:"storageRoot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{8}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *State) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *State) GetCodeHash() []byte {
	if m != nil {
		return m.CodeHash
	}
	return nil
}

func (m *State) GetStorageRoot() []byte {
	if m != nil {
		return m.StorageRoot
	}
	return nil
}

type Receipt struct {
	ContractAddress      []byte   `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Ret                  string   `protobuf:"bytes,3,opt,name=ret,proto3" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{9}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receipt.Unmarshal(m, b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
}
func (dst *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(dst, src)
}
func (m *Receipt) XXX_Size() int {
	return xxx_messageInfo_Receipt.Size(m)
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetContractAddress() []byte {
	if m != nil {
		return m.ContractAddress
	}
	return nil
}

func (m *Receipt) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Receipt) GetRet() string {
	if m != nil {
		return m.Ret
	}
	return ""
}

type FnArgument struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FnArgument) Reset()         { *m = FnArgument{} }
func (m *FnArgument) String() string { return proto.CompactTextString(m) }
func (*FnArgument) ProtoMessage()    {}
func (*FnArgument) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{10}
}
func (m *FnArgument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FnArgument.Unmarshal(m, b)
}
func (m *FnArgument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FnArgument.Marshal(b, m, deterministic)
}
func (dst *FnArgument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FnArgument.Merge(dst, src)
}
func (m *FnArgument) XXX_Size() int {
	return xxx_messageInfo_FnArgument.Size(m)
}
func (m *FnArgument) XXX_DiscardUnknown() {
	xxx_messageInfo_FnArgument.DiscardUnknown(m)
}

var xxx_messageInfo_FnArgument proto.InternalMessageInfo

func (m *FnArgument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Function struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arguments            []*FnArgument `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{11}
}
func (m *Function) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Function.Unmarshal(m, b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Function.Marshal(b, m, deterministic)
}
func (dst *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(dst, src)
}
func (m *Function) XXX_Size() int {
	return xxx_messageInfo_Function.Size(m)
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetArguments() []*FnArgument {
	if m != nil {
		return m.Arguments
	}
	return nil
}

type ABI struct {
	Version              string      `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Language             string      `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Functions            []*Function `protobuf:"bytes,3,rep,name=functions,proto3" json:"functions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ABI) Reset()         { *m = ABI{} }
func (m *ABI) String() string { return proto.CompactTextString(m) }
func (*ABI) ProtoMessage()    {}
func (*ABI) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_66eb8125c1c706ae, []int{12}
}
func (m *ABI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ABI.Unmarshal(m, b)
}
func (m *ABI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ABI.Marshal(b, m, deterministic)
}
func (dst *ABI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABI.Merge(dst, src)
}
func (m *ABI) XXX_Size() int {
	return xxx_messageInfo_ABI.Size(m)
}
func (m *ABI) XXX_DiscardUnknown() {
	xxx_messageInfo_ABI.DiscardUnknown(m)
}

var xxx_messageInfo_ABI proto.InternalMessageInfo

func (m *ABI) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ABI) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *ABI) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "types.Block")
	proto.RegisterType((*BlockHeader)(nil), "types.BlockHeader")
	proto.RegisterType((*BlockBody)(nil), "types.BlockBody")
	proto.RegisterType((*TxList)(nil), "types.TxList")
	proto.RegisterType((*Tx)(nil), "types.Tx")
	proto.RegisterType((*TxBody)(nil), "types.TxBody")
	proto.RegisterType((*TxIdx)(nil), "types.TxIdx")
	proto.RegisterType((*TxInBlock)(nil), "types.TxInBlock")
	proto.RegisterType((*State)(nil), "types.State")
	proto.RegisterType((*Receipt)(nil), "types.Receipt")
	proto.RegisterType((*FnArgument)(nil), "types.FnArgument")
	proto.RegisterType((*Function)(nil), "types.Function")
	proto.RegisterType((*ABI)(nil), "types.ABI")
	proto.RegisterEnum("types.TxType", TxType_name, TxType_value)
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_blockchain_66eb8125c1c706ae) }

var fileDescriptor_blockchain_66eb8125c1c706ae = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x7d, 0x4e, 0xec, 0xb4, 0xbe, 0x49, 0xdb, 0xbc, 0xd1, 0xd3, 0x93, 0xdf, 0x87, 0x50, 0xb0,
	0x0a, 0x8a, 0x2a, 0x91, 0x4a, 0x65, 0xc1, 0x82, 0x55, 0x82, 0x5a, 0x08, 0x94, 0x54, 0x1a, 0x22,
	0x16, 0x48, 0x2c, 0x26, 0xf6, 0x34, 0x19, 0x88, 0x67, 0x2c, 0xcf, 0xb8, 0x72, 0x7e, 0x33, 0xbf,
	0x01, 0x09, 0xcd, 0xf5, 0x24, 0x0e, 0x55, 0x61, 0x95, 0x39, 0xe7, 0x7e, 0x9f, 0x7b, 0x1d, 0xe8,
	0x2f, 0xd6, 0x2a, 0xf9, 0x9a, 0xac, 0x98, 0x90, 0xa3, 0xbc, 0x50, 0x46, 0x91, 0xc0, 0x6c, 0x72,
	0xae, 0xe3, 0x0c, 0x82, 0x89, 0x35, 0x11, 0x02, 0xfe, 0x8a, 0xe9, 0x55, 0xe4, 0x0d, 0xbc, 0x61,
	0x8f, 0xe2, 0x9b, 0x9c, 0x41, 0x67, 0xc5, 0x59, 0xca, 0x8b, 0xa8, 0x35, 0xf0, 0x86, 0xdd, 0x0b,
	0x32, 0xc2, 0xa0, 0x11, 0x46, 0xbc, 0x41, 0x0b, 0x75, 0x1e, 0xe4, 0x14, 0xfc, 0x85, 0x4a, 0x37,
	0x51, 0x1b, 0x3d, 0xfb, 0xfb, 0x9e, 0x13, 0x95, 0x6e, 0x28, 0x5a, 0xe3, 0xef, 0x1e, 0x74, 0xf7,
	0xa2, 0xc9, 0x29, 0x1c, 0xe5, 0x05, 0xbf, 0xab, 0xa9, 0xa6, 0xfc, 0xcf, 0x24, 0x89, 0xe0, 0x00,
	0xfb, 0x9f, 0x29, 0x6c, 0xc4, 0xa7, 0x5b, 0x48, 0xfe, 0x87, 0xd0, 0x88, 0x8c, 0x6b, 0xc3, 0xb2,
	0x1c, 0x4b, 0xb7, 0x69, 0x43, 0x90, 0xa7, 0x70, 0x8c, 0x8e, 0x9a, 0x2a, 0x65, 0x30, 0xbd, 0x8f,
	0xe9, 0xef, 0xb1, 0x64, 0x00, 0x5d, 0x53, 0x35, 0x4e, 0x01, 0x3a, 0xed, 0x53, 0xe4, 0x5f, 0x38,
	0x4c, 0x94, 0xbc, 0x15, 0x45, 0xa6, 0xa3, 0x0e, 0xb6, 0xb0, 0xc3, 0xe4, 0x6f, 0xe8, 0xe4, 0xe5,
	0xe2, 0x1d, 0xdf, 0x44, 0x07, 0x18, 0xe8, 0x90, 0x55, 0x54, 0x8b, 0xa5, 0x8c, 0x0e, 0x6b, 0x45,
	0xed, 0x3b, 0x1e, 0x42, 0xb8, 0x93, 0x84, 0xfc, 0x07, 0x6d, 0x53, 0xe9, 0xc8, 0x1b, 0xb4, 0x87,
	0xdd, 0x8b, 0xd0, 0x29, 0x36, 0xaf, 0xa8, 0x65, 0xe3, 0x27, 0xd0, 0x99, 0x57, 0xd7, 0x42, 0x9b,
	0xdf, 0xbb, 0xbd, 0x84, 0xd6, 0xbc, 0x7a, 0x70, 0x79, 0x8f, 0xdd, 0x42, 0xea, 0xd5, 0x1d, 0xed,
	0xe2, 0xf6, 0xb6, 0xf1, 0xcd, 0xb3, 0x45, 0xb0, 0x97, 0xbf, 0x20, 0x90, 0x4a, 0x26, 0x1c, 0x53,
	0xf8, 0xb4, 0x06, 0x56, 0x78, 0x96, 0x24, 0xaa, 0x94, 0x06, 0xd3, 0xf4, 0xe8, 0x16, 0x5a, 0xe1,
	0x0b, 0x9e, 0x88, 0x5c, 0x70, 0x69, 0x50, 0xf8, 0x1e, 0x6d, 0x08, 0x2b, 0x09, 0xcb, 0x30, 0xcc,
	0xc7, 0x74, 0x0e, 0xd9, 0x7c, 0x39, 0xdb, 0xac, 0x15, 0x4b, 0x9d, 0xc8, 0x5b, 0x68, 0xeb, 0xaf,
	0x45, 0x26, 0x8c, 0x53, 0xb7, 0x06, 0x96, 0xcd, 0x0b, 0x91, 0x70, 0x54, 0xd6, 0xa7, 0x35, 0x78,
	0x48, 0x58, 0x3b, 0xad, 0x1d, 0x30, 0x0a, 0x07, 0xde, 0xf0, 0x78, 0x6f, 0xda, 0xf9, 0x26, 0xe7,
	0x14, 0x4d, 0xf1, 0x0b, 0x08, 0xe6, 0xd5, 0x34, 0xad, 0x6c, 0xef, 0x8b, 0x7b, 0x07, 0xd7, 0x10,
	0xa4, 0x0f, 0x6d, 0x91, 0x56, 0x38, 0x6f, 0x40, 0xed, 0x33, 0x7e, 0x0b, 0xe1, 0xbc, 0x9a, 0xca,
	0xfa, 0x3b, 0x89, 0x21, 0x30, 0x36, 0x0b, 0x06, 0x76, 0x2f, 0x7a, 0xbb, 0x4a, 0xd3, 0xb4, 0xa2,
	0xb5, 0x89, 0xfc, 0x03, 0x2d, 0x53, 0x39, 0xe1, 0xf7, 0x16, 0xd6, 0x32, 0x55, 0x5c, 0x42, 0xf0,
	0xc1, 0x30, 0xc3, 0x7f, 0x2d, 0xf8, 0x82, 0xad, 0x99, 0xe5, 0xb7, 0x97, 0x5e, 0xc3, 0xfa, 0x02,
	0x53, 0x8e, 0x3d, 0xd7, 0x7a, 0xef, 0xb0, 0xbd, 0x5f, 0x6d, 0x54, 0xc1, 0x96, 0xdc, 0x1e, 0xac,
	0x3b, 0xf2, 0x7d, 0x2a, 0xfe, 0x0c, 0x07, 0x94, 0x27, 0x5c, 0xe4, 0x86, 0x0c, 0xe1, 0x24, 0x51,
	0xd2, 0x14, 0x2c, 0x31, 0xe3, 0x34, 0x2d, 0xb8, 0xd6, 0x4e, 0x83, 0xfb, 0xb4, 0xdd, 0xa2, 0x36,
	0xcc, 0x94, 0x1a, 0x7b, 0x09, 0xa9, 0x43, 0x56, 0xa1, 0x82, 0xd7, 0x5b, 0x0f, 0xa9, 0x7d, 0xc6,
	0x03, 0x80, 0x2b, 0x39, 0x2e, 0x96, 0x65, 0x66, 0xb7, 0x4f, 0xc0, 0x97, 0x2c, 0xab, 0x27, 0x0b,
	0x29, 0xbe, 0xe3, 0x1b, 0x38, 0xbc, 0x2a, 0x65, 0x62, 0x84, 0x92, 0x0f, 0xd9, 0xc9, 0x39, 0x84,
	0xcc, 0xc5, 0xdb, 0x72, 0xf6, 0xd4, 0xff, 0x74, 0xca, 0x35, 0x99, 0x69, 0xe3, 0x13, 0x7f, 0x81,
	0xf6, 0x78, 0x32, 0xb5, 0x82, 0xdd, 0xf1, 0x42, 0x0b, 0x25, 0x5d, 0xba, 0x2d, 0xb4, 0x82, 0xad,
	0x99, 0x5c, 0x96, 0x6c, 0xc9, 0x5d, 0xff, 0x3b, 0x4c, 0x9e, 0x41, 0x78, 0xeb, 0xba, 0xd1, 0x51,
	0x1b, 0xab, 0x9d, 0x6c, 0xab, 0x39, 0x9e, 0x36, 0x1e, 0x67, 0xa7, 0xf6, 0x33, 0xb1, 0x97, 0x44,
	0x00, 0x3a, 0xb3, 0x1b, 0xfa, 0x7e, 0x7c, 0xdd, 0xff, 0x83, 0x1c, 0x03, 0xbc, 0xbe, 0xf9, 0x78,
	0x49, 0x67, 0xe3, 0xd9, 0xab, 0xcb, 0xbe, 0x37, 0x19, 0x7c, 0x7a, 0xb4, 0x14, 0x66, 0x55, 0x2e,
	0x46, 0x89, 0xca, 0xce, 0x19, 0x2f, 0x96, 0x4a, 0xa8, 0xfa, 0xf7, 0x1c, 0x73, 0x2f, 0x3a, 0xf8,
	0xd7, 0xfb, 0xfc, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x6f, 0xb7, 0xfb, 0x8e, 0x05, 0x00,
	0x00,
}
