package main

import (
	"fmt"
)

func main() {
	var tamanho int

	for {
		_, err := fmt.Scanf("%d", &tamanho)
		if err != nil || tamanho == 0 {
			break
		}

		matriz := make([][]int, tamanho)
		for i := range matriz {
			matriz[i] = make([]int, tamanho)
		}

		elemento := 1
		for linha := 0; linha < tamanho; linha++ {
			for coluna := 0; coluna < tamanho; coluna++ {
				matriz[linha][coluna] = elemento
				elemento *= 2
			}
			elemento = matriz[linha][0] * 2
		}

		digitos := 0
		t := matriz[tamanho-1][tamanho-1]

		// Count the digits in the largest number in the matrix
		for t > 0 {
			t /= 10
			digitos++
		}

		for linha := 0; linha < tamanho; linha++ {
			for coluna := 0; coluna < tamanho; coluna++ {
				if coluna == 0 {
					fmt.Printf("%*d", digitos, matriz[linha][coluna])
				} else {
					fmt.Printf(" %*d", digitos, matriz[linha][coluna])
				}
			}
			fmt.Println()
		}

		fmt.Println()
	}
}
