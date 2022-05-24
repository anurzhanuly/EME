package xor

import (
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

	header, _ := headerUtil.GetFileHeader(filepathOrigin)
	t.Log("Original file is:", header)

	for index, char := range header {
		header[index] = char ^ key
	}

	newHeader := make([]byte, 2)
	newHeader[0] = header[0]
	newHeader[1] = header[1]

	t.Log("Modified file is:", header)

	newFile, _ := os.Create(newFilepath)
	_, err := newFile.Write(newHeader)
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
