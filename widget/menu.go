package klib

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ReplaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

func TruncateStr(str string, num int) string {
	s := []rune(str)
	if len(s) > num {
		s := []rune(str)

		return string(s[0:num])
	}

	return str
}

func TrimStr(str string, start int) string {
	s := []rune(str)
	if len(s) > start {
		s := []rune(str)

		return string(s[start:])
	}

	return ""
}

func FuzzySign(text string, n int) string {
	return TrimStr(text, n) + TruncateStr(text, n)
}
