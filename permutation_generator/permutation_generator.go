package permutation_generator

import (
	"log"
	"strings"
)

func generate(wordLists [][]string, words []string, outputChannel chan string) {
	if len(wordLists) == 0 {
		phrase := strings.Join(words, " ")
		outputChannel <- phrase
	} else {
		for _, word := range wordLists[0] {
			wordsCpy := make([]string, len(words))
			copy(wordsCpy, words)
			wordsCpy = append(wordsCpy, word)
			generate(wordLists[1:], wordsCpy, outputChannel)
		}
	}
}

func Run(dictionary map[int][]string, inputChannel chan []int, outputChannel chan string) {
	for {
		wordLengths := <-inputChannel

		if len(wordLengths) == 0 {
			break
		}
		log.Printf("Checking permutation: %v", wordLengths)
		wordLists := [][]string{}
		for _, length := range wordLengths {
			words := make([]string, len(dictionary[length]))
			copy(words, dictionary[length])
			wordLists = append(wordLists, words)
		}
		generate(wordLists, []string{}, outputChannel)
		log.Printf("Checking permutation %v done", wordLengths)
	}
	outputChannel <- ""
}
