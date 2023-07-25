package wc

import (
	"regexp"
	"sort"
	"strings"
)

type Counter map[string]int

var splitter *regexp.Regexp = regexp.MustCompile(" ")

func WordCountRegexp(s string) Counter {
	counter := make(Counter)
	for _, word := range splitter.Split(s, -1) {
		word := strings.ToLower(word)
		counter[word]++
	}
	return counter
}

func WordCountFields(s string) Counter {
	counter := make(Counter)
	for _, word := range strings.Fields(s) {
		word = strings.ToLower(word)
		counter[word]++
	}
	return counter
}

func WordCountSplit(s string) Counter {
	counter := make(Counter)
	for _, word := range strings.Split(s, " ") {
		word := strings.ToLower(word)
		counter[word]++
	}
	return counter
}

func WordCountPhrase(s string) Counter {
	phrase := strings.ToLower(s)
	counter := make(Counter)
	for _, word := range strings.Split(phrase, " ") {
		counter[word]++
	}
	return counter
}

func WordCountAllocate(s string) Counter {
	words := strings.Split(s, " ")
	size := len(words)
	if size > 10000 {
		size = 10000
	}

	counter := make(Counter, size)
	for _, word := range words {
		word = strings.ToLower(word)
		counter[word]++
	}
	return counter
}

type Words struct {
	str map[int]string
}

func MakeWords(s string) Words {
	words := strings.Fields(s)
	wordsMap := make(map[int]string, len(words))

	for key, word := range words {
		wordsMap[key] = word
	}

	return Words{wordsMap}
}

func (w Words) Index(word string) int {
	var keys []int
	for key, val := range w.str {
		if val == word {
			keys = append(keys, key)
		}
	}

	length := len(keys)
	if length > 0 {
		sort.Ints(keys)
		return keys[0]
	}

	return -1
}

/*// Words represents words in a string.
type Words struct {
	str string
}

// MakeWords creates words from a string.
func MakeWords(s string) Words {
	return Words{s}
}

// Index returns the index of the first instance of word in words,
// or -1 if word is not present in words.
func (w Words) Index(word string) int {
	words := strings.Fields(w.str)
	for idx, item := range words {
		if item == word {
			return idx
		}
	}
	return -1
}*/
