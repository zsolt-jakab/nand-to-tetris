package main

import (
	"github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/analyzer"
	"github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/io"
	"os"
)

func main() {
	filePath := os.Args[1]
	doMain(filePath)
}

func doMain(filePath string) {
	var reader io.FileReader = &io.DefaultFileReader{}
	var writer io.FileWriter = &io.DefaultFileWriter{}

	var sourceAccessor io.FileAccessor = &io.DefaultFileAccessor{FileReader: reader, FileWriter: writer}
	var syntaxAnalyzer analyzer.SyntaxAnalyzer = *analyzer.NewSyntaxAnalyzer(&sourceAccessor)
	syntaxAnalyzer.Analyze(filePath)
}
