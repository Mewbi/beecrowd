package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	var count, n int
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if count >= 3 {
			break
		}

		scanner.Scan()
		input = strings.Trim(scanner.Text(), " ")

		if input == "caw caw" {
			fmt.Println(n)
			count += 1
			n = 0
			continue
		}

		for i, ch := range input {
			if string(ch) == "*" {
				n += int(math.Pow(2, float64(len(input)-1)-float64(i)))
			}
		}
	}
}
