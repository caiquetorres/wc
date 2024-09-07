package main

import (
	"log"
	"os"

	"github.com/caiquetorres/wc/internal"
	"github.com/caiquetorres/wc/internal/stats"
	streams "github.com/caiquetorres/wc/internal/stream"
)

func main() {
	options := internal.ReadOptions()
	var readers []*streams.StreamReader
	if len(options.FilePaths) == 0 {
		reader := streams.NewStdinStreamReader()
		defer reader.Close()
		readers = append(readers, reader)
	} else {
		for _, filePath := range options.FilePaths {
			reader, err := streams.NewStreamReader(filePath)
			defer reader.Close()
			if err != nil {
				stderr := new(log.Logger)
				stderr.SetOutput(os.Stderr)
				stderr.Println("ccwc error: no such file or directory")
				return
			}
			readers = append(readers, reader)
		}
	}
	statsArr := stats.NewStatsArr(readers, options)
	for _, stats := range statsArr {
		stats.Print(options)
	}
}
