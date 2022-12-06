package belajar_golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println("Test data")
	fmt.Println(version)
}
