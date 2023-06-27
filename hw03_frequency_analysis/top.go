package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var splitter = regexp.MustCompile(`\s+`)

const punctuationMarks = ",.-;!"

func Top10(input string) []string {
	words := splitter.Split(input, -1)
	wordsFrequency := map[string]int{}
	for _, word := range words {
		word = cleanWord(word)
		if word == "" {
			continue
		}
		wordsFrequency[word]++
	}
	cleanedWords := make([]string, 0, len(wordsFrequency))
	for word := range wordsFrequency {
		cleanedWords = append(cleanedWords, word)
	}
	sort.Slice(cleanedWords, func(i, j int) bool {
		w1, w2 := cleanedWords[i], cleanedWords[j]
		if wordsFrequency[w1] == wordsFrequency[w2] {
			return strings.Compare(w1, w2) < 1
		}
		return wordsFrequency[w1] > wordsFrequency[w2]
	})
	if len(cleanedWords) > 10 {
		cleanedWords = cleanedWords[:10]
	}
	return cleanedWords
}

func Top10V2(input string) []string {
	words := splitter.Split(input, -1)
	wordsFrequency := map[string]int{}
	for _, word := range words {
		word = cleanWord(word)
		if word == "" {
			continue
		}
		wordsFrequency[word]++
	}
	t := newTop(10)
	for word, count := range wordsFrequency {
		t.add(word, count)
	}
	return t.result()
}

type stat struct {
	word  string
	count int
}

func (s *stat) less(r *stat) bool {
	if s == nil || s.count < r.count {
		return true
	}
	if s.count == r.count {
		return strings.Compare(s.word, r.word) > -1
	}
	return false
}

type top struct {
	stats []*stat
}

func newTop(size int) *top {
	return &top{
		stats: make([]*stat, size),
	}
}

func (t *top) add(word string, count int) {
	n := &stat{
		word:  word,
		count: count,
	}
	for index, s := range t.stats {
		if s.less(n) {
			t.insert(index, n)
			break
		}
	}
}

func (t *top) result() []string {
	r := make([]string, 0, len(t.stats))
	for _, s := range t.stats {
		if s == nil {
			break
		}
		r = append(r, s.word)
	}
	return r
}

func (t *top) insert(index int, s *stat) {
	if index >= len(t.stats) {
		return
	}
	for i := len(t.stats) - 1; i > index; i-- {
		t.stats[i] = t.stats[i-1]
	}
	t.stats[index] = s
}

func cleanWord(word string) string {
	word = strings.Trim(word, punctuationMarks)
	return strings.ToLower(word)
}
