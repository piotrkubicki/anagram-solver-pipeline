package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestMapWord(t *testing.T) {
	word := "this is test"
	expected := map[rune]int{
		't': 3,
		'h': 1,
		'i': 2,
		's': 3,
		'e': 1,
	}

	actual := mapWord(word)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Fail! %v not equal %v", actual, expected)
	}
}

func TestIsValid(t *testing.T) {
	allowedChars := map[rune]int{
		't': 2,
		'e': 1,
		's': 1,
		'a': 3,
		'd': 2,
		'l': 2,
		'o': 1,
		'w': 1,
		'g': 5,
	}
	cases := []struct {
		word     string
		expected bool
	}{
		{"test", true},
		{"tests", false},
		{"aloe", true},
		{"vault", false},
		{"go", false},
		{"allowed", false},
	}

	for _, tc := range cases {
		if isValid(tc.word, allowedChars, 3, 6) != tc.expected {
			t.Fatalf("Fail! isValid not returned %v for %v", tc.expected, tc.word)
		}
	}
}

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
	allowedChars := map[rune]int{
		'h': 1,
		'e': 1,
		'y': 1,
		't': 2,
		's': 1,
		'w': 1,
		'b': 1,
		'f': 1,
		'l': 2,
		'o': 1,
	}
	actual, err := mapWords(reader, allowedChars, 3, 7)

	if err != nil {
		t.Fatal("mapWords return error")
	}
	expected := map[int][]string{
		3: []string{"hey"},
		4: []string{"test", "west", "best", "fest"},
		5: []string{"hello"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("mapWords: %v different than %v!", actual, expected)
	}
}
