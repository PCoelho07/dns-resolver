package message

import "testing"

func TestHeaderFlagBytesParsing(t *testing.T) {
	hf := NewHeaderFlag(true, 1, true, true, true, true, 0, 0)
	hfb := hf.ToBytes()

	expectedResult := uint16(36736)
	if hfb != expectedResult {
		t.Errorf("got %d, want %d", hfb, expectedResult)
	}
}

func TestHeaderFlagFromBytesParsing(t *testing.T) {
	plainH := []byte{1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0}
	h := HeaderFlagFromBytes(plainH)

	if h.QR != true && h.OpCode != 1 && h.AA != true && h.TC != true && h.RD != true && h.RA != true && h.Z != 0 && h.RCode != 0 {
        t.Error("fail when parse header flag from bytes")
	}
}
