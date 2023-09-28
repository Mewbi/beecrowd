package main

import "fmt"

func printMatrix(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := i + 1
			y := j + 1

			if x > n/2 {
				x = n - x + 1
			}

			if y > n/2 {
				y = n - y + 1
			}

			v := x
			if v > y {
				v = y
			}

			whiteSpace := 4
			if j == 0 {
				whiteSpace = 3
			}

			whiteSpace -= ((v / 10) + 1)
			var output string
			for k := 0; k < whiteSpace; k++ {
				output += fmt.Sprintf(" ")
			}
			output += fmt.Sprintf("%d", v)
			fmt.Printf(output)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var n int

	for {
		fmt.Scanf("%d", &n)

		if n == 0 {
			break
		}

		printMatrix(n)
	}
}
