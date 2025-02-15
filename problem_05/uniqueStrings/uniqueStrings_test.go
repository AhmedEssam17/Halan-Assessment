package uniqueStrings_test

import (
	"main/uniqueStrings"
	"reflect"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"apples", "oranges", "flowers", "apples"}, []string{"oranges", "flowers"}},
		{[]string{"apples", "apples"}, []string{}},
		{[]string{"apple", "banana", "cherry"}, []string{"apple", "banana", "cherry"}},
		{[]string{"Apple", "apple", "Banana", "banana"}, []string{}}, // Case insensitive
		{[]string{}, []string{}}, // Empty array
	}

	for _, tc := range testCases {
		got := uniqueStrings.UniqueStrings(tc.input)
		if len(got) == 0 && len(tc.expected) == 0 {
			continue // Test case passes if both are empty slices
		}
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("UniqueStrings(%v) = %v; want %v", tc.input, got, tc.expected)
		}
	}
}
