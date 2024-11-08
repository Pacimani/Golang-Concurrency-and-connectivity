package nlp

import (
	"regexp"
	"strings"
)

var (
	wordRe = regexp.MustCompile(`\w+`)
)

// Tokenize returns slice of tokens found in text (lower case)

func Tokenize(text string) []string {
	return wordRe.FindAllString(strings.ToLower(text), -1)
}
