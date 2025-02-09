package message

import (
	"bytes"
	"encoding/binary"
)

type ResourceRecord struct {
	Name     string
	Type     uint16
	Class    uint16
	TTL      uint32
	RDLength uint16
	RData    []byte
}

func NewResourceRecord(name string) *ResourceRecord {
    return &ResourceRecord{
        Name: name,
    }
}

func (rr *ResourceRecord) ToBytes() []byte {
	buffer := new(bytes.Buffer)

	buffer.Write([]byte(encodeName(rr.Name)))
	buffer.Write(rr.RData)
	binary.Write(buffer, binary.BigEndian, rr.Type)
	binary.Write(buffer, binary.BigEndian, rr.Class)
	binary.Write(buffer, binary.BigEndian, rr.TTL)

	return buffer.Bytes()

}
