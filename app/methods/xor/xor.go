package xor

import (
	"ABA/EME/app/methods"
	headerUtil "ABA/EME/app/methods/utils/header"
	"fmt"
	"strings"
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
		fmt.Printf("%s, ", string(i))

		result := make([]byte, len(fileContents))

		for index, char := range fileContents {
			result[index] = i ^ char
		}

		strResult := string(result)

		if strings.Contains(strResult, methods.KeyWordForIdentification) {
			d.ResultKey = string(i)
			fmt.Println()

			return true
		}
	}

	fmt.Println()
	return false
}

func (d Detector) Present() {

}
