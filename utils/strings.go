package utils

import (
	"strings"
)

// StringSplit2 splits a string into two strings using a seperator.
func StringSplit2(s, sep string) (head, tail string) {
	index := strings.Index(s, sep)
	if index < 0 {
		return s, ""
	}
	return string(s[:index]), string(s[index+1:])
}
