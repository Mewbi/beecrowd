package main

import (
	"fmt"
	"math"
)

func main() {
	var r, pi, a float64
	pi = 3.14159
	fmt.Scanf("%f", &r)
	a = pi * math.Pow(r, 2)
	fmt.Printf("A=%.4f\n", a)
}
