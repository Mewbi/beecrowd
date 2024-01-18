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
		v, _ := strconv.Atoi(scanner.Text())
		if v%2 == 0 {
			fmt.Println(0)
			continue
		}
		fmt.Println(1)
	}
}
