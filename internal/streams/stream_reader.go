package streams

import (
	"bufio"
	"io"
)

type StreamReader struct {
	name     string
	currByte byte
	currErr  error
	reader   io.ReadCloser
	source   *bufio.Reader
}

func NewStreamReader(rd io.ReadCloser, rdName string) *StreamReader {
	reader := bufio.NewReader(rd)
	currByte, currErr := reader.ReadByte()
	return &StreamReader{
		name:     rdName,
		currByte: currByte,
		currErr:  currErr,
		source:   reader,
		reader:   rd,
	}
}

func (s *StreamReader) Name() string {
	return s.name
}

func (s *StreamReader) Close() error {
	return s.reader.Close()
}

func (s *StreamReader) Peek() (byte, error) {
	return s.currByte, s.currErr
}

func (s *StreamReader) Next() (byte, error) {
	currByte, currErr := s.currByte, s.currErr
	s.currByte, s.currErr = s.source.ReadByte()
	return currByte, currErr
}
