package message

import (
	"slices"
	"testing"
)

func TestHeaderToBytesParsing(t *testing.T) {
	hf := NewHeaderFlag(true, 1, true, true, true, true, 0, 0)
	h := NewHeader(10, hf, 1, 1, 1, 1)
	hBytes := h.ToBytes()

	expectedResult := []byte{0, 10, 143, 128, 0, 1, 0, 1, 0, 1, 0, 1}
	if !slices.Equal(hBytes, expectedResult) {
		t.Errorf("got %d, want %d", hBytes, expectedResult)
	}
}

func TestHeaderFromBytesParsing(t *testing.T) {
    hBytes := []byte{0, 10, 143, 128, 0, 1, 0, 1, 0, 1, 0, 1}
    h := HeaderFromBytes(hBytes)

    if h.ID != 10 {
        t.Error("fail when parse header from bytes")
    }
} 
