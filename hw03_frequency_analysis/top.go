package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordStats struct {
	word  string
	count int
}

func Top10(text string) []string {
	var lastWordIndexInText int
	counted := make(map[string]int)

	for k, s := range strings.Fields(text) {
		lastWordIndexInText = k
		counted[s]++
	}

	stats := make([]wordStats, 0, lastWordIndexInText)
	top := make([]string, 0, lastWordIndexInText)

	for word, cnt := range counted {
		stats = append(stats, wordStats{word: word, count: cnt})
	}

	sort.Slice(stats, func(i, s int) bool {
		return stats[s].count < stats[i].count || stats[s].count == stats[i].count && strings.Compare(stats[i].word, stats[s].word) == -1
	})

	for _, v := range stats {
		top = append(top, v.word)
	}

	if len(top) > 10 {
		return top[:10]
	}

	return top
}
