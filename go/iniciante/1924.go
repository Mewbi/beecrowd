package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
	}
	fmt.Println("Ciencia da Computacao")
}
