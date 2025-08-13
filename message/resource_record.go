package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type ResourceRecord struct {
	Name        string
	Type        uint16
	Class       uint16
	TTL         uint32
	RDLength    uint16
	RData       []byte
	RDataParsed string
}

func NewResourceRecord(name string, rType uint16, class uint16, ttl uint32, rdLength uint16, rData []byte) ResourceRecord {
	parsedRData, _ := parseRData(rType, rData)
	return ResourceRecord{
		Name:        name,
		Type:        rType,
		Class:       class,
		TTL:         ttl,
		RDLength:    rdLength,
		RData:       rData,
		RDataParsed: parsedRData,
	}
}

func parseRData(rType uint16, rData []byte) (string, error) {
	if rType == TypeA {
		return parseA(rData)
	}

	return "", fmt.Errorf("unknown resource type: %d", rType)
}

func parseA(rData []byte) (string, error) {
	if len(rData) != 4 {
		return "", fmt.Errorf("invalid A record length: %d", rData)
	}

	ip := net.IP(rData)
	return ip.String(), nil
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

func ResourceRecordFromBytes(data []byte) ResourceRecord {
    name := "resource"
    if int(data[0]) >= 192 {
        name = "the same"
    }

	return ResourceRecord{
        Name: name,
    }
}

func decodeName(data []byte) (name string) {
    return ""
}
