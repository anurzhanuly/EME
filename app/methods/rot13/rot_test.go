package rot13

import (
	"bytes"
	"testing"
)

func TestRot13(t *testing.T) {
	input := []byte("Lbh penpxrq gur pbqr!")
	output := []byte("You cracked the code!")
	result := make([]byte, len(input))

	for index, value := range input {
		result[index] = Rot13(value, 13)
	}

	if bytes.Contains(result, output) {
		t.Log("KRASAVA! ROT PASHET")
	}
}
