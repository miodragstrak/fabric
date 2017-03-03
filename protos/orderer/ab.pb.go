// Code generated by protoc-gen-go.
// source: orderer/ab.proto
// DO NOT EDIT!

/*
Package orderer is a generated protocol buffer package.

It is generated from these files:
	orderer/ab.proto
	orderer/configuration.proto
	orderer/kafka.proto

It has these top-level messages:
	BroadcastResponse
	SeekNewest
	SeekOldest
	SeekSpecified
	SeekPosition
	SeekInfo
	DeliverResponse
	ConsensusType
	BatchSize
	BatchTimeout
	CreationPolicy
	ChainCreationPolicyNames
	KafkaBrokers
	KafkaMessage
	KafkaMessageRegular
	KafkaMessageTimeToCut
	KafkaMessageConnect
	KafkaMetadata
*/
package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/hyperledger/fabric/protos/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type SeekInfo_SeekBehavior int32

const (
	SeekInfo_BLOCK_UNTIL_READY SeekInfo_SeekBehavior = 0
	SeekInfo_FAIL_IF_NOT_READY SeekInfo_SeekBehavior = 1
)

var SeekInfo_SeekBehavior_name = map[int32]string{
	0: "BLOCK_UNTIL_READY",
	1: "FAIL_IF_NOT_READY",
}
var SeekInfo_SeekBehavior_value = map[string]int32{
	"BLOCK_UNTIL_READY": 0,
	"FAIL_IF_NOT_READY": 1,
}

func (x SeekInfo_SeekBehavior) String() string {
	return proto.EnumName(SeekInfo_SeekBehavior_name, int32(x))
}
func (SeekInfo_SeekBehavior) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type BroadcastResponse struct {
	Status common.Status `protobuf:"varint,1,opt,name=status,enum=common.Status" json:"status,omitempty"`
}

func (m *BroadcastResponse) Reset()                    { *m = BroadcastResponse{} }
func (m *BroadcastResponse) String() string            { return proto.CompactTextString(m) }
func (*BroadcastResponse) ProtoMessage()               {}
func (*BroadcastResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SeekNewest struct {
}

func (m *SeekNewest) Reset()                    { *m = SeekNewest{} }
func (m *SeekNewest) String() string            { return proto.CompactTextString(m) }
func (*SeekNewest) ProtoMessage()               {}
func (*SeekNewest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SeekOldest struct {
}

func (m *SeekOldest) Reset()                    { *m = SeekOldest{} }
func (m *SeekOldest) String() string            { return proto.CompactTextString(m) }
func (*SeekOldest) ProtoMessage()               {}
func (*SeekOldest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SeekSpecified struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *SeekSpecified) Reset()                    { *m = SeekSpecified{} }
func (m *SeekSpecified) String() string            { return proto.CompactTextString(m) }
func (*SeekSpecified) ProtoMessage()               {}
func (*SeekSpecified) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SeekPosition struct {
	// Types that are valid to be assigned to Type:
	//	*SeekPosition_Newest
	//	*SeekPosition_Oldest
	//	*SeekPosition_Specified
	Type isSeekPosition_Type `protobuf_oneof:"Type"`
}

func (m *SeekPosition) Reset()                    { *m = SeekPosition{} }
func (m *SeekPosition) String() string            { return proto.CompactTextString(m) }
func (*SeekPosition) ProtoMessage()               {}
func (*SeekPosition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isSeekPosition_Type interface {
	isSeekPosition_Type()
}

type SeekPosition_Newest struct {
	Newest *SeekNewest `protobuf:"bytes,1,opt,name=newest,oneof"`
}
type SeekPosition_Oldest struct {
	Oldest *SeekOldest `protobuf:"bytes,2,opt,name=oldest,oneof"`
}
type SeekPosition_Specified struct {
	Specified *SeekSpecified `protobuf:"bytes,3,opt,name=specified,oneof"`
}

func (*SeekPosition_Newest) isSeekPosition_Type()    {}
func (*SeekPosition_Oldest) isSeekPosition_Type()    {}
func (*SeekPosition_Specified) isSeekPosition_Type() {}

func (m *SeekPosition) GetType() isSeekPosition_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SeekPosition) GetNewest() *SeekNewest {
	if x, ok := m.GetType().(*SeekPosition_Newest); ok {
		return x.Newest
	}
	return nil
}

func (m *SeekPosition) GetOldest() *SeekOldest {
	if x, ok := m.GetType().(*SeekPosition_Oldest); ok {
		return x.Oldest
	}
	return nil
}

func (m *SeekPosition) GetSpecified() *SeekSpecified {
	if x, ok := m.GetType().(*SeekPosition_Specified); ok {
		return x.Specified
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SeekPosition) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SeekPosition_OneofMarshaler, _SeekPosition_OneofUnmarshaler, _SeekPosition_OneofSizer, []interface{}{
		(*SeekPosition_Newest)(nil),
		(*SeekPosition_Oldest)(nil),
		(*SeekPosition_Specified)(nil),
	}
}

func _SeekPosition_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SeekPosition)
	// Type
	switch x := m.Type.(type) {
	case *SeekPosition_Newest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Newest); err != nil {
			return err
		}
	case *SeekPosition_Oldest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Oldest); err != nil {
			return err
		}
	case *SeekPosition_Specified:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Specified); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SeekPosition.Type has unexpected type %T", x)
	}
	return nil
}

