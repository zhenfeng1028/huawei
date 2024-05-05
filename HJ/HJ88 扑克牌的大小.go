package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cards = map[string]int{
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"10":    10,
	"J":     11,
	"Q":     12,
	"K":     13,
	"A":     14,
	"2":     15,
	"joker": 16,
	"JOKER": 17,
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs := strings.Split(input.Text(), "-")
	card1, card2 := strs[0], strs[1]
	cards1, cards2 := strings.Split(card1, " "), strings.Split(card2, " ")
	if len(cards1) != len(cards2) {
		if card1 == "joker JOKER" || card2 == "joker JOKER" {
			fmt.Println("joker JOKER")
		} else if len(cards1) == 4 {
			fmt.Println(card1)
		} else if len(cards2) == 4 {
			fmt.Println(card2)
		} else {
			fmt.Println("ERROR")
		}
	} else {
		if cards[cards1[0]] > cards[cards2[0]] {
			fmt.Println(card1)
		} else {
			fmt.Println(card2)
		}
	}
}
