package streams

import (
	"bytes"
	"io"
	"testing"

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

func TestNewStreamReader(t *testing.T) {
	data := []byte("hello")
	mockReader := &mockReadCloser{Reader: bytes.NewReader(data)}
	stream := NewStreamReader(mockReader, "testStream")
	peekByte, peekErr := stream.Peek()

	assert.Equal(t, stream.name, "testStream")
	assert.Equal(t, peekByte, data[0])
	assert.Nil(t, peekErr)
}

func TestStreamReader_Name(t *testing.T) {
	stream := NewStreamReader(&mockReadCloser{Reader: bytes.NewReader([]byte("hello"))}, "testStream")
	assert.Equal(t, stream.Name(), "testStream")
}

func TestStreamReader_Peek(t *testing.T) {
	data := []byte("hello")
	stream := NewStreamReader(&mockReadCloser{Reader: bytes.NewReader(data)}, "testStream")
	peekByte, peekErr := stream.Peek()

	assert.Equal(t, peekByte, data[0])
	assert.Nil(t, peekErr)
}

func TestStreamReader_Next(t *testing.T) {
	data := []byte("hello")
	stream := NewStreamReader(&mockReadCloser{Reader: bytes.NewReader(data)}, "testStream")

	firstByte, firstErr := stream.Next()

	assert.Equal(t, firstByte, data[0])
	assert.Nil(t, firstErr)

	secondByte, secondErr := stream.Next()
	assert.Equal(t, secondByte, data[1])
	assert.Nil(t, secondErr)
}

func TestStreamReader_Next_EndOfStream(t *testing.T) {
	data := []byte("hi")
	stream := NewStreamReader(&mockReadCloser{Reader: bytes.NewReader(data)}, "testStream")

	stream.Next() // "h"
	stream.Next() // "i"

	endByte, endErr := stream.Next()

	assert.NotNil(t, endErr)
	assert.NotEqual(t, endByte, 0)
}

func TestStreamReader_Close(t *testing.T) {
	mockReader := &mockReadCloser{Reader: bytes.NewReader([]byte("hello"))}
	stream := NewStreamReader(mockReader, "testStream")

	err := stream.Close()
	assert.Nil(t, err)
	assert.True(t, mockReader.wasClosed())
}
