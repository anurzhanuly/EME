package methods

const PositionOfM = 0
const PositionOfZ = 1
const HexM = 0x4d
const HexZ = 0x5a
const AsciiM = 77
const AsciiZ = 90
const AsciiLimit = 255

type Detector interface {
	Detect() bool
	Present()
}
