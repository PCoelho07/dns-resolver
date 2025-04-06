package message

import (
	"bytes"
	"encoding/binary"
)

type DnsMessage struct {
	Header    HeaderType
	Questions []QuestionType
	Answers   []ResourceRecord
}

func NewMessage(questions []QuestionType) *DnsMessage {
	headerFlag := HeaderFlag{
		QR:     true,
		OpCode: 1,
		AA:     false,
		TC:     false,
		RD:     false,
		RA:     true,
		Z:      3,
		RCode:  10,
	}
	header := NewHeader(10, headerFlag, 1, 0, 1, 0)
	answers := []ResourceRecord{}

	return &DnsMessage{
		Header:    header,
		Questions: questions,
		Answers:   answers,
	}
}

func (dnsMessage *DnsMessage) DecodeFromBytes(data []byte) (*HeaderType, error) {
    headerResult, _ := dnsMessage.Header.DecodeFromBytes(data)
    return headerResult, nil
}

func (dnsMessage *DnsMessage) ToBytes() []byte {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.BigEndian, dnsMessage.Header.ToBytes())

	for _, q := range dnsMessage.Questions {
		binary.Write(buffer, binary.BigEndian, q.ToBytes())
	}

	for _, a := range dnsMessage.Answers {
		binary.Write(buffer, binary.BigEndian, a.ToBytes())
	}

	return buffer.Bytes()
}

func (dnsMessage *DnsMessage) ReadFromBytes(message []byte) ([]byte, error) {
	return []byte("response"), nil
}
