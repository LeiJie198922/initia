// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/fetchprice/v1/types.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CurrencyPrice is the information necessary for initialization of a
// CurrencyPrice.
type CurrencyPrice struct {
	// The currency id is the string with "BASE/QUOTE" format.
	CurrencyId string     `protobuf:"bytes,1,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	QuotePrice QuotePrice `protobuf:"bytes,2,opt,name=quote_price,json=quotePrice,proto3" json:"quote_price"`
}

func (m *CurrencyPrice) Reset()         { *m = CurrencyPrice{} }
func (m *CurrencyPrice) String() string { return proto.CompactTextString(m) }
func (*CurrencyPrice) ProtoMessage()    {}
func (*CurrencyPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_853a8402b4739f0c, []int{0}
}
func (m *CurrencyPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrencyPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrencyPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrencyPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyPrice.Merge(m, src)
}
func (m *CurrencyPrice) XXX_Size() int {
	return m.Size()
}
func (m *CurrencyPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyPrice.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyPrice proto.InternalMessageInfo

func (m *CurrencyPrice) GetCurrencyId() string {
	if m != nil {
		return m.CurrencyId
	}
	return ""
}

func (m *CurrencyPrice) GetQuotePrice() QuotePrice {
	if m != nil {
		return m.QuotePrice
	}
	return QuotePrice{}
}

// QuotePrice is the representation of the aggregated prices for a CurrencyPair,
// where price represents the price of Base in terms of Quote
type QuotePrice struct {
	Price cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=price,proto3,customtype=cosmossdk.io/math.Int" json:"price"`
	// decimals represents the number of decimals that the quote-price is
	// represented in. For Pairs where ETHEREUM is the quote this will be 18,
	// otherwise it will be 8.
	Decimals uint64 `protobuf:"varint,2,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// BlockTimestamp tracks the block height associated with this price update.
	// We include block timestamp alongside the price to ensure that smart
	// contracts and applications are not utilizing stale oracle prices.
	BlockTimestamp time.Time `protobuf:"bytes,3,opt,name=block_timestamp,json=blockTimestamp,proto3,stdtime" json:"block_timestamp"`
	// BlockHeight is block height of provider chain.
	BlockHeight uint64 `protobuf:"varint,4,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *QuotePrice) Reset()         { *m = QuotePrice{} }
func (m *QuotePrice) String() string { return proto.CompactTextString(m) }
func (*QuotePrice) ProtoMessage()    {}
func (*QuotePrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_853a8402b4739f0c, []int{1}
}
func (m *QuotePrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuotePrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuotePrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuotePrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotePrice.Merge(m, src)
}
func (m *QuotePrice) XXX_Size() int {
	return m.Size()
}
func (m *QuotePrice) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotePrice.DiscardUnknown(m)
}

var xxx_messageInfo_QuotePrice proto.InternalMessageInfo

func (m *QuotePrice) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *QuotePrice) GetBlockTimestamp() time.Time {
	if m != nil {
		return m.BlockTimestamp
	}
	return time.Time{}
}

func (m *QuotePrice) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*CurrencyPrice)(nil), "ibc.applications.fetchprice.v1.CurrencyPrice")
	proto.RegisterType((*QuotePrice)(nil), "ibc.applications.fetchprice.v1.QuotePrice")
}

func init() {
	proto.RegisterFile("ibc/applications/fetchprice/v1/types.proto", fileDescriptor_853a8402b4739f0c)
}

