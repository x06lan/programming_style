package main

import (
	"os"
	"testing"
)

// Unit Tests
func TestWords(t *testing.T) {
	dsm := &DataStorageManager{data: "hello world hello"}
	words := dsm.Words()
	if len(words) != 3 || words[0] != "hello" || words[1] != "world" || words[2] != "hello" {
		t.Errorf("Expected ['hello', 'world', 'hello'], got %v", words)
	}
}
func TestIsStopWord(t *testing.T) {
	swm := &StopWordManager{stopWords: map[string]bool{
		"hello": true,
		"world": true,
	}}
	if !swm.IsStopWord("hello") || !swm.IsStopWord("world") || swm.IsStopWord("hello world") {
		t.Errorf("Expected true for 'hello' and 'world'")
	}
}

func TestIncrementCount(t *testing.T) {
	wfm := &WordFrequencyManager{wordFreqs: map[string]int{}}
	wfm.IncrementCount("hello")
	if wfm.wordFreqs["hello"] != 1 || wfm.wordFreqs["world"] != 0 {
		t.Errorf("Expected 'hello' count to be 2 and 'world' count to be 0")
	}
	wfm.IncrementCount("world")
	wfm.IncrementCount("hello")
	if wfm.wordFreqs["hello"] != 2 || wfm.wordFreqs["world"] != 1 {
		t.Errorf("Expected 'hello' count to be 2 and 'world' count to be 1")
	}
}

func TestSorted(t *testing.T) {
	wfm := &WordFrequencyManager{wordFreqs: map[string]int{
		"hello": 2,
		"world": 1,
	}}

	sorted := wfm.Sorted()

	if len(sorted) != 2 {
		t.Errorf("sorted count to be 2")
	}
	if sorted[0].word != "hello" || sorted[0].count != 2 {
		t.Errorf("Expected 'hello' count to be 2")
	}
	if sorted[1].word != "world" || sorted[1].count != 1 {
		t.Errorf("expected 'world' count to be 1")
	}
}

func TestNewDataStorageManager(t *testing.T) {
	content := "Hello, World! Hello."
	filePath := "test.txt"
	os.WriteFile(filePath, []byte(content), 0644)
	defer os.Remove(filePath)

	dsm := NewDataStorageManager(filePath)
	words := dsm.Words()

	expected := []string{"hello", "world", "hello"}
	if len(words) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, words)
	}
}

func TestNewStopWordManager(t *testing.T) {
	content := "a,an,the,and"
	filePath := "stopwords.txt"
	os.WriteFile(filePath, []byte(content), 0644)
	defer os.Remove(filePath)

	swm := NewStopWordManager(filePath)

	if !swm.IsStopWord("a") || !swm.IsStopWord("the") || swm.IsStopWord("banana") {
		t.Errorf("StopWordManager failed to load stop words correctly")
	}
}

func TestWordFrequencyController_Run(t *testing.T) {
	// Prepare test files
	textContent := "Hello world! Hello again. Goodbye world."
	stopWordsContent := "a,an,the,and,goodbye"
	textFile := "test_text.txt"
	stopFile := "test_stopwords.txt"

	os.WriteFile(textFile, []byte(textContent), 0644)
	os.WriteFile(stopFile, []byte(stopWordsContent), 0644)
	defer os.Remove(textFile)
	defer os.Remove(stopFile)

	// Run controller
	controller := NewWordFrequencyController(textFile, stopFile)
	controller.Run()

	// Validate output (mocking fmt.Printf would be better in real-world tests)
	expected := []struct {
		word  string
		count int
	}{
		{"hello", 2},
		{"world", 2},
		{"again", 1},
	}

	result := controller.wordFreqManager.Sorted()
	// fmt.Println(result)
	for i, exp := range expected {
		if result[i].word != exp.word || result[i].count != exp.count {
			t.Errorf("Expected %s - %d, got %s - %d", exp.word, exp.count, result[i].word, result[i].count)
		}
	}
}
