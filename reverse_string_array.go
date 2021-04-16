package main

import (
	"strings"
)

// reverseStringArray will receive an string as a input and return the reversed string.
func reverseStringArray(input string) string {
	values := strings.Fields(input)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}

	return strings.Join(values, " ")
}
