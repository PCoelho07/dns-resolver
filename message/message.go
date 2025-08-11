package message

import (
	"bytes"
	"encoding/binary"
)

type DnsMessage struct {
	Header        HeaderType
	Questions     []QuestionType
	Answers       []ResourceRecord
	AuthoritysRRs []ResourceRecord
	AdditionalRRs []ResourceRecord
}

func NewMessage(query string) *DnsMessage {
	headerFlag := NewHeaderFlag(
		false,
		0,
		false,
		false,
		true,
		false,
		0,
		0,
	)

	questions := NewQuestion(query, TypeA, ClassIN)
	header := NewHeader(10, headerFlag, 1, 0, 0, 0)

	answers := make([]ResourceRecord, 0)
	authorityRR := make([]ResourceRecord, 0)
	additionalRR := make([]ResourceRecord, 0)

	return &DnsMessage{
		Header:    header,
		Questions: []QuestionType{*questions},
		Answers:   answers,
        AuthoritysRRs: authorityRR,
        AdditionalRRs: additionalRR,
	}
}

func DnsMessageFromBytes(data []byte) (*DnsMessage, error) {
    headerResult := HeaderFromBytes(data[0:12])
    answer := ResourceRecordFromBytes(data)
	return &DnsMessage{
        Header: headerResult,
        Answers: []ResourceRecord{answer},
    }, nil
}

func (dnsMessage *DnsMessage) HasError() bool {
    return dnsMessage.Header.Flags.RCode != 0
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

	for _, arr := range dnsMessage.AuthoritysRRs {
		binary.Write(buffer, binary.BigEndian, arr.ToBytes())
	}

	for _, addRR := range dnsMessage.AdditionalRRs {
		binary.Write(buffer, binary.BigEndian, addRR.ToBytes())
	}

	return buffer.Bytes()
}

func (dnsMessage *DnsMessage) ReadFromBytes(message []byte) ([]byte, error) {
	return []byte("response"), nil
}
