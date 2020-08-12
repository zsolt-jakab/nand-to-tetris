package main

import (
	"os"

	"github.com/zsolt-jakab/nand-to-tetris/assembler/io"
	"github.com/zsolt-jakab/nand-to-tetris/assembler/parser"
)

func main() {
	fileName := os.Args[1]
	doMain(fileName)
}

func doMain(fileName string) {
	var reader io.FileReader = &io.DefaultFileReader{}
	var writer io.FileWriter = &io.DefaultFileWriter{}
	var sourceAccessor io.FileAccessor = &io.DefaultFileAccessor{FileReader: reader, FileWriter: writer}

	codeLines, codeLineIndexes := sourceAccessor.ReadCodeLines(fileName)

	binaryLines := parser.Translate(codeLines, codeLineIndexes)
	sourceAccessor.CreateHackFile(fileName, binaryLines)
}
