package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	/*// kevinhwang strings.Fields(s) 相当于 java中的s.split("\\s+")
	wordList := strings.Fields(s)
	for _, v := range wordList {
		ret[v]++
	}*/

	var tmp []byte
	for _, v := range s {
		if v == ' ' {
			ret[string(tmp)]++
			tmp = make([]byte, 0)
		} else {
			tmp = append(tmp, byte(v))
		}
	}
	if len(tmp) > 0 {
		ret[string(tmp)]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
