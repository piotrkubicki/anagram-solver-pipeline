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

func mapWord(word string) map[rune]int {
	chars := make(map[rune]int)

	for _, char := range word {
		if char == ' ' {
			continue
		}
		if _, ok := chars[char]; ok {
			chars[char]++
		} else {
			chars[char] = 1
		}
	}

	return chars
}

func cleanWord(word string) string {
	word = strings.TrimSuffix(word, "\n")
	word = strings.Split(word, "'")[0]

	return word
}

func isValid(word string, allowedChars map[rune]int, minLength int, maxLength int) bool {
	wordLength := len(word)
	if wordLength < minLength || wordLength > maxLength {
		return false
	}

	mappedWord := mapWord(word)

	for key, val := range mappedWord {
		if allowedCharsCount, ok := allowedChars[key]; ok {
			if val > allowedCharsCount {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func mapWords(reader *bufio.Reader, allowedCharacters map[rune]int, minLength int, maxLength int) (map[int][]string, error) {
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
				if isValid(cleanLine, allowedCharacters, minLength, maxLength) == true {
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
	}

	return wordsByLength, nil
}

func run(fileName string, phrase string, minChars int, maxChars int, maxWords int, targetLength int, hashes []string) {
	passwords := []string{}
	allowedChars := mapWord(phrase)
	wordLengthPermutationChannel := make(chan []int, 35)
	permutationsChannel := make(chan string, 1000)

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	dictionary, err := mapWords(reader, allowedChars, minChars, maxChars)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		word_length_permutation_generator.Generate(minChars, maxChars, targetLength, maxWords, wordLengthPermutationChannel)
	}()

	for i := 0; i < 10; i++ {
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

func main() {
	originalAnagram := "poultry outwits ants"
	minWordLength := 2
	maxWordLength := 15
	targetLength := 18
	maxWords := 4
	hashes := []string{
		"e4820b45d2277f3844eac66c903e84be",
		"23170acc097c24edb98fc5488ab033fe",
		"665e5bcb0c20062fe8abaaf4628bb154",
	}

	log.SetPrefix("anagram_solver: ")
	log.SetFlags(0)

	run("wordlist", originalAnagram, minWordLength, maxWordLength, maxWords, targetLength, hashes)
}
