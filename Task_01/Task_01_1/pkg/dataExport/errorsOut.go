package dataExport

import (
	"fmt"
	"os"
)

type ErrorFileWriter struct {
	_outputFilePath string
}

func NewErrorFileWriter(outputFile string) (res *ErrorFileWriter, err error) {
	file, err := os.Create(outputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return &ErrorFileWriter{_outputFilePath: outputFile}, nil
}

func (w *ErrorFileWriter) WriteString(msg string) (int, error) {
	file, err := os.OpenFile(w._outputFilePath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	n, err := fmt.Fprint(file, msg)
	return n, err
}
