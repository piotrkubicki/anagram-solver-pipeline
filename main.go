package main

import (
	"bufio"
	"github.com/piotrkubicki/anagram-solver-pipeline/checker"
	"github.com/piotrkubicki/anagram-solver-pipeline/permutation_generator"
	"github.com/piotrkubicki/anagram-solver-pipeline/word_length_permutation_generator"
	"log"
	"os"
	"strings"
	"sync"
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
	minWordLength = 3
	maxWordLength = 10
	targetLength := 18
	maxWords := 3
	hashes := []string{
		"e4820b45d2277f3844eac66c903e84be",
		"23170acc097c24edb98fc5488ab033fe",
		"665e5bcb0c20062fe8abaaf4628bb154",
	}
	passwords := []string{}
	wordLengthPermutationChannel := make(chan []int, 5)
	permutationsChannel := make(chan string, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		word_length_permutation_generator.Generate(minWordLength, maxWordLength, targetLength, maxWords, wordLengthPermutationChannel)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			permutation_generator.Run(dictionary, wordLengthPermutationChannel, permutationsChannel)
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			checker.Run(hashes, &passwords, permutationsChannel)
		}()
	}

	wg.Wait()
	log.Printf("Found passwords: %v", passwords)
}
