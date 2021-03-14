package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main() {
	//var password string
	// fmt.Scanf("%s\n", &password)
	password := "iLnLnLeLb"
	spass := strings.Split(password, "")
	mypass := strings.Builder{}
	operations := strings.Builder{}

	for i := 0; i < len(spass); i++ {
		key := spass[i]
		fmt.Println(key)
		if strings.ToLower(key) == key || isNumeric(key) {
			if operations.Len() == 0 {
				mypass.WriteString(key)
			} else {
				ops := strings.Split(operations.String(), "")
				for _, op := range ops {
					switch op {
					case "L":
						fmt.Println("It's L")
					case "R":
						fmt.Println("It's R")
					case "B":
						fmt.Println("It's B")
					}
				}
				operations.Reset()
			}
		}
		if strings.ToUpper(key) == key {
			operations.WriteString(key)
		}
	}
	fmt.Println(mypass.String())
	fmt.Println(operations.String())
	mypass.Reset()
	operations.Reset()
}
