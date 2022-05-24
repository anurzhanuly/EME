package file

import (
	"ABA/EME/app/methods"
	"io/ioutil"
	"os"
)

func GetFileHeader(filepath string) ([2]byte, error) {
	fileContents, err := GetFileContent(filepath)
	if err != nil {
		return [2]byte{}, err
	}

	var header [2]byte
	header[0] = fileContents[0]
	header[1] = fileContents[1]

	return header, err
}

func GetFileWithoutHeader(filepath string) ([]byte, error) {
	result, err := GetFileContent(filepath)
	if err != nil {
		return result, err
	}

	return result[methods.TrimLengthForExeHeader:], err
}

func GetFileContent(filepath string) ([]byte, error) {
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

	return fileContents, err
}
