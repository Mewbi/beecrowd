package main

import "fmt"

func main() {
	for {
		var tamanho uint16
		fmt.Scan(&tamanho)

		if tamanho == 0 {
			return
		}

		matriz := make([][]uint16, tamanho)
		for i := range matriz {
			matriz[i] = make([]uint16, tamanho)
		}

		for linha := uint16(0); linha < tamanho; linha++ {
			for coluna := uint16(0); coluna < tamanho; coluna++ {
				if linha == coluna {
					matriz[linha][coluna] = 1
				} else if linha < coluna {
					matriz[linha][coluna] = coluna - linha + 1
				} else {
					matriz[linha][coluna] = linha - coluna + 1
				}
			}
		}

		// Print the matrix
		for linha := uint16(0); linha < tamanho; linha++ {
			for coluna := uint16(0); coluna < tamanho; coluna++ {
				if coluna == 0 {
					fmt.Printf("%3d", matriz[linha][coluna])
				} else {
					fmt.Printf(" %3d", matriz[linha][coluna])
				}
			}
			fmt.Println()
		}

		fmt.Println()
	}
}
