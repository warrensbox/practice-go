package main

import (
	"fmt"
	"math/rand"
)

func main() {
	cards := []int{5, 4, 2, 6, 8, 9, 10, 5, 7, 9, 3, 6, 9, 12, 14, 24, 34, 50}
	shuffleCards(cards)
	fmt.Println(cards)
}

func shuffleCards(cards []int) {
	for i := 0; i < len(cards); i++ {
		k := rand.Intn(i + 1)
		cards[k], cards[i] = cards[i], cards[k]
	}
}
