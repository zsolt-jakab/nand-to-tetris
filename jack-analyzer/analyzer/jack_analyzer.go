package analyzer

import (
	"github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/analyzer/compeng"
	"github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/analyzer/tokenizer"
	"github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/io"
	"strings"
)

type SyntaxAnalyzer struct {
	fileAccessor io.FileAccessor
}

func NewSyntaxAnalyzer(fileAccessor *io.FileAccessor) *SyntaxAnalyzer {
	return &SyntaxAnalyzer{fileAccessor: *fileAccessor}
}

func (sa *SyntaxAnalyzer) Analyze(fileSystemEntry string) {
	filePaths, _ := sa.fileAccessor.FindPaths(fileSystemEntry, "*.jack")

	for _, filePath := range filePaths {
		var tokenizer tokenizer.Tokenizer = tokenizer.NewJackTokenizer(sa.fileAccessor.ReadCode(filePath))
		CompilationEngine := compeng.NewJackCompilationEngine(&tokenizer)
		compiledXml := CompilationEngine.CompileClass()
		xmlName := strings.TrimRight(filePath, ".jack") + ".xml"

		sa.fileAccessor.CreateFileFromContent(xmlName, compiledXml)
	}

}
