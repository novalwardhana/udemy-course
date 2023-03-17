package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

//go:embed files/data1.txt
var embedByte1 []byte

//go:embed files/data2.txt
var embedByte2 []byte

//go:embed files/data3.txt
var embedByte3 []byte

//go:embed files/data4.png
var embedByte4 []byte

func TestEmbedFileByte(t *testing.T) {
	fmt.Println(embedByte1)
	fmt.Println(embedByte2)
	fmt.Println(embedByte3)
	fmt.Println(embedByte4)

	arrData := []string{
		string(embedByte1),
		string(embedByte2),
		string(embedByte3),
		string(embedByte4),
	}
	fmt.Println(arrData)

	arrDataByte := [][]byte{
		embedByte1,
		embedByte2,
		embedByte3,
		embedByte4,
	}
	for index, data := range arrDataByte {
		newIndex := index + 1
		switch newIndex {
		case 1, 2, 3:
			{
				filename := fmt.Sprintf("files/data%d_copy.txt", newIndex)
				if err := ioutil.WriteFile(filename, data, os.ModePerm); err != nil {
					panic(err)
				}
				fmt.Println("success")
			}
		case 4:
			{
				filename := fmt.Sprintf("files/data%d_copy.png", newIndex)
				if err := ioutil.WriteFile(filename, data, os.ModePerm); err != nil {
					panic(err)
				}
				fmt.Println("Success")
			}
		}
	}
}
