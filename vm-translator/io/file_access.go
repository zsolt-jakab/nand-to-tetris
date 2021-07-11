package io

import (
	"os"
)

/*
FileAccess is an interface for checking file info
*/
type FileAccess interface {
	Stat(fileName string) (os.FileInfo, error)
}

/*
DefaultFileAccess is the base implementation of FileAccess
*/
type DefaultFileAccess struct {
}

/*
Stat uses inside the os. Stat function to get fileInfo
*/
func (sc *DefaultFileAccess) Stat(fileName string) (os.FileInfo, error) {
	return os.Stat(fileName)
}
