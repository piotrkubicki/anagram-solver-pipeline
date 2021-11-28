package permutation_generator

import (
	"reflect"
	"sync"
	"testing"
)

func TestGenerate(t *testing.T) {
	wordLists := [][]string{
		[]string{"uno", "due", "one"},
		[]string{"wake", "sure"},
	}
	outputChannel := make(chan string, 50)
	expected := []string{
		"uno wake",
		"uno sure",
		"due wake",
		"due sure",
		"one wake",
		"one sure",
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		generate(wordLists, []string{}, outputChannel)
	}()
	wg.Wait()
	close(outputChannel)
	actual := []string{}
	for output := range outputChannel {
		actual = append(actual, output)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Fail! %v not equal %v", actual, expected)
	}
}

func TestRun(t *testing.T) {
	dictionary := map[int][]string{
		2: []string{"on", "at"},
		3: []string{"one", "two", "six"},
		4: []string{"four", "five"},
		5: []string{"seven", "eight"},
	}
	combinations := [][]int{
		[]int{2, 3, 4},
		[]int{3, 5},
		[]int{},
	}
	inputChannel := make(chan []int, 10)
	outputChannel := make(chan string, 10)
	actual := []string{}
	expected := []string{
		"on one four",
		"on one five",
		"on two four",
		"on two five",
		"on six four",
		"on six five",
		"at one four",
		"at one five",
		"at two four",
		"at two five",
		"at six four",
		"at six five",
		"one seven",
		"one eight",
		"two seven",
		"two eight",
		"six seven",
		"six eight",
	}
	go Run(dictionary, inputChannel, outputChannel)

	for _, combination := range combinations {
		inputChannel <- combination
	}

	for {
		permutation := <-outputChannel
		if permutation == "" {
			break
		}
		actual = append(actual, permutation)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Fail! %v not equal %v", actual, expected)
	}
}
