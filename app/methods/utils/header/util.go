package header

import (
	"io"
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
