package main

import (
	"fmt"
	"math"
)

// LinkedList позволяет хранить данные в формате связанного списка.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// Flatten выводит строку с последовательностью значений элементов связанного
// списка, разделенных символом стрелки, например: (7 -> 0 -> 8)
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

// ReverseLinkedListIteratively переворачивает связанный список сверху вниз с
// помощью последовательных операций.
func ReverseLinkedListIteratively(l *LinkedList) *LinkedList {
	var next bool
	var reversed *LinkedList
	for {
		if next {
			reversed = &LinkedList{
				Value: l.Value,
				Next:  reversed,
			}
		} else {
			reversed = &LinkedList{
				Value: l.Value,
				Next:  nil,
			}
			next = true
		}
		if l.Next != nil {
			l = l.Next
		} else {
			break
		}
	}
	return reversed
}

// SumNumbersAsLinkedLists считает сумму всех введенных чисел в виде связанных
// списков из цифр в обратном порядке, выводит ее в виде такого же связанного
// списка с цифрами в обратном порядке.
// Пример: (2 -> 4 -> 3) + (5 -> 6 -> 4) = 7 -> 0 -> 8
// Пояснение: 342 + 465 = 807.
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
