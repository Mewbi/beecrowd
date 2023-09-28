package main

import "fmt"

func getLevel(v int) int {
	if v < 10 {
		return 1
	}

	if v < 20 {
		return 2
	}

	return 3
}

func main() {
	var n, v, level, l int

	for {
		if _, err := fmt.Scanf("%d", &n); err != nil {
			break
		}

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &v)
			l = getLevel(v)

			if l > level {
				level = l
			}
		}

		fmt.Println(level)
		level = 0
	}
}
