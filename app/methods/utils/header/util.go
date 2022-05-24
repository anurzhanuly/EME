package header

import (
	"io"
	"io/ioutil"
	"os"
)

func GetFileHeader(filepath string) [2]byte {
	r, _ := os.Open(filepath)
	defer func(r *os.File) {
		err := r.Close()
		if err != nil {
			return
		}
	}(r)

	var header [2]byte
	_, _ = io.ReadFull(r, header[:])

	return header
}

func GetFileWithoutHeader(filepath string) ([]byte, error) {
	r, _ := os.Open(filepath)
	defer func(r *os.File) {
		err := r.Close()
		if err != nil {
			return
		}
	}(r)

	var fileContents []byte
	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return []byte{}, err
	}

	return fileContents[2:], err
}
