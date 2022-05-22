package xor

import (
	"ABA/EME/app/methods"
	"fmt"
	"io"
	"os"
)

func Detect(filepath string) bool {
	header := GetFileHeader(filepath)

	fmt.Println("Getting initial header: ", header)

	for i := byte(0); i < methods.AsciiLimit; i++ {
		result := make([]byte, 2)

		for index, char := range header {
			result[index] = i ^ char
		}

		if result[methods.PositionOfM] == methods.AsciiM && result[methods.PositionOfZ] == methods.AsciiZ {
			fmt.Printf("The key is: %s", string(i))
			fmt.Println()

			return true
		}
	}

	return false
}

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
