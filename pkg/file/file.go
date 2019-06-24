package file

import (
	"io/ioutil"
)

func ReadFile(path string) ([]byte, error) {
	file_byte, err := ioutil.ReadFile(path)
	return file_byte, err
}
