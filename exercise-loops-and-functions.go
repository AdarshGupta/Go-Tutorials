package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	times := 0
	var prevZ float64 = 0.0
	z := 1.0

	for times < 10 {
		fmt.Println(z)

		if z == prevZ {
			return z
		}

		if z*z == x {
			return z
		}

		prevZ = z
		z -= (z*z - x) / (2 * z)

		times += 1
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2000))
}
