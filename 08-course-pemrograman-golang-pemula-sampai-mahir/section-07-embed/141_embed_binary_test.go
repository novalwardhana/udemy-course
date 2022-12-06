package belajar_golang_embed

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed sample.jpg
var sampleImage []byte

func TestByte(t *testing.T) {
	if err := ioutil.WriteFile("sample_copy.jpg", sampleImage, fs.ModePerm); err != nil {
		panic(err)
	}
}
