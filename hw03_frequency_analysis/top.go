package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	frequencyWord := make(map[string]int)
	words := strings.Fields(text)
	if len(words) == 0 {
		return nil
	}
	for _, word := range words {
		word = strings.ToLower(word)
		word = strings.TrimRight(word, "-,.:;")
		if word != "" {
			frequencyWord[word]++
		}
	}
	frequencyMap := make(map[int][]string)
	for word, frequency := range frequencyWord {
		frequencyMap[frequency] = append(frequencyMap[frequency], word)
	}
	frequencySlice := make([]int, 0)
	for frequency, wordSlice := range frequencyMap {
		sort.Slice(wordSlice, func(d, e int) bool {
			return wordSlice[d] < wordSlice[e]
		})
		frequencySlice = append(frequencySlice, frequency)
	}
	sort.Slice(frequencySlice, func(d, e int) bool {
		return frequencySlice[d] > frequencySlice[e]
	})
	var frequencySortSlice []string
	for _, frequency := range frequencySlice {
		frequencySortSlice = append(frequencySortSlice, frequencyMap[frequency]...)
	}
	return frequencySortSlice[:10]
}
