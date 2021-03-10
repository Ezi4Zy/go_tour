package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"sun fear", 30}
	z := Person{"sun yu", 4}
	fmt.Println(a, z)
}
