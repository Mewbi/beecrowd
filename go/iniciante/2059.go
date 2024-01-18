package main

import "fmt"

const (
	PAR   = 1
	IMPAR = 0
)

func main() {
	var p, j1, j2, r, a, j1Choose, v int
	fmt.Scanf("%d %d %d %d %d", &p, &j1, &j2, &r, &a)

	if p == PAR {
		j1Choose = PAR
	} else {
		j1Choose = IMPAR
	}

	if r == 1 && a == 1 {
		v = 2
	} else if r == 1 && a == 0 {
		v = 1
	} else {
		res := (j1 + j2) % 2

		if res == 0 && j1Choose == PAR {
			v = 1
		} else if res == 1 && j1Choose == IMPAR {
			v = 1
		} else {
			v = 2
		}
	}

	fmt.Printf("Jogador %d ganha!\n", v)
}
