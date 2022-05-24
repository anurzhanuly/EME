package methods

const (
	PositionOfM              = 0
	PositionOfZ              = 1
	HexM                     = 0x4d
	HexZ                     = 0x5a
	AsciiM                   = 77
	AsciiZ                   = 90
	AsciiLimit               = 255
	KeyWordForIdentification = "This is a program"
	TrimLengthForExeHeader   = 90
)

type Detector interface {
	Detect() (bool, error)
	Present() // Beautiful results of operation
}
