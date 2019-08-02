package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
)

const (
	ListeningPort = ":8080"
)

type Data struct {
	Text string `json:"Text"`
}

func sortAlphabet2(word string) string {
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
	return stringVowels
}

func postService(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t Data
	err := decoder.Decode(&t)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("2"))
	}

	word := t.Text
	sortLetters := sortAlphabet2(word)

	var hasil bytes.Buffer
	hasil.WriteString("Hasil : ")
	hasil.WriteString(sortLetters)

	w.Write([]byte(hasil.String()))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/countVowels", postService).Methods(http.MethodPost)

	log.Printf("Starting http server at %v", ListeningPort)
	err := http.ListenAndServe(ListeningPort, r)
	if err != nil {
		log.Fatalf("Unable to run http server: %v", err)
	}
	log.Println("Stopping API Service...")
}
