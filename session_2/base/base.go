package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func sum(array []int) int {
// 	result := 0
// 	for _, v := range array {
// 		result += v
// 	}
// 	return result
// }

// func sliceAtoi(sa []string) ([]int, error) {
// 	si := []int{}
// 	for _, a := range sa {
// 		i, err := strconv.Atoi(a)
// 		if err != nil {
// 			return si, err
// 		}
// 		si = append(si, i)
// 	}
// 	return si, nil
// }

// func check1(list []string) bool {
// 	for _, b := range list {
// 		if b != "1" {
// 			return false
// 		}
// 	}
// 	return true
// }

func val(c string) int {
	if c[0] >= "0"[0] && c[0] <= "9"[0] {
		return int(c[0] - "0"[0])
	}
	return int(c[0] - "A"[0] - 22)
}

func toDeci(Inputnbr string, base int) int {
	len := len(Inputnbr)
	res3 := strings.SplitN(Inputnbr, "", len)
	power := 1 // Initialize power of base
	num := 0   // Initialize result

	for i := len - 1; i >= 0; i-- {
		if base == 1 {
			if val(res3[i]) > base {
				return -1
			}
		}
		if base > 1 {
			if val(res3[i]) >= base {
				return -1
			}
		}
		num += val(res3[i]) * power
		power = power * base
	}

	return num
}

func main() {

	var usrInputs, equation string
	var usrInputInt int
	AllBases := list.New()

	// var aInt, bInt, cInt int

	BaseDict := map[int]string{
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
		16: "g",
		17: "h",
		18: "i",
		19: "j",
		20: "k",
		21: "l",
		22: "m",
		23: "n",
		24: "o",
		25: "p",
		26: "q",
		27: "r",
		28: "s",
		29: "t",
		30: "u",
		31: "v",
		32: "w",
		33: "x",
		34: "y",
		35: "z",
		36: "0"}

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		usrInputs = scanner.Text()
		usrInputInt, _ = strconv.Atoi(usrInputs)

	}

	for i := 0; i < usrInputInt; i++ {
		if scanner.Scan() {
			equation = scanner.Text()
		}
		var bases []string
		myInput := strings.Split(equation, " ")

		// a, b, c := strings.Split(myInput[0], ""), strings.Split(myInput[2], ""), strings.Split(myInput[4], "")

		// if check1(a) {
		// 	if check1(b) {
		// 		if check1(c) {
		// 			aArr, _ := sliceAtoi(a)
		// 			bArr, _ := sliceAtoi(b)
		// 			cArr, _ := sliceAtoi(c)
		// 			aInt := sum(aArr)
		// 			bInt := sum(bArr)
		// 			cInt := sum(cArr)
		// 			if myInput[1] == "+" {
		// 				if cInt == aInt+bInt {
		// 					bases = append(bases, "1")
		// 				}
		// 			}
		// 			if myInput[1] == "-" {
		// 				if cInt == aInt-bInt {
		// 					bases = append(bases, "1")
		// 				}
		// 			}
		// 			if myInput[1] == "*" {
		// 				if cInt == aInt*bInt {
		// 					bases = append(bases, "1")
		// 				}
		// 			}
		// 			if myInput[1] == "/" {
		// 				if cInt == aInt/bInt {
		// 					bases = append(bases, "1")
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		for CurrentBase := 0; CurrentBase < 37; CurrentBase++ {
			aInt := toDeci(myInput[0], CurrentBase)
			bInt := toDeci(myInput[2], CurrentBase)
			cInt := toDeci(myInput[4], CurrentBase)

			if aInt != -1 && bInt != -1 && cInt != -1 {
				if myInput[1] == "+" {
					if cInt == aInt+bInt {
						bases = append(bases, BaseDict[CurrentBase])
					}
				}
				if myInput[1] == "-" {
					if cInt == aInt-bInt {
						bases = append(bases, BaseDict[CurrentBase])
					}
				}
				if myInput[1] == "*" {
					if cInt == aInt*bInt {
						bases = append(bases, BaseDict[CurrentBase])
					}
				}
				if myInput[1] == "/" {
					if float32(cInt) == float32(aInt)/float32(bInt) {
						bases = append(bases, BaseDict[CurrentBase])
					}
				}
			}
		}

		if len(bases) > 0 {
			AllBases.PushBack(strings.Join(bases, ""))
		} else {
			AllBases.PushBack("invalid")
		}

	}
	for e := AllBases.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
