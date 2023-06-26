package hw03frequencyanalysis

import (
	"regexp"
	"strings"
)

var splitter = regexp.MustCompile("\\s+")

const punctuationMarks = ",.-;!"

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

type Top struct {
	stats []*stat
}

func NewTop(size int) *Top {
	return &Top{
		stats: make([]*stat, size),
	}
}

func (t *Top) Add(word string, count int) {
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

func (t *Top) Result() []string {
	var r []string
	for _, s := range t.stats {
		if s == nil {
			break
		}
		r = append(r, s.word)
	}
	return r
}

func (t *Top) insert(index int, s *stat) {
	if index >= len(t.stats) {
		return
	}
	for i := len(t.stats) - 1; i > index; i-- {
		t.stats[i] = t.stats[i-1]
	}
	t.stats[index] = s
}

func CleanWord(word string) string {
	word = strings.Trim(word, punctuationMarks)
	return strings.ToLower(word)
}

func Top10(input string) []string {
	words := splitter.Split(input, -1)
	wordsFrequency := map[string]int{}
	for _, word := range words {
		word = CleanWord(word)
		if word == "" {
			continue
		}
		wordsFrequency[word]++
	}
	t := NewTop(10)
	for word, count := range wordsFrequency {
		t.Add(word, count)
	}
	return t.Result()
}
