package main

import (
	"reflect"
	"testing"
)

func TestFilterCharsAndNormalize(t *testing.T) {
	input := "Hello, World! 123."
	expected := "hello  world  123 "
	output := filterCharsAndNormalize(input)
	if output != expected {
		t.Errorf("Expected %q but got %q", expected, output)
	}
}

func TestTokenize(t *testing.T) {
	input := "hello world 123"
	expected := []string{"hello", "world", "123"}
	output := tokenize(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v but got %v", expected, output)
	}
}

func TestRemoveStopWords(t *testing.T) {
	words := []string{"hello", "world", "this", "is", "a", "test"}
	stopWordsFile := "../../example/stop_words.txt"
	stopWords := "this,is,a"

	// Mock reading stop words from file
	mockReadFile := func(string) string { return stopWords }
	stopWordsContent := mockReadFile(stopWordsFile)

	stopWordsList := tokenize(stopWordsContent)
	stopWordsMap := make(map[string]struct{})
	for _, word := range stopWordsList {
		stopWordsMap[word] = struct{}{}
	}

	expected := []string{"hello", "world", "test"}
	output := removeStopWords(words, stopWordsFile)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v but got %v", expected, output)
	}
}

func TestComputeFrequencies(t *testing.T) {
	words := []string{"hello", "world", "hello", "test"}
	expected := map[string]int{"hello": 2, "world": 1, "test": 1}
	output := computeFrequencies(words)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v but got %v", expected, output)
	}
}

func TestSortFrequencies(t *testing.T) {
	freqs := map[string]int{"hello": 2, "world": 1, "test": 3}
	expected := []WordFrequency{{"test", 3}, {"hello", 2}, {"world", 1}}
	output := sortFrequencies(freqs)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v but got %v", expected, output)
	}
}
