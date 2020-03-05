package main

import (
	"bytes"
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// LongestSubstringWithoutRepeatingCharacters находит в заданной строке
// длину максимальной подстроки, у которой нет повторящихся символов.
func LongestSubstringWithoutRepeatingCharacters(s string) int {
	var stopIndex, maxLength, bufferIndex int
	var buffer []byte
	input := []byte(s)

	for i, char := range input {
		bufferIndex = bytes.IndexByte(buffer, char)
		if bufferIndex > -1 {
			maxLength = max(maxLength, i-stopIndex)
			stopIndex += bufferIndex + 1
		}
		buffer = input[stopIndex : i+1]
	}

	if stopIndex == 0 {
		maxLength = len(input)
	} else {
		maxLength = max(maxLength, len(buffer))
	}

	return maxLength
}

func main() {
	fmt.Println("hello")
}
