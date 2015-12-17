package stringutils

import (
	"regexp"
	"unicode"
	"unicode/utf8"
)

// Matches checks whether a textual regular expression matches a string. More complicated queries need to use regex.Compile
// and the full Regexp interface.
func Matches(pattern string, s string) (matched bool, err error) {
	return regexp.MatchString(pattern, s)
}

// Reverse reverses a string but does not preserve combining characters. This function interprets its argument as UTF-8
// and ignores bytes that do not form valid UTF-8. The return value is UTF-8.
func Reverse(utf8Str string) string {
	runes := make([]rune, len(utf8Str))
	start := len(utf8Str)
	for _, rune := range utf8Str {
		// quietly skip invalid UTF-8
		if rune != utf8.RuneError {
			start--
			runes[start] = rune
		}
	}
	return string(runes[start:])
}

// ReverseCC reverses a string while preserving combining characters. Interprets its argument as UTF-8 and ignores bytes
// that do not form valid UTF-8. The return value is UTF-8. This version of function is about three times slower than Reverse.
// You should only use this function if you require preservation of combining characters.
func ReverseCC(utf8Str string) string {
	if utf8Str == "" {
		return ""
	}
	p := []rune(utf8Str)
	r := make([]rune, len(p))
	start := len(r)
	for i := 0; i < len(p); {
		// quietly skip invalid UTF-8
		if p[i] == utf8.RuneError {
			i++
			continue
		}
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) || unicode.Is(unicode.Me, p[j]) || unicode.Is(unicode.Mc, p[j])) {
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}
		i = j
	}
	return (string(r[start:]))
}
