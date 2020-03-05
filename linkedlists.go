package main

import (
	"fmt"
	"math"
)

// LinkedList allows to store connected data as a linked list.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// Flatten returns a string with a sequence of linked list elements values
// separated by arrow, for example: (7 -> 0 -> 8)
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

// SumNumbersAsLinkedLists calculates the sum of entered numbers (represented by
// linked lists with digits in reverse order) and returns it as a linked list
// with digits in reverse order.
// Example: (2 -> 4 -> 3) + (5 -> 6 -> 4) = 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
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
