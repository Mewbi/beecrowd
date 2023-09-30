package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	T := 0

	for scanner.Scan() {
		N, _ := strconv.Atoi(scanner.Text())

		if N == 0 {
			break
		}

		if T > 0 {
			fmt.Println()
		}

		totalX, totalY := 0.0, 0.0
		consumos := make(map[int]int)

		for i := 0; i < N; i++ {
			scanner.Scan()
			line := scanner.Text()
			values := splitToIntegers(line)

			X, Y := values[0], values[1]
			totalX += float64(X)
			totalY += float64(Y)

			if _, ok := consumos[Y/X]; ok {
				consumos[Y/X] += X
			} else {
				consumos[Y/X] = X
			}
		}

		T++

		fmt.Printf("Cidade# %d:\n", T)

		var keys []int
		for key := range consumos {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		var output []string
		for _, key := range keys {
			output = append(output, fmt.Sprintf("%d-%d", consumos[key], key))
		}

		fmt.Println(strings.Join(output, " "))

		fmt.Printf("Consumo medio: %.2f m3.\n", math.Trunc(totalY*100/totalX)/100)
	}
}

func splitToIntegers(s string) []int {
	var result []int
	values := strings.Fields(s)
	for _, v := range values {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}
	return result
}
