package message

import (
	"slices"
	"testing"
)

func TestQuestionEncodeNameParsing(t *testing.T) {
    q := encodeName("google.com")
    encodedName := []byte{6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0}

    if !slices.Equal(q, encodedName) {
        t.Errorf("got %d, want %d", q, encodedName)
    }
}

func TestQuestionToByteParsing(t *testing.T) {
    q := NewQuestion("google.com", TypeA, ClassIN)
    qBytes := q.ToBytes()
    expectedQBytes := []byte{6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0, 0, 1, 0, 1}
    
    if !slices.Equal(qBytes, expectedQBytes) {
        t.Errorf("got %d, want %d", qBytes, expectedQBytes)
    }
}
