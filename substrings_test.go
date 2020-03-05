package main

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	var result int
	tests := map[string]int{
		"A":                   1,
		"AA":                  1,
		"AB":                  2,
		"ABC":                 3,
		"ABCDAFGAHI":          6,
		"ABCADGEAFGHAHIJK":    6,
		"abrkaabcdefghijjxxx": 10,
	}

	for input, answer := range tests {
		result = LongestSubstringWithoutRepeatingCharacters(input)
		if result != answer {
			t.Errorf("Error! Got: %v, wanted: %v", result, answer)
		}
	}
}
