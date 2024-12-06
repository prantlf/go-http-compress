package compress

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

var raw = []byte(`raw text for testing compression encodings, which has to be very long,
because some encoding have a significant overhead and make sense to be applied for input
with the size at least a hundred bytes, apart from snappy and s2, which are efficient only
if the input contains repeated parts, not like this text`)

func TestGZIP(t *testing.T) {
	testCompression(t, GZIP, true)
}

func TestDEFLATE(t *testing.T) {
	testCompression(t, DEFLATE, true)
}

func TestBROTLI(t *testing.T) {
	testCompression(t, BROTLI, true)
}

func TestZSTD(t *testing.T) {
	testCompression(t, ZSTD, true)
}

func TestSNAPPY(t *testing.T) {
	testCompression(t, SNAPPY, false)
}

func TestS2(t *testing.T) {
	testCompression(t, S2, false)
}

func testCompression(t *testing.T, encoding string, smallerOutput bool) {
	var compressed bytes.Buffer
	cw, err := NewWriter(&compressed, encoding, -1)
	if err != nil {
		t.Fatal(err)
	}
	cw.Write(raw)
	cw.Close()
	if smallerOutput {
		assert.Less(t, compressed.Len(), len(raw))
	}

	cr, err := NewReader(&compressed, encoding)
	if err != nil {
		t.Fatal(err)
	}
	defer cr.Close()
	out, err := io.ReadAll(cr)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, raw, out)
}
