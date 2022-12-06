package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var versionCompile string

//go:embed sample.jpg
var imageCompile []byte

//go:embed files/*.txt
var pathCompile embed.FS

func main() {
	fmt.Println(versionCompile)

	if err := ioutil.WriteFile("sample_copy.jpg", imageCompile, fs.ModePerm); err != nil {
		panic(err)
	}

	dirEntries, err := pathCompile.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			content, err := pathCompile.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))
		}
	}
}
