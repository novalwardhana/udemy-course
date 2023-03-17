package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//go:embed files/data1.txt
var embedCompileData1 string

//go:embed files/data2.txt
var embedCompileData2 string

//go:embed files/data3.txt
var embedCompileData3 string

//go:embed files/data4.png
var embedCompileData4 string

//go:embed files/json/data1.json
var embedCompileByte1 []byte

//go:embed files/json/data2.json
var embedCompileByte2 []byte

//go:embed files/json/data3.json
var embedCompileByte3 []byte

//go:embed files/json/data4.json
var embedCompileByte4 []byte

//go:embed files/*
var embedCompileMultiple embed.FS

func main() {
	fmt.Println(embedCompileData1)
	fmt.Println(embedCompileData2)
	fmt.Println(embedCompileData3)
	fmt.Println(embedCompileData4)

	var arrByte = [][]byte{
		embedCompileByte1,
		embedCompileByte2,
		embedCompileByte3,
		embedCompileByte4,
	}
	for index, dataByte := range arrByte {
		newIndex := index + 1
		filename := fmt.Sprintf("data%d_copy.json", newIndex)
		dataString := string(dataByte)
		dataString = strings.TrimSpace(dataString)
		dataString = strings.ReplaceAll(dataString, "\n", "")
		if err := ioutil.WriteFile(path.Join("files", "json", filename), []byte(dataString), os.ModePerm); err != nil {
			panic(err)
		}
	}

	embedEntries, err := embedCompileMultiple.ReadDir("files")
	if err != nil {
		panic(err)
	}
	for _, embed := range embedEntries {
		fmt.Println(embed.Name(), embed.IsDir())
		if embed.IsDir() {
			embedChildEntries, err := embedCompileMultiple.ReadDir(path.Join("files", embed.Name()))
			if err != nil {
				panic(err)
			}
			for _, embedChild := range embedChildEntries {
				fmt.Println(embedChild.Name())
				data, err := embedCompileMultiple.ReadFile(path.Join("files", embed.Name(), embedChild.Name()))
				if err != nil {
					panic(err)
				}
				fmt.Println(string(data))
				fmt.Println("")
			}
		} else {
			data, err := embedCompileMultiple.ReadFile(path.Join("files", embed.Name()))
			if err != nil {
				panic(err)
			}
			fmt.Println(embed.Name())
			fmt.Println(data)
			fmt.Println("")
		}
	}
}
