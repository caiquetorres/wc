package tool

import (
	"fmt"
	"os"
	"testing"

	"github.com/caiquetorres/wc/internal/options"
	"github.com/stretchr/testify/assert"
)

func TestWcNoFilesProvided(t *testing.T) {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	r, w, _ := os.Pipe()
	defer r.Close()
	os.Stdin = r

	mockInput := "test input"
	w.WriteString(mockInput)
	w.Close()

	opt := options.NewToolOptions()
	_, err := Wc(opt)
	assert.NoError(t, err, "Expected no error when reading from stdin")
}

func TestWcSingleValidFile(t *testing.T) {
	tmpFile, _ := os.CreateTemp("", "testfile")
	defer os.Remove(tmpFile.Name())

	content := "test input"
	_, err := tmpFile.WriteString(content)
	tmpFile.Close()

	opt := options.NewToolOptions()
	opt.FilePaths = []string{tmpFile.Name()}

	_, err = Wc(opt)
	assert.NoError(t, err, "Expected no error when processing a valid file")
}

func TestWcMultipleFiles(t *testing.T) {
	files := []string{}
	for i := 0; i < 2; i++ {
		tmpFile, _ := os.CreateTemp("", fmt.Sprintf("testfile%d", i))
		defer os.Remove(tmpFile.Name())

		tmpFile.WriteString(fmt.Sprintf("content for file %d", i))
		tmpFile.Close()
		files = append(files, tmpFile.Name())
	}

	opt := options.NewToolOptions()
	opt.FilePaths = files

	_, err := Wc(opt)
	assert.NoError(t, err, "Expected no error when processing multiple files")
}

func TestWcFileNotFound(t *testing.T) {
	invalidPath := "nonexistentfile.txt"

	opt := options.NewToolOptions()
	opt.FilePaths = append(opt.FilePaths, invalidPath)

	_, err := Wc(opt)

	assert.Error(t, err, "Expected error for nonexistent file")
	assert.Contains(t, err.Error(), "no such file or directory", "Error message should mention file not found")
}

func TestWcWordCount(t *testing.T) {
	tmpFile, _ := os.CreateTemp("", "testfile")
	defer os.Remove(tmpFile.Name())

	content := "hello world\nthis is a test file\n"
	tmpFile.WriteString(content)
	tmpFile.Close()

	opt := options.NewToolOptions()
	opt.FilePaths = []string{tmpFile.Name()}
	opt.WordsCount = true

	statArr, err := Wc(opt)
	assert.Equal(t, 7, int(statArr[0].WordsCount))
	assert.NoError(t, err, "Expected no error when counting words")
}
