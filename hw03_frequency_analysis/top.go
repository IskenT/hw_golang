package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func Top10(input string) []string {
	splitedWords := strings.Fields(input)
	// Count frequency with dictionary
	wordFreq := make(map[string]int)
	for _, word := range splitedWords {
		if wordFreq[word] != 0 {
			wordFreq[word]++
		} else {
			wordFreq[word] = 1
		}
	}
	wordCounts := make([]wordCount, 0, len(input))
	// Sort by frequency
	for word, count := range wordFreq {
		wordCounts = append(wordCounts, wordCount{word, count})
	}
	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].count == wordCounts[j].count {
			return wordCounts[i].word < wordCounts[j].word
		}
		return wordCounts[i].count > wordCounts[j].count
	})
	// Take top 10
	var topWords []string
	for i := 0; i < len(wordCounts) && i < 10; i++ {
		topWords = append(topWords, wordCounts[i].word)
	}
	return topWords
}
