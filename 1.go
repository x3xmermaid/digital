package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortAlphabet(word string) {
	wordArray := strings.Split(word, "")
	vowelsArray := []string{"a", "i", "u", "e", "o"}
	vowels := []string{}
	consonants := []string{}

	for i := 0; i < len(wordArray); i++ {
		for j := 0; j < len(vowelsArray); j++ {
			if wordArray[i] == vowelsArray[j] {
				vowels = append(vowels, wordArray[i])
				break
			} else if j == len(vowelsArray)-1 {
				consonants = append(consonants, wordArray[i])
			}
		}
	}

	sort.Strings(vowels)
	sort.Strings(consonants)
	vowels = append(vowels, consonants...)
	stringVowels := strings.Join(vowels, "")
	fmt.Printf("Hasil Sorting : %s", stringVowels)
}

func main() {
	var kata string
	fmt.Printf("Masukkan Kata : ")
	fmt.Scanf("%s", &kata)
	sortAlphabet(kata)
}
