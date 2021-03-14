package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getScalar() []int {
	var scalar []string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		scalar = strings.Split(scanner.Text(), " ")
	}

	ints := make([]int, len(scalar))
	for i, s := range scalar {
		ints[i], _ = strconv.Atoi(s)
	}

	return ints
}

func dot(x []int, y []int, scalarLength int) (r int) {
	for i, xi := range x {
		r += xi * y[scalarLength-(i+1)]
	}
	return
}

func main() {
	var cases int
	fmt.Scanln(&cases)

	myScalars := make([]int, cases)

	for i := 0; i < cases; i++ {
		var scalarLength int
		fmt.Scanln(&scalarLength)

		scalar1 := getScalar()
		scalar2 := getScalar()

		sort.Ints(scalar1)
		sort.Ints(scalar2)

		dotty := dot(scalar1, scalar2, scalarLength)
		myScalars[i] = dotty
	}

	for j := 0; j < cases; j++ {
		fmt.Printf("Case #%d: %d\n", j+1, myScalars[j])
	}
}
