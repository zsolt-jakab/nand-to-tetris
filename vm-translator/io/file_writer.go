package io

import (
	"io/ioutil"
)

/*
FileWriter is an interface for creating and writing a file with a given name
*/
type FileWriter interface {
	Write(name string, data []byte) error
}

/*
DefaultFileWriter is the base implementation of FileWriter
*/
type DefaultFileWriter struct {
}

/*
Write writes the data to a file with the given file name, it creates the file if it does not exists.
It truncates the file if it already exists.
*/
func (sc *DefaultFileWriter) Write(name string, data []byte) error {
	return ioutil.WriteFile(name, data, 0644)
}
