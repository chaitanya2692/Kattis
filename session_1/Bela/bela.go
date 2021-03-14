package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	dom := map[string]int{"A": 11, "K": 4, "Q": 3, "J": 20, "T": 10, "9": 14, "8": 0, "7": 0}
	notDom := map[string]int{"A": 11, "K": 4, "Q": 3, "J": 2, "T": 10, "9": 0, "8": 0, "7": 0}

	totalScore := 0

	var hands, score int

	var cardInfo, domSuit string

	for {
		_, err := fmt.Scanf("%d %s", &hands, &domSuit)
		if err == io.EOF {
			break
		}
		break
	}

	for i := 0; i < 4*hands; i++ {
		_, err := fmt.Scanf("%s", &cardInfo)
		if err == io.EOF {
			break
		}

		if strings.Split(cardInfo, "")[1] == domSuit {
			score = dom[strings.Split(cardInfo, "")[0]]
		} else {
			score = notDom[strings.Split(cardInfo, "")[0]]
		}
		totalScore = totalScore + score
	}
	fmt.Println(totalScore)
}
