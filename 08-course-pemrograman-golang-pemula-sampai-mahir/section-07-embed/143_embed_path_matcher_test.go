package belajar_golang_embed

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed files/*.txt
var path embed.FS

func TestEmbedPathMatcher(t *testing.T) {
	dir, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		if !entry.IsDir() {
			content, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))
		}
	}
}
