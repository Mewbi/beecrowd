package main

import (
	"fmt"
	"math"
)

func bhaskara(a, b, c float64) float64 {
	x := (math.Pow(b, 2)) - (4 * a * c)
	x = math.Sqrt(x)
	x1 := (-b + x) / (2 * a)
	x2 := (-b - x) / (2 * a)

	if x1 > 0 {
		return x1
	}

	return x2
}

func calc(h, p1, p2, a, v, pi, g float64) {
	vy := v * math.Sin(a*pi/180)
	vx := v * math.Cos(a*pi/180)

	t := bhaskara(-g/2, vy, h)

	alc := vx * t

	res := "NUCK"
	if (alc >= p1) && (alc <= p2) {
		res = "DUCK"
	}

	fmt.Printf("%.5f -> %s\n", alc, res)
}

func main() {
	var h, p1, p2, a, v float64
	var n int

	for {
		if _, err := fmt.Scanf("%f", &h); err != nil {
			break
		}
		if _, err := fmt.Scanf("%f %f", &p1, &p2); err != nil {
			break
		}
		if _, err := fmt.Scanf("%d", &n); err != nil {
			break
		}

		for i := 0; i < n; i++ {
			if _, err := fmt.Scanf("%f %f", &a, &v); err != nil {
				break
			}

			calc(h, p1, p2, a, v, 3.14159, 9.80665)
		}
	}
}
