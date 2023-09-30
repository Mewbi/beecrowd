package main

import (
	"fmt"
)

func main() {
	var tamanho uint16
	var impar bool

	for {
		_, err := fmt.Scanf("%d", &tamanho)
		if err != nil {
			break
		}

		impar = false
		if tamanho%2 != 0 {
			impar = true
		}
		matriz := make([][]int, tamanho)
		for i := range matriz {
			matriz[i] = make([]int, tamanho)
		}

		for linha := uint16(0); linha < tamanho; linha++ {
			for coluna := uint16(0); coluna < tamanho; coluna++ {
				if (linha == coluna) && (linha == tamanho/2 && impar) {
					matriz[linha][coluna] = 2
				} else if linha == coluna {
					matriz[linha][coluna] = 1
				} else if linha == tamanho-coluna-1 {
					matriz[linha][coluna] = 2
				} else {
					matriz[linha][coluna] = 3
				}
			}
		}

		for linha := uint16(0); linha < tamanho; linha++ {
			for coluna := uint16(0); coluna < tamanho; coluna++ {
				fmt.Print(matriz[linha][coluna])
			}
			fmt.Println()
		}
	}
}
