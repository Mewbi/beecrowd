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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		if input[0] == "Thor" {
			fmt.Println("Y")
		} else {
			fmt.Println("N")
		}
	}
}
