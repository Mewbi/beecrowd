package main

import "fmt"

func nonzero(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}

	x := n / 5
	y := n % 5

	potX := x % 4
	switch potX {
	case 0:
		potX = 1
	case 1:
		potX = 2
	case 2:
		potX = 4
	case 3:
		potX = 8
	}

	var lastY uint64
	switch y {
	case 0:
		lastY = 1
	case 1:
		lastY = 1
	case 2:
		lastY = 2
	case 3:
		lastY = 6
	case 4:
		lastY = 4
	default:
		lastY = 0
	}

	return potX * nonzero(x) * lastY
}

func main() {
	var n, result, instancia uint64
	instancia = 1

	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		fmt.Printf("Instancia %d\n", instancia)
		instancia++

		result = nonzero(n)
		for result >= 10 {
			result = result % 10
		}

		fmt.Printf("%d\n\n", result)
	}
}
