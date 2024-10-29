package stats

import (
	"bytes"
	"io"
	"testing"

	"github.com/caiquetorres/wc/internal/options"
	"github.com/caiquetorres/wc/internal/streams"
	"github.com/stretchr/testify/assert"
)

type mockReadCloser struct {
	io.Reader
	closed bool
}

func (m *mockReadCloser) Close() error {
	m.closed = true
	return nil
}

func (m *mockReadCloser) wasClosed() bool {
	return m.closed
}

func TestMergeStats(t *testing.T) {
	stats := []Stat{
		{ByteCount: 10, LineCount: 2, WordsCount: 3, CharsCount: 10, Name: "file1"},
		{ByteCount: 20, LineCount: 5, WordsCount: 6, CharsCount: 15, Name: "file2"},
	}
	expected := Stat{ByteCount: 30, LineCount: 7, WordsCount: 9, CharsCount: 25, Name: "total"}
	mergedStat := MergeStats(stats)

	assert.Equal(t, mergedStat, expected)
}

func TestNewStat_AllOptions(t *testing.T) {
	data := []byte("Just a test ã\n")
	mockReader := &mockReadCloser{Reader: bytes.NewReader(data)}
	stream := streams.NewStreamReader(mockReader, "testStream")

	op := options.NewToolOptions()
	op.LineCount = true
	op.WordsCount = true
	op.ByteCount = true
	op.CharsCount = true

	stat := NewStat(stream, op)

	assert.EqualValues(t, stat.LineCount, 1)
	assert.EqualValues(t, stat.WordsCount, 4)
	assert.EqualValues(t, stat.ByteCount, 15)
	assert.EqualValues(t, stat.CharsCount, 14)
}

func TestNewStat_NoOptions(t *testing.T) {
	data := []byte("Just a test\n")
	mockReader := &mockReadCloser{Reader: bytes.NewReader(data)}
	stream := streams.NewStreamReader(mockReader, "testStream")

	op := options.NewToolOptions()
	stat := NewStat(stream, op)

	assert.EqualValues(t, stat.LineCount, 1)
	assert.EqualValues(t, stat.WordsCount, 3)
	assert.EqualValues(t, stat.ByteCount, len(data))
}
