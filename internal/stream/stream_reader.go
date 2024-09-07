package streams

import (
	"bufio"
	"os"
)

type StreamReader struct {
	fileName    string
	currentChar byte
	currentErr  error
	reader      *bufio.Reader
	file        *os.File
}

// Creates a new stream reader instance.
func NewStdinStreamReader() *StreamReader {
	reader := bufio.NewReader(os.Stdin)
	currentChar, currentErr := reader.ReadByte()
	return &StreamReader{
		fileName:    "",
		currentChar: currentChar,
		currentErr:  currentErr,
		reader:      reader,
		file:        os.Stdin,
	}
}

func NewStreamReader(filePath string) (*StreamReader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	currentChar, currentErr := reader.ReadByte()
	return &StreamReader{
		fileName:    filePath,
		currentChar: currentChar,
		currentErr:  currentErr,
		reader:      reader,
		file:        file,
	}, nil
}

func (stream *StreamReader) Close() error {
	return stream.file.Close()
}

func (stream *StreamReader) Peek() (byte, error) {
	return stream.currentChar, stream.currentErr
}

func (stream *StreamReader) Next() (byte, error) {
	currentChar, currentErr := stream.currentChar, stream.currentErr
	stream.currentChar, stream.currentErr = stream.reader.ReadByte()
	return currentChar, currentErr
}
