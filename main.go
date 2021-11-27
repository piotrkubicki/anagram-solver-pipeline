package main

import (
	"bufio"
	"github.com/piotrkubicki/anagram-solver-pipeline/permutation_generator"
	"github.com/piotrkubicki/anagram-solver-pipeline/word_length_permutation_generator"
	"log"
	"os"
	"strings"
)

func cleanWord(word string) string {
	word = strings.TrimSuffix(word, "\n")
	word = strings.Split(word, "'")[0]

	return word
}

func mapWords(reader *bufio.Reader) (map[int][]string, error) {
	wordsByLength := make(map[int][]string)
	seenWords := make(map[string]bool)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		cleanLine := cleanWord(line)
		lineLen := len(cleanLine)

		if lineLen > 0 {
			if _, value := seenWords[cleanLine]; !value {
				if val, ok := wordsByLength[lineLen]; ok {
					val = append(val, cleanLine)
					wordsByLength[lineLen] = val
				} else {
					wordsByLength[lineLen] = []string{cleanLine}
				}
				seenWords[cleanLine] = true
			}
		}
	}

	return wordsByLength, nil
}

func findMinMaxWordLength(dictionary map[int][]string) (int, int) {
	minWordLength := 1000
	maxWordLength := 0

	for key := range dictionary {
		if key < minWordLength {
			minWordLength = key
		}
		if key > maxWordLength {
			maxWordLength = key
		}
	}

	return minWordLength, maxWordLength
}

func main() {
	log.SetPrefix("anagram_solver: ")
	log.SetFlags(0)

	file, err := os.Open("wordlist")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	dictionary, err := mapWords(reader)
	if err != nil {
		log.Fatal(err)
	}

	// originalAnagram := "poultry outwits ants"
	minWordLength, maxWordLength := findMinMaxWordLength(dictionary)
	targetLength := 18
	maxWords := 3
	wordLengthPermutationChannel := make(chan []int, 5)
	permutationsChannel := make(chan string, 10)

	go word_length_permutation_generator.Generate(minWordLength, maxWordLength, targetLength, maxWords, wordLengthPermutationChannel)
	for i := 0; i < 5; i++ {
		go permutation_generator.Run(dictionary, wordLengthPermutationChannel, permutationsChannel)
	}

	for {
		phrase := <-permutationsChannel
		log.Printf("Phrase: %v", phrase)
	}
}
