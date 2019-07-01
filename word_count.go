package main

import (
	"regexp"
	"strings"
)

// P = punctuation, Sm = math symbols, Sc = currency symbols
// C = other (invisible control characters and unused code points)
var replaceSpecialCharsRegex = regexp.MustCompile("[\\p{P}\\p{Sm}\\p{Sc}]")

var whiteSpaceRegex = regexp.MustCompile("\\s+")

// Returns a copy of text with all punctuation, maths, currency, and control
// characters symbols replaced with replacement.
func ReplaceSpecialCharacters(text string, replacement string) string {
	return replaceSpecialCharsRegex.ReplaceAllString(text, replacement)
}

func SplitByWhiteSpace(text string) []string {
	return whiteSpaceRegex.Split(strings.Trim(text, " \n\r\t"), -1)
}

func CountWords(text string) map[string]int {
	wordCount := make(map[string]int)
	cleanedText := ReplaceSpecialCharacters(text, " ")
	words := SplitByWhiteSpace(cleanedText)

	for _, word := range words {
		_, ok := wordCount[word]
		if ok {
			wordCount[word]++
		} else {
			wordCount[word] = 1
		}
	}
	return wordCount
}
