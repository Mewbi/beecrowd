package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	x := make([]int, n)

	scanner.Scan()
	input = scanner.Text()

	strNumbers := strings.Split(input, " ")

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(strNumbers[i])
		x[i] = num
	}

	minimum := x[0]
	result := 1

	for i := 1; i < n; i++ {
		if x[i] < minimum {
			minimum = x[i]
			result = i + 1
		}
	}

	fmt.Println(result)
}
