package main

import "testing"

func TestFirstLastIndices(t *testing.T) {
	type testCase struct {
		target int
		first  int
		last   int
		array  []int
	}
	tests := []testCase{
		testCase{
			target: 9,
			first:  6,
			last:   8,
			array:  []int{1, 3, 3, 5, 7, 8, 9, 9, 9, 15},
		},
		testCase{
			target: 150,
			first:  1,
			last:   2,
			array:  []int{100, 150, 150, 153},
		},
		testCase{
			target: 150,
			first:  -1,
			last:   -1,
			array:  []int{1, 2, 3, 4, 5, 6, 10},
		},
		testCase{
			target: 2,
			first:  1,
			last:   4,
			array:  []int{1, 2, 2, 2, 2, 3, 4, 7, 8, 8},
		},
	}

	for _, tc := range tests {
		first, last := FirstLastIndices(tc.target, tc.array)
		if first != tc.first || last != tc.last {
			t.Errorf("Error! Got: [%d, %d], wanted: [%d, %d]", first, last, tc.first, tc.last)
		}
	}
}

func TestFindDuplicateNumber(t *testing.T) {
	type testCase struct {
		input  []int
		output int
	}
	tests := []testCase{
		testCase{
			input:  []int{1, 3, 4, 2, 2},
			output: 2,
		},
		testCase{
			input:  []int{3, 1, 3, 4, 2},
			output: 3,
		},
		testCase{
			input:  []int{1, 2, 3, 4, 5, 6, 3},
			output: 3,
		},
	}
	for _, tc := range tests {
		got := FindDuplicateNumber(tc.input)
		if got != tc.output {
			t.Errorf("Error! Got: %d, wanted: %d", got, tc.output)
		}
	}
}
