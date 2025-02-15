package tree_test

import (
	"main/tree"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		root     *tree.Node
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single Node",
			root:     &tree.Node{Value: 10, Children: []*tree.Node{}},
			expected: 10,
		},
		{
			name: "Multiple levels",
			root: &tree.Node{Value: 10, Children: []*tree.Node{
				{Value: 5, Children: []*tree.Node{
					{Value: 2, Children: []*tree.Node{}},
					{Value: 1, Children: []*tree.Node{}},
				}},
				{Value: 3, Children: []*tree.Node{
					{Value: 4, Children: []*tree.Node{}},
				}},
			}},
			expected: 10 + 5 + 2 + 1 + 3 + 4, // 25
		},
		{
			name: "Deeply nested tree",
			root: &tree.Node{Value: 1, Children: []*tree.Node{
				{Value: 2, Children: []*tree.Node{
					{Value: 3, Children: []*tree.Node{
						{Value: 4, Children: []*tree.Node{
							{Value: 5, Children: []*tree.Node{}},
						}},
					}},
				}},
			}},
			expected: 1 + 2 + 3 + 4 + 5, // 15
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := tree.Sum(test.root)
			if result != test.expected {
				t.Errorf("For test '%s', expected %d but got %d", test.name, test.expected, result)
			}
		})
	}
}
