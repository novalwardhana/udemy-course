package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed files/data1.txt
var embedString1 string

//go:embed files/data2.txt
var embedString2 string

//go:embed files/data3.txt
var embedString3 string

//go:embed files/data4.png
var embedString4 string

func TestEmbedFileString(t *testing.T) {
	fmt.Println(embedString1)
	fmt.Println(embedString2)
	fmt.Println(embedString3)
	fmt.Println(embedString4)

	arrData := []string{
		embedString1,
		embedString2,
		embedString3,
		embedString4,
	}
	fmt.Println(arrData)
}
