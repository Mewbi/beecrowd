package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")

	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])

	fmt.Println(math.Max(float64(a), float64(b)))
}
