package stats

import (
	"fmt"

	"github.com/caiquetorres/wc/internal"
	streams "github.com/caiquetorres/wc/internal/stream"
)

type Stats struct {
	byteCount  uint
	lineCount  uint
	wordsCount uint
	charsCount uint
	fileName   string
}

func NewStatsArr(readers []*streams.StreamReader, options *internal.ToolOptions) []*Stats {
	var statsArr []*Stats
	totalStats := newStats("total")
	for i, reader := range readers {
		var currStats *Stats
		if options.IsStdin() {
			currStats = newStats("stdin")
		} else {
			currStats = newStats(options.FilePaths[i])
		}
		prevByte := byte(0)
		currByte, err := reader.Next()
		for {
			if err != nil {
				break
			}
			if options.NoOptions() || options.LineCount {
				if currByte == '\n' {
					currStats.lineCount++
					totalStats.lineCount++
				}
			}
			if options.NoOptions() || options.WordsCount {
				if isWhitespace(currByte) && !isWhitespace(prevByte) {
					currStats.wordsCount++
					totalStats.wordsCount++
				}
			}
			if options.NoOptions() || options.ByteCount {
				currStats.byteCount++
				totalStats.byteCount++
			}
			if options.CharsCount {
				if (currByte & 0xC0) != 0x80 {
					currStats.charsCount++
					totalStats.charsCount++
				}
			}
			prevByte = currByte
			currByte, err = reader.Next()
		}
		if options.NoOptions() || options.WordsCount {
			if isWhitespace(currByte) && !isWhitespace(prevByte) {
				currStats.wordsCount++
				totalStats.wordsCount++
			}
		}
		statsArr = append(statsArr, currStats)
	}
	if len(statsArr) > 1 {
		statsArr = append(statsArr, totalStats)
	}
	return statsArr
}

func (stats *Stats) Print(options *internal.ToolOptions) {
	if options.NoOptions() || options.LineCount {
		fmt.Printf("% 8d", stats.lineCount)
	}
	if options.NoOptions() || options.WordsCount {
		fmt.Printf("% 8d", stats.wordsCount)
	}
	if options.NoOptions() || options.ByteCount {
		fmt.Printf("% 8d", stats.byteCount)
	}
	if options.CharsCount {
		fmt.Printf("% 8d", stats.charsCount)
	}
	fmt.Printf(" %s\n", stats.fileName)
}

func newStats(fileName string) *Stats {
	return &Stats{
		byteCount:  0,
		lineCount:  0,
		wordsCount: 0,
		charsCount: 0,
		fileName:   fileName,
	}
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\v' || char == '\f' || char == '\r'
}
