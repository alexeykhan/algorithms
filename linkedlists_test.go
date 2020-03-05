package main

import "testing"

func TestSumNumbersAsLinkedLists(t *testing.T) {
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

	got := SumNumbersAsLinkedLists(l1, l2).Flatten()
	wanted := "(7 -> 0 -> 8)"
	if got != wanted {
		t.Errorf("Error! Got: %v, wanted: %v", got, wanted)
	}
}
