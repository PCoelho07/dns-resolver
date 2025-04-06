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
	z := uint16(hFlag.Z)

	return uint16(qr<<15 | opCode<<11 | aa<<10 | tc<<9 | rd<<8 | ra<<7 | z<<4 | rCode)
}

func extractBits(value uint16, offset, length uint8) uint16 {
    return uint16(value >> offset) & ((1 << length) - 1)
}

func (hflag *HeaderFlag) DecodeFromBytes(data []byte) HeaderFlag {
    readedData := binary.BigEndian.Uint16(data)

    return HeaderFlag{
        QR: extractBits(readedData, 15, 1) == 1,
        OpCode: uint8(extractBits(readedData, 11, 4)),
        AA: extractBits(readedData, 10, 1) == 1,
        TC: extractBits(readedData, 9, 1) == 1,
        RD: extractBits(readedData, 8, 1) == 1,
        RA: extractBits(readedData, 7, 1) == 1,
        Z: uint8(extractBits(readedData, 4, 3)),
        RCode: uint8(extractBits(readedData, 0, 4)),
    }
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

func (h *HeaderType) DecodeFromBytes(data []byte) (*HeaderType, error) {
    id := binary.BigEndian.Uint16(data[0:2])
    flags := h.Flags.DecodeFromBytes(data[2:4])
    qdCount := binary.BigEndian.Uint16(data[4:6])
    anCount := binary.BigEndian.Uint16(data[6:8])
    nsCount := binary.BigEndian.Uint16(data[8:10])
    arCount := binary.BigEndian.Uint16(data[10:12])

    headerResult := &HeaderType{
        ID: id,
        Flags: flags,
        QdCount: qdCount,
        NsCount: nsCount,
        AnCount: anCount,
        ArCount: arCount,
    }

    return headerResult, nil
}
