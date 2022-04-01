package base32

import (
	"bytes"
	"testing"
)

var (
	testBuf = []byte{
		0xd8, 0x6b, 0x33, 0x92, 0xc1, 0x20, 0x2e, 0x8f,
		0xf5, 0xa4, 0x23, 0xb3, 0x02, 0xe6, 0x28, 0x4d,
		0xb7, 0xf8, 0xf4, 0x35, 0xea, 0x9f, 0x39, 0xb5,
		0xb1, 0xb2, 0x0f, 0xd3, 0xac, 0x36, 0xdf, 0xcb,
	}
	testBase32 = "1jyz6snd63xjn6skk7za6psgidsd53k05cr3lksqybi0q6936syq"
)

func TestEncodedLen(t *testing.T) {
	if expected, actual := len(testBase32), EncodedLen(len(testBuf)); expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestEncode(t *testing.T) {
	if expected, actual := testBase32, Encode(testBuf); expected != actual {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	buf := make([]byte, len(testBuf))
	if ok := Decode(buf, testBase32); !ok {
		t.Error("failed to decode")
	}
	if expected, actual := testBuf, buf; !bytes.Equal(expected, actual) {
		t.Errorf("expected %#v, got %#v", expected, actual)
	}
}
