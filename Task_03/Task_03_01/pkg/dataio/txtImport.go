package dataio

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"example.com/tykoon/pkg/usersystem"
)

type UsersFReader struct {
	_sourceFilePath string
	_errorOut       io.StringWriter
}

func NewUsersFReader(path string, errOut io.StringWriter) *UsersFReader {
	return &UsersFReader{
		_sourceFilePath: path,
		_errorOut:       errOut,
	}
}

type emptyStruct struct{}

func (r *UsersFReader) ImportAll() (res usersystem.UserList, err error) {
	resMap := make(map[string]emptyStruct)
	res = make([]*usersystem.User, 0)

	file, err := os.Open(r._sourceFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	user, lineCount, ok := &usersystem.User{}, 0, true

	for fileScanner.Scan() {
		str := fileScanner.Text()
		lineCount++

		user, ok = usersystem.NewUserFromString(str)
		if !ok {
			errorMsg := fmt.Sprintf("%d %s\n", lineCount, str)
			r._errorOut.WriteString(errorMsg)
			continue
		}
		if _, found := resMap[user.Email]; found {
			errorMsg := fmt.Sprintf("%d %s -- duplicate\n", lineCount, str)
			r._errorOut.WriteString(errorMsg)
			continue
		}
		resMap[user.Email] = emptyStruct{}
		res = append(res, user)
	}
	err = fileScanner.Err()
	return
}
