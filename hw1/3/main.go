package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var data []rune
var words []string
var freqs map[string]int
var sortedFreqs []struct {
	Word  string
	Count int
}

func readFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + " ")
	}
	data = []rune(content.String())
}

func FilterCharsAndNormalize() {
	var filtered []rune
	for _, r := range data {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			filtered = append(filtered, unicode.ToLower(r))
		} else {
			filtered = append(filtered, ' ')
		}
	}
	data = filtered
}
func scan() {
	words = strings.Fields(string(data))
}

func removeStopWords(stopWordsFile string) {
	file, err := os.Open(stopWordsFile)
	if err != nil {
		fmt.Println("Error opening stop words file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stopWordsList := strings.Split(scanner.Text(), ",")
	stopWords := make(map[string]bool)
	for _, word := range stopWordsList {
		stopWords[word] = true
	}

	filtered := []string{}
	for _, word := range words {
		if _, found := stopWords[word]; !found {
			filtered = append(filtered, word)
		}
	}
	words = filtered
}

func frequencies() {
	freqs_buff := make(map[string]int)
	for _, word := range words {
		freqs_buff[word]++
	}
	// save freqs to global variable
	freqs = freqs_buff

}

func sortFrequencies() {
	sortedFreqs_buff := make([]struct {
		Word  string
		Count int
	}, 0)
	// iter freq hash map
	for word, count := range freqs {
		sortedFreqs_buff = append(sortedFreqs_buff, struct {
			Word  string
			Count int
		}{word, count})
	}
	// sort
	for i := 0; i < len(sortedFreqs_buff); i++ {
		for j := i + 1; j < len(sortedFreqs_buff); j++ {
			if sortedFreqs_buff[i].Count < sortedFreqs_buff[j].Count {
				sortedFreqs_buff[i], sortedFreqs_buff[j] = sortedFreqs_buff[j], sortedFreqs_buff[i]
			}
		}
	}
	sortedFreqs = sortedFreqs_buff

}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <text_file> <stop_words_file>")
		return
	}

	textFile := os.Args[1]
	stopWordsFile := os.Args[2]

	readFile(textFile)
	FilterCharsAndNormalize()
	scan()
	removeStopWords(stopWordsFile)
	frequencies()
	sortFrequencies()

	for _, item := range sortedFreqs[:25] {
		fmt.Println(item.Word, "-", item.Count)
	}
}
