package main

import "fmt"

func isHappy(t1, t2, t3 int) bool {
	if t1 > t2 && t2 <= t3 {
		return true
	}

	if t1 < t2 && t2 >= t3 {
		return false
	}

	if t1 < t2 && t2 < t3 && t3-t2 < t2-t1 {
		return false
	}

	if t1 < t2 && t2 < t3 && t3-t2 >= t2-t1 {
		return true
	}

	if t1 > t2 && t2 > t3 && t2-t3 < t1-t2 {
		return true
	}

	if t1 > t2 && t2 > t3 && t2-t3 >= t1-t2 {
		return false
	}

	if t1 == t2 && t2 < t3 {
		return true
	}

	return false
}

func main() {
	var t1, t2, t3 int
	fmt.Scanf("%d %d %d", &t1, &t2, &t3)
	if isHappy(t1, t2, t3) {
		fmt.Println(":)")
	} else {
		fmt.Println(":(")
	}
}
