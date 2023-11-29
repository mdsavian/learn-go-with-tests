package iteration

import "strings"

func Repeat(character string, times int, useLibString bool) string {

	if useLibString {
		return strings.Repeat(character, times)
	}

	var repeated string

	for i:= 0; i<times; i++ {
		repeated += character
	}

	return repeated
}