func _SeekPosition_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SeekPosition)
	switch tag {
	case 1: // Type.newest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SeekNewest)
		err := b.DecodeMessage(msg)
		m.Type = &SeekPosition_Newest{msg}
		return true, err
	case 2: // Type.oldest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SeekOldest)
		err := b.DecodeMessage(msg)
		m.Type = &SeekPosition_Oldest{msg}
		return true, err
	case 3: // Type.specified
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SeekSpecified)
		err := b.DecodeMessage(msg)
		m.Type = &SeekPosition_Specified{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SeekPosition_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SeekPosition)
	// Type
	switch x := m.Type.(type) {
	case *SeekPosition_Newest:
		s := proto.Size(x.Newest)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SeekPosition_Oldest:
		s := proto.Size(x.Oldest)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SeekPosition_Specified:
		s := proto.Size(x.Specified)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// SeekInfo specifies the range of requested blocks to return
// If the start position is not found, an error is immediately returned
// Otherwise, blocks are returned until a missing block is encountered, then behavior is dictated
// by the SeekBehavior specified.  If BLOCK_UNTIL_READY is specified, the reply will block until
// the requested blocks are available, if FAIL_IF_NOT_READY is specified, the reply will return an
// error indicating that the block is not found.  To request that all blocks be returned indefinitely
// as they are created, behavior should be set to BLOCK_UNTIL_READY and the stop should be set to
// specified with a number of MAX_UINT64
type SeekInfo struct {
	Start    *SeekPosition         `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	Stop     *SeekPosition         `protobuf:"bytes,2,opt,name=stop" json:"stop,omitempty"`
	Behavior SeekInfo_SeekBehavior `protobuf:"varint,3,opt,name=behavior,enum=orderer.SeekInfo_SeekBehavior" json:"behavior,omitempty"`
}

func (m *SeekInfo) Reset()                    { *m = SeekInfo{} }
func (m *SeekInfo) String() string            { return proto.CompactTextString(m) }
func (*SeekInfo) ProtoMessage()               {}
func (*SeekInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SeekInfo) GetStart() *SeekPosition {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *SeekInfo) GetStop() *SeekPosition {
	if m != nil {
		return m.Stop
	}
	return nil
}

type DeliverResponse struct {
	// Types that are valid to be assigned to Type:
	//	*DeliverResponse_Status
	//	*DeliverResponse_Block
	Type isDeliverResponse_Type `protobuf_oneof:"Type"`
}

func (m *DeliverResponse) Reset()                    { *m = DeliverResponse{} }
func (m *DeliverResponse) String() string            { return proto.CompactTextString(m) }
func (*DeliverResponse) ProtoMessage()               {}
func (*DeliverResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isDeliverResponse_Type interface {
	isDeliverResponse_Type()
}

type DeliverResponse_Status struct {
	Status common.Status `protobuf:"varint,1,opt,name=status,enum=common.Status,oneof"`
}
type DeliverResponse_Block struct {
	Block *common.Block `protobuf:"bytes,2,opt,name=block,oneof"`
}

func (*DeliverResponse_Status) isDeliverResponse_Type() {}
func (*DeliverResponse_Block) isDeliverResponse_Type()  {}

func (m *DeliverResponse) GetType() isDeliverResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DeliverResponse) GetStatus() common.Status {
	if x, ok := m.GetType().(*DeliverResponse_Status); ok {
		return x.Status
	}
	return common.Status_UNKNOWN
}

func (m *DeliverResponse) GetBlock() *common.Block {
	if x, ok := m.GetType().(*DeliverResponse_Block); ok {
		return x.Block
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeliverResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DeliverResponse_OneofMarshaler, _DeliverResponse_OneofUnmarshaler, _DeliverResponse_OneofSizer, []interface{}{
		(*DeliverResponse_Status)(nil),
		(*DeliverResponse_Block)(nil),
	}
}

func _DeliverResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeliverResponse)
	// Type
	switch x := m.Type.(type) {
	case *DeliverResponse_Status:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Status))
	case *DeliverResponse_Block:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DeliverResponse.Type has unexpected type %T", x)
	}
	return nil
}

func _DeliverResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeliverResponse)
	switch tag {
	case 1: // Type.status
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &DeliverResponse_Status{common.Status(x)}
		return true, err
	case 2: // Type.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.Block)
		err := b.DecodeMessage(msg)
		m.Type = &DeliverResponse_Block{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DeliverResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DeliverResponse)
	// Type
	switch x := m.Type.(type) {
	case *DeliverResponse_Status:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Status))
	case *DeliverResponse_Block:
		s := proto.Size(x.Block)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BroadcastResponse)(nil), "orderer.BroadcastResponse")
	proto.RegisterType((*SeekNewest)(nil), "orderer.SeekNewest")
	proto.RegisterType((*SeekOldest)(nil), "orderer.SeekOldest")
	proto.RegisterType((*SeekSpecified)(nil), "orderer.SeekSpecified")
	proto.RegisterType((*SeekPosition)(nil), "orderer.SeekPosition")
	proto.RegisterType((*SeekInfo)(nil), "orderer.SeekInfo")
	proto.RegisterType((*DeliverResponse)(nil), "orderer.DeliverResponse")
	proto.RegisterEnum("orderer.SeekInfo_SeekBehavior", SeekInfo_SeekBehavior_name, SeekInfo_SeekBehavior_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for AtomicBroadcast service

type AtomicBroadcastClient interface {
	// broadcast receives a reply of Acknowledgement for each common.Envelope in order, indicating success or type of failure
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_BroadcastClient, error)
	// deliver first requires an Envelope of type DELIVER_SEEK_INFO with Payload data as a mashaled SeekInfo message, then a stream of block replies is received.
	Deliver(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_DeliverClient, error)
}

type atomicBroadcastClient struct {
	cc *grpc.ClientConn
}

func NewAtomicBroadcastClient(cc *grpc.ClientConn) AtomicBroadcastClient {
	return &atomicBroadcastClient{cc}
}

func (c *atomicBroadcastClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_BroadcastClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AtomicBroadcast_serviceDesc.Streams[0], c.cc, "/orderer.AtomicBroadcast/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &atomicBroadcastBroadcastClient{stream}
	return x, nil
}

type AtomicBroadcast_BroadcastClient interface {
	Send(*common.Envelope) error
	Recv() (*BroadcastResponse, error)
	grpc.ClientStream
}

type atomicBroadcastBroadcastClient struct {
	grpc.ClientStream
}

func (x *atomicBroadcastBroadcastClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *atomicBroadcastBroadcastClient) Recv() (*BroadcastResponse, error) {
	m := new(BroadcastResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *atomicBroadcastClient) Deliver(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_DeliverClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AtomicBroadcast_serviceDesc.Streams[1], c.cc, "/orderer.AtomicBroadcast/Deliver", opts...)
	if err != nil {
		return nil, err
	}
	x := &atomicBroadcastDeliverClient{stream}
	return x, nil
}

type AtomicBroadcast_DeliverClient interface {
	Send(*common.Envelope) error
	Recv() (*DeliverResponse, error)
	grpc.ClientStream
}

type atomicBroadcastDeliverClient struct {
	grpc.ClientStream
}

func (x *atomicBroadcastDeliverClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *atomicBroadcastDeliverClient) Recv() (*DeliverResponse, error) {
	m := new(DeliverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AtomicBroadcast service

type AtomicBroadcastServer interface {
	// broadcast receives a reply of Acknowledgement for each common.Envelope in order, indicating success or type of failure
	Broadcast(AtomicBroadcast_BroadcastServer) error
	// deliver first requires an Envelope of type DELIVER_SEEK_INFO with Payload data as a mashaled SeekInfo message, then a stream of block replies is received.
	Deliver(AtomicBroadcast_DeliverServer) error
}

func RegisterAtomicBroadcastServer(s *grpc.Server, srv AtomicBroadcastServer) {
	s.RegisterService(&_AtomicBroadcast_serviceDesc, srv)
}

func _AtomicBroadcast_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AtomicBroadcastServer).Broadcast(&atomicBroadcastBroadcastServer{stream})
}

type AtomicBroadcast_BroadcastServer interface {
	Send(*BroadcastResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type atomicBroadcastBroadcastServer struct {
	grpc.ServerStream
}

func (x *atomicBroadcastBroadcastServer) Send(m *BroadcastResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *atomicBroadcastBroadcastServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AtomicBroadcast_Deliver_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AtomicBroadcastServer).Deliver(&atomicBroadcastDeliverServer{stream})
}

type AtomicBroadcast_DeliverServer interface {
	Send(*DeliverResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type atomicBroadcastDeliverServer struct {
	grpc.ServerStream
}

func (x *atomicBroadcastDeliverServer) Send(m *DeliverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *atomicBroadcastDeliverServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AtomicBroadcast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderer.AtomicBroadcast",
	HandlerType: (*AtomicBroadcastServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Broadcast",
			Handler:       _AtomicBroadcast_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Deliver",
			Handler:       _AtomicBroadcast_Deliver_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("orderer/ab.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x63, 0x48, 0xdd, 0x76, 0x68, 0xd3, 0x74, 0xab, 0x56, 0x51, 0x0e, 0x08, 0x59, 0x02,
	0x82, 0x00, 0x1b, 0x05, 0x89, 0x03, 0x45, 0x42, 0x31, 0x6d, 0x95, 0x88, 0x28, 0x41, 0x4e, 0x38,
	0xc0, 0x25, 0xf2, 0xc7, 0xa4, 0x31, 0x75, 0xbc, 0xd6, 0xee, 0x26, 0xa8, 0x4f, 0xc1, 0x8b, 0xf0,
	0x48, 0x3c, 0x0c, 0xda, 0x0f, 0x3b, 0x0a, 0x14, 0x4e, 0xc9, 0xcc, 0xfc, 0xfe, 0xb3, 0xff, 0x19,
	0x8d, 0xa1, 0x49, 0x59, 0x82, 0x0c, 0x99, 0x17, 0x46, 0x6e, 0xc1, 0xa8, 0xa0, 0x64, 0xd7, 0x64,
	0xda, 0x27, 0x31, 0x5d, 0x2e, 0x69, 0xee, 0xe9, 0x1f, 0x5d, 0x75, 0xce, 0xe1, 0xd8, 0x67, 0x34,
	0x4c, 0xe2, 0x90, 0x8b, 0x00, 0x79, 0x41, 0x73, 0x8e, 0xe4, 0x09, 0xd8, 0x5c, 0x84, 0x62, 0xc5,
	0x5b, 0xd6, 0x23, 0xab, 0xd3, 0xe8, 0x36, 0x5c, 0xa3, 0x99, 0xa8, 0x6c, 0x60, 0xaa, 0xce, 0x01,
	0xc0, 0x04, 0xf1, 0x66, 0x84, 0xdf, 0x91, 0x8b, 0x32, 0x1a, 0x67, 0x89, 0x8c, 0x9e, 0xc2, 0xa1,
	0x8c, 0x26, 0x05, 0xc6, 0xe9, 0x3c, 0xc5, 0x84, 0x9c, 0x81, 0x9d, 0xaf, 0x96, 0x11, 0x32, 0xd5,
	0xb4, 0x1e, 0x98, 0xc8, 0xf9, 0x69, 0xc1, 0x81, 0x24, 0x3f, 0x51, 0x9e, 0x8a, 0x94, 0xe6, 0xe4,
	0x25, 0xd8, 0xb9, 0xea, 0xa8, 0xc0, 0x07, 0xdd, 0x13, 0xd7, 0x4c, 0xe0, 0x6e, 0x1e, 0xeb, 0xd7,
	0x02, 0x03, 0x49, 0x9c, 0xaa, 0x27, 0x5b, 0xf7, 0xee, 0xc0, 0xb5, 0x1b, 0x89, 0x6b, 0x88, 0xbc,
	0x81, 0x7d, 0x5e, 0x7a, 0x6a, 0xdd, 0x57, 0x8a, 0xb3, 0x2d, 0x45, 0xe5, 0xb8, 0x5f, 0x0b, 0x36,
	0xa8, 0x6f, 0x43, 0x7d, 0x7a, 0x5b, 0xa0, 0xf3, 0xcb, 0x82, 0x3d, 0x89, 0x0d, 0xf2, 0x39, 0x25,
	0xcf, 0x61, 0x87, 0x8b, 0x90, 0x95, 0x4e, 0x4f, 0xb7, 0x1a, 0x95, 0x03, 0x05, 0x9a, 0x21, 0xcf,
	0xa0, 0xce, 0x05, 0x2d, 0x8c, 0xcd, 0x7f, 0xb0, 0x0a, 0x21, 0x6f, 0x61, 0x2f, 0xc2, 0x45, 0xb8,
	0x4e, 0x29, 0x53, 0x1e, 0x1b, 0xdd, 0x87, 0x5b, 0xb8, 0x7c, 0x5c, 0xfd, 0xf1, 0x0d, 0x15, 0x54,
	0xbc, 0xf3, 0x4e, 0xaf, 0xb3, 0xac, 0x90, 0x53, 0x38, 0xf6, 0x87, 0xe3, 0x0f, 0x1f, 0x67, 0x9f,
	0x47, 0xd3, 0xc1, 0x70, 0x16, 0x5c, 0xf6, 0x2e, 0xbe, 0x34, 0x6b, 0x32, 0x7d, 0xd5, 0x1b, 0x0c,
	0x67, 0x83, 0xab, 0xd9, 0x68, 0x3c, 0x35, 0x69, 0xcb, 0xf9, 0x06, 0x47, 0x17, 0x98, 0xa5, 0x6b,
	0x64, 0xd5, 0x35, 0x74, 0xfe, 0x7f, 0x0d, 0x72, 0xb7, 0xba, 0x4e, 0x1e, 0xc3, 0x4e, 0x94, 0xd1,
	0xf8, 0xc6, 0x8c, 0x78, 0x58, 0x82, 0xbe, 0x4c, 0xf6, 0x6b, 0x81, 0xae, 0x96, 0xab, 0xec, 0xfe,
	0xb0, 0xe0, 0xa8, 0x27, 0xe8, 0x32, 0x8d, 0xab, 0x13, 0x24, 0xef, 0x61, 0x7f, 0x13, 0x34, 0xcb,
	0x06, 0x97, 0xf9, 0x1a, 0x33, 0x5a, 0x60, 0xbb, 0x5d, 0xad, 0xe1, 0xaf, 0xab, 0x75, 0x6a, 0x1d,
	0xeb, 0x95, 0x45, 0xce, 0x61, 0xd7, 0x0c, 0x70, 0x87, 0xbc, 0x55, 0xc9, 0xff, 0x18, 0x52, 0x8b,
	0x7d, 0xf7, 0xeb, 0x8b, 0xeb, 0x54, 0x2c, 0x56, 0x91, 0x54, 0x7a, 0x8b, 0xdb, 0x02, 0x59, 0x86,
	0xc9, 0x35, 0x32, 0x6f, 0x1e, 0x46, 0x2c, 0x8d, 0x3d, 0xf5, 0xd1, 0x70, 0xcf, 0x74, 0x89, 0x6c,
	0x15, 0xbf, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x79, 0xcc, 0xb7, 0x76, 0x03, 0x00, 0x00,
}
