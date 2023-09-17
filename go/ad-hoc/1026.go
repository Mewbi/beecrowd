package main

import "fmt"

func main() {
	var x, y int

	for {
		_, err := fmt.Scanf("%d %d", &x, &y)
		if err != nil {
			break
		}

		fmt.Println(x ^ y)
	}
}
