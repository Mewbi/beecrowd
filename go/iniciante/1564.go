package main

import "fmt"

func main() {
	var n int
	for {
		if _, err := fmt.Scanf("%d", &n); err != nil {
			break
		}

		out := "vai ter copa!"
		if n > 0 {
			out = "vai ter duas!"
		}
		fmt.Println(out)
	}
}
