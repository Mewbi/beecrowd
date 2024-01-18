package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		name := strings.Split(scanner.Text(), " ")
		scanner.Scan()
		values := strings.Split(scanner.Text(), " ")

		n1, _ := strconv.Atoi(values[0])
		n2, _ := strconv.Atoi(values[1])

		if (n1+n2)%2 == 0 && name[1] == "PAR" {
			fmt.Println(name[0])
		} else if (n1+n2)%2 != 0 && name[1] == "IMPAR" {
			fmt.Println(name[0])
		} else {
			fmt.Println(name[2])
		}
	}
}
