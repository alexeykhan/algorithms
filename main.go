package main

import (
	"bytes"
	"fmt"
	"math"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
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

// SumNumbersAsLinkedLists считает сумму всех введенных чисел в виде связанных
// списков из цифр в обратном порядке, выводит ее в виде такого же связанного
// списка с цифрами в обратном порядке.
func SumNumbersAsLinkedLists(list ...*LinkedList) int {
	var sum, digit int
	for _, iterator := range list {
		for {
			sum += iterator.Value * int(math.Pow10(digit))
			if iterator.Next != nil {
				iterator = iterator.Next
				digit++
			} else {
				digit = 0
				break
			}
		}
	}

	// Осталось перевернуть сумму и вывести ее как связанный список
	// + написать тесты

	return sum
}

func main() {
	l1 := &LinkedList{
		Value: 2,
		Next: &LinkedList{
			Value: 4,
			Next: &LinkedList{
				Value: 3,
				Next:  nil,
			},
		},
	}
	l2 := &LinkedList{
		Value: 5,
		Next: &LinkedList{
			Value: 6,
			Next: &LinkedList{
				Value: 4,
				Next:  nil,
			},
		},
	}
	fmt.Println(SumNumbersAsLinkedLists(l1, l2))
}
