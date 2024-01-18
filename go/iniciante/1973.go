package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, i uint64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.ParseUint(scanner.Text(), 10, 64)

	star := make([]uint64, n)
	sheep := make([]uint64, n)
	var sumTrak, sumSheep uint64

	for i = 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.ParseUint(scanner.Text(), 10, 64)
		sheep[i] = v
		star[i] = 0
	}

	i = 0
	for {
		if i == 0 && sheep[i]%2 == 0 {
			star[i] = 1
			if sheep[i] > 0 {
				sheep[i]--
			}
			break
		} else if i == n-1 && sheep[i]%2 == 1 {
			star[i] = 1
			if sheep[i] > 0 {
				sheep[i]--
			}
			break
		} else if sheep[i]%2 == 1 {
			sheep[i]--
			star[i] = 1
			i++
		} else if sheep[i]%2 == 0 {
			star[i] = 1
			if sheep[i] > 0 {
				sheep[i]--
			}
			i--
		}
	}

	for i = 0; i < n; i++ {
		sumSheep += sheep[i]
		sumTrak += star[i]
	}

	fmt.Printf("%d %d\n", sumTrak, sumSheep)
}
