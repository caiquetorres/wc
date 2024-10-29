package options

import (
	"flag"
)

type ToolOptions struct {
	ByteCount  bool
	LineCount  bool
	WordsCount bool
	CharsCount bool
	FilePaths  []string
}

func NewToolOptions() *ToolOptions {
	return &ToolOptions{
		ByteCount:  false,
		LineCount:  false,
		WordsCount: false,
		CharsCount: false,
		FilePaths:  []string{},
	}
}

func ReadOptions() *ToolOptions {
	opts := NewToolOptions()
	flag.BoolVar(&opts.ByteCount, "c", false, "Count bytes")
	flag.BoolVar(&opts.LineCount, "l", false, "Count lines")
	flag.BoolVar(&opts.WordsCount, "w", false, "Count words")
	flag.BoolVar(&opts.CharsCount, "m", false, "Count characters")
	flag.Parse()
	opts.FilePaths = flag.Args()
	return opts
}

func (o *ToolOptions) IsStdin() bool {
	return len(o.FilePaths) == 0
}

func (o *ToolOptions) NoOptions() bool {
	return !o.ByteCount && !o.LineCount && !o.WordsCount && !o.CharsCount
}
