package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		if i < n-1 {
			fmt.Printf("Ho ")
			continue
		}
		fmt.Printf("Ho")
	}
	fmt.Println("!")
}
