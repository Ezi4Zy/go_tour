// Package main provides ...
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("Mac OS!")
	case "linux":
		fmt.Println("Linux!")
	default:
		fmt.Println(os + "!")
	}
}
