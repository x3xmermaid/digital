package main

import (
	"fmt"
	"strings"
)

// func getIndex()

func RemoveIndex(s []string, index string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == index {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func countAlphabet(word string) {
	wordArray := strings.Split(word, "")
	vowelsArray := []string{"a", "i", "u", "e", "o"}
	totalVowels := 0
	total := 0

	for i := 0; i < len(wordArray); i++ {
		fmt.Println(wordArray[i])
		for j := 0; j < len(vowelsArray); j++ {
			if wordArray[i] == vowelsArray[j] {
				totalVowels++
				if wordArray[i] != wordArray[i+1] {
					wordArray = RemoveIndex(wordArray, wordArray[i])
					i--
				} else {
					wordArray = RemoveIndex(wordArray, wordArray[i])
				}
				break
			} else if j == len(vowelsArray)-1 {
				total++
			}
			fmt.Println(wordArray)
		}
	}

	fmt.Println(totalVowels)
	fmt.Println(total)
}

func main() {
	countAlphabet("uuumaaamama")
}
