package xor

import (
	"ABA/EME/app/methods"
	"io"
	"os"
)

type Detector struct {
	Result    bool
	Filepath  string
	ResultKey string
}

func (d *Detector) Detect() bool {
	header := GetFileHeader(d.Filepath)

	for i := byte(0); i < methods.AsciiLimit; i++ {
		result := make([]byte, 2)

		for index, char := range header {
			result[index] = i ^ char
		}

		if result[methods.PositionOfM] == methods.AsciiM && result[methods.PositionOfZ] == methods.AsciiZ {
			return true
		}
	}

	return false
}

func (d Detector) Present() {

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
