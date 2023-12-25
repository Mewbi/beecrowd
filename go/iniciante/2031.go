package main

import (
	"fmt"
)

func game(op1, op2 string) string {
	if (op1 == "ataque" && op2 == "pedra") ||
		(op1 == "ataque" && op2 == "papel") ||
		(op1 == "pedra" && op2 == "papel") {
		return "Jogador 1 venceu"
	}

	if (op2 == "ataque" && op1 == "pedra") ||
		(op2 == "ataque" && op1 == "papel") ||
		(op2 == "pedra" && op1 == "papel") {
		return "Jogador 2 venceu"
	}

	if op1 == "papel" && op2 == op1 {
		return "Ambos venceram"
	}

	if op1 == "pedra" && op2 == op1 {
		return "Sem ganhador"
	}

	return "Aniquilacao mutua"
}

func main() {
	var n int
	var op1, op2 string

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &op1)
		fmt.Scanf("%s", &op2)

		res := game(op1, op2)
		fmt.Println(res)
	}
}
