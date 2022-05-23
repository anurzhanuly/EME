package rot13

import headerUtil "ABA/EME/app/methods/utils/header"

type Detector struct {
	Filepath string
	Result   bool
}

func (d *Detector) Detect() bool {
	header := headerUtil.GetFileHeader(d.Filepath)

	return d.isROTed(header)
}

func (d *Detector) isROTed(header [2]byte) bool {

	return false
}

func (d Detector) Present() {

}
