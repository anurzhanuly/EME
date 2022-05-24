package rot13

import (
	"bytes"
	"fmt"
	"github.com/anurzhanuly/EME/app/methods"
	headerUtil "github.com/anurzhanuly/EME/app/methods/utils/file"
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
	fmt.Println("Checking for ROT13 cipher")

	return d.isROTed(fileContent), err
}

func (d *Detector) isROTed(fileContent []byte) bool {
	result := make([]byte, len(fileContent))

	fmt.Print("keys used to shift the letter:")

	for i := 1; i < methods.AsciiLettersLength; i++ {
		fmt.Printf("%d, ", i)

		for index, val := range fileContent {
			result[index] = Rot13(val, i)
		}

		if bytes.Contains(result, []byte(methods.KeyWordForIdentification)) {
			d.ResultKey = i

			return true
		}
	}

	fmt.Println()

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
	fmt.Println("ROT13 - is the encryption type")
	fmt.Printf("%d - is the used key for ROT13", d.ResultKey)
	fmt.Println()
}
