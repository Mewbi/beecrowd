package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var v, d, h, a float64
	scanner := bufio.NewScanner(os.Stdin)

	pi := 3.14
	for scanner.Scan() {
		v, _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		d, _ = strconv.ParseFloat(scanner.Text(), 64)

		a = math.Pow((d/2), 2) * pi
		h = v / a
		fmt.Printf("ALTURA = %.2f\n", h)
		fmt.Printf("AREA = %.2f\n", a)
	}
}
