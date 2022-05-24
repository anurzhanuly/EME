package rot13

import (
	"ABA/EME/app/methods"
	headerUtil "ABA/EME/app/methods/utils/file"
	"bytes"
	"fmt"
)

type Detector struct {
	Filepath  string
	Result    bool
	ResultKey int
}

func (d *Detector) Detect() (bool, error) {
	fileContent, err := headerUtil.GetFileWithoutHeader(d.Filepath)
	if err != nil {
		return false, err
	}
	fmt.Println("Checking flags for ROT13")

	return d.isROTed(fileContent), err
}

func (d *Detector) isROTed(fileContent []byte) bool {
	result := make([]byte, len(fileContent))

	for i := 1; i < methods.AsciiLettersLength; i++ {
		for index, val := range fileContent {
			result[index] = Rot13(val, i)
		}
	}

	return false
}

func Rot13(b byte, shift int) byte {
	pos := bytes.IndexByte([]byte(methods.AsciiUppercase), b)
	if pos != -1 {
		return methods.AsciiUppercase[(pos+shift)%methods.AsciiLettersLength]
	}

	pos = bytes.IndexByte([]byte(methods.AsciiLowercase), b)
	if pos != -1 {
		return methods.AsciiLowercase[(pos+shift)%methods.AsciiLettersLength]
	}

	return b
}

func (d Detector) Present() {

}
