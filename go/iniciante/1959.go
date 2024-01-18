package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, l int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	n, _ = strconv.Atoi(input[0])
	l, _ = strconv.Atoi(input[1])
	fmt.Println(n * l)
}
