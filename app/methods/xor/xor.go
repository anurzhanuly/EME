package xor

import (
	"bytes"
	"fmt"
	"github.com/anurzhanuly/EME/app/methods"
	headerUtil "github.com/anurzhanuly/EME/app/methods/utils/file"
)

type Detector struct {
	Result    bool
	Filepath  string
	ResultKey string
}

func (d *Detector) Detect() (bool, error) {
	fmt.Println("Starting XOR checking...")
	fmt.Println("Reading from file...")

	fileContents, err := headerUtil.GetFileWithoutHeader(d.Filepath)
	if err != nil {
		fmt.Println("ERROR: Cannot read from file")
		return false, err
	}

	return d.isXORed(fileContents), err
}

func (d *Detector) isXORed(fileContents []byte) bool {
	fmt.Print("Starting XOR brute-forcing with key: ")

	for i := byte(0); i < methods.AsciiLimit; i++ {
		fmt.Printf("%x, ", string(i))

		result := make([]byte, len(fileContents))

		for index, char := range fileContents {
			result[index] = i ^ char
		}

		if bytes.Index(result, []byte(methods.KeyWordForIdentification)) != -1 {
			d.ResultKey = string(i)
			fmt.Println()

			return true
		}
	}

	fmt.Println()
	return false
}

func (d Detector) Present() {
	fmt.Println("XOR - Encryption type was found")
	fmt.Printf("%s - is the key that was used", d.ResultKey)
	fmt.Println()
}
