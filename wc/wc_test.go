package wc

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkWordCountRegexp(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Regexp-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountAllocate(phrase)
			}
		})
	}
}

func randomPhrase(n int) string {
	words := make([]string, n)
	for i := range words {
		words[i] = randomWord(3)
	}
	return strings.Join(words, " ")
}

func randomWord(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}
