package palindrome_test

import (
	"main/palindrome"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"rotator", true},
		{"Level", true},    // Case insensitive
		{"race car", true}, // Ignores spaces
		{"apple", false},   // Not a palindrome
		{"Essam", false},   // Not a palindrome
		{"Anna", true},     // Case insensitive
		{"a", true},        // Single character
		{"", true},         // Empty string
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := palindrome.IsPalindrome(test.input)
			if result != test.expected {
				t.Errorf("IsPalindrome(%q) = %t; expected %t", test.input, result, test.expected)
			}
		})
	}
}
