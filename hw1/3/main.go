package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type WordFrequency struct {
	Word  string
	Count int
}

func readFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + " ")
	}
	return content.String()
}

func filterCharsAndNormalize(text string) string {
	var filtered strings.Builder
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			filtered.WriteRune(unicode.ToLower(r))
		} else {
			filtered.WriteRune(' ')
		}
	}
	return filtered.String()
}

func tokenize(text string) []string {
	return strings.Fields(text)
}

func removeStopWords(words []string, stopWordsFile string) []string {
	file, err := os.Open(stopWordsFile)
	if err != nil {
		fmt.Println("Error opening stop words file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stopWordsList := strings.Split(scanner.Text(), ",")
	stopWords := make(map[string]struct{})
	for _, word := range stopWordsList {
		stopWords[word] = struct{}{}
	}

	var filtered []string
	for _, word := range words {
		if _, found := stopWords[word]; !found {
			filtered = append(filtered, word)
		}
	}
	return filtered
}

func computeFrequencies(words []string) map[string]int {
	freqs := make(map[string]int)
	for _, word := range words {
		freqs[word]++
	}
	return freqs
}

func sortFrequencies(freqs map[string]int) []WordFrequency {
	sortedFreqs := make([]WordFrequency, 0, len(freqs))
	for word, count := range freqs {
		sortedFreqs = append(sortedFreqs, WordFrequency{word, count})
	}

	sort.Slice(sortedFreqs, func(i, j int) bool {
		return sortedFreqs[i].Count > sortedFreqs[j].Count
	})

	return sortedFreqs
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <text_file> <stop_words_file>")
		return
	}

	textFile := os.Args[1]
	stopWordsFile := os.Args[2]

	sortedFrequencies := sortFrequencies(computeFrequencies(removeStopWords(tokenize(filterCharsAndNormalize(readFile(textFile))), stopWordsFile)))

	for _, item := range sortedFrequencies[:25] {
		fmt.Println(item.Word, "-", item.Count)
	}
}
