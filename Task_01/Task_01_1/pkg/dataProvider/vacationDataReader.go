package dataProvider

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"example.com/tykoon/pkg/model"
)

type VacationsCsvReader struct {
	_sourceFilePath string
	_dateFormat     string
	_errorOut       io.StringWriter
}

func NewVacationsCsvReader(path string, dateFormat string, errOut io.StringWriter) *VacationsCsvReader {
	return &VacationsCsvReader{
		_sourceFilePath: path,
		_dateFormat:     dateFormat,
		_errorOut:       errOut,
	}
}

func (r *VacationsCsvReader) GetAllVacations() (res model.VacationSlice, err error) {
	res = make([]model.Vacation, 0)

	file, err := os.Open(r._sourceFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	emp, lineCount := &model.Vacation{}, 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		lineCount++

		emp, err = ParseVacationEntry(str, r._dateFormat)
		if err != nil {
			errorMsg := fmt.Sprintf("%d %s\n", lineCount, str)
			r._errorOut.WriteString(errorMsg)
			continue
		}

		res = append(res, *emp)
	}

	err = fileScanner.Err()
	return
}

func ParseVacationEntry(str string, dateFormat string) (*model.Vacation, error) {
	res := &model.Vacation{}
	entries := strings.Split(str, ",")
	if len(entries) != 3 {
		return res, errors.New("too many entries")
	}
	names := strings.Split(entries[0], " ")
	if len(names) != 2 {
		return res, errors.New("too many names")
	}
	start, err := time.Parse(dateFormat, entries[1])
	if err != nil {
		return res, errors.New("wrong date format")
	}
	end, err := time.Parse(dateFormat, entries[2])
	if err != nil {
		return res, errors.New("wrong date format")
	}
	res.Name, res.LastName = names[0], names[1]
	res.Start, res.End = start, end

	return res, nil
}
