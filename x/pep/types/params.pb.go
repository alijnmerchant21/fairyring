// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/pep/params.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Params defines the parameters for the module.
type Params struct {
	// option (gogoproto.equal) = true;
	KeyshareChannelId     string                 `protobuf:"bytes,1,opt,name=keyshare_channel_id,json=keyshareChannelId,proto3" json:"keyshare_channel_id,omitempty" yaml:"keyshare_channel_id"`
	IsSourceChain         bool                   `protobuf:"varint,2,opt,name=is_source_chain,json=isSourceChain,proto3" json:"is_source_chain,omitempty" yaml:"is_source_chain"`
	TrustedCounterParties []*TrustedCounterParty `protobuf:"bytes,3,rep,name=trusted_counter_parties,json=trustedCounterParties,proto3" json:"trusted_counter_parties,omitempty"`
	TrustedAddresses      []string               `protobuf:"bytes,4,rep,name=trusted_addresses,json=trustedAddresses,proto3" json:"trusted_addresses,omitempty" yaml:"trusted_addresses"`
	MinGasPrice           *types.Coin            `protobuf:"bytes,5,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price,omitempty" yaml:"min_gas_price"`
	PrivateKeysharePrice  *types.Coin            `protobuf:"bytes,6,opt,name=private_keyshare_price,json=privateKeysharePrice,proto3" json:"private_keyshare_price,omitempty" yaml:"private_keyshare_price"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a32cf7d58c7a431, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetKeyshareChannelId() string {
	if m != nil {
		return m.KeyshareChannelId
	}
	return ""
}

func (m *Params) GetIsSourceChain() bool {
	if m != nil {
		return m.IsSourceChain
	}
	return false
}

func (m *Params) GetTrustedCounterParties() []*TrustedCounterParty {
	if m != nil {
		return m.TrustedCounterParties
	}
	return nil
}

func (m *Params) GetTrustedAddresses() []string {
	if m != nil {
		return m.TrustedAddresses
	}
	return nil
}

func (m *Params) GetMinGasPrice() *types.Coin {
	if m != nil {
		return m.MinGasPrice
	}
	return nil
}

func (m *Params) GetPrivateKeysharePrice() *types.Coin {
	if m != nil {
		return m.PrivateKeysharePrice
	}
	return nil
}

// TrustedCounterParty defines the structure to store the ibc info
// of the source chain (fairyring) to reliably fetch active keys and
// aggregated/encrypted keyshares
type TrustedCounterParty struct {
	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ChannelId    string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (m *TrustedCounterParty) Reset()         { *m = TrustedCounterParty{} }
func (m *TrustedCounterParty) String() string { return proto.CompactTextString(m) }
func (*TrustedCounterParty) ProtoMessage()    {}
func (*TrustedCounterParty) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a32cf7d58c7a431, []int{1}
}
func (m *TrustedCounterParty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrustedCounterParty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrustedCounterParty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrustedCounterParty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustedCounterParty.Merge(m, src)
}
func (m *TrustedCounterParty) XXX_Size() int {
	return m.Size()
}
func (m *TrustedCounterParty) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustedCounterParty.DiscardUnknown(m)
}

var xxx_messageInfo_TrustedCounterParty proto.InternalMessageInfo

func (m *TrustedCounterParty) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *TrustedCounterParty) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *TrustedCounterParty) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "fairyring.pep.Params")
	proto.RegisterType((*TrustedCounterParty)(nil), "fairyring.pep.TrustedCounterParty")
}

func init() { proto.RegisterFile("fairyring/pep/params.proto", fileDescriptor_9a32cf7d58c7a431) }

var fileDescriptor_9a32cf7d58c7a431 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6e, 0x12, 0x41,
	0x1c, 0xc6, 0xd9, 0x52, 0x49, 0x19, 0x24, 0xca, 0x14, 0x71, 0xa5, 0x76, 0xc1, 0xf5, 0x42, 0x7a,
	0xd8, 0x4d, 0xeb, 0xad, 0x37, 0x17, 0xa3, 0x21, 0x26, 0x86, 0xac, 0x26, 0x26, 0xbd, 0x6c, 0x86,
	0xd9, 0x11, 0x26, 0x65, 0x67, 0x26, 0x33, 0x03, 0x91, 0x57, 0xf0, 0xe4, 0x8b, 0x98, 0xf8, 0x18,
	0x1e, 0x7b, 0xf4, 0x44, 0x0c, 0x1c, 0xbc, 0xf3, 0x04, 0x86, 0x99, 0x85, 0xda, 0x96, 0xa4, 0x97,
	0xcd, 0xcc, 0xf7, 0xff, 0xed, 0xff, 0x9b, 0x7c, 0xf9, 0x40, 0xf3, 0x0b, 0xa2, 0x72, 0x26, 0x29,
	0x1b, 0x86, 0x82, 0x88, 0x50, 0x20, 0x89, 0x32, 0x15, 0x08, 0xc9, 0x35, 0x87, 0xd5, 0xed, 0x2c,
	0x10, 0x44, 0x34, 0x6b, 0x28, 0xa3, 0x8c, 0x87, 0xe6, 0x6b, 0x89, 0x66, 0x7d, 0xc8, 0x87, 0xdc,
	0x1c, 0xc3, 0xf5, 0x29, 0x57, 0x3d, 0xcc, 0x55, 0xc6, 0x55, 0x38, 0x40, 0x8a, 0x84, 0xd3, 0xd3,
	0x01, 0xd1, 0xe8, 0x34, 0xc4, 0x9c, 0x32, 0x3b, 0xf7, 0x7f, 0xec, 0x83, 0x52, 0xdf, 0x18, 0xc1,
	0x0f, 0xe0, 0xf0, 0x92, 0xcc, 0xd4, 0x08, 0x49, 0x92, 0xe0, 0x11, 0x62, 0x8c, 0x8c, 0x13, 0x9a,
	0xba, 0x4e, 0xdb, 0xe9, 0x94, 0x23, 0x6f, 0x35, 0x6f, 0x35, 0x67, 0x28, 0x1b, 0x9f, 0xfb, 0x3b,
	0x20, 0x3f, 0xae, 0x6d, 0xd4, 0xae, 0x15, 0x7b, 0x29, 0x8c, 0xc0, 0x23, 0xaa, 0x12, 0xc5, 0x27,
	0x12, 0x1b, 0x96, 0x32, 0x77, 0xaf, 0xed, 0x74, 0x0e, 0xa2, 0xe6, 0x6a, 0xde, 0x6a, 0xd8, 0x5d,
	0xb7, 0x00, 0x3f, 0xae, 0x52, 0xf5, 0xd1, 0x08, 0xdd, 0xf5, 0x1d, 0x5e, 0x80, 0xa7, 0x5a, 0x4e,
	0x94, 0x26, 0x69, 0x82, 0xf9, 0x84, 0x69, 0x22, 0x13, 0x81, 0xa4, 0xa6, 0x44, 0xb9, 0xc5, 0x76,
	0xb1, 0x53, 0x39, 0xf3, 0x83, 0x1b, 0xc1, 0x04, 0x9f, 0x2c, 0xdd, 0xb5, 0x70, 0x1f, 0x49, 0x3d,
	0x8b, 0x9f, 0xe8, 0x3b, 0x22, 0x25, 0x0a, 0xf6, 0x40, 0x6d, 0xb3, 0x1b, 0xa5, 0xa9, 0x24, 0x4a,
	0x11, 0xe5, 0xee, 0xb7, 0x8b, 0x9d, 0x72, 0xf4, 0x7c, 0x35, 0x6f, 0xb9, 0xf6, 0x85, 0x77, 0x10,
	0x3f, 0x7e, 0x9c, 0x6b, 0xaf, 0x37, 0x12, 0xfc, 0x0c, 0xaa, 0x19, 0x65, 0xc9, 0x10, 0xa9, 0x44,
	0x48, 0x8a, 0x89, 0xfb, 0xa0, 0xed, 0x74, 0x2a, 0x67, 0xcf, 0x02, 0x9b, 0x7e, 0xb0, 0x4e, 0x3f,
	0xc8, 0xd3, 0x0f, 0xba, 0x9c, 0xb2, 0xc8, 0x5d, 0xcd, 0x5b, 0x75, 0xeb, 0x70, 0xe3, 0x4f, 0x3f,
	0xae, 0x64, 0x94, 0xbd, 0x43, 0xaa, 0xbf, 0xbe, 0x41, 0x09, 0x1a, 0x42, 0xd2, 0x29, 0xd2, 0x24,
	0xd9, 0xc6, 0x6e, 0x1d, 0x4a, 0xf7, 0x39, 0xbc, 0x58, 0xcd, 0x5b, 0xc7, 0xd6, 0x61, 0xf7, 0x0a,
	0x3f, 0xae, 0xe7, 0x83, 0xf7, 0xb9, 0x6e, 0x3c, 0xcf, 0x8f, 0xbe, 0xfd, 0xfd, 0x79, 0xd2, 0xb8,
	0xee, 0xe2, 0x57, 0xd3, 0x46, 0x5b, 0x12, 0x7f, 0x0a, 0x0e, 0x77, 0x44, 0x0c, 0x8f, 0x40, 0x19,
	0x8f, 0x29, 0x61, 0x7a, 0xdb, 0x98, 0xf8, 0xc0, 0x0a, 0xbd, 0x14, 0xbe, 0x04, 0x55, 0xcc, 0x19,
	0x23, 0x58, 0x53, 0xce, 0xd6, 0xc0, 0x9e, 0x01, 0x1e, 0x5e, 0x8b, 0xbd, 0x14, 0x1e, 0x03, 0xf0,
	0x5f, 0xe9, 0x8a, 0x86, 0x28, 0xe3, 0x4d, 0x99, 0xa2, 0x37, 0xbf, 0x16, 0x9e, 0x73, 0xb5, 0xf0,
	0x9c, 0x3f, 0x0b, 0xcf, 0xf9, 0xbe, 0xf4, 0x0a, 0x57, 0x4b, 0xaf, 0xf0, 0x7b, 0xe9, 0x15, 0x2e,
	0x4e, 0x86, 0x54, 0x8f, 0x26, 0x83, 0x00, 0xf3, 0x2c, 0x7c, 0x8b, 0xa8, 0x1c, 0x8c, 0x39, 0xbe,
	0x0c, 0x6f, 0x3f, 0x5f, 0xcf, 0x04, 0x51, 0x83, 0x92, 0x29, 0xfd, 0xab, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0x7b, 0x9e, 0xa6, 0x6a, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PrivateKeysharePrice != nil {
		{
			size, err := m.PrivateKeysharePrice.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.MinGasPrice != nil {
		{
			size, err := m.MinGasPrice.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TrustedAddresses) > 0 {
		for iNdEx := len(m.TrustedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TrustedAddresses[iNdEx])
			copy(dAtA[i:], m.TrustedAddresses[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.TrustedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.TrustedCounterParties) > 0 {
		for iNdEx := len(m.TrustedCounterParties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TrustedCounterParties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.IsSourceChain {
		i--
		if m.IsSourceChain {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.KeyshareChannelId) > 0 {
		i -= len(m.KeyshareChannelId)
		copy(dAtA[i:], m.KeyshareChannelId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.KeyshareChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TrustedCounterParty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrustedCounterParty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TrustedCounterParty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyshareChannelId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.IsSourceChain {
		n += 2
	}
	if len(m.TrustedCounterParties) > 0 {
		for _, e := range m.TrustedCounterParties {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.TrustedAddresses) > 0 {
		for _, s := range m.TrustedAddresses {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.MinGasPrice != nil {
		l = m.MinGasPrice.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.PrivateKeysharePrice != nil {
		l = m.PrivateKeysharePrice.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *TrustedCounterParty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyshareChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyshareChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSourceChain", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsSourceChain = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedCounterParties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustedCounterParties = append(m.TrustedCounterParties, &TrustedCounterParty{})
			if err := m.TrustedCounterParties[len(m.TrustedCounterParties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustedAddresses = append(m.TrustedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MinGasPrice == nil {
				m.MinGasPrice = &types.Coin{}
			}
			if err := m.MinGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivateKeysharePrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrivateKeysharePrice == nil {
				m.PrivateKeysharePrice = &types.Coin{}
			}
			if err := m.PrivateKeysharePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TrustedCounterParty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TrustedCounterParty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrustedCounterParty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
