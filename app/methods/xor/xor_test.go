package xor

import (
	"ABA/EME/app/methods"
	headerUtil "ABA/EME/app/methods/utils/file"
	"os"
	"testing"
)

func TestDetect(t *testing.T) {
	filepathOrigin := "/home/nurzhanuly/Documents/Personal/mona/AITU/diploma/petest/crackme.exe"
	newFilepath := "/home/nurzhanuly/Documents/Personal/mona/AITU/diploma/petest/xor/crackme.exe"
	key := byte('Z')
	detector := Detector{
		Filepath: filepathOrigin,
	}

	fileContent, err := headerUtil.GetFileContent(filepathOrigin)
	if err != nil {
		t.Error(err)
	}

	t.Log("Original file is:", fileContent[:10])

	for index, char := range fileContent {
		fileContent[index] = char ^ key
	}

	newHeader := make([]byte, methods.TrimLengthForExeHeader)
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

	if result, err := detector.Detect(); err != nil || !result {
		t.Error("XOR type is not found")
	}
}
