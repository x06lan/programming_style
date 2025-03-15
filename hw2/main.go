package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type DataStorageManager struct {
	data string
}

func NewDataStorageManager(filePath string) *DataStorageManager {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	reg := regexp.MustCompile(`[\W_]+`)
	data := strings.ToLower(reg.ReplaceAllString(string(file), " "))
	return &DataStorageManager{data: data}
}

func (d *DataStorageManager) Words() []string {
	// split by space
	return strings.Fields(d.data)
}

type StopWordManager struct {
	stopWords map[string]bool
}

func NewStopWordManager(filePath string) *StopWordManager {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(file), ",")
	stopWords := make(map[string]bool)
	for _, word := range words {
		stopWords[word] = true
	}
	for ascii_lowercase := 'a'; ascii_lowercase <= 'z'; ascii_lowercase++ {
		stopWords[string(ascii_lowercase)] = true
	}
	return &StopWordManager{stopWords: stopWords}
}

func (s *StopWordManager) IsStopWord(word string) bool {
	return s.stopWords[word]
}

type WordFrequencyManager struct {
	wordFreqs map[string]int
}

func NewWordFrequencyManager() *WordFrequencyManager {
	return &WordFrequencyManager{wordFreqs: make(map[string]int)}
}

func (w *WordFrequencyManager) IncrementCount(word string) {
	w.wordFreqs[word]++
}

func (w *WordFrequencyManager) Sorted() []struct {
	word  string
	count int
} {
	wordCounts := make([]struct {
		word  string
		count int
	}, 0, len(w.wordFreqs))

	for k, v := range w.wordFreqs {
		wordCounts = append(wordCounts, struct {
			word  string
			count int
		}{k, v})
	}
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})
	return wordCounts
}

type WordFrequencyController struct {
	storageManager  *DataStorageManager
	stopWordManager *StopWordManager
	wordFreqManager *WordFrequencyManager
}

func NewWordFrequencyController(filePath string, stopWordFile string) *WordFrequencyController {
	return &WordFrequencyController{
		storageManager:  NewDataStorageManager(filePath),
		stopWordManager: NewStopWordManager(stopWordFile),
		wordFreqManager: NewWordFrequencyManager(),
	}
}

func (c *WordFrequencyController) Run() {
	for _, w := range c.storageManager.Words() {
		if !c.stopWordManager.IsStopWord(w) {
			c.wordFreqManager.IncrementCount(w)
		}
	}

	wordFreqs := c.wordFreqManager.Sorted()
	for i, wc := range wordFreqs {
		// todo:change back to 25
		if i >= 25 {
			break
		}
		fmt.Printf("%s - %d\n", wc.word, wc.count)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <file_path> ")
		return
	}
	controller := NewWordFrequencyController(os.Args[1], os.Args[2])
	controller.Run()
}
