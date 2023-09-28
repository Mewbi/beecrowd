package main

import (
	"fmt"
)

func winner(p1, p2 string) string {
	if p1 == p2 {
		return "none"
	}

	if p1 == "tesoura" {
		if p2 == "papel" {
			return p1
		}

		if p2 == "lagarto" {
			return p1
		}

		if p2 == "pedra" {
			return p2
		}

		if p2 == "Spock" {
			return p2
		}
	}

	if p1 == "papel" {
		if p2 == "pedra" {
			return p1
		}

		if p2 == "Spock" {
			return p1
		}

		if p2 == "lagarto" {
			return p2
		}

		if p2 == "tesoura" {
			return p2
		}
	}

	if p1 == "pedra" {
		if p2 == "tesoura" {
			return p1
		}

		if p2 == "lagarto" {
			return p1
		}

		if p2 == "papel" {
			return p2
		}

		if p2 == "Spock" {
			return p2
		}
	}

	if p1 == "lagarto" {
		if p2 == "Spock" {
			return p1
		}

		if p2 == "papel" {
			return p1
		}

		if p2 == "pedra" {
			return p2
		}

		if p2 == "tesoura" {
			return p2
		}
	}

	if p1 == "Spock" {
		if p2 == "pedra" {
			return p1
		}

		if p2 == "tesoura" {
			return p1
		}

		if p2 == "papel" {
			return p2
		}

		if p2 == "lagarto" {
			return p2
		}
	}

	return "none"
}

func main() {
	var n int
	var op1, op2 string

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &op1)
		fmt.Scanf("%s", &op2)

		win := winner(op1, op2)
		if win == "none" {
			fmt.Printf("Caso #%d: De novo!\n", i+1)
			continue
		}

		if win == op1 {
			fmt.Printf("Caso #%d: Bazinga!\n", i+1)
			continue
		}

		if win == op2 {
			fmt.Printf("Caso #%d: Raj trapaceou!\n", i+1)
			continue
		}
	}
}
