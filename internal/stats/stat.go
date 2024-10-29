package stats

import (
	"fmt"

	"github.com/caiquetorres/wc/internal/options"
	"github.com/caiquetorres/wc/internal/streams"
)

type Stat struct {
	ByteCount  uint
	LineCount  uint
	WordsCount uint
	CharsCount uint
	Name       string
}

func newStat(name string) Stat {
	return Stat{
		ByteCount:  0,
		LineCount:  0,
		WordsCount: 0,
		CharsCount: 0,
		Name:       name,
	}
}

func MergeStats(statArr []Stat) Stat {
	totalStat := newStat("total")
	for _, stat := range statArr {
		totalStat.LineCount += stat.LineCount
		totalStat.WordsCount += stat.WordsCount
		totalStat.ByteCount += stat.ByteCount
		totalStat.CharsCount += stat.CharsCount
	}
	return totalStat
}

func NewStat(reader *streams.StreamReader, op *options.ToolOptions) Stat {
	stat := newStat(reader.Name())
	prevByte := byte(0)
	for {
		currByte, err := reader.Next()
		if err != nil {
			break
		}
		if op.NoOptions() || op.LineCount {
			if currByte == '\n' {
				stat.LineCount++
			}
		}
		if op.NoOptions() || op.WordsCount {
			if isWhitespace(currByte) && !isWhitespace(prevByte) {
				stat.WordsCount++
			}
		}
		if op.NoOptions() || op.ByteCount {
			stat.ByteCount++
		}
		if op.CharsCount {
			if (currByte & 0xC0) != 0x80 {
				stat.CharsCount++
			}
		}
		prevByte = currByte
	}
	if op.NoOptions() || op.WordsCount {
		if !isWhitespace(prevByte) {
			stat.WordsCount++
		}
	}
	return stat
}

func (s *Stat) Print(op *options.ToolOptions) {
	if op.NoOptions() || op.LineCount {
		fmt.Printf("% 8d", s.LineCount)
	}
	if op.NoOptions() || op.WordsCount {
		fmt.Printf("% 8d", s.WordsCount)
	}
	if op.NoOptions() || op.ByteCount {
		fmt.Printf("% 8d", s.ByteCount)
	}
	if op.CharsCount {
		fmt.Printf("% 8d", s.CharsCount)
	}
	fmt.Printf(" %s\n", s.Name)
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\v' || char == '\f' || char == '\r'
}
