package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	s_fields = s.Fields()
	for _, word := range s_fields {
		counter[word] += 1
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
