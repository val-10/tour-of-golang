package main

import (
	"golang.org/x/tour/wc"
	"strings"	
)

func WordCount(s string) map[string]int {
	words:= make(map[string]int)
	N:=strings.Fields(s)
	for _, word:=range N{
		words[word]++}
	return words
}
func main() {
	wc.Test(WordCount)
}