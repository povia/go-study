package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for i := range cardSuits {
		for j := range cardValues {
			cards = append(cards, cardValues[j]+" of "+cardSuits[i])
		}
	}
	return cards
}

/*
*
  - 영상에서는 index, slice의 요소를 for의 변수로 사용
    for i, card := range cards {
    fmt.Println(i, card)
    }

TODO 왜 인덱스만 사용하는 지는 속도 테스트를 해봐야 할 듯 하다.
*/
func (d deck) print() {
	for i := range d {
		fmt.Println(i, d[i])
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// ioutil 의 WriteFile 이 deprecated 여서 os 패키지의 WriteFile 사용
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
