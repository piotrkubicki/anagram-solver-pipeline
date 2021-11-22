package main

import (
	"bufio"
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

func main() {
	log.SetPrefix("anagram_solver: ")
	log.SetFlags(0)

	file, err := os.Open("wordlist")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	_, err = mapWords(reader)
	if err != nil {
		log.Fatal(err)
	}
}
