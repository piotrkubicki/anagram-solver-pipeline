package checker

import (
	"reflect"
	"sync"
	"testing"
)

func TestCheck(t *testing.T) {
	hashes := []string{
		"8c6d115258631625b625486f81b09532",
	}
	cases := []struct {
		phrase   string
		expected bool
	}{
		{"this is test", true},
		{"this one fail", false},
	}
	for _, tc := range cases {
		if check(tc.phrase, hashes) != tc.expected {
			t.Fatal("Fail! Password check failed")
		}
	}
}

func TestRun(t *testing.T) {
	hashes := []string{
		"5e4fe0155703dde467f3ab234e6f966f",
		"75429ca6672f64684f72ab35558ce4d5",
	}
	actual := []string{}
	phrases := []string{
		"one two three", "this this one", "this one is wrong", "this is it", "this is test", "this one not",
	}
	expected := []string{"one two three", "this is it"}

	phrasesChannel := make(chan string, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Run(hashes, &actual, phrasesChannel)
	}()
	for _, phrase := range phrases {
		phrasesChannel <- phrase
	}
	wg.Wait()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Fail! %v not equal %v", actual, expected)
	}
}
