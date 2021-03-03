// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 5, 6, 7}
	fmt.Println(primes)

	fmt.Println(primes[1:4])

	slice1 := primes[1:4]
	slice2 := primes[3:5]
	primes[3] = 100
	fmt.Println(primes, slice1, slice2)

}
