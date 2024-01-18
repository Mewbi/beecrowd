package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, d int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	n, _ = strconv.Atoi(input)

	for i := 0; i < n; i++ {
		scanner.Scan()
		d, _ = strconv.Atoi(scanner.Text())

		dateConv := d - 2015

		if dateConv > 0 {
			fmt.Printf("%d A.C.\n", dateConv+1)
		} else if dateConv < 0 {
			fmt.Printf("%d D.C.\n", -dateConv)
		} else {
			fmt.Printf("1 A.C.\n")
		}
	}
}
