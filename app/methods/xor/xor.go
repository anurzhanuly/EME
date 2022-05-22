package xor

import (
	"ABA/EME/app/methods"
	"io"
	"os"
)

func Detect(filepath string) bool {
	header := getFileHeader(filepath)

	for i := byte(0); i < methods.AsciiLimit; i++ {
		var result [2]byte

		for _, char := range header {
			result[0] = i ^ char
		}

		if result[0] == methods.AsciiM && result[1] == methods.AsciiZ {
			return true
		}
	}

	return false
}

func getFileHeader(filepath string) [2]byte {
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
