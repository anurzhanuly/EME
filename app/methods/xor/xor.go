package xor

import (
	"ABA/EME/app/methods"
	headerUtil "ABA/EME/app/methods/utils/header"
)

type Detector struct {
	Result    bool
	Filepath  string
	ResultKey string
}

func (d *Detector) Detect() bool {
	header := headerUtil.GetFileHeader(d.Filepath)

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
