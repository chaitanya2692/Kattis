package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
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

func dot(x, y []int) (r int, err error) {
	if len(x) != len(y) {
		return 0, errors.New("incompatible lengths")
	}
	for i, xi := range x {
		r += xi * y[i]
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

		sort.Sort(sort.Reverse(sort.IntSlice(scalar1)))
		sort.Sort(sort.IntSlice(scalar2))

		fmt.Printf("%#v\n", scalar1)
		fmt.Printf("%#v\n", scalar2)

		d, err := dot(scalar1, scalar2)
		if err != nil {
			log.Fatal(err)
		}
		myScalars[i] = d
	}

	for j := 0; j < cases; j++ {
		fmt.Printf("Case #%d: %d\n", j+1, myScalars[j])
	}
}
