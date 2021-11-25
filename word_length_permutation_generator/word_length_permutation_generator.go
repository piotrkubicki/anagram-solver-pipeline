package word_length_permutation_generator

import (
	"log"
)

func sum(intSlice []int) int {
	res := 0
	for _, v := range intSlice {
		res += v
	}

	return res
}

func increment(slice []int, max int) []int {
	if len(slice) == 0 {
		return slice
	}
	lastValue := slice[len(slice)-1] + 1
	if lastValue > max {
		slice = increment(slice[:len(slice)-1], max)
	} else {
		slice = slice[:len(slice)-1]
		slice = append(slice, lastValue)
	}
	return slice
}

func Generate(min int, max int, targetLength int, maxWords int, outputChannel chan []int) {
	seed := []int{}
	for {
		seedTotal := sum(seed)
		if seedTotal == targetLength {
			comb := make([]int, len(seed))
			copy(comb, seed)
			outputChannel <- comb
			seed = increment(seed, max)
		} else if len(seed) < maxWords {
			seed = append(seed, min)
		} else {
			seed = increment(seed, max)
		}
		if len(seed) == 0 {
			outputChannel <- seed
			break
		}
	}
	log.Println("Generator finished!")
}
