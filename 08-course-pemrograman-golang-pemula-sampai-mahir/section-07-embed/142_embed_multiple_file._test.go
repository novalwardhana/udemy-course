package belajar_golang_embed

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed files/data1.txt
//go:embed files/data2.txt
//go:embed files/data3.txt
var files embed.FS

func TestEmbedMultipleFiles(t *testing.T) {
	data1, err := files.ReadFile("files/data1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data1))

	data2, err := files.ReadFile("files/data2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data2))

	data3, err := files.ReadFile("files/data3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data3))
}
