package main

import (
	"fmt"
)

func main() {
	for {
		var N int
		_, err := fmt.Scanf("%d", &N)
		if err != nil || N == 0 {
			break
		} else {
			j := 1
			p := 0
			q := 0
			r := 0
			i := N
			if N%2 == 0 {
				r = N / 2
			} else if N%2 == 1 {
				r = (N / 2) + 1
			}
			ara := make([][]int, N)
			for i := range ara {
				ara[i] = make([]int, N)
			}
			for c := 1; c <= r; c++ {
				for a := p; a < i; a++ {
					for b := q; b < i; b++ {
						ara[a][b] = j
					}
				}
				j++
				p++
				q++
				i--
			}
			for g := 0; g < N; g++ {
				for h := 0; h < N; h++ {
					if h == 0 {
						fmt.Printf("%3d", ara[g][h])
					} else {
						fmt.Printf(" %3d", ara[g][h])
					}
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
	}
}
