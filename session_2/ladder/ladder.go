package main

import (
	"fmt"
	"io"
	"math"
	// "github.com/pkg/profile"
)

func degreesToRadians(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func main() {
	// defer profile.Start(profile.CPUProfile).Stop()

	var side, angleDeg float64

	for {
		_, err := fmt.Scanf("%f %f", &side, &angleDeg)
		if err == io.EOF {
			break
		}
		break
	}

	angleRad := degreesToRadians(angleDeg)
	hypot := side / math.Sin(angleRad)

	fmt.Println(math.Ceil(hypot))
}
