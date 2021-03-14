package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PerfOp performs addition or subtraction
func PerfOp(numOne int, numTwo int, op string) (int, error) {
	switch op {
	case "+":
		return numOne + numTwo, nil
	case "-":
		return numOne - numTwo, nil
	}
	return 0, errors.New("invalid")
}

func getInitResult(equation []string, values map[string]int) (int, error) {
	numOne, ok1 := values[equation[0]]
	numTwo, ok2 := values[equation[2]]

	if ok1 && ok2 {
		if result, err := PerfOp(numOne, numTwo, equation[1]); err == nil {
			return result, nil
		}
	}
	return 0, errors.New("invalid")
}

func calc(equation []string, values map[string]int) (int, error) {
	result, err := getInitResult(equation, values)

	if err == nil && len(equation) > 4 {
		for i := 4; i < len(equation); i = i + 2 {
			if val, ok := values[equation[i]]; ok {
				result, _ = PerfOp(result, val, equation[i-1])
			} else {
				return 0, errors.New("invalid")
			}
		}
	}
	return result, err
}

func main() {

	var usrInput []string
	var dict = map[string]int{}
	var revDict = map[int]string{}
	scanner := bufio.NewScanner(os.Stdin)
	myOutput := list.New()

forLoop:
	for {
		if scanner.Scan() {
			usrInput = strings.Split(scanner.Text(), " ")
		}

		op := usrInput[0]

		switch op {

		case "def":
			intVal, _ := strconv.Atoi(usrInput[2])
			dict[usrInput[1]] = intVal
			revDict[intVal] = usrInput[1]

		case "calc":
			result, err := calc(usrInput[1:], dict)

			if strVal, ok := revDict[result]; ok && err == nil {
				outStr := strings.Join(usrInput[1:], " ") + " " + strVal
				myOutput.PushBack(outStr)
			} else {
				outStr := strings.Join(usrInput[1:], " ") + " " + "unknown"
				myOutput.PushBack(outStr)
			}

		case "clear":
			for e := myOutput.Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
			break forLoop
		}
	}
}
