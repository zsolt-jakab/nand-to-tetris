package main

import (
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/io"
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/parser"
	"os"
)

func main() {
	filePath := os.Args[1]
	doMain(filePath)
}

func doMain(filePath string) {
	fileInfo, _ := os.Stat(filePath)
	var reader io.FileReader = &io.DefaultFileReader{}
	var writer io.FileWriter = &io.DefaultFileWriter{}
	var access io.FileAccess = &io.DefaultFileAccess{}
	var sourceAccessor io.FileAccessor = &io.VMTranslatorFileAccessor{FileReader: reader, FileWriter: writer, FileAccess: access}

	codeLines := sourceAccessor.ReadSourceLines(filePath)

	binaryLines := parser.Translate(codeLines, fileInfo.IsDir())
	sourceAccessor.CreateTargetFile(filePath, binaryLines)
}
