package main

import "fmt"

func roman(n int) string {
	dec := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	rom := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var romanNumber string
	var i int
	for {
		if n <= 0 {
			break
		}

		for {
			if n/dec[i] <= 0 {
				break
			}

			romanNumber += rom[i]
			n -= dec[i]
		}

		i++
	}

	return romanNumber
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(roman(n))
}
