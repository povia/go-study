package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	/**
	 * 영상에서는 index, slice의 요소를 for의 변수로 사용
	 for i, card := range cards {
		fmt.Println(i, card)
	 }
	TODO 왜 인덱스만 사용하는 지는 속도 테스트를 해봐야 할 듯 하다.
	*/
	for i := range cards {
		fmt.Println(i, cards[i])
	}

}

func newCard() string {
	return "Five of Diamonds"
}
