package main

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed files/data1.txt
//go:embed files/data2.txt
//go:embed files/data3.txt
//go:embed files/data4.png
var embedMultipleFiles embed.FS

func TestEmbedMultipleFile(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover when: ", r)
		}
	}()

	file1, err := embedMultipleFiles.ReadFile("files/data1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(file1)

	file2, err := embedMultipleFiles.ReadFile("files/data2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file2))

	file3, err := embedMultipleFiles.ReadFile("files/data3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file3))

	file4, err := embedMultipleFiles.ReadFile("files/data4.png")
	if err != nil {
		panic(err)
	}
	fmt.Println(file4)

	file5, err := embedMultipleFiles.ReadFile("files/data5.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(file5)
}
