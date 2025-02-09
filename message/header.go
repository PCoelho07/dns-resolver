package message

import (
	"bytes"
	"encoding/binary"
)

type HeaderType struct {
	ID      uint16
	Flags   HeaderFlag
	QdCount uint16
	AnCount uint16
	NsCount uint16
	ArCount uint16
}

type HeaderFlag struct {
	QR     bool
	OpCode uint8
	AA     bool
	TC     bool
	RD     bool
	RA     bool
	Z      uint8
	RCode  uint8
}

func NewHeader(id uint16, flag HeaderFlag, qdCount, anCount, nsCount, arCount uint16) HeaderType {
    return HeaderType{
        ID: id,
        Flags: flag,
        QdCount: qdCount,
        AnCount: anCount,
        NsCount: nsCount,
        ArCount: arCount,
    }
}

func (hFlag *HeaderFlag) GenerateFlags() uint16 {
	qr := uint16(BoolToInt(hFlag.QR))
	opCode := uint16(hFlag.OpCode)
	aa := uint16(BoolToInt(hFlag.AA))
	tc := uint16(BoolToInt(hFlag.TC))
	rd := uint16(BoolToInt(hFlag.RD))
	ra := uint16(BoolToInt(hFlag.RA))
	rCode := uint16(hFlag.RCode)
	z := uint16(hFlag.RCode)

	return uint16(qr<<15 | opCode<<11 | aa<<10 | tc<<9 | rd<<8 | ra<<7 | z<<4 | rCode)
}

func (h *HeaderType) ToBytes() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, h.ID)
	binary.Write(buf, binary.BigEndian, h.Flags.GenerateFlags())
	binary.Write(buf, binary.BigEndian, h.QdCount)
	binary.Write(buf, binary.BigEndian, h.AnCount)
	binary.Write(buf, binary.BigEndian, h.NsCount)
	binary.Write(buf, binary.BigEndian, h.ArCount)

	return buf.Bytes()
}
