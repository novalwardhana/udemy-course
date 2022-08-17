package main

import (
	"fmt"
	"io"
	"os"
)

type readFile struct{}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	read := readFile{}
	io.Copy(read, file)

	//io.Copy(os.Stdout, file)
}

func (readFile) Write(byteData []byte) (int, error) {
	fmt.Println("readFile: ", string(byteData))
	return len(byteData), nil
}
