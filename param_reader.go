package amugine

import (
	"io/ioutil"
	"strings"
)

type paramReader struct {
	Param string
}

func (pr *paramReader) GetValue() ([]byte, error) {
	if pr.isFile() {
		return pr.readFile()
	}
	return []byte(pr.Param), nil
}

func (pr *paramReader) isFile() bool {
	return strings.HasPrefix(pr.Param, "@")
}

func (pr *paramReader) readFile() ([]byte, error) {
	return ioutil.ReadFile(pr.Param[1:])
}
