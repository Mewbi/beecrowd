package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int

	for {
		_, err := fmt.Scanf("%d", &a)
		if err != nil || a == 0 {
			break
		}

		fmt.Scanf("%d%d", &b, &c)

		y := (float64(a) * float64(b) * 100) / float64(c)
		fmt.Printf("%.0f\n", math.Floor(math.Sqrt(y)))
	}
}
