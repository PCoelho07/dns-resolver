package message

import "encoding/binary"


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

func NewHeaderFlag(qr bool, opCode uint8, aa, tc, rd, ra bool, z, rCode uint8) HeaderFlag {
    return HeaderFlag{
		QR:     qr,
		OpCode: opCode,
		AA:     aa,
		TC:     tc,
		RD:     rd,
		RA:     ra,
		Z:      z,
		RCode:  rCode,
    }
}

func (hFlag HeaderFlag) ToBytes() uint16 {
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

func HeaderFlagFromBytes(data []byte) HeaderFlag {
	readData := binary.BigEndian.Uint16(data)

	return HeaderFlag{
		QR:     extractBits(readData, 15, 1) == 1,
		OpCode: uint8(extractBits(readData, 11, 4)),
		AA:     extractBits(readData, 10, 1) == 1,
		TC:     extractBits(readData, 9, 1) == 1,
		RD:     extractBits(readData, 8, 1) == 1,
		RA:     extractBits(readData, 7, 1) == 1,
		Z:      uint8(extractBits(readData, 4, 3)),
		RCode:  uint8(extractBits(readData, 0, 4)),
	}
}

func extractBits(value uint16, offset, length uint8) uint16 {
    return uint16(value>>offset) & ((1 << length) - 1)
}
