package tool

import (
	"fmt"
	"os"

	"github.com/caiquetorres/wc/internal/options"
	"github.com/caiquetorres/wc/internal/stats"
	"github.com/caiquetorres/wc/internal/streams"
)

func Wc(opt *options.ToolOptions) ([]stats.Stat, error) {
	var readers []*streams.StreamReader
	var statArr []stats.Stat
	if opt.IsStdin() {
		reader := streams.NewStreamReader(os.Stdin, "stdin")
		defer reader.Close()
		readers = append(readers, reader)
	} else {
		for _, filePath := range opt.FilePaths {
			file, err := os.Open(filePath)
			if err != nil {
				return statArr, fmt.Errorf("no such file or directory")
			}
			reader := streams.NewStreamReader(file, file.Name())
			defer reader.Close()
			readers = append(readers, reader)
		}
	}
	for _, reader := range readers {
		stat := stats.NewStat(reader, opt)
		statArr = append(statArr, stat)
	}
	return statArr, nil
}
