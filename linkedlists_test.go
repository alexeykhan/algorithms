package main

import (
	"testing"
)

func TestReverseLinkedListIteratively(t *testing.T) {
	type testCase struct {
		input  *LinkedList
		output string
	}
	tests := []testCase{
		testCase{
			input: &LinkedList{
				Value: 2,
				Next: &LinkedList{
					Value: 4,
					Next: &LinkedList{
						Value: 3,
						Next:  nil,
					},
				},
			},
			output: "(3 -> 4 -> 2)",
		},
		testCase{
			input: &LinkedList{
				Value: 5,
				Next: &LinkedList{
					Value: 1,
					Next: &LinkedList{
						Value: 8,
						Next:  nil,
					},
				},
			},
			output: "(8 -> 1 -> 5)",
		},
		testCase{
			input: &LinkedList{
				Value: 9,
				Next: &LinkedList{
					Value: 2,
					Next:  nil,
				},
			},
			output: "(2 -> 9)",
		},
		testCase{
			input: &LinkedList{
				Value: 7,
				Next:  nil,
			},
			output: "(7)",
		},
	}

	for _, tc := range tests {
		got := ReverseLinkedListIteratively(tc.input).Flatten()
		if got != tc.output {
			t.Errorf("Error! Got: %v, wanted: %v", got, tc.output)
		}
	}
}

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
