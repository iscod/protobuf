// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: goods.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetGoodsAck struct {
	Goods                *Goods   `protobuf:"bytes,1,opt,name=goods,proto3" json:"goods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGoodsAck) Reset()         { *m = GetGoodsAck{} }
func (m *GetGoodsAck) String() string { return proto.CompactTextString(m) }
func (*GetGoodsAck) ProtoMessage()    {}
func (*GetGoodsAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30593c5487368b0, []int{0}
}
func (m *GetGoodsAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetGoodsAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetGoodsAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetGoodsAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGoodsAck.Merge(m, src)
}
func (m *GetGoodsAck) XXX_Size() int {
	return m.Size()
}
func (m *GetGoodsAck) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGoodsAck.DiscardUnknown(m)
}

var xxx_messageInfo_GetGoodsAck proto.InternalMessageInfo

func (m *GetGoodsAck) GetGoods() *Goods {
	if m != nil {
		return m.Goods
	}
	return nil
}

func init() {
	proto.RegisterType((*GetGoodsAck)(nil), "pb.GetGoodsAck")
}

func init() { proto.RegisterFile("goods.proto", fileDescriptor_a30593c5487368b0) }

var fileDescriptor_a30593c5487368b0 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcf, 0xcf, 0x4f,
	0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x49, 0xce, 0xcf,
	0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x28, 0xe9, 0x71, 0x71, 0xbb, 0xa7, 0x96, 0xb8, 0x83, 0xd4, 0x38,
	0x26, 0x67, 0x0b, 0xc9, 0x73, 0xb1, 0x82, 0xd5, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x71,
	0xea, 0x15, 0x24, 0xe9, 0x81, 0x25, 0x83, 0x20, 0xe2, 0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0x83,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x45, 0x80, 0x70, 0x69, 0x00, 0x00, 0x00,
}

func (m *GetGoodsAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetGoodsAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetGoodsAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Goods != nil {
		{
			size, err := m.Goods.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGoods(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGoods(dAtA []byte, offset int, v uint64) int {
	offset -= sovGoods(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetGoodsAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Goods != nil {
		l = m.Goods.Size()
		n += 1 + l + sovGoods(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGoods(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGoods(x uint64) (n int) {
	return sovGoods(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetGoodsAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoods
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
			return fmt.Errorf("proto: GetGoodsAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetGoodsAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Goods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoods
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
				return ErrInvalidLengthGoods
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGoods
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Goods == nil {
				m.Goods = &Goods{}
			}
			if err := m.Goods.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGoods(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGoods
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGoods
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGoods(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGoods
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
					return 0, ErrIntOverflowGoods
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
					return 0, ErrIntOverflowGoods
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
				return 0, ErrInvalidLengthGoods
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGoods
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGoods
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGoods        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGoods          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGoods = fmt.Errorf("proto: unexpected end of group")
)
