package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"path"
	"testing"
)

//go:embed files/*.txt
var embedPath embed.FS

//go:embed files/*
var embedPathAll embed.FS

func TestPatchMatcher(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover when: ", r)
		}
	}()

	embedPathEntries, err := embedPath.ReadDir("files")
	if err != nil {
		panic(err)
	}
	for _, embed := range embedPathEntries {
		if embed.IsDir() {
			continue
		}
		data, err := embedPath.ReadFile(path.Join("files", embed.Name()))
		if err != nil {
			panic(err)
		}
		info, err := embed.Info()
		if err != nil {
			panic(err)
		}
		fmt.Println("Filename: ", embed.Name(), " | Info: ", info)
		fmt.Println(string(data))
		fmt.Println("")
	}

	embedPathJsonEntries, err := embedPathAll.ReadDir("files/json")
	if err != nil {
		panic(err)
	}
	for _, embed := range embedPathJsonEntries {
		if embed.IsDir() {
			continue
		}
		data, err := embedPathAll.ReadFile(path.Join("files", "json", embed.Name()))
		if err != nil {
			panic(err)
		}
		type dataJSONStruct struct {
			Name    string `json:"name"`
			Age     int    `json:"age"`
			Address string `json:"address"`
		}
		var dataJSON dataJSONStruct
		if err := json.Unmarshal(data, &dataJSON); err != nil {
			panic(err)
		}
		info, err := embed.Info()
		if err != nil {
			panic(err)
		}
		fmt.Println("Filename: ", embed.Name(), " | Info: ", info)
		fmt.Println(dataJSON)
		fmt.Println("")
	}
}
