package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	// kevinhwang strings.Fields(s) 相当于 java中的s.split("\\s+")
	wordList := strings.Fields(s)
	for _, v := range wordList {
		ret[v]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
