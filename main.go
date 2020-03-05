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

// Flatten выводит строку с последовательностью значений элементов списка,
// разделенных символом стрелки, например: (7 -> 0 -> 8)
func (l *LinkedList) Flatten() string {
	var answer string
	var next bool

	iterator := l
	for {
		if next {
			answer = fmt.Sprintf("%s -> %d", answer, iterator.Value)
		} else {
			answer = fmt.Sprintf("%d", iterator.Value)
			next = true
		}
		if iterator.Next != nil {
			iterator = iterator.Next
		} else {
			break
		}
	}

	return fmt.Sprintf("(%s)", answer)
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
func SumNumbersAsLinkedLists(list ...*LinkedList) *LinkedList {
	var digit int
	var total float64
	for _, iterator := range list {
		for {
			total += float64(iterator.Value) * math.Pow10(digit)
			if iterator.Next != nil {
				iterator = iterator.Next
				digit++
			} else {
				digit = 0
				break
			}
		}
	}

	var answer, current *LinkedList
	var value float64

	for total > 0 {
		value = math.Mod(total, 10)
		total = (total - value) / 10

		if answer == nil {
			answer = &LinkedList{
				Value: int(value),
				Next:  &LinkedList{},
			}
			current = answer.Next
			fmt.Println(total, value, answer)
			continue
		}

		current.Value = int(value)
		if total > 0 {
			current.Next = &LinkedList{}
			current = current.Next
		} else {
			current.Next = nil
		}
	}

	return answer
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

	sum := SumNumbersAsLinkedLists(l1, l2)
	fmt.Println(sum.Flatten())
}
