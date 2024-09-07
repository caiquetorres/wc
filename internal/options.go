package internal

import (
	"flag"
	"maps"
	"slices"
)

type ToolOptions struct {
	ByteCount  bool
	LineCount  bool
	WordsCount bool
	CharsCount bool
	FilePaths  []string
}

func ReadOptions() *ToolOptions {
	options := ToolOptions{
		ByteCount:  false,
		LineCount:  false,
		WordsCount: false,
		CharsCount: false,
		FilePaths:  make([]string, 0),
	}
	flag.BoolVar(&options.ByteCount, "c", false, "Get the bytes count")
	flag.BoolVar(&options.LineCount, "l", false, "Get the lines count")
	flag.BoolVar(&options.WordsCount, "w", false, "Get the words count")
	flag.BoolVar(&options.CharsCount, "m", false, "Get the chars count")
	flag.Parse()
	files := flag.Args()
	options.FilePaths = removeDuplicates(files)
	return &options
}

func (options *ToolOptions) IsStdin() bool {
	return len(options.FilePaths) == 0
}

func (options *ToolOptions) NoOptions() bool {
	return !options.ByteCount && !options.LineCount && !options.WordsCount && !options.CharsCount
}

func removeDuplicates[T comparable](arr []T) []T {
	type void struct{}
	set := make(map[T]void)
	for _, el := range arr {
		set[el] = void{}
	}
	return slices.Collect(maps.Keys(set))
}
