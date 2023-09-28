package main

import (
	"fmt"
)

func euclideanDivision(a, b int) (int, int) {
	var q, r, limit int
	limit = b
	if b < 0 {
		limit = -b
	}
	for r = 0; r < limit; r++ {
		if (a-r)%b == 0 {
			q = (a - r) / b
			break
		}
	}
	return q, r
}

func main() {
	var a, b, q, r int
	fmt.Scanf("%d %d", &a, &b)

	q, r = euclideanDivision(a, b)

	fmt.Printf("%d %d\n", q, r)
}
