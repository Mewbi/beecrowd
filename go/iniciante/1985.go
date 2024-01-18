package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	items := map[int]float64{
		1001: 1.5,
		1002: 2.5,
		1003: 3.5,
		1004: 4.5,
		1005: 5.5,
	}

	var n, item, q int
	var tot float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		item, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		q, _ = strconv.Atoi(scanner.Text())

		tot += items[item] * float64(q)
	}

	fmt.Printf("%.2f\n", tot)
}
