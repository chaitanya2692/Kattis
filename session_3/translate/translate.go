package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var translation = map[byte]string{
	97:  "@",
	98:  "8",
	99:  "(",
	100: "|)",
	101: "3",
	102: "#",
	103: "6",
	104: "[-]",
	105: "|",
	106: "_|",
	107: "|<",
	108: "1",
	109: "[]\\/[]",
	110: "[]\\[]",
	111: "0",
	112: "|D",
	113: "(,)",
	114: "|Z",
	115: "$",
	116: "']['",
	117: "|_|",
	118: "\\/",
	119: "\\/\\/",
	120: "}{",
	121: "`/",
	122: "2"}

func main() {

	var textLower string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		textLower = strings.ToLower(scanner.Text())
	}

	for i := 0; i < len(textLower); i++ {
		if textLower[i] > 96 && textLower[i] < 123 {
			fmt.Print(translation[textLower[i]])
		} else {
			fmt.Print(string(textLower[i]))
		}
	}

}
