package utils

import (
	"regexp"
	"strings"
)

//DescInGroup reverse string in group of 2 letters
func DescInGroup(text string) string {
	reg := regexp.MustCompile(`([\d|a-zA-Z]{2})`)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes))
	for i, element := range indexes {
		value := text[laststart:(element[0] + len(element))]
		result[i] = value
		laststart = element[1]
	}
	return strings.Join(reverse(result)[:], "")
}

func reverse(ss []string) []string {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}

	return ss
}
