package dataProvider

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/tykoon/pkg/knnModel"
)

type dataFileReader struct {
	_sourceFileName string
	_attributeCount int
}

func NewDataFileReader(sourse string, attrCount int) (*dataFileReader, error) {
	_, err := os.Stat(sourse)
	if err != nil {
		return nil, err
	}
	return &dataFileReader{sourse, attrCount}, nil
}

func (r *dataFileReader) ReadAllEntities() (res []knnModel.Entity, err error) {
	res = make([]knnModel.Entity, 0)

	file, err := os.Open(r._sourceFileName)
	if err != nil {
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var entryStr knnModel.Entity

	for fileScanner.Scan() {
		entryStr, err = ParseEntityString(fileScanner.Text(), r._attributeCount)
		if err == nil {
			res = append(res, entryStr)
		}
	}
	err = fileScanner.Err()
	return
}

func ParseEntityString(str string, n int) (res knnModel.Entity, err error) {
	items := strings.Split(str, ",")

	if len(items) != n+1 {
		err = errors.New("wrong source string format")
		return
	}

	res.ClassName = items[0]
	var temp float64

	for i := 1; i <= n; i++ {
		_, err = fmt.Sscan(items[i], &temp)
		if err != nil {
			return res, errors.New("wrong source string format")
		}
		res.AttributeValues = append(res.AttributeValues, temp)
	}
	return
}
