package duplicateIndex_test

import (
	"main/duplicateIndex"
	"testing"
)

func TestFirstDuplicateIdx(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Duplicate at index 3", []int{5, 17, 3, 17, 4, -1}, 3},
		{"No duplicates", []int{1, 2, 3, 4, 5}, -1},
		{"Duplicate at index 3", []int{9, 7, 5, 7, 8, 9}, 3},
		{"First element duplicate", []int{1, 1, 2, 3, 4, 5}, 1},
		{"All unique elements", []int{10, 20, 30, 40, 50}, -1},
		{"Duplicate at index 5", []int{4, 3, 5, 6, 7, 3}, 5},
		{"Single element", []int{1}, -1},
		{"Empty array", []int{}, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := duplicateIndex.FirstDuplicateIdx(test.input)
			if result != test.expected {
				t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}
