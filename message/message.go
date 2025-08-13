package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
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
		Header:        header,
		Questions:     []QuestionType{*questions},
		Answers:       answers,
		AuthoritysRRs: authorityRR,
		AdditionalRRs: additionalRR,
	}
}

func (dnsMessage *DnsMessage) DnsMessageFromBytes(data []byte) (*DnsMessage, error) {
	headerResult := HeaderFromBytes(data[0:12])
	rrOffset := len(dnsMessage.Questions[0].ToBytes()) + 12

    fmt.Println("questions bytes length ", rrOffset)
    fmt.Println("data rrOffset value", data[rrOffset:])

    answer := ResourceRecord{}
    rr := data[rrOffset:] 
    if rr[0] >= 192 { 
        n := dnsMessage.Questions[0].Name
        t := binary.BigEndian.Uint16(rr[2:4])
        c := binary.BigEndian.Uint16(rr[4:6])
        ttl := binary.BigEndian.Uint32(rr[6:10])
        rdLen := binary.BigEndian.Uint16(rr[10:12])
        rData, err := parseRData(t, rr[12:])
        if err != nil {
            log.Fatalf("parse R data fails: %v", err)
        }

        answer = ResourceRecord{
            Name: n,
            Type: t,
            Class: c,
            TTL: ttl,
            RDLength: rdLen,
            RData: []byte(rData),
            RDataParsed: rData,
        }
    }

	return &DnsMessage{
		Header:  headerResult,
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
