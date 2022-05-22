package methods

type detector interface {
	Detect(filePath string) bool
}
