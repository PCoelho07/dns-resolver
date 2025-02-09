package message

import (
	"bytes"
	"encoding/binary"
)

type DnsMessage struct {
	Header HeaderType
    Question QuestionType
    Answer ResourceRecord
}

func NewMessage(query string, queryType uint16) *DnsMessage {
    question := NewQuestion(query, queryType, ClassIN) 
    headerFlag := HeaderFlag{
    	QR:     false,
    	OpCode: 0,
    	AA:     false,
    	TC:     false,
    	RD:     false,
    	RA:     false,
    	Z:      0,
    	RCode:  0,
    }
    header := NewHeader(22, headerFlag, 1, 0, 0, 0)
    answer := ResourceRecord{} 

    return &DnsMessage{
        Question: question,
        Header: header,
        Answer: answer,
    }
}

func (dnsMessage *DnsMessage) ToBytes() []byte {
    buffer := new(bytes.Buffer)

    binary.Write(buffer, binary.BigEndian, dnsMessage.Header)
    binary.Write(buffer, binary.BigEndian, dnsMessage.Question)

    return buffer.Bytes()
}

func (dnsMessage *DnsMessage) ReadFromBytes(message []byte) ([]byte, error) {
   return []byte("response"), nil
}
