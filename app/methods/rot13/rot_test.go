package rot13

import (
	"bytes"
	"github.com/anurzhanuly/EME/app/methods"
	headerUtil "github.com/anurzhanuly/EME/app/methods/utils/file"
	"os"
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
		t.Log("ROT13 algorithm is working!")
	}
}

func TestFileForROT(t *testing.T) {
	filepathOrigin := "/home/nurzhanuly/Documents/Personal/mona/AITU/diploma/petest/crackme.exe"
	newFilepath := "/home/nurzhanuly/Documents/Personal/mona/AITU/diploma/petest/xor/rotTestFile.exe"
	shift := 13

	fileContent, err := headerUtil.GetFileContent(filepathOrigin)
	if err != nil {
		t.Error(err)
	}

	for index, char := range fileContent {
		fileContent[index] = Rot13(char, shift)
	}

	newHeader := fileContent[:methods.TrimLengthForExeHeader]
	newHeader = append(newHeader, fileContent...)

	newFile, _ := os.Create(newFilepath)
	_, err = newFile.Write(newHeader)
	if err != nil {
		return
	}

	err = newFile.Close()
	if err != nil {
		return
	}
}
