package main

import (
	"fmt"
	"math"
)

func main() {
	var nrStatues float64
	fmt.Scanf("%f\n", &nrStatues)

	fmt.Println(math.Ceil(math.Log2(nrStatues) + 1))
}
