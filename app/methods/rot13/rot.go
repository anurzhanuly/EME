package rot13

type Detector struct {
	Filepath string
	Result   bool
}

func (d *Detector) Detect() bool {

	return false
}

func (d Detector) Present() {

}
