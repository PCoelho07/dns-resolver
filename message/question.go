package message

import (
	"bytes"
	"encoding/binary"
	"strings"
)

type QuestionType struct {
    Name string
    QName []byte
    QType uint16
    QClass uint16
}

func NewQuestion(name string, qType, qClass uint16) *QuestionType {
    q := &QuestionType{
        Name: name,
        QType: qType,
        QClass: qClass,
    }

    q.QName = encodeName(name)
    return q
}

func encodeName(name string) []byte {
    var buffer bytes.Buffer
    labels := strings.Split(name, ".")

    for _, part := range labels {
        buffer.WriteByte(byte(len(part)))
        buffer.WriteString(part)
    }

    buffer.WriteByte(0)
    return buffer.Bytes()
}

func (q *QuestionType) ToBytes() []byte {
    buffer := new(bytes.Buffer)

    buffer.Write(q.QName)
    binary.Write(buffer, binary.BigEndian, q.QType)
    binary.Write(buffer, binary.BigEndian, q.QClass)

    return buffer.Bytes()
}
