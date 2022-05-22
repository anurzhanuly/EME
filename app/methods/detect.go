package methods

const HexM = 0x4d
const HexZ = 0x5a
const AsciiM = 77
const AsciiZ = 90
const AsciiLimit = 255

type detector interface {
	Detect(filepath string) bool
}
