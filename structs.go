// Package main provides ...
package main

import "fmt"

type Vector struct {
	X int
	Y int
}

func main() {
	v := Vector{1, 2}
	fmt.Println(v)
	v.X = 3
	fmt.Println(v)
	p := &v
	p.Y = 4
	fmt.Println(v)
}
