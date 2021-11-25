package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestCleanWord(t *testing.T) {
	cases := []struct{ word, expected string }{
		{"a's", "a"},
		{"abacus", "abacus"},
		{"abacus's", "abacus"},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Test: %v", tc.word), func(t *testing.T) {
			actual := cleanWord(tc.word)
			if actual != tc.expected {
				t.Fatalf("Fail: %v not equal %v!", actual, tc.expected)
			}
		})
	}
}

func TestMapWords(t *testing.T) {
	data := "hello\nhey\ntest\nwest\nchest\nbest\nfest\nchest's"
	reader := bufio.NewReader(strings.NewReader(data))
	actual, err := mapWords(reader)

	if err != nil {
		t.Fatal("mapWords return error")
	}
	expected := map[int][]string{
		3: []string{"hey"},
		4: []string{"test", "west", "best", "fest"},
		5: []string{"hello", "chest"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("mapWords: %v different than %v!", actual, expected)
	}
}

func TestFindMinMaxWordLength(t *testing.T) {
	cases := []struct {
		dictionary  map[int][]string
		expectedMin int
		expectedMax int
	}{
		{
			map[int][]string{
				3: []string{"one", "two"},
				4: []string{"four", "five", "zero"},
				5: []string{"eight"},
			},
			3,
			5,
		},
		{
			map[int][]string{
				1: []string{"a"},
				3: []string{"one", "two"},
				5: []string{"eight"},
			},
			1,
			5,
		},
		{
			map[int][]string{
				3: []string{"one", "two"},
			},
			3,
			3,
		},
	}

	for _, tc := range cases {
		actualMin, actualMax := findMinMaxWordLength(tc.dictionary)

		if actualMin != tc.expectedMin && actualMax != tc.expectedMax {
			t.Fatalf("Fail: actualMin: %v expectedMin: %v actualMax: %v expectedMax: %v", actualMin, tc.expectedMin, actualMax, tc.expectedMax)
		}
	}
}
