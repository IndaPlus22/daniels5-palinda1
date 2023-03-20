package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordC := make(map[string]int)
	splitted := strings.Split(s, " ")
	for _, word := range splitted {
		wordC[word] += 1
	}

	return wordC
}

func main() {
	wc.Test(WordCount)
}
