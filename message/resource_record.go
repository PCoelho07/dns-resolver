package message

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strings"
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
	parsedRData, _ := parseRData(rType, rData, []byte{})
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

func parseRData(rType uint16, rData []byte, fullData []byte) (string, error) {
	if rType == TypeA {
		return parseA(rData)
	}

	if rType == TypeCNAME {
		return parseCNAME(rData, fullData)
	}

	return "", fmt.Errorf("unknown resource type: %d", rType)
}

func parseCNAME(rData []byte, fullData []byte) (string, error) {
    if rData[0] >= 192 {
        offset := rData[1]
        return decodeName(fullData[offset:])
    }

    return decodeName(rData)
}

func decodeName(data []byte) (string, error) { 
    if len(data) <= 0 {
        return "", errors.New("resource record decode name: name segment is empty.")
    }

    result := []string{}

    i := 0
    for i < len(data) {
        sLength := int(data[i])

        if sLength == 0 { 
            break
        }

        labelBytes := data[(i+1) : (sLength+i+1)]
        label := []string{}
        for _, r := range labelBytes {
            label = append(label, string(rune(r)))
        }

        result = append(result, strings.Join(label, ""))
        i += sLength + 1
    }

    return strings.Join(result, "."), nil
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
