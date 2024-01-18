package main

import "fmt"

func main() {
	var valorInicial, valorFinal float64
	fmt.Scanf("%f %f", &valorInicial, &valorFinal)

	if valorInicial == valorFinal {
		fmt.Println("0.00%")
	} else {
		resultado := ((valorFinal * 100.0) / valorInicial) - 100
		fmt.Printf("%.2f%%\n", resultado)
	}
}
