package word_length_permutation_generator

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		slice    []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{2, 5, 1, 7}, 15},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Test: %v", tc.slice), func(t *testing.T) {
			actual := sum(tc.slice)
			if actual != tc.expected {
				t.Fatalf("Fail: %v not equal %v!", actual, tc.expected)
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	cases := []struct {
		input    []int
		max      int
		expected []int
	}{
		{[]int{2, 4, 4}, 4, []int{3}},
		{[]int{2, 4, 3}, 4, []int{2, 4, 4}},
		{[]int{2, 4, 5}, 5, []int{2, 5}},
		{[]int{1, 2, 7}, 7, []int{1, 3}},
		{[]int{5, 2, 5}, 6, []int{5, 2, 6}},
	}

	for _, tc := range cases {
		actual := increment(tc.input, tc.max)

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("Fail: %v not equal %v!", actual, tc.expected)
		}
	}
}

func TestGenerate(t *testing.T) {
	cases := []struct {
		min           int
		max           int
		maxWords      int
		targetLength  int
		outputChannel chan []int
		expected      [][]int
	}{
		{2, 4, 3, 10, make(chan []int),
			[][]int{
				[]int{2, 4, 4},
				[]int{3, 3, 4},
				[]int{3, 4, 3},
				[]int{4, 2, 4},
				[]int{4, 3, 3},
				[]int{4, 4, 2},
			},
		},
		{2, 5, 4, 15, make(chan []int),
			[][]int{
				[]int{2, 3, 5, 5},
				[]int{2, 4, 4, 5},
				[]int{2, 4, 5, 4},
				[]int{2, 5, 3, 5},
				[]int{2, 5, 4, 4},
				[]int{2, 5, 5, 3},
				[]int{3, 2, 5, 5},
				[]int{3, 3, 4, 5},
				[]int{3, 3, 5, 4},
				[]int{3, 4, 3, 5},
				[]int{3, 4, 4, 4},
				[]int{3, 4, 5, 3},
				[]int{3, 5, 2, 5},
				[]int{3, 5, 3, 4},
				[]int{3, 5, 4, 3},
				[]int{3, 5, 5, 2},
				[]int{4, 2, 4, 5},
				[]int{4, 2, 5, 4},
				[]int{4, 3, 3, 5},
				[]int{4, 3, 4, 4},
				[]int{4, 3, 5, 3},
				[]int{4, 4, 2, 5},
				[]int{4, 4, 3, 4},
				[]int{4, 4, 4, 3},
				[]int{4, 4, 5, 2},
				[]int{4, 5, 2, 4},
				[]int{4, 5, 3, 3},
				[]int{4, 5, 4, 2},
				[]int{5, 2, 3, 5},
				[]int{5, 2, 4, 4},
				[]int{5, 2, 5, 3},
				[]int{5, 3, 2, 5},
				[]int{5, 3, 3, 4},
				[]int{5, 3, 4, 3},
				[]int{5, 3, 5, 2},
				[]int{5, 4, 2, 4},
				[]int{5, 4, 3, 3},
				[]int{5, 4, 4, 2},
				[]int{5, 5, 2, 3},
				[]int{5, 5, 3, 2},
				[]int{5, 5, 5},
			},
		},
	}

	for _, tc := range cases {
		actual := [][]int{}
		go Generate(tc.min, tc.max, tc.targetLength, tc.maxWords, tc.outputChannel)
		for {
			output := <-tc.outputChannel
			if len(output) == 0 {
				break
			}
			actual = append(actual, output)
		}
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("Fail: %v not equal %v", actual, tc.expected)
		}
	}
}
