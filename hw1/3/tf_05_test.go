package main

import (
	"testing"
)

func TestFilterCharsAndNormalize(t *testing.T) {
	data = []rune("Hello, World! 123")
	FilterCharsAndNormalize()
	if string(data) != "hello  world  123" {
		t.Errorf("Expected 'hello  world  123', got '%s'", string(data))
	}
}

func TestScan(t *testing.T) {
	data = []rune("hello world 123")
	scan()
	expected := []string{"hello", "world", "123"}
	for i, word := range words {
		if word != expected[i] {
			t.Errorf("Expected '%s', got '%s'", expected[i], word)
		}
	}
}

func TestFrequencies(t *testing.T) {
	words = []string{"hello", "world", "hello"}
	frequencies()
	expected := map[string]int{"hello": 2, "world": 1}
	for word, count := range expected {
		if freqs[word] != count {
			t.Errorf("Expected count %d for word '%s', got %d", count, word, freqs[word])
		}
	}
}

func TestSortFrequencies(t *testing.T) {
	freqs = map[string]int{"hello": 2, "world": 1}
	sortFrequencies()
	if sortedFreqs[0].Word != "hello" || sortedFreqs[0].Count != 2 {
		t.Errorf("Expected 'hello' with count 2, got '%s' with count %d", sortedFreqs[0].Word, sortedFreqs[0].Count)
	}
}
