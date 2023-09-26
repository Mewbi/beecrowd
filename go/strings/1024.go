package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func encript(text string) string {
	var newText string

	for _, c := range text {
		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			c += 3
		}
		newText += fmt.Sprintf("%c", c)
	}

	text = reverse(newText)
	newText = ""
	size := len(text) / 2

	for i, c := range text {
		if i >= size {
			c -= 1
		}
		newText += fmt.Sprintf("%c", c)
	}

	return newText
}

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var text string
		if scanner.Scan() {
			text = scanner.Text()
		}

		fmt.Println(encript(text))
	}
}
