package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MIN_NOTE = 8.0

func main() {
	var n, idMax, id int
	var max, note float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		id, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		note, _ = strconv.ParseFloat(scanner.Text(), 64)

		if note > max {
			max = note
			idMax = id
		}
	}

	if max < MIN_NOTE {
		fmt.Println("Minimum note not reached")
		return
	}

	fmt.Println(idMax)
}
