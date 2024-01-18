package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPossible(a, b, c, d int) bool {
	if (a < b+c && b < a+c && c < a+b) ||
		(a < b+d && b < a+d && d < a+b) ||
		(a < c+d && c < a+d && d < a+c) ||
		(b < c+d && c < b+d && d < b+c) {
		return true
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")

	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])
	d, _ := strconv.Atoi(input[3])

	if isPossible(a, b, c, d) {
		fmt.Println("S")
		return
	}
	fmt.Println("N")
}
