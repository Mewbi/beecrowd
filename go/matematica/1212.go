package main

import "fmt"

func NumCarryAdd(n1, n2 int64) int64 {
	var a, b, c, t int64
	c = 0
	t = 0
	for {
		a = n1 % 10
		b = n2 % 10
		n1 = n1 / 10
		n2 = n2 / 10
		if (a + b + c) >= 10 {
			t++
			c = 1
		} else {
			c = 0
		}
		if n1 == 0 && n2 == 0 {
			break
		}
	}
	return t
}

func main() {
	var x, y, carry int64
	for {
		_, err := fmt.Scanf("%d %d", &x, &y)
		if err != nil || (x == 0 && y == 0) {
			break
		}
		carry = NumCarryAdd(x, y)
		if carry == 0 {
			fmt.Println("No carry operation.")
		} else if carry == 1 {
			fmt.Println("1 carry operation.")
		} else {
			fmt.Printf("%d carry operations.\n", carry)
		}
	}
}
