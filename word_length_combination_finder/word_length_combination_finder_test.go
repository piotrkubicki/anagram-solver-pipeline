package wordLengthCombinationFinder

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

func TestGenerateSeed(t *testing.T) {
	cases := []struct {
		min          int
		max          int
		targetLength int
		maxWords     int
		expected     []int
	}{
		{2, 7, 11, 3, []int{2, 2, 7}},
		{3, 4, 12, 4, []int{3, 3, 3, 3}},
		{1, 11, 34, 4, []int{1, 11, 11, 11}},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Test: %v", i), func(t *testing.T) {
			actual := generateSeed(tc.min, tc.max, tc.targetLength, tc.maxWords)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("Fail: %v not equal %v!", actual, tc.expected)
			}
		})
	}
}
