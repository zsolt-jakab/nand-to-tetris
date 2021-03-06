package main

import (
	"os"
	"path/filepath"

	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/io"
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/parser"
)

func main() {
	fileName := os.Args[1]
	doMain(fileName)
}

func doMain(fileNameWithPath string) {
	var reader io.FileReader = &io.DefaultFileReader{}
	var writer io.FileWriter = &io.DefaultFileWriter{}
	var sourceAccessor io.FileAccessor = &io.VMTranslatorFileAccessor{FileReader: reader, FileWriter: writer}

	codeLines, codeLineIndexes := sourceAccessor.ReadSourceLines(fileNameWithPath)

	binaryLines := parser.Translate(filepath.Base(fileNameWithPath), codeLines, codeLineIndexes)
	sourceAccessor.CreateTargetFile(fileNameWithPath, binaryLines)
}
