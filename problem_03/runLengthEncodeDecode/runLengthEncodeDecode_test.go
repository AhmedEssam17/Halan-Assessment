package runLengthEncodeDecode_test

import (
	"main/runLengthEncodeDecode"
	"testing"
)

func TestRunLengthEncode(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"aaaaaaaaaabbbaxxxxyyyzyx", "a10b3a1x4y3z1y1x1"},
		{"abcd", "a1b1c1d1"},
		{"aaa", "a3"},
		{"", ""},
		{"x", "x1"},
		{"zzzzzzzzzzz", "z11"},
		{"aabcccccaaa", "a2b1c5a3"},
	}

	for _, tc := range testCases {
		got := runLengthEncodeDecode.RunLengthEncode(tc.input)
		if got != tc.expected {
			t.Errorf("RunLengthEncode(%q) = %q; want %q", tc.input, got, tc.expected)
		}
	}
}

func TestRunLengthDecode(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"a10b3a1x4y3z1y1x1", "aaaaaaaaaabbbaxxxxyyyzyx"},
		{"a1b1c1d1", "abcd"},
		{"a3", "aaa"},
		{"", ""},
		{"x1", "x"},
		{"z11", "zzzzzzzzzzz"},
		{"a2b1c5a3", "aabcccccaaa"},
	}

	for _, tc := range testCases {
		got := runLengthEncodeDecode.RunLengthDecode(tc.input)
		if got != tc.expected {
			t.Errorf("RunLengthDecode(%q) = %q; want %q", tc.input, got, tc.expected)
		}
	}
}

func TestRunLengthEncode_InvalidInput(t *testing.T) {
	invalidInputs := []string{
		"a1b2c3d4", // Should not allow numbers in input
		"hello123", // Mixed letters and numbers
		"9999",     // Only numbers
	}

	for _, input := range invalidInputs {
		got := runLengthEncodeDecode.RunLengthEncode(input)
		if got != "" {
			t.Errorf("Expected empty string for invalid input %q, got %q", input, got)
		}
	}
}

func TestRunLengthDecode_InvalidCases(t *testing.T) {
	invalidCases := []string{
		"abc",    // No numbers
		"a-5b10", // Invalid negative number
		"z10x",   // No count for 'x'
		"p5q3a",  // 'a' without count
	}

	for _, input := range invalidCases {
		got := runLengthEncodeDecode.RunLengthDecode(input)
		if got != "" {
			t.Errorf("Expected empty string for invalid input %q, got %q", input, got)
		}
	}
}
