package main

import "fmt"

func main() {
	var s, t, f int
	fmt.Scanf("%d %d %d", &s, &t, &f)

	final := s + t + f
	if final >= 24 {
		final -= 24
	} else if final < 0 {
		final += 24
	}

	fmt.Println(final)
}
