package captcha

import (
	"bufio"
	"os"
	"testing"
)

func TestDraw(t *testing.T) {
	s := "1234"
	data, err := DrawToImg(s)
	if err != nil {
		t.Error("Error: " + err.Error())
	}
	outFile, err := os.Create("out.png")
	if err != nil {
		t.Error("Create file error: " + err.Error())
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	_, err = b.Write(data)
	if err != nil {
		t.Error("Write error: " + err.Error())
	}
	err = b.Flush()
	if err != nil {
		t.Error("Flush error: " + err.Error())
	}
}
