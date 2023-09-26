package main

import (
	"fmt"
	"math"
)

func kamenetsky(n float64) int {
	if n < 0 {
		return 0
	}

	if n <= 1 {
		return 1
	}

	x := ((n * math.Log10(n/math.E)) +
		(math.Log10(2*math.Pi*n) / 2.0))

	return int(math.Floor(x) + 1)
}

func main() {
	var n float64
	fmt.Scanf("%f", &n)
	fmt.Println(kamenetsky(n))
}
