package io

import (
	"io/ioutil"
)

const (
	startOfcomment = "//"
)

/*
FileReader is an interface for reading files
*/
type FileReader interface {
	Read(fileName string) ([]byte, error)
}

/*
DefaultFileReader is the base implementation of FileReader
*/
type DefaultFileReader struct {
}

/*
Read uses inside the ioutil.ReadFile function to read files
*/
func (sc *DefaultFileReader) Read(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}
