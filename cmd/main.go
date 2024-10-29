package main

import (
	"log"
	"os"

	"github.com/caiquetorres/wc/internal/options"
	"github.com/caiquetorres/wc/internal/stats"
	"github.com/caiquetorres/wc/internal/tool"
)

var logger = log.New(os.Stderr, "ccwc error: ", 0)

func main() {
	opt := options.ReadOptions()
	statArr, err := tool.Wc(opt)
	if err != nil {
		logger.Println(err.Error())
	}
	for _, stat := range statArr {
		stat.Print(opt)
	}
	if len(statArr) > 1 {
		stat := stats.MergeStats(statArr)
		stat.Print(opt)
	}
}
