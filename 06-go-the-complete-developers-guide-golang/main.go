package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func main() {
	myDeck := deck{"Arsenal", "Tottenham", "Barcelona"}
	myDeck.saveToFile("sample.txt")

	myDeckFile := newDeckFromFile("sample.txt")
	fmt.Println("Test: ", myDeckFile)

	myDeck.shuffle()
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	return err
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		fmt.Println(newPosition, i)
	}
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}
	data := string(bs)
	dataSplit := strings.Split(data, ",")
	return deck(dataSplit)
}
