package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var idx int

	scanner := bufio.NewScanner(os.Stdin)
	for {
		var n int
		var nums []string
		if !scanner.Scan() {
			break
		}

		n, _ = strconv.Atoi(strings.Trim(scanner.Text(), " "))
		idx++
		nums = append(nums, "0")
		for i := 0; i <= n; i++ {
			for j := 0; j < i; j++ {
				nums = append(nums, strconv.FormatInt(int64(i), 10))
			}
		}

		var plural string
		if len(nums) > 1 {
			plural = "s"
		}
		fmt.Printf("Caso %d: %d numero%s\n%s\n\n", idx, len(nums), plural, strings.Join(nums, " "))
	}
}
