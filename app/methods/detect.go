package methods

const (
	PositionOfM              = 0
	PositionOfZ              = 1
	HexM                     = 0x4d
	HexZ                     = 0x5a
	AsciiM                   = 77
	AsciiZ                   = 90
	AsciiLimit               = 255
	KeyWordForIdentification = "This program"
	TrimLengthForExeHeader   = 90
	AsciiUppercase           = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AsciiLowercase           = "abcdefghijklmnopqrstuvwxyz"
	AsciiLettersLength       = len(AsciiLowercase)
)

type Detector interface {
	Detect() (bool, error) // Basic algorithm for finding the encryption type
	Present()              // Beautiful results of operation
}
