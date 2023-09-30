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
	text := "LIFE IS NOT A PROBLEM TO BE SOLVED"
	fmt.Println(text[:n])
}
