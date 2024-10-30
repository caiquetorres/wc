package streams

import (
	"fmt"
	"io"
)

const chunkSize = 1024

type StreamReader struct {
	name   string
	ptr    uint
	err    error
	size   uint
	buffer [chunkSize]byte
	reader io.ReadCloser
}

func NewStreamReader(rd io.ReadCloser, rdName string) *StreamReader {
	stream := StreamReader{
		name:   rdName,
		ptr:    0,
		err:    nil,
		size:   0,
		buffer: [chunkSize]byte{},
		reader: rd,
	}
	stream.err = stream.fill()
	return &stream
}

func (s *StreamReader) Name() string {
	return s.name
}

func (s *StreamReader) Close() error {
	return s.reader.Close()
}

func (s *StreamReader) Peek() (byte, error) {
	if int(s.ptr) < len(s.buffer) {
		return s.buffer[s.ptr], nil
	}
	return 0, s.err
}

func (s *StreamReader) Next() (byte, error) {
	// If the current position plus 1 exceeds the buffer size, refill the buffer.
	if s.ptr+1 > s.size {
		if err := s.fill(); err != nil {
			return 0, err
		}
	}
	ch, _ := s.Peek()
	s.ptr++
	return ch, nil
}

func (s *StreamReader) fill() error {
	n, err := s.reader.Read(s.buffer[:])
	// If err is not nil or no new data is read into the buffer, it indicates the
	// end of the input.
	if err != nil && err != io.EOF || n == 0 {
		return fmt.Errorf("end of stream")
	}
	s.size = uint(n)
	s.ptr = 0
	return nil
}
