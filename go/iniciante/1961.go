package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var j, n, h, prevH, t int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	j, _ = strconv.Atoi(input[0])
	n, _ = strconv.Atoi(input[1])

	scanner.Scan()
	input = strings.Split(scanner.Text(), " ")
	prevH, _ = strconv.Atoi(input[0])

	win := true
	for i := 1; i < n; i++ {
		h, _ = strconv.Atoi(input[i])

		t = h - prevH
		if t < 0 {
			t *= -1
		}

		if t > j {
			win = false
			break
		}
		prevH = h
	}

	if win {
		fmt.Println("YOU WIN")
		return
	}
	fmt.Println("GAME OVER")
}