var fileDescriptor_853a8402b4739f0c = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xbd, 0xce, 0xd3, 0x30,
	0x14, 0x8d, 0xa1, 0xa0, 0x0f, 0x87, 0x1f, 0x29, 0x02, 0x29, 0x64, 0x48, 0x3e, 0x3a, 0x55, 0x45,
	0xb5, 0x55, 0xe0, 0x05, 0x08, 0x0b, 0x1d, 0x40, 0x34, 0x62, 0x62, 0x89, 0x12, 0xc7, 0x4d, 0xac,
	0x26, 0x71, 0x1a, 0x3b, 0x15, 0x9d, 0x79, 0x81, 0x3e, 0x0c, 0x0f, 0xd1, 0xb1, 0x62, 0x42, 0x48,
	0x14, 0xd4, 0xbe, 0x08, 0x8a, 0x9d, 0xa6, 0x9d, 0xd8, 0x7c, 0xcf, 0x3d, 0xc7, 0xe7, 0xdc, 0xab,
	0x0b, 0xc7, 0x2c, 0x26, 0x38, 0xaa, 0xaa, 0x9c, 0x91, 0x48, 0x32, 0x5e, 0x0a, 0xbc, 0xa0, 0x92,
	0x64, 0x55, 0xcd, 0x08, 0xc5, 0xeb, 0x29, 0x96, 0x9b, 0x8a, 0x0a, 0x54, 0xd5, 0x5c, 0x72, 0xcb,
	0x65, 0x31, 0x41, 0xd7, 0x5c, 0x74, 0xe1, 0xa2, 0xf5, 0xd4, 0x79, 0x9a, 0xf2, 0x94, 0x2b, 0x2a,
	0x6e, 0x5f, 0x5a, 0xe5, 0x3c, 0x27, 0x5c, 0x14, 0x5c, 0x84, 0xba, 0xa1, 0x8b, 0xae, 0xe5, 0xa5,
	0x9c, 0xa7, 0x39, 0xc5, 0xaa, 0x8a, 0x9b, 0x05, 0x96, 0xac, 0xa0, 0x42, 0x46, 0x45, 0xa5, 0x09,
	0xc3, 0x6f, 0x00, 0x3e, 0x7a, 0xd7, 0xd4, 0x35, 0x2d, 0xc9, 0xe6, 0x53, 0x6b, 0x63, 0x79, 0xd0,
	0x24, 0x1d, 0x10, 0xb2, 0xc4, 0x06, 0xb7, 0x60, 0xf4, 0x20, 0x80, 0x67, 0x68, 0x96, 0x58, 0x73,
	0x68, 0xae, 0x1a, 0x2e, 0x69, 0xa8, 0x62, 0xd9, 0x77, 0x6e, 0xc1, 0xc8, 0x7c, 0x35, 0x46, 0xff,
	0x8f, 0x8e, 0xe6, 0xad, 0x44, 0x39, 0xf8, 0x83, 0xdd, 0xc1, 0x33, 0x02, 0xb8, 0xea, 0x91, 0xe1,
	0x6f, 0x00, 0xe1, 0x85, 0x60, 0xbd, 0x85, 0xf7, 0xf4, 0xdf, 0xca, 0xdc, 0x7f, 0xd9, 0xf2, 0x7f,
	0x1d, 0xbc, 0x67, 0x7a, 0x34, 0x91, 0x2c, 0x11, 0xe3, 0xb8, 0x88, 0x64, 0x86, 0x66, 0xa5, 0xfc,
	0xf1, 0x7d, 0x02, 0xbb, 0x99, 0x67, 0xa5, 0x0c, 0xb4, 0xd2, 0x72, 0xe0, 0x4d, 0x42, 0x09, 0x2b,
	0xa2, 0x5c, 0xa8, 0x84, 0x83, 0xa0, 0xaf, 0xad, 0x0f, 0xf0, 0x49, 0x9c, 0x73, 0xb2, 0x0c, 0xfb,
	0x65, 0xd8, 0x77, 0xd5, 0x10, 0x0e, 0xd2, 0xeb, 0x42, 0xe7, 0x75, 0xa1, 0xcf, 0x67, 0x86, 0x7f,
	0xd3, 0x86, 0xd8, 0xfe, 0xf1, 0x40, 0xf0, 0x58, 0x89, 0xfb, 0x8e, 0xf5, 0x02, 0x3e, 0xd4, 0xdf,
	0x65, 0x94, 0xa5, 0x99, 0xb4, 0x07, 0xca, 0xce, 0x54, 0xd8, 0x7b, 0x05, 0xf9, 0x1f, 0x77, 0x47,
	0x17, 0xec, 0x8f, 0x2e, 0xf8, 0x7b, 0x74, 0xc1, 0xf6, 0xe4, 0x1a, 0xfb, 0x93, 0x6b, 0xfc, 0x3c,
	0xb9, 0xc6, 0x97, 0x37, 0x29, 0x93, 0x59, 0x13, 0x23, 0xc2, 0x0b, 0xcc, 0x4a, 0x26, 0x59, 0x34,
	0xc9, 0xa3, 0x58, 0x74, 0x6f, 0xfc, 0x15, 0xb7, 0xd7, 0x73, 0x75, 0x30, 0xea, 0x5a, 0xe2, 0xfb,
	0x2a, 0xe0, 0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x7d, 0xf6, 0xf4, 0x5c, 0x02, 0x00,
	0x00,
}

func (m *CurrencyPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrencyPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrencyPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.QuotePrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.CurrencyId) > 0 {
		i -= len(m.CurrencyId)
		copy(dAtA[i:], m.CurrencyId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CurrencyId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuotePrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuotePrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuotePrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.BlockTimestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.BlockTimestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if m.Decimals != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CurrencyPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CurrencyId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.QuotePrice.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *QuotePrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Price.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Decimals != 0 {
		n += 1 + sovTypes(uint64(m.Decimals))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.BlockTimestamp)
	n += 1 + l + sovTypes(uint64(l))
	if m.BlockHeight != 0 {
		n += 1 + sovTypes(uint64(m.BlockHeight))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CurrencyPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: CurrencyPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrencyPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrencyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuotePrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuotePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *QuotePrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: QuotePrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotePrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.BlockTimestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
