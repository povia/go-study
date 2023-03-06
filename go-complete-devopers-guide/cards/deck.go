package main

import "fmt"

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

/**
 * 영상에서는 index, slice의 요소를 for의 변수로 사용
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
