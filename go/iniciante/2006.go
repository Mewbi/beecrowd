package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var correct int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	input := scanner.Text()

	for i := 0; i < 5; i++ {
		scanner.Scan()
		item := scanner.Text()
		if input == item {
			correct += 1
		}
	}

	fmt.Println(correct)
}
