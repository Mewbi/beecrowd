package main

import (
	"fmt"
)

func matrix(tamanho int) {
	m := make([][]int, tamanho)
	for i := range m {
		m[i] = make([]int, tamanho)
	}

	for linha := 0; linha < tamanho; linha++ {
		for coluna := 0; coluna < tamanho; coluna++ {
			if linha == coluna {
				m[linha][coluna] = 2
			} else if linha == tamanho-coluna-1 {
				m[linha][coluna] = 3
			} else {
				m[linha][coluna] = 0
			}
		}
	}

	start := tamanho / 3
	end := tamanho - start

	for linha := start; linha < end; linha++ {
		for coluna := start; coluna < end; coluna++ {
			m[linha][coluna] = 1
		}
	}

	m[tamanho/2][tamanho/2] = 4

	for linha := 0; linha < tamanho; linha++ {
		for coluna := 0; coluna < tamanho; coluna++ {
			fmt.Print(m[linha][coluna])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var tamanho int

	for {
		_, err := fmt.Scanf("%d", &tamanho)
		if err != nil {
			break
		}

		matrix(tamanho)
	}
}
