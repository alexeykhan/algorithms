package main

import (
	"bytes"
	"fmt"
	"math"
)

type linkedList struct {
	value int
	next  *linkedList
}

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

func AddTwoNumbersAsLinkedLists(l1, l2 *linkedList) int {
	var number1, number2, sum, digit int
	var iterator *linkedList
	var next bool

	iterator = l1
	for {
		if next {
			number2 += iterator.value * int(math.Pow10(digit))
		} else {
			number1 += iterator.value * int(math.Pow10(digit))
		}

		if iterator.next != nil {
			iterator = iterator.next
			digit++
		} else if !next {
			iterator = l2
			next = true
			digit = 0
		} else {
			digit = 0
			break
		}
	}

	sum = number1 + number2

	// Осталось перевернуть сумму и вывести ее как связанный список
	// + написать тесты

	return sum
}

func main() {
	l1 := &linkedList{
		value: 2,
		next: &linkedList{
			value: 4,
			next: &linkedList{
				value: 3,
				next:  nil,
			},
		},
	}
	l2 := &linkedList{
		value: 5,
		next: &linkedList{
			value: 6,
			next: &linkedList{
				value: 4,
				next:  nil,
			},
		},
	}
	fmt.Println(AddTwoNumbersAsLinkedLists(l1, l2))
}
