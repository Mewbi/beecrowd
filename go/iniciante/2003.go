package main

import (
	"fmt"
	"time"
)

func main() {
	var input string

	timeMeet, _ := time.Parse("2006-01-02 15:04", "2006-01-02 8:00")
	for {
		var delay float64
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		wake := fmt.Sprintf("2006-01-02 %s", input)
		wakeUp, _ := time.Parse("2006-01-02 15:04", wake)

		wakeUp = wakeUp.Add(time.Minute * 60)

		diff := wakeUp.Sub(timeMeet)
		if diff.Minutes() > 0 {
			delay = diff.Minutes()
		}

		fmt.Printf("Atraso maximo: %.0f\n", delay)
	}
}
