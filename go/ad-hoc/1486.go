package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for {
		// Read n, m, and p using bufio.Scanner
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		p, _ := strconv.Atoi(scanner.Text())

		if n == 0 && m == 0 && p == 0 {
			break
		}

		// Initialize a 2D array for the vetor
		vetor := make([][]int, m+1)
		for i := 1; i <= m; i++ {
			vetor[i] = make([]int, n+1)
		}

		// Read vetor values directly into the array
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				scanner.Scan()
				val, _ := strconv.Atoi(scanner.Text())
				vetor[i][j] = val
			}
		}

		resp := 0
		for i := 1; i <= n; i++ {
			jaforam := 0
			for j := 1; j <= m; j++ {
				if vetor[j][i] == 1 {
					jaforam++
				} else {
					if jaforam >= p {
						resp++
					}
					jaforam = 0
				}
			}
			if jaforam >= p {
				resp++
			}
		}

		fmt.Printf("%d\n", resp)
	}
}
