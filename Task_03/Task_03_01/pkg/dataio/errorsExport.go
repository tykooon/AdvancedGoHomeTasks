package dataio

import (
	"fmt"
	"os"
)

type ErrorFWriter struct {
	_errorFilePath string
}

func NewErrorFWriter(fname string) (*ErrorFWriter, error) {
	if file, err := os.Create(fname); err != nil {
		return nil, err
	} else {
		file.Close()
		return &ErrorFWriter{_errorFilePath: fname}, nil
	}
}

func (w *ErrorFWriter) WriteString(msg string) (int, error) {
	file, err := os.OpenFile(w._errorFilePath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return fmt.Fprint(file, msg)
}
