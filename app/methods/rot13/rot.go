package rot13

import (
	headerUtil "ABA/EME/app/methods/utils/file"
	"fmt"
)

type Detector struct {
	Filepath string
	Result   bool
}

func (d *Detector) Detect() bool {
	header := headerUtil.GetFileHeader(d.Filepath)
	fmt.Println("Checking flags for ROT13")

	return d.isROTed(header)
}

func (d *Detector) isROTed(header [2]byte) bool {

	return false
}

func rot13(r rune) rune {
	if r >= 'a' && r <= 'z' {
		if r >= 'm' {
			return r - 13
		} else {
			return r + 13
		}
	} else if r >= 'A' && r <= 'Z' {
		if r >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}

	return r
}

func (d Detector) Present() {

}
