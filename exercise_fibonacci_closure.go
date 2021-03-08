// Package main provides ...
package main

import (
	"fmt"
)

func fibonacci() func() int {
	pre, last := 0, 1
	return func() int {
		ret := pre
		pre, last = last, pre+last
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
