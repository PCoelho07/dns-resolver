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

func NewHeader(id uint16, flag HeaderFlag, qdCount, anCount, nsCount, arCount uint16) HeaderType {
	return HeaderType{
		ID:      id,
		Flags:   flag,
		QdCount: qdCount,
		AnCount: anCount,
		NsCount: nsCount,
		ArCount: arCount,
	}
}

func (h HeaderType) ToBytes() []byte {
    buf := new(bytes.Buffer)

    binary.Write(buf, binary.BigEndian, h.ID)
    binary.Write(buf, binary.BigEndian, h.Flags.ToBytes())
    binary.Write(buf, binary.BigEndian, h.QdCount)
    binary.Write(buf, binary.BigEndian, h.AnCount)
    binary.Write(buf, binary.BigEndian, h.NsCount)
    binary.Write(buf, binary.BigEndian, h.ArCount)

    return buf.Bytes()
}

func HeaderFromBytes(data []byte) HeaderType {
	id := binary.BigEndian.Uint16(data[0:2])
	flags := HeaderFlagFromBytes(data[2:4])
	qdCount := binary.BigEndian.Uint16(data[4:6])
	anCount := binary.BigEndian.Uint16(data[6:8])
	nsCount := binary.BigEndian.Uint16(data[8:10])
	arCount := binary.BigEndian.Uint16(data[10:12])

	headerResult := HeaderType{
		ID:      id,
		Flags:   flags,
		QdCount: qdCount,
		NsCount: nsCount,
		AnCount: anCount,
		ArCount: arCount,
	}

	return headerResult
}
