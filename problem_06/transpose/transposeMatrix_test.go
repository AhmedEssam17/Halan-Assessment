package transpose_test

import (
	"main/transpose"
	"reflect"
	"testing"
)

func TestTransposeMatrix(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "2x2 Matrix",
			input:    [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{1, 3}, {2, 4}},
		},
		{
			name:     "2x4 Matrix",
			input:    [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
			expected: [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}},
		},
		{
			name:     "1x4 Matrix",
			input:    [][]int{{1, 2, 3, 4}},
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:     "4x1 Matrix",
			input:    [][]int{{1}, {2}, {3}, {4}},
			expected: [][]int{{1, 2, 3, 4}},
		},
		{
			name:     "Empty Matrix",
			input:    [][]int{},
			expected: [][]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := transpose.TransposeMatrix(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
