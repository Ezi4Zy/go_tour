package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		diff := (z*z - x) / (2 * z)
		if diff >= -0.00001 && diff <= 0.00001 {
			return z
		} else {
			z = z - diff
		}
	}

}

func Sqrt2(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt2(2))
}
