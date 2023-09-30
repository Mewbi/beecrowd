package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	x := make([]int, 10)
	y := 0

	for i := 0; i < 10; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &x[i])

		if x[i] > 0 {
			fmt.Printf("X[%d] = %d\n", y, x[i])
		} else {
			fmt.Printf("X[%d] = 1\n", y)
		}
		y++
	}
}
