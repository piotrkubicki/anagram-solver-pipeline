package wordLengthCombinationFinder

import (
	"log"
	"sort"
)

type wordLengthCombinationFinder struct {
	min           int
	max           int
	maxWords      int
	targetLength  int
	outputChannel chan []int
}

func sum(intSlice []int) int {
	res := 0
	for _, v := range intSlice {
		res += v
	}

	return res
}

func increment(slice []int, max int) []int {
	lastValue := slice[len(slice)-1] + 1
	if lastValue > max {
		slice = increment(slice[:len(slice)-1], max)
	} else {
		slice = slice[:len(slice)-1]
		slice = append(slice, lastValue)
	}

	return slice
}

func generateSeed(min int, max int, targetLength int, maxWords int) []int {
	seed := []int{}
	for {
		seedTotal := sum(seed)
		if seedTotal == targetLength {
			break
		} else if len(seed) < maxWords {
			seed = append(seed, min)
		} else if seedTotal < targetLength {
			seed = increment(seed, max)
		}
	}

	return seed
}

func FindCombination(wordLengths []int, targetLength int, maxWords int, output_channel chan []int) bool {
	sort.Ints(wordLengths)
	min := wordLengths[0]
	max := wordLengths[len(wordLengths)-1]
	seed := generateSeed(min, max, targetLength, maxWords)
	log.Println(seed)

	return true
}